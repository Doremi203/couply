package main

import (
	"context"
	"github.com/Doremi203/couply/backend/auth/pkg/token"
	"google.golang.org/grpc/metadata"
	"strings"
	"sync"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/auth/pkg/webapp"
	blocker_service "github.com/Doremi203/couply/backend/blocker/internal/app/blocker-service"
	telegram_client "github.com/Doremi203/couply/backend/blocker/internal/client/telegram"
	"github.com/Doremi203/couply/backend/blocker/internal/client/user"
	blocker_usecase "github.com/Doremi203/couply/backend/blocker/internal/usecase/blocker-service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type connectionManager struct {
	mu          sync.Mutex
	connections []*grpc.ClientConn
}

func (cm *connectionManager) Add(conn *grpc.ClientConn) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.connections = append(cm.connections, conn)
}

func (cm *connectionManager) CloseAll() {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	for _, conn := range cm.connections {
		if conn != nil {
			_ = conn.Close()
		}
	}
}

func main() {
	cm := &connectionManager{}
	defer cm.CloseAll()

	webapp.Run(func(ctx context.Context, app *webapp.App) error {
		app.AddGatewayOptions(
			runtime.WithIncomingHeaderMatcher(func(s string) (string, bool) {
				switch s = strings.ToLower(s); s {
				case "x-api-token", "user-token":
					return s, true
				default:
					return runtime.DefaultHeaderMatcher(s)
				}
			}),
		)

		var telegramConfig struct {
			Token       string `secret:"telegram-token"`
			AdminChatID int64  `secret:"telegram-admin-chat-id"`
		}
		err := app.Config.ReadSection("telegram", &telegramConfig)
		if err != nil {
			return errors.WrapFail(err, "read telegram config")
		}

		tokenConfig := token.Config{}
		err = app.Config.ReadSection("user-token", &tokenConfig)
		if err != nil {
			return err
		}

		var userServiceConfig struct {
			Address string `yaml:"address"`
		}
		if err := app.Config.ReadSection("user_service", &userServiceConfig); err != nil {
			return errors.WrapFail(err, "read user service config")
		}

		userServiceClient, conn, err := user.NewClient(userServiceConfig.Address)
		if err != nil {
			return errors.WrapFail(err, "create user service client")
		}
		cm.Add(conn)

		bot, err := telegram_client.NewBotClient(telegramConfig.Token, telegramConfig.AdminChatID)
		if err != nil {
			return errors.WrapFail(err, "create telegram bot")
		}

		blockUseCase := blocker_usecase.NewUseCase(
			userServiceClient,
			bot,
			app.Log,
		)

		blockService := blocker_service.NewImplementation(
			blockUseCase,
			app.Log,
		)

		app.RegisterGRPCServices(
			blockService,
		)

		app.AddGatewayHandlers(
			blockService,
		)

		go bot.StartCallbackHandler(func(userID string) {
			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

			token, found := blockUseCase.GetCache().Get(userID)
			if !found {
				app.Log.Infof("token not found for user %s", userID)
				return
			}

			md := metadata.New(map[string]string{"user-token": token.(string)})
			ctx = metadata.NewIncomingContext(ctx, md)

			if err := userServiceClient.UpdateUserV1(ctx, true); err != nil {
				app.Log.Infof("failed to block user %s with error %v", userID, err)
			} else {
				app.Log.Infof("successfully blocked user %s", userID)
			}
		})

		return nil
	})
}

package blocker_service

import (
	"context"
	"testing"
	"time"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"github.com/Doremi203/couply/backend/blocker/internal/domain/blocker"
	"github.com/Doremi203/couply/backend/blocker/internal/dto"
	mock_blocker_service "github.com/Doremi203/couply/backend/blocker/internal/mocks/usecase/blocker"
	user_service "github.com/Doremi203/couply/backend/matcher/gen/api/user-service/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type loggerStub struct{}

func (l *loggerStub) Infof(_ string, _ ...any) {}

func (l *loggerStub) Error(_ error) {}

func (l *loggerStub) Warn(_ error) {}

func TestUseCase_ReportUser(t *testing.T) {
	t.Parallel()

	now := time.Now()

	type mocks struct {
		userServiceClient    *mock_blocker_service.MockuserClient
		bot                  *mock_blocker_service.MockbotClient
		blockerStorageFacade *mock_blocker_service.MockblockerStorageFacade
	}
	type args struct {
		in *dto.ReportUserV1Request
	}
	tests := []struct {
		name    string
		setup   func(mocks)
		args    args
		want    *dto.ReportUserV1Response
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "tx error",
			setup: func(m mocks) {
				m.userServiceClient.EXPECT().GetUserByIDV1(gomock.Any(), "11111111-1111-1111-1111-111111111111").
					Return(&user_service.User{
						Id: "11111111-1111-1111-1111-111111111111",
					}, nil)

				m.bot.EXPECT().SendReportMessage(
					&user_service.User{
						Id: "11111111-1111-1111-1111-111111111111",
					},
					[]blocker.ReportReason{
						blocker.ReportReasonSpam,
					},
					"",
					uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				).Return(nil)

				m.blockerStorageFacade.EXPECT().ReportUserTx(gomock.Any(),
					&blocker.UserBlock{
						ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						BlockedID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						Message:   "",
						Reasons: []blocker.ReportReason{
							blocker.ReportReasonSpam,
						},
						Status:    blocker.BlockStatusPending,
						CreatedAt: now,
					},
				).Return(blocker.ErrDuplicateUserBlock)
			},
			args: args{
				in: &dto.ReportUserV1Request{
					BlockID:      uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					TargetUserID: "11111111-1111-1111-1111-111111111111",
					ReportReasons: []blocker.ReportReason{
						blocker.ReportReasonSpam,
					},
					Message:   "",
					CreatedAt: now,
				},
			},
			want: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, blocker.ErrDuplicateUserBlock)
			},
		},
		{
			name: "send message error",
			setup: func(m mocks) {
				m.userServiceClient.EXPECT().GetUserByIDV1(gomock.Any(), "11111111-1111-1111-1111-111111111111").
					Return(&user_service.User{
						Id: "11111111-1111-1111-1111-111111111111",
					}, nil)

				m.bot.EXPECT().SendReportMessage(
					&user_service.User{
						Id: "11111111-1111-1111-1111-111111111111",
					},
					[]blocker.ReportReason{
						blocker.ReportReasonSpam,
					},
					"",
					uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				).Return(errors.Error("some error"))
			},
			args: args{
				in: &dto.ReportUserV1Request{
					BlockID:      uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					TargetUserID: "11111111-1111-1111-1111-111111111111",
					ReportReasons: []blocker.ReportReason{
						blocker.ReportReasonSpam,
					},
					Message:   "",
					CreatedAt: now,
				},
			},
			want: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return true
			},
		},
		{
			name: "dto error",
			setup: func(m mocks) {
				m.userServiceClient.EXPECT().GetUserByIDV1(gomock.Any(), "1fwefewfrefr").
					Return(&user_service.User{
						Id: "1fwefewfrefr",
					}, nil)
			},
			args: args{
				in: &dto.ReportUserV1Request{
					BlockID:      uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					TargetUserID: "1fwefewfrefr",
					ReportReasons: []blocker.ReportReason{
						blocker.ReportReasonSpam,
					},
					Message:   "",
					CreatedAt: now,
				},
			},
			want: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return true
			},
		},
		{
			name: "get user err",
			setup: func(m mocks) {
				m.userServiceClient.EXPECT().GetUserByIDV1(gomock.Any(), "11111111-1111-1111-1111-111111111111").
					Return(nil, errors.Error("some error"))
			},
			args: args{
				in: &dto.ReportUserV1Request{
					BlockID:      uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					TargetUserID: "11111111-1111-1111-1111-111111111111",
					ReportReasons: []blocker.ReportReason{
						blocker.ReportReasonSpam,
					},
					Message:   "",
					CreatedAt: now,
				},
			},
			want: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return true
			},
		},
		{
			name: "success",
			setup: func(m mocks) {
				m.userServiceClient.EXPECT().GetUserByIDV1(gomock.Any(), "11111111-1111-1111-1111-111111111111").
					Return(&user_service.User{
						Id: "11111111-1111-1111-1111-111111111111",
					}, nil)

				m.bot.EXPECT().SendReportMessage(
					&user_service.User{
						Id: "11111111-1111-1111-1111-111111111111",
					},
					[]blocker.ReportReason{
						blocker.ReportReasonSpam,
					},
					"",
					uuid.MustParse("11111111-1111-1111-1111-111111111111"),
				).Return(nil)

				m.blockerStorageFacade.EXPECT().ReportUserTx(gomock.Any(),
					&blocker.UserBlock{
						ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						BlockedID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						Message:   "",
						Reasons: []blocker.ReportReason{
							blocker.ReportReasonSpam,
						},
						Status:    blocker.BlockStatusPending,
						CreatedAt: now,
					},
				).Return(nil)
			},
			args: args{
				in: &dto.ReportUserV1Request{
					BlockID:      uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					TargetUserID: "11111111-1111-1111-1111-111111111111",
					ReportReasons: []blocker.ReportReason{
						blocker.ReportReasonSpam,
					},
					Message:   "",
					CreatedAt: now,
				},
			},
			want:    &dto.ReportUserV1Response{},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)

			mocks := mocks{
				userServiceClient:    mock_blocker_service.NewMockuserClient(ctrl),
				bot:                  mock_blocker_service.NewMockbotClient(ctrl),
				blockerStorageFacade: mock_blocker_service.NewMockblockerStorageFacade(ctrl),
			}

			if tt.setup != nil {
				tt.setup(mocks)
			}

			logger := &loggerStub{}

			usecase := NewUseCase(
				mocks.userServiceClient,
				mocks.bot,
				mocks.blockerStorageFacade,
				logger,
			)

			got, err := usecase.ReportUser(context.Background(), tt.args.in)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

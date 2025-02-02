package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"sync"
)

type ResourceCloser func() error

type BackgroundProcessor func(ctx context.Context) error

type Context struct {
	wg sync.WaitGroup

	log     *slog.Logger
	closers []ResourceCloser

	grpcServer *grpcServer
	httpServer *http.Server
	router     *gin.Engine

	backgroundCtx        context.Context
	backgroundStopper    context.CancelCauseFunc
	backgroundProcessors []BackgroundProcessor

	readinessRegistered bool
}

func newContext(
	log *slog.Logger,
) *Context {
	backgroundCtx, backgroundStopper := context.WithCancelCause(context.Background())
	router := newGinRouter(log)

	return &Context{
		log:               log,
		backgroundCtx:     backgroundCtx,
		backgroundStopper: backgroundStopper,
		grpcServer:        newGRPCServer(log),
		router:            router,
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", httpConfig.Port),
			Handler: router,
		},
	}
}

func (c *Context) AddCloser(closer ResourceCloser) {
	c.closers = append(c.closers, closer)
}

func (c *Context) AddBackgroundProcessor(processor BackgroundProcessor) {
	c.backgroundProcessors = append(c.backgroundProcessors, processor)
}

func (c *Context) GRPCServer() *grpcServer {
	return c.grpcServer
}

func (c *Context) HTTPRouter() *gin.Engine {
	return c.router
}

func (c *Context) AddReadinessCheck(readiness func() bool) error {
	if readiness == nil {
		return errors.New("readiness function is required")
	}
	if c.readinessRegistered {
		return errors.New("readiness check is already registered")
	}

	c.readinessRegistered = true
	c.router.GET("/readiness", func(c *gin.Context) {
		if !readiness() {
			c.AbortWithStatus(http.StatusServiceUnavailable)
			return
		}

		c.Status(http.StatusOK)
	})

	return nil
}

func (c *Context) close() {
	c.log.Info("closing app context")

	c.stopBackgroundProcessors()
	c.closeResources()

	c.log.Info("app context closed")
}

func (c *Context) closeResources() {
	c.log.Info("closing resources")

	for i := len(c.closers) - 1; i >= 0; i-- {
		err := c.closers[i]()
		if err != nil {
			c.log.Error("failed to close resource", "err", err)
		}
	}

	c.log.Info("resources closed")
}

func (c *Context) stopBackgroundProcessors() {
	c.log.Info("stopping background processors")

	c.backgroundStopper(errors.New("background processors stopped"))
	c.wg.Wait()

	c.log.Info("background processors stopped")
}

func (c *Context) startGrpcServer() {

}

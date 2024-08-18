package httpserver

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func (srv HTTPServer) Run() error {
	err := srv.mapHandlers()
	if err != nil {
		return err
	}

	ctx := context.Background()
	go func() {
		srv.gin.Run(fmt.Sprintf(":%d", srv.port))
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	srv.l.Infof(ctx, "Started server on :%d", srv.port)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	srv.l.Info(ctx, <-ch)
	srv.l.Info(ctx, "Stopping API server.")

	return nil
}

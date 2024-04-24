package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"

	"git.corp.zgcszkw.com/authhub/config"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().Msgf("cannot load config: %s", err)
	}

	gin.SetMode(gin.ReleaseMode)
	// 开发环境日志输出
	if config.Environment == "local" || config.Environment == "development" {
		gin.SetMode(gin.DebugMode)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	waitGroup, ctx := errgroup.WithContext(ctx)
	runGinServer(ctx, &config, waitGroup)
	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("connot start server")
	}
}

func runGinServer(
	ctx context.Context,
	config *config.Config,
	waitGroup *errgroup.Group,
) {

	server, err := initApp(ctx, config)
	// server, err := server.NewServer(config, service, logger)
	if err != nil {
		log.Fatal().Msgf("cannot new server: %s", err)
	}
	waitGroup.Go(func() error {
		log.Info().Msgf("start http server at %s", config.HTTPServerAddress)

		if err := server.Start(config.HTTPServerAddress); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}

			log.Error().Err(err).Msg("cannot start http server")
			return err
		}
		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done() // 从通道中读出信号，表示已经关闭

		log.Info().Msg("graceful shutdown gRPC server")
		_ = server.GracefulShutdown(ctx)
		log.Info().Msg("gRPC server is stopped")

		return nil
	})
}

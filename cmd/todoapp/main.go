package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/skankhunter/todo-go/internal/core/logger"
	core_pgx_pool "github.com/skankhunter/todo-go/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/skankhunter/todo-go/internal/core/transport/http/middleware"
	core_http_server "github.com/skankhunter/todo-go/internal/core/transport/http/server"
	users_postgres_repository "github.com/skankhunter/todo-go/internal/features/users/repository/postgres"
	users_service "github.com/skankhunter/todo-go/internal/features/users/service"
	users_transport_http "github.com/skankhunter/todo-go/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGALRM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())

	if err != nil {
		fmt.Println("Failed to init app logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("Init postgres connection pool")

	pool, err := core_pgx_pool.NewPool(
		ctx,
		core_pgx_pool.NewConfigMust(),
	)

	if err != nil {
		logger.Fatal("Failed to init postgres connection pool", zap.Error(err))
	}

	defer pool.Close()

	logger.Debug("Init feature", zap.String("feature", "users"))

	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransportHTTP := users_transport_http.NewUsersHTTPHanlder(usersService)

	logger.Debug("Init HTTP server", zap.String("feature", "users"))

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)
	apiVersionRouterV1 := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouterV1.RegisterRoutes(usersTransportHTTP.Routes()...)

	// apiVersionRouterV2 := core_http_server.NewAPIVersionRouter(
	// 	core_http_server.ApiVersion2,
	// 	core_http_middleware.Dummy("API v2 middleware"),
	// )
	// apiVersionRouterV2.RegisterRoutes(usersTransportHTTP.Routes()...)

	httpServer.RegisterAPIRouters(
		apiVersionRouterV1,
		// apiVersionRouterV2,
	)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}

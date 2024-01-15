package main

import (
	"go.uber.org/fx"
	"wallet/db"
	"wallet/handler"
	"wallet/internal/config"
	"wallet/internal/locale"
	"wallet/internal/logger"
	"wallet/server"
	memberService "wallet/service/member"
	transService "wallet/service/transaction"
	walletService "wallet/service/wallet"
	memberStorage "wallet/storage/member"
	transStorage "wallet/storage/transaction"
	walletStorage "wallet/storage/wallet"
)

func main() {
	fx.New(
		fx.Provide(
			postgresDB,
			redisDB,

			// clients
			externalClients,

			// storages
			fx.Annotate(
				memberStorage.NewStorage,
				fx.As(new(memberStorage.Repository)),
			),
			fx.Annotate(
				walletStorage.NewStorage,
				fx.As(new(walletStorage.Repository)),
			),
			fx.Annotate(
				transStorage.NewStorage,
				fx.As(new(transStorage.Repository)),
			),

			// services
			fx.Annotate(
				transService.New,
				fx.As(new(transService.UseCase)),
			),

			fx.Annotate(
				walletService.New,
				fx.As(new(walletService.UseCase)),
			),

			fx.Annotate(
				memberService.New,
				fx.As(new(memberService.UseCase)),
			),

			// handlers
			handler.NewMemberHandler,
			handler.NewWalletHandler,

			// server
			server.NewServer,
		),
		fx.Supply(),
		fx.Invoke(
			config.Init,
			logger.SetupLogger,
			locale.Init,
			db.Migrate,
			setupServer,
			handler.SetupMemberRoutes,
			handler.SetupWalletRoutes,
			server.Run,
		),
	).Run()
}

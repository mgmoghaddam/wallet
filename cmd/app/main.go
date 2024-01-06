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
			memberStorage.New,
			walletStorage.New,
			transStorage.New,

			// services
			transService.New,
			walletService.New,
			memberService.New,

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

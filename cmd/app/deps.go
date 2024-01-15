package main

import (
	"database/sql"
	"log"
	"wallet/client/discount"
	"wallet/db"
	"wallet/internal/config"
	"wallet/server"
)

func postgresDB() *sql.DB {
	psql, err := db.NewPostgres(
		config.DBName(), config.DBUser(), config.DBPassword(), config.DBHost(), config.DBPort(),
		config.DBMaxOpenConn(), config.DBMaxIdleConn(),
	)
	if err != nil {
		log.Fatalf("failed to initalize db: %v", err)
	}
	return psql
}

func redisDB() db.RedisClient {
	rdb, err := db.NewRedis(config.RDBHost(), config.RDBPassword(), config.RDBPort(), config.RDB())
	if err != nil {
		log.Fatalf("failed to initalize redis: %v", err)
	}
	return rdb
}

func externalClients() discount.Client {
	giftRepo := discount.NewHTTPClient(config.APIDiscount())
	return giftRepo
}

func setupServer(s *server.Server, psql *sql.DB) {
	s.SetHealthFunc(healthFunc(psql)).
		SetupRoutes()
}

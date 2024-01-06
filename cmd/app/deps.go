package main

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
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

func redisDB() *redis.Client {
	rdb, err := db.NewRedis(config.RDBHost(), config.RDBPassword(), config.RDBPort(), config.RDB())
	if err != nil {
		log.Fatalf("failed to initalize redis: %v", err)
	}
	return rdb
}

func externalClients() *discount.Client {
	discountClient := discount.New(config.APIDiscount())
	return discountClient
}

func setupServer(s *server.Server, psql *sql.DB) {
	s.SetHealthFunc(healthFunc(psql)).
		SetupRoutes()
}

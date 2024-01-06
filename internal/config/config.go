package config

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"time"
)

func ServerPort() int {
	return viper.GetInt("server.port")
}

func ServerDebug() bool {
	return viper.GetBool("server.debug")
}

func DBName() string {
	return viper.GetString("db.postgres.name")
}

func DBPassword() string {
	return viper.GetString("db.postgres.password")
}

func DBUser() string {
	return viper.GetString("db.postgres.user")
}

func DBPort() string {
	return viper.GetString("db.postgres.port")
}

func DBHost() string {
	return viper.GetString("db.postgres.host")
}

func DBMaxIdleConn() int {
	return viper.GetInt("db.postgres.maxIdleConn")
}

func DBMaxOpenConn() int {
	return viper.GetInt("db.postgres.maxOpenConn")
}

func DBMigrationsPath() string {
	return viper.GetString("db.postgres.migrationsPath")
}

// ---- Redis

func RDBTimeOut() time.Duration {
	return viper.GetDuration("db.redis.timeout")
}

func RDBPrefix() string {
	return viper.GetString("db.redis.prefix")
}

func RDBPassword() string {
	return viper.GetString("db.redis.password")
}

func RDBPort() string {
	return viper.GetString("db.redis.port")
}

func RDBHost() string {
	return viper.GetString("db.redis.host")
}

func RDB() int {
	return viper.GetInt("db.redis.db")
}

// ---- App

func LogLevel() string {
	return viper.GetString("app.log.level")
}

func APIDiscount() string {
	return viper.GetString("api.discount.url")
}

func Init() {
	viper.SetConfigName(getEnv("CONFIG_NAME", "conf"))
	viper.SetConfigType("yaml")              // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./resources/conf")  // optionally look for config in the working directory
	viper.AddConfigPath("../resources/conf") // optionally look for config in the working directory
	err := viper.ReadInConfig()              // Find and read the config file
	if err != nil {                          // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

func getEnv(key, fallback string) string {
	log.Info().Msg("getting environment")
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

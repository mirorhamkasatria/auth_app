package configs

import (
	"os"
	"strconv"
)

type Config struct {
	DbUsername      string
	DbPassword      string
	DbHost          string
	DbName          string
	DbPort          string
	SvPort          string
	MaxConn         int
	MaxIdleConn     int
	MaxLifetimeConn int
}

func LoadConfig() *Config {
	// database
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	//server
	svPort := os.Getenv("SERVER_PORT")

	return &Config{
		DbUsername:      dbUsername,
		DbPassword:      dbPassword,
		DbHost:          dbHost,
		DbName:          dbName,
		DbPort:          dbPort,
		SvPort:          svPort,
		MaxConn:         maxConn,
		MaxIdleConn:     maxIdleConn,
		MaxLifetimeConn: maxLifetimeConn,
	}
}

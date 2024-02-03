package main

import (
	"github.com/auth_app/configs"
	"github.com/auth_app/routes"
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	validation := validator.New()
	cfg := configs.LoadConfig()

	gormDbInit, err := configs.GORMOpenConn(cfg)
	if err != nil {
		panic(err)
	}

	server := routes.NewServer(validation, gormDbInit)
	server.ListenAndServer(cfg.SvPort)
}

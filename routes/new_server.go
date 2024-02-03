package routes

import (
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

type ApiServer struct {
	GormDB    *gorm.DB
	App       *fiber.App
	Validator *validator.Validate
}

func NewServer(validation *validator.Validate, gormDB *gorm.DB) *ApiServer {
	app := fiber.New(fiber.Config{
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		BodyLimit:    100 * 1024 * 1024,
	})

	app.Use(cors.New(cors.ConfigDefault))
	return &ApiServer{
		App:       app,
		Validator: validation,
		GormDB:    gormDB,
	}
}

func (server *ApiServer) ListenAndServer(port string) {
	registerRouters(server)
	// Run server.
	if err := server.App.Listen(port); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}

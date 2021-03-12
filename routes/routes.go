package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//Routes is function all router
func init() {
	var port string = ":8080"
	FiBe := fiber.New()
	FiBe.Use(logger.New())
	FiBe.Use(cors.New())
	UseRoutes(FiBe)
	FiBe.Listen(port)
}

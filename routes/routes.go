package routes

import "github.com/gofiber/fiber/v2"
//Routes is function all router
func init() {
	FiBe:=fiber.New()
	UseRoutes(FiBe)
	FiBe.Listen(":8080")
}
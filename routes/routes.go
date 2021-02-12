package routes

import "github.com/gofiber/fiber/v2"
//Routes is function all router
func init() {
	var port string =":8000"
	FiBe:=fiber.New()
	UseRoutes(FiBe)
	FiBe.Listen(port)
}
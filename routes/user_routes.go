package routes

import (
	"github.com/fluk27/StockMagageSysyem/controllers"
	"github.com/gofiber/fiber/v2"
)

//UseRoutes is function all method router of user
func UseRoutes(routes *fiber.App) {

	uc := controllers.UserContrillers{}
	routes.Post("/Register", uc.Register)
	routes.Post("/Login", uc.Login)
}

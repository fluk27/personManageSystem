package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/fluk27/StockMagageSysyem/controllers"
)

//UseRoutes is function all method router of user
func UseRoutes(routes *fiber.App) {
	
uc:=controllers.UserContrillers{}
	routes.Post("/Login", uc.Login)
}

package controllers

import (
	"log"
	"net/http"

	"github.com/fluk27/StockMagageSysyem/models"
	"github.com/fluk27/StockMagageSysyem/services"
	"github.com/gofiber/fiber/v2"
)

//UserContrillers zip all varible and function
type UserContrillers struct {
}

//Login is functon login(user,pasword)
func (uc *UserContrillers) Login(c *fiber.Ctx) error {
	user := &models.UserModel{}
	us := &services.UserServices{}
	err := c.BodyParser(user)
	if err != nil {
		log.Println("errorUC:",err.Error())
		return c.JSON(err.Error())
	}
	
	re,err:=us.InstertDataUsers("users",user)
if err != nil {
	log.Println("error intertData:",err.Error())
	return c.Status(http.StatusInternalServerError).JSON(err.Error())
}

	return c.Status(http.StatusOK).JSON(re)
}

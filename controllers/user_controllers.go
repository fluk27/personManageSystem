package controllers

import (
	"github.com/fluk27/StockMagageSysyem/services"
	"github.com/gofiber/fiber/v2"
)

//UserContrillers zip all varible and function
type UserContrillers struct {
}
type Student struct {
	Name         string  `json:"name"`
	Age          int64   `json:"age"`
	AverageScore float64 `json:"average_score"`
}

//Login is functon login(user,pasword)
func (uc *UserContrillers) Login(c *fiber.Ctx) error {
	es:=&services.ELKServices{}
	es.GetData()
	return c.SendString("hello login")
}

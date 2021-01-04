package controllers

import "github.com/gofiber/fiber/v2"

//UserContrillers zip all varible and function 
type UserContrillers struct {
}
//Login is functon login(user,pasword)
func (uc *UserContrillers) Login(c *fiber.Ctx)  error{ 
return c.SendString("hello login")
}
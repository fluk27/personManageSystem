package controllers

import (
	"github.com/fluk27/StockMagageSysyem/models"
	"github.com/fluk27/StockMagageSysyem/services"
	outSer "github.com/fluk27/testgo2020/services"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

//UserContrillers zip all varible and function
type UserContrillers struct {
}

//Login is functon login(user,pasword)
func (uc *UserContrillers) Login(c *fiber.Ctx) error {
	user := &models.UserModel{}
	//US := &services.UserServices{}
	err := c.BodyParser(user)
	if err != nil {
		log.Println("errorUC:", err.Error())
		return c.JSON(err.Error())
	}
	log.Println(user.Username)
	if user.Username != "flk12345" || user.Password != "Ws0844038001" {
		return c.Status(http.StatusUnauthorized).JSON("error:'unauthen'")
	}

	// result,err:=US.GetdataUsers("users",nil)
	// if err != nil {
	// 	log.Println("error intertData:", err.Error())
	// 	return c.Status(http.StatusInternalServerError).JSON(err.Error())
	// }

	return c.Status(http.StatusOK).JSON("result")
}

//Register is function insert data to system
func (uc *UserContrillers) Register(c *fiber.Ctx) error {
	UM := &models.UserModel{}

	US := &services.UserServices{}

	// map body to struct
	if err := c.BodyParser(UM); err != nil {
		return err
	}
	// validate struct
	Status, result := US.ValidateStruct(UM)
	if Status != false {
		return c.Status(http.StatusBadRequest).JSON(result)
	}
	query := map[string]string{
		"username": UM.Username,
	}
	resultUser, err := US.GetdataUsers("users", query)
	if err != nil {
		//	log.Fatalln("error getdataUser=",err.Error())
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	//log.Fatalln("result user ELK:",resultUser)
	if len(*resultUser) == 0 {
		// set up RSA key
		encpt := &outSer.RSAKey{
			PathPublicKey:      "./",
			FileNamePublicKey:  "publicKey.pem",
			PathPrivateKey:     "./",
			FileNamePrivateKey: "privateKey.pem"}
			// hash function sha 256
		 result,err:=US.Hashfunction512(UM.Password)
		 if err != nil {
			return c.Status(http.StatusBadRequest).JSON(err.Error())
		 }
		
		UM.Password, err = encpt.EncyptDataWithPKC(UM.Password, result)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(err.Error())
		}
		US.InstertDataUsers("users", UM)
		return c.Status(http.StatusCreated).JSON(UM)
	} else {
		return c.Status(http.StatusConflict).JSON("user duplicate ")
	}

}

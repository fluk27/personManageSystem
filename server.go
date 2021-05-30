package main

import (
	"log"
	"reflect"
// "strings"
	MO "github.com/fluk27/StockMagageSysyem/models"
	//	_ "github.com/fluk27/StockMagageSysyem/routes"
)

func main() {
 readTagInModels()
}

func readTagInModels()  {
	t := reflect.TypeOf(MO.UserModel{})
	
	result:=t.Field(1)
	log.Println("555:",result.Tag.Get("elk"))
	log.Println(t.NumField())

}
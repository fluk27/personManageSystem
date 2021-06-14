package services

import (
	"log"
	"net/http"
	"regexp"
)


type SatrangProServices struct {
}

func (SatrangProServices)GetPriceSatrangProOneItem(nameCoin string) {
	match,err:=regexp.Match("[A-Za-z]*$/g",[] byte (nameCoin))
if err != nil {
	log.Fatalln(err.Error())
}
log.Println("word:",match)
	URL:="https://satangcorp.com/api/orders/?pair="+nameCoin+"_thb"
	http.NewRequest(http.MethodGet,URL,nil)
}
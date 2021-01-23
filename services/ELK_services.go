package services

import (
	"fmt"
	"log"

	elastic "github.com/olivere/elastic/v7"
	// Import the Elasticsearch library packages
)

//ELKServices is stuct
type ELKServices struct {
	UrlELK string
}

//GetESClient is function
func (ELK *ELKServices) initELK() (*elastic.Client, error) {

	client, err := elastic.NewClient(elastic.SetURL(ELK.UrlELK), elastic.SetSniff(false), elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")
	if err != nil {
		log.Fatalln("error init ELK:", err.Error())
	}

	return client, err

}

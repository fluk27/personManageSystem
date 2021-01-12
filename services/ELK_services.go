package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/fluk27/StockMagageSysyem/models"
	elastic "github.com/olivere/elastic/v7"
	// Import the Elasticsearch library packages
)

//ELKServices is stuct
type ELKServices struct {
	UrlELK string
}

//GetESClient is function
func (ELK *ELKServices) initELK() (*elastic.Client, error) {

	client, err := elastic.NewClient(elastic.SetURL(ELK.UrlELK),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")
	if err != nil {
		log.Fatalln("error init ELK:", err.Error())
	}

	return client, err

}

//GetData is function getdata
func (ELK *ELKServices) GetData() {

	ctx := context.Background()
	ELK.UrlELK = "http://localhost:9200"
	esclient, err := ELK.initELK()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}

	//creating student object
	newStudent := &models.Student{
		Name:         "Gopher doe",
		Age:          10,
		AverageScore: 99.9,
	}
	dataJSON, err := json.Marshal(newStudent)
	if err != nil {
		log.Fatalln("err stuct to json:", err)
	}
	js := string(dataJSON)
	log.Println("js=", js)
	ind, err := esclient.Index().
		Index("student").
		BodyJson(js).
		Do(ctx)
		
	if err != nil {
		log.Fatalln("error insert:", err.Error())
	}
	
	fmt.Println("[Elastic][InsertProduct]Insertion Successful",ind.Version)

}

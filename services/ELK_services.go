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

	client, err := elastic.NewClient(elastic.SetURL(ELK.UrlELK),elastic.SetSniff(false), elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")
	if err != nil {
		log.Fatalln("error init ELK:", err.Error())
	}

	return client, err

}

//InstertData is function InstertData
func (ELK *ELKServices) InstertData() {

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

// GetData is function get data from index in ELK
func (ELK *ELKServices) GetData(indexName string)  {

		ctx := context.Background()
		ELK.UrlELK = "http://localhost:9200"
		esclient, err := ELK.initELK()
		if err != nil {
			fmt.Println("Error initializing : ", err)
			panic("Client fail ")
		}
	
		var students [] models.Student
	
		searchSource := elastic.NewSearchSource()
		searchSource.Query(elastic.NewMatchQuery("name", "Doe"))
	
		/* this block will basically print out the es query */
		// queryStr, err1 := searchSource.Source()
		// queryJs, err2 := json.Marshal(queryStr)
	
		// if err1 != nil || err2 != nil {
		// 	fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
		// }
		// fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))
		/* until this block */
	
		searchService := esclient.Search().Index("students").SearchSource(searchSource)
		
		searchResult, err :=  searchService.Do(ctx)
		if err != nil {
			fmt.Println("[ProductsES][GetPIds]Error=", err)
			return
		}
	
		for _, hit := range searchResult.Hits.Hits {
			var student models.Student
			err := json.Unmarshal(hit.Source, &student)
			if err != nil {
				fmt.Println("[Getting Students][Unmarshal] Err=", err)
			}
	
			students = append(students, student)
		}
	
		if err != nil {
			fmt.Println("Fetching student fail: ", err)
		} else {
			for _, s := range students {
				fmt.Printf("Student found Name: %s, Age: %d, Score: %f \n", s.Name, s.Age, s.AverageScore)
			}
		}
	
	
}

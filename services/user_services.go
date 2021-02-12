package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/fluk27/StockMagageSysyem/models"
	outlib "github.com/fluk27/testgo2020/services"
	elastic "github.com/olivere/elastic/v7"
)

//UserServices is stuct of userServices
type UserServices struct {
}

//InstertDataUsers is function add data user to ELK
func (us *UserServices) InstertDataUsers(index string, user *models.UserModel) (result string, err error) {
	rsa:=outlib.RSAKey{
		PathPublicKey:"./",
		FileNamePublicKey: "publicKey.pem",
		PathPrivateKey: "./",
		FileNamePrivateKey: "privateKey.pem",
	}
	result,err=rsa.EncyptDataWithPKC("0844038001")
	if err != nil {
		return " ", err
	}
	log.Println("rr:",result)
	ex,err:=rsa.DncyptDataWithPKC(result)
	log.Println("ex:",ex)
	// ELK := &ELKServices{}

	// ctx := context.Background()
	// ELK.UrlELK = "http://localhost:9200"
	// esclient, err := ELK.initELK()
	// if err != nil {
	// 	log.Println("Error initializing : ", err)
	// 	return "", err
	// }

	// //creating student object

	// dataJSON, err := json.Marshal(user)
	// if err != nil {
	// 	log.Fatalln("err stuct to json:", err)
	// 	return "", err
	// }
	// js := string(dataJSON)
	// _, err = esclient.Index().
	// 	Index(index).
	// 	BodyJson(js).
	// 	Do(ctx)

	// if err != nil {
	// 	log.Fatalln("error insert:", err.Error())
	// 	return "", err
	// }

	return "insert successfull", nil

}

// getdataUsers is function get data from index in ELK
func (us *UserServices) getdataUsers(indexName string, query map[string]string) (*[]models.UserModel, error) {
	ELK := &ELKServices{}
	ctx := context.Background()
	ELK.UrlELK = "http://localhost:9200"
	esclient, err := ELK.initELK()
	if err != nil {
		log.Println("Error initializing : ", err)
		panic("Client fail ")
	}

	var students []models.UserModel

	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("name", "Doe"))

	searchService := esclient.Search().Index(indexName).SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return nil, err
	}

	for _, hit := range searchResult.Hits.Hits {
		var student models.UserModel
		err := json.Unmarshal(hit.Source, &student)
		if err != nil {
			fmt.Println("[Getting Students][Unmarshal] Err=", err)
		}

		students = append(students, student)
	}

	if err != nil {
		fmt.Println("Fetching student fail: ", err)
	}
	return &students, nil

}

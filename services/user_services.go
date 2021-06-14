package services

import (
	"context"
	//"crypto/sha256"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"log"

	"github.com/fluk27/personManageSystem/models"
	elastic "github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	validator "gopkg.in/validator.v2"
)

//UserServices is stuct of userServices
type UserServices struct {
}

//InstertDataUsers is function add data user to ELK
func (us *UserServices) InstertDataUsers(index string, user *models.UserModel) (result string, err error) {

	ELK := &ELKServices{}

	ctx := context.Background()
	ELK.UrlELK = "http://localhost:9200"
	esclient, err := ELK.initELK()
	if err != nil {
		log.Println("Error initializing : ", err)
		return "", err
	}

	//creating  object

	dataJSON, err := json.Marshal(user)
	if err != nil {
		//	log.Fatalln("err stuct to json:", err)
		return "", errors.New(err.Error())
	}
	js := string(dataJSON)
	_, err = esclient.Index().
		Index(index).
		BodyJson(js).
		Do(ctx)

	if err != nil {
		log.Fatalln("error insert:", err.Error())
		return "", errors.New(err.Error())
	}

	return "insert successfull", nil

}

// GetdataUsers is function get data from index in ELK
func (us *UserServices) GetdataUsers(indexName string, query map[string]string) (*[]models.UserModel, error) {
	var users []models.UserModel
	ELK := &ELKServices{}
	ctx := context.Background()
	ELK.UrlELK = "http://localhost:9200"
	esclient, err := ELK.initELK()
	if err != nil {
		log.Println("Error initializing : ", err)
		errors.New("Client fail ")
	}

	_, err = ELK.checkELK(indexName)
	if err != nil {
		err := ELK.CreateIndex(indexName)
		if err != nil {

			return nil, errors.New(err.Error())
		}

	}
	searchSource := elastic.NewSearchSource()
	if len(query) == 0 {
		// get all data	
		searchSource.Query(elastic.NewMatchAllQuery())
		} else {
			// name rsult from query map
			var Name string
			
			for i := range query {
				Name = i
			}
		
		//	elastic.NewAdjacencyMatrixAggregation().Filters()
			searchSource.Query(elastic.NewMatchQuery(Name, query[Name]))
			elastic.NewBoolQuery().Filter()
	}
	searchService := esclient.Search().Index(indexName).SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)

		return nil, errors.New(err.Error())
	}

	for _, hit := range searchResult.Hits.Hits {
		var user models.UserModel
		err := json.Unmarshal(hit.Source, &user)
		if err != nil {
			fmt.Println("[Getting Students][Unmarshal] Err=", err)
		}

		users = append(users, user)
	}

	if err != nil {
		fmt.Println("Fetching student fail: ", err)
	}

	return &users, nil

}

func (us *UserServices) ValidateStruct(user *models.UserModel) (bool, []error) {
	var errorValidate []error
	var satutsValldate bool
	satutsValldate = false
	err := validator.WithPrintJSON(true).Validate(user)
	if err != nil {
		satutsValldate = true
		log.Println("err from validate:", err.Error())
		errorValidate = append(errorValidate, err)
		// values not valid, deal with errors here
		return satutsValldate, errorValidate
	}
	return satutsValldate, nil
}

func (us *UserServices) Hashfunction512(data string) (string, error) {
	hashFuknc := sha512.New()
	result, err := hashFuknc.Write([]byte(data))
	if err != nil {
		return " ", errors.New(err.Error())
	}
	return string(rune(result)), nil
}

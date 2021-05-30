package services

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

//	"github.com/fluk27/StockMagageSysyem/models"
	elastic "github.com/olivere/elastic/v7"
	// Import the Elasticsearch library packages
)

//ELKServices is stuct
type ELKServices struct {
	UrlELK string
}

//GetESClient is function
func (ELK *ELKServices) initELK() (*elastic.Client, error) {

	client, err := elastic.NewClient(elastic.SetURL(ELK.UrlELK), elastic.SetSniff(false), elastic.SetHealthcheck(true), elastic.SetHealthcheckTimeout(1*time.Millisecond))

	//fmt.Println("ES initialized...")
	if err != nil {
		//	log.Fatalln("error init ELK:", err.Error())
		return nil, errors.New("elasticSearch cann't connect")
	}

	return client, err

}

// func (ELK *ELKServices) Siglequery(indexName string, query map[string]string) (result map[string]interface{}, err error) {
// 	var resultELK map[string]interface{}
// 	el, err := ELK.initELK()
// 	if err != nil {
// 		return nil, errors.New("elasticsearch don't connect")
// 	}
// 	searchSource := elastic.NewSearchSource()
// 	if len(query) == 0 {
// 		// get all data
// 		searchSource.Query(elastic.NewMatchAllQuery())
// 	} else {
// 		// name rsult from query map
// 		var Name string

// 		for i := range query {
// 			Name = i
// 		}

// 		//	elastic.NewAdjacencyMatrixAggregation().Filters()
// 		searchSource.Query(elastic.NewMatchQuery(Name, query[Name]))
// 		elastic.NewBoolQuery().Filter()
// 	}
// 	searchService := el.Search().Index(indexName).SearchSource(searchSource)

// 	searchResult, err := searchService.Do(context.Background())
// 	if err != nil {
// 		fmt.Println("[ProductsES][GetPIds]Error=", err)

// 		return nil, errors.New(err.Error())
// 	}

// 	for _, hit := range searchResult.Hits.Hits {
// 		var user models.UserModel
// 		err := json.Unmarshal(hit.Source, &user)
// 		if err != nil {
// 			fmt.Println("[Getting Students][Unmarshal] Err=", err)
// 		}

// 		resultELK = append(resultELK, user)
// 	}

// 	if err != nil {
// 		fmt.Println("Fetching student fail: ", err)
// 	}
// }

//getIndexElastic get indexName from ELK
// func (ELK *ELKServices) getIndexElastic(name string) (string, error) {
// 	ELK.UrlELK = "http://localhost:9200"
// 	esclient, err := ELK.initELK()
// 	if err != nil {
// 		return "", errors.New(err.Error())
// 	}
// 	//esclient.IndexGet("users").Do(ctx)
// 	nameIndex, err := esclient.IndexNames()
// 	if err != nil {
// 		log.Fatalln("error get indexName:", err)
// 		return " ", err
// 	}
// 	if len(nameIndex) == 0 {
// 		//	log.Fatalln("count index in elasticSeach=",nameIndex)
// 		return "", errors.New("don't have index")
// 	}
// 	for i := 0; i < len(nameIndex); i++ {
// 		if nameIndex[i] == name {
// 			return name, nil
// 		}
// 	}
// 	return name, nil
// }

//CreateIndex is function create index in ELK
// func (ELK *ELKServices) CreateIndex(name string) error {
// 	ctx := context.Background()
// 	esclient, err := ELK.initELK()
// 	if err != nil {
// 		return errors.New(err.Error())
// 		//	panic("Client fail ")
// 	}
// 	esclient.CreateIndex(name).Do(ctx)
// 	return nil
// }


// checkELK is function check index name is elk is empty
func (ELK *ELKServices) checkELK(indexName string) (map[string]interface{}, error){
	var result map[string]interface{}
url:="http://localhost:9200"

var URL = url+`/`+indexName+`?pretty`
resp, err := http.Head(URL)
if err != nil {
log.Fatalln("errors from ",err.Error())
}
defer resp.Body.Close()

 if resp.StatusCode==http.StatusOK{
	body, err := io.ReadAll(resp.Body)
if err != nil {
	return nil, errors.New(err.Error())
}
	json.Unmarshal(body,&result)
	return result,nil
 }else{
	return nil ,errors.New(indexName+" find not found")
 }
}

// CreateIndex is create index name in elk
func (ELK *ELKServices) CreateIndex(indexName string) (error){
url:="http://localhost:9200"
client:=&http.Client{}
var URL = url+`/`+indexName+`?pretty`
req,err:=http.NewRequest(http.MethodPut,URL,nil)
if err != nil {
	return errors.New(err.Error())
}
resp,err:=client.Do(req)
if err != nil {
	return errors.New(err.Error())
}

if resp.StatusCode!=http.StatusOK{

	return errors.New(indexName+"not created")
 }
 return nil
}
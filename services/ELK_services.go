package services

import (
	"context"
	"errors"
	"log"
	"time"

	elastic "github.com/olivere/elastic/v7"
	// Import the Elasticsearch library packages
)

//ELKServices is stuct
type ELKServices struct {
	UrlELK string
}

//GetESClient is function
func (ELK *ELKServices) initELK() (*elastic.Client, error) {

	client, err := elastic.NewClient(elastic.SetURL(ELK.UrlELK), elastic.SetSniff(false), elastic.SetHealthcheck(true),elastic.SetHealthcheckTimeout(1 *time.Millisecond))

	//fmt.Println("ES initialized...")
	if err != nil {
	//	log.Fatalln("error init ELK:", err.Error())
	return nil,errors.New("elasticSearch cann't connect")
	}

	return client, err

}

func (ELK *ELKServices) getIndexElastic(name string)  (string,error){
	ELK.UrlELK = "http://localhost:9200"
	esclient, err := ELK.initELK()
	if err != nil {
		return "",errors.New(err.Error())
	}
	//esclient.IndexGet("users").Do(ctx)
	nameIndex,err:=esclient.IndexNames()
	if err != nil {
		log.Fatalln("error get indexName:" ,err)
		return " ", err
	}
	if len(nameIndex)==0 {
	//	log.Fatalln("count index in elasticSeach=",nameIndex)
		return "",errors.New("don't have index")
	}
for i := 0; i < len(nameIndex); i++ {
	if nameIndex[i]== name {
		return name,nil
	}
}
return name,nil
}
//CreateIndex is function create index in ELK
func(ELK *ELKServices) CreateIndex(name string)  error{
	ctx:=context.Background()
	esclient, err := ELK.initELK()
	if err != nil {
	return errors.New(err.Error())
	//	panic("Client fail ")
	}
	esclient.CreateIndex(name).Do(ctx)
	return nil
}

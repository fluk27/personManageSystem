package services

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v3"
)

type SendMailServices struct {
	Sender     string
	DataMail   string
	HeaderMail string
}

type mailConfig struct {
	user     string `yaml:"SENDER"`
	password string `yaml:"SECRET"`
	port     string `yaml:"PORT"`
	host     string `yaml:"SERVER"`
}
// vaildateConfig 
func (sm *SendMailServices) vaildateConfig() {
	buf, err := ioutil.ReadFile("./configs/g_emaiL.env.yaml")
	if err != nil {
		log.Println("err read file yaml:",err.Error())
	}

	c := &mailConfig{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		log.Println("err map yaml to struct:",err.Error())
	}

	// return c, nil
}
func (sm *SendMailServices) validateDataMail() {

}

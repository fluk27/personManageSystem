package services

import (
	"errors"
	"io/ioutil"
	"log"

	validator "github.com/go-playground/validator/v10"
	yaml "gopkg.in/yaml.v2"
)

type SendMailServices struct {
	Sender     string `validate:"required"`
	DataMail   string `validate:"required"`
	HeaderMail string `validate:"required"`
}

type mailConfig struct {
	User     string `yaml:"SENDER" validate:"required"`
	Password string `yaml:"SECRET" validate:"required"`
	Port     string `yaml:"PORT" validate:"required"`
	Host     string `yaml:"SERVER" validate:"required"`
}

// vaildateConfig
func (sm *SendMailServices) vaildateConfig() (*mailConfig, error) {
	buf, err := ioutil.ReadFile("./helper/email/configs/g_emaiL.env.yaml")
	if err != nil {
		log.Println("err read file yaml:", err.Error())
		return nil, errors.New(err.Error())
	}

	mailCon := &mailConfig{}
	err = yaml.Unmarshal(buf, mailCon)
	if err != nil {
		log.Println("err map yaml to struct:", err.Error())
		return nil, errors.New(err.Error())
	}
	var validate *validator.Validate
	errr := validate.Struct(mailConfig)
	if errr != nil {
		return nil, err
	}
	return mailCon, nil
}
func (sm *SendMailServices) validateDataMail() error {
	log.Println("dataMail:", sm)
	return nil
}

func (sm *SendMailServices) SendMail() error {
	_, err := sm.vaildateConfig()
	if err != nil {
		return errors.New(err.Error())
	}
	err = sm.validateDataMail()
	return nil
}

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

var validate *validator.Validate

// vaildateConfig form yaml fie
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
	validate = validator.New()
	errr := validate.Struct(mailCon)
	if errr != nil {
		log.Println("error struct mailConfig:", err.Error())
		return nil, errors.New(err.Error())
	}
	return mailCon, nil
}
func (sm *SendMailServices) validateDataMail() error {
	validate = validator.New()
	err := validate.Struct(sm)
	if err != nil {
		log.Println("error struct mailConfig:", err.Error())
		return  errors.New(err.Error())
	}
	return nil
}
// SendMail config data with stuct 
func (sm *SendMailServices) SendMail() error {
	_, err := sm.vaildateConfig()
	if err != nil {
		return errors.New(err.Error())
	}
	err = sm.validateDataMail()
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

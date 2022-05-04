package services

type SendMailServices struct {
	Sender     string
	DataMail   string
	HeaderMail string
}

type mailConfig struct {
	user     string
	password string
	port     string
	host     string
}

func (sm *SendMailServices) vaildateConfig() {

}
func (sm *SendMailServices) validateDataMail() {

}

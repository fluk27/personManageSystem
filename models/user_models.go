package models

//UserModel is set interface all atrribute user
type UserModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FistName string `json:"fristName"`
	TelNumber string`json:"telNumber"`
}

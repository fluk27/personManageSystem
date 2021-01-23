package models

//UserModel is set interface all atrribute users
type UserModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FistName string `json:"fristName"`
	LastName string `json:"lastName"`
	Contact  *Contact
}
type Contact struct {
	TelNumber    string  `json:"telNo."`
	EmailAddress string `json:"email"`
	LineID       string
	address      *Address
}

//Address is set interface address users
type Address struct {
	AddressNumber string `json:"addNo."`
	Moo           int8   `json:"moo"`
	Soi           string `json:"soi"`

	PostCode int16 `json:"posCode"`
}

package models

//UserModel is set interface all atrribute users
type UserModel struct {
	Username string `json:"username"validate:"required"`
	Password string `json:"password"validate:"required"`
	FistName string `json:"fristName"validate:"required"`
	LastName string `json:"lastName"validate:"required"`
	Contact  *contact
}

// Contact is set interface all contact of users
type contact struct {
	TelNumber    string `json:"telNo."`
	EmailAddress string `json:"email"`
	LineID       string `json:"lineID"`
}

//Address is set interface address users
type Address struct {
	HomeAddress  *homeAddress
	TowerAddress *towerAddress
}

//HomeAddress is set interface address HomeAddress of users
type homeAddress struct {
	AddressNumber string `json:"addNo."`      // บ้านเลขที่
	Moo           int    `json:"moo"`         // หมู่
	Alley         string `json:"alley"`       // ตรอก
	Lane          string `json:"Lane"`        // ซอย
	Road          string `json:"road"`        // ถนน
	SubDistrict   string `json:"subDistrict"` // ตำบล
	District      string `json:"district"`    //อำเภอ
	Province      string `json:"province"`    // จังหวัด
	PostalCode    int16  `json:"postalCode"`  // รหัสไปรษณีย์
}

//TowerAddress is set interface address TowerAddress of users
type towerAddress struct {
	Name          string `json:"towerName"`   //ชื่ออาคารหรือตึก
	Floor         string `int:"floor"`        // ชั้น
	AddressNumber string `json:"addNo."`      // บ้านเลขที่
	Moo           int    `json:"moo"`         // หมู่
	Alley         string `json:"alley"`       // ตรอก
	Lane          string `json:"Lane"`        // ซอย
	Road          string `json:"road"`        // ถนน
	SubDistrict   string `json:"subDistrict"` // ตำบล
	District      string `json:"district"`    //อำเภอ
	Province      string `json:"province"`    // จังหวัด
	PostalCode    int16  `json:"postalCode"`  // รหัสไปรษณีย์
}

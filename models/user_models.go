package models

//UserModel is set interface all atrribute users
type UserModel struct {
	EmailAddress string `json:"email"validate:"nonzero" elk:"email"`
	Password string `json:"password"validate:"nonzero" elk:"password"` 
	FistName string `json:"fristName"validate:"nonzero"`
	LastName string `json:"lastName"validate:"nonzero"`
	Contact  *contact
}

// Contact is set interface all contact of users
type contact struct {
	TelNumber    string `json:"telNo."validate:"nonzero"`
	LineID       string `json:"lineID"validate:"nonzero"`
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

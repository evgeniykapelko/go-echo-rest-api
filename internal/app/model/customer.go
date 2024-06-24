package model

type Customer struct {
	Id        int    `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Country   string `json:"country"`
	Password  string `json:"password"`
}

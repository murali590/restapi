package model

import (
	"D:\restapi\util"
	"gorm.io/gorm"
)

// Album - represent a movie album in Inflix store
type user struct {
	gorm.Model
	userid        string          `json:"title"`
	Password      string          `json:"password"`
	name          string          `json:"name"`
	date_of_birth util.SimpleDate `json:"date_of_birth"`
	address       string          `json:"address"`
}

//TableName - to reriteve the db table name
func (us *user) TableName() string {
	return "us"
}

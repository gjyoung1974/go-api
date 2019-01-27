package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
)


type Person struct {
	gorm.Model
	Status bool   `json:"status"`
	Firstname string   `json:"firstname,omitempty`
	Lastname  string   `json:"lastname,omitempty"`
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}


func (e *Person) Disable() {
	e.Status = false
}

func (p *Person) Enable() {
	p.Status = true
}


// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
//func DBMigrate(db *gorm.DB) *gorm.DB {
func DBMigrate() *gorm.DB {
	db, _ := gorm.Open("sqlite3", ".//people.db")
	db.AutoMigrate(&Person{})
	defer db.Close()
	return db
}
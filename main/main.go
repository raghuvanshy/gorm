package main

import (
	"fmt"
	"github.com/raghuvanshy/gorm"
	"strconv"
)
import _ "github.com/raghuvanshy/gorm/dialects/mysql"

type Customer struct {
	ID        uint `gorm:"primary_key"`
	Ph        int
	PhNo	  string `gorm:"SOURCE:Ph;CONVERTER:PhNoConvert"`
}

func (c Customer) PhNoConvert(ph int) (phNo string) {
	return strconv.Itoa(ph)
}

func main() {
	var dbDSN string
	if dbDSN == "" {
		dbDSN = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True"
	}
	db, err := gorm.Open("mysql", dbDSN)
	if err != nil {
		fmt.Println("ERROR WHILE CONNECTING TO DATABASE")
		fmt.Println(err.Error())
	}
	db.AutoMigrate(Customer{})
	db.Create(&Customer{Ph:998877})
	db.Create(&Customer{Ph:887766})
	db.Create(&Customer{Ph:776655})
	db.Create(&Customer{Ph:998877})
	db.Create(&Customer{Ph:887766})
	db.Create(&Customer{Ph:776655})
}
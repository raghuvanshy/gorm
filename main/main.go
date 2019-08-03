package main

import (
	"fmt"
	"github.com/raghuvanshy/gorm"
	"github.com/raghuvanshy/gorm/Models"
)
import _ "github.com/raghuvanshy/gorm/dialects/mysql"

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
	db.AutoMigrate(Models.Customer{})
	/*db.Create(&Models.Customer{Ph:998877})
	db.Create(&Models.Customer{Ph:887766})
	db.Create(&Models.Customer{Ph:776655})
	db.Create(&Models.Customer{Ph:998877})
	db.Create(&Models.Customer{Ph:887766})
	db.Create(&Models.Customer{Ph:776655})*/
}
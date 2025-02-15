package main

import (
	"server/main.go/database"
	"server/main.go/models"
)

func main() {

	var db = database.Connect()

	if err := db.Debug().AutoMigrate(&models.User{}).Error; err != nil {
		panic(err)
	}
	if err := db.Debug().AutoMigrate(&models.CreditCards{}).Error; err != nil {
		panic(err)
	}
	if err := db.Debug().AutoMigrate(&models.CreditCardApplication{}).Error; err != nil {
		panic(err)
	}
}

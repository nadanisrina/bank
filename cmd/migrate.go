package main

import (
	"bank/user"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=bank port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	// Add name field
	errMigrate := db.Migrator().AddColumn(&user.User{}, "token")

	if errMigrate != nil {
		log.Panic(errMigrate)
	}

	fmt.Println("success migrate")
}

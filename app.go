package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/ws", WebsocketHandler)
}

type Messages struct {
	gorm.Model
	Message string
}

var DB *gorm.DB

func main(){
	dsn := "host=localhost user=postgres password=242090 dbname=golang port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return
	 } else {
		fmt.Println("go")
	 }
	 
	 if !DB.Migrator().HasTable(&Messages{}) {
		DB.Migrator().CreateTable(&Messages{})
	 }

	 DB.AutoMigrate(&Messages{})

	log.Fatal(http.ListenAndServe("192.168.0.101:8000", nil))
}
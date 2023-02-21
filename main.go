package main

import (
	"PAK/handler"
	"PAK/komponen"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pak?charset=utf8mb4&parseTime=True&loc=Local" //memasukkan hak akses
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	komponenpakRepository := komponen.NewRepository(db)
	komponenpakService := komponen.NewServiceKomponenPAK(komponenpakRepository)

	komponenpakhandler := handler.NewkomponenpakHandler(komponenpakService)

	router := gin.Default()

	api := router.Group("api/komponenpak")

	api.POST("/create", komponenpakhandler.Create)

	router.Run()

}

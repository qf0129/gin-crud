package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qf0129/gin-crud/pkg"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	Id    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{})

	pkg.Init(db)

	app := gin.Default()
	group1 := app.Group("/api")
	pkg.CreateRouter[Product](group1)
	app.Run()
}

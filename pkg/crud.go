package pkg

import (
	"gorm.io/gorm"
)

type GormModel any

var DB *gorm.DB

const (
	PRIMARY_KEY        = "id"
	DEFAULT_PAGE_INDEX = 1
	DEFAULT_PAGE_SIZE  = 10
)

func Init(db *gorm.DB) {
	DB = db
}

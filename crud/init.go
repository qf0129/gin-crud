package crud

import (
	"gorm.io/gorm"
)

var db *gorm.DB
var conf = &Config{}

func Init(d *gorm.DB, confs ...*Config) {
	if d == nil {
		panic("[Error] Invalid gorm.DB")
	}
	db = d

	for _, c := range confs {
		if c != nil {
			if c.PrimaryKey == "" {
				conf.PrimaryKey = defaultConf.PrimaryKey
			}
			if c.DefaultPageIndex == 0 {
				conf.DefaultPageIndex = defaultConf.DefaultPageIndex
			}
			if c.defaultPageSize == 0 {
				conf.defaultPageSize = defaultConf.defaultPageSize
			}
		}
	}

}

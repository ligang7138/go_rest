package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"sync"
)

const DB_URL = "ys_test:ys123456@tcp(47.93.205.0)/qy_shop?charset=utf8&parseTime=true&loc=Local"
const Time_Format  = "2006-01-02 15:04:05"

var DB *gorm.DB


func Init()  {
	/*db, err := gorm.Open("mysql", DB_URL)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", "ys_shop")
	}
	DB = db*/
	DB = getInstance()
}

// 获取单例
/**

 */
func getInstance() *gorm.DB {
	var (
		db *gorm.DB
		err error
		once = &sync.Once{}
	)
	once.Do(func() {
		db, err = gorm.Open("mysql", DB_URL)
		if err != nil {
			log.Errorf(err, "Database connection failed. Database name: %s", "ys_shop")
		}
	})
	return db
}

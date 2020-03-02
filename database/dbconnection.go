package database

import (
	"github.com/jinzhu/gorm"
"fmt"
//  _ "github.com/go-sql-driver/mysql"

// 	_ "github.com/jinzhu/gorm/dialects/mysql"
		_ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func init(){
	var err error
	DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}

	fmt.Println("hello")

// DB, err := gorm.Open("mysql", "sushil:sushil1234@(localhost)/apidb?charset=utf8&parseTime=True&loc=Local")
// 	if err != nil {
// 			fmt.Println(DB)
// 			panic(err)
// 		}
}
func ConnectDB() *gorm.DB {
	return DB
}

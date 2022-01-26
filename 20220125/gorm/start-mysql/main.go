package main

import (
	"fmt"
	//"github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	gorm.Model
	passport  string
	password string
	nickname string
}

func main() {
	fmt.Println("=================")

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:my-secret-pw@tcp(10.252.37.64:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	var user User

	result  := db.Find(&user)
	fmt.Println("result.RowsAffected =", result.RowsAffected)



	//rows := db.Row()
	//fmt.Println("rows = ", rows)


}
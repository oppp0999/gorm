package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type Product struct {
	gorm.Model
	code  string
	Price uint
}

func main() {
	db, err := gorm.Open("mysql", "root:mysqlroot@tcp(localhost:3306)/orm_test?charset=utf8")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("finish")

	//insert
	defer db.Close()
	db.AutoMigrate(&Product{})                 //최조 한 번만 실행도 됨. 테이블 자동생성
	db.Create(&Product{code: "M", Price: 100}) //생성
	fmt.Println(time.Now())

}

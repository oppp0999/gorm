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

/*gorm.Model을 포함하고 있는데 Model 구조체는 아래와 같이 생겼으며, 데이터 CRUD시에 gorm이 create, udpate, deleted 값을 넣어준다.
type Model struct {
	ID uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt DeletedAt `gorm:"index"`
}*/

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

	//read -> 결과를 p에 저장한다.
	p := &Product{}
	db.First(&p) // select * from table limit 1
	db.Last(&p)  // select * from table order by desc limit 1
	/*
		조건 지정하는 경우
		db.Where("필드명 = ?", "값").First(&p)
		이외에도 직관적인 여러 메소드를 메소드 체이닝으로 사용할 수 있다.
		db.Where().Or().Not().Limit().Offset().Order().Group().Having().Find()
	*/

	//UPDATE GORM은 아쉽게도 Dirty Checking 기능이 없다.
	db.Model(&Product).Update("Price", 200) //price -> 200으로 수정
	// 풀어쓴 형태
	db.First(&Product**)
	Product.Price = 200
	db.Save(&Product)
	//값을 변경한 후 Save() 메소드를 사용하여 저장해 준다.

	//DELETE
	db.Where("price = ?", 122).Delete(&Product{})
	// UPDATE "products" SET "deleted_at"='2021-08-29 22:04:44.911' WHERE price = 122 AND "products"."deleted_at" IS NULL
	//soft delete가 실행된다.

}

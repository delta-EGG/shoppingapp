package db

import (
	// フォーマットI/O
	"fmt"

	// GORM
	"github.com/jinzhu/gorm"

	// エンティティ(データベースの行)
	entity "mvc/models/entity"
)

// open DB接続
func open() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "Shopping"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	// DBエンジンヲInnoDBに
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	// 詳細なログを表示
	db.LogMode(true)

	// 登録するテーブルを単数形に
	db.SingularTable(true)

	// マイグレーション
	db.AutoMigrate(&entity.Product{})

	fmt.Println("db connected:", &db)
	return db
}

// FindAllProducts は商品テーブルのレコードを全権取得
func FindAllProducts() []entity.Product {
	products := []entity.Product{}
	db := open()
	// select
	db.Order("ID asc").Find(&products)

	defer db.Close()
	return products
}

// FindProduct は商品のレコードを1つ取得
func FindProduct(productID int) []entity.Product {
	product := []entity.Product{}

	db := open()
	// select
	db.First(&product, productID)
	defer db.Close()

	return product
}

// InsertProduct は商品の追加
func InsertProduct(registerProduct *entity.Product) {
	db := open()
	// insert
	db.Create(&registerProduct)
	defer db.Close()
}

// UpdateStateProduct はレコードの状態変更
func UpdateStateProduct(productID int, productState int) {
	product := []entity.Product{}
	db := open()
	// update
	db.Model(&product).Where("ID = ?", productID).Update("State", productState)
	defer db.Close()
}

// DeleteProduct はレコードの削除
func DeleteProduct(productID int) {
	product := []entity.Product{}
	db := open()
	// delete
	db.Delete(&product, productID)
	defer db.Close()
}

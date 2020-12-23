package controller

import (
	// 文字列と基本データ型の変換
	strconv "strconv"

	//Gin
	"github.com/gin-gonic/gin"

	// エンティティ（DBのテーブルの行）
	entity "mvc/models/entity"

	// DBアクセスモジュール
	db "mvc/models/db"
)

// 商品の購入状態を定義
const (
	// NotPurchased=未購入
	NotPurchased = 0

	// Purchased=購入済
	Purchased = 1
)

// FetchAllProducts は全ての商品除法を取得
func FetchAllProducts(c *gin.Context) {
	resultProducts := db.FindAllProducts()

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultProducts)
}

// FindProduct は指定IDの商品情報を取得
func FindProduct(c *gin.Context) {
	productIDStr := c.Query("productID")
	productID, _ := strconv.Atoi(productIDStr)
	resultProduct := db.FindProduct(productID)

	// URLヘのアクセスに対してJSON
	c.JSON(200, resultProduct)
}

// AddProduct は商品をDBに登録
func AddProduct(c *gin.Context) {
	productName := c.PostForm("productName")
	productMemo := c.PostForm("productMemo")

	var product = entity.Product{
		Name:  productName,
		Memo:  productMemo,
		State: NotPurchased,
	}
	db.InsertProduct(&product)
}

// ChangeStateProduct は商品情報の状態を変更
func ChangeStateProduct(c *gin.Context) {
	reqProductID := c.PostForm("productID")
	reqProductState := c.PostForm("productState")

	productID, _ := strconv.Atoi(reqProductID)
	productState, _ := strconv.Atoi(reqProductState)
	changeState := NotPurchased

	// 商品状態が未購入の場合
	if productState == NotPurchased {
		changeState = Purchased
	} else {
		changeState = NotPurchased
	}
	db.UpdateStateProduct(productID, changeState)
}

// DeleteProduct は商品情報をDBから削除
func DeleteProduct(c *gin.Context) {
	productIDStr := c.PostForm("productID")
	productID, _ := strconv.Atoi(productIDStr)
	db.DeleteProduct(productID)
}

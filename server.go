package main

// サーバー周りの処理
import (
	// ロギングを行う
	"log"

	// HTTPを扱う
	"net/http"

	// Gin
	"github.com/gin-gonic/gin"

	// MySQLドライバ
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// コントローラ
	controller "mvc/controllers/controller"
)

func main() {
	// サーバを起動
	serve()
}

// サーバを起動する関数
func serve() {
	// デフォルトのミドルウェアでginのルータを作成
	// LoggerとクラッシュをキャッチするRecoveryを保有
	router := gin.Default()

	// 静的ファイルパス
	router.Static("/views", "./views")

	// ルータの設定
	// URLへのアクセスに対して静的ページを返す
	router.StaticFS("/shoppingapp", http.Dir("./views/static"))

	// 全ての商品情報のJSONを返す
	router.GET("/fetchAllProducts", controller.FetchAllProducts)

	// 1つの商品情報の状態のJSONを返す
	router.GET("/fetchProduct", controller.FindProduct)

	// 商品情報をDBへ登録
	router.POST("/addProduct", controller.AddProduct)

	// 商品情報の状態を変更
	router.POST("/changeStateProduct", controller.ChangeStateProduct)

	// 商品情報を削除
	router.POST("/deleteProduct", controller.DeleteProduct)

	// ルータ走らせてエラーが起きたらエラー処理
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed:", err)
	}
}

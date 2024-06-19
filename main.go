package main

import (
	"log"
	"test_api/controllers"
	"test_api/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// .envファイルから環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// データベース接続の初期化
	database.Init()

	// Ginルーターの作成
	r := gin.Default()

	// ルーティング設定
	r.GET("/users", controllers.GetUsers)

	// サーバーの起動
	r.Run(":8080")
}

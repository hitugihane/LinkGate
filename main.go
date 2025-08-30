package main

import (
	"context"
	"fuagfuga-2025-LinkGate/src/router"
	"fuagfuga-2025-LinkGate/src/usecase"
	"fuagfuga-2025-LinkGate/src/usecase/discord"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// MongoDB 接続 URI を環境変数から取得します。設定されていない場合はデフォルトを使用します。
	mongoURI := os.Getenv("MONGODB_URI")

	// MongoDB に接続するための context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// MongoDB クライアントの作成
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	// アプリ終了時にクライアントを必ず切断する
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Printf("MongoDBの切断に失敗🥺: %v", err)
		}
	}()

	// 使用するデータベースとコレクションを選択
	db := client.Database("linkgate")
	collection := db.Collection("posts")

	// Gin エンジンを初期化
	r := gin.Default()

	// ルーティングのセットアップ
	router.SetupRoutes(r, collection, ctx, client)

	go usecase.WatchChanges(collection)

	go discord.InitializeDiscordBot(collection)

	// サーバーを起動
	if err := r.Run(":8080"); err != nil {
		log.Fatal("サーバーの起動に失敗🥺:", err)
	}
}

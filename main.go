package main

import (
	"go-mysql-api/pkg/controllers"
	"go-mysql-api/pkg/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	utils.LoggingSettings()

	router := gin.Default()
	// router.Use() // 共通でやりたい処理があればここに入れる

	router.StaticFile("/", "./index.html")

	// User routing set up
	user := controllers.NewUserController()
	router.GET("api/user", user.GetUsers)
	router.GET("api/user/:id", user.GetUser)
	router.POST("api/user", user.CreateUser)
	router.PUT("api/user", user.UpdateUser)
	router.DELETE("api/user/:id", user.DeleteUser)

	// POSTで更新したい場合↓のように書ける
	// router.POST("/user/*action", user.UpdateUser)

	router.Run(":3000")
}
package handlers

import (
	"main/database"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getMain(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
}

func postLogin(router *gin.Engine) {
	router.POST("/login", func(ctx *gin.Context) {
		var name string = ctx.Query("Name")
		var email string = ctx.Query("Email")
		var password string = ctx.Query("Password")

		var db *gorm.DB = database.Connect()
		var user model.User = model.User{
			Name:      name,
			Email:     email,
			IpAddress: "localhost",
			Password:  password,
		}
		db.Create(&user)

		ctx.JSON(http.StatusOK, gin.H{
			"message": model.CreateAccessToken(user),
		})
	})
}

func Start() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*.html")

	getMain(router)
	postLogin(router)

	router.Run("localhost:8080")
}

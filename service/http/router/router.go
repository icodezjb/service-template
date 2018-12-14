package router

import (
	"github.com/buchenglei/service-template/service/http/controller"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	// 用户登录相关路由
	passport := r.Group("/passport")
	{
		passportController := controller.NewPassportController()
		passport.POST("/login", passportController.Login)
		passport.POST("/register", passportController.Register)
		passport.DELETE("/logout", passportController.Logout)
	}

	// 文章相关路由
	article := r.Group("/article")
	{
		article.POST("/create", nil)
		article.POST("/query", nil)
		article.DELETE("/delete", nil)
	}

	// 用户相关路由
	user := r.Group("/user")
	{
		user.GET("/info", nil)
		user.GET("/articles", nil)
	}

	return r
}

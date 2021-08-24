package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nnnk7777/mysql-file-server/controller"
)

// InitRouter Ginとルーティングの設定
func InitRouter(Db *gorm.DB) {
	g := gin.Default()

	g.Use(func(c *gin.Context) {
		c.Set("db", Db)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	// g.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Hello, World!!")
	// })
	g.Static("/", "./public")

	// g.GET("/save", )
	// g.GET("/load", )
	ctrl := controller.Controller{}
	g.POST("/file", ctrl.HandlePostFile)

	// u := g.Group("/users")
	// {
	// 	u.GET("/me", ctrl.HandleGetUserByEmail)
	// }

	g.Run(":8080")
}
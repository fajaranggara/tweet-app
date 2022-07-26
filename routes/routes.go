package routes

import (
	"net/http"
	"tweet-app/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Static("/assets", "./assets")

	r.LoadHTMLGlob("views/*")

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html" , gin.H{"Data": ""})
	})

	r.GET("/signup", controllers.SignUp)
	r.POST("/signup", controllers.CreateUser)
	r.GET("/users", controllers.GetAllUser)

	
	return r
}
package controllers

import (
	"net/http"
	"tweet-app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html" , gin.H{"Data": "Sign Up"})
}

func CreateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	user := models.User{
		Name: c.PostForm("name"),
		Email: c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	_, err := user.SaveUser(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/users")
}

func GetAllUser(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var user []models.User

	db.Find(&user)

	c.HTML(http.StatusOK, "user.html" , gin.H{"Data": user})
}
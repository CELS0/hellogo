package controllers

import (
	"strconv"

	"github.com/CELS0/hellogo/database"
	"github.com/CELS0/hellogo/models"
	"github.com/CELS0/hellogo/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	db := database.GetDatabase()

	var p models.User

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	p.Password = services.SHA256Encoder(p.Password)

	err = db.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func ShowUser(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}
	db := database.GetDatabase()

	var p models.User
	err = db.First(&p, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find user:" + err.Error(),
		})
	}
	c.JSON(200, p)
}

func ShowUsers(c *gin.Context) {
	db := database.GetDatabase()

	var ps []models.User
	err := db.Find(&ps).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list users: " + err.Error(),
		})
	}
	c.JSON(200, ps)
}

func UpdateUser(c *gin.Context) {
	db := database.GetDatabase()

	var p models.User

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	err = db.Save(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update user" + err.Error(),
		})
		return
	}
	c.JSON(200, p)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
	}

	db := database.GetDatabase()

	err = db.Delete(&models.User{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete user" + err.Error(),
		})
		return
	}
	c.Status(200)
}

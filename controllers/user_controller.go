package controllers

import (
	"hygya-api/database"
	"hygya-api/models"
	"hygya-api/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	db := database.GetDataBase()

	var p models.User

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não é possível vincular o JSON: " + err.Error(),
		})
		return
	}

	p.Password = services.SHA256Encoder(p.Password)

	err = db.Create(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não pode criar um Paciente: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func ShowUsers(c *gin.Context) {
	db := database.GetDataBase()

	var user []models.User
	err := db.Find(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list users: " + err.Error(),
		})
		return
	}

	c.JSON(200, user)

}

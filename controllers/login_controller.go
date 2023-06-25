package controllers

import (
	"hygya-api/database"
	"hygya-api/models"
	"hygya-api/services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db := database.GetDataBase()

	var p models.Login

	err := c.ShouldBindJSON(&p)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não é possível vincular o JSON: " + err.Error(),
		})
		return
	}

	//VALIDAÇÕES
	var user models.User
	dbError := db.Where("email = ?", p.Email).First(&user).Error

	if dbError != nil {
		c.JSON(400, gin.H{
			"error": "Não pode encontrar um Usuário: ",
		})
		return
	}

	if user.Password != services.SHA256Encoder(p.Password) {
		c.JSON(401, gin.H{
			"error": "Credenciais inválidas",
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"id":       user.ID,
		"email":    user.Email,
		"token":    token,
		"password": user.Password,
	})
}

package controllers

import (
	"hygya-api/database"
	"hygya-api/models"
	"hygya-api/services"
	"strconv"

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

func ShowUser(c *gin.Context) {
	id := c.Param("id")

	//Atoi converte para INT
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID tem que ser um inteiro",
		})
		return
	}

	db := database.GetDataBase()

	var user models.User

	err = db.First(&user, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não pode encontrar o usuário: " + err.Error(),
		})
		return
	}
	c.JSON(200, user)
}

func ShowUsers(c *gin.Context) {
	db := database.GetDataBase()

	var user []models.User
	err := db.Find(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não pode listar os usuários: " + err.Error(),
		})
		return
	}

	c.JSON(200, user)
}

// FUNÇÃO PARA ATUALIZAR UM PACIENTE
func UpdateUsers(c *gin.Context) {
	db := database.GetDataBase()

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não é possível vincular o JSON: " + err.Error(),
		})
		return
	}
	err = db.Save(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Erro": "Não pode atualizar o usuário: " + err.Error(),
		})
		return
	}

	c.JSON(200, user)
}

// FUNÇÃO PARA DELETAR UM PACIENTE
func DeleteUsers(c *gin.Context) {
	id := c.Param("id")

	//Atoi converte para INT
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID tem que ser inteiro",
		})
		return
	}

	db := database.GetDataBase()

	err = db.Delete(&models.User{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não pode deletar o usuário: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

package controllers

import (
	"hygya-api/database"
	"hygya-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FUNÇÃO PARA MOSTRAR UM PACIENTE
func ShowPaciente(c *gin.Context) {
	id := c.Param("id")

	//Atoi converte para INT
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDataBase()

	var paciente models.Paciente

	err = db.First(&paciente, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não pode encontrar o paciente: " + err.Error(),
		})
		return
	}
	c.JSON(200, paciente)
}

// FUNÇÃO PARA CRIAR OS PACIENTES
func CreatePaciente(c *gin.Context) {
	db := database.GetDataBase()

	var paciente models.Paciente

	err := c.ShouldBindJSON(&paciente)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não é possível vincular o JSON: " + err.Error(),
		})
		return
	}
	err = db.Create(&paciente).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não pode criar um Paciente: " + err.Error(),
		})
		return
	}

	c.JSON(200, paciente)
}

// FUNÇÃO PARA MOSTRAR OS PACIENTES
func ShowPacientes(c *gin.Context) {
	db := database.GetDataBase()

	var pacientes []models.Paciente
	err := db.Find(&pacientes).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list pacientes: " + err.Error(),
		})
		return
	}

	c.JSON(200, pacientes)

}

// FUNÇÃO PARA ATUALIZAR UM PACIENTE
func UpdatePacientes(c *gin.Context) {
	db := database.GetDataBase()

	var paciente models.Paciente

	err := c.ShouldBindJSON(&paciente)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não é possível vincular o JSON: " + err.Error(),
		})
		return
	}
	err = db.Save(&paciente).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Erro": "Não pode atualizar o paciente: " + err.Error(),
		})
		return
	}

	c.JSON(200, paciente)
}

// FUNÇÃO PARA DELETAR UM PACIENTE
func DeletePacientes(c *gin.Context) {
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

	err = db.Delete(&models.Paciente{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não pode deletar o Paciente: " + err.Error(),
		})
		return
	}

	c.Status(204)

}

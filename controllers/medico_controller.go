package controllers

import (
	"hygya-api/database"
	"hygya-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FUNÇÃO PARA MOSTRAR UM MEDICO
func ShowMedico(c *gin.Context) {
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

	var medico models.Medico

	err = db.First(&medico, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não pode encontrar o médico: " + err.Error(),
		})
		return
	}
	c.JSON(200, medico)
}

// FUNÇÃO PARA CRIAR OS MEDICOS
func CreateMedico(c *gin.Context) {
	db := database.GetDataBase()

	var medico models.Medico

	err := c.ShouldBindJSON(&medico)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não é possível vincular o JSON: " + err.Error(),
		})
		return
	}
	err = db.Create(&medico).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não pode criar um médico: " + err.Error(),
		})
		return
	}

	c.JSON(200, medico)
}

// FUNÇÃO PARA MOSTRAR OS PACIENTES
func ShowMedicos(c *gin.Context) {
	db := database.GetDataBase()

	var medico []models.Medico
	err := db.Find(&medico).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não pode listar os médicos: " + err.Error(),
		})
		return
	}

	c.JSON(200, medico)

}

// FUNÇÃO PARA ATUALIZAR UM PACIENTE
func UpdateMedicos(c *gin.Context) {
	db := database.GetDataBase()

	var medico models.Medico

	err := c.ShouldBindJSON(&medico)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não é possível vincular o JSON: " + err.Error(),
		})
		return
	}
	err = db.Save(&medico).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Erro": "Não pode atualizar o médico: " + err.Error(),
		})
		return
	}

	c.JSON(200, medico)
}

// FUNÇÃO PARA DELETAR UM PACIENTE
func DeleteMedicos(c *gin.Context) {
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

	err = db.Delete(&models.Medico{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não pode deletar o médico: " + err.Error(),
		})
		return
	}

	c.Status(204)

}

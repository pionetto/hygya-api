package controllers

import (
	"hygya-api/database"
	"hygya-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FUNÇÃO PARA MOSTRAR UM PACIENTE
func ShowEspecialidade(c *gin.Context) {
	id := c.Param("id")

	//Atoi converte para INT
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID deve ser um inteiro",
		})
		return
	}

	db := database.GetDataBase()

	var especialidade models.Especialidade

	err = db.First(&especialidade, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não pode encontrar a especialidade: " + err.Error(),
		})
		return
	}
	c.JSON(200, especialidade)
}

// FUNÇÃO PARA CRIAR AS ESPECIALIDADES
func CreateEspecialidade(c *gin.Context) {
	db := database.GetDataBase()

	var especialidade models.Especialidade

	err := c.ShouldBindJSON(&especialidade)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não é possível vincular o JSON: " + err.Error(),
		})
		return
	}
	err = db.Create(&especialidade).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não pode criar uma especialidade: " + err.Error(),
		})
		return
	}

	c.JSON(200, especialidade)
}

// FUNÇÃO PARA MOSTRAR AS ESPECIALIDADES
func ShowEspecialidades(c *gin.Context) {
	db := database.GetDataBase()

	var especialidades []models.Especialidade
	err := db.Find(&especialidades).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não pode listar as especialidades: " + err.Error(),
		})
		return
	}

	c.JSON(200, especialidades)

}

// FUNÇÃO PARA ATUALIZAR UMA ESPECIALIDADE
func UpdateEspecialidades(c *gin.Context) {
	db := database.GetDataBase()

	var especialidade models.Especialidade

	err := c.ShouldBindJSON(&especialidade)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não é possível vincular o JSON: " + err.Error(),
		})
		return
	}
	err = db.Save(&especialidade).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Erro": "Não pode atualizar a especialidade: " + err.Error(),
		})
		return
	}

	c.JSON(200, especialidade)
}

// FUNÇÃO PARA DELETAR UMA ESPECIALIDADE
func DeleteEspecialidades(c *gin.Context) {
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

	err = db.Delete(&models.Especialidade{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não pode deletar a Especialidade: " + err.Error(),
		})
		return
	}

	c.Status(204)

}

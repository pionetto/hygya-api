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

// FUNÇÃO PARA MOSTRAR OS PACIENTES
func ShowPacientes(c *gin.Context) {
	db := database.GetDataBase()

	var pacientes []models.Paciente
	err := db.Find(&pacientes).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não pode listar os pacientes: " + err.Error(),
		})
		return
	}

	c.JSON(200, pacientes)

}

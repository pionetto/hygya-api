package controllers

import (
	"hygya-api/database"
	"hygya-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
			"error": "cannot find paciente: " + err.Error(),
		})
		return
	}
	c.JSON(200, paciente)
}

func CreateBook(c *gin.Context) {
	db := database.GetDataBase()

	var book models.Paciente

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	err = db.Create(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

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

func UpdateBook(c *gin.Context) {
	db := database.GetDataBase()

	var book models.Paciente

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	err = db.Save(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update book: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
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

	err = db.Delete(&models.Paciente{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete book: " + err.Error(),
		})
		return
	}

	c.Status(204)

}

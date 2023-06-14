package routes

import (
	"hygya-api/controllers"

	"github.com/gin-gonic/gin"
)

// Função ConfigRoutes recebe router do tipo *gin.Engine devolvendo *gin.Engine
func ConfigRoutes(router *gin.Engine) *gin.Engine {

	//Criando Grupo Main, que deixa o api/v1 como padrão da rota
	main := router.Group("api/v1")
	{
		//Criando o Grupo de Livros (books)
		books := main.Group("paciente")
		{
			books.GET("/:id", controllers.ShowPaciente)
			books.GET("/", controllers.ShowPacientes)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}
	return router
}

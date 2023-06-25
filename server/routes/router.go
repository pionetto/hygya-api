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
		user := main.Group("users")
		{
			user.POST("/", controllers.CreateUser)
			user.GET("/", controllers.ShowUsers)
		}
		//Criando o Grupo de Pacientes
		// pacientes := main.Group("paciente", middlewares.Auth())
		pacientes := main.Group("paciente")
		{
			pacientes.GET("/:id", controllers.ShowPaciente)
			pacientes.GET("/", controllers.ShowPacientes)
			pacientes.POST("/", controllers.CreatePaciente)
			pacientes.PUT("/", controllers.UpdatePacientes)
			pacientes.DELETE("/:id", controllers.DeletePacientes)
		}
		medicos := main.Group("medico")
		{
			medicos.GET("/:id", controllers.ShowMedico)
			medicos.GET("/", controllers.ShowMedicos)
			medicos.POST("/", controllers.CreateMedico)
			medicos.PUT("/", controllers.UpdateMedicos)
			medicos.DELETE("/:id", controllers.DeleteMedicos)
		}
		especialidades := main.Group("especialidades")
		{
			especialidades.GET("/:id", controllers.ShowEspecialidade)
			especialidades.GET("/", controllers.ShowEspecialidades)
			especialidades.POST("/", controllers.CreateEspecialidade)
			especialidades.PUT("/", controllers.UpdateEspecialidades)
			especialidades.DELETE("/:id", controllers.DeleteEspecialidades)
		}

		main.POST("login", controllers.Login)
	}
	return router
}

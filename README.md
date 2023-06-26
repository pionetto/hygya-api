<h1 align="center">:file_cabinet: HYGYA PACS - Sua Saúde nas nuvens</h1>

## :memo: Descrição
Este projeto é um software para consumo de dados de um servidor de imagens [DICOM](https://www.dicomstandard.org/) [PACS](https://pt.wikipedia.org/wiki/PACS).
Ele serve para que o médico possa receber as imagens DICOM para laudo de exames de 
imagem (Ultrassonografia, Tomografia, Raio-X e Mamografias) de forma mais rápida
e de qualquer lugar, pelo fato da aplicação ser web.

## :books: Funcionalidades
* <b>Funcionalidade 1</b>
</br>
Esta é uma API desenvolvida na linguagem [Golang](https://go.dev/).

Ela realiza o cadastro, edição, visualização e remoção de:

* Usuários;
* Médicos (controllers/user_controller.go);
* Especialidades;

Ela realiza visualização de:

* Pacientes/Exames;

Pois os exames e pacientes são dados que não podem ser apagados.

O usuário realiza o login via middleware, através de token de autentitação 
via JWT com Bearer no Header ( arquivo em server/middlewares/auth_middleware.go ).

Estes itens são acessados através das seguintes rotas:

```go
func ConfigRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("api/v1")
	{
		user := main.Group("users")
		{
			user.POST("/", controllers.CreateUser)
			user.GET("/", controllers.ShowUsers)
			user.PUT("/", controllers.UpdateUsers)
			user.DELETE("/:id", controllers.DeleteUsers)
		}
		 pacientes := main.Group("paciente", middlewares.Auth())
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

```





## :wrench: Tecnologias utilizadas
* [Golang](https://go.dev/);
* [Gin Framework](https://github.com/gin-gonic/gin);
* [Gorm](gorm.io/gorm);
* [Gorilla Mux](github.com/gorilla/mux);
* [PostgreSQL](https://go.dev/);
* [Github](https://go.dev/);
* [Visual Studio Code](https://go.dev/);
* [S.O. Linux Mint](https://go.dev/);
* [Insomnia](https://go.dev/) - Para alguns testes da API;

## :rocket: Rodando o projeto
São necessários alguns requisitos para rodar a aplicação.

* Ter o PostgreSQL instalado e configurado, tem um guia clicando [aqui](https://www.edivaldobrito.com.br/como-instalar-o-postgresql-no-ubuntu-20-04-lts-e-derivados/).
* Ter instalado o Golang.

Para instalar o golang, faça:

```
sudo apt install -y golang
```

Crie um diretório (mkdir) em seu lugar de preferência.

```
sudo apt install -y golang
```
Para rodar o repositório é necessário clonar o mesmo, dar o seguinte comando para iniciar o projeto:

```
git clone https://github.com/pionetto/hygya-api.git
```

Em seguida:

```
git clone https://github.com/pionetto/hygya-api.git
```


## :soon: Implementação futura
* O que será implementado na próxima sprint?

## :handshake: Colaboradores
<table>
  <tr>
    <td align="center">
      <a href="http://github.com/pionetto">
        <img src="https://avatars.githubusercontent.com/u/5672555?v=4" width="100px;" alt="Foto de Pio Netto no GitHub"/><br>
        <sub>
          <b>@pionetto</b>
        </sub>
      </a>
    </td>
  </tr>
</table>

## :dart: Status do projeto
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

Abaixo estão as rotas criadas:

```go
func ConfigRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("api/v1")
	{
		user := main.Group("users")
		{
      user.GET("/:id", controllers.ShowUser)
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
E para acessa-las, podemos utilizar uma ferramenta chamada Insomnia(clique [aqui](https://insomnia.rest/download) para baixar),
para testar se os endpoints estão funcionando corretamente.
Ela serve para que possamos fazer requisições (GET, POST, PUT, DELETE..).


## Sobre a API
* Todos os caminhos da API poderão ser acessados a partir do link http://localhost:5000/api/v1;
* Todos os caminhos da api utilizam o metodo de requisição GET;
* As respostas das requisições feitas a api são em formato JSON;

# User
</br>

**Criar Usuário**
</br>

**Método:** POST
</br>

**Endpoint:** http://localhost:5000/api/v1/users
</br>

**Objeto JSON a ser enviado:**

```JSON
{
		"name": "Pio Lima",
		"email": "pio@pio.com",
		"password": "123456"
}
```

**Listar Usuários**
</br>

**Método:** GET
</br>

**Endpoint:** http://localhost:5000/api/v1/users
</br>
</br>

**Listar um Usuário**
</br>

**Método:** GET
</br>

**Endpoint:** http://localhost:5000/api/v1/users/{id}
</br>
</br>

**Atualizar Usuários**
</br>

**Método:** PUT
</br>

**Endpoint:** http://localhost:5000/api/v1/users
</br>

**Objeto JSON a ser enviado:**

```JSON
{
	"id": 3,
	"name": "Pandora Cunha",
	"email": "pandora@pandora.com",
	"password": "pandora"
}
```

**Deletar um Usuário**
</br>

**Método:** DELETE
</br>

**Endpoint:** http://localhost:5000/api/v1/users/{id}
</br>


# Medico
</br>

**Criar Médico**
</br>

**Método:** POST
</br>

**Endpoint:** http://localhost:5000/api/v1/medico
</br>

**Objeto JSON a ser enviado:**

```JSON
{
		"nome": "Hipócrates da Grecia Antiga",
		"sexo": "M",
		"data_nasc": "1900-06-20T00:00:00Z",
		"cpf": "99999999999",
		"especialidade": "1",
		"rqe": "44664644",
		"crm_estado": "PARÁ",
		"crm_numero": "121212",
		"end_logradouro": "RUA ABCD",
		"end_numero": "11111",
		"end_complemento": "ESQUINA COM RUA B",
		"end_bairro": "CENTRO",
		"end_cidade": "SANTARÉM",
		"end_uf": "PARÁ",
		"end_cep": "68000-000",
		"contato_email": "teste@teste.com",
		"contato_telefone1": "939999-9999",
		"contato_telefone2": " ",
		"contato_telefone3": " ",
		"contato_obs": " "
}
```

**Listar Médicos** </br>
</br>

**Método:** GET 
</br>

**Endpoint:** http://localhost:5000/api/v1/medico
</br>

**Listar um Médico**
</br>

**Método:** GET
</br>

**Endpoint:** http://localhost:5000/api/v1/medico/{id}

**Atualizar Médico**
</br>

**Método:** PUT
</br>

**Endpoint:** http://localhost:5000/api/v1/medico
</br>

**Objeto JSON a ser enviado:**

```JSON
{
	"id": 1,
	"nome": "Hipócrates da Grecia Nova",
	"sexo": "M",
	"data_nasc": "1900-06-20T00:00:00Z",
	"cpf": "99999999999",
	"especialidade": "Radiologia",
	"rqe": "44664644",
	"crm_estado": "PARÁ",
	"crm_numero": "121212",
	"end_logradouro": "RUA ABCD",
	"end_numero": "11111",
	"end_complemento": "ESQUINA COM RUA B",
	"end_bairro": "CENTRO",
	"end_cidade": "SANTARÉM",
	"end_uf": "PARÁ",
	"end_cep": "68000-000",
	"contato_email": "teste@teste.com",
	"contato_telefone1": "939999-9999",
	"contato_telefone2": " ",
	"contato_telefone3": " ",
	"contato_obs": " "
}
```

**Deletar um Médico** 
</br>

**Método:** DELETE 
</br>

**Endpoint:** http://localhost:5000/api/v1/medico/{id}
</br>



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
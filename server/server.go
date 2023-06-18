package server

import (
	"hygya-api/server/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("X-Requested-With", "Content-Type", "Authorization")
	config.AddAllowMethods("GET", "POST", "PUT", "DELETE")
	r.Use(cors.New(config))
	r.Use(cors.New(config))

	return Server{
		port:   "5000",
		server: r,
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Print("server is running at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}

// package server

// import (
// 	"hygya-api/server/routes"
// 	"log"

// 	"github.com/gin-gonic/gin"
// )

// type Server struct {
// 	port   string
// 	server *gin.Engine
// }

// func NewServer() Server {
// 	return Server{
// 		port:   "5000",
// 		server: gin.Default(),
// 	}
// }

// func (s *Server) Run() {
// 	router := routes.ConfigRoutes(s.server)

// 	log.Print("server is running at port: ", s.port)
// 	log.Fatal(router.Run(":" + s.port))
// }

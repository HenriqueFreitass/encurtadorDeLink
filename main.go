package main

import (
	"database/sql"
	"log"
	"net/http"

	"encurtador-de-link/backend/config"
	"encurtador-de-link/backend/handlers"
	"encurtador-de-link/backend/repository"
	"encurtador-de-link/backend/routes"
	"encurtador-de-link/backend/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.InitDB()
	defer db.Close()

	// Conectar ao banco
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/linkShortener")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Criar camadas da aplicação
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	shortenerRepo := repository.NewShortenerRepository(db)
	shortenerService := service.NewService(shortenerRepo)
	shortenerHandler := handlers.NewShortenerHandler(shortenerService)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Permite apenas o front-end
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Configurar rotas
	routes.SetupUserRoutes(router, userHandler, shortenerHandler)

	router.Run(":8080")
	// Iniciar servidor
	log.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

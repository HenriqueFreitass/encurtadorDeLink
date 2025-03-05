package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	DB *sql.DB
}

// LoadEnv carrega as variáveis do arquivo .env
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠  Arquivo .env não encontrado, carregando variáveis de ambiente do sistema.")
	}
}

// InitDB inicializa a conexão com o banco de dados
func InitDB() *sql.DB {
	LoadEnv()

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("❌ Erro ao conectar no banco de dados: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("❌ Banco de dados inacessível: %v", err)
	}

	log.Println("✅ Conexão com banco de dados estabelecida!")
	return db
}

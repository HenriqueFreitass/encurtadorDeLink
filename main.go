package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql" // Importa o driver MySQL
)
func main() {
    // String de conexão com o banco de dados
    dsn := "root:senha@tcp(localhost:3306)/linkShortener"
    
    // Abre a conexão com o banco de dados
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Verifica se a conexão foi estabelecida corretamente
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
}




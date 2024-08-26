package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "meu-projeto/db"
    "meu-projeto/handlers"
)

func main() {
    r := gin.Default()

    // Conecta ao banco de dados
    db.ConnectDB()

    // Define as rotas
    r.POST("/users", handlers.CreateUser)
    r.GET("/users/:id", handlers.GetUser)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Inicia o servidor
    if err := r.Run(":" + port); err != nil {
        log.Fatal(err)
    }
}

package handlers

import (
    "net/http"
    "strconv"
    "os"

    "github.com/gin-gonic/gin"
    "meu-projeto/db"
    "meu-projeto/models"
)

// init configura o ambiente e a conexão com o banco de dados antes da execução dos testes e da aplicação.
// O TDD exige que o ambiente de teste esteja configurado corretamente para garantir que os testes sejam realizados em um ambiente controlado e consistente, refletindo o comportamento do sistema em condições reais.
func init() {
    os.Setenv("MYSQL_DSN", "root:4002@tcp(127.0.0.1:3306)/ponderada_db?charset=utf8mb4&parseTime=True&loc=Local")
    db.ConnectDB() // Estabelece a conexão com o banco de dados usando o DSN configurado
}

// CreateUser é o manipulador para a rota POST /users.
// Ele lida com a criação de um novo usuário no banco de dados.
// No contexto de TDD, o objetivo é garantir que o endpoint de criação de usuário processe corretamente as requisições e retorne a resposta esperada.
// Isso inclui a verificação de dados recebidos, a criação do usuário e a resposta da API.
func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }
    c.JSON(http.StatusOK, user)
}

// GetUser é o manipulador para a rota GET /users/:id.
// Ele lida com a recuperação de um usuário do banco de dados usando o ID fornecido.
// O TDD foca na garantia de que o endpoint de recuperação de usuári funcione corretamente ao buscar um usuário pelo ID e retornar a resposta esperada.
func GetUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        // Verifica se o ID fornecido é válido. Se não for, retorna um status 400 Bad Request com uma mensagem de erro indicando que o ID do usuário é inválido.
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    var user models.User
    if err := db.DB.First(&user, id).Error; err != nil {
        // Tenta buscar o usuário no banco de dados pelo ID. Se não encontrar o usuário ou ocorrer um erro, retorna um status 500 Internal Server Error com uma mensagem indicando que o usuário não foi encontrado.
        c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
        return
    }

    // Se o usuário for encontrado com sucesso, retorna o usuário com um status 200 OK.
    c.JSON(http.StatusOK, user)
}

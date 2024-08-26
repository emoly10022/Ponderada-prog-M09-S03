package handlers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "meu-projeto/db"
    "meu-projeto/models"
)

// setupTestRouter configura o roteador para os testes, incluindo as rotas e os manipuladores necessários.
// Em TDD, a configuração do roteador de teste é crucial para garantir que os testes interajam com o sistema da mesma forma que um usuário real faria, validando o comportamento da API em condições controladas.
func setupTestRouter() *gin.Engine {
    db.ConnectDB() // Conecta ao banco de dados para os testes
    router := gin.Default()
    router.POST("/users", CreateUser) // Rota para criar usuário
    router.GET("/users/:id", GetUser)  // Rota para obter usuário por ID
    return router
}

// TestCreateUser testa o endpoint de criação de usuário para garantir que um usuário seja criado corretamente e que a resposta seja a esperada. Este teste verifica o comportamento do sistema quando um pedido de criação de usuário é feito e valida se a resposta corresponde ao esperado.
// A criação e a validação do usuário garantem que o endpoint de criação está funcionando como pretendido.
func TestCreateUser(t *testing.T) {
    router := setupTestRouter()

    // Cria um novo usuário para ser enviado na requisição POST
    user := models.User{
        Name:  "Test User",
        Email: "testuser@example.com",
    }
    jsonValue, _ := json.Marshal(user) 
    req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue)) // Cria uma nova requisição POST
    req.Header.Set("Content-Type", "application/json") 

    w := httptest.NewRecorder() // Cria um gravador de resposta
    router.ServeHTTP(w, req) // Envia a requisição para o roteador

    assert.Equal(t, http.StatusOK, w.Code) 

    var responseUser models.User
    err := json.Unmarshal(w.Body.Bytes(), &responseUser) // Converte a resposta JSON para um objeto User
    assert.Nil(t, err) // Verifica se houve erro ao deserializar a resposta
    assert.Equal(t, "Test User", responseUser.Name) // Verifica se o nome do usuário na resposta é o esperado
    assert.Equal(t, "testuser@example.com", responseUser.Email) // Verifica se o e-mail do usuário na resposta é o esperado
}

// TestGetUser testa o endpoint de obtenção de usuário para garantir que um usuário possa ser recuperado corretamente usando seu ID. Este teste assegura que o sistema está respondendo adequadamente a uma solicitação GET e que o usuário retornado corresponde ao que foi criado.
// Criar um usuário antes de fazer a requisição GET valida que a recuperação de dados está funcionando corretamente.
func TestGetUser(t *testing.T) {
    router := setupTestRouter()

    // Cria um usuário para buscar depois
    user := models.User{Name: "Test User", Email: "testuser@example.com"}
    db.DB.Create(&user) // Salva o usuário no banco de dados

    req, _ := http.NewRequest("GET", "/users/1", nil) // Cria uma nova requisição GET para o usuário com ID 1
    w := httptest.NewRecorder() // Cria um gravador de resposta
    router.ServeHTTP(w, req) // Envia a requisição para o roteador

    assert.Equal(t, http.StatusOK, w.Code) // Verifica se o status da resposta é 200 OK

    var responseUser models.User
    err := json.Unmarshal(w.Body.Bytes(), &responseUser) // Converte a resposta JSON para um objeto User
    assert.Nil(t, err) // Verifica se houve erro ao deserializar a resposta
    assert.Equal(t, "Test User", responseUser.Name) // Verifica se o nome do usuário na resposta é o esperado
}

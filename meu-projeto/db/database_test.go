package db

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "meu-projeto/models"
)

// TestConnectDB é um teste unitário que verifica a conexão com o banco de dados.
// Este teste é fundamental no TDD, pois estabelece uma base para garantir que o ambiente de testes possa se conectar ao banco de dados antes de realizar qualquer operação.
// Asserção: Verifica se a conexão DB não é nula, o que significa que a conexão foi estabelecida com sucesso.
func TestConnectDB(t *testing.T) {
    ConnectDB()
    assert.NotNil(t, DB)
}

// TestCreateAndRetrieveUser é um teste que valida duas funcionalidades principais: 
// a criação de um registro de usuário no banco de dados e a recuperação desse mesmo registro.
// No TDD, a prática é escrever testes para funcionalidades antes mesmo de implementá-las, garantindo que o código funcione conforme o esperado antes de ser liberado.
func TestCreateAndRetrieveUser(t *testing.T) {
    // Conecta ao banco de dados para garantir que estamos prontos para interagir com o banco
    ConnectDB()
    
    // Cria um novo usuário no banco de dados.
    // Este teste simula a operação de inserção de dados e é essencial para garantir que a funcionalidade de criação de usuários esteja funcionando corretamente.
    user := models.User{Name: "Test DB User", Email: "testdbuser@example.com"}
    err := DB.Create(&user).Error
    // Asserção: Verifica se a operação de criação não retornou erro, garantindo que o usuário foi criado com sucesso.
    assert.Nil(t, err)

    // Recupera o usuário recém-criado do banco de dados.
    // Este teste valida que os dados inseridos podem ser corretamente recuperados, o que é crucial para a integridade das operações de CRUD (Create, Read, Update, Delete).
    var retrievedUser models.User
    err = DB.First(&retrievedUser, user.ID).Error
    // Asserção: Verifica se a operação de recuperação não retornou erro, garantindo que o usuário foi encontrado.
    assert.Nil(t, err)
    // Asserção: Verifica se o nome do usuário recuperado é o mesmo que foi inserido.
    // Esta verificação assegura que os dados não foram corrompidos ou alterados indevidamente durante a inserção.
    assert.Equal(t, "Test DB User", retrievedUser.Name)
}

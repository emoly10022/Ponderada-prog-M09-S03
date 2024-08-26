package db

import (
    "log"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "meu-projeto/models"
)

// DB é uma variável global que armazena a conexão com o banco de dados.
var DB *gorm.DB

// ConnectDB é uma função que configura a conexão com o banco de dados MySQL.
// No contexto do TDD, uma função como esta é crucial para estabelecer um ambiente de teste estável e confiável,permitindo que os testes interajam com um banco de dados real ou simulado conforme a necessidade.
func ConnectDB() {
    // Definindo a variável de ambiente MYSQL_DSN que contém informações de conexão com o banco de dados.
    // Em um cenário de TDD, isso permite configurar o ambiente de teste para interagir com o banco de dados durante os testes de forma controlada e consistente.
    os.Setenv("MYSQL_DSN", "root:4002@tcp(localhost:3306)/ponderada_db?charset=utf8mb4&parseTime=True&loc=Local")

    // Obtém o DSN do MySQL da variável de ambiente.
    // A obtenção do DSN é uma prática comum em TDD para garantir que a configuração do banco de dados seja flexível e possa ser facilmente modificada sem alterar o código-fonte diretamente.
    dsn := os.Getenv("MYSQL_DSN")

    // Abre a conexão com o banco de dados usando Gorm e MySQL.
    // Essa etapa é fundamental para TDD, pois estabelece a base sobre a qual todos os testes que interagem com o banco de dados serão executados. Se a conexão falhar, os testes que dependem dela também falharão.
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the database: ", err)
    }

    // Executa a migração automática para a estrutura de User.
    // A migração automática é uma técnica que garante que a estrutura do banco de dados esteja sincronizada com os modelos definidos no código. Em TDD, isso assegura que os testes sejam executados em um banco de dados com a estrutura correta, minimizando falhas relacionadas a inconsistências no esquema do banco de dados.
    db.AutoMigrate(&models.User{})

    // Define a conexão de banco de dados global.
    // Definir a variável global DB permite que diferentes partes do código acessem a conexão com o banco de dados, facilitando a realização de operações de CRUD durante os testes e garantindo que a configuração de conexão seja centralizada.
    DB = db
}

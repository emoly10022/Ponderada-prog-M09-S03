package models

import "gorm.io/gorm"

// User representa o modelo de usuário no banco de dados.
// Em TDD, a definição do modelo é fundamental, pois os testes dependem da estrutura dos dados que estão sendo manipulados.
type User struct {
    gorm.Model // Inclui campos padrão como ID, CreatedAt, UpdatedAt e DeletedAt fornecidos pelo GORM

    // Nome do usuário.
    // A validação e os testes devem garantir que o nome esteja presente e siga as regras de negócios aplicáveis.
    Name string `json:"name"`

    // Email do usuário.
    // Os testes devem validar que o email está corretamente formatado e é único, conforme as regras de validação de dados.
    Email string `json:"email"`
}

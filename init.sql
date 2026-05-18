CREATE TABLE pacientes (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    cpf VARCHAR(14) NOT NULL UNIQUE,
    data_nascimento DATE NOT NULL,
    diagnostico VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
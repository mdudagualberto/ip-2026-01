# Sistema de Gerenciamento de Pacientes com GO

## Visão Geral
Este projeto é um servidor HTTP desenvolvido em GoLang para gerenciamento de pacientes. O sistema permite cadastrar, listar, atualizar e deletar pacientes através de uma interface web simples e estilizada.

O projeto utiliza PostgreSQL como banco de dados e Docker para facilitar a configuração do ambiente.

---

## Estrutura do Projeto

### Diretórios Principais:
- **`app/`**: Contém a lógica do servidor, incluindo handlers e utilitários.
  - **`handlers/`**: Responsável por processar as requisições HTTP do CRUD de pacientes.
  - **`utils/`**: Contém funções auxiliares, como conexão com banco de dados e estrutura dos pacientes.
- **`static/`**: Contém os arquivos estáticos da aplicação.
  - **`forms/`**: Formulários HTML para cadastro, listagem, atualização e exclusão de pacientes.
  - **`css/`**: Arquivos CSS responsáveis pela estilização do sistema.

---

## Configuração do Ambiente

### Pré-requisitos
1. **GoLang** instalado na máquina.
2. **PostgreSQL** para armazenamento dos dados.
3. **Docker** (opcional, mas recomendado).

---

## Passos para Configuração

### 1. Clone o repositório

```bash
git clone <URL_DO_REPOSITORIO>
cd servidorHTTP
```

---

### 2. Configure o arquivo `.env`

Crie um arquivo `.env` na raiz do projeto contendo:

```env
DB_USER=<seu_usuario>
DB_PASSWORD=<sua_senha>
DB_NAME=<nome_do_banco>
DB_HOST=<host_do_banco>
DB_PORT=<porta_do_banco>
```

---

### 3. Configure o Banco de Dados

Se estiver utilizando Docker:

```bash
sudo docker compose up -d
```

O banco será inicializado automaticamente utilizando o arquivo:

```txt
init.sql
```

Tabela utilizada:

```sql
CREATE TABLE pacientes (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    cpf VARCHAR(14) NOT NULL UNIQUE,
    data_nascimento DATE NOT NULL,
    diagnostico VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

### 4. Instale as dependências

```bash
go mod tidy
```

---

## Executando o Projeto

### 1. Inicie o servidor

```bash
go run app/main.go
```

---

### 2. Acesse a aplicação

O servidor estará disponível em:

```txt
http://localhost:8080/
```

---

## Funcionalidades

### Rotas Principais
- **`/`**: Dashboard principal do sistema.
- **`/static/forms/criar.html`**: Cadastro de pacientes.
- **`/static/forms/listar.html`**: Listagem de pacientes.
- **`/static/forms/atualizar.html`**: Atualização de pacientes.
- **`/static/forms/deletar.html`**: Exclusão de pacientes.

---

## Handlers

- **`CriarPacienteHandler`**: Responsável pelo cadastro de pacientes.
- **`ListarPacientesHandler`**: Exibe os pacientes cadastrados.
- **`AtualizarPacienteHandler`**: Atualiza os dados do paciente.
- **`DeletarPacienteHandler`**: Remove pacientes do banco de dados.

---

## Estrutura de Pastas

```txt
.env
docker-compose.yml
go.mod
go.sum
init.sql

app/
  main.go

  handlers/
    createHandler.go
    readHandler.go
    updateHandler.go
    deleteHandler.go

  utils/
    connectToDB.go
    paciente.go

static/
  index.html

  css/
    style.css

  forms/
    criar.html
    listar.html
    atualizar.html
    deletar.html
```

---

## Observações

- Certifique-se de que o PostgreSQL esteja rodando antes de iniciar o servidor.
- O projeto utiliza o driver `github.com/lib/pq` para conexão com PostgreSQL.
- O sistema possui interface estilizada utilizando HTML e CSS puro.
- Caso ocorram erros relacionados ao Docker no Linux, utilize `sudo` nos comandos Docker.

---

## Desenvolvido por

Maria Eduarda Gualberto

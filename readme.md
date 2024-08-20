
# GoBooks

GoBooks é um projeto desenvolvido durante o curso intensivo de Golang do Dev FullCycle. Este projeto tem como objetivo criar uma API simples para gerenciar livros utilizando o banco de dados SQLite3. A API oferece endpoints para criar, listar, atualizar e excluir livros.

## Tecnologias Utilizadas

-   **Golang**: Linguagem de programação utilizada para desenvolver o projeto.
-   **SQLite3**: Banco de dados relacional utilizado para armazenar as informações dos livros.
-   **net/http**: Pacote nativo do Go para lidar com requisições HTTP.
-   **mattn/go-sqlite3**: Driver SQLite3 para Go.

## Estrutura do Projeto

O projeto é dividido em duas camadas principais:

1.  **service**: Camada responsável pela lógica de negócios, como operações de CRUD para livros.
2.  **web**: Camada responsável por lidar com as requisições HTTP e interagir com os serviços.

## Configuração e Execução

### Pré-requisitos

-   Go 1.18 ou superior instalado.
-   SQLite3 instalado.

### Passos para Rodar o Projeto

1.  **Clone o repositório:**
            
    `git clone https://github.com/seu-usuario/gobooks.git
    cd gobooks` 
    
2.  **Instale as dependências:**
    
    Use o Go Modules para instalar as dependências do projeto:
            
    `go mod tidy` 
    
3.  **Execute o projeto:**
    
    Rode o projeto com o comando:
    
    `go run main.go` 
    
    O servidor estará disponível em `http://localhost:8080`.
    

### Endpoints da API

-   **GET /books**: Retorna uma lista de todos os livros.
-   **POST /books**: Cria um novo livro.
-   **GET /books/{id}**: Retorna os detalhes de um livro específico pelo ID.
-   **PUT /books/{id}**: Atualiza as informações de um livro específico.
-   **DELETE /books/{id}**: Deleta um livro pelo ID.

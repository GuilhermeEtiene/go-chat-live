# Go Chat Live - Demonstração de Conhecimentos em Golang

Este é um projeto simples em Go desenvolvido para demonstrar conhecimentos práticos na linguagem Golang, aplicando conceitos importantes e melhores práticas de desenvolvimento.

## 🎯 Objetivo

Demonstrar proficiência em Go através de uma aplicação real de chat em tempo real, showcasing conhecimentos em:

- **Clean Architecture** - Separação clara de responsabilidades
- **Goroutines e Concorrência** - WebSocket com processamento assíncrono
- **Interfaces** - Repository pattern e abstrações
- **Dependency Injection** - Desacoplamento de componentes
- **Middleware** - CORS e autenticação JWT
- **ORM e Database** - GORM com PostgreSQL
- **Criptografia** - Hash de senhas com bcrypt
- **JSON Marshaling/Unmarshaling** - APIs RESTful
- **Error Handling** - Tratamento adequado de erros
- **Testing** - Testes unitários (service_test.go)

## 🏗️ Arquitetura

```
go-chat-live/
├── cmd/                    # Pontos de entrada da aplicação
│   ├── server/            # Servidor REST API
│   └── wsserver/          # Servidor WebSocket
├── internal/              # Código interno da aplicação
│   ├── chat/             # Domínio do chat em tempo real
│   ├── database/         # Configuração do banco de dados
│   └── user/             # Domínio de usuários
└── docker-compose.yml    # Infraestrutura PostgreSQL
```

### Padrões Aplicados

- **Repository Pattern**: Abstração da camada de dados
- **Service Layer**: Lógica de negócio centralizada
- **Handler Layer**: Controllers HTTP/WebSocket
- **Dependency Injection**: Inversão de dependências
- **Clean Architecture**: Separação de concerns

## 🚀 Tecnologias

- **Go 1.23.5** - Linguagem principal
- **Gin Gonic** - Framework web performático
- **GORM** - ORM para Go com PostgreSQL
- **Gorilla WebSocket** - Comunicação real-time
- **JWT** - Autenticação stateless
- **bcrypt** - Hash seguro de senhas
- **PostgreSQL** - Banco de dados relacional
- **Docker Compose** - Containerização

## 🔧 Configuração

### Pré-requisitos
- Go 1.23.5+
- Docker e Docker Compose
- Git

### Instalação

1. **Clone o repositório**
```bash
git clone <repository-url>
cd go-chat-live
```

2. **Configure variáveis de ambiente**
```bash
# Crie o arquivo .env na raiz do projeto
cp .env.example .env

# Configure as variáveis necessárias
JWT_SECRET=seu_jwt_secret_aqui
DB_HOST=localhost
DB_PORT=5433
DB_USER=chatuser
DB_PASSWORD=chatpass
DB_NAME=chatdb
```

3. **Inicie o PostgreSQL**
```bash
docker-compose up -d
```

4. **Instale dependências**
```bash
go mod tidy
```

5. **Execute os servidores**

Terminal 1 - API REST:
```bash
go run cmd/server/main.go
```

Terminal 2 - WebSocket:
```bash
go run cmd/wsserver/main.go
```

## 📡 API Endpoints

### Autenticação
- `POST /users` - Criar usuário
- `POST /login` - Autenticar usuário
- `GET /users` - Listar usuários
- `GET /users/:id` - Buscar usuário
- `PUT /users/:id` - Atualizar usuário
- `DELETE /users/:id` - Remover usuário

### WebSocket
- `WS /ws?room=<room_id>&token=<jwt_token>` - Conectar ao chat

## 🎮 Como Usar

1. **Create user**
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John","email":"john@test.com","password":"123456"}'
```

2. **Login**
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"john@test.com","password":"123456"}'
```

3. **Connect to WebSocket**
Use the JWT token returned from login to connect:
```
ws://localhost:8081/ws?room=room1&token=<your_jwt_token>
```

## 🔍 Conceitos Go Demonstrados

### 1. Goroutines e Concorrência
```go
// Hub processando eventos de forma assíncrona
func (h *Hub) Run() {
    for {
        select {
        case client := <-h.register:
            // Processar registro
        case client := <-h.unregister:
            // Processar desconexão
        case msg := <-h.broadcast:
            // Distribuir mensagem
        }
    }
}
```

### 2. Interfaces e Dependency Injection
```go
type UserRepository interface {
    Create(user *User) error
    FindAll() ([]User, error)
    FindById(id int) (*User, error)
    // ...
}
```

### 3. Middleware Patterns
```go
// CORS middleware com Gin
router.Use(cors.New(cors.Config{
    AllowOrigins: []string{"*"},
    AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
}))
```

### 4. Error Handling
```go
if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
    return nil, errors.New("senha incorreta")
}
```

### 5. JSON Marshal/Unmarshal
```go
type ChatMessage struct {
    Content  string `json:"content"`
    UserName string `json:"userName"`
}
```

## 🧪 Testes

Execute os testes unitários:
```bash
go test ./internal/user/
```

## 🔒 Segurança

- **JWT Tokens** para autenticação stateless
- **bcrypt** para hash de senhas
- **Validação** de dados de entrada
- **CORS** configurado adequadamente

## 📚 Aprendizados Demonstrados

Este projeto demonstra compreensão sólida de:

- ✅ Sintaxe e idiomas Go
- ✅ Concorrência com goroutines e channels
- ✅ Interfaces e composition over inheritance
- ✅ Error handling idiomático
- ✅ Package organization e módulos
- ✅ HTTP servers e REST APIs
- ✅ WebSocket real-time communication
- ✅ Database integration com GORM
- ✅ Testing practices
- ✅ Security best practices

## 🚀 Melhorias Futuras

- [ ] Implementar rate limiting
- [ ] Adicionar logs estruturados
- [ ] Metrics e monitoring
- [ ] Deployment com Docker
- [ ] CI/CD pipeline
- [ ] GraphQL API
- [ ] Redis para sessions
- [ ] Kubernetes deployment

---

**Desenvolvido como demonstração de conhecimentos em Go/Golang** 🐹

Este projeto serve como portfólio para demonstrar proficiência na linguagem Go, aplicando conceitos fundamentais e padrões de desenvolvimento profissional.

### 2. Instalar dependências
```bash
go mod tidy
```

### 3. Executar servidores

**Terminal 1 - Servidor REST:**
```bash
cd cmd/server
go run main.go
```

**Terminal 2 - Servidor WebSocket:**
```bash
cd cmd/wsserver  
go run main.go
```

### 4. Abrir o chat
Abra o arquivo `chat-auth.html` no navegador.

## 🗃️ Banco de dados

O projeto usa PostgreSQL via Docker. As configurações estão no `.env`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=chatuser
DB_PASSWORD=chatpass
DB_NAME=chatdb
```

## 🧪 Executar testes

```bash
# Todos os testes
go test -v ./...

# Apenas testes do módulo user
go test -v ./internal/user/
```

## 📡 API Endpoints

### Usuários
- `POST /users` - Cadastrar usuário
- `POST /login` - Fazer login
- `GET /users` - Listar usuários
- `GET /users/:id` - Buscar usuário por ID
- `PUT /users/:id` - Atualizar usuário
- `DELETE /users/:id` - Deletar usuário

### WebSocket
- `ws://localhost:8081/ws?room=SALA&token=JWT_TOKEN`

## 🏗️ Arquitetura

```
cmd/
├── server/     # Servidor REST (porta 8080)
└── wsserver/   # Servidor WebSocket (porta 8081)

internal/
├── chat/       # Lógica do chat WebSocket
├── database/   # Conexão PostgreSQL  
└── user/       # CRUD usuários + autenticação
```

## 🔐 Autenticação

1. Cadastre um usuário via `POST /users`
2. Faça login via `POST /login` para obter JWT
3. Use o token para conectar no WebSocket
4. Chat protegido - apenas usuários autenticados

## 🐳 Docker

O `docker-compose.yml` fornece:
- PostgreSQL 15
- Dados persistentes
- Rede isolada
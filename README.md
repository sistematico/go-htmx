# go-htmx

Projeto mínimo com **Go** (apenas biblioteca padrão) no backend e **htmx** no frontend.

## Características

- Zero dependências de terceiros — usa somente a stdlib do Go (`net/http`, `html/template`).
- Roteamento com o `http.ServeMux` nativo do Go 1.22+ (padrões como `GET /{$}` e rota catch-all para 404).
- htmx salvo localmente em `static/js/htmx.min.js` (sem uso de CDN).
- Arquivos estáticos (CSS e htmx) servidos a partir da pasta `static/`, conforme sugerido pela documentação do Go.

## Estrutura do projeto

```
go-htmx/
├── go.mod                       # módulo Go
├── cmd/server/main.go           # ponto de entrada da aplicação
├── internal/
│   ├── router/router.go         # definição das rotas
│   └── handlers/handlers.go     # handlers (Home, Hello, NotFound)
├── templates/
│   ├── index.html               # página inicial
│   └── 404.html                 # página de erro 404
└── static/
    ├── css/style.css
    └── js/htmx.min.js
```

## Pré-requisitos

- [Go](https://go.dev/dl/) 1.22 ou superior instalado (`go version`).

## Instalação

Não há dependências externas para instalar. Basta clonar/copiar o projeto e entrar na pasta:

```bash
cd go-htmx
```

## Executando em modo desenvolvimento

```bash
go run ./cmd/server
```

O servidor sobe em `http://localhost:8080`.

## Build (binário de produção)

```bash
go build -o bin/go-htmx ./cmd/server
```

Executar o binário gerado:

```bash
./bin/htmx-go
```

## Rotas disponíveis

| Método | Rota                | Descrição                                   |
|--------|----------------------|----------------------------------------------|
| GET    | `/`                  | Página inicial                                |
| GET    | `/hello`             | Endpoint de exemplo consumido via htmx (`hx-get`) |
| GET    | `/static/*`          | Arquivos estáticos (CSS, htmx.min.js)         |
| *      | qualquer outra rota  | Página 404                                    |

## Configuração

A porta do servidor está definida em `cmd/server/main.go` (constante `addr`, padrão `:8080`). Altere esse valor conforme necessário.

## Atualizando o htmx

O arquivo `static/js/htmx.min.js` é uma cópia local. Para atualizar a versão, baixe o novo release e substitua o arquivo:

```bash
curl -sL https://unpkg.com/htmx.org/dist/htmx.min.js -o static/js/htmx.min.js
```
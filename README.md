# go-htmx

Projeto mínimo com **Go** (apenas biblioteca padrão) no backend e **htmx** no frontend.

## Características

- Zero dependências de terceiros — usa somente a stdlib do Go (`net/http`, `html/template`).
- Roteamento com o `http.ServeMux` nativo do Go 1.22+ (padrões como `GET /{$}` e rota catch-all para 404).
- Sistema de templates com layout padrão (`html/template`): `base.html` define header/main/footer e cada página injeta seu conteúdo em `{{template "content" .}}`.
- htmx **2.0.10** salvo localmente em `static/js/htmx.min.js` (sem uso de CDN).
- Arquivos estáticos guardados em `static/` mas expostos sem o prefixo `/static`, via uma única rota curinga (`GET /{asset...}`) que cai no handler de 404 quando o arquivo não existe (ex.: `/css/style.css`, `/js/htmx.min.js`).

## Estrutura do projeto

```
go-htmx/
├── go.mod                          # módulo Go (github.com/sistematico/go-htmx)
├── cmd/go-htmx/main.go             # ponto de entrada da aplicação
├── internal/
│   ├── router/router.go            # definição das rotas
│   ├── handlers/handlers.go        # handlers (Home, Hello, NotFound)
│   └── templates/templates.go      # parsing e renderização dos templates
├── templates/
│   ├── base.html                   # layout padrão (header, main, footer)
│   ├── index.html                  # conteúdo da página inicial
│   └── 404.html                    # conteúdo da página de erro 404
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
go run ./cmd/go-htmx
```

O servidor sobe em `http://localhost:8080`.

## Build (binário de produção)

```bash
go build -o bin/go-htmx ./cmd/go-htmx
```

Executar o binário gerado:

```bash
./bin/go-htmx
```

## Rotas disponíveis

| Método | Rota                | Descrição                                   |
|--------|----------------------|----------------------------------------------|
| GET    | `/`                  | Página inicial                                |
| GET    | `/hello`             | Endpoint de exemplo consumido via htmx (`hx-get`) |
| GET    | `/{asset...}`        | Arquivos estáticos servidos a partir de `static/` (CSS, htmx.min.js), sem o prefixo `/static` |
| *      | qualquer outra rota  | Página 404                                    |

## Configuração

A porta do servidor está definida em `cmd/go-htmx/main.go` (constante `addr`, padrão `:8080`). Altere esse valor conforme necessário.

## Layout e templates

O layout padrão fica em `templates/base.html`, que define o template `base` (header, main, footer) e carrega o CSS e o htmx. Cada página (`index.html`, `404.html`) define apenas o bloco `content`, injetado dentro do layout via `internal/templates/templates.go`.

## Atualizando o htmx

O arquivo `static/js/htmx.min.js` é uma cópia local (versão 2.0.10). Para atualizar, baixe o novo release e substitua o arquivo:

```bash
curl -sL https://unpkg.com/htmx.org/dist/htmx.min.js -o static/js/htmx.min.js
```
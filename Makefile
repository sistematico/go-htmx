.PHONY: dev air build clean help

APP := go-htmx

dev: ## rodar com `go run` (desenvolvimento)
	go run ./cmd/$(APP)

air: ## rodar com live-reload via air (desenvolvimento)
	air -c .air.toml

build: ## compilar o binário de produção
	go build -o ./tmp/$(APP) ./cmd/$(APP)

clean: ## remover artefatos de build
	rm -rf ./tmp

help: ## listar os comandos disponíveis
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2}'
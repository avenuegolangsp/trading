# Sistema de Trading - Makefile
.PHONY: help setup build run test test-unit test-integration test-coverage clean lint docker-build docker-run

# Variáveis
APP_NAME=trading-server
BINARY_PATH=bin/$(APP_NAME)
GO_FILES=$(shell find . -name "*.go" -type f)

# Help
help: ## Mostra esta ajuda
	@echo "🏆 Sistema de Trading - Avenue"
	@echo ""
	@echo "Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Setup
setup: ## Configura o ambiente de desenvolvimento
	@echo "🔧 Configurando ambiente..."
	go mod tidy
	go mod download
	@echo "✅ Ambiente configurado!"

# Build
build: ## Compila o projeto
	@echo "🔨 Compilando..."
	mkdir -p bin
	go build -o $(BINARY_PATH) cmd/server/main.go
	@echo "✅ Compilação concluída: $(BINARY_PATH)"

# Run
run: run-web ## Alias para run-web

run-web: ## Executa o web service
	@echo "🚀 Iniciando web service..."
	go run internal/services/web/cmd/main.go

run-engine: ## Executa o engine service (TODO: implementar)
	@echo "🚀 Engine service não implementado ainda..."
	@echo "TODO: Implementar internal/services/engine/cmd/main.go"

# Run com build
run-binary: build ## Compila e executa o binário
	@echo "🚀 Executando binário..."
	./$(BINARY_PATH)

# Testes
test: ## Executa todos os testes
	@echo "🧪 Executando todos os testes..."
	go test -v ./...

test-unit: ## Executa testes unitários
	@echo "🧪 Executando testes unitários..."
	go test -v ./tests/unit/...

test-integration: ## Executa testes de integração
	@echo "🧪 Executando testes de integração..."
	go test -v ./tests/integration/...

test-coverage: ## Executa testes com cobertura
	@echo "🧪 Executando testes com cobertura..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "📊 Relatório de cobertura: coverage.html"

# Qualidade de código
lint: ## Executa linter
	@echo "🔍 Executando linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "⚠️ golangci-lint não encontrado. Instalando..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
		golangci-lint run; \
	fi

fmt: ## Formata o código
	@echo "🎨 Formatando código..."
	go fmt ./...
	@echo "✅ Código formatado!"

vet: ## Executa go vet
	@echo "🔍 Executando go vet..."
	go vet ./...

# Dados
load-data: ## Verifica se os dados estão carregados
	@echo "📊 Verificando dados..."
	@if [ -f "data/users.json" ] && [ -f "data/stocks.json" ]; then \
		echo "✅ Arquivos de dados encontrados"; \
		echo "👥 Usuários: $$(cat data/users.json | grep -o '"id"' | wc -l)"; \
		echo "📈 Ações: $$(cat data/stocks.json | grep -o '"company"' | wc -l)"; \
	else \
		echo "❌ Arquivos de dados não encontrados!"; \
		echo "Certifique-se de que data/users.json e data/stocks.json existem"; \
		exit 1; \
	fi

# Health checks
health: ## Verifica se o servidor está rodando
	@echo "❤️ Verificando saúde do servidor..."
	@curl -s http://localhost:8080/api/health || echo "❌ Servidor não está respondendo"

# Exemplos
example-order: ## Cria uma ordem de exemplo
	@echo "📝 Criando ordem de exemplo..."
	curl -X POST http://localhost:8080/api/orders \
		-H "Content-Type: application/json" \
		-d '{"user_id":"ana-silva","symbol":"AAPL","side":"BUY","quantity":10,"price":220.00}'

example-orderbook: ## Consulta order book de exemplo
	@echo "📚 Consultando order book AAPL..."
	curl -s http://localhost:8080/api/orderbook/AAPL

example-portfolio: ## Consulta portfolio de exemplo
	@echo "👤 Consultando portfolio ana-silva..."
	curl -s http://localhost:8080/api/portfolio/ana-silva

example-all: ## Testa todos os endpoints
	@echo "🧪 Testando todos os endpoints..."
	@echo "\n📍 Health Check:"
	@curl -s http://localhost:8080/api/health
	@echo "\n📍 Stocks:"
	@curl -s http://localhost:8080/api/stocks
	@echo "\n📍 Market Status:"
	@curl -s http://localhost:8080/api/market/status
	@echo "\n📍 Portfolio:"
	@curl -s http://localhost:8080/api/portfolio/ana-silva
	@echo "\n📍 Order Book:"
	@curl -s http://localhost:8080/api/orderbook/AAPL

# Limpeza
clean: ## Limpa arquivos gerados
	@echo "🧹 Limpando..."
	rm -rf bin/
	rm -f coverage.out coverage.html
	go clean -cache
	@echo "✅ Limpeza concluída!"

clean-all: clean docker-stop ## Limpeza completa
	@echo "🧹 Limpeza completa..."
	docker system prune -f
	@echo "✅ Limpeza completa concluída!"

# Informações
info: ## Mostra informações do projeto
	@echo "🏆 Sistema de Trading - Avenue"
	@echo ""
	@echo "📊 Estatísticas do projeto:"
	@echo "  Arquivos Go: $$(find . -name '*.go' | wc -l)"
	@echo "  Linhas de código: $$(find . -name '*.go' -exec cat {} \; | wc -l)"
	@echo ""
	@echo "🔧 Versão Go: $$(go version)"
	@echo "📁 Diretório: $$(pwd)"
	@echo ""
	@echo "🌐 URLs importantes:"
	@echo "  Health Check: http://localhost:8080/health"
	@echo "  API Orders: http://localhost:8080/orders"
	@echo "  Order Book: http://localhost:8080/orderbook/{symbol}"
	@echo "  Portfolio: http://localhost:8080/portfolio/{user_id}"

# Default
.DEFAULT_GOAL := help

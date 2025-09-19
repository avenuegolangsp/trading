package domain

import "errors"

// Erros de validação de negócio
var (
	// Portfolio errors
	ErrInsufficientBalance  = errors.New("saldo insuficiente")
	ErrInsufficientPosition = errors.New("posição insuficiente")

	// User errors
	ErrUserNotFound = errors.New("usuário não encontrado")
	ErrInvalidUser  = errors.New("usuário inválido")

	// Stock errors
	ErrInvalidSymbol   = errors.New("símbolo inválido")
	ErrPriceTooLow     = errors.New("preço abaixo do mínimo permitido")
	ErrInvalidPrice    = errors.New("preço inválido")
	ErrInvalidQuantity = errors.New("quantidade inválida")

	// Market errors
	ErrMarketClosed = errors.New("mercado fechado")

	// Order errors
	ErrInvalidOrder     = errors.New("ordem inválida")
	ErrOrderNotFound    = errors.New("ordem não encontrada")
	ErrExceedsLimit     = errors.New("ordem excede limite do perfil")
	ErrInvalidOrderSide = errors.New("lado da ordem inválido")

	// Matching errors
	ErrNoMatch = errors.New("nenhuma correspondência encontrada")
)

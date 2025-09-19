package validators

import (
	"trading/internal/domain"
)

// BusinessValidator implementa validações de regras de negócio
type BusinessValidator struct {
	// TODO: Implementar campos necessários (ex: stocks map)
}

// NewBusinessValidator cria um novo validador de negócio
func NewBusinessValidator() *BusinessValidator {
	validator := &BusinessValidator{
		// TODO: Inicializar campos
	}

	// TODO: Carregar dados de ações do arquivo JSON
	return validator
}

// ValidateOrder valida uma ordem completa
func (v *BusinessValidator) ValidateOrder(order *domain.Order) error {
	// TODO: Implementar validações completas
	// 1. Validar símbolo existe
	// 2. Validar preço mínimo
	// 3. Validar horário de mercado
	// 4. Validar campos básicos

	return nil
}

// ValidateSymbol valida se o símbolo existe
func (v *BusinessValidator) ValidateSymbol(symbol string) error {
	// TODO: Implementar validação de símbolo
	// 1. Verificar se símbolo está na lista de 20 ações
	// 2. Retornar erro se não encontrado

	return nil
}

// ValidateMinPrice valida se o preço está acima do mínimo
func (v *BusinessValidator) ValidateMinPrice(symbol string, price float64) error {
	// TODO: Implementar validação de preço mínimo
	// 1. Obter preço mínimo do símbolo
	// 2. Comparar com preço fornecido
	// 3. Retornar erro se abaixo do mínimo

	return nil
}

// ValidateMarketHours valida se o mercado está aberto
func (v *BusinessValidator) ValidateMarketHours() error {
	// TODO: Implementar validação de horário
	// 1. Obter horário atual em EST
	// 2. Verificar se é dia útil (segunda a sábado para o evento)
	// 3. Verificar se está no horário 9:30-16:00 EST
	// 4. Verificar feriados da NYSE

	return nil
}

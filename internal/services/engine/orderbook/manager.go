package orderbook

import (
	"trading/internal/domain"
)

// OrderBook representa o livro de ofertas de um símbolo
type OrderBook struct {
	Symbol string          `json:"symbol"`
	Bids   []*domain.Order `json:"bids"` // Ordens de compra (preço decrescente)
	Asks   []*domain.Order `json:"asks"` // Ordens de venda (preço crescente)
}

// Manager gerencia livros de ofertas
type Manager struct {
	// TODO: Implementar campos necessários
}

// NewManager cria um novo manager de order book
func NewManager() *Manager {
	return &Manager{
		// TODO: Inicializar campos
	}
}

// GetOrderBook retorna o livro de ofertas de um símbolo
func (s *Manager) GetOrderBook(symbol string) *OrderBook {
	// TODO: Implementar lógica do order book
	// 1. Buscar ou criar livro para o símbolo
	// 2. Organizar bids por preço decrescente
	// 3. Organizar asks por preço crescente
	// 4. Retornar livro atualizado

	return &OrderBook{
		Symbol: symbol,
		Bids:   []*domain.Order{},
		Asks:   []*domain.Order{},
	}
}

// AddOrder adiciona uma ordem ao livro
func (s *Manager) AddOrder(order *domain.Order) {
	// TODO: Implementar adição de ordem
	// 1. Obter livro do símbolo
	// 2. Adicionar ordem na lista correta (bids ou asks)
	// 3. Ordenar lista por price-time priority
}

// RemoveOrder remove uma ordem do livro
func (s *Manager) RemoveOrder(symbol, orderID string) {
	// TODO: Implementar remoção de ordem
	// 1. Encontrar ordem no livro
	// 2. Remover da lista correspondente
}

// FindBestMatch encontra a melhor correspondência para uma ordem
func (s *Manager) FindBestMatch(order *domain.Order) *domain.Order {
	// TODO: Implementar busca de correspondência
	// 1. Para ordem BUY: buscar menor preço de venda (ask)
	// 2. Para ordem SELL: buscar maior preço de compra (bid)
	// 3. Verificar se preços são compatíveis
	// 4. Retornar ordem correspondente ou nil

	return nil
}

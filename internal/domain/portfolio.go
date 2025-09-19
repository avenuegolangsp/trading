package domain

import (
	"sync"
	"time"
)

// Portfolio representa o portfolio de um usuário
type Portfolio struct {
	UserID    string         `json:"user_id"`
	Cash      float64        `json:"cash"`
	Positions map[string]int `json:"positions"` // symbol -> quantity
	UpdatedAt time.Time      `json:"updated_at"`
	mutex     sync.RWMutex   `json:"-"`
}

// NewPortfolio cria um novo portfolio
func NewPortfolio(userID string, initialCash float64) *Portfolio {
	return &Portfolio{
		UserID:    userID,
		Cash:      initialCash,
		Positions: make(map[string]int),
		UpdatedAt: time.Now().UTC(),
	}
}

// GetCash retorna o saldo em dinheiro (thread-safe)
func (p *Portfolio) GetCash() float64 {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.Cash
}

// GetPosition retorna a posição de um símbolo (thread-safe)
func (p *Portfolio) GetPosition(symbol string) int {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.Positions[symbol]
}

// HasSufficientCash verifica se há saldo suficiente
func (p *Portfolio) HasSufficientCash(amount float64) bool {
	return p.GetCash() >= amount
}

// HasSufficientPosition verifica se há posição suficiente
func (p *Portfolio) HasSufficientPosition(symbol string, quantity int) bool {
	return p.GetPosition(symbol) >= quantity
}

// ExecuteBuy executa uma compra (debita cash, credita posição)
func (p *Portfolio) ExecuteBuy(symbol string, quantity int, price float64) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	cost := float64(quantity) * price

	if p.Cash < cost {
		return ErrInsufficientBalance
	}

	p.Cash -= cost
	p.Positions[symbol] += quantity
	p.UpdatedAt = time.Now().UTC()

	return nil
}

// ExecuteSell executa uma venda (credita cash, debita posição)
func (p *Portfolio) ExecuteSell(symbol string, quantity int, price float64) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.Positions[symbol] < quantity {
		return ErrInsufficientPosition
	}

	proceeds := float64(quantity) * price

	p.Cash += proceeds
	p.Positions[symbol] -= quantity

	// Remove posição se chegou a zero
	if p.Positions[symbol] == 0 {
		delete(p.Positions, symbol)
	}

	p.UpdatedAt = time.Now().UTC()

	return nil
}

// GetTotalValue calcula o valor total do portfolio (cash + posições)
func (p *Portfolio) GetTotalValue(stockPrices map[string]float64) float64 {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	total := p.Cash

	for symbol, quantity := range p.Positions {
		if price, exists := stockPrices[symbol]; exists {
			total += float64(quantity) * price
		}
	}

	return total
}

package domain

import (
	"time"
)

// Trade representa uma negociação executada
type Trade struct {
	ID         string    `json:"id"`
	BuyerID    string    `json:"buyer_id"`
	SellerID   string    `json:"seller_id"`
	Symbol     string    `json:"symbol"`
	Quantity   int       `json:"quantity"`
	Price      float64   `json:"price"`
	Value      float64   `json:"value"`
	ExecutedAt time.Time `json:"executed_at"`

	// Referências às ordens originais
	BuyOrderID  string `json:"buy_order_id"`
	SellOrderID string `json:"sell_order_id"`
}

// NewTrade cria uma nova negociação
func NewTrade(buyOrder, sellOrder *Order, quantity int, price float64) *Trade {
	value := float64(quantity) * price

	return &Trade{
		ID:          generateTradeID(),
		BuyerID:     buyOrder.UserID,
		SellerID:    sellOrder.UserID,
		Symbol:      buyOrder.Symbol,
		Quantity:    quantity,
		Price:       price,
		Value:       value,
		ExecutedAt:  time.Now().UTC(),
		BuyOrderID:  buyOrder.ID,
		SellOrderID: sellOrder.ID,
	}
}

// generateTradeID gera um ID único para a negociação
func generateTradeID() string {
	return "trade-" + time.Now().Format("20060102150405") + "-" + randomString(6)
}

package domain

import (
	"time"
)

// OrderSide representa o lado da ordem (compra ou venda)
type OrderSide string

const (
	BUY  OrderSide = "BUY"
	SELL OrderSide = "SELL"
)

// OrderStatus representa o status da ordem
type OrderStatus string

const (
	PENDING  OrderStatus = "PENDING"
	FILLED   OrderStatus = "FILLED"
	REJECTED OrderStatus = "REJECTED"
	PARTIAL  OrderStatus = "PARTIAL"
)

// Order representa uma ordem de compra ou venda
type Order struct {
	ID        string      `json:"id"`
	UserID    string      `json:"user_id"`
	Symbol    string      `json:"symbol"`
	Side      OrderSide   `json:"side"`
	Quantity  int         `json:"quantity"`
	Price     float64     `json:"price"`
	Status    OrderStatus `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`

	// Campos para matching
	RemainingQuantity int `json:"remaining_quantity,omitempty"`
}

// NewOrder cria uma nova ordem
func NewOrder(userID, symbol string, side OrderSide, quantity int, price float64) *Order {
	now := time.Now().UTC()
	return &Order{
		ID:                generateOrderID(),
		UserID:            userID,
		Symbol:            symbol,
		Side:              side,
		Quantity:          quantity,
		Price:             price,
		Status:            PENDING,
		CreatedAt:         now,
		UpdatedAt:         now,
		RemainingQuantity: quantity,
	}
}

// IsComplete verifica se a ordem foi completamente executada
func (o *Order) IsComplete() bool {
	return o.RemainingQuantity == 0
}

// GetValue retorna o valor total da ordem
func (o *Order) GetValue() float64 {
	return float64(o.Quantity) * o.Price
}

// generateOrderID gera um ID único para a ordem
func generateOrderID() string {
	return "order-" + time.Now().Format("20060102150405") + "-" + randomString(6)
}

// randomString gera uma string aleatória
func randomString(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[time.Now().UnixNano()%int64(len(chars))]
	}
	return string(result)
}

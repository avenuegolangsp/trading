package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sync"
	"time"
	"trading/internal/domain"

	"github.com/google/uuid"
)

// OrderRepository defines basic operations for orders
type OrderRepository interface {
	Create(order *domain.Order) error
	GetByID(id string) (*domain.Order, error)
	Update(order *domain.Order) error
	Delete(id string) error
	List() ([]*domain.Order, error)
}

// TradeRepository defines basic operations for trades
type TradeRepository interface {
	Create(trade *domain.Trade) error
	GetByID(id string) (*domain.Trade, error)
	List() ([]*domain.Trade, error)
}

// UserRepository defines basic operations for users
type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id string) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id string) error
	List() ([]*domain.User, error)
}

// StockRepository defines basic operations for stocks
type StockRepository interface {
	Create(stock *domain.Stock) error
	GetBySymbol(symbol string) (*domain.Stock, error)
	Update(stock *domain.Stock) error
	Delete(symbol string) error
	List() ([]*domain.Stock, error)
}

// InMemoryOrderRepository is an in-memory implementation of OrderRepository
type InMemoryOrderRepository struct {
	mu     sync.RWMutex
	orders map[string]*domain.Order
}

func NewInMemoryOrderRepository() *InMemoryOrderRepository {
	return &InMemoryOrderRepository{
		orders: make(map[string]*domain.Order),
	}
}

func (r *InMemoryOrderRepository) Create(order *domain.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if order.ID == "" {
		order.ID = uuid.New().String()
	}
	r.orders[order.ID] = order
	return nil
}

func (r *InMemoryOrderRepository) GetByID(id string) (*domain.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	order, ok := r.orders[id]
	if !ok {
		return nil, errors.New("order not found")
	}
	return order, nil
}

func (r *InMemoryOrderRepository) Update(order *domain.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.orders[order.ID]; !ok {
		return errors.New("order not found")
	}
	r.orders[order.ID] = order
	return nil
}

func (r *InMemoryOrderRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.orders[id]; !ok {
		return errors.New("order not found")
	}
	delete(r.orders, id)
	return nil
}

func (r *InMemoryOrderRepository) List() ([]*domain.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []*domain.Order
	for _, o := range r.orders {
		result = append(result, o)
	}
	return result, nil
}

// InMemoryTradeRepository is an in-memory implementation of TradeRepository
type InMemoryTradeRepository struct {
	mu     sync.RWMutex
	trades map[string]*domain.Trade
}

func NewInMemoryTradeRepository() *InMemoryTradeRepository {
	return &InMemoryTradeRepository{
		trades: make(map[string]*domain.Trade),
	}
}

func (r *InMemoryTradeRepository) Create(trade *domain.Trade) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if trade.ID == "" {
		trade.ID = uuid.New().String()
	}
	r.trades[trade.ID] = trade
	return nil
}

func (r *InMemoryTradeRepository) GetByID(id string) (*domain.Trade, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	trade, ok := r.trades[id]
	if !ok {
		return nil, errors.New("trade not found")
	}
	return trade, nil
}

func (r *InMemoryTradeRepository) List() ([]*domain.Trade, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []*domain.Trade
	for _, t := range r.trades {
		result = append(result, t)
	}
	return result, nil
}

// InMemoryUserRepository is an in-memory implementation of UserRepository
type InMemoryUserRepository struct {
	mu    sync.RWMutex
	users map[string]*domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*domain.User),
	}
}

func (r *InMemoryUserRepository) Create(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if user.ID == "" {
		user.ID = uuid.New().String()
	}
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) GetByID(id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, ok := r.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *InMemoryUserRepository) Update(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.users[user.ID]; !ok {
		return errors.New("user not found")
	}
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.users[id]; !ok {
		return errors.New("user not found")
	}
	delete(r.users, id)
	return nil
}

func (r *InMemoryUserRepository) List() ([]*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []*domain.User
	for _, u := range r.users {
		result = append(result, u)
	}
	return result, nil
}

// InMemoryStockRepository is an in-memory implementation of StockRepository
type InMemoryStockRepository struct {
	mu     sync.RWMutex
	stocks map[string]*domain.Stock
}

func NewInMemoryStockRepository() *InMemoryStockRepository {
	return &InMemoryStockRepository{
		stocks: make(map[string]*domain.Stock),
	}
}

func (r *InMemoryStockRepository) Create(stock *domain.Stock) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if stock.Symbol == "" {
		return errors.New("stock symbol required")
	}
	r.stocks[stock.Symbol] = stock
	return nil
}

func (r *InMemoryStockRepository) GetBySymbol(symbol string) (*domain.Stock, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	stock, ok := r.stocks[symbol]
	if !ok {
		return nil, errors.New("stock not found")
	}
	return stock, nil
}

func (r *InMemoryStockRepository) Update(stock *domain.Stock) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if stock.Symbol == "" {
		return errors.New("stock symbol required")
	}
	if _, ok := r.stocks[stock.Symbol]; !ok {
		return errors.New("stock not found")
	}
	r.stocks[stock.Symbol] = stock
	return nil
}

func (r *InMemoryStockRepository) Delete(symbol string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.stocks[symbol]; !ok {
		return errors.New("stock not found")
	}
	delete(r.stocks, symbol)
	return nil
}

func (r *InMemoryStockRepository) List() ([]*domain.Stock, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []*domain.Stock
	for _, s := range r.stocks {
		result = append(result, s)
	}
	return result, nil
}

// NewInMemoryStockRepositoryFromJSON loads stocks from a JSON file (e.g., /data/stocks.json)
func NewInMemoryStockRepositoryFromJSON(path string) (*InMemoryStockRepository, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var data struct {
		Stocks map[string]struct {
			Company     string  `json:"company"`
			Sector      string  `json:"sector"`
			MinPrice    float64 `json:"min_price"`
			MarketCap   string  `json:"market_cap"`
			Description string  `json:"description"`
		} `json:"stocks"`
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	repo := NewInMemoryStockRepository()
	for symbol, s := range data.Stocks {
		repo.stocks[symbol] = &domain.Stock{
			Symbol:      symbol,
			Company:     s.Company,
			Sector:      s.Sector,
			MinPrice:    s.MinPrice,
			MarketCap:   s.MarketCap,
			Description: s.Description,
		}
	}
	return repo, nil
}

// NewInMemoryUserRepositoryFromJSON loads users from a JSON file (e.g., /data/users.json)
func NewInMemoryUserRepositoryFromJSON(path string) (*InMemoryUserRepository, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var data struct {
		Users []struct {
			ID        string  `json:"id"`
			Name      string  `json:"name"`
			Email     string  `json:"email"`
			Cash      float64 `json:"cash"`
			Status    string  `json:"status"`
			CreatedAt string  `json:"created_at"`
		} `json:"users"`
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	repo := NewInMemoryUserRepository()
	for _, u := range data.Users {
		t, _ := time.Parse(time.RFC3339, u.CreatedAt)
		repo.users[u.ID] = &domain.User{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			Cash:      u.Cash,
			Status:    u.Status,
			CreatedAt: t,
		}
	}
	return repo, nil
}

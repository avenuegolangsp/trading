package domain

import "time"

// User represents a trading user (simplified for repository)
type User struct {
	ID        string
	Name      string
	Email     string
	Cash      float64
	Status    string
	CreatedAt time.Time
}

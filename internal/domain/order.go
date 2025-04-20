package domain

import "time"

type Order struct {
	ID        int64
	UserID    int64
	Items     []OrderItem
	CreatedAt time.Time
}

type OrderItem struct {
	ProductID int64
	Quantity  int
	Price     float64
}

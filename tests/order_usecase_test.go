package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/Helltale/tz-telecom/internal/domain"
	"github.com/Helltale/tz-telecom/internal/usecase"
	"github.com/stretchr/testify/assert"
)

type mockOrderRepo struct {
	stock        bool
	stockErr     error
	price        float64
	priceErr     error
	createCalled bool
	createErr    error
}

func (m *mockOrderRepo) ProductInStock(ctx context.Context, id int64, qty int) (bool, error) {
	return m.stock, m.stockErr
}

func (m *mockOrderRepo) GetProductPrice(ctx context.Context, id int64) (float64, error) {
	return m.price, m.priceErr
}

func (m *mockOrderRepo) Create(ctx context.Context, o *domain.Order) error {
	m.createCalled = true
	return m.createErr
}

func TestCreateOrder(t *testing.T) {
	tests := []struct {
		name        string
		repo        *mockOrderRepo
		items       []domain.OrderItem
		expectError string
	}{
		{
			name: "valid order",
			repo: &mockOrderRepo{
				stock: true, price: 123.45,
			},
			items: []domain.OrderItem{{ProductID: 1, Quantity: 2}},
		},
		{
			name: "out of stock",
			repo: &mockOrderRepo{
				stock: false,
			},
			items:       []domain.OrderItem{{ProductID: 1, Quantity: 1}},
			expectError: "product out of stock",
		},
		{
			name: "price fetch error",
			repo: &mockOrderRepo{
				stock: true, priceErr: errors.New("price error"),
			},
			items:       []domain.OrderItem{{ProductID: 1, Quantity: 1}},
			expectError: "price error",
		},
		{
			name: "create failed",
			repo: &mockOrderRepo{
				stock: true, price: 99.99, createErr: errors.New("create failed"),
			},
			items:       []domain.OrderItem{{ProductID: 1, Quantity: 1}},
			expectError: "create failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := usecase.NewOrderUseCase(tt.repo)
			err := uc.CreateOrder(context.Background(), 42, tt.items)
			if tt.expectError != "" {
				assert.EqualError(t, err, tt.expectError)
			} else {
				assert.NoError(t, err)
				assert.True(t, tt.repo.createCalled)
			}
		})
	}
}

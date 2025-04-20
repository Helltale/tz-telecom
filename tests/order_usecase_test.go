package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Helltale/tz-telecom/internal/domain"
	"github.com/Helltale/tz-telecom/internal/usecase"
	"github.com/stretchr/testify/assert"
)

type mockOrderRepo struct {
	inStockFunc       func(id int64, qty int) bool
	priceFunc         func(id int64) (float64, error)
	createShouldError bool
}

func (m *mockOrderRepo) ProductInStock(ctx context.Context, id int64, qty int) (bool, error) {
	if m.inStockFunc != nil {
		return m.inStockFunc(id, qty), nil
	}
	return true, nil
}

func (m *mockOrderRepo) GetProductPrice(ctx context.Context, id int64) (float64, error) {
	if m.priceFunc != nil {
		return m.priceFunc(id)
	}
	return 100.0, nil
}

func (m *mockOrderRepo) Create(ctx context.Context, order *domain.Order) error {
	if m.createShouldError {
		return errors.New("create failed")
	}
	return nil
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
			repo: &mockOrderRepo{},
			items: []domain.OrderItem{
				{ProductID: 1, Quantity: 2},
			},
		},
		{
			name: "product out of stock",
			repo: &mockOrderRepo{
				inStockFunc: func(id int64, qty int) bool {
					return false
				},
			},
			items: []domain.OrderItem{
				{ProductID: 2, Quantity: 5},
			},
			expectError: "product out of stock",
		},
		{
			name: "price fetch error",
			repo: &mockOrderRepo{
				priceFunc: func(id int64) (float64, error) {
					return 0, errors.New("price not found")
				},
			},
			items: []domain.OrderItem{
				{ProductID: 3, Quantity: 1},
			},
			expectError: "price not found",
		},
		{
			name: "create error",
			repo: &mockOrderRepo{
				createShouldError: true,
			},
			items: []domain.OrderItem{
				{ProductID: 1, Quantity: 1},
			},
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
			}
		})
	}
}

package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Helltale/tz-telecom/internal/domain"
)

type OrderRepository interface {
	ProductInStock(ctx context.Context, id int64, qty int) (bool, error)
	GetProductPrice(ctx context.Context, id int64) (float64, error)
	Create(ctx context.Context, order *domain.Order) error
}

type OrderUseCase struct {
	repo OrderRepository
}

func NewOrderUseCase(r OrderRepository) *OrderUseCase {
	return &OrderUseCase{repo: r}
}

func (uc *OrderUseCase) CreateOrder(ctx context.Context, userID int64, items []domain.OrderItem) error {
	enrichedItems := make([]domain.OrderItem, 0, len(items))

	for _, item := range items {
		ok, err := uc.repo.ProductInStock(ctx, item.ProductID, item.Quantity)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("product out of stock")
		}

		price, err := uc.repo.GetProductPrice(ctx, item.ProductID)
		if err != nil {
			return err
		}

		enrichedItems = append(enrichedItems, domain.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     price,
		})
	}

	order := &domain.Order{
		UserID:    userID,
		Items:     enrichedItems,
		CreatedAt: time.Now(),
	}
	return uc.repo.Create(ctx, order)
}

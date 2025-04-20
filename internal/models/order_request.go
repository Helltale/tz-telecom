package models

import "github.com/Helltale/tz-telecom/internal/domain"

type OrderRequest struct {
	UserID int64            `json:"user_id"`
	Items  []OrderItemInput `json:"items"`
}

type OrderItemInput struct {
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

func (r OrderRequest) ToDomainItems() []domain.OrderItem {
	items := make([]domain.OrderItem, 0, len(r.Items))
	for _, i := range r.Items {
		items = append(items, domain.OrderItem{
			ProductID: i.ProductID,
			Quantity:  i.Quantity,
			// price will be filled in usecase, not here
		})
	}
	return items
}

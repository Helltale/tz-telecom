package postgresrepo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Helltale/tz-telecom/internal/domain"
)

type OrderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) ProductInStock(ctx context.Context, id int64, qty int) (bool, error) {
	var stock int
	err := r.db.QueryRowContext(ctx, `
		SELECT quantity FROM products WHERE id = $1
	`, id).Scan(&stock)

	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return stock >= qty, nil
}

func (r *OrderRepo) Create(ctx context.Context, order *domain.Order) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// inserting an order
	err = tx.QueryRowContext(ctx, `
		INSERT INTO orders (user_id, created_at)
		VALUES ($1, $2)
		RETURNING id
	`, order.UserID, time.Now()).Scan(&order.ID)
	if err != nil {
		return err
	}

	// inserting products and updating the warehouse
	for _, item := range order.Items {
		// check and update warehouse
		res, err := tx.ExecContext(ctx, `
			UPDATE products
			SET quantity = quantity - $1
			WHERE id = $2 AND quantity >= $1
		`, item.Quantity, item.ProductID)
		if err != nil {
			return err
		}
		affected, _ := res.RowsAffected()
		if affected == 0 {
			return errors.New("not enough stock")
		}

		// insert record into order_items
		_, err = tx.ExecContext(ctx, `
			INSERT INTO order_items (order_id, product_id, quantity, price)
			VALUES ($1, $2, $3, $4)
		`, order.ID, item.ProductID, item.Quantity, item.Price)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *OrderRepo) GetProductPrice(ctx context.Context, id int64) (float64, error) {
	var price float64
	err := r.db.QueryRowContext(ctx, `
		SELECT price FROM products WHERE id = $1
	`, id).Scan(&price)
	if err != nil {
		return 0, err
	}
	return price, nil
}

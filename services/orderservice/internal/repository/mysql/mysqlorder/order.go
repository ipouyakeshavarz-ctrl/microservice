package mysqlorder

import (
	"context"
	"orderapp/internal/domain"
)

func (d *DB) ExistsByCheckoutID(ctx context.Context, checkoutID string) (bool, error) {

	query := `
	SELECT EXISTS(
		SELECT 1 FROM orders WHERE checkout_id = ?
	)`

	var exists bool

	err := d.conn.Conn().QueryRowContext(ctx, query, checkoutID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (d *DB) Create(ctx context.Context, order *domain.Order) error {

	tx, err := d.conn.Conn().BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	query := `
	INSERT INTO orders (checkout_id, user_id, status, created_at)
	VALUES (?, ?, ?, ?)
	`

	res, err := tx.ExecContext(
		ctx,
		query,
		order.CheckoutID,
		order.UserID,
		order.Status,
		order.CreatedAt,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	order.ID = uint(id)

	for _, item := range order.Items {

		_, err = tx.ExecContext(ctx, `
			INSERT INTO order_items (order_id, product_id, quantity)
			VALUES (?, ?, ?)
		`,
			order.ID,
			item.ProductID,
			item.Quantity,
		)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

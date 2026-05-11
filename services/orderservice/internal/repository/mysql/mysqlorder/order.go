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

func (d *DB) GetByID(ctx context.Context, id uint) (*domain.Order, error) {
	query := `
	SELECT id, checkout_id, user_id, status, created_at
	FROM orders
	WHERE id = ?
	LIMIT 1
	`

	var o domain.Order

	err := d.conn.Conn().QueryRowContext(ctx, query, id).Scan(
		&o.ID,
		&o.CheckoutID,
		&o.UserID,
		&o.Status,
		&o.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	items, err := d.getItemsByOrderID(ctx, o.ID)
	if err != nil {
		return nil, err
	}

	o.Items = items

	return &o, nil
}

func (d *DB) ListByUserID(ctx context.Context, userID uint) ([]domain.Order, error) {
	query := `
	SELECT id, checkout_id, user_id, status, created_at
	FROM orders
	WHERE user_id = ?
	ORDER BY created_at DESC
	`

	rows, err := d.conn.Conn().QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order

	for rows.Next() {
		var o domain.Order

		err := rows.Scan(
			&o.ID,
			&o.CheckoutID,
			&o.UserID,
			&o.Status,
			&o.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		items, err := d.getItemsByOrderID(ctx, o.ID)
		if err != nil {
			return nil, err
		}

		o.Items = items

		orders = append(orders, o)
	}

	return orders, rows.Err()
}

func (d *DB) getItemsByOrderID(
	ctx context.Context,
	orderID uint,
) ([]domain.OrderItem, error) {

	query := `
	SELECT product_id, quantity
	FROM order_items
	WHERE order_id = ?
	`

	rows, err := d.conn.Conn().QueryContext(ctx, query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.OrderItem

	for rows.Next() {
		var item domain.OrderItem

		err := rows.Scan(
			&item.ProductID,
			&item.Quantity,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, rows.Err()
}

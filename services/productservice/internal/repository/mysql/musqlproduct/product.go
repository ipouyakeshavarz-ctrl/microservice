package mysqlproduct

import (
	"context"
	"database/sql"
	"myapp/pkg/errmsg"
	"myapp/pkg/richerror"
	"productapp/internal/entity"
)

func (d *DB) Create(ctx context.Context, p entity.Product) (*entity.Product, error) {
	const op = "ProductRepository.Create"

	query := `
		INSERT INTO products (store_id, name, description, category, price, stock, sku, image_url, is_active, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())
	`
	res, err := d.conn.Conn().ExecContext(ctx, query,
		p.StoreID,
		p.Name,
		p.Description,
		p.Category,
		p.Price,
		p.Stock,
		p.SKU,
		p.ImageURL,
		p.IsActive,
	)
	if err != nil {
		return nil, richerror.New(op).WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithErr(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, richerror.New(op).WithKind(richerror.KindUnexpected).WithErr(err)
	}
	p.ID = uint(id)
	return &p, nil
}

// Update modifies an existing product.
func (d *DB) Update(ctx context.Context, p entity.Product) (*entity.Product, error) {
	const op = "ProductRepository.Update"

	query := `
		UPDATE products
		SET name=?, description=?, category=?, price=?, stock=?, sku=?, image_url=?, is_active=?, updated_at=NOW()
		WHERE id=? AND store_id=?
	`
	res, err := d.conn.Conn().ExecContext(ctx, query,
		p.Name,
		p.Description,
		p.Category,
		p.Price,
		p.Stock,
		p.SKU,
		p.ImageURL,
		p.IsActive,
		p.ID,
		p.StoreID,
	)
	if err != nil {
		return nil, richerror.New(op).WithKind(richerror.KindUnexpected).
			WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, richerror.New(op).WithKind(richerror.KindUnexpected).WithErr(err)
	}
	if rows == 0 {
		return nil, richerror.New(op).WithKind(richerror.KindNotFound).
			WithMessage(errmsg.ErrorMsgProductNotFound)
	}
	return &p, nil
}

// Delete removes a product by ID.
func (d *DB) Delete(ctx context.Context, id uint) error {
	const op = "ProductRepository.Delete"

	query := `DELETE FROM products WHERE id=?`
	res, err := d.conn.Conn().ExecContext(ctx, query, id)
	if err != nil {
		return richerror.New(op).WithKind(richerror.KindUnexpected).
			WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return richerror.New(op).WithKind(richerror.KindNotFound).
			WithMessage(errmsg.ErrorMsgProductNotFound)
	}
	return nil
}

// GetByID fetches a product by its ID.
func (d *DB) GetByID(ctx context.Context, id uint) (*entity.Product, error) {
	const op = "ProductRepository.GetByID"
	query := `
		SELECT id, store_id, name, description, category, price, stock, sku, image_url, is_active, created_at, updated_at
		FROM products WHERE id=?
	`
	row := d.conn.Conn().QueryRowContext(ctx, query, id)

	var p entity.Product
	err := row.Scan(
		&p.ID,
		&p.StoreID,
		&p.Name,
		&p.Description,
		&p.Category,
		&p.Price,
		&p.Stock,
		&p.SKU,
		&p.ImageURL,
		&p.IsActive,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, richerror.New(op).WithKind(richerror.KindNotFound).
			WithMessage(errmsg.ErrorMsgProductNotFound)
	}
	if err != nil {
		return nil, richerror.New(op).WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithErr(err)
	}
	return &p, nil
}

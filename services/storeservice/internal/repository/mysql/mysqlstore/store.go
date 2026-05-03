package mysqlstore

import (
	"context"
	"myapp/pkg/errmsg"
	"myapp/pkg/richerror"
	"storeapp/internal/domain"
	"storeapp/internal/param"
	"storeapp/internal/repository/mysql"
	"time"
)

func (d DB) CreateStore(ctx context.Context, s domain.Store) (*domain.Store, error) {
	const op = "StoreRepo.CreateStore"
	query := `
		INSERT INTO stores (user_id, name, description, logo_url, street, city, province, postal_code, phone_number, is_active, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	result, err := d.conn.Conn().ExecContext(ctx, query,
		s.UserID, s.Name, s.Description, s.LogoURL, s.Address.Street, s.Address.City,
		s.Address.Province, s.Address.PostalCode, s.PhoneNumber, s.IsActive)

	if err != nil {
		return nil, richerror.New(op).WithKind(richerror.KindUnexpected).
			WithErr(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, richerror.New(op).WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithErr(err)
	}
	s.ID = uint(id)
	s.IsActive = true

	return &s, nil

}

func (d DB) UpdateStore(ctx context.Context, s domain.Store) (*domain.Store, error) {
	const op = "StoreRepo.UpdateStore"
	query := `
		UPDATE stores SET user_id=?, name=?, description=?, logo_url=?, street=?, city=?, province=?, postal_code=?, phone_number=?, is_active=?, updated_at=CURRENT_TIMESTAMP
		WHERE id=?
	`
	_, err := d.conn.Conn().ExecContext(ctx, query,
		s.UserID, s.Name, s.Description, s.LogoURL,
		s.Address.Street, s.Address.City, s.Address.Province, s.Address.PostalCode,
		s.PhoneNumber, s.IsActive, s.ID)
	if err != nil {
		return nil, richerror.New(op).WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithErr(err)
	}
	return &s, nil
}

func (d DB) DeleteStore(ctx context.Context, id uint) error {
	const op = "StoreRepo.DeleteStore"
	query := `DELETE FROM stores WHERE id=?`

	_, err := d.conn.Conn().ExecContext(ctx, query, id)

	return richerror.New(op).WithKind(richerror.KindUnexpected).
		WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithErr(err)

}

func (d DB) GetStoreByID(ctx context.Context, id uint) (*domain.Store, error) {
	const op = "StoreRepo.GetStoreByID"
	query := `SELECT id, user_id, name, description, logo_url, street, city, province, postal_code, phone_number, is_active, created_at, updated_at FROM stores WHERE id=?`
	row := d.conn.Conn().QueryRowContext(ctx, query, id)

	s, err := scanStore(row)

	if err != nil {
		return nil, richerror.New(op).WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithErr(err)
	}
	return &s, nil
}

func (d DB) ListStoresByUser(ctx context.Context, userID uint) ([]param.StoreInfo, error) {
	const op = "StoreRepo.ListStoresByUser"

	query := `SELECT id, user_id, name, description, logo_url, street, city, province, postal_code, phone_number, is_active, created_at, updated_at FROM stores WHERE user_id=?`
	rows, err := d.conn.Conn().QueryContext(ctx, query, userID)
	if err != nil {
		return nil, richerror.New(op).WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithErr(err)
	}

	var stores []param.StoreInfo
	for rows.Next() {

		s, err := scanStore2(rows)

		if err != nil {
			return nil, richerror.New(op).WithKind(richerror.KindUnexpected).
				WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithErr(err)
		}
		stores = append(stores, s)
	}
	return stores, richerror.New(op).WithKind(richerror.KindUnexpected).
		WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithErr(err)
}

func scanStore(scanner mysql.Scanner) (domain.Store, error) {
	const op = "storeRepo.scanStore"
	var createdAt time.Time
	var updatedAt time.Time
	var store domain.Store

	err := scanner.Scan(&store.ID, &store.UserID, &store.Name, &store.Description, &store.LogoURL,
		&store.Address.Street, &store.Address.City, &store.Address.Province, &store.Address.PostalCode,
		&store.PhoneNumber, &store.IsActive, &createdAt, &updatedAt)

	return store, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgCantScanQueryResult)
}

func scanStore2(scanner mysql.Scanner) (param.StoreInfo, error) {
	var createdAt time.Time
	var updatedAt time.Time
	var store param.StoreInfo

	err := scanner.Scan(&store.ID, &store.UserID, &store.Name, &store.Description, &store.LogoURL,
		&store.Address.Street, &store.Address.City, &store.Address.Province, &store.Address.PostalCode,
		&store.PhoneNumber, &store.IsActive, &createdAt, &updatedAt)

	return store, err
}

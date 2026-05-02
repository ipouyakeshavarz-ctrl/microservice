package mysqluser

import (
	"context"
	"database/sql"
	"fmt"
	"myapp/pkg/errmsg"
	"myapp/pkg/richerror"
	"time"
	"userapp/internal/domain"
	"userapp/internal/repository/mysql"
)

func (d *DB) IsPhoneNumberUnique(ctx context.Context, phoneNumber string) (bool, error) {
	const op = "mysql.IsPhoneNumberUnique"

	row := d.conn.Conn().QueryRowContext(ctx, `select * from users where phone_number = ?`, phoneNumber)

	_, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	return false, nil
}

func (d *DB) Register(ctx context.Context, u domain.User) (domain.User, error) {
	res, err := d.conn.Conn().ExecContext(ctx, `insert into users(name, phone_number, password, role) values(?, ?, ?, ?)`,
		u.Name, u.PhoneNumber, u.Password, u.Role.String())
	if err != nil {
		return domain.User{}, fmt.Errorf("can't execute command: %w", err)
	}

	// error is always nil
	id, _ := res.LastInsertId()
	u.ID = uint(id)

	return u, nil
}

func (d *DB) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (domain.User, error) {
	const op = "mysql.GetUserByPhoneNumber"

	row := d.conn.Conn().QueryRowContext(ctx, `select * from users where phone_number = ?`, phoneNumber)

	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}

		// TODO - log unexpected error for better observability
		return domain.User{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	return user, nil
}

func (d *DB) GetUserByID(ctx context.Context, userID uint) (domain.User, error) {
	const op = "mysql.GetUserByID"

	row := d.conn.Conn().QueryRowContext(ctx, `select * from users where id = ?`, userID)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, richerror.New(op).WithErr(err).
				WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}

		return domain.User{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	return user, nil
}

func scanUser(scanner mysql.Scanner) (domain.User, error) {
	var createdAt time.Time
	var user domain.User

	var roleStr string

	err := scanner.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt, &user.Password, &roleStr)

	user.Role = domain.MapToRoleEntity(roleStr)

	return user, err
}

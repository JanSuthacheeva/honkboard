package models

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type ValidationCode struct {
	ID      int
	Code    int
	Expires time.Time
	Type    string
}

type ValidationCodeModel struct {
	DB *sql.DB
}

func (m *ValidationCodeModel) Insert(userId, code int, codeType string) (int, error) {
	query := `INSERT INTO validation_codes (user_id, code, type, expires)
		VALUES(?, ?, ?, DATE_ADD(UTC_TIMESTAMP(), INTERVAL 5 MINUTES))`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, userId, code, codeType)
	if err != nil {
		return 0, nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *ValidationCodeModel) Get(code int) (ValidationCode, error) {
	query := `SELECT id, code, type, expires FROM validation_codes
		WHERE code = ?
		AND expires > UTC_TIMESTAMP()`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var validationCode ValidationCode

	err := m.DB.QueryRowContext(ctx, query, code).Scan(&validationCode.ID, &validationCode.Code, &validationCode.Type, &validationCode.Expires)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ValidationCode{}, ErrNoRecord
		default:
			return ValidationCode{}, err
		}
	}

	return validationCode, nil
}

func (m *ValidationCodeModel) Delete(id int) error {
	query := `DELETE FROM validation_codes
		WHERE id = ?`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNoRecord
	}

	return nil
}

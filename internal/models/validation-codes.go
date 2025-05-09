package models

import (
	"context"
	"database/sql"
	"errors"
	"math/rand/v2"
	"time"
)

type ValidationCode struct {
	ID      int
	UserID  int
	Code    int
	Expires time.Time
	Type    string
}

type ValidationCodeModel struct {
	DB *sql.DB
}

func (m *ValidationCodeModel) Insert(userId int, codeType string) (ValidationCode, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		return ValidationCode{}, err
	}

	defer tx.Rollback()

	insertQuery := `INSERT INTO validation_codes (user_id, code, type, expires)
		VALUES(?, ?, ?, DATE_ADD(UTC_TIMESTAMP(), INTERVAL 5 MINUTE))`

	result, err := m.DB.ExecContext(ctx, insertQuery, userId, createCode(), codeType)
	if err != nil {
		return ValidationCode{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return ValidationCode{}, err
	}

	var validationCode ValidationCode

	getQuery := `SELECT id, user_id, code, type, expires FROM validation_codes
	WHERE id = ?`

	err = tx.QueryRowContext(ctx, getQuery, id).Scan(&validationCode.ID, &validationCode.UserID, &validationCode.Code, &validationCode.Type, &validationCode.Expires)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ValidationCode{}, ErrNoRecord
		default:
			return ValidationCode{}, err
		}
	}

	if err = tx.Commit(); err != nil {
		return ValidationCode{}, err
	}

	return validationCode, nil
}

func (m *ValidationCodeModel) GetByCode(code int) (ValidationCode, error) {

	query := `SELECT id, user_id, code, type, expires FROM validation_codes
		WHERE code = ?
		AND expires > UTC_TIMESTAMP()`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var validationCode ValidationCode

	err := m.DB.QueryRowContext(ctx, query, code).Scan(&validationCode.ID, &validationCode.UserID, &validationCode.Code, &validationCode.Type, &validationCode.Expires)
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

func createCode() int {
	randInt := rand.IntN(999999-100000) + 100000
	return randInt
}

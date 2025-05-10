package models

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"time"
)

type PasswordResetToken struct {
	ID      int
	Email   string
	Token   string
	Expires time.Time
}

type PasswordResetTokenModel struct {
	DB *sql.DB
}

func generateToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (m *PasswordResetTokenModel) Insert(email string) (PasswordResetToken, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		return PasswordResetToken{}, err
	}

	defer tx.Rollback()

	insertQuery := `INSERT INTO password_reset_tokens (email, token, expires)
		VALUES(?, ?, DATE_ADD(UTC_TIMESTAMP(), INTERVAL 10 MINUTE))`

	token, err := generateToken()
	if err != nil {
		return PasswordResetToken{}, err
	}
	result, err := tx.ExecContext(ctx, insertQuery, email, token)
	if err != nil {
		return PasswordResetToken{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return PasswordResetToken{}, err
	}

	var passwordResetToken PasswordResetToken

	getQuery := `SELECT id, email, token, expires FROM password_reset_tokens
	WHERE id = ?`

	err = tx.QueryRowContext(ctx, getQuery, id).Scan(&passwordResetToken.ID, &passwordResetToken.Email, &passwordResetToken.Token, &passwordResetToken.Expires)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return PasswordResetToken{}, ErrNoRecord
		default:
			return PasswordResetToken{}, err
		}
	}

	if err = tx.Commit(); err != nil {
		return PasswordResetToken{}, err
	}

	return passwordResetToken, nil
}

func (m *PasswordResetTokenModel) Get(token string) (PasswordResetToken, error) {
	query := `SELECT id, email, token, expires FROM password_reset_tokens
	WHERE token = ?
	AND expires > UTC_TIMESTAMP()`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var passwordResetToken PasswordResetToken

	err := m.DB.QueryRowContext(ctx, query, token).Scan(&passwordResetToken.ID, &passwordResetToken.Email, &passwordResetToken.Token, &passwordResetToken.Expires)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return PasswordResetToken{}, ErrNoRecord
		default:
			return PasswordResetToken{}, err
		}
	}

	return passwordResetToken, nil
}

func (m *PasswordResetTokenModel) Delete(id int) error {
	query := `DELETE FROM password_reset_tokens
		WHERE id = ?`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
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

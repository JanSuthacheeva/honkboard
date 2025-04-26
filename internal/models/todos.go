package models

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jansuthacheeva/honkboard/internal/enums"
)

type TodoModelInterface interface {
	Insert(title, typeString string) (int, error)
	Delete(id int) error
	DeleteCompleted(listType string) error
	ToggleStatus(id int) (Todo, error)
	GetAll(listType string) ([]Todo, error)
}

type Todo struct {
	ID      int
	Title   string
	Type    enums.TodoType
	Status  enums.TodoStatus
	Created time.Time
}

type TodoModel struct {
	DB *sql.DB
}

func (m *TodoModel) Insert(title, typeString string) (int, error) {

	todoType, ok := enums.ParseTodoType(typeString)
	if !ok {
		return 0, ErrUnknownType
	}

	query := `INSERT INTO todos (title, status, type, created)
		VALUES(?, "not done", ?, UTC_TIMESTAMP())`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, title, todoType.String())
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *TodoModel) Delete(id int) error {
	if id <= 0 {
		return ErrNoRecord
	}

	query := `DELETE FROM todos
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

func (m *TodoModel) DeleteCompleted(listType string) error {

	todoType, ok := enums.ParseTodoType(listType)
	if !ok {
		return ErrUnknownType
	}

	query := `DELETE FROM todos
			WHERE type = ?
			AND status = "done"`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, todoType.String())
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

func (m *TodoModel) ToggleStatus(id int) (Todo, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		return Todo{}, err
	}

	defer tx.Rollback()

	statusQuery := `SELECT status FROM todos
		WHERE id = ?`

	var statusString string

	row := tx.QueryRowContext(ctx, statusQuery, id)
	err = row.Scan(&statusString)

	oldStatus, ok := enums.ParseTodoStatus(statusString)
	if !ok {
		return Todo{}, ErrUnknownStatus
	}

	var status enums.TodoStatus
	switch {
	case oldStatus == enums.TodoStatusDone:
		status = enums.TodoStatusNotDone
	case oldStatus == enums.TodoStatusNotDone:
		status = enums.TodoStatusDone
	default:
		return Todo{}, ErrUnknownStatus
	}

	updateQuery := `UPDATE todos SET status = ?
		WHERE id = ?`

	result, err := tx.ExecContext(ctx, updateQuery, status.String(), id)
	if err != nil {
		return Todo{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Todo{}, err
	}

	if rowsAffected == 0 {
		return Todo{}, ErrNoRecord
	}

	getQuery := `SELECT id, title, status, created FROM todos WHERE id = ?`

	row = tx.QueryRowContext(ctx, getQuery, id)

	var todo Todo
	err = row.Scan(&todo.ID, &todo.Title, &todo.Status, &todo.Created)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return Todo{}, ErrNoRecord
		default:
			return Todo{}, err
		}
	}
	if err = tx.Commit(); err != nil {
		return Todo{}, err
	}

	return todo, nil
}

func (m *TodoModel) GetAll(listType string) ([]Todo, error) {

	todoType, ok := enums.ParseTodoType(listType)
	if !ok {
		return nil, ErrUnknownType
	}

	query := `SELECT id, title, status, created FROM todos
			WHERE type = ?`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, todoType.String())
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var t Todo

		var statusString *string

		err = rows.Scan(&t.ID, &t.Title, &statusString, &t.Created)
		if err != nil {
			return nil, err
		}

		status, ok := enums.ParseTodoStatus(*statusString)
		if !ok {
			return nil, ErrUnknownStatus
		}
		t.Status = status

		todos = append(todos, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

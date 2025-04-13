package models

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Todo struct {
	ID      int
	Title   string
	Type    string
	Status  TodoStatus
	Created time.Time
}

type TodoStatus string

const (
	StatusDone    TodoStatus = "done"
	StatusNotDone TodoStatus = "not done"
)

var statusName = map[TodoStatus]string{
	StatusDone:    "done",
	StatusNotDone: "not done",
}

func (ts TodoStatus) String() string {
	return statusName[ts]
}

func parseStatus(s string) (TodoStatus, bool) {
	for status, name := range statusName {
		if name == s {
			return status, true
		}
	}
	return "", false
}

type TodoModel struct {
	DB *sql.DB
}

func (m *TodoModel) Insert(todo Todo) error {
	return nil
}

func (m *TodoModel) Delete(id int) error {
	return nil
}

func (m *TodoModel) Update(id int, title, status string) (*Todo, error) {
	return nil, nil
}

func (m *TodoModel) GetAll(todoType string) ([]Todo, error) {
	query := `SELECT id, title, status, created FROM todos
			WHERE type = ?`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, todoType)
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

		status, ok := parseStatus(*statusString)
		if !ok {
			return nil, errors.New("Error parsing status string")
		}
		t.Status = status

		todos = append(todos, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

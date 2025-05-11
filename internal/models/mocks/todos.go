package mocks

import (
	"time"

	"github.com/jansuthacheeva/honkboard/internal/enums"
	"github.com/jansuthacheeva/honkboard/internal/models"
)

var mockTodo = models.Todo{
	ID:      1,
	Title:   "Mock Todo",
	Type:    enums.TodoTypePersonal,
	Status:  enums.TodoStatusDone,
	Created: time.Now(),
}

type TodoModel struct{}

func (m *TodoModel) Insert(userId int, title, typeString string) (int, error) {
	return 2, nil
}

func (m *TodoModel) Delete(userId, id int) error {
	switch {
	case id == 1:
		return nil
	default:
		return models.ErrNoRecord
	}
}

func (m *TodoModel) DeleteCompleted(userId int, listType string) error {
	switch {
	case listType == enums.TodoTypePersonal.String():
		return nil
	case listType == "noCompleted":
		return models.ErrNoRecord
	default:
		return models.ErrUnknownType
	}
}

func (m *TodoModel) ToggleStatus(id int) (models.Todo, error) {
	switch {
	case id == 1:
		return models.Todo{
			ID:      mockTodo.ID,
			Title:   mockTodo.Title,
			Type:    mockTodo.Type,
			Status:  enums.TodoStatusNotDone,
			Created: mockTodo.Created,
		}, nil
	case id == 2:
		return models.Todo{}, models.ErrUnknownStatus
	default:
		return models.Todo{}, models.ErrNoRecord
	}
}

func (m *TodoModel) GetAll(userId int, listType string) ([]models.Todo, error) {
	var todos []models.Todo
	switch {
	case listType == enums.TodoTypePersonal.String():
		todos = append(todos, mockTodo)
		return todos, nil
	case listType == enums.TodoTypeProfessional.String():
		return todos, nil
	default:
		return nil, models.ErrUnknownType
	}
}

package todo

import (
	"time"
)

// Todo represents a single task
type Todo struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TodoManager handles all todo operations
type TodoManager struct {
	storage Storage
	todos   []Todo
}

// NewTodoManager creates a new TodoManager
func NewTodoManager(storage Storage) (*TodoManager, error) {
	tm := &TodoManager{
		storage: storage,
	}
	
	// Load existing todos
	if err := tm.load(); err != nil {
		return nil, err
	}
	
	return tm, nil
}

// Add creates a new todo
func (tm *TodoManager) Add(description string) error {
	todo := Todo{
		ID:          tm.generateID(),
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	
	tm.todos = append(tm.todos, todo)
	return tm.save()
}

// List returns all todos
func (tm *TodoManager) List() ([]Todo, error) {
	return tm.todos, nil
}

// Complete marks a todo as done
func (tm *TodoManager) Complete(id int) error {
	for i := range tm.todos {
		if tm.todos[i].ID == id {
			tm.todos[i].Completed = true
			tm.todos[i].UpdatedAt = time.Now()
			return tm.save()
		}
	}
	return ErrTodoNotFound
}

// Delete removes a todo
func (tm *TodoManager) Delete(id int) error {
	for i, todo := range tm.todos {
		if todo.ID == id {
			tm.todos = append(tm.todos[:i], tm.todos[i+1:]...)
			return tm.save()
		}
	}
	return ErrTodoNotFound
}

// Clear removes all todos
func (tm *TodoManager) Clear() error {
	tm.todos = []Todo{}
	return tm.save()
}

// load todos from storage
func (tm *TodoManager) load() error {
	return tm.storage.Load(&tm.todos)
}

// save todos to storage
func (tm *TodoManager) save() error {
	return tm.storage.Save(tm.todos)
}

// generateID creates a new unique ID
func (tm *TodoManager) generateID() int {
	if len(tm.todos) == 0 {
		return 1
	}
	maxID := 0
	for _, todo := range tm.todos {
		if todo.ID > maxID {
			maxID = todo.ID
		}
	}
	return maxID + 1
}
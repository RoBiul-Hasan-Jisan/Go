package todo

import (
	"encoding/json"
	"errors"
	"os"
)

// Storage interface defines methods for data persistence
type Storage interface {
	Load(interface{}) error
	Save(interface{}) error
}

// FileStorage implements Storage using JSON files
type FileStorage struct {
	filepath string
}

// NewFileStorage creates a new FileStorage
func NewFileStorage(filepath string) *FileStorage {
	return &FileStorage{filepath: filepath}
}

// Load reads data from JSON file
func (fs *FileStorage) Load(data interface{}) error {
	file, err := os.Open(fs.filepath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// File doesn't exist yet, start with empty data
			return nil
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(data)
}

// Save writes data to JSON file
func (fs *FileStorage) Save(data interface{}) error {
	// Create directory if it doesn't exist
	dir := "data"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	file, err := os.Create(fs.filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

// Errors
var (
	ErrTodoNotFound = errors.New("todo not found")
)
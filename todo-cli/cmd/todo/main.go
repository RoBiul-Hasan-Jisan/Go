package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yourusername/todo-cli/internal/todo"
	"github.com/yourusername/todo-cli/internal/utils"
)

func main() {
	if len(os.Args) < 2 {
		utils.PrintUsage()
		os.Exit(1)
	}

	// Initialize storage and manager
	storage := todo.NewFileStorage("data/todos.json")
	manager, err := todo.NewTodoManager(storage)
	if err != nil {
		log.Fatal(err)
	}

	// Handle commands
	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task description")
			os.Exit(1)
		}
		err := manager.Add(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("✅ Task added successfully")

	case "list":
		todos, err := manager.List()
		if err != nil {
			log.Fatal(err)
		}
		utils.PrintTodos(todos)

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task ID")
			os.Exit(1)
		}
		id, err := utils.ParseID(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		err = manager.Complete(id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("✅ Task marked as complete")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task ID")
			os.Exit(1)
		}
		id, err := utils.ParseID(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		err = manager.Delete(id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("🗑️ Task deleted successfully")

	case "clear":
		err := manager.Clear()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("🧹 All todos cleared")

	default:
		utils.PrintUsage()
		os.Exit(1)
	}
}
package utils

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/yourusername/todo-cli/internal/todo"
)

// PrintUsage displays help message
func PrintUsage() {
	cyan := color.New(color.FgCyan).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()
	
	fmt.Printf("%s - A simple CLI Todo Manager\n\n", bold("Todo CLI"))
	fmt.Println("Usage:")
	fmt.Printf("  %s add <task>     %s\n", cyan("todo"), "Add a new task")
	fmt.Printf("  %s list           %s\n", cyan("todo"), "List all tasks")
	fmt.Printf("  %s complete <id>  %s\n", cyan("todo"), "Mark task as complete")
	fmt.Printf("  %s delete <id>    %s\n", cyan("todo"), "Delete a task")
	fmt.Printf("  %s clear          %s\n", cyan("todo"), "Clear all tasks")
	fmt.Println("\nExamples:")
	fmt.Println("  todo add \"Buy groceries\"")
	fmt.Println("  todo list")
	fmt.Println("  todo complete 1")
	fmt.Println("  todo delete 2")
}

// PrintTodos displays todos in a formatted way
func PrintTodos(todos []todo.Todo) {
	if len(todos) == 0 {
		fmt.Println("📭 No tasks found")
		return
	}

	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	
	for _, t := range todos {
		status := red("✗")
		if t.Completed {
			status = green("✓")
		}
		
		// Format time
		timeStr := t.CreatedAt.Format("Jan 02")
		
		fmt.Printf("%s [%s] %-6d %s %s\n", 
			status, 
			yellow(timeStr), 
			t.ID, 
			t.Description,
			getPriorityColor(t))
	}
}

// ParseID converts string to integer ID
func ParseID(idStr string) (int, error) {
	return strconv.Atoi(idStr)
}

// getPriorityColor returns color based on task age
func getPriorityColor(t todo.Todo) string {
	days := time.Since(t.CreatedAt).Hours() / 24
	
	if days > 7 && !t.Completed {
		return color.RedString("(urgent)")
	} else if days > 3 && !t.Completed {
		return color.YellowString("(pending)")
	}
	return ""
}
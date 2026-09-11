package main

/*
File I/O + JSON (encoding/json)
Two very common real-world tasks: reading/writing files and
serializing Go structs to/from JSON.
*/

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"` // omitted from JSON if empty
}

func main() {
	// ---------- Writing a file ----------
	content := []byte("Hello from Go!\nThis is line two.\n")
	err := os.WriteFile("demo.txt", content, 0644)
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}
	fmt.Println("Wrote demo.txt")

	// ---------- Reading a file ----------
	data, err := os.ReadFile("demo.txt")
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}
	fmt.Println("File content:\n" + string(data))

	// ---------- Appending to a file ----------
	f, err := os.OpenFile("demo.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err == nil {
		defer f.Close()
		f.WriteString("Appended line.\n")
	}

	// cleanup for the demo
	defer os.Remove("demo.txt")

	// ---------- JSON: struct -> JSON (Marshal) ----------
	user := User{Name: "Sadia", Age: 24, Email: "sadia@example.com"}
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Marshal error:", err)
		return
	}
	fmt.Println("JSON:", string(jsonBytes))

	// Pretty-printed JSON
	prettyJSON, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println("Pretty JSON:\n" + string(prettyJSON))

	// ---------- JSON -> struct (Unmarshal) ----------
	rawJSON := `{"name":"Tanvir","age":30}`
	var decoded User
	if err := json.Unmarshal([]byte(rawJSON), &decoded); err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}
	fmt.Println("Decoded struct:", decoded)

	// JSON into a generic map when the shape is unknown
	var generic map[string]interface{}
	json.Unmarshal([]byte(rawJSON), &generic)
	fmt.Println("Decoded into map:", generic)
}

package main

/*
Maps: key-value store, unordered.
Declare: map[KeyType]ValueType
*/

import "fmt"

func main() {
	// Create with make
	ages := make(map[string]int)
	ages["Alice"] = 30
	ages["Bob"] = 25

	// Map literal
	capitals := map[string]string{
		"Bangladesh": "Dhaka",
		"Japan":      "Tokyo",
		"France":     "Paris",
	}

	fmt.Println("Ages map:", ages)
	fmt.Println("Capitals map:", capitals)

	// Access
	fmt.Println("Bob's age:", ages["Bob"])

	// Access with "comma ok" idiom to check existence
	value, exists := ages["Charlie"]
	fmt.Println("Charlie exists?", exists, "Value:", value)

	// Update
	ages["Bob"] = 26

	// Delete
	delete(ages, "Alice")
	fmt.Println("After delete:", ages)

	// Iterate (order is NOT guaranteed)
	for country, capital := range capitals {
		fmt.Println(country, "->", capital)
	}

	// len() on map
	fmt.Println("Number of capitals:", len(capitals))

	// Nested map
	users := map[string]map[string]string{
		"u1": {"name": "Nadia", "role": "admin"},
	}
	fmt.Println("Nested map:", users["u1"]["name"])
}

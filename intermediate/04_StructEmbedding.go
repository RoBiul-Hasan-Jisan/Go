package main

/*
Go has no classical inheritance. Instead it uses COMPOSITION via embedding:
embed one struct (or interface) inside another to "promote" its fields/methods.
*/

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Describe() string {
	return a.Name + " makes a sound"
}

func (a Animal) Eat() {
	fmt.Println(a.Name, "is eating")
}

// Dog EMBEDS Animal (no field name -> promoted fields/methods)
type Dog struct {
	Animal // embedded struct
	Breed  string
}

// Dog can override a promoted method by defining its own
func (d Dog) Describe() string {
	return d.Name + " (a " + d.Breed + ") barks"
}

// Embedding interfaces works too - useful for wrapping/extending behavior
type Reader interface{ Read() string }
type Writer interface{ Write(string) }

type ReadWriter interface {
	Reader // embedded interface
	Writer
}

type File struct{ content string }

func (f *File) Read() string      { return f.content }
func (f *File) Write(s string)    { f.content += s }

func main() {
	d := Dog{
		Animal: Animal{Name: "Rex"},
		Breed:  "Labrador",
	}

	// Promoted field access - looks like inheritance, but it's composition
	fmt.Println("Name (promoted field):", d.Name)

	// Overridden method
	fmt.Println(d.Describe())

	// Promoted method (not overridden)
	d.Eat()

	// Access the embedded struct explicitly if needed
	fmt.Println("Original Animal.Describe():", d.Animal.Describe())

	var rw ReadWriter = &File{}
	rw.Write("hello ")
	rw.Write("world")
	fmt.Println("File content:", rw.Read())
}

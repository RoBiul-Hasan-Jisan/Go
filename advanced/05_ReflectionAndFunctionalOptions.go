package main

/*
Two "pro-level" tricks that show up constantly in real Go codebases:

1. reflect: inspect types/values at runtime. Powers libraries like
   encoding/json, ORMs, validators. Use sparingly - it's slower and less
   type-safe than normal Go code.

2. Functional Options Pattern: the idiomatic Go way to build configurable
   constructors without a giant parameter list or requiring every caller
   to know every option.
*/

import (
	"fmt"
	"reflect"
)

type Product struct {
	Name  string
	Price float64
}

func inspect(v interface{}) {
	t := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	fmt.Println("Type:", t.Name(), "| Kind:", t.Kind())

	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			value := val.Field(i)
			fmt.Printf("  Field: %s (%s) = %v\n", field.Name, field.Type, value)
		}
	}
}

// ---------- Functional Options Pattern ----------
type Server struct {
	host    string
	port    int
	timeout int
	debug   bool
}

type ServerOption func(*Server)

func WithHost(host string) ServerOption {
	return func(s *Server) { s.host = host }
}

func WithPort(port int) ServerOption {
	return func(s *Server) { s.port = port }
}

func WithTimeout(seconds int) ServerOption {
	return func(s *Server) { s.timeout = seconds }
}

func WithDebug() ServerOption {
	return func(s *Server) { s.debug = true }
}

// NewServer has sensible defaults; callers only specify what they want to change
func NewServer(opts ...ServerOption) *Server {
	s := &Server{ // defaults
		host:    "localhost",
		port:    8080,
		timeout: 30,
		debug:   false,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func main() {
	// ---------- reflect ----------
	p := Product{Name: "Keyboard", Price: 49.99}
	inspect(p)
	inspect(42)
	inspect("hello")

	// ---------- functional options ----------
	s1 := NewServer() // all defaults
	fmt.Printf("s1: %+v\n", s1)

	s2 := NewServer(
		WithHost("0.0.0.0"),
		WithPort(9000),
		WithDebug(),
	)
	fmt.Printf("s2: %+v\n", s2)
}

package main

/*
Strings in Go are IMMUTABLE, read-only slices of bytes (UTF-8 encoded).
A "rune" is Go's term for a Unicode code point (int32).
len("héllo") counts BYTES, not characters - watch out with multi-byte chars.
*/

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	s := "Hello, Go!"

	// ---------- strings package ----------
	fmt.Println("Upper:", strings.ToUpper(s))
	fmt.Println("Lower:", strings.ToLower(s))
	fmt.Println("Contains 'Go'?", strings.Contains(s, "Go"))
	fmt.Println("Replace:", strings.Replace(s, "Go", "World", 1))
	fmt.Println("Split by ', ':", strings.Split(s, ", "))
	fmt.Println("Join:", strings.Join([]string{"a", "b", "c"}, "-"))
	fmt.Println("TrimSpace:", strings.TrimSpace("   padded   "))
	fmt.Println("HasPrefix:", strings.HasPrefix(s, "Hello"))
	fmt.Println("Index of 'Go':", strings.Index(s, "Go"))

	// Efficient string building (avoid += in a loop, it reallocates every time)
	var sb strings.Builder
	for i := 0; i < 5; i++ {
		sb.WriteString("Go ")
	}
	fmt.Println("Builder result:", sb.String())

	// ---------- runes ----------
	word := "héllo"
	fmt.Println("len() in bytes:", len(word))
	fmt.Println("rune count:", len([]rune(word)))

	for i, r := range word { // range over a string yields (byteIndex, rune)
		fmt.Printf("byteIndex %d -> rune %q (code point %d)\n", i, r, r)
	}

	// unicode helpers
	fmt.Println("Is 'A' upper?", unicode.IsUpper('A'))
	fmt.Println("Is '5' a digit?", unicode.IsDigit('5'))

	// ---------- conversions ----------
	n, err := strconv.Atoi("42") // string -> int
	fmt.Println("Atoi:", n, err)

	str := strconv.Itoa(99) // int -> string
	fmt.Println("Itoa:", str)

	f, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println("ParseFloat:", f)
}

package calculator

import "testing"

// A basic test function: must start with "Test", take *testing.T
func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5
	if got != want {
		t.Errorf("Add(2, 3) = %d; want %d", got, want)
	}
}

// Table-driven tests: the idiomatic Go way to test many cases cleanly
func TestDivide(t *testing.T) {
	cases := []struct {
		name    string
		a, b    int
		want    int
		wantErr bool
	}{
		{"normal division", 10, 2, 5, false},
		{"divide by zero", 10, 0, 0, true},
		{"negative numbers", -10, 2, -5, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Divide(tc.a, tc.b)
			if (err != nil) != tc.wantErr {
				t.Fatalf("unexpected error state: err=%v, wantErr=%v", err, tc.wantErr)
			}
			if !tc.wantErr && got != tc.want {
				t.Errorf("Divide(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("madam") {
		t.Error("expected 'madam' to be a palindrome")
	}
	if IsPalindrome("hello") {
		t.Error("expected 'hello' to NOT be a palindrome")
	}
}

// Benchmark: run with `go test -bench=.`
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

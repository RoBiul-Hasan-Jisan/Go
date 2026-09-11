package main

/*
Arrays vs Slices
----------------
Array:  fixed size, size is part of the type      -> var a [5]int
Slice:  dynamic size, backed by an array under the hood -> []int
append() may allocate a new underlying array when capacity is exceeded.
*/

import "fmt"

func main() {
	// ---------- ARRAYS ----------
	var arr [5]int // fixed size, zero-valued: [0 0 0 0 0]
	arr[0] = 10
	arr[1] = 20
	fmt.Println("Array:", arr, "Length:", len(arr))

	arr2 := [3]string{"Go", "Rust", "Python"} // array literal
	fmt.Println("Array2:", arr2)

	// ---------- SLICES ----------
	nums := []int{1, 2, 3, 4, 5} // slice literal
	fmt.Println("Slice:", nums, "Len:", len(nums), "Cap:", cap(nums))

	// Slicing: s[low:high] -> excludes high
	sub := nums[1:4]
	fmt.Println("Sub-slice nums[1:4]:", sub)

	// append (grows the slice, may reallocate)
	nums = append(nums, 6, 7)
	fmt.Println("After append:", nums)

	// make() creates a slice with length & capacity
	scores := make([]int, 3, 5) // len=3, cap=5
	scores[0] = 90
	fmt.Println("make() slice:", scores, "len:", len(scores), "cap:", cap(scores))

	// copy()
	dst := make([]int, len(nums))
	copy(dst, nums)
	fmt.Println("Copied slice:", dst)

	// 2D slice
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2D slice:", grid)

	// Removing an element at index i (no built-in remove)
	i := 2
	nums = append(nums[:i], nums[i+1:]...)
	fmt.Println("After removing index 2:", nums)

	// range over slice
	for idx, val := range nums {
		fmt.Println("idx:", idx, "val:", val)
	}
}

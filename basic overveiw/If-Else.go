package main

import "fmt"

func main() {

   var num int 


   fmt.Print("Enter Your Number : ")
   fmt.Scan(&num)

   if num > 0 {
	if num%2 == 0 {
		fmt.Println("positive Even  Number")
	} else {
		fmt.Println("Positive Odd Number")
	}
   } else if num < 0 {
	fmt.Println("Negative Number ")
   } else {
	fmt.Println("Zero")
   }
}

/*

if condition1 {
    if condition2 {
        // code runs when condition1 AND condition2 are true
    } else {
        // code runs when condition1 is true but condition2 is false
    }
} else if condition3 {
    // code runs when condition1 is false AND condition3 is true
} else {
    // code runs when all above conditions are false
}


*/
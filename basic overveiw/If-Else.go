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

package main

import "fmt"

func addNum(a , b int) (int) {
	return a+b;
}

// func addNum(a , b int) (result int) {
// 	result = a+b;
// 	return
// }




func main(){

	ans := addNum(3,3)

	fmt.Println("The addition of 2 numbers is:",ans)
}
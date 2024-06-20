package main

import "fmt"


func main(){

	var name string = "Shubhankit"
	var age int  = 20
	var height float32 = 6.129383

	// always prints in a new line and in this format specifiers doens't works
	fmt.Println("My name is %s and i'm %d year old and my height is %.2f", name, age, height) ❌
	// Gives us an error: fmt.Println call has possible Printf formatting directive %sprintfdefault
	fmt.Println("My name is ", name," and i'm ", age," year old and my height is ",height) ✅
	// output: My name is  Shubhankit  and i'm  20  year old and my height is  6.129383
	
	// this always prints the string in the same line and format specifiers will also works in this.
	fmt.Printf("My name is %s and i'm %d year old and my height is %.2f", name, age, height) ✅
	// My name is Shubhankit and i'm 20 year old and my height is 6.13%                        
	fmt.Println("My name is ", name," and i'm ", age," year old and my height is ",height) ✅
	// output: My name is  Shubhankit  and i'm  20  year old and my height is  6.129383
}
package main

import (
	"fmt"
)

func main()  {

	var name string  = "Shubhankit"; // String
	var version int  = 30; // Number
	var isOkay bool = false // boolean
	var amount float32 = 193.98 // Decimanl

	const pi = 3.14 // Constant

	var PublicVar = 99 // Can be exported  ( can be imported in other files too. )
	var privateVar = 9 // cannot be exported ( cannot be imported in other files. )


	fmt.Println(name);
	fmt.Println(version);
	fmt.Println(isOkay);
	fmt.Println(amount);
	
	fmt.Println(pi)

	fmt.Println(PublicVar)
	fmt.Println(privateVar)


}
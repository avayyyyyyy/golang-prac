package main

import (
	"bufio"
	"fmt"
	"os"
)


func main(){

	var name string;

	fmt.Println("Tell me your name!")
	
	fmt.Scan(&name)
	
	fmt.Printf("Hii, %s",name)
	
	var reader = bufio.NewReader(os.Stdin)

	var name, _  = reader.ReadString('\n')

	fmt.Printf("Hii %s",name)

}
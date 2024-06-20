package main

import (
	"bufio"
	"fmt"
	"os"
)


func main(){

	fmt.Println("Tell me your name!")
	var reader = bufio.NewReader(os.Stdin)

	var name, _  = reader.ReadString('\n')

	fmt.Printf("Hii %s",name)

}
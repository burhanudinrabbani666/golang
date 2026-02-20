package main

import "fmt"

func main(){
	const(
	firstName string = "Burhanudin" // Not Error if not used
	lastName = "Rabbani"
	)
	
	fmt.Println(firstName, lastName) // Burhanudin Rabbani
}

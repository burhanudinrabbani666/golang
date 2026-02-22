package main

import "fmt"

func sayHelloTo() (string, int) {
	return "Burhanudin", 23
}

func main() {
	firstName, _ := sayHelloTo()

	fmt.Println(firstName)
}

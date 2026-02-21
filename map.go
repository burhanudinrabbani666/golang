package main

import "fmt"

func main(){

	book := make(map[string]string)
	book["Title"] = "Hello World"
	book["Authoer"] = "Bani"
	book["ups"] = "salah"

 	fmt.Println(book) // map[Authoer:Bani Title:Hello World ups:salah]

	delete(book, "ups")
 	fmt.Println(book) // map[Authoer:Bani Title:Hello World]
}
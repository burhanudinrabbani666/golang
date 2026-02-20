package main

import "fmt"

func main(){
	var  name string = "Burhanudin Rabbani"
	var isByte = name[0]
	var byteToString = string(isByte)

	fmt.Println(name) // Burhanudin Rabbani
	fmt.Println(isByte) // 66
	fmt.Println(byteToString) // B

}
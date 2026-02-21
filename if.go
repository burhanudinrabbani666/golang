package main

import "fmt"

func main() {

	name := "Jokow"

	if length := len(name); length >= 5 { // Short Statement
		fmt.Println("Name Terlalu Panjang")
	} else {
		fmt.Println("Name Sudah Benar")
	}

}

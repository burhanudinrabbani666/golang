package main

import "fmt"

func main(){

	var number = [...]int{ // ... ketika data array nya ngga pasti berapa tapi harus langsung dideklarsi kan 
		14, 11, 12,
	}
	
	fmt.Println(len(number))
}
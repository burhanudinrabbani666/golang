package main

import "fmt"

func main(){
	type NoKTP string

	var KTPBani NoKTP = "111111"

	var contoh string = "2222222"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(KTPBani)
	fmt.Println(contohKTP)

}
package main

import "fmt"

func main() {

	names := []string{"Bani", "Agus", "Heri"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	fmt.Println("Selesai")

}

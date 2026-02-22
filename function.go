package main

import (
	"fmt"
	"strings"
)

func sayHelloWIthFilter(name string, filter func(string) string) {
	filteredName := filter(name)

	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if strings.ToLower(name) == "anjing" {
		return "***"
	} else {
		return name
	}

}

func main() {
	sayHelloWIthFilter("Bani", spamFilter)

	sayHelloWIthFilter("AnjinG", spamFilter)
}

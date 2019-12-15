package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		age  = 29
		name = "Paulin"
	)

	fmt.Println("Je suis " + name + ", j'ai " + strconv.Itoa(age) + " ans et je suis " + majorityStatus(age))
}

func majorityStatus(age int) string {
	if age >= 18 {
		return "majeur"
	}
	return "minor"
}

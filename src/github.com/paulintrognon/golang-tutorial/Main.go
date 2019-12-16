package main

import (
	"fmt"
	"strconv"
)

const (
	isMale = iota + 1
	isFemale
)

func main() {
	var (
		age  = 29
		name = "Paulin"
		sex  = isMale
	)

	fmt.Println("Je suis " + name + ", j'ai " + strconv.Itoa(age) + " (je suis " + majorityStatus(age) + ")")

	if sex == isMale {
		fmt.Println("Je suis un homme")
	} else {
		fmt.Println("Je suis une femme")
	}
}

func majorityStatus(age int) string {
	if age >= 18 {
		return "majeur"
	}
	return "mineur"
}

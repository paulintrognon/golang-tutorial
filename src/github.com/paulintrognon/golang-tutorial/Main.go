package main

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	male = iota + 1
	female
)

type person struct {
	name  string
	age   int
	sex   int
	books []string
}

func main() {

	me := person{
		name:  "Paulin",
		age:   29,
		sex:   male,
		books: []string{"Dune", "Foundation"},
	}

	fmt.Println("Je suis " + me.name + ", j'ai " + strconv.Itoa(me.age))

	majorityStatus, err := getMajorityStatus(me.age)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("je suis " + majorityStatus)
	}

	if me.sex == male {
		fmt.Println("Je suis un homme")
	} else {
		fmt.Println("Je suis une femme")
	}

	listBooks(me.books)

	inc(&me.age)

	fmt.Println("J'ai maintenant " + strconv.Itoa(me.age) + " ans !")
}

func getMajorityStatus(age int) (string, error) {
	if age < 0 {
		return "", errors.New("L'age ne peut pas être négatif")
	}
	if age >= 18 {
		return "majeur", nil
	}
	return "mineur", nil
}

func listBooks(books []string) {
	fmt.Println("Mes livres :")
	for index := range books {
		fmt.Println(books[index])
	}
}

func inc(x *int) {
	*x++
}

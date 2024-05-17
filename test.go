package main

import (
	"fmt"
	"unicode/utf8"
)

func calc(x, y int) int {
	return x - y
}

func name(firstName, lastName string, age int) string {
	name := firstName + " " + lastName

	if length := utf8.RuneCountInString(name); length > 1 {
		fmt.Println("\nMy name is:", name)
		fmt.Println("My age is:", age)
		fmt.Printf("str length: %v", length)
	}
	return name
}

func main() {
	const firstName = "Alireza"
	const lastName = "Esmaeili"

	age := calc(1403, 1379)

	if age < 18 {
		fmt.Printf("you are still young, enjoy your carefree days while they last")
	} else if age < 20 {
		fmt.Printf("you need to find and passion and work on it for your carrier")
	} else {
		fmt.Printf("hurry up before its too late")
	}

	name(firstName, lastName, age)

	fmt.Printf("\nhi: %v\n", "go")
	fmt.Printf("hi: %s\n", "go")
	fmt.Printf("hi: %v\n", 42)
	fmt.Printf("hi: %q\n", "go")
	fmt.Printf("hi: %#q\n", "go")
}

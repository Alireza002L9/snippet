package main

import (
	"fmt"
	"unicode/utf8"
)

type skills struct {
	stacks []string
	moreInfo string
}
type jobSpecification struct {
	skills
	name string
	employed bool
	job string
	location string
	salary string
	hoursWorked int
}

func (x jobSpecification) generateResume() string {
	introduceMsg := fmt.Sprintf("Hello I'm  %s", x.name)
	return introduceMsg
}




func jobDes(x jobSpecification) {
	if x.employed {
		fmt.Printf("you work as %s in %s for %s amount of money for the total of %v hours of work", x.job, x.location, x.salary, x.hoursWorked)
	} else {
		fmt.Printf("you are currently unemployed")
	}
}

func calc(x, y int) int {
	return x - y
}

func name(firstName, lastName string, age int) string {
	fullname := firstName + " " + lastName
	name := jobSpecification{
		name: fullname,
	} 


	if length := utf8.RuneCountInString(fullname); length > 1 {
		fmt.Println("\nMy name is:", fullname)
		fmt.Println(jobSpecification.generateResume(name))
		fmt.Println("My age is:", age)
		fmt.Printf("str length: %v", length)
	}
	return fullname
}

func main() {
	const firstName = "Alireza"
	const lastName = "Esmaeili"

	workInfo := jobSpecification{
		employed: true,
		job: "developer",
		location: "Technology and science estate",
		salary: "unkown",
		hoursWorked: 204,
		skills: skills{
			stacks: []string{"react", "three.js"},
		},
	}
	resume := workInfo.generateResume()

	
	age := calc(1403, 1379)

	if age < 18 {
		fmt.Printf("\nyou are still young, enjoy your carefree days while they last\n")
	} else if age < 20 {
		fmt.Printf("\nyou need to find and passion and work on it for your carrier\n")
	} else {
		fmt.Printf("\nhurry up before its too late\n")
	}


	jobDes(workInfo)
	name(firstName, lastName, age)

	fmt.Printf("\nhi: %v\n", "go")
	fmt.Printf("hi: %s\n", "go")
	fmt.Printf("hi: %v\n", 42)
	fmt.Printf("hi: %q\n", "go")
	fmt.Printf("hi: %#q\n", "go")
	fmt.Println(resume)
}


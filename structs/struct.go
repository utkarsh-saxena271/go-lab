package main

import "fmt"

type student struct {
	name       string
	rollNumber int
	section    string
}

func main() {
	var stu1 student
	stu1.name = "utkarsh"
	stu1.rollNumber = 10
	stu1.section = "A"

	fmt.Println("Name: ", stu1.name);
	fmt.Println("Roll Number: ", stu1.rollNumber);
	fmt.Println("Section: ", stu1.section);

}

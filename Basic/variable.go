package main

import "fmt"

func main() {

	//explicit declaration

	var name string = "Aakash Kumar"
	var age int = 28

	//implicit declaration // type is inferred by the compiler

	country := "India"
	city := "Nalanda"

	fmt.Println("Hello my name is ", name)
	fmt.Println("Age is ", age)
	fmt.Println("Country is ", country)
	fmt.Println("City is ", city)
}

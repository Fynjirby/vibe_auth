package main

import "fmt"

func main() {
	var answer string

	fmt.Println("Enter password: ")
	fmt.Scanln(&answer)
	if answer == "password" { } else { return }
	fmt.Println("Password is incorrect ")
	fmt.Scanln(&answer)
	if answer == "incorrect" { } else { return }
	fmt.Println("Try again. ")
	fmt.Scanln(&answer)
	if answer == "again" { } else { return }
	fmt.Println("Please try again later")
	fmt.Scanln(&answer)
	return
}

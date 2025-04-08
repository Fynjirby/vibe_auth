package main

import "fmt"

func main() {
	var first, second, third, fourth string

	fmt.Println("Enter password: ")
	fmt.Scanln(&first)
	if first == "password" { } else { return }
	fmt.Println("Password is incorrect ")
	fmt.Scanln(&second)
	if second == "incorrect" { } else { return }
	fmt.Println("Try again. ")
	fmt.Scanln(&third)
	if third == "again" { } else { return }
	fmt.Println("Please try again later")
	fmt.Scanln(&fourth)
	if fourth == "again later" { } else { return }
}

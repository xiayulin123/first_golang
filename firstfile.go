package main

import "fmt"

func main(){
	fmt.Println("Welcome to the quiz game")
	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello, %v, welcome to the game \n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay you can play the game!")

	} else {
		fmt.Println("Sorry you have to be greater than 10 to play this game")
		return
	}
	fmt.Println("continue...")
	fmt.Printf("Is RTX 3080 better or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)
	if answer + " " + answer2 == "RTX 3090"{
		fmt.Println("Correct!")
	}else{
		fmt.Println("False!")
	}

}
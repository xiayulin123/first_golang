package main

import "fmt"

func main(){
	fmt.Println("Welcome to the quiz game")
	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello, %v, welcome to the game \n", name)
}
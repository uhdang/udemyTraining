package main

import "fmt"

func main() {
	var myGreeting map[string]string
	fmt.Println(myGreeting)
	fmt.Println(myGreeting == nil)

	// Doing below will give error
	myGreeting["Tim"] = "Good morning"
	myGreeting["Jenny"] = "Bonjour"
}

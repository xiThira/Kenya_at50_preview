package main

import "fmt"

func main() {
	fmt.Println("Kenya at 50 Quiz game")

	fmt.Printf("What is your name: ")

	var name string

	fmt.Scan(&name)

	fmt.Printf("Hello %v, Welcome to the quiz!\n", name)

	fmt.Printf("How old are you: ")
	var age uint16

	fmt.Scan(&age)
	if age >= 10 {
		fmt.Printf("Yay you can play!\n")
	} else {
		fmt.Printf("Oops,sorry you can't play")
		return
	}

	score := 0
	num_questions := 3

	fmt.Printf("Who is the first Vice President of Kenya? ")
	var answer string
	var answer2 string
	var answer3 string
	fmt.Scan(&answer, &answer2, &answer3)

	if answer+" "+answer2+" "+answer3 == "Jaramogi Oginga Odinga" {
		fmt.Println("Correct")
		score += 1
	} else if answer+" "+answer2+" "+answer3 == "jaramogi oginga odinga" {
		fmt.Println("Correct")
		score += 1
	} else {
		fmt.Println("Incorrect")
	}

	fmt.Printf("T/f: Turkana is the second largest county in Kenya:  ")
	var county_response bool
	fmt.Scan(&county_response)
	if county_response == true {
		fmt.Println("Incorrect")
	} else {
		fmt.Println("Correct")
		score++
	}
	fmt.Printf("How many years has it been since the introduction of the new Constitution? ")
	var years uint
	fmt.Scan(&years)
	if years == 13 {
		fmt.Println("Correct")
		score = score + 1
	} else {
		fmt.Println("Incorrect")
	}
	fmt.Printf("You scored %v out of %v.\n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("You scored %v%%.", percent)

}

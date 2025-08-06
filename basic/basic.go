package basic

import (
	"fmt"
	"tutordek/calculation" // Replace with your actual module path
)

func aaa() {
	result := calculation.Add(9, 9)
	fmt.Printf("The result of calculation is %d", result)

	name := "kussoy"
	fmt.Printf("\n%s", name)

	var surname = "Ica"
	fmt.Printf("\n%s", surname)

	title := "Pejabat"
	for _, letter := range title {
		fmt.Printf("\nletter %d", letter)
	}

	fruits := [...]string{"banana", "apple", "orange", "grape"}

	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// append and arrays

	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	// get average value
	scoresLength := len(scores)
	var scoresTotal float64

	for _, score := range scores {
		scoresTotal = float64(score) + scoresTotal
	}

	fmt.Println(scoresTotal / float64(scoresLength))

	// score greater equal than 90
	var goodScores []int

	for _, score := range scores {
		if score >= 90 {
			goodScores = append(goodScores, score)
		}
	}

	fmt.Println(goodScores)
}

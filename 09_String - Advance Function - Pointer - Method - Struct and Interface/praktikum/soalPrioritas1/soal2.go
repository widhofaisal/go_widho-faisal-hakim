package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) average() {
	var minimum, maximum, average int = 100, 0, 0
	var minName, maxName string
	for i, eachScore := range s.score {
		// codition minimun
		if eachScore < minimum {
			minimum = eachScore
			minName = s.name[i]
		}
		// condition maximum
		if eachScore > maximum {
			maximum = eachScore
			maxName = s.name[i]
		}
		// average
		average += eachScore
	}
	average = average / len(s.score)
	fmt.Printf("Average Score\t: %d\n", average)
	fmt.Printf("Min Score of Students\t: %s (%d)\n", minName, minimum)
	fmt.Printf("Max Score of Students\t: %s (%d)\n", maxName, maximum)
}

func main() {
	var names = []string{}
	var scores = []int{}

	for i := range [5]int{} {
		var name string = ""
		var score int = 0

		fmt.Printf("Input-%d_Student's_Name\t: ", i+1)
		fmt.Scanln(&name)
		fmt.Printf("Input-%d_Student's_Score\t: ", i+1)
		fmt.Scanln(&score)

		names = append(names, name)
		scores = append(scores, score)
	}
	s1 := Student{names, scores}
	s1.average()
}

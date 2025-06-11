package utils

import "fmt"

func GetInputInt(prompt string) int {
	fmt.Print(prompt)
	var input int
	fmt.Scanln(&input)
	return input
}

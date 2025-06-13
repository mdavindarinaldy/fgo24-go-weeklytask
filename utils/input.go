package utils

import (
	"fmt"
	"strings"
	"time"
)

func GetInputInt(prompt string) int {
	fmt.Print(prompt)
	var input int
	fmt.Scanln(&input)
	return input
}

func GetInputString(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	return input
}

func InvalidInput() {
	fmt.Print("\nInput tidak valid! Silakan coba lagi\n\n")
	time.Sleep(time.Second)
}

package utils

import (
	"fmt"
	"time"
)

func GetInputInt(prompt string) int {
	fmt.Print(prompt)
	var input int
	fmt.Scanln(&input)
	return input
}

func InvalidInput() {
	fmt.Print("\nInput tidak valid! Silakan coba lagi\n\n")
	time.Sleep(time.Second)
}

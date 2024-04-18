package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Введите фразу без пробелов:")
	fmt.Scanln(&input)

	input = strings.ToLower(input)

	frequencyMap := make(map[rune]int)
	for _, char := range input {
		frequencyMap[char]++
	}

	totalChars := float64(len(input))

	for char, frequency := range frequencyMap {
		percentage := (float64(frequency) / totalChars) * 100
		percentageStr := fmt.Sprintf("%.2f", percentage)
		fmt.Printf("%c - %d %s%%\n", char, frequency, percentageStr)
	}
}

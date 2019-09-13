package main

import (
	"fmt"
)

// Complete the isValid function below.
func isValid(s string) string {

	m := countApearTime(countChar(s))

	switch len(m) {
	case 1:
		return "YES"
	case 2:
		for k, v := range m {
			if v == 1 {
				if k == 1 {
					return "YES"
				}
				for k1, _ := range m {
					if k != k1 && k-k1 == 1 {
						return "YES"
					}
				}
			}
		}
		return "NO"
	default:
		return "NO"
	}
}

func countChar(s string) map[string]int {
	m := make(map[string]int)
	for _, value := range s {
		m[string(value)]++
	}
	return m
}

func countApearTime(m map[string]int) map[int]int {
	m2 := make(map[int]int)
	for _, value := range m {
		m2[value]++
	}

	return m2
}

func main() {
	fmt.Println("Hello world!")

	input := "abcdefghhgfedecba"
	fmt.Println(isValid(input))
}

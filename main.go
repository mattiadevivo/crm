package main

import (
	"fmt"
)

func getRectangleArea(width int, length int) string {
	if product := width * length; length < 50 {
		return fmt.Sprintf("The area is %d, which is less than 50", product)
	} else {
		return fmt.Sprintf("The area is %d, which is greater than or equal to 50", product)
	}
}

func arrayAndSlices() {
	var languages = []string{"Go", "JavaScript", "Ruby", "Python"}
	fmt.Println(languages)
	fmt.Println(len(languages))
	fmt.Println(languages[0])
	fmt.Println(languages[1:3])
	languages = append(languages, "PHP")
	fmt.Println(languages)
}

func fizzbuzz(n int) []string {
	buzzs := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				buzzs = append(buzzs, "Fizzbuzz")
				continue
			}
			buzzs = append(buzzs, "Fizz")
		} else if i%5 == 0 {
			buzzs = append(buzzs, "Buzz")
		} else {
			buzzs = append(buzzs, fmt.Sprint(i))
		}
	}
	return buzzs
}

func main() {
	fmt.Println("Hello world")
	fmt.Println(fizzbuzz(15))
}

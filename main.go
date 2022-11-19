package main

import (
	"fmt"
	"strings"
	"time"
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

func mapExercise() {
	courses := map[int]string{1: "Calculus", 2: "Biology", 3: "Chemistry", 4: "Computer Science", 5: "Communications", 6: "English", 7: "Cantonese"}
	for id, course := range courses {
		if strings.HasPrefix(course, "C") {
			fmt.Println(id, course)
		}
	}

	courses[4] = "Algorithms"
	courses[8] = "Spanish"
	delete(courses, 1)
}

// L3.Structs
type Student struct {
	id   int
	name string
}

type Classroom struct {
	id          int
	capacity    int
	subject     string
	studentList []Student
}

func structExercise() {
	c1 := Classroom{id: 1, capacity: 28, subject: "Algorithm", studentList: []Student{{id: 1, name: "First Student"}, {id: 2, name: "Second Student"}}}
	c2 := new(Classroom)
	c2.id = 2
	c2.capacity = 100
	c2.subject = "Theater"
	c2.studentList = []Student{
		{
			id:   40,
			name: "Vince",
		},
		{
			id:   50,
			name: "Johnny",
		},
	}

	fmt.Println(c1)
	fmt.Println(c2)
}

// L3.Interfaces
type Rectangle struct {
	length, width float64
}

type Square struct {
	side float64
}

type Shape interface {
	perimeter() float64
}

func (r Rectangle) perimeter() float64 {
	return (2 * r.length) + (2 * r.width)
}

func (s Square) perimeter() float64 {
	return s.side * 4
}

func getPerimeter(s Shape) float64 {
	return s.perimeter()
}

func interfaceMain() {
	rectangle := Rectangle{length: 2, width: 4}
	square := Square{side: 2}

	fmt.Println("Rectangle perimeter: ", rectangle.perimeter())
	fmt.Println("Square perimeter: ", square.perimeter())
}

// L3.Goroutines
func routinePrinter(a []string) {
	for _, elem := range a {
		fmt.Println(elem)
	}
}

func goroutineMain() {
	startTime := time.Now()
	colorNames := []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"}
	colorCodes := []string{"#FF0000", "#FF7F00", "#FFFF00", "#00FF00", "#0000FF", "#4B0082", "#9400D3"}
	go routinePrinter(colorNames)
	go routinePrinter(colorCodes)

	duration := time.Since(startTime)

	fmt.Println("Duration of execution: " + duration.String())

	time.Sleep(time.Second)
}

func main() {
	fmt.Println("Hello world")
	fmt.Println(fizzbuzz(15))
}

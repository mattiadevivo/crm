package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
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

// L4.Handlers I
var cities = []string{"Tokyo", "Delhi", "Shanghai", "Sao Paulo", "Mexico City"}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world :]")
}

func cityList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, cities)
}

func handlers1Main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/citylist", cityList)

	fmt.Println("Server is starting...")
	// Instructing this HTTP server to listen for incoming requests on port 3000
	http.ListenAndServe(":3000", nil)
}

// L4.Handlers II
var cityPopulations = map[string]uint32{
	"Tokyo":       37435191,
	"Delhi":       29399141,
	"Shanghai":    26317104,
	"Sao Paulo":   21846507,
	"Mexico City": 21671908,
}

func index2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	for k, v := range cityPopulations {
		fmt.Fprintf(w, "<h2>%s, %d</h2>", k, v)
	}
}

func handlers2Main() {
	http.HandleFunc("/index", index2)

	fmt.Println("Server is starting on port 3000")
	http.ListenAndServe(":3000", nil)
}

// L4.Handlers III
func index3(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cityPopulations)
}

func handlers3Main() {
	http.HandleFunc("/index", index3)

	fmt.Println("Server is starting on port 3000")
	http.ListenAndServe(":3000", nil)
}

// L4.Handlers III with Gorilla
var dictionary = map[string]string{
	"Go":     "A programming language created by Google.",
	"Gopher": "A software engineer who builds with Go.",
	"Golang": "Another name for Go.",
}

func getDictionary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(dictionary)
}

func createDictionary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newEntry map[string]string

	// Read request
	reqBody, _ := ioutil.ReadAll(r.Body)
	// Parse JSON body
	json.Unmarshal(reqBody, &newEntry)
	// Add new entry to dictionary
	for k, v := range newEntry {
		if _, isPresent := dictionary[k]; isPresent {
			w.WriteHeader(http.StatusConflict)
		} else {
			dictionary[k] = v
			w.WriteHeader(http.StatusCreated)
		}
	}

	json.NewEncoder(w).Encode(dictionary)
}

func handlers3MuxMain() {
	router := mux.NewRouter()

	router.HandleFunc("/dictionary", getDictionary).Methods("GET")
	router.HandleFunc("/dictionary", createDictionary).Methods("POST")
	fmt.Println("Server is starting on port 3000")
	http.ListenAndServe(":3000", router)
}

// L4.Routing
var members = map[string]string{
	"1": "Andy",
	"2": "Peter",
	"3": "Gabriella",
	"4": "Jordy",
}

func getMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(members)
}

func deleteMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if _, ok := members[id]; ok {
		delete(members, id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(members)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(members)
	}
}

func routingMain() {
	router := mux.NewRouter()

	router.HandleFunc("/members", getMembers).Methods("GET")
	router.HandleFunc("/members/{id}", deleteMember).Methods("DELETE")
	fmt.Println("Server is starting on port 3000")
	http.ListenAndServe(":3000", router)
}

func main() {
	fmt.Println("Hello world")
	fmt.Println(fizzbuzz(15))
	routingMain()
}

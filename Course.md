# Go fundamentals

## Go installation

Install binary file from official site and put them under `/usr/local/bin`.

Add the folder to the path inside `/etc/profile`: `export PATH=/usr/local/bin:$PATH`.

Check that Golang has been correctly installed by running `go version`.

Add the go workspace `bin` folder to the PATH with `export PATH=$PATH:$(go env GOPATH)/bin`.

You can optionally also export `GOPATH` folder: `export GOPATH=$(go env GOPATH)`.


## Go Workspace

Go workspace is the folder/local workspace where all the golang sources, binaries and dependencies are stored. 
Run `go env GOPATH` to see the Go workspace directory.

Inside `GOPATH` directory, Go workspaces have a particular directory:
```
- go
  - src
     - (VERSION_CONTROL)
       - (USERNAME)
          - (PROJECT)
  - bin
  - pkg
```

The `go` folder is the first folder inside your local workspace. Within it are three folders:
- The `src` folder will contain your source code
- The `bin` folder will contain executable files
- The `pkg` folder will contain Go package objects compiled from source code (including third party libraries).

As an example: if I used GitHub for version control, and my username is `udacity`, and I've build a `sample_project`, that project would be located here: `GOPATH/go/src/github.com/udacity/sample_project`

Then you can build and install the project with: `cd $GOPATH/src/github.com/user/hello && go install`, the commands builds the executable binary and then install it inside workspace's `bin` directory.

## Package names

The first statement in a Go source file must be
```go
package name
```
where *name* is the package's default name for imports. (All files in a package must use the same name.)

Go's convention is that the package name is the last element of the import path: the package imported as "crypto/rot13" should be named rot13.

Executable commands must always use package main.

There is no requirement that package names be unique across all packages linked into a single binary, only that the import paths (their full file names) 

## Testing

Go has a lightweight test framework composed of the `go test` command and the **testing** package.

You write a test by creating a file with a name ending in `_test.go` that contains functions named `TestXXX` with signature `func (t *testing.T)`. The test framework runs each such function; if the function calls a failure function such as `t.Error` or `t.Fail`, the test is considered to have failed.

Example of test:
```go
package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
```

Then run the test with go test: `go test `.

## Remote packages

An import path can describe how to obtain the package source code using a revision control system such as Git or Mercurial. The go tool uses this property to automatically fetch packages from remote repositories. For instance, the examples described in this document are also kept in a Git repository hosted at GitHub golang.org/x/example. If you include the repository URL in the package's import path, go get will fetch, build, and install it automatically:
```bash
go get golang.org/x/example/hello
$GOPATH/bin/hello
Hello, Go examples!
```

If the specified package is not present in a workspace, `go get` will place it inside the first workspace specified by `GOPATH`. (If the package does already exist, go get skips the remote fetch and behaves the same as `go install`.)

## Variables

```go
var state string = "California"
var state = "California"
state := "California"
```
### Variable Types

Boolean:
- `bool`
  - This refers to either a true or false value

String:
- `string`

Numeric:
- `int int8 int16 int32 int64`
    - Signed integers (e.g., int8 is a signed 8-bit integer with values from -128 to 127)
- `uint uint8 uint16 uint32 uint64 uintptr`
    - Unsigned integers (e.g., uint8 is an unsigned, or positive, 8-bit integer with values from 0 to 255)
- `byte` (alias for uint8)
- `rune` (alias for int32)
- `float32 float64`
  - Used for numbers with decimals
- `complex64 complex128`
    - Used for extremely large numbers

## Conditionals

```go
if CONDITION {
   // Code to be run if condition is true
} else if CONDITION {
  // Code to be run if another condition is true
} else {
  // Code to be run in all other cases
}
```

**Comparison operators**:
```go
==    equal
!=    not equal
<     less
<=    less or equal
>     greater
>=    greater or equal
```

**Switch operator** are like C, C++, Java, JS switch case with the big difference that **only the match case** is executed so 
we don't need to use `break` operator.
```go
switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
```

## Function 

Function signature:
```go
func <name>(<parameters>) <return type> {
    // Body of the function
}
```

## Array and Slices

Overall, the key takeaways of **arrays** are:
- An array is a data structure
- An array contains an ordered list of elements of a single type
- An array contains a specific number of elements
- Because an array is ordered, the first index is 0, the second index is 1… etc.

Array:
```go
var nums = [5]int{1, 2, 3, 4, 5}
```

**Slices**
For the most part, arrays and slices are quite similar, and solve for mostly the same types of use cases. However, one key difference between them is that slices do not have a fixed length.

To recap the takeaways for slices:
- A slice is a data structure
- A slice contains an ordered list of elements of a single type
- A slice contains a list of items that can expand and shrink
    - This is the key difference between a slice and an array!

```go
// Array
var numsArray = [5]int{1, 2, 3, 4, 5}

// Slice
var numsSlice = []int{1, 2, 3, 4, 5}
```

```go
var numsSlice = []int{1, 2, 3, 4, 5}

numsSlice = append(numsSlice, 6)
```

## Loops

```go
for INITIALIZATION; CONDITION; INCREMENT {
    // Code to be repeated
}
```

## Ranges

Ranges allow us to iterate over all the elements with a variety of data structures.

Here's the signature for using a range on an array or slice:
```go
for i, <element> := range <data structure> {
    // Your code here
}
```

Step by step:

- We first use the for keyword. After all, we're still looping through a sequential set of data!
- We follow this by a variable i, referring to the index of the element during iteration.
- Note that this is completely optional. If you don't need access to the index, an underscore _ can be used as a placeholder instead
- We then declare a variable for what we want to name each element in the iteration. This way, we can refer to it and perform any operations we choose - on that element.
- We use the range keyword to signify that we're indeed using a range.
- We provide the name of the data structure onto which we are looking to iterate through. In this case, it would be the variable name pointing to the - array or slice
- And finally, within the curly braces, we include our custom code that we want to run!

## Maps

Items in map are not ordered and are accessible by a key.

Maps are data structures with key-value pairs. That is, for every key, there is an associated value. Additionally:
- Maps are unordered. That is, there is no "index" or position that each value is associated with. Values are - associated to their assigned key
- Keys are unique. Since values are accessed by referring to their key, a map must have unique keys. Note the - difference with an array or slice -- where an array (or a slice) can have elements that are the same! (e.g., var `mySlice = []in{1, 2, 2, 3})`
- Lookup is fast! In an array, it's possible that we may have to iterate through every single element to find - something. With a map, as long as we know the key, we have near-instant access to its associated value

Signature of a map:
```go
map[key_type]value_type{key1: value1, key2: value2, ...}
```

Map demo:
```go
courses := map[uint16]string{

        1: "Calculus",
        2: "Biology",
        3: "Chemistry",
        4: "Computer Science",
        5: "Communications",
        7: "English",
        8: "Cantonese",
    }

    // Using a range, we can through the "courses" map (just like we can for arrays and slices)
    // Rather than an index and element, we are declaring variables for a key (id) and a value (course) 
    for id, course := range courses {
        // If the course, which is a string, begins with the letter "C," it will be printed to the console
        if strings.HasPrefix(course, "C") {
            fmt.Println(id, course)
        }
    }

	// Removing an item from the "courses" map via the built-in "delete" function
	delete(courses, 1)
```

## Structs

Struct are similar to classes, they provide a way to reuse things using a common template and instanciate it with different values.

Structs have the following signature:
```go
type <name> struct {
  // field1 type
  // field2 type
  // field3 type
  // ...
}
```

We cannot define **methods** like in classes but instead we can attach function of 2 different types:
- **value receiver** can only access value in read   
- **pointer receiver** has access to the pointer so they can also modify values

```go
import "srtconv"
type Car struct {
	make string
	year string
	used bool
}

// This only can do read access
func (c Car) describe() string {
	used := ""
	if c.used {
		used = "a used car"
	} else {
		used = "a new car" 
	}
	return "This " + strconv.Itoa(int(c.year)) + " " + c.make + "is " + 
}
```

## Interface

In Go, an interface is custom type that we can use to list out one or more methods, along with their method signatures. But unlike what we’ve seen with a struct, it’s a bit more abstract. That is, we don’t quite create an instance of an interface. While it’s a custom type, it’s really still just a collection of methods that we can refer to.

Interface signature
```go
type <name> interface {
	method1 return_type
	method2 return_type
	method3 return_type
}
```

Example:
```go
package main

import "fmt"

// A "Baseball" struct with two fields
type Baseball struct {
    mass         int
    acceleration int
}

// A "Football" struct with no fields
type Football struct {
}

// A "getForce" method that accepts a "Baseball" type
func (b Baseball) getForce() int {
    return b.mass * b.acceleration
}

// A "getForce" method that accepts a "Football" type
func (f Football) getForce() int {
    return 50
}

// A "Force" interface
// Any type -- such as the Baseball or Football -- that satisfies this interface is said to implement it
// Because both the Baseball and Football structs indeed satisfy the interface via their "getForce" methods...
// ... both the Baseball type and Football type implement the "Force" interface
type Force interface {
    getForce() int
}

// As such, rather than having separate "compareForce" functions, both the Baseball and Football's forces...
// ... can now be compared by used this common compareForce function, since both types implement the Force interface
func compareForce(a, b Force) bool {
    return a.getForce() > b.getForce()
}

func main() {
    // Create an instance of each type
    b := Baseball{2, 20}
    f := Football{}

    // Print the return value of invoking "compareForce" after passing in each instance
    fmt.Println(compareForce(b, f))
}
```

## Goroutines

A goroutine is a function that executes simultaneously and independently of other Goroutines in the program code. This allows for concurrency in Go. In fact, just about any concurrently-executing activity processes in Go can be referred to as a Goroutine.

To create a new Goroutine you just need to add the keyword `go` in front of a function:
```go
func myFunction(){

}
go myFunction()
```

Demo: 
```go
package main

import (
    "fmt"
    "time"
)

func showMessage(message string) {
    for i := 0; i < 5; i++ {
        // "Pause" for one second, then...
        time.Sleep(1 * time.Second)
        // ... print the message to the console
        fmt.Println(message)
    }
}

func main() {
    // Since this is a Goroutine, this will run concurrently (i.e., simultaneous and indpenedent of other program code)
    // As such, this Goroutine will not wait until it is finished executing
    // As a result, the `showMessage` call below will begin running at about the same time as the previous call!
    go showMessage("Go is a great programming language.")

    // This function call does not require the above function call to finish before running itself
    showMessage("Welcome, Gophers!")
}

```

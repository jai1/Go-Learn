# Go-Learn
Few Sample Go applications

In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.pizza and pi do not start with a capital letter, so they are not exported.
When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

package main

import "fmt"

func add10(x int, y int) int {
	var z int = 5 // explicitly declare the data type
	a := 5 // implicit data type detection
	fmt.Printf("a is of type %T\n", a)
	return x + y + z + a
}

func main() {
	fmt.Println(add10(42, 13))
}

A return statement without arguments returns the named return values. This is known as a "naked" return.
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var i, j int = 1, 2
func main() {
	fmt.Println(split(17))

	// 
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}




////////////////////////////////////////////////////////
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.

/////////////////////////////////////////////////////////////////
Variables declared without an explicit initial value are given their zero value.
Constants can't take zero value
The zero value is:
- 0 for numeric types,
- false for the boolean type, and
- "" (the empty string) for strings.

package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
/////////////////////////////////////////////////////////////
The expression T(v) converts the value v to the type T.

Some numeric conversions:

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
Or, put more simply:

i := 42
f := float64(i)
u := uint(f)
///////////////////////////////////////////////////////////
Constants are declared like variables, but with the const keyword.
Constants can be character, string, boolean, or numeric values.
Constants can be implicitly infered but we need (=) to declare it, can't use (:=)
An untyped constant takes the type needed by its context.


package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
	
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// 	fmt.Println(needInt(Big)) // error since int is 64 bit while Big requires 100 bits
}
///////////////////////////////////////////////////////
package main

import "fmt"

func main() {
	sum := 0
	// for init; condition; post
	// The init and post statement are optional.
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	/*
	// Infinte for loop
	for {
	}
	*/
}
/////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// v := math.Pow(x, n);
	// if v < lim {
	// Like for, the if statement can start with a short statement to execute before the condition.
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
//////////////////// Switch /////////////////////////
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	// switch like if can take initializing statement
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
		// cases don't require break
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
//////////////////////// DEFER 
A defer statement defers the execution of a function until the enclosing function returns.
The defer statement execution is bottom up. Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.


package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("jai")
	defer fmt.Println("mohit")

	fmt.Println("hello")
}
/*
hello
mohit
jai
world
*/

///////////////////// POINTERS //////////////////
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
	
	z := &p
	fmt.Println(z) // see the new value of z
}

//////////////////// STRUCTS ///////////////////////
Struct fields can be accessed through a struct pointer.

To access the field X of a struct when we have the struct pointer p we could write (*p).X. 
However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{10, -2})
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	v = Vertex{X: 1}  // jist initializing X to 1 - Y:0 is implicit

	p := &v
	p.X = 1e9
	fmt.Println(v)
}

///////////////////////////// ARRAYS AND SLICES
An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. 
In practice, slices are much more common than arrays.
Slices are reference to an array hence mutable i.e Arrays are passed by copy while slices are passed by reference.

package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	
	var s []int = primes[1:4]
	fmt.Println(s)
	
	// Changing a slice value reflects on the original array
	s[0] = 1 // index 0 of s which is index 1 of primes will be changed
	fmt.Println(primes)
}

// SLICE LITERALS - A slice literal is like an array literal without the length.
package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	
	// make (dataType, length = capacity) all elements initialized to zero value
	sliceInt := make([]int, 5) // make a integer slice of length 5
	
	// A nil slice has a length and capacity of 0 and has no underlying array.
	var nilSlice []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// Confusing make syntax
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4

package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

		
	// make (dataType, length, capacity)
	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

/// Appending to a slice and iterating
The range form of the for loop iterates over a slice or map.
When ranging over a slice, two values are returned for each iteration. 
The first is the index, and the second is a copy of the element at that index.

package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
// Appending to a slice
package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
	
	for i, v := range s {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	
	pow := make([]int, 10)
	// Drop the value
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	
	// Drop the range
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
// 2D Slice
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    pic := make([][]uint8, dy) /* type declaration */
    for i := range pic {
        pic[i] = make([]uint8, dx) /* again the type? */
        for j := range pic[i] {
            pic[i][j] = uint8(i*j)
        }
    }
    return pic
}

func main() {
	pic.Show(Pic)
}


///////////////////// Map ///////////////
The zero value of a map is nil. A nil map has no keys, nor can keys be added.
The make function returns a map of the given type, initialized and ready for use.
package main
Map literals are like struct literals, but the keys are required.


import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	m1 = make(map[string]Vertex)
	m1["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m1["Bell Labs"])
	
	var m2 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
	}
	var m3 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
	}
}
//////// Functions and closures
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}










Why use Go over Java?
https://javax0.wordpress.com/2016/04/27/comparing-golang-with-java/
- Smaller Objects since Go doesn't require Object Header like Java
- can put structs on Stack
- faster GC but no compaction
Golang has a garbage collection but this is not a full GC as in Java, there is no memory compaction. It is not necessarily bad. It can go a long way running servers very long time and still not having the memory fragmented. Some of the JVM garbage collectors also skip the compacting steps to decrease GC pause when cleaning old generations and do the compacting only as a last resort. This last resort step in Go is missing and it may cause some problem is rare cases. You are not likely to face the problem while learning the language.
- can return two values
- Kind of between OO and functional 
- Threads and channels are built in
Threads and queues are built into the language, there is a drastic distinctions between Java and Go concurrency models and how they handle threads. They are called goroutines and channels. To start a goroutine you only have to write go functioncall() and the function will be started in a different thread. Although there are methods/functions in the standard Go library to lock “objects” the native multi-thread programming is using channels. Channel is a built-in type in Go that is a fixed size FIFO channel of any other type. You can push a value into a channel and a goroutine can pull it off. If the channel is full pushing blocks and in case the channel is empty the pull is blocking.

In general Go is an interesting language. It is not a replacement for Java even on a language level. They are not supposed to serve the same type of tasks. Java is enterprise development language, Go is a system programming language. Go, just as well as Java, is continuously developing so we may see some change in that in the future.



Why use Go over C++
- returning reference to temporary variables is OK
- 
Drawbacks of Go

Good things in Go

Why is Go good for multi threading

Why is Go good for distributed systems

///////////////////////////// Plan
2 - 3:30 => Hair cut and reach company


















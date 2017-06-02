package main

import (
	"fmt"
)

func main() {
	a()
	b()
}

func a() {
    for i := 0; i < 4; i++ {
        defer fmt.Print(i)
    }
}

func b() {
    for i := 0; i < 4; i++ {
        defer func() {
		fmt.Print(i)
	}()
    }
}

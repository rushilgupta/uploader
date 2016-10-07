package main

import "fmt"
import "time"

type WorkRequest struct {
	Name  string
	Delay time.Duration
}

func main() {
	fmt.Println("Hello World")
	// Create a new WorkRequest struct instance.
    loc := new(WorkRequest)

    // Set a field and then print its value.
    loc.Name = "boom"
    fmt.Println(loc.Name)
}

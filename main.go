package main

import (
	"fmt"
)

var jobs [2]job

func main() {

	jobs[0] = job{Id: "1", Company: "Amazon", Location: "Boston"}
	jobs[1] = job{Id: "2", Company: "Apple", Location: "Mountain View"}

	fmt.Println("Hello World!")

	runJobController()

}

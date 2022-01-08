package main

import "fmt"

var jobs []Job

func main() {

	jobs = append(jobs, Job{Id: "1", Company: "Amazon", Location: "Boston"})
	jobs = append(jobs, Job{Id: "2", Company: "Apple", Location: "Mountain View"})

	fmt.Println("Starting Server...")
	runJobController()

}

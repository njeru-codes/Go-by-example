/**
worker -pools  -->



*/

package main

import (
	"fmt"
	"time"
)

func worker(id int, job int) {
	fmt.Println("Worker", id, "is working on job", job)
	time.Sleep(1 * time.Second) // pretend it's doing something
	fmt.Println("Worker", id, "finished job", job)
}

func main() {
	jobs := []int{1, 2, 3, 4, 5} // 5 jobs

	for i, job := range jobs {
		go worker(i+1, job) // start a goroutine for each job
	}

	time.Sleep(3 * time.Second) // wait for all workers to finish
}

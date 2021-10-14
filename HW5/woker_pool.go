package main

import "fmt"

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		results <- job
		fmt.Printf("Worker %d finised job %d\n", id, job)
	}
}

func main() {
	numJobs := 10
	numWorkers := 4
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 1; i <= numWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		<-results
	}
}

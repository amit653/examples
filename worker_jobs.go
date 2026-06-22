package main
import (
	"fmt"
	"sync"
	"time"
)
func worker(id int, task <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range task { // task receiver channel from jobs<- 1

		fmt.Printf("Worker %d  started job %d\n", id, j)
		//result := j * 2 // simulate work
		time.Sleep(10 * time.Second) // simualte delay in task completion
		results <- j * 2             // send only channel
		fmt.Printf("Worker %d  finished job %d\n", id, j)
		//results <- result
	}
}
func main() {
	jobs := make(chan int, 3)
	results := make(chan int, 3)
	var wg sync.WaitGroup
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}
	for j := 1; j <= 3; j++ {
		jobs <- j //send only channel
	}
	close(jobs)
	wg.Wait() /// <-- blocks until worker finishes
	close(results)
	for r := range results {
		fmt.Println("results ", r)
	}
}
/*output
Worker 2  started job 2
Worker 1  started job 1
Worker 1  finished job 1
Worker 1  started job 3
Worker 2  finished job 2
Worker 1  finished job 3
results  2
results  4
results  6
*/

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	ID  int
	URL string
}

type Result struct {
	JobID  int
	Status string
}

func worker(
	ctx context.Context,
	id int,
	jobs <-chan Job,
	results map[int]Result,
	mu *sync.Mutex,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d: shutting down\n", id)
			return

		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("worker %d: jobs channel closed\n", id)
				return
			}

			fmt.Printf("worker %d: processing job %d\n", id, job.ID)
			time.Sleep(time.Millisecond * time.Duration(500+rand.Intn(500)))

			res := Result{
				JobID:  job.ID,
				Status: "resized",
			}

			mu.Lock()
			results[job.ID] = res
			mu.Unlock()
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const numWorkers = 4
	const numJobs = 10

	jobs := make(chan Job)
	results := make(map[int]Result)

	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, results, &mu, &wg)
	}

	go func() {
		defer close(jobs)
		for i := 1; i <= numJobs; i++ {
			jobs <- Job{
				ID:  i,
				URL: fmt.Sprintf("image_%d.jpg", i),
			}
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println(">>> cancelling context")
		cancel()
	}()

	wg.Wait()

	fmt.Println("\nFinal Results:")
	mu.Lock()
	for id, res := range results {
		fmt.Printf("Job %d -> %s\n", id, res.Status)
	}
	mu.Unlock()
}
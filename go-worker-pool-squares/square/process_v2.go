package square

import (
	"math"
	"sync"
)

func publishData(numbers []int) <-chan int {
	workerCh := make(chan int)
	go func() {
		defer close(workerCh)

		for _, number := range numbers {
			workerCh <- number
		}
	}()

	return workerCh
}

func ProcessDataWorker(numbers []int, workerCount int) []Data {
	output := make([]Data, len(numbers))
	outputCh := make(chan Data)

	if workerCount <= 0 {
		workerCount = 5
	}

	var wg sync.WaitGroup
	// producer worker
	workerCh := publishData(numbers)

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		// consumer worker
		go func() {
			defer wg.Done()

			for number := range workerCh {
				d := Data{
					Number: number,
					Square: math.Pow(float64(number), 2.0),
				}

				outputCh <- d
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outputCh)
	}()

	for d := range outputCh {
		output = append(output, d)
	}

	return output
}

func worker(_ int, job <-chan int, result chan<- Data) {
	for number := range job {
		d := Data{
			Number: number,
			Square: math.Pow(float64(number), 2.0),
		}

		result <- d
	}

}

func ProcessDataWorkerV2(numbers []int, workerCount int) []Data {
	output := make([]Data, len(numbers))

	if workerCount <= 0 {
		workerCount = 5
	}

	jobCh := make(chan int, len(numbers))
	outputCh := make(chan Data, len(numbers))

	for i := 0; i < workerCount; i++ {
		go worker(i, jobCh, outputCh)
	}

	for _, number := range numbers {
		jobCh <- number
	}

	close(jobCh)

	for i := 0; i < len(numbers); i++ {
		d := <-outputCh
		output = append(output, d)
	}

	close(outputCh)

	return output
}

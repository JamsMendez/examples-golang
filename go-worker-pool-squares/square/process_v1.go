package square

import (
	"math"
	"sync"
)

func ProcessDataGoroutines(numbers []int) []Data {
	outputs := make([]Data, len(numbers))
	outputCh := make(chan Data)

	var wg sync.WaitGroup

	for _, number := range numbers {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			d := Data{
				Number: num,
				Square: math.Pow(float64(num), 2.0),
			}

			outputCh <- d
		}(number)
	}

	go func() {
		wg.Wait()
		close(outputCh)
	}()

	for d := range outputCh {
		outputs = append(outputs, d)
	}

	return outputs
}

func ProcessDataGoroutinesV2(numbers []int) []Data {
	outputs := make([]Data, len(numbers))
	outputCh := make(chan Data)

	for _, number := range numbers {
		go func(num int) {
			d := Data{
				Number: num,
				Square: math.Pow(float64(num), 2.0),
			}

			outputCh <- d
		}(number)
	}

	for i := 0; i < len(numbers); i++ {
		d := <-outputCh
		outputs = append(outputs, d)
	}

	close(outputCh)

	return outputs
}

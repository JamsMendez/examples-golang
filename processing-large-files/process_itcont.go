package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
}

type result struct {
	numRows           int
	peopleCount       int
	commonName        string
	commonNameCount   int
	donationMonthFreq map[string]int
}

// processRow takes a pipe-separated line and returns the firstName, fullName and month
// this functions was created to be somewhat compute intensive and not accurate
func processRow(text string) (firstName, fullName, month string) {
	row := strings.Split(text, "|")

	fullName = strings.ReplaceAll(strings.TrimSpace(row[7]), " ", "")

	name := strings.TrimSpace(row[7])
	if name != "" {
		startOfName := strings.Index(name, ", ") + 2
		if endOfName := strings.Index(name[startOfName:], " "); endOfName < 0 {
			firstName = name[startOfName:]
		} else {
			firstName = name[startOfName : startOfName+endOfName]
		}

		if strings.HasSuffix(firstName, ",") {
			firstName = strings.ReplaceAll(firstName, ",", "")
		}
	}

	date := strings.TrimSpace(row[13])
	if len(date) >= 8 {
		month = date[:2]
	} else {
		month = "--"
	}

	return firstName, fullName, month
}

func sequential(fileName string) result {
	res := result{
		donationMonthFreq: map[string]int{},
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fullNameRegister := make(map[string]bool)
	firstNameMap := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		firstName, fullName, month := processRow(row)

		fullNameRegister[fullName] = true

		firstNameMap[firstName]++

		if firstNameMap[firstName] > res.commonNameCount {
			res.commonName = firstName
			res.commonNameCount = firstNameMap[firstName]
		}

		res.donationMonthFreq[month]++
		res.numRows++
		res.peopleCount = len(fullNameRegister)
	}

	return res
}

// process data file concurrently
func concurrent(fileName string, numWorkers, batchSize int) (res result) {
	res = result{
		donationMonthFreq: map[string]int{},
	}

	type processed struct {
		numRows    int
		fullNames  []string
		firstNames []string
		months     []string
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// reader creates and returns a channel that recieves
	// batchs of rows (of length batchSize) from the file
	reader := func(ctx context.Context, rowsBatch *[]string) <-chan []string {
		out := make(chan []string)

		scanner := bufio.NewScanner(file)

		go func() {
			defer close(out)

			for {
				scanned := scanner.Scan()

				select {
				case <-ctx.Done():
					// finish goroutine
					return

				default:
					row := scanner.Text()
					// if batch size is complete or end of file. send batch out
					if len(*rowsBatch) == batchSize || !scanned {
						out <- *rowsBatch
						// clear batch
						*rowsBatch = []string{}
					}

					// if nothing else to scan return
					if !scanned {
						return
					}

					*rowsBatch = append(*rowsBatch, row)
				}

			}
		}()

		// out is a channel and a pointer
		return out
	}

	// worker takes in a read-only channel to recieve batches of rows
	// after it processes each row-batch it sends out the processed output
	// on its channel
	worker := func(_ context.Context, rowBatch <-chan []string) <-chan processed {
		out := make(chan processed)

		go func() {
			defer close(out)

			p := processed{}

			for rowBatch := range rowBatch {
				for _, row := range rowBatch {
					firstName, fullName, month := processRow(row)
					p.fullNames = append(p.fullNames, fullName)
					p.firstNames = append(p.firstNames, firstName)
					p.months = append(p.months, month)
					p.numRows++
				}
			}

			out <- p
		}()

		return out
	}

	// combiner takes in multiples read-only channels that receive processed output
	// (from workers) and sends it out on its own channel via multiplexer
	combiner := func(ctx context.Context, inputs []<-chan processed) <-chan processed {
		out := make(chan processed)

		var wg sync.WaitGroup
		multiplexer := func(p <-chan processed) {
			defer wg.Done()

			for input := range p {
				select {
				case <-ctx.Done():
				case out <- input:
				}
			}
		}

		// add length of input channels to be consumed by multiplexer
		wg.Add(len(inputs))
		for _, input := range inputs {
			go multiplexer(input)
		}

		// close channel after all inputs channels are closed
		go func() {
			wg.Wait()
			close(out)
		}()

		return out
	}

	// create main context, and call cancel at the end, to ensure all out
	// goroutines exit leaving leaks.
	// particularly, if this function becomes part of a program with a longer
	// lifetime that this function
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// STAGE 1: start reader
	rowsBatch := []string{}
	rowsCh := reader(ctx, &rowsBatch)

	// STAGE 2: create a slice of processed output channels with size of numWorkers
	// add assign each slot with the out channel from each worker
	workerCh := make([]<-chan processed, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workerCh[i] = worker(ctx, rowsCh)
	}

	firstNameCount := map[string]int{}
	fullNameCount := map[string]bool{}

	// STAGE 3: read from the combined cnannel and calculate the final result
	// this will end once all channels from workers are closed
	for processed := range combiner(ctx, workerCh) {
		res.numRows += processed.numRows

		// add months processed by worker
		for _, month := range processed.months {
			res.donationMonthFreq[month]++
		}

		// use full names to count people
		for _, fullName := range processed.fullNames {
			fullNameCount[fullName] = true
		}

		res.peopleCount = len(fullNameCount)

		// update most common first name based on processed results
		for _, firstName := range processed.firstNames {
			firstNameCount[firstName]++

			if firstNameCount[firstName] > res.commonNameCount {
				res.commonName = firstName
				res.commonNameCount = firstNameCount[firstName]
			}
		}
	}

	return
}

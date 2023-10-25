package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"

	"github.com/pkg/profile"
)

func runAsyncPipelineWorkPool() {
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	//defer profile.Start(profile.MemProfile, profile.MemProfileRate(1), profile.ProfilePath(".")).Stop()

	// defer the renaming of the profile file
	defer func() {
		err := os.Rename("trace.out", "traceAsyncWorkerPool.out")
		if err != nil {
			log.Fatal(err)
		}
	}()

	for _, pipeline := range []func(*csv.Reader){asynchronousPipelineWorkPool} {
		file, err := os.Open("data.csv")
		if err != nil {
			panic(err)
		}

		reader := csv.NewReader(file)
		pipeline(reader)
		file.Close()
	}
}

func workerPoolPipelineState[in any, out any](input <-chan in, output chan<- out, process func(in) out, numWorkers int) {
	defer close(output)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for data := range input {
				output <- process(data)
			}
		}()
	}

	wg.Wait()
}

func asynchronousPipelineWorkPool(input *csv.Reader) {
	parseInputCh := make(chan []string)
	convertInputCh := make(chan Record)
	encodeInputCh := make(chan Record)

	outputCh := make(chan []byte)

	done := make(chan struct{})

	numWorkers := 2

	go workerPoolPipelineState(parseInputCh, convertInputCh, parse, numWorkers)
	go workerPoolPipelineState(convertInputCh, encodeInputCh, convert, numWorkers)
	go workerPoolPipelineState(encodeInputCh, outputCh, encode, numWorkers)

	go func() {
		file, err := os.OpenFile("converted.json", os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		encode := json.NewEncoder(file)

		for data := range outputCh {
			var record Record
			if err := json.Unmarshal(data, &record); err != nil {
				panic(err)
			}

			if err := encode.Encode(record); err != nil {
				panic(err)
			}
		}

		close(done)
	}()

	// ignore the first row or header
	input.Read()

	for {
		rec, err := input.Read()
		if err == io.EOF {
			close(parseInputCh)
			break
		}

		if err != nil {
			panic(err)
		}

		parseInputCh <- rec
	}

	<-done
}

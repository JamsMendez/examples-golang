package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Record struct {
	Row    int     `json:"row"`
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
}

func runAsyncPipeline() {
	for _, pipeline := range []func(*csv.Reader){asynchronousPipeline} {
		file, err := os.Open("data.csv")
		if err != nil {
			panic(err)
		}

		reader := csv.NewReader(file)
		pipeline(reader)
		file.Close()
	}
}

func newRecord(in []string) (Record, error) {
	var rec Record
	var err error

	rec.Row, err = strconv.Atoi(in[0])
	if err != nil {
		return rec, err
	}

	rec.Height, err = strconv.ParseFloat(in[1], 64)
	if err != nil {
		return rec, err
	}

	rec.Weight, err = strconv.ParseFloat(in[2], 64)
	if err != nil {
		return rec, err
	}

	return rec, err
}

func parse(input []string) Record {
	rec, err := newRecord(input)
	if err != nil {
		panic(err)
	}

	return rec
}

func convert(input Record) Record {
	input.Height = 2.54 * input.Height
	input.Weight = 0.454 * input.Weight
	return input
}

func encode(input Record) []byte {
	data, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}

	return data
}

func pipelineStage[in any, out any](input <-chan in, output chan<- out, process func(in) out) {
	defer close(output)

	for data := range input {
		output <- process(data)
	}
}

func asynchronousPipeline(input *csv.Reader) {
	parseInputCh := make(chan []string)
	convertInputCh := make(chan Record)
	encodeInputCh := make(chan Record)

	outputCh := make(chan []byte)
	done := make(chan struct{})

	go pipelineStage(parseInputCh, convertInputCh, parse)
	go pipelineStage(convertInputCh, encodeInputCh, convert)
	go pipelineStage(encodeInputCh, outputCh, encode)

	go func() {
		for data := range outputCh {
			s := string(data)
			fmt.Println(s)
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

		// send input to pipeline
		parseInputCh <- rec

		// wait until the last output is printed
		<-done
	}
}

package main

import (
	"errors"
	"fmt"
	"log"
)

type PipelineData struct {
	Input   any
	Error   error
	Context string
}

var (
	ErrorValidation = errors.New("validation error")
	ErrorProcessing = errors.New("processing error")
)

func runPipelineData() {
	data := []any{
		"Input 1",
		"Input 2",
		"Validation Error",
		"Processing Error",
		"Input 5",
	}

	for _, input := range data {
		result := validationStage(input)
		result = processingState(result)
		responseState(result)
	}
}

func validationStage(input any) PipelineData {
	if input == "Validation Error" {
		return PipelineData{
			Input:   input,
			Error:   ErrorValidation,
			Context: "Validation State",
		}
	}

	return PipelineData{
		Input:   input,
		Context: "Validation State",
	}
}

func processingState(data PipelineData) PipelineData {
	if data.Error != nil {
		return data
	}

	if data.Input == "Processing Error" {
		return PipelineData{
			Input:   data.Input,
			Error:   ErrorProcessing,
			Context: "Processing State",
		}
	}

	return PipelineData{Input: data.Input, Context: "Processing State"}
}
func responseState(data PipelineData) {
	if data.Error != nil {
		log.Printf("Error: %s, Input: %s, Context: %s", data.Error, data.Input, data.Context)
		return
	}

	fmt.Printf("Successfully processed input: %v, Context: %s\n", data.Input, data.Context)
}

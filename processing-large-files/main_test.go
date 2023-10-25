package main

import (
	"fmt"
	"reflect"
	"testing"
)

func BenchmarkProcessCSV(b *testing.B) {
	b.ResetTimer()

	fileName := "./students.csv"
	for i := 0; i < b.N; i++ {
		processSequentialFile(fileName)
	}
}

func BenchmarkProcessConcurrentCSV(b *testing.B) {
	b.ResetTimer()

	fileName := "./students.csv"
	for i := 0; i < b.N; i++ {
		processConcurrentFile(fileName)
	}
}

func TestProcessItCont(t *testing.T) {
	tableTest := []struct {
		file     string
		expected result
	}{
		{
			file: "./itcont_sample_40.txt",
			expected: result{
				numRows:           40,
				peopleCount:       35,
				commonName:        "LAURA",
				commonNameCount:   4,
				donationMonthFreq: map[string]int{"01": 7, "02": 2, "03": 6, "04": 2, "05": 3, "06": 2, "07": 1, "08": 2, "11": 15},
			},
		},
		{
			file: "./itcont_sample_4000.txt",
			expected: result{
				numRows:           4000,
				peopleCount:       35,
				commonName:        "LAURA",
				commonNameCount:   400,
				donationMonthFreq: map[string]int{"01": 700, "02": 200, "03": 600, "04": 200, "05": 300, "06": 200, "07": 100, "08": 200, "11": 1500},
			},
		},
	}

	t.Run("sequential", func(t *testing.T) {
		for _, item := range tableTest {
			got := sequential(item.file)
			if !reflect.DeepEqual(item.expected, got) {
				t.Errorf("expected %v, want %v", got, item)
			}
		}
	})

	t.Run("concurrent", func(t *testing.T) {
		for _, item := range tableTest {
			got := concurrent(item.file, 4, 10)
			if !reflect.DeepEqual(item.expected, got) {
				t.Errorf("expected %v, want %v", got, item)
			}
		}
	})

}

func BenchmarkMain(b *testing.B) {
	tableBenchmarks := []struct {
		name    string
		file    string
		inputs  [][]int
		benchFn func(file string, numRows, batchSize int) result
	}{
		{
			name:   "Sequential",
			file:   "./itcont_sample_4000.txt",
			inputs: [][]int{{0, 0}},
			benchFn: func(file string, numWorkers, batchSize int) result {
				return sequential(file)
			},
		},
		{
			name:   "Concurrent",
			file:   "./itcont_sample_4000.txt",
			inputs: [][]int{{1, 1}, {1, 1000}, {10, 1000}, {10, 10000}, {10, 100000}},
			benchFn: func(file string, numWorkers, batchSize int) result {
				return concurrent(file, numWorkers, batchSize)
			},
		},
	}

	for _, item := range tableBenchmarks {
		for _, input := range item.inputs {
			numWorkers := input[0]
			batchSize := input[1]

			bName := fmt.Sprintf("%s %03d workers %04d batchSize", item.name, numWorkers, batchSize)
			b.Run(bName, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					item.benchFn(item.file, numWorkers, batchSize)
				}
			})
		}
	}
}

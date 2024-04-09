package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filename := ""
	chunkSize := 4096

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	chunkChan, errChan := producerChuck(file, chunkSize)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatalln(err)
			}
		case chunk := <-chunkChan:
			fmt.Println("Chunk Size: ", len(chunk))
		}
	}
}

// Chunk is a chunk of file
type Chunk []byte

type State struct {
	buffer []byte
}

func producerChuck(r io.Reader, chunkSize int) (<-chan Chunk, <-chan error) {
	chunkChan := make(chan Chunk)
	errChan := make(chan error)

	go func() {
		defer close(chunkChan)

		reader := bufio.NewReader(r)
		var state State

		for {
			buffer := make([]byte, chunkSize)
			n, err := reader.Read(buffer)
			if err != nil {
				if !errors.Is(err, io.EOF) {
					errChan <- err
				}

				break
			}

			chunk, empty := polishChunk(buffer, n, &state)
			if !empty {
				chunkChan <- chunk
			}
		}

		if len(state.buffer) > 0 {
			chunkChan <- state.buffer
		}
	}()

	return chunkChan, errChan
}

func polishChunk(buffer []byte, size int, state *State) (Chunk, bool) {
	chunk := buffer[:size]
	// concatenate the previous buffer
	chunk = append(chunk, state.buffer...)
	// restart buffer state
	state.buffer = make([]byte, 0)

	lastNewLineIndex := bytes.LastIndexByte(chunk, byte('\n'))
	if lastNewLineIndex != -1 {
		// store chunk remaining
		state.buffer = append(state.buffer, chunk[lastNewLineIndex+1:]...)

		// send chunk to new line
		return chunk[:lastNewLineIndex+1], false
	}

	// not found new line, store chunk remaining
	state.buffer = append(state.buffer, chunk...)

	// return buffer empty
	return make([]byte, 0), true
}

package fetch

import (
	"errors"
	"io"
	"net/http"
)

type Worker struct {
	id         int
	taskQueue  <-chan string
	resultChan chan<- Result
}

func (w *Worker) Start() {
	go func() {
		for url := range w.taskQueue {
			data, err := fetchAndProcess(url)
			result := Result{workerID: w.id, url: url, data: data, err: err}

			w.resultChan <- result
		}
	}()
}

func fetchAndProcess(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to fetch the URL")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

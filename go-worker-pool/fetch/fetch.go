package fetch

import "log"

func Run() {
	urls := []string{}

	workerPool := NewWorkerPool(3)
	workerPool.Start()

	for _, url := range urls {
		workerPool.Submit(url)
	}

	for i := 0; i < len(urls); i++ {
		result := workerPool.Results()
		if result.err != nil {
			log.Println("error fetching URL:", result.url, "error:", result.err)

			continue
		}

		// storage data
	}
}

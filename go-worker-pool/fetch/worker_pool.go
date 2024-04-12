package fetch

type WorkerPool struct {
	taskQueue   chan string
	resultChan  chan Result
	workerCount int
}

func NewWorkerPool(workerCount int) *WorkerPool {
	return &WorkerPool{
		taskQueue:   make(chan string),
		resultChan:  make(chan Result),
		workerCount: workerCount,
	}
}

func (w *WorkerPool) Start() {
	for i := 0; i < w.workerCount; i++ {
		worker := Worker{id: i, taskQueue: w.taskQueue, resultChan: w.resultChan}
		worker.Start()
	}
}

func (wp *WorkerPool) Submit(url string) {
	wp.taskQueue <- url
}

func (w *WorkerPool) Results() Result {
	return <-w.resultChan
}

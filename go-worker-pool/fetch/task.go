package fetch

type Task func(url string) (string, error)

type Result struct {
	workerID int
	url      string
	data     string
	err      error
}

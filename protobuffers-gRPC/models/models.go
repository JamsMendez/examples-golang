package models

type Student struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type Test struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Question struct {
	ID       string `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
	TestID   string `json:"testId"`
}

type Enrollment struct {
	StudentID string
	TestID    string
}

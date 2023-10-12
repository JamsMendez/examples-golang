package factory

import "fmt"

func Run() {
	m1, _ := newPublication("magazine", "M name 1", "M publisher 1", 100)
	m2, _ := newPublication("magazine", "M name 2", "M publisher 2", 75)
	n1, _ := newPublication("newspaper", "N name 1", "N publisher 1", 100)
	n2, _ := newPublication("newspaper", "N name 2", "N publisher 2", 100)

	getDetails(m1)
	getDetails(m2)
	getDetails(n1)
	getDetails(n2)
}

func getDetails(p Publication) {
	fmt.Println(p.getName())
}

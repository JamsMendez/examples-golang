package arraysslices

// [01]
/* func Sum(numbers [5]int) (sum int) {
  size := len(numbers)
  for i := 0; i < size; i++ {
    sum += numbers[i]
  }

  return
} */

func Sum(numbers []int) (sum int) {
	for _, v := range numbers {
		sum += v
	}

	return
}

// [02]
// func SumAll(numbersToSum ...[]int) []int {
/* func SumAll(numbersToSum [][]int) []int {
	lenOfNumbers := len(numbersToSum)
	sumList := make([]int, lenOfNumbers)

	for i, s := range numbersToSum {
		sumList[i] = Sum(s)
	}

	return sumList
} */

func SumAll(numbersOfSum [][]int) []int {
	var sums []int

	for _, numbers := range numbersOfSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersOfSum [][]int) []int {
	var sumsTails []int

	for _, numbers := range numbersOfSum {
		lengthNumbers := len(numbers)

		if lengthNumbers == 0 {
			sumsTails = append(sumsTails, 0)
		} else {
			tail := numbers[1:]
			sumsTails = append(sumsTails, Sum(tail))
		}
	}

	return sumsTails
}

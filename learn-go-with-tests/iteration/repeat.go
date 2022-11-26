package iteration

import (
	"fmt"
)

const repeatCount = 5

func Repeat(character string) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated = fmt.Sprintf("%s%s", repeated, character)
	}

	return
}

package file

import (
	"fmt"
	"log"
	"os"
)

func OpenCloseFile() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	n, err := file.WriteString("Jose Angel")
	if os.IsPermission(err) {
		fmt.Println(err)
	}

	fmt.Println(n, err)

	file.Close()

	// os_O_XXX | os.O_XXX

	// os.O_RDONLY  Read only
	// os.O_WRONLY  Write only
	// os.O_RDWR 	Read and write
	// os.O_APPEND	Append to end of file
	// os.O_CREATE  Create is none exist
	// os.O_TRUNC   Truncate file when opening, empty the entire file before writing

	// file, err = os.OpenFile("file.txt", os.O_RDONLY, 0o666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// defer file.Close()
}

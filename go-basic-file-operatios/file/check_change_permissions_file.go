package file

import (
	"log"
	"os"
	"time"
)

func CheckPermissionsFile() {
	file, err := os.OpenFile("file.txt", os.O_WRONLY, 0o666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("write permission denied")
		}
	}

	file.Close()

	file, err = os.OpenFile("file.txt", os.O_RDONLY, 0o666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("write permission denied")
		}
	}

	file.Close()
}

func ChangePermissionsFile() {
	err := os.Chmod("file.txt", 0o777)
	if err != nil {
		log.Println(err)
	}

	err = os.Chown("file.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Println(err)
	}

	timestamp := time.Now().Add(48 * time.Hour)
	lastAccessTime := timestamp
	lastModifyTime := timestamp
	err = os.Chtimes("file.txt", lastAccessTime, lastModifyTime)
	if err != nil {
		log.Println(err)
	}
}

package cmd

import (
	"fmt"
	"log"
	"os"
)

func Install(version, dstInstall string) error {
	fmt.Println("Downloading ...")

	err := downloadFile(version)
	if err != nil {
		return err
	}

	fmt.Println("Downloading complete")

	var file *os.File
	file, err = getFile(version)
	if err != nil {
		return err
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Println("install file close: ", err)
		}

		err = removeFile(version)
		if err != nil {
			log.Println("install file remove: ", err)
		}
	}()

	fmt.Printf(`Decompress "%s" in "%s" ...%s`, file.Name(), dstInstall, "\n")

	err = decompress(dstInstall, file)
	if err != nil {
		if errDir := removeGoDir(dstInstall); errDir != nil {
			log.Println(errDir)
		}

		return err
	}

	err = removeFile(version)
	if err != nil {
		log.Println("install file remove: ", err)
	}

	fmt.Println("Decompress complete")

	return err
}

func ListVersion() (string, error) {
	list, err := fetchVersions()
	if err != nil {
		return "", err
	}

	return list.String(), nil
}

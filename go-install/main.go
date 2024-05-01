package main

import (
	"flag"
	"fmt"
	"log"

	"go-install/cmd"
)

func main() {
	version := flag.String("version", "", "-version specify the version of go to download and install")
	destination := flag.String("d", cmd.DstInstall, "-d specify the destination of the go installation")

	flag.Parse()

	if *version == "" {
		s, err := cmd.ListVersion()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("-version is empty. Stable versions:\n%s\n", s)
		return
	}

	fmt.Printf("Starting Go installation version %s ...\n", *version)

	fmt.Println(*version, *destination)

	err := cmd.Install(*version, *destination)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Go Installation complete ... :D!")
}

package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	tmplURL = "https://go.dev/dl/go%s.linux-amd64.tar.gz"
)

func getURL(version string) (*url.URL, error) {
	uri := fmt.Sprintf(tmplURL, version)
	return url.Parse(uri)
}

func downloadFile(version string) error {
	uri, err := getURL(version)
	if err != nil {
		return err
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("installer version not found")
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Println("response body close: ", err)
		}
	}()

	file, err := getFile(version)
	if err != nil {
		return err
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Println("download file close: ", err)
		}
	}()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		if errrf := removeFile(version); errrf != nil {
			log.Println("download file remove: ", errrf)
		}

		return err
	}

	return nil
}

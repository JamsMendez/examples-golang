package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	urlAPI = "https://golang.org/dl/?mode=json"
)

type listStableVersion []stableVersion

func (l *listStableVersion) String() string {
	var elems []string
	for _, item := range *l {
		elems = append(elems, item.Version)
	}

	return strings.Join(elems, "\n")
}

type stableVersion struct {
	Version string `json:"version"`
}

func fetchVersions() (listStableVersion, error) {
	uri, err := url.Parse(urlAPI)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("golang api version not found")
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Println("golang api response body close: ", err)
		}
	}()

	var list listStableVersion

	err = json.NewDecoder(resp.Body).Decode(&list)
	if err != nil {
		log.Println("golang api decode json error: ", err)
	}

	return list, nil
}

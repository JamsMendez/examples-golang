package factory

import "fmt"

func newPublication(pubType, name, pub string, pg int) (Publication, error) {
	if pubType == "newspaper" {
		return createNewspaper(name, pub, pg), nil
	}

	if pubType == "magazine" {
		return createMagazine(name, pg, pub), nil
	}

	return nil, fmt.Errorf("no such publication type")
}

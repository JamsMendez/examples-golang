package factory

import "fmt"

type magazine struct {
	publication
}

func (m magazine) String() string {
	return fmt.Sprintf("Magazine name: %s", m.name)
}

func createMagazine(name string, pages int, publisher string) Publication {
	return &magazine{
		publication{
			name:      name,
			pages:     pages,
			publisher: publisher,
		},
	}
}

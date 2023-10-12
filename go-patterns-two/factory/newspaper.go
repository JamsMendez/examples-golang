package factory

import "fmt"

type newspaper struct {
	publication
}

func (n newspaper) String() string {
	return fmt.Sprintf("newspaper name: %s", n.name)
}

func createNewspaper(name, publisher string, pages int) Publication {
	return &newspaper{
		publication{
			name:      name,
			pages:     pages,
			publisher: publisher,
		},
	}
}

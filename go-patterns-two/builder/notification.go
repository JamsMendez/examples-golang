package builder

type Notification struct {
	title    string
	subtitle string
	message  string
	image    string
	icon     string
	priority int
	notType  string
}

func (n *Notification) GetTitle() string {
	return n.title
}

func (n *Notification) GetMessage() string {
	return n.message
}

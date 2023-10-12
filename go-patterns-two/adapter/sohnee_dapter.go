package adapter

type sammysangAdapter struct {
	isOn bool
	st   *SammysangTV
}

func (s sammysangAdapter) volumeUp() int {
	volume := s.st.getVolume() + 1
	s.st.setVolume(volume)

	return s.st.getVolume()
}

func (s sammysangAdapter) volumeDown() int {
	volume := s.st.getVolume() - 1
	s.st.setVolume(volume)

	return s.st.getVolume()
}

func (s sammysangAdapter) turnOn() {
	s.isOn = true

}

func (s sammysangAdapter) turnOff() {
	s.isOn = false
}

package adapter

type SammysangTV struct {
	currentChan   int
	currentVolume int
	tvOn          bool
}

func (tv *SammysangTV) getVolume() int {
	return tv.currentVolume
}

func (tv *SammysangTV) setVolume(volume int) {
	tv.currentVolume = volume
}

func (tv *SammysangTV) getChannel() int {
	return tv.currentChan
}

func (tv *SammysangTV) setChannel(channel int) {
	tv.currentChan = channel
}

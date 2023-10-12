package adapter

type SohneeTV struct {
	vol     int
	channel int
	isOn    bool
}

func (st *SohneeTV) turnOn() {
	st.isOn = true
}

func (st *SohneeTV) turnOff() {
	st.isOn = false
}

func (st *SohneeTV) volumeUp() int {
	st.vol++
	return st.vol
}

func (st *SohneeTV) volumeDown() int {
	st.vol--
	return st.vol
}

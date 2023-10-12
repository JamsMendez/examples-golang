package adapter

func Run() {
	tv1 := &SammysangTV{
		currentChan:   200,
		currentVolume: 45,
		tvOn:          true,
	}

	tv2 := &SohneeTV{
		vol:     48,
		channel: 202,
		isOn:    true,
	}

	// turnOnTv(tv1) not working
	turnOnTv(tv2)

	tv1Adapter := &sammysangAdapter{
		isOn: false,
		st:   tv1,
	}

	turnOnTv(tv1Adapter)
}

func turnOnTv(t television) {
	t.turnOn()
	t.volumeUp()
}

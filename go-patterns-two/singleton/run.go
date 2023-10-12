package singleton

func Run() {
	log := getLoggerInstance()

	log.SetLogLevel(1)
	log.Log("log with message level 1")

	nLog := getLoggerInstance()
	nLog.SetLogLevel(2)
	nLog.Log("nLog with message level 2")

	log.Log("log with message level 2")
}

package st_config_log

import (
	"github.com/kpango/glg"
)

// NetWorkLogger sample network logger
type NetWorkLogger struct{}

func (n NetWorkLogger) Write(b []byte) (int, error) {
	http.Get(VtLogAddr+"/log")
	glg.Success("Requested")
	glg.Infof("RawString is %s", glg.RawString(b))
	return 1, nil
}
var infoLog *os.File
var errLog *os.File

func getTodayDate() string {
	t := time.Now()
	return fmt.Sprintf("%d-%d-%d", t.Year(),
		t.Month(), t.Day())
}

func Start() {
	// var errWriter io.Writer
	// var customWriter io.Writer

	infoLog = glg.FileWriter( "./log/"+getTodayDate()+"/info.log", os.ModeAppend)

	customTag := "FINE"
	customErrTag := "CRIT"

	errLog = glg.FileWriter("./log/"+getTodayDate()+"/error.log", os.ModeAppend)

	glg.Get().
		SetMode(glg.BOTH). // default is STD
		// DisableColor().
		// SetMode(glg.NONE).
		// SetMode(glg.WRITER).
		// SetMode(glg.BOTH).
		// InitWriter().
		// AddWriter(customWriter).
		// SetWriter(customWriter).
		// AddLevelWriter(glg.LOG, customWriter).
		// AddLevelWriter(glg.INFO, customWriter).
		// AddLevelWriter(glg.WARN, customWriter).
		// AddLevelWriter(glg.ERR, customWriter).
		// SetLevelWriter(glg.LOG, customWriter).
		// SetLevelWriter(glg.INFO, customWriter).
		// SetLevelWriter(glg.WARN, customWriter).
		// SetLevelWriter(glg.ERR, customWriter).
		AddLevelWriter(glg.INFO, infoLog). // add info log file destination
		AddLevelWriter(glg.ERR, errLog). // add error log file destination
		AddStdLevel(customTag, glg.STD, false).                    //user custom log level
		AddErrLevel(customErrTag, glg.STD, true).                  // user custom error log level
		SetLevelColor(glg.TagStringToLevel(customTag), glg.Cyan).  // set color output to user custom level
		SetLevelColor(glg.TagStringToLevel(customErrTag), glg.Red) // set color output to user custom level
	/*
		glg.Info("info")
		glg.Infof("%s : %s", "info", "formatted")
		glg.Log("log")
		glg.Logf("%s : %s", "info", "formatted")
		glg.Debug("debug")
		glg.Debugf("%s : %s", "info", "formatted")
		glg.Warn("warn")
		glg.Warnf("%s : %s", "info", "formatted")
		glg.Error("error")
		glg.Errorf("%s : %s", "info", "formatted")
		glg.Success("ok")
		glg.Successf("%s : %s", "info", "formatted")
		glg.Fail("fail")
		glg.Failf("%s : %s", "info", "formatted")
		glg.Print("Print")
		glg.Println("Println")
		glg.Printf("%s : %s", "printf", "formatted")
		glg.CustomLog(customTag, "custom logging")
		glg.CustomLog(customErrTag, "custom error logging")
	*/
	glg.Get().AddLevelWriter(glg.DEBG, NetWorkLogger{}) // add info log file destination

	http.Handle("/glg", glg.HTTPLoggerFunc("glg sample", func(w http.ResponseWriter, r *http.Request) {
		glg.New().
			AddLevelWriter(glg.INFO, NetWorkLogger{}).
			AddLevelWriter(glg.INFO, w).
			Info("glg HTTP server logger")
	}))

	go http.ListenAndServe(VtLogAddr, nil)
}

func End() {
	infoLog.Close()
	errLog.Close()
}
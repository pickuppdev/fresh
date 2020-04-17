package runner

import (
	"fmt"
	logPkg "log"

	"github.com/mattn/go-colorable"
)

type logFunc func(string, ...interface{})

var stdoutLogger = logPkg.New(colorable.NewColorableStdout(), "", 0)
var stderrLogger = logPkg.New(colorable.NewColorableStderr(), "", 0)

func newLogFunc(prefix string) func(string, ...interface{}) {
	color, clear := "", ""
	if settings["colors"] == "1" {
		color = fmt.Sprintf("\033[%sm", logColor(prefix))
		clear = fmt.Sprintf("\033[%sm", colors["reset"])
	}
	prefix = fmt.Sprintf("%-11s", prefix)

	return func(format string, v ...interface{}) {
		format = fmt.Sprintf("%s%s |%s %s", color, prefix, clear, format)
		stdoutLogger.Printf(format, v...)
	}
}

func fatal(err error) {
	stderrLogger.Fatal(err)
}

type appLogWriter struct{}

func (a appLogWriter) Write(p []byte) (n int, err error) {
	appLog(string(p))

	return len(p), nil
}

package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/fatih/color"
)

// Logging Global Vars
var logLevelInt int
var appEnv string
var envLogLevel string

type Log struct {
	Uuid        string
	Message     string
	LongMessage string
	Level       string
	Timestamp   string
	ErrorCode   string
	Module      string
	Error       error
}

func init() {
	// Logging
	// var appEnv = os.Getenv("PHOENIX_APP_ENV")
	// var envLogLevel = os.Getenv("PHOENIX_LOG_LEVEL")
	// var SilentUnusedWS = os.Getenv("PHOENIX_LOG_WS_SILENT")

	//DEBUG move to ENV
	// 'dev' appEnv enables human readable logging
	appEnv = "dev"
	envLogLevel = "5"

	var err error
	logLevelInt, err = strconv.Atoi(envLogLevel)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid Log Level set in env!")
	}
}

func Info(message Log) {
	message.Level = "INFO"

	if logLevelInt >= 4 {
		if appEnv == "dev" {

			message.Timestamp = time.Now().Format(time.RFC3339)

			if message.LongMessage != "" {
				message.LongMessage = color.CyanString(
					"(" + message.LongMessage + ")",
				)
			}
			fmt.Fprintln(
				os.Stdout,
				color.CyanString("["+message.Level+"]"),
				"- "+message.Timestamp+" -",
				message.Uuid,
				"-",
				color.CyanString(message.Message),
				message.LongMessage,
				"<-- "+message.Module,
			)

		} else {
			debugJson, err := json.Marshal(message)
			if err != nil {
				fmt.Println(os.Stderr, "Failed to marshall log message")
				return
			}
			fmt.Fprintln(os.Stdout, string(debugJson))
		}
	}
}

func Error(message Log) {
	message.Level = "ERROR"

	if logLevelInt >= 2 {
		if appEnv == "dev" {

			message.Timestamp = time.Now().Format(time.RFC3339)

			if message.LongMessage != "" {
				message.LongMessage = color.RedString(
					"(" + message.LongMessage + ")",
				)
			}

			if message.Error != nil {
				message.Message = message.Message + " (Error: " + message.Error.Error() + ")"
			}

			fmt.Fprintln(
				os.Stdout,
				color.RedString("["+message.Level+"]"),
				"- "+message.Timestamp+" -",
				message.Uuid,
				"-",
				color.RedString(message.Message),
				message.LongMessage,
				"<-- "+message.Module,
			)

		} else {
			debugJson, err := json.Marshal(message)
			if err != nil {
				fmt.Println(os.Stderr, "Failed to marshall log message")
				return
			}
			fmt.Fprintln(os.Stdout, string(debugJson))
		}
	}
}

func Trace() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])

	return "[" + file + ":" + strconv.Itoa(line) + "] " + f.Name()
}

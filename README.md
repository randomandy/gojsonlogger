# gojsonlogger
A simple logger for Go, supporting JSON and colorful pretty output

# Install
  go get github.com/randomandy/gojsonlogger
  
# Usage

  import "github.com/randomandy/gojsonlogger"

	logger.Info(logger.Log{
		Uuid:        "SYSTEM PROCESS",
		Module:      logger.Trace(),
		Message:     "Foobar",
    LongMessage: "Some more Foobar",
	})
 
	logger.Error(logger.Log{
		Uuid:    "de12592e-3ba8-4908-80e2-94857778d221",
		Module:  logger.Trace(),
		Message: "Some error",
    Error:   err,
	})

## Output

  [INFO] - 2017-05-30T11:15:33+04:00 - SYSTEM PROCESS - Foobar (Some more Foobar) <-- [/Users/andy/checkouts/golang/src/github.com/ASWATFZLLC/claptrap/main.go:59] main.main

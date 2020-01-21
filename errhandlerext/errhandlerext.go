package errhandlerext

import (
	"os"

	"github.com/payfazz/go-errors"
	"github.com/payfazz/stdlog"
)

// LogAndExit is handler for https://godoc.org/github.com/payfazz/go-errors/errhandler#With
func LogAndExit(err error) {
	errors.PrintTo(stdlog.Err, err)
	os.Exit(1)
}

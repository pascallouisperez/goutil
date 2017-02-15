package errors

import (
	go_errors "errors"
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
)

var thisFilePattern = regexp.MustCompile(`util/errors/errors\.go$`)

// New formats the message, and prepends the filename and line number of the
// immediate caller.
func New(format string, args ...interface{}) error {
	message := fmt.Sprintf(format, args...)
	for i := 0; i < 10; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if !thisFilePattern.MatchString(file) {
			_, filename := filepath.Split(file)
			return fmt.Errorf("%s:%d: %s", filename, line, message)
		}
	}
	return go_errors.New(message)
}

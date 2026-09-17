package errs

import "fmt"

func Errorf(function string, format string, a ...any) error {
	return fmt.Errorf(function+": "+format, a...)
}

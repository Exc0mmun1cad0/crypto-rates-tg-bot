package error

import "fmt"

func Wrap(msg string, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %s", msg, err)
	} else {
		return nil
	}
}

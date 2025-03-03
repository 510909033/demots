package fn

import "fmt"

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Recover(fn func()) error {
	if err := recover(); err != nil {
		return fmt.Errorf("recover: %+v", err)
	}
	fn()
	return nil
}

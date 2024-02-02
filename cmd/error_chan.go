package cmd

import (
	"go.uber.org/fx"
)

var errorModule = fx.Module("error", fx.Provide(newErrorChan))

func newErrorChan() chan error {
	return make(chan error)
}

//
//func errorListener(errorChannel chan<- error) {
//	if err := <-errorChannel; err != nil {
//		fmt.Println("Error:", err)
//	} else {
//		fmt.Println("Operation completed successfully.")
//	}
//}

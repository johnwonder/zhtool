package util

import (
	"fmt"
	"os"
)

func PrintErr(str string, a ...interface{}) {
	fmt.Fprintln(os.Stderr,str,a)
}

func Error(str string, a ...interface{}) {
	fmt.Fprintln(os.Stderr,str,a)
}

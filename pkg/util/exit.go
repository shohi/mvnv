package util

import (
	"fmt"
	"os"
)

func PrintErrorAndExit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

package main

import (
	"os"
	"path/filepath"

	"github.com/atrox/homedir"
	"github.com/shohi/mvnv/pkg/wrapper"
)

func main() {
	// MVNV_BIN points to the maven binary
	// where MVNV_HOME directs to the location of
	// default maven version
	var binName = os.Getenv("MVNV_BIN")
	if binName == "" {
		binName = "mvn"
	}

	var mvnvHome = os.Getenv("MVNV_HOME")
	if mvnvHome == "" {
		h, _ := homedir.Dir()
		mvnvHome = filepath.Join(h, ".mvnv")
	} else {
		mvnvHome, _ = homedir.Expand(mvnvHome)
	}

	wrapper.Wrapper(mvnvHome, binName)
}

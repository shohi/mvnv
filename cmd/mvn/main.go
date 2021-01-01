package main

import "github.com/shohi/mvnv/pkg/wrapper"

func main() {
	// TODO: support using env to override
	var binName string = "mvn"

	wrapper.Wrapper(".mvnv", binName)
}

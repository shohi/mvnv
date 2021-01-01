package main

import "github.com/shohi/mvnv/cmd/mvnv/root"

// mvnv will creates a subfolder - `~/.mvnv` to
// hold all related information, including
// - versions/ - unzipped maven binary
// - downloads/ - downloaded maven tar.gz
// - .mvn-version - current activated maven version, which points to `versions`
func main() {
	root.Execute()
}

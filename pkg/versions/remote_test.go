package versions

import "testing"

func TestCollector_Remote(t *testing.T) {
	// TODO: skip

	apiURL := "https://archive.apache.org/dist/maven/maven-3"
	verPrefix := "3"

	c := NewCollector(apiURL, verPrefix)
	c.Remote()
}

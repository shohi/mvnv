package versions

import (
	"fmt"
	"os"
	"sort"

	"github.com/fatih/color"
)

func Print(versions []string, inuseVersion string) {
	// sort by descent order, e.g. [3.6, 3.5, ...]
	sort.Sort(sort.Reverse(sort.StringSlice(versions)))

	for _, v := range versions {
		if v == inuseVersion {
			color.New(color.FgGreen).Fprintf(os.Stdout, "* %s\n", v)
		} else {
			fmt.Fprintf(os.Stdout, "  %s\n", v)
		}
	}
}

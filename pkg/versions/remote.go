package versions

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Collector struct {
	client http.Client
	apiURL string // used to get all versions info

	// only matched version are colleted
	// TODO: use more flexible strategies
	versionPrefix string
}

func NewCollector(apiURL, versionPrefix string) *Collector {
	return &Collector{
		apiURL:        apiURL,
		versionPrefix: versionPrefix,
		client:        http.Client{Timeout: 30 * time.Second},
	}
}

func (c *Collector) Remote() ([]string, error) {
	resp, err := c.client.Get(c.apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return c.parseVersions(content)
}

// parseVersions parses versions from html page.
func (c *Collector) parseVersions(page []byte) ([]string, error) {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(page))

	if err != nil {
		return nil, err
	}

	var versions []string

	// Find the version items
	doc.Find("pre a").Each(func(i int, s *goquery.Selection) {
		val, ok := s.Attr("href")
		if !ok || !strings.HasPrefix(val, c.versionPrefix) {
			return
		}

		ver := strings.TrimSuffix(val, "/")
		if ver != "" {
			versions = append(versions, ver)
		}
	})

	return versions, err
}

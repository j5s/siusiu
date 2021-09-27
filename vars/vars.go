package vars

import "regexp"

var (
	BingURL    = "https://cn.bing.com"
	AtagRe     = regexp.MustCompile(`<a[^>]+href="(http[^>"]+)"[^>]+>`)
	NextPageRe = regexp.MustCompile(`<a[^>]+href="(/search\?q=[^>]+)"[^>]+>`)
	Debug      = false
)

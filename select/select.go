package racer

import (
	"net/http"
	"time"
)

func Racer(url1, url2 string) (fastURL string) {
	durationURL1 := measureResponseTime(url1)
	durationURL2 := measureResponseTime(url2)

	if durationURL1 < durationURL2 {
		return url1
	}

	return url2
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

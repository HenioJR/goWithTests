package racer

import (
	"net/http"
	"time"
)

func Racer(url1, url2 string) (fastURL string) {
	startURL1 := time.Now()
	http.Get(url1)
	durationURL1 := time.Since(startURL1)

	startURL2 := time.Now()
	http.Get(url2)
	durationURL2 := time.Since(startURL2)

	if durationURL1 < durationURL2 {
		return url1
	}

	return url2
}

package racer

import (
	"net/http"
)

func Racer(url1, url2 string) (fastURL string) {
	//constructor select helps us synchronise processes easily and clearly
	//select wait on multiple channels. The first one to send a value "wins"
	select {
	case <-ping(url1):
		return url1
	case <-ping(url2):
		return url2
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}

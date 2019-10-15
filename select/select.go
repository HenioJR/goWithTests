package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(url1, url2 string) (fastURL string, err error) {
	return RacerConfigurableTimeout(url1, url2, tenSecondTimeout)
}

func RacerConfigurableTimeout(url1, url2 string, timeout time.Duration) (fastURL string, err error) {
	//constructor select helps us synchronise processes easily and clearly
	//select wait on multiple channels. The first one to send a value "wins"
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", url1, url2)
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

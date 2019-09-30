package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	//channels are a Go data structure that can both receive and send values
	//channels allow communication between different processes
	resultChannel := make(chan result)

	for _, url := range urls {
		//anonymous function
		go func(u string) {
			//we're sending a result to the result channel
			// <- is a operator to send statement
			resultChannel <- result{u, wc(u)}
		}(url)

	}

	//by sending the results into a channel, we can control the timing of each write into result map
	for i := 0; i < len(urls); i++ {
		//receive expression
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}

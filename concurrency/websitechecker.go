package concurrency

import "time"

type result struct {
	string
	bool
}

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	checkResult := make(chan result)

	for _, url := range urls {
		go func() {
			checkResult <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-checkResult
		results[r.string] = r.bool
	}

	time.Sleep(3 * time.Second)
	return results
}

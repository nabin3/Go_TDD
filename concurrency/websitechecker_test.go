package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "wwa://vwdg.com"
}

func TestCheckWebsites(t *testing.T) {
	websiteList := []string{
		"https://google.co.in",
		"https://flipkart.com",
		"wwa://vwdg.com",
	}

	got := CheckWebsites(mockWebsiteChecker, websiteList)

	want := map[string]bool{
		"https://google.co.in": true,
		"https://flipkart.com": true,
		"wwa://vwdg.com":       false,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %v, but got %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

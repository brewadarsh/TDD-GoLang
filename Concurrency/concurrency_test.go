package concurrency

import (
	"testing"
	"time"
)

func mockValidator(website string) bool {
	time.Sleep(50 * time.Millisecond)
	return true
}

func BenchmarkWebsiteValidator(b *testing.B) {
	urls := make([]string, 100)
	for i := range urls { // Create a large data set.
		urls[i] = "website_url"
	}
	b.ResetTimer()
	for b.Loop() {
		WebsiteValidator(mockValidator, urls)
	}
}

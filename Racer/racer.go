package racer

import "net/http"

// Slower method.
/*func WebsiteRacer(first, second string) string {
	firstStart := time.Now()
	http.Get(first)
	firstDuration := time.Since(firstStart)

	secondStart := time.Now()
	http.Get(second)
	secondDuration := time.Since(secondStart)

	if firstDuration < secondDuration {
		return first
	}
	return second
}*/

// Faster method.
func WebsiteRacer(first, second string) string {
	select {
	case <-ping(first):
		return first
	case <-ping(second):
		return second
	}
}

func ping(website string) chan struct{} {
	channel := make(chan struct{})
	go func() {
		http.Get(website)
		close(channel)
	}()
	return channel
}

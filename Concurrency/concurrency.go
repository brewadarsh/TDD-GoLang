package concurrency

// Interface for website checker.
type Validator func(string) bool

// The validation result used for channel.
type validationResult struct {
	website string
	result  bool
}

// Checks the validity of an website.
func WebsiteValidator(validator Validator, websites []string) map[string]bool {
	results := make(map[string]bool) // Create the empty map.
	channels := make(chan validationResult)
	for _, website := range websites {
		go func() {
			// results[website] = validator(website)
			channels <- validationResult{website: website, result: validator(website)}
		}()
	}
	for range len(websites) {
		r := <-channels
		results[r.website] = r.result
	}
	return results
}

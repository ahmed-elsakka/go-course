package quote

import "math/rand/v2"

var quotes = []string{
	"Keep going, you’re doing great!",
	"Don’t watch the clock; do what it does. Keep going.",
	"Success is not final, failure is not fatal.",
}

func GetQuote() string {
	return quotes[rand.IntN(len(quotes))]
}

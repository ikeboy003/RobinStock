package urls

// EarningsURL returns the URL for earnings
func EarningsURL() string {
	return "https://api.robinhood.com/marketdata/earnings/"
}

// EventsURL returns the URL for events
func EventsURL() string {
	return "https://api.robinhood.com/options/events/"
}

// FundamentalsURL returns the URL for fundamentals
func FundamentalsURL() string {
	return "https://api.robinhood.com/fundamentals/"
}

// HistoricalsURL returns the URL for historicals
func HistoricalsURL() string {
	return "https://api.robinhood.com/quotes/historicals/"
}

// InstrumentsURL returns the URL for instruments
func InstrumentsURL() string {
	return "https://api.robinhood.com/instruments/"
}

// NewsURL returns the URL for news
func NewsURL(symbol string) string {
	return "https://api.robinhood.com/midlands/news/" + symbol + "/?"
}

// PopularityURL returns the URL for popularity
func PopularityURL(symbol string) string {
	return "https://api.robinhood.com/instruments/" + symbol + "/popularity/"
}

// QuotesURL returns the URL for quotes
func QuotesURL() string {
	return "https://api.robinhood.com/quotes/"
}

// RatingsURL returns the URL for ratings
func RatingsURL(symbol string) string {
	return "https://api.robinhood.com/midlands/ratings/" + symbol + "/"
}

// SplitsURL returns the URL for splits
func SplitsURL(symbol string) string {
	return "https://api.robinhood.com/instruments/" + symbol + "/splits/"
}

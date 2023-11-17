package urls

// CurrencyURL returns the URL for currency
func CurrencyURL() string {
	return "https://nummus.robinhood.com/currency_pairs/"
}

// MarketsURL returns the URL for markets
func MarketsURL() string {
	return "https://api.robinhood.com/markets/"
}

// MarketHoursURL returns the URL for market hours
func MarketHoursURL(market, date string) string {
	return "https://api.robinhood.com/markets/" + market + "/hours/" + date + "/"
}

// MoversSP500URL returns the URL for movers sp500
func MoversSP500URL() string {
	return "https://api.robinhood.com/midlands/movers/sp500/"
}

// Get100MostPopularURL returns the URL for get 100 most popular
func Get100MostPopularURL() string {
	return "https://api.robinhood.com/midlands/tags/tag/100-most-popular/"
}

// MoversTopURL returns the URL for movers top
func MoversTopURL() string {
	return "https://api.robinhood.com/midlands/tags/tag/top-movers/"
}

// MarketCategoryURL returns the URL for market category
func MarketCategoryURL(category string) string {
	return "https://api.robinhood.com/midlands/tags/tag/" + category + "/"
}

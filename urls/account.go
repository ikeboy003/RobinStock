package urls

// PhoenixURL returns the URL for phoenix
func PhoenixURL() string {
	return "https://phoenix.robinhood.com/accounts/unified"
}

// PositionsURL returns the URL for positions
func PositionsURL(accountNumber string) string {
	if accountNumber != "" {
		return "https://api.robinhood.com/positions/?account_number=" + accountNumber
	}
	return "https://api.robinhood.com/positions/"
}

// CardTransactionsURL returns the URL for card transactions
func CardTransactionsURL() string {
	return "https://minerva.robinhood.com/history/transactions/"
}

// DayTradesURL returns the URL for day trades
func DayTradesURL(account string) string {
	return "https://api.robinhood.com/accounts/" + account + "/recent_day_trades/"
}

// DividendsURL returns the URL for dividends
func DividendsURL() string {
	return "https://api.robinhood.com/dividends/"
}

// DocumentsURL returns the URL for documents
func DocumentsURL() string {
	return "https://api.robinhood.com/documents/"
}

// WithdrawlURL returns the URL for withdrawl
func WithdrawlURL(bankID string) string {
	return "https://api.robinhood.com/ach/relationships/" + bankID + "/"
}

// LinkedURL returns the URL for linked
func LinkedURL(id string, unlink bool) string {
	if unlink {
		return "https://api.robinhood.com/ach/relationships/" + id + "/unlink/"
	}
	if id != "" {
		return "https://api.robinhood.com/ach/relationships/" + id + "/"
	}
	return "https://api.robinhood.com/ach/relationships/"
}

// MarginURL returns the URL for margin
func MarginURL() string {
	return "https://api.robinhood.com/margin/calls/"
}

// MarginInterestURL returns the URL for margin interest
func MarginInterestURL() string {
	return "https://api.robinhood.com/cash_journal/margin_interest_charges/"
}

// NotificationsURL returns the URL for notifications
func NotificationsURL(tracker bool) string {
	if tracker {
		return "https://api.robinhood.com/midlands/notifications/notification_tracker/"
	}
	return "https://api.robinhood.com/notifications/devices/"
}

// ReferralURL returns the URL for referral
func ReferralURL() string {
	return "https://api.robinhood.com/midlands/referral/"
}

// StockLoanURL returns the URL for stock loan
func StockLoanURL() string {
	return "https://api.robinhood.com/stock_loan/payments/"
}

// SubscriptionURL returns the URL for subscription
func SubscriptionURL() string {
	return "https://api.robinhood.com/subscription/subscription_fees/"
}

// WireTransfersURL returns the URL for wire transfers
func WireTransfersURL() string {
	return "https://api.robinhood.com/wire/transfers"
}

// WatchlistsURL returns the URL for watchlists
func WatchlistsURL(name string, add bool) string {
	if name != "" {
		return "https://api.robinhood.com/midlands/lists/items/"
	}
	return "https://api.robinhood.com/midlands/lists/default/"
}

// BankTransfersURL returns the URL for bank transfers
func BankTransfersURL(direction string) string {

	if direction == "received" {
		return "https://api.robinhood.com/ach/received/transfers/"
	} else {
		return "https://api.robinhood.com/ach/transfers/"
	}

}

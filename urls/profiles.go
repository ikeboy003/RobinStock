package urls

// AccountProfileURL returns the URL for account profile
func AccountProfileURL(accountNumber string) string {
	if accountNumber != "" {
		return "https://api.robinhood.com/accounts/" + accountNumber
	}
	return "https://api.robinhood.com/accounts/"
}

// BasicProfileURL returns the URL for basic profile
func BasicProfileURL() string {
	return "https://api.robinhood.com/user/basic_info/"
}

//	InvestmentProfileURL returns the URL for investment profile
func InvestmentProfileURL() string {
	return "https://api.robinhood.com/user/investment_profile/"
}

// PortfolioProfileURL returns the URL for portfolio profile	
func PortfolioProfileURL() string {
	return "https://api.robinhood.com/portfolios/"
}

// SecurityProfileURL returns the URL for security profile
func SecurityProfileURL() string {
	return "https://api.robinhood.com/user/additional_info/"
}

// UserProfileURL returns the URL for user profile
func UserProfileURL() string {
	return "https://api.robinhood.com/user/"
}

// PortfolisHistoricalsURL returns the URL for portfolio historicals
func PortfolisHistoricalsURL(accountNumber string) string {
	return "https://api.robinhood.com/portfolios/historicals/" + accountNumber + "/"
}
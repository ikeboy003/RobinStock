package urls

// LoginURL returns the URL for login
func LoginURL() string {
	return "https://api.robinhood.com/oauth2/token/"
}

// ChallengeURL returns the URL for challenge
func ChallengeURL(challengeID string) string {
	return "https://api.robinhood.com/challenge/" + challengeID + "/respond/"
}

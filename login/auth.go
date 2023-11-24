package login

import (
	"RobinStock/helpers"
	"RobinStock/urls"
	"errors"
	"fmt"
)

type LoginOptions struct {
	Username        string
	Password        string
	ExpiresIn       int
	Scope           string
	BySMS           bool
	StoreSession    bool
	LoadSession     bool
	MfaCode         string
	LoadFileName    string
	SessionFileName string
}

// RespondToChallenge posts to the challenge URL with the given challenge ID and SMS code.
func RespondToChallenge(challengeID, smsCode string) (interface{}, error) {
	url := urls.ChallengeURL(challengeID)
	payload := map[string]interface{}{
		"response": smsCode,
	}

	// Using your custom RequestPost function
	return helpers.RequestPost(url, payload, 16, false, true)
}

// Login logs in to Robinhood with the given username and password.
func Login2(opts LoginOptions) (map[string]interface{}, error) {
	// If username or password is not provided, prompt the user for them

	deviceToken := GenerateDeviceToken() // Assuming you have a similar function in Go

	// Prepare payload for login
	payload := map[string]interface{}{
		"client_id":      "c82SH0WZOsabOXGP2sxqcj34FxkvfnWRZBKlBjFS",
		"expires_in":     opts.ExpiresIn,
		"grant_type":     "password",
		"password":       opts.Password,
		"scope":          opts.Scope,
		"username":       opts.Username,
		"challenge_type": "sms", // Assuming bySMS is true for simplicity
		"device_token":   deviceToken,
	}

	if opts.MfaCode != "" {
		payload["mfa_code"] = opts.MfaCode
	}

	// Attempt to load from stored session if loadSession is true
	if opts.LoadSession {
		tokenData, err := loadSessionData(opts.LoadFileName)
		if err == nil {
			accessToken, ok := tokenData["access_token"].(string)
			if !ok {
				return nil, errors.New("invalid access token in session data")
			} else {
				helpers.UpdateSession("Authorization", fmt.Sprintf("Bearer %s", accessToken))
				helpers.SetLoginState(true)

				// Verify the token is still valid
				url := urls.PositionsURL("")                            // Assuming PositionsURL returns the correct URL
				_, err := helpers.GetRequest(url, "results", nil, true) // Adjust params as necessary
				if err != nil {
					return nil, fmt.Errorf("error verifying session: %w", err)
				}
				return tokenData, nil
			}
		} else {
			return nil, fmt.Errorf("error loading session data: %w", err)
		}
	}

	// Make the login request
	url := urls.LoginURL()
	response, err := helpers.RequestPost(url, payload, 16, false, true)
	if err != nil {
		return nil, err
	}
	// Asserting response to be a map
	responseData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format")
	}
	// Handle MFA if required
	if mfaRequired, ok := responseData["mfa_required"].(bool); ok && mfaRequired {
		// Prompt for MFA code
		mfaToken := promptForInput("Enter MFA code:")
		payload["mfa_code"] = mfaToken

		// Retry the login request with MFA code
		response, err = helpers.RequestPost(url, payload, 16, false, true)
		if err != nil {
			return nil, fmt.Errorf("error during MFA authentication: %w", err)
		}

		responseData, ok = response.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid response format after MFA")
		}
	}

	// Save session data if storeSession is true
	if opts.StoreSession {
		err := saveSessionData(opts.SessionFileName, responseData)
		if err != nil {
			return nil, err
		}
	}

	return responseData, nil
}

func Login(opts LoginOptions) (interface{}, error) {
	opts = promptCredentials(opts)
	deviceToken := GenerateDeviceToken()

	payload := preparePayload(opts, deviceToken)

	if opts.LoadSession {
		tokenData, err := handleSessionLoad(opts)
		if err != nil {
			return nil, err
		}
		return tokenData, nil
	}

	responseData, err := handleLoginRequest(opts, payload)
	if err != nil {
		return nil, err
	}

	if opts.StoreSession {
		err := saveSessionData(opts.SessionFileName, responseData)
		if err != nil {
			return nil, err
		}
	}

	return responseData, nil
}

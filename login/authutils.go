package login

import (
	"RobinStock/helpers"
	"RobinStock/urls"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func saveSessionData(filename string, data interface{}) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dataDir := filepath.Join(homeDir, ".tokens")
	if err := os.MkdirAll(dataDir, 0700); err != nil {
		return err
	}

	filePath := filepath.Join(dataDir, filename+".json")
	fileData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, fileData, 0600)
}

// loadSessionData loads session data from a file.
func loadSessionData(filename string) (map[string]interface{}, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	filePath := filepath.Join(homeDir, ".tokens", filename+".json")
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(fileData, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// GenerateDeviceToken generates a random device token.
func GenerateDeviceToken() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	var rands [16]int
	for i := range rands {
		rands[i] = r.Intn(256)
	}

	var id string
	for i, rand := range rands {
		id += fmt.Sprintf("%02x", rand)
		if i == 3 || i == 5 || i == 7 || i == 9 {
			id += "-"
		}
	}
	return id
}

func promptCredentials(opts LoginOptions) LoginOptions {
	if opts.Username == "" {
		opts.Username = promptForInput("Enter username:")
	}
	if opts.Password == "" {
		opts.Password = promptForInput("Enter password:")
	}
	return opts
}

// promptForInput displays a prompt to the user and returns their input.
func promptForInput(promptMessage string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(promptMessage)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func preparePayload(opts LoginOptions, deviceToken string) map[string]interface{} {
	payload := map[string]interface{}{
		"client_id":      "c82SH0WZOsabOXGP2sxqcj34FxkvfnWRZBKlBjFS",
		"expires_in":     opts.ExpiresIn,
		"grant_type":     "password",
		"password":       opts.Password,
		"scope":          opts.Scope,
		"username":       opts.Username,
		"challenge_type": "sms",
		"device_token":   deviceToken,
	}

	if opts.MfaCode != "" {
		payload["mfa_code"] = opts.MfaCode
	}

	return payload
}

func handleSessionLoad(opts LoginOptions) (map[string]interface{}, error) {
	tokenData, err := loadSessionData(opts.LoadFileName)
	if err != nil {
		return nil, fmt.Errorf("error loading session data: %w", err)
	}

	accessToken, ok := tokenData["access_token"].(string)
	if !ok {
		return nil, errors.New("invalid access token in session data")
	}

	helpers.UpdateSession("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	helpers.SetLoginState(true)

	url := urls.PositionsURL("")
	_, err = helpers.GetRequest(url, "results", nil, true)
	if err != nil {
		return nil, fmt.Errorf("error verifying session: %w", err)
	}

	return tokenData, nil
}

func handleLoginRequest(opts LoginOptions, payload map[string]interface{}) (interface{}, error) {
	url := urls.LoginURL()
	response, err := helpers.RequestPost(url, payload, 16, false, true)
	if err != nil {
		return nil, err
	}

	responseData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format")
	}

	if mfaRequired, ok := responseData["mfa_required"].(bool); ok && mfaRequired {
		responseData, err = handleMfa(opts, payload)
		if err != nil {
			return nil, err
		}
	} else if challenge, ok := responseData["challenge"].(map[string]interface{}); ok {
		_, err = handleChallenge(opts, challenge)
		if err != nil {
			return nil, err
		}
		fmt.Println("Successfully Passed Challenge ")
		// Reattempt login after successfully passing the challenge
		newResponse, err := helpers.RequestPost(url, payload, 16, false, true)
		if err != nil {
			return nil, err
		}
		// Set newResponse to responseData
		if newResponseData, ok := newResponse.(map[string]interface{}); ok {
			responseData = newResponseData
		} else {
			return nil, fmt.Errorf("invalid response format after reattempt")
		}
	}

	if accessToken, ok := responseData["access_token"].(string); ok {
		tokenType, _ := responseData["token_type"].(string)
		token := fmt.Sprintf("%s %s", tokenType, accessToken)

		// Update the session with the new access token
		helpers.UpdateSession("Authorization", token)
		helpers.SetLoginState(true)

		// Optionally store the session data
		if opts.StoreSession {
			if err := saveSessionData(opts.SessionFileName, responseData); err != nil {
				return nil, fmt.Errorf("error saving session data: %w", err)
			}
		}

		// Adding a detail message to the response data
		responseData["detail"] = "logged in with brand new authentication code."
	}

	return responseData, nil
}

func handleMfa(opts LoginOptions, payload map[string]interface{}) (map[string]interface{}, error) {
	mfaToken := promptForInput("Enter MFA code:")
	payload["mfa_code"] = mfaToken

	url := urls.LoginURL()
	response, err := helpers.RequestPost(url, payload, 16, false, true)
	if err != nil {
		return nil, fmt.Errorf("error during MFA authentication: %w", err)
	}

	responseData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format after MFA")
	}

	return responseData, nil
}

func handleChallenge1(opts LoginOptions, challengeData map[string]interface{}) (map[string]interface{}, error) {
	// Extract the challenge ID from the challenge data
	challengeID, ok := challengeData["id"].(string)
	if !ok {
		return nil, fmt.Errorf("challenge ID not found in challenge data")
	}

	// Prompt the user for the SMS code
	smsCode := promptForInput("Enter Robinhood code for validation: ")

	// Respond to the challenge
	response, err := RespondToChallenge(challengeID, smsCode)
	if err != nil {
		return nil, fmt.Errorf("error responding to challenge: %w", err)
	}

	// Asserting the response to be a map
	responseData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format after responding to challenge")
	}

	return responseData, nil
}

func handleChallenge(opts LoginOptions, challengeData map[string]interface{}) (interface{}, error) {
	// Extract the challenge ID from the challenge data
	challengeID, ok := challengeData["id"].(string)
	if !ok {
		return nil, fmt.Errorf("challenge ID not found in challenge data")
	}

	var smsCode string
	var response interface{}
	var err error

	for {
		// Prompt the user for the SMS code
		smsCode = promptForInput("Enter Robinhood code for validation: ")

		// Respond to the challenge
		response, err = RespondToChallenge(challengeID, smsCode)
		if err != nil {
			return nil, fmt.Errorf("error responding to challenge: %w", err)
		}

		// Asserting the response to be a map
		responseData, ok := response.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid response format after responding to challenge")
		}

		// Check the status of the challenge response
		if challengeResponse, ok := responseData["challenge"].(map[string]interface{}); ok {
			if status, ok := challengeResponse["status"].(string); ok && status == "validated" {
				// Update the session with the challenge ID upon successful validation
				helpers.UpdateSession("X-ROBINHOOD-CHALLENGE-RESPONSE-ID", challengeID)
				return responseData, nil
			} else if remainingAttempts, ok := challengeResponse["remaining_attempts"].(float64); ok && remainingAttempts > 0 {
				fmt.Printf("That code was not correct. %d tries remaining. Please type in another code.\n", int(remainingAttempts))
				continue
			} else if remainingAttempts, ok := challengeResponse["remaining_attempts"].(int); ok && remainingAttempts == 0 {
				fmt.Println("You have run out of tries. Please try again later.")
				return nil, errors.New("challenge failed")
			}
		}

		return responseData, nil
	}
}

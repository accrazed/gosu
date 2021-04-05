package gosu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const tokenURL = "https://osu.ppy.sh/oauth/token"
const authCodeURL = "https://osu.ppy.sh/oauth/authorize"
const baseURL = "https://osu.ppy.sh/api/v2/"
const tokenName = "oauth_token.json"

type GosuClient struct {
	clientSecret string
	clientID     string
	token        string
}

type oauthTokenResponse struct {
	AccessToken    string
	ExpiresIn      float64
	LastAccessTime time.Time
	TokenType      string
}

func RevalidateOAuthToken(clientSecret string, clientID string) (*oauthTokenResponse, error) {
	var ret *oauthTokenResponse
	var err error

	// if oauth token doesn't exist in file dir
	if _, err = os.Stat(tokenName); os.IsNotExist(err) {
		fmt.Println("Oauth token not detected. requesting new token...")

		// Request a new token
		if ret, err = requestOAuthToken(clientSecret, clientID); err != nil {
			fmt.Println("Error requesting Oauth token: %w", err)
		}
		ret.LastAccessTime = time.Now()
		writeOAuthToken(tokenName, ret)

		return ret, nil
	} else {
		fmt.Println("Oauth token found. Validating and updating timestamp...")
		var previousToken *oauthTokenResponse
		// In this case, the oauth token file already exists
		if previousToken, err = readOAuthToken(tokenName); err != nil {
			fmt.Println("Error in reading OAuth token: %w", err)
			return nil, err
		}
		// If the token is expired, get a new one
		if float64(time.Since(previousToken.LastAccessTime).Seconds()) > previousToken.ExpiresIn {
			fmt.Println("Token expired. Requesting new token...")
			if ret, err = requestOAuthToken(clientSecret, clientID); err != nil {
				fmt.Println("Error requesting Oauth token: %w", err)
			}

			ret.LastAccessTime = time.Now()
			writeOAuthToken(tokenName, ret)
			return ret, nil
		} else {
			return previousToken, nil
		}
	}
}

// contacts the osu api to request a new OAuth token
func requestOAuthToken(clientSecret string, clientID string) (*oauthTokenResponse, error) {
	var ret oauthTokenResponse

	// Make token request
	resp, err := http.PostForm(tokenURL, url.Values{
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"grant_type":    {"client_credentials"},
		"scope":         {"public"},
	})
	if err != nil {
		fmt.Println("Error in getOauthToken POST: %w", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in reading POST response: %w", err)
		return nil, err
	}

	// Change formatting so that the variable structure is easily unmarshalable
	oauthFormatter := strings.NewReplacer(
		"token_type", "TokenType",
		"expires_in", "ExpiresIn",
		"access_token", "AccessToken")
	formattedBody := oauthFormatter.Replace(string(body))

	// Unmarshal the json into the oauthTokenResponse struct
	if err := json.Unmarshal([]byte(formattedBody), &ret); err != nil {
		fmt.Println("Error in unmarshaling body: %w", err)
		return nil, err
	}

	// Timestamp the token
	ret.LastAccessTime = time.Now()
	return &ret, nil
}

func readOAuthToken(path string) (*oauthTokenResponse, error) {
	var ret *oauthTokenResponse

	// If the file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("file %s not found", path)
	} else {
		// Read the file
		if body, err := os.ReadFile(path); err != nil {
			fmt.Println("Error reading file: %w", err)
			return nil, err
		} else {
			// Decode body into ret
			if err := json.Unmarshal(body, &ret); err != nil {
				fmt.Println("Error in unmarshaling body: %w", err)
				return nil, err
			}
			return ret, nil
		}
	}
}

func writeOAuthToken(path string, token *oauthTokenResponse) error {
	if data, err := json.Marshal(token); err != nil {
		fmt.Println("Error converting struct oauthTokenResponse to json: %w", err)
		return err
	} else if f, err := os.Create(path); err != nil {
		fmt.Println("Error creating file: %w", err)
		return err
	} else {
		defer f.Close()
		f.WriteString(string(data))
	}
	return nil
}

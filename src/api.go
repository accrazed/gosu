package gosu

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const tokenURL = "https://osu.ppy.sh/oauth/token"
const authCodeURL = "https://osu.ppy.sh/oauth/authorize"
const baseURL = "https://osu.ppy.sh/api/v2/"

type GosuClient struct {
	clientSecret string
	clientID     string
	token        string
}

func GetOauthToken(clientSecret string, clientID string) *GosuClient {
	resp, err := http.PostForm(tokenURL, url.Values{"client_id": {clientID}, "client_secret": {clientSecret}, "grant_type": {"client_credentials"}, "scope": {"public"}})

	if err != nil {
		fmt.Println("Error in getOauthToken POST: %w", err)
		return nil
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error in reading POST response: %w", err)
		return nil
	}

	ioutil.WriteFile("dump", body, 0600)

	respStr := string(body)

	fmt.Println(respStr) // json

	token := respStr

	return &GosuClient{clientSecret: clientSecret, clientID: clientID, token: token}
}

package gosu

const tokenURL = "https://osu.ppy.sh/oauth/token"
const authCodeURL = "https://osu.ppy.sh/oauth/authorize"
const baseURL = "https://osu.ppy.sh/api/v2/"

type GosuClient struct {
	token string
}

func createGosuClient(token string) GosuClient {
	return GosuClient{token: token}
}

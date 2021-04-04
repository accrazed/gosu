package main

import (
	"fmt"
	"os"

	gosu "accrazed.com/gosu/src"
	"github.com/joho/godotenv"
)

func getEnvVar(key string) string {
	err := godotenv.Load("src/test/.env")

	if err != nil {
		fmt.Println("Error loading env file: %w", err)
		return ""
	}
	return os.Getenv(key)
}

func main() {
	godotenv.Load()
	test := gosu.GetOauthToken(getEnvVar("GOSU_CLIENT_SECRET"), getEnvVar("GOSU_CLIENT_ID"))

	fmt.Println(test)
	fmt.Printf("%d %d %d %d\n%s", gosu.Graveyard, gosu.Loved, gosu.Mania, gosu.Taiko, getEnvVar("GOSU_CLIENT_SECRET"))
}

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
	test, err := gosu.RevalidateOAuthToken(getEnvVar("GOSU_CLIENT_SECRET"), getEnvVar("GOSU_CLIENT_ID"))

	if err != nil {
		panic(err)
	}
	fmt.Println(*test)
}

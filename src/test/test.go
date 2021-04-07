package main

import (
	"fmt"
	"os"

	gosu "accrazed.com/gosu/src"
	"github.com/joho/godotenv"
)

func getEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading env file: %w", err)
		return ""
	}
	return os.Getenv(key)
}

func main() {
	godotenv.Load()
	client, _ := gosu.CreateGosuClient(getEnvVar("GOSU_CLIENT_SECRET"), getEnvVar("GOSU_CLIENT_ID"))

	beatmap, _ := client.LookupBeatmap("2775178", "id")

	fmt.Println(beatmap)

}

package main

import (
	"fmt"
	"os"

	gosu "github.com/AccraZed/gosu"
	"github.com/joho/godotenv"
)

func getEnvVar(key string) string {
	err := godotenv.Load("src/test/.env")

	if err != nil {
		fmt.Println("Error loading env file:", err)
		return ""
	}
	return os.Getenv(key)
}

func main() {
	godotenv.Load()
	client, _ := gosu.CreateGosuClient(getEnvVar("GOSU_CLIENT_SECRET"), getEnvVar("GOSU_CLIENT_ID"))

	// beatmap, _ := client.LookupBeatmap("2775178", "id")
	// score, _ := client.GetUserBeatmapScore("1781697", "6725659", map[string]interface{}{})
	// scores, _ := client.GetBeatmapScores("1781697", map[string]interface{}{})
	beatmap, _ := client.GetBeatmap("1781697")

	// fmt.Println(beatmap)
	fmt.Println(beatmap)
}

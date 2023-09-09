package main

import (
	"os"
	"strings"

	"github.com/codecrafters-io/grep-tester/internal"
	"github.com/joho/godotenv"
)

func main() {
	os.Exit(internal.RunCLI(envMap()))
}

func envMap() map[string]string {
	_ = godotenv.Load(".env")

	result := make(map[string]string)
	for _, keyVal := range os.Environ() {
		split := strings.SplitN(keyVal, "=", 2)
		key, val := split[0], split[1]
		result[key] = val
	}

	return result
}

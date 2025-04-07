package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	fmt.Println("Hi, your OPENAI_API_KEY is:", os.Getenv("OPENAI_API_KEY"))
}

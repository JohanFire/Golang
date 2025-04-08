/*
https://github.com/openai/openai-go
go get -u 'github.com/openai/openai-go@v0.1.0-beta.7'
*/

package main

import (
	"context"
	"fmt"

	// "os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go"
	// "github.com/openai/openai-go/option"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	client := openai.NewClient(
	// option.WithAPIKey("os.Getenv("OPENAI_API_KEY")"), // defaults to os.LookupEnv("OPENAI_API_KEY")
	)

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("Hi, I'm Johan. Tell me what model are you"),
		},
		Model: openai.ChatModelGPT3_5Turbo,
	})
	if err != nil {
		panic(err.Error())
	}
	println(chatCompletion.Choices[0].Message.Content)
}

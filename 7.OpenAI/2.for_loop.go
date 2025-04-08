/*
https://github.com/openai/openai-go
go get -u 'github.com/openai/openai-go@v0.1.0-beta.7'
*/

package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

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

	fmt.Println()
	for true {
		fmt.Print("Send a message: ")

		inputScanner := bufio.NewScanner(os.Stdin)
		inputScanner.Scan()

		// var userInput string = inputScanner.Text()

		chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(
					// userInput,
					inputScanner.Text(),
				),
			},
			Model: openai.ChatModelGPT3_5Turbo,
		})
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("AI:", chatCompletion.Choices[0].Message.Content)
		fmt.Println()
	}

}

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/azure"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	var azureOpenAIAPIKey string = os.Getenv("AZURE_OPENAI_API_KEY")
	var azureOpenAIEndpoint string = os.Getenv("AZURE_OPENAI_ENDPOINT")
	var azureOpenAIAPIVersion string = os.Getenv("AZURE_OPENAI_API_VERSION")
	var azureChatDeploymentModel string = os.Getenv("AZURE_CHAT_DEPLOYMENT")

	client := openai.NewClient(
		azure.WithEndpoint(azureOpenAIEndpoint, azureOpenAIAPIVersion),
		azure.WithAPIKey(azureOpenAIAPIKey),
	)

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("Hi, I'm Johan. Tell me what model are you"),
		},
		// Model: openai.ChatModelGPT3_5Turbo,
		Model: azureChatDeploymentModel,
	})
	if err != nil {
		panic(err.Error())
	}
	println(chatCompletion.Choices[0].Message.Content)
}

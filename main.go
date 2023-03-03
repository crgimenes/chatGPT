package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Payload struct {
	Model    string     `json:"model"`
	Messages []Messages `json:"messages"`
}

func main() {

	data := Payload{
		Model: "gpt-3.5-turbo",
		Messages: []Messages{
			{
				Role:    "user",
				Content: "What is the OpenAI mission?",
			},
		},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	req.Header.Set("Authorization", os.ExpandEnv("Bearer $OPENAI_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(b))

}

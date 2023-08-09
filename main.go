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
		Model: "gpt-4",
		Messages: []Messages{
			{
				Role: "user",
				Content: `estamos criando bitmaps com imagens em pixelart para um jogo incluindo letras, personagens, itens, etc. todos os desenhos são feitos em uma tabela em que o 0 representa o pixel preto e 1 representa pixel o branco. Todas as imagens da grade são 8x8 pixels.

Como no exemplo abaixo em que desenhei a letra A:

00000000
01111110
01000010
01000010
01111110
01000010
01000010
00000000

seguindo esse mesmo padrão desenhe a letra F`,
			},
		},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest(
		http.MethodPost,
		"https://api.openai.com/v1/chat/completions", body)
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

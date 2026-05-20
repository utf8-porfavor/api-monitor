package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type API struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"descricao"`
}

func checaAPI(client *http.Client, api API) {
	resp, err := client.Get(api.URL)

	if err != nil {
		fmt.Printf("[ERRO] %s -> %v\n", api.Name, err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Printf("[OK] %s -> status %d\n", api.Name, resp.StatusCode)
	} else {
		fmt.Printf("[FALHA] %s -> status %d\n", api.Name, resp.StatusCode)
	}
}

func main() {
	data, err := os.ReadFile("apis.json")
	if err != nil {
		fmt.Println("Erro ao ler arquivo:", err)
		return
	}

	var apis []API

	if err := json.Unmarshal(data, &apis); err != nil {
		fmt.Println("Erro ao converter JSON:", err)
		return
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for _, api := range apis {
		checaAPI(client, api)
	}
}

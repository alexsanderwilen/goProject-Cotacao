package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatal("Erro criando requisição:", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Erro chamando servidor:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Erro lendo resposta:", err)
	}

	var data map[string]string
	fmt.Println("Resposta do servidor:", string(body))
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal("Erro interpretando JSON:", err)
	}

	bid := data["bid"]

	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Fatal("Erro criando arquivo:", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Dólar: %s\n", bid))
	if err != nil {
		log.Fatal("Erro escrevendo no arquivo:", err)
	}

	fmt.Println("Cotação salva com sucesso:", bid)
}

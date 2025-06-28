package main

import (
	"Cotacao/config"
	"Cotacao/handler"
	"Cotacao/repository"
	"Cotacao/service"
	"log"
	"net/http"
)

func main() {
	db := config.InitDB()
	repo := repository.NewCotacaoRepository(db)
	svc := service.NewCotacaoService(repo)
	h := handler.NewCotacaoHandler(svc)

	http.HandleFunc("/cotacao", h.HandleCotacao)

	log.Println("Servidor na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

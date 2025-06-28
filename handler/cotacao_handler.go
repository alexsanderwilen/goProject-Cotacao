package handler

import (
	"Cotacao/service"
	"encoding/json"
	"log"
	"net/http"
)

type CotacaoHandler struct {
	service service.CotacaoService
}

func NewCotacaoHandler(s service.CotacaoService) *CotacaoHandler {
	return &CotacaoHandler{service: s}
}

func (h *CotacaoHandler) HandleCotacao(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cotacao, err := h.service.BuscarECadastrarCotacao(ctx)
	if err != nil {
		log.Println("Erro ao processar cotação:", err)
		http.Error(w, "Erro ao buscar cotação", http.StatusRequestTimeout)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"bid": cotacao.Bid,
	})
}

package external

import (
	"Cotacao/model"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type ResponseAPI struct {
	USDBRL model.Cotacao `json:"USDBRL"`
}

func BuscarCotacao(ctx context.Context) (model.Cotacao, error) {
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return model.Cotacao{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return model.Cotacao{}, err
	}
	defer resp.Body.Close()

	var r ResponseAPI
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return model.Cotacao{}, err
	}

	return r.USDBRL, nil
}

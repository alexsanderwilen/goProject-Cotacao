package service

import (
	"Cotacao/external"
	"Cotacao/model"
	"Cotacao/repository"
	"context"
)

type CotacaoService interface {
	BuscarECadastrarCotacao(ctx context.Context) (model.Cotacao, error)
}

type cotacaoService struct {
	repo repository.CotacaoRepository
}

func NewCotacaoService(repo repository.CotacaoRepository) CotacaoService {
	return &cotacaoService{repo: repo}
}

func (s *cotacaoService) BuscarECadastrarCotacao(ctx context.Context) (model.Cotacao, error) {
	cotacao, err := external.BuscarCotacao(ctx)
	if err != nil {
		return model.Cotacao{}, err
	}

	err = s.repo.SalvarCotacao(ctx, cotacao)
	if err != nil {
		// Continua mesmo se der erro na gravação
	}

	return cotacao, nil
}

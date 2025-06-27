package repository

import (
	"Cotacao/model"
	"context"
	"database/sql"
	"log"
	"time"
)

type CotacaoRepository interface {
	SalvarCotacao(ctx context.Context, c model.Cotacao) error
}

type cotacaoRepository struct {
	db *sql.DB
}

func NewCotacaoRepository(db *sql.DB) CotacaoRepository {
	return &cotacaoRepository{db: db}
}

func (r *cotacaoRepository) SalvarCotacao(ctx context.Context, c model.Cotacao) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	query := "INSERT INTO cotacoes (bid) VALUES (?)"
	_, err := r.db.ExecContext(ctx, query, c.Bid)
	if err != nil {
		log.Println("Erro ao salvar no banco:", err)
	}
	return err
}

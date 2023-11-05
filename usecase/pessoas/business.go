package pessoas

import (
	"context"
	"github.com/google/uuid"
	"spark/database"
	domain "spark/domain/pessoas"
	"spark/infraestructure/repository/pessoas"
)

// Adicionar define a regra de negocio para adicionar uma pessoa à base de dados
func Adicionar(ctx context.Context, req *domain.Pessoa) (id *uuid.UUID, err error) {
	tx, err := database.NewTransaction(ctx, false)
	if err != nil {
		return id, err
	}
	defer tx.Rollback()

	repo := pessoas.NewRepo(tx)

	id, err = repo.AdicionarPessoa(req)
	if err != nil {
		return id, err
	}

	if req.Stack != nil {
		for i := range req.Stack {
			if err = repo.AdicionarTech(id, &req.Stack[i]); err != nil {
				return id, err
			}
		}
	}

	if err = tx.Commit(); err != nil {
		return id, err
	}

	return
}

package pessoas

import (
	"context"
	"spark/database"
	domain "spark/domain/pessoas"
	"spark/infraestructure/repository/pessoas"
)

// Adicionar define a regra de negocio para adicionar uma pessoa à base de dados
func Adicionar(ctx context.Context, req *domain.Pessoa) (err error) {
	tx, err := database.NewTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	repo := pessoas.NewRepo(tx)

	id, err := repo.Adicionar(req)
	if err != nil {
		return err
	}

	if req.Stack != nil {
		for i := range req.Stack {
			if err = repo.AdicionarTech(id, &req.Stack[i]); err != nil {
				return err
			}
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return
}

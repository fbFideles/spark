package pessoas

import (
	"github.com/google/uuid"
	"spark/database"
	domain "spark/domain/pessoas"
)

func NewRepo(tx database.Transaction) domain.RepoPessoas {
	return &repository{tx: tx}
}

type repository struct {
	tx database.Transaction
}

// AdicionarPessoa define um metodo para adicionar um registro de pessoa
func (r *repository) AdicionarPessoa(req *domain.Pessoa) (id *uuid.UUID, err error) {
	id = new(uuid.UUID)
	if err := r.tx.QueryRow(`INSERT INTO public.t_pessoa (apelido, nome) 
			VALUES ($1, $2) RETURNING "id";`, req.Apelido, req.Nome).Scan(&id); err != nil {
		return id, err
	}
	return
}

// AdicionarTech define um metodo para adicionar a tecnologia de uma pessoa
func (r *repository) AdicionarTech(pessoaID *uuid.UUID, tech *string) error {
	if _, err := r.tx.Exec(`INSERT INTO public.t_tecnologia(pessoa_id, nome) VALUES ($1, $2);`, pessoaID, tech); err != nil {
		return err
	}
	return nil
}

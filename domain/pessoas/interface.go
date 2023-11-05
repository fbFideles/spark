package pessoas

import "github.com/google/uuid"

// RepoPessoas define uma interface
type RepoPessoas interface {
	Adicionar(req *Pessoa) (id *uuid.UUID, err error)
	AdicionarTech(pessoaID *uuid.UUID, tech *string) (err error)
}

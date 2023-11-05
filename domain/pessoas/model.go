package pessoas

import "time"

// Pessoa define uma entidade de pessoa
type Pessoa struct {
	Apelido    *string    `json:"apelido"`
	Nome       *string    `json:"nome"`
	Nascimento *time.Time `json:"nascimento"`
	Stack      []string   `json:"stack"`
}

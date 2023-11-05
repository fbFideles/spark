package pessoas

import (
	"bou.ke/monkey"
	"context"
	"github.com/google/uuid"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"spark/database"
	domain "spark/domain/pessoas"
	"spark/mocks"
	"spark/utils"
	"time"

	repository "spark/infraestructure/repository/pessoas"
	"testing"
)

func TestPessoasUseCase(t *testing.T) {
	transactionMock := mocks.NewTransaction(t)
	transactionMock.On("Rollback")
	transactionMock.On("Commit").
		Return(nil)

	connMock := mocks.NewConnectionsPool(t)
	connMock.On("NewTransaction", mock.Anything, mock.Anything).
		Return(transactionMock, nil)

	ctx := context.WithValue(context.Background(), utils.DatabaseKey, connMock)

	t.Run("Deve registrar com sucesso uma pessoa", func(t *testing.T) {
		uuidMock := uuid.New()

		reqTest := &domain.Pessoa{
			Apelido:    utils.GetPointer("Fideles"),
			Nome:       utils.GetPointer("Felipe"),
			Nascimento: utils.GetPointer(time.Now()),
			Stack: []string{
				"Postgres",
				"Golang",
			},
		}

		repoMock := mocks.NewRepoPessoas(t)
		repoMock.On("AdicionarPessoa", reqTest).Return(&uuidMock, nil)
		repoMock.On("AdicionarTech", &uuidMock, &reqTest.Stack[0]).Return(nil)
		repoMock.On("AdicionarTech", &uuidMock, &reqTest.Stack[1]).Return(nil)

		monkey.Patch(repository.NewRepo, func(tx database.Transaction) domain.RepoPessoas {
			return repoMock
		})
		defer monkey.Unpatch(repository.NewRepo)

		id, err := Adicionar(ctx, reqTest)
		require.Nil(t, err)

		assert.Equal(t, &uuidMock, id)
	})
}

package pessoas

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"spark/config"
	"spark/database"
	domain "spark/domain/pessoas"
	"spark/handler/pessoas"
	"spark/middleware"
	"spark/utils"
	"testing"
	"time"
)

func TestIntegracaoPessoas(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test short")
	}

	config.LoadConfig()
	db, err := database.OpenConnection()
	require.Nil(t, err)
	ginEngine := gin.New()
	v1 := ginEngine.Group("/v1/")
	v1.Use(middleware.SetupConnections(db))
	pessoas.Router(v1)

	t.Run("Integration Testing Adicionar Pessoa", func(t *testing.T) {
		w := httptest.NewRecorder()

		reqTest := domain.Pessoa{
			Apelido:    utils.GetPointer("Fideles"),
			Nome:       utils.GetPointer("Felipe"),
			Nascimento: utils.GetPointer(time.Now()),
			Stack: []string{
				"Postgres",
				"Golang",
			},
		}

		dados, err := json.Marshal(reqTest)
		require.Nil(t, err)

		req, _ := http.NewRequest("POST", "/v1/pessoas", bytes.NewBuffer(dados))
		ginEngine.ServeHTTP(w, req)
		assert.Equal(t, w.Code, http.StatusCreated)
	})
}

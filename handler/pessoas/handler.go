package pessoas

import (
	"github.com/gin-gonic/gin"
	"net/http"
	domain "spark/domain/pessoas"
	"spark/usecase/pessoas"
)

func listarPessoas(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "from ListarPessoas",
	})
}

func obterPessoa(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "from ObterPessoa",
	})
}

func obterContagemPessoas(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "from ObterContagemPessoas",
	})
}

func adicionarPessoa(c *gin.Context) {
	var req domain.Pessoa

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	if err = pessoas.Adicionar(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusCreated, nil)
}

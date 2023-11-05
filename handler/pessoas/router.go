package pessoas

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("pessoas", adicionarPessoa)
	r.GET("pessoas/:id", obterPessoa)
	r.GET("pessoas", listarPessoas)
	r.GET("contagem-pessoas", obterContagemPessoas)
}

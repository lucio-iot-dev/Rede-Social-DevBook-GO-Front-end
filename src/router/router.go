package router

import (
	"webapp/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar retorna um router com todas as rotas configuradas
func Gerar() *mux.Router {

	// Primeira forma que trazer a rota.congigurar para fazer o mapeamento das rotas

	r := mux.NewRouter()
	return rotas.Configurar(r)
	/*-------------------------------------------------------------------------*/
	// Segunda forma que trazer a rota.congigurar para fazer o mapeamento das rotas
	// router := mux.NewRouter()
	// rotas.Configurar(router)
	// return router
}

// return mux.NewRouter() Apresentou erro

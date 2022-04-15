package modelos

import (
	"net/http"
	"time"
)

// Usuario representa um a pessoa utilizando a rede social
type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	CriadoEm    time.Time    `json:"criadoEm"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}


// BuscarUsuarioCompleto faz 4 requisições na API para montar o usuário
func BuscarUsuarioCompleto(usuarioID uint64, r *http.Request) (Usuario, error) {
    canalUsuario := make(chan Usuario)
		canalSeguidores := make(chan []Usuario)
		canalSeguindo := make(chan []Usuario)
		canalPublicacoes := make(chan []Publicacao)

		go BuscarDadosDoUsuario(canalUsuario, usuarioID, r)
		go BuscarSeguidores(canalSeguidores, usuarioID, r)
		go BuscarSDeguindo(canalSeguindo, usuarioID,r)
		go BuscarPublicacoes(canalPublicacoes, usuarioID, r)
	
}

func BuscarDadosDoUsuario(canal <-chan Usuario, usuarioID uint64, r *http.Request) {

}


func BuscarSeguidores(canal <-chan []Usuario, usuarioID uint64, r *http.Request) {
	
}


func BuscarSDeguindo(canal <-chan []Usuario, usuarioID uint64, r *http.Request) {
	
}


func BuscarPublicacoes(canal <-chan []Publicacao, usuarioID uint64, r *http.Request) {
	
}
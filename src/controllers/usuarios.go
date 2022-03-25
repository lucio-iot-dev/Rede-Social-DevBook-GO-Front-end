package controllers

import (
	"fmt"
	"net/http"
)

//CriarUsuario chama a API para cadastrar um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	nome := r.FormValue("nome")

	fmt.Println(nome)
}

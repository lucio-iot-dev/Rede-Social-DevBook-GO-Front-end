package controllers

import (
	"fmt"
	"net/http"
)

// CarregarTelaDeLogin vai renderizar a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println([]byte("Tela de login"))
	w.Write([]byte("Tela de login"))
}
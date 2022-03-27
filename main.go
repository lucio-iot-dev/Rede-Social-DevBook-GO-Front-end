package main

import (
	// "encoding/hex"
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"

	// "github.com/gorilla/securecookie"
)
//FUNÇÃO UTILIZADA SOMENTE UMA VEZ PARA GERAR AS CHAVES
// func init() {
// 	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(hashKey)

// 	Blockkey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(Blockkey)
// }

func main() {
  config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}

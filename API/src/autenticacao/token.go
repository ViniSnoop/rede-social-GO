package autenticacao

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken retorna um token assinado com a permissoes do usuario
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioID"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte("Secret"))
}

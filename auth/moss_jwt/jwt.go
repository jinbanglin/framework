package moss_jwt

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("")

type MossClaims struct {
	UserId string
	Client string
	AppId  string
	jwt.StandardClaims
}

func NewJwtToken(claims jwt.Claims, jwtKey []byte) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
}

func GetJwtTokenString(val string) (token string, ok bool) {
	authHeaderParts := strings.Split(val, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", false
	}
	return authHeaderParts[1], true
}

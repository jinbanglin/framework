package moss_jwt

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/segmentio/ksuid"
)

var MossJwtExpiresDelta int64 = 7200
var JwtKey = []byte("8784e410796b279afea776524a6a464d7f9c153b")

type Claims struct {
	UserName string
	UserId   string
	jwt.StandardClaims
}

func genJwtClaims(userName, userId, audience string) Claims {
	now := time.Now()
	return Claims{
		UserName: userName,
		UserId:   userId,
		StandardClaims: jwt.StandardClaims{
			Audience:  audience,
			ExpiresAt: now.Unix() + MossJwtExpiresDelta,
			Id:        ksuid.New().String() + userId,
			IssuedAt:  now.Unix(),
			Issuer:    "moss",
			NotBefore: now.Unix(),
			Subject:   userId + "." + userName,
		},
	}
}

func NewJwtToken(userName, userId, audience string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, genJwtClaims(
		userName,
		userId,
		audience,
	))
	return token.SignedString(JwtKey)
}

func GetJwtTokenString(val string) (token string, ok bool) {
	authHeaderParts := strings.Split(val, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", false
	}
	return authHeaderParts[1], true
}

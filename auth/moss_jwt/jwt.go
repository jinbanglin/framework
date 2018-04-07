package moss_jwt

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const EXPIRES_DELTA = 3600 * 12

//==============================jwt sha1==============================
var JwtKey = []byte("8784e410796b279afea776524a6a464d7f9c153b")

//==============================pwd sha1==============================
var PasswordSalt = []byte("17160fcec5bd355fb8b770630d5a13401c9cd6cb")

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
			ExpiresAt: now.Unix() + EXPIRES_DELTA,
			Id:        now.Format("150405") + userId,
			IssuedAt:  now.Unix(),
			Issuer:    "bixin_user",
			NotBefore: now.Unix(),
			Subject:   userId + "_" + userName,
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

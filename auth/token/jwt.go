package token

import (
	"crypto/rsa"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTTokenGen struct {
	privateKey *rsa.PrivateKey;
	issuer string;
	nowFunc func() time.Time;
}

func NewJWTTokenGen(issuer string, privateKey *rsa.PrivateKey) *JWTTokenGen {
	return &JWTTokenGen {
		issuer:issuer,
		nowFunc: time.Now,
		privateKey: privateKey,
	}
}

func (t *JWTTokenGen) GenerateToken(accountID string, expire time.Duration)(string, error) {
	nowSec := t.nowFunc().Unix()
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512,jwt.StandardClaims{
		Issuer: t.issuer,
		IssuedAt: nowSec,
		ExpiresAt: nowSec+int64(expire),
		Subject: accountID,
	})

	return tkn.SignedString(t.privateKey)
}
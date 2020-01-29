package auth

import (
	"crypto/rsa"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	currentKey    *rsa.PrivateKey
	oldKey        *rsa.PrivateKey
	tokenLifeTime time.Duration
)

func Init(current *rsa.PrivateKey, old *rsa.PrivateKey, lifeTime time.Duration) {
	currentKey = current
	oldKey = old
	tokenLifeTime = lifeTime
}

// Please see the documentation: http://jwt.io/
func Verify(authToken string) (userId string, err error) {
	// parse and vertify the token string
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(authToken, claims, func(t *jwt.Token) (interface{}, error) {
		// make sure the JWT token is using RSA alg
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return &currentKey.PublicKey, nil
	})
	if err != nil {
		return ``, err
	}

	if token.Valid == false { // make sure token is Valid
		return ``, errors.New("Wrong jwt token")
	}
	// check LifeTime
	switch ts := claims["lifeTime"].(type) {
	default:
		return ``, errors.New("Improper JWT Token")
	case float64:
		timestamp := time.Unix(int64(ts), 0)
		if timestamp.Before(time.Now()) {
			return ``, errors.New("JWT Token has expired")
		}
	}

	if s, ok := claims["userId"].(string); !ok {
		return ``, errors.New("Improper JWT Token")
	} else {
		userId = s
	}
	return userId, nil
}

func ToekenSign(userId string) (authToken string, err error) {
	token := jwt.New(jwt.SigningMethodRS512)
	claims := make(jwt.MapClaims)
	claims["userId"] = userId
	claims["lifeTime"] = time.Now().Add(tokenLifeTime).Unix()
	token.Claims = claims
	return token.SignedString(currentKey)
}

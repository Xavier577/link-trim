package tokenizer

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenVerifyOptions struct {
	Secret      string
	SignedToken string
}

func VerifyToken(tokenVerifyOptions TokenVerifyOptions) (bool, JWTClaim, error) {

	var isValid bool
	var verificationError error
	var jwtClaim JWTClaim

	token, err := jwt.ParseWithClaims(
		tokenVerifyOptions.SignedToken,
		&JWTClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(tokenVerifyOptions.Secret), nil
		},
	)

	if err != nil {
		isValid = false
		verificationError = err
	}

	if claims, ok := token.Claims.(*JWTClaim); ok && token.Valid {
		fmt.Println(claims.Payload)
		isValid = true
		jwtClaim = *claims
	} else if !ok {
		verificationError = errors.New("TOKEN_PARSE_ERR")
	} else if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		verificationError = errors.New("TOKEN_EXPIRED")
	}

	return isValid, jwtClaim, verificationError
}

package utils

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	SubjectClaim   = "sub"
	ExpiryClaim    = "exp"
	IssuedAtClaim  = "iat"
	NotBeforeClaim = "nbf"
)

func CreateJWTToken(ttl time.Duration, payload interface{}, privateKey string) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("could not decode private key: %w", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	now := time.Now().UTC()

	claims := jwt.MapClaims{
		SubjectClaim:   payload,
		ExpiryClaim:    now.Add(ttl).Unix(),
		IssuedAtClaim:  now.Unix(),
		NotBeforeClaim: now.Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return token, nil
}

func ValidateJWTToken(token string, publicKey string) (interface{}, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode public key: %w", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims format")
	}

	return claims[SubjectClaim], nil
}

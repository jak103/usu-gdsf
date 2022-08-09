package auth

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"time"

	"github.com/jak103/usu-gdsf/config"
)

type UserType int8

const (
	ADMIN_USER   UserType = iota
	REGULAR_USER
)


type TokenType int8

const (
	ACCESS_TOKEN  TokenType = iota
	REFRESH_TOKEN
)

const (
	TOKEN_EXPIRED = "Expired token"
)

type TokenClaims struct {
	Type       TokenType
	Expiration int64    
	UserId     uint64
	UserType   UserType
	UserEmail  string
}

type TokenParams struct {
	Type      TokenType
	UserId    uint64
	UserType  UserType
	UserEmail string
}

func GenerateToken(params TokenParams) string {
	expiration := time.Now().UnixMilli()

	switch params.Type {
	case ACCESS_TOKEN:
		expiration += config.AccessTokenLifetimeMins * 60 * 1000
	case REFRESH_TOKEN:
		expiration += config.RefreshTokenLifetimeDays * 24 * 60 * 60 * 1000
	default:
		panic("Unrecognized token type")
	}

	claims := TokenClaims{
		Type:       params.Type,
		Expiration: expiration,
		UserId: params.UserId,
		UserType: params.UserType,
		UserEmail: params.UserEmail,
	}

	claimsJson, err := json.Marshal(claims)

	if err != nil {
		panic("Failed to serialize token claims as JSON")
	}

	mac := hmac.New(sha256.New, []byte(config.TokenHashingKey))
	mac.Write(claimsJson)

	hash := mac.Sum(nil)
	encodedHash := hex.EncodeToString(hash)

	unencodedToken := string(claimsJson) + "|" + encodedHash
	encodedToken := base64.RawURLEncoding.EncodeToString([]byte(unencodedToken))

	return encodedToken
}

func DecodeAndVerifyToken(token string, tokenType TokenType) (*TokenClaims, error) {
	if len(token) == 0 {
		return nil, errors.New("Token cannot be empty")
	}

	currentTimestamp := time.Now().UnixMilli()
	decodedToken, err := base64.RawURLEncoding.DecodeString(token)

	if err != nil {
		return nil, errors.New("Invalid token")
	}

	indexOfDelimeter := bytes.LastIndexByte(decodedToken, byte('|'))

	if indexOfDelimeter < 1 || len(decodedToken) <= indexOfDelimeter {
		return nil, errors.New("Invalid token format")
	}

	providedHexHash := decodedToken[indexOfDelimeter+1:]
	providedClaimsJson := decodedToken[:indexOfDelimeter]

	var claims TokenClaims

	if err = json.Unmarshal(providedClaimsJson, &claims); err != nil {
		return nil, errors.New("Invalid claims payload")
	}

	if claims.Expiration <= currentTimestamp {
		return nil, errors.New(TOKEN_EXPIRED)
	}

	if claims.Type != tokenType {
		return nil, errors.New("Incorrect token type")
	}

	providedHash, err := hex.DecodeString(string(providedHexHash))

	if err != nil {
		return nil, errors.New("Invalid signature")
	}

	mac := hmac.New(sha256.New, []byte(config.TokenHashingKey))
	mac.Write(providedClaimsJson)

	realHash := mac.Sum(nil)

	if !hmac.Equal(realHash, providedHash) {
		return nil, errors.New("Incorrect signature")
	}

	return &claims, nil
}

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

type TokenType int8

const (
	ACCESS_TOKEN TokenType = iota
	REFRESH_TOKEN
)

type TokenClaims struct {
	Type       TokenType
	Expiration int64    
	UserId     int64    
	UserEmail  string   
}

type TokenParams struct {
	Type      TokenType
	UserId    int64
	UserEmail string
}

func GenerateToken(params TokenParams) (string, error) {
	expiration := time.Now().UnixMilli()

	switch params.Type {
	case ACCESS_TOKEN:
		expiration += config.AccessTokenLifetimeMins * 60 * 1000
	case REFRESH_TOKEN:
		expiration += config.RefreshTokenLifetimeDays * 24 * 60 * 60 * 1000
	default:
		return "", errors.New("Unrecognized token type")
	}

	claims := TokenClaims{
		Type: params.Type,
		Expiration: expiration,
		UserId: params.UserId,
		UserEmail: params.UserEmail,
	}
	
	claimsJson, err := json.Marshal(claims)

	if err != nil {
		return "", errors.New("Failed to serialize token claims as JSON")
	}

	mac := hmac.New(sha256.New, []byte(config.TokenHashingKey))
	mac.Write(claimsJson)

	hash := mac.Sum(nil)
	encodedHash := hex.EncodeToString(hash)

	unencodedToken := string(claimsJson) + "|" + encodedHash
	encodedToken := base64.RawURLEncoding.EncodeToString([]byte(unencodedToken))

	return encodedToken, nil
}

func DecodeAndVerifyTokenForUser(token string, userId int64, tokenType TokenType) (*TokenClaims, error) {
	currentTimestamp := time.Now().UnixMilli()
	decodedToken, err := base64.RawURLEncoding.DecodeString(token)

	if err != nil {
		return nil, errors.New("Invalid token encoding")
	}
	
	indexOfDelimeter := bytes.LastIndexByte(decodedToken, byte('|'))

	if indexOfDelimeter < 1 || len(decodedToken) <= indexOfDelimeter {
		return nil, errors.New("Invalid token format")
	}

	providedHexHash := decodedToken[indexOfDelimeter + 1:]
	providedClaimsJson := decodedToken[:indexOfDelimeter]

	var claims TokenClaims

	if err = json.Unmarshal(providedClaimsJson, &claims); err != nil {
		return nil, errors.New("Invalid claims payload")
	}

	if claims.Expiration <= currentTimestamp {
		return nil, errors.New("Expired token")
	}

	if claims.Type != tokenType {
		return nil, errors.New("Incorrect token type")
	}

	if claims.UserId != userId {
		return nil, errors.New("Token does not belong to user")
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

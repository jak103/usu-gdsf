package auth

import (
	"encoding/base64"
	"testing"
	"time"

	"github.com/jak103/usu-gdsf/config"
	"github.com/stretchr/testify/assert"
)

func TestGenerateTokenDecodeAndValidate(t *testing.T) {
	params := TokenParams{
		Type: ACCESS_TOKEN,
		UserId: 42,
		UserType: REGULAR_USER,
		UserEmail: "testing@example.com",
	}

	token := GenerateToken(params)
	claims, _ := DecodeAndVerifyToken(token, params.Type)

	assert.Equal(t, claims.Type, params.Type)
	assert.Equal(t, claims.UserType, params.UserType)
	assert.Equal(t, claims.UserId, params.UserId)
	assert.Equal(t, claims.UserEmail, params.UserEmail)

	expectedExperationTime := time.Now().UnixMilli() + config.AccessTokenLifetimeMins*60*1000
	assert.True(t, claims.Expiration == expectedExperationTime || claims.Expiration == expectedExperationTime-1)

	params = TokenParams{
		Type: REFRESH_TOKEN,
		UserId: 999999,
		UserType: REGULAR_USER,
		UserEmail: "te|sting2@example.com|",
	}

	token = GenerateToken(params)
	claims, _ = DecodeAndVerifyToken(token, params.Type)

	assert.Equal(t, claims.Type, params.Type)
	assert.Equal(t, claims.UserType, params.UserType)
	assert.Equal(t, claims.UserId, params.UserId)
	assert.Equal(t, claims.UserEmail, params.UserEmail)

	expectedExperationTime = time.Now().UnixMilli() + config.RefreshTokenLifetimeDays*24*60*60*1000
	assert.True(t, claims.Expiration == expectedExperationTime || claims.Expiration == expectedExperationTime-1)
}

func TestEmptyToken(t *testing.T) {
	claims, err := DecodeAndVerifyToken("", ACCESS_TOKEN)
	assert.Nil(t, claims)
	assert.NotNil(t, err)
}

func TestInvalidTokenEncoding(t *testing.T) {
	params := TokenParams{
		Type: ACCESS_TOKEN,
		UserId: 42,
		UserType: REGULAR_USER,
		UserEmail: "testing@example.com",
	}

	token := GenerateToken(params)
	decodedToken, _ := base64.RawURLEncoding.DecodeString(token)

	claims, err := DecodeAndVerifyToken(string(decodedToken), params.Type)
	assert.Nil(t, claims)
	assert.NotNil(t, err)
}

func TestInvalidTokenFormat(t *testing.T) {
	fakeToken := "| this is a fake token string"
	encodedFakeToken := base64.RawURLEncoding.EncodeToString([]byte(fakeToken))

	claims, err := DecodeAndVerifyToken(encodedFakeToken, ACCESS_TOKEN)
	assert.Nil(t, claims)
	assert.NotNil(t, err)
}

func TestInvalidClaimsPayload(t *testing.T) {
	// Missing comma after type
	invalidJsonToken := "{\"Type\":0\"Expiration\":1658548929402,\"UserId\":42,\"UserEmail\":\"testing@example.com\"}|c0c7b1bc9c8eaae5ab5732fa5bddf9704fc71fbe3afb6af1a797af20fa0040db"
	encodedInvalidToken := base64.RawURLEncoding.EncodeToString([]byte(invalidJsonToken))

	claims, err := DecodeAndVerifyToken(encodedInvalidToken, ACCESS_TOKEN)
	assert.Nil(t, claims)
	assert.NotNil(t, err)
}

func TestExpiredToken(t *testing.T) {
	// These params were used to generate the following token
	// TokenParams{
	// 	Type: ACCESS_TOKEN,
	// 	UserId: 42,
	//      UserType: REGULAR_USER,
	// 	UserEmail: "testing@example.com",
	// }
	// The expiration timestamp is 1659402529985
	// The hashing key is ZUDx!8@76Dri!TI#v1O#Zh!IrzS%8FJZtCNmxqMo
		
	token := "eyJUeXBlIjowLCJFeHBpcmF0aW9uIjoxNjU5NDAyNTI5OTg1LCJVc2VySWQiOjQyLCJVc2VyVHlwZSI6MSwiVXNlckVtYWlsIjoidGVzdGluZ0BleGFtcGxlLmNvbSJ9fGE5N2QwMjViYzhkNzc2YjQwZjkyNmExNTBkMzg3OWEyY2Q5YTNmZTVkYTA1YTIyM2VlMDRmMDZlNGQxODFkZGI"

	claims, err := DecodeAndVerifyToken(token, ACCESS_TOKEN)
	assert.Nil(t, claims)
	assert.NotNil(t, err)
}

func TestIncorrectTokenType(t *testing.T) {
	params := TokenParams{
		Type:      ACCESS_TOKEN,
		UserId:    42,
		UserEmail: "testing@example.com",
	}

	token := GenerateToken(params)

	claims, err := DecodeAndVerifyToken(token, REFRESH_TOKEN)
	assert.Nil(t, claims)
	assert.NotNil(t, err)

	claims, err = DecodeAndVerifyToken(token, ACCESS_TOKEN)
	assert.NotNil(t, claims)
	assert.Nil(t, err)
}

func TestInvalidSignature(t *testing.T) {
	// The signature here ends with a "g" which is not a legal character for a hex encoding
	invalidSignatureToken := "{\"Type\":0,\"Expiration\":1659402529985,\"UserId\":42,\"UserType\":1,\"UserEmail\":\"testing@example.com\"}|a97d025bc8d776b40f926a150d3879a2cd9a3fe5da05a223ee04f06e4d181ddg"
	encodedInvalidToken := base64.RawURLEncoding.EncodeToString([]byte(invalidSignatureToken))

	claims, err := DecodeAndVerifyToken(encodedInvalidToken, ACCESS_TOKEN)
	assert.Nil(t, claims)
	assert.NotNil(t, err)
}

func TestIncorrectSignature(t *testing.T) {
	params := TokenParams{
		Type:      ACCESS_TOKEN,
		UserId:    39,
		UserEmail: "testing@example.com",
	}

	token := GenerateToken(params)
	decodedToken, _ := base64.RawURLEncoding.DecodeString(token)
	reencodedToken := base64.RawURLEncoding.EncodeToString([]byte(decodedToken))

	alteredToken := decodedToken

	if alteredToken[len(alteredToken)-1] != '1' {
		alteredToken[len(alteredToken)-1] = '1'
	} else {
		alteredToken[len(alteredToken)-1] = '2'
	}

	encodedAlteredToken := base64.RawURLEncoding.EncodeToString([]byte(alteredToken))

	claims, err := DecodeAndVerifyToken(reencodedToken, params.Type)
	assert.NotNil(t, claims)
	assert.Nil(t, err)

	claims, err = DecodeAndVerifyToken(encodedAlteredToken, params.Type)
	assert.Nil(t, claims)
	assert.NotNil(t, err)
}

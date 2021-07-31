package token

import (
	"fmt"
	"time"

	"github.com/architagr/golang-microservice-tutorial/authentication/models"

	"github.com/dgrijalva/jwt-go"
)

func Init() *models.ErrorDetail {
	flags, err := models.GetFlags()
	if err != nil {
		return err
	}

	ip, _ = flags.GetApplicationUrl()
	return nil
}

var ip *string

const (
	jWTPrivateToken = "SecrteTokenSecrteToken"
)

func GenrateToken(claims *models.JwtClaims, expirationTime time.Time) (string, *models.ErrorDetail) {


	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().UTC().Unix()
	claims.Issuer = *ip

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(jWTPrivateToken))
	if err != nil {
		return "", &models.ErrorDetail{
			ErrorType: models.ErrorTypeError,
			ErrorMessage: err.Error(),
		}
	}
	return tokenString, nil
}

func VerifyToken(tokenString, origin string) (bool, *models.JwtClaims) {
	claims := &models.JwtClaims{}
	token, _ := getTokenFromString(tokenString, claims)
	
	if token.Valid {
		if claims.VerifyAudience(origin) {
			return true, claims
		}
	}
	return false, nil
}

func GetClaims(tokenString string) models.JwtClaims {
	claims := &models.JwtClaims{}

	_, err := getTokenFromString(tokenString, claims)
	if err == nil {
		return *claims
	}
	return *claims
}
func getTokenFromString(tokenString string, claims *models.JwtClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(jWTPrivateToken), nil
	})
}

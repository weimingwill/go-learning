package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	AudienceAPP    = "io.flexrewards.app.development"
	AudienceAdmin  = "admin-dev.flextechdev.com"
	AudiencePortal = "portal-dev.flextechdev.com"
)

var hmacSecret []byte

type JWTClaims struct {
	Merchants []int64 `json:"merchants,omitempty"`
	jwt.StandardClaims
}

func main() {
	os.Setenv("hmac_secret", "this is hmac secret string")
	hmacSecret = []byte(os.Getenv("hmac_secret"))

	token := CreateJWT()
	ParseJWT(token)
}

func NewClaim() jwt.Claims {
	claims := JWTClaims{
		[]int64{},
		jwt.StandardClaims{
			Audience:  AudiencePortal,
			IssuedAt:  time.Now().UTC().Unix(),
			Subject:   "userID",
			ExpiresAt: time.Date(2018, 11, 17, 20, 34, 58, 651387237, time.UTC).Unix(),
			Id:        "jti uuid",
		},
	}

	return claims
}

func CreateJWT() string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, NewClaim())

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSecret)

	fmt.Println(tokenString, err)
	return tokenString
}

func ParseJWT(tokenString string) {
	// sample token is expired.  override time so it parses as valid
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSecret, nil
	})

	if token.Valid {
		claims, ok := token.Claims.(*JWTClaims)
		if !ok {
			fmt.Println("claim format error")
		}
		fmt.Printf("%v %v", claims.Id, claims.StandardClaims.ExpiresAt)
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("other errors")
	}

	// if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
	// } else {
	// 	fmt.Println(err)
	// }

}

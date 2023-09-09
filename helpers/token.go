package helpers

import (
	"GO-JWT/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)


var mySigninKey = []byte("mysecretkey")


type MyCustomClaims struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateToken(user *models.User)(string, error){
	claims := MyCustomClaims{
		user.ID,
		user.Name,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2*time.Minute)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
		
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	ss, err := token.SignedString(mySigninKey)

	return ss, err
}

func ValidateToken(tokenString string)(any, error){
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return mySigninKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}

	claims, ok := token.Claims.(*MyCustomClaims)

	if !ok || !token.Valid {
		return nil, fmt.Errorf("unauthorized")
	}

	return claims, nil
}
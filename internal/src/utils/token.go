package utils

import (
	"log"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)



type Claims struct {
	UserId uint   `json:"userid"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// accesstoken
func AccessToken(UserId uint, Email string, role string) (string, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	//key of enocode
	KEY := []byte(os.Getenv("DB_SECRET_KEY"))

	Claim := Claims{
		UserId: UserId,
		Email:  Email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	enodeToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claim)
	AccessToken, err := enodeToken.SignedString(KEY)
	return AccessToken, err
}

//Refersh token
func RefershToken(UserId uint, Email string, role string) (string, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	//key of enocode
	KEY := []byte(os.Getenv("DB_SECRET_KEY"))

	claim := Claims{
		UserId: UserId,
		Email:  Email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}

	EncodeToken:=jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	RefershToken,err:= EncodeToken.SignedString(KEY)
	return RefershToken,err

}

package jwt

import (
	"errors"
	"fmt"
	"ginApi/common/config"
	"github.com/golang-jwt/jwt"
	"time"
)

type Jwt struct {
}

type MyCustomClaims struct {
	UserId int `json:"UserId"`
	jwt.StandardClaims
}

var mySigningKey []byte

func init() {
	mySigningKey = []byte(config.Viper.GetString("token.key"))
}

func (this Jwt) CreateToken(UserId int) (string, error) {
	claims := MyCustomClaims{
		UserId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + config.Viper.GetInt64("token.expire"),
			Issuer:    "admin",
		},
	}
	result := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := result.SignedString(mySigningKey)
	return ss, err
}

func (this Jwt) ValidateToken(token string) (int, error) {
	result, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Printf("%s", err)
		return 0, err
	}

	if data, ok := result.Claims.(*MyCustomClaims); ok {
		fmt.Println(data)
		return data.UserId, nil
	} else {
		return 0, errors.New("assert error")
	}
}

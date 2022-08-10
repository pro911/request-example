package util

import (
	"github.com/dgrijalva/jwt-go"
	"request-example/config"
	"time"
)

//var jwtSecret = []byte("23347$040412")
//var jwtSecret = "23347$040412"
//var jwtSecret = []byte(config.JwtSecret)

var jwtSecret = []byte(config.AppConf.JwtSecret)

type Claims struct {
	Mobile  string `json:"mobile"`
	VerCode string `json:"ver_code"`
	jwt.StandardClaims
}

func GenerateToken(mobile, ver_code string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		Mobile:  mobile,
		VerCode: ver_code,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "pro911",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

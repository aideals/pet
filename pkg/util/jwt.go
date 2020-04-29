package util

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"

	"Pet/setting"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	OpenId string `json:"openid"`
	SessionKey string `json:"sessionkey"`
	jwt.StandardClaims
}

func GenerateToken(openid, sessionkey string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

    claims := Claims{
		  openid,
			sessionkey,
			jwt.StandardClaims {
				 ExpiresAt: expireTime.Unix(),
				 Issuer: "gin-pet",
			},
		}
 

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token,err := tokenClaims.SignedString(jwtSecret)
	
	return token,err
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
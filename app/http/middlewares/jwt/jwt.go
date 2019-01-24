package jwt

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/imofei/gin-easy/app/enums/codes"
	"github.com/imofei/gin-easy/app/http/response"
)

var jwtSecret = []byte("kucoin-kcs-oa")

type Claims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(UserId string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(8 * time.Hour)

	claims := Claims{
		UserId,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "kucoin-kcs",
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


func RefreshToken(token string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})


	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {

			jwt.TimeFunc = time.Now
			claims.StandardClaims.ExpiresAt = time.Now().Add(3 * time.Hour).Unix()
			return GenerateToken(token)

		}
	}
	return "", err
}



func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Token")
		if len(token) == 0 {
			response.FailJSON(c, codes.ParamError, "missing token")
			c.Abort()
			return
		}

		claims, err := ParseToken(token)
		if err != nil {
			response.FailJSON(c, codes.AuthCheckTokenFail, "")
			c.Abort()
			return
		}

		if time.Now().Unix() > claims.ExpiresAt {
			response.FailJSON(c, codes.AuthCheckTokenTimeout, "")
			c.Abort()
			return
		}

		c.Set("staff", "UserInfo")

		c.Next()
	}
}

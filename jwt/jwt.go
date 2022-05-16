package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func Create(data map[string]interface{}) (string, error) {
	data["exp"] = time.Now().Add(time.Minute * 60).Unix()
	var claims jwt.MapClaims = data
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func Validate(data string) (map[string]interface{}, error) {
	token, err := jwt.Parse(data, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// fmt.Println(claims["foo"], claims["nbf"])
		return claims, nil
	} else {
		return nil, err
	}
}

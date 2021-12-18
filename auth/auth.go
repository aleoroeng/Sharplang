package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

type MyClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

var myClaims MyClaims = MyClaims{
	"appname",
	jwt.StandardClaims{
		ExpiresAt: jwt.NewTime(15000),
		Issuer:    "test",
	},
}

var PrivateKey *ecdsa.PrivateKey

func Init(c *gin.Context) {
	PrivateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		log.Fatal(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, myClaims)
	ss, tokenErr := token.SignedString(PrivateKey)
	fmt.Printf("%v %v", ss, tokenErr)
	c.JSON(http.StatusOK, ss)
}

func GetToken(c *gin.Context) {

	rawToken, queryErr := c.GetQuery("token")

	if queryErr != false {
		fmt.Println(rawToken)
	}

	token, err := jwt.ParseWithClaims(rawToken, myClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		fmt.Println("Hi", PrivateKey.D)
		return PrivateKey, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if token.Valid {
		fmt.Println("valid token")
	} else {
		fmt.Println("not valid")
	}
}

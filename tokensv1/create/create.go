package main

import (
	"fmt"
	"time"
	"tokensv1/constants"
	"tokensv1/val"

	"github.com/dgrijalva/jwt-go"
)



func createToken(email, customerid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email, "customerid": customerid,
		"exp": time.Now().Add(time.Second*30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(constants.SecretKey))
	if err != nil {
		return "", err
	}
  //  fmt.Println("token: ->  ", tokenString)
	return tokenString, nil
}

func main() {

    token,err:=createToken("example@mail","Id1")
	if err!= nil {
		fmt.Println("token not generated")
        fmt.Println(err)
		return
    }
	fmt.Println("--> ",token)
	data,_:=val.ExtractCustomerID(token,constants.SecretKey)
	fmt.Println(data)

    
}

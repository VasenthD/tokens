package val

import (
	"fmt"
	"tokensv1/models"

	"github.com/dgrijalva/jwt-go"
)

func ExtractCustomerID(jwtToken string, secretKey string) (*models.Result, error) {

	// Parse the JWT token
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err // Handle token parsing errors
	}
	var result models.Result
	// Check if the token is valid
	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Extract the customer ID from the claims
			
			result.Customerid, _ = claims["customerid"].(string)
			result.Email, _ = claims["email"].(string)
			if ok {
				return &result, nil
			}
		}
	}
	fmt.Println(result)
	return &result, fmt.Errorf("Invalid or expired JWT token")
}

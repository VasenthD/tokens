package main

import (
	"fmt"
	"net/http"
	"tokensv1/constants"
	"tokensv1/val"

	"github.com/gin-gonic/gin"
)

func main() {

	routes := gin.Default()

	routes.POST("/check", func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		res, err := val.ExtractCustomerID(token, constants.SecretKey)
		if err != nil {
			fmt.Println("error in extracting customer ID: ", err)
			c.JSON(http.StatusBadRequest,gin.H{"invalied_token":token})
	

		}
		fmt.Println("valid customer ID: ", res.Customerid)
		c.JSON(http.StatusOK, gin.H{ "customer ID": res.Customerid})
	})

	routes.Run(":8880")
}

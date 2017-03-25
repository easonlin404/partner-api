package main

import (
	"gopkg.in/gin-gonic/gin.v1"

	"fmt"
)

// Binding from JSON
type ProvisionReq struct {
	CellPhone     string `json:"cellPhone" binding:"required"`
	PartnerUserId string `json:"partnerUserId" binding:"required"`
}

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	r.POST("/admin/partners/general/userProvision", func(c *gin.Context) {
		var json ProvisionReq
		if c.BindJSON(&json) == nil {
			//if json.CellPhone
			fmt.Println(json.CellPhone)
			fmt.Println(json.PartnerUserId)
			c.JSON(200, gin.H{
				"code":      "0",
				"accountId": "95a8b0b4-1aca-4bcc-b993-3d434d61482a",
				"cellPhone": json.CellPhone,
				"message":   "OK.",
			})
		} else {
			c.JSON(200, gin.H{

				"message": "error",
			})
		}

	})

	r.POST("/admin/partners/general/userActivation", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": "004",
			"message":"OK.",
		})
	})

	r.Run(":8081") // listen and serve on 0.0.0.0:8080

}

package main

import (
	pal "leetcode/internal"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	ginEngine := gin.Default()
	ginEngine.GET("/health", healthCheck)
	ginEngine.POST("/ispal", checkIsPalindrome)
	ginEngine.Run() // listen and serve on 0.0.0.0:8080

}

func healthCheck(ginCtx *gin.Context) {
	ginCtx.JSON(200, gin.H{
		"status":  "OK",
		"message": "Server is ready!",
	})
}

func checkIsPalindrome(ginCtx *gin.Context) {

	if ginCtx.Request.Method != http.MethodPost {
		ginCtx.JSON(400, gin.H{
			"message": "invalid request method",
		})
		return
	}

	palText := new(pal.Text)
	// Parse request body into Order struct
	if err := ginCtx.BindJSON(&palText); err != nil {
		log.Println(err)
		ginCtx.JSON(400, gin.H{
			"message": "invalid request body",
		})
		return
	}

	if pal.IsPalindromeStr(palText.Name) {
		ginCtx.JSON(200, gin.H{
			"status":  "OK",
			"message": palText.Name + " is a palindrome.",
		})
	} else {
		ginCtx.JSON(200, gin.H{
			"status":  "OK",
			"message": palText.Name + " is not a palindrome.",
		})
	}
}

package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AccessToken(c *gin.Context) {

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
	// 	ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
	// })

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
	})

	ss, err := token.SignedString([]byte("==signature=="))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": ss,
	})
}

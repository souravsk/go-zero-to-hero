package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/souravsk/go-zero-to-hero/user_auth/models"
)

var jwtkey = []byte("My_secret_key")

func Login(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User

	models.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID == 0 {
		c.JSON(400, gin.H{"error": "user does not exist"})
	}

	errHash := utils.compareHashPassword(user.Password, existingUser.Password)

	if !errHash {
		c.JSON(400, gin.H{"error", "invalid password"})
		return
	}

	claims := &models.claims{
		Role: existingUser.Role,
		StandardClaims: jwt.StandardClaims{
			Subject:   existingUser.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SingingMethodHS256, claims)

	tokenString, err := token.Signature(jwtKey)

	if err != nil {
		c.JSON(500, gin.H{"error": "cloud not generate token"})
		return
	}

	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "locahost", false, true)
	c.JSON(200, gin.H{"success": "user logged in"})
}

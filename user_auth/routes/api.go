// routes for a web application using the Gin HTTP framework.

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/souravsk/go-zero-to-hero/user_auth/controllers"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)
	r.POST("/singup", controllers.Signup)
	r.GET("/home", controllers.Home)
	r.GET("/premium", controllers.Premium)
	r.GET("/logout", controllers.Logout)
}

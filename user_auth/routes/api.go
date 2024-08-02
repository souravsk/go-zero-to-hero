// routes for a web application using the Gin HTTP framework.

package routes

import (
	"github.com/gin-gonic/gin"
	"k8s.io/apiserver/pkg/server/options/encryptionconfig/controller"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/login", controller.Login)
	r.POST("/singup", controller.Singup)
	r.GET("/home", controller.Home)
	r.GET("/premium", controller.Premium)
	r.GET("/logout", controller.Logout)
}

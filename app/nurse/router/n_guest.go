package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/nurse/apis"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerGuestRouter)
}

// registerGuestRouter
func registerGuestRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	guest := apis.NGuest{}
	r := v1.Group("/guest") //.Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", guest.GetPage)
		r.GET("/:id", guest.Get)
		r.POST("", guest.Insert)
		r.PUT("/:id", guest.Update)
	}
}

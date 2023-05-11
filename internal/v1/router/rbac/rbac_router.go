package rbac

import (
	rbac "d2din.com/internal/v1/presenter/rbac"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	//controller := rbac.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("policy")
	{
		v10.POST("", rbac.AddPolicy)
		v10.GET("", rbac.GetPolicy)
		v10.DELETE("", rbac.DeletePolicy)
	}

	return route
}

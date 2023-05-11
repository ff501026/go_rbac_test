package account

import (
	"d2din.com/internal/v1/middleware"
	"d2din.com/internal/v1/middleware/auth"
	presenter "d2din.com/internal/v1/presenter/account"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("account")
	{
		v10.POST("",middleware.Verify(),middleware.Transaction(db), controller.Created)
		v10.GET("", middleware.Verify(), controller.List)
		v10.GET(":accountID", middleware.Verify(), controller.GetByID)
		v10.DELETE(":accountID", auth.AuthCheckRole(), controller.Delete)
		v10.PATCH(":accountID", middleware.Verify(), controller.Updated)
	}

	return route
}

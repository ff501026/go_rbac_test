package rbac

import (
	"net/http"

	"d2din.com/internal/pkg/code"
	"d2din.com/internal/pkg/log"
	//"d2din.com/internal/v1/resolver/rbac"
	//"d2din.com/internal/v1/structure/jwe"
	//"d2din.com/internal/v1/structure/rbac"
	rbac "d2din.com/internal/v1/middleware/auth"
	"github.com/gin-gonic/gin"
	_"gorm.io/gorm"
)

type Presenter interface {
	AddPolicy(ctx *gin.Context)
	GetPolicy(ctx *gin.Context)
	DeletePolicy(ctx *gin.Context)
}

// type presenter struct {
// 	rbac rbac.Resolver
// }

// func New(db *gorm.DB) Presenter {
// 	return &presenter{
// 		rbac: rbac.New(db),
// 	}
// }

// AddPolicy
// @Summary 新增策略
// @description 新增策略
// @Tags rbac
// @version 1.0
// @Accept json
// @produce json
// @param * body rbacs.rbac true "登入帶入"
// @success 200 object code.SuccessfulMessage{body=jwe.Token} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/rbac/web [post]
func AddPolicy(ctx *gin.Context) {
	input := &rbac.CasbinModel{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))
		return
	}

	result,err := rbac.AddCasbin(*input)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.InternalServerError, err.Error()))
		return
	}
	if result != true {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.DoesNotExist, "Policy已經存在"))
		return
	}
	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.Successful, "新增成功"))
}

// AddPolicy
// @Summary 取得策略
// @description 取得策略
// @Tags rbac
// @version 1.0
// @Accept json
// @produce json
// @param * body rbacs.rbac true "登入帶入"
// @success 200 object code.SuccessfulMessage{body=jwe.Token} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/rbac/web [post]
func GetPolicy(ctx *gin.Context) {

	result := rbac.GetAllPolice()

	if result == nil {
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.DoesNotExist, "查無Policy"))
		return
	}
	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.Successful, result))
}

// // Refresh
// // @Summary 刪除策略
// // @description 刪除策略
// // @Tags rbac
// // @version 1.0
// // @Accept json
// // @produce json
// // @param * body jwe.Refresh true "登入帶入"
// // @success 200 object code.SuccessfulMessage{body=jwe.Token} "成功後返回的值"
// // @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// // @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// // @Router /authority/v1.0/rbac/refresh [post]
func DeletePolicy(ctx *gin.Context) {
	input := &rbac.CasbinModel{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))
		return
	}

	result,err := rbac.DeleteCasbin(*input)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.InternalServerError, err.Error()))
		return
	}
	if result != true {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.DoesNotExist, "Policy不存在"))
		return
	}
	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.Successful, "刪除成功"))
}

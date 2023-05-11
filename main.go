package main

// main is run all api form localhost port 8080
import (
	"fmt"
	"net/http"

	_ "d2din.com/api"
	"d2din.com/internal/pkg/dao/gorm"
	"d2din.com/internal/pkg/log"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"d2din.com/internal/v1/router"
	//account "d2din.com/internal/v1/router/account"
	//company "d2din.com/internal/v1/router/company"
	rbac "d2din.com/internal/v1/router/rbac"


	_"d2din.com/internal/v1/structure/accounts"
	_"d2din.com/internal/v1/structure/companies"
	//"d2din.com/internal/v1/router/login"
)

type Presenter interface {
	BydProductToPos(ctx *gin.Context)
}

// @title          hsmaster SYSTEM API
// @version        0.1
// @description    企業系統整合管理平台
// @termsOfService https://hamastar.app/

// @contact.name  API System Support
// @contact.url   https://hamastar.app/
// @contact.email mingzong.lyu@gmail.com

// @license.name AGPL 3.0
// @license.url  https://www.gnu.org/licenses/agpl-3.0.en.html

// @host     api.hamastar.app
// @BasePath /
// @schemes  https
func main() {
	db, err := gorm.New()
	db.AutoMigrate(
		// &accounts.Table{},
		// &companies.Table{},
	)
	if err != nil {
		log.Error(err)
		return
	}

	route := router.Default()
	//account.GetRoute(route, db)
	//login.GetRoute(route, db)
	//company.GetRoute(route, db)
	rbac.GetRoute(route, db)


	url := ginSwagger.URL(fmt.Sprintf("http://localhost:8080/swagger/doc.json"))
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	log.Fatal(http.ListenAndServe(":8080", route))
	//Force update
}

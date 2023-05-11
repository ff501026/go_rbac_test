package main

// authority rbac
import (
	"d2din.com/internal/pkg/dao/gorm"
	"d2din.com/internal/pkg/log"
	"d2din.com/internal/v1/router"
	// "d2din.com/internal/v1/router/account"
	// "d2din.com/internal/v1/router/company"
	"d2din.com/internal/v1/router/rbac"

	"github.com/apex/gateway"
)

func main() {
	db, err := gorm.New()
	if err != nil {
		log.Error(err)
		return
	}

	route := router.Default()
	//route = account.GetRoute(route, db)
	//route = company.GetRoute(route, db)
	route = rbac.GetRoute(route, db)
	log.Fatal(gateway.ListenAndServe(":8080", route))
}

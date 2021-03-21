package routes

import (
	"superstar/bootstrap"
	"superstar/services"
	"superstar/web/controllers"
	"superstar/web/middleware"

	"github.com/kataras/iris/v12/mvc"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	superstarService := services.NewSuperStarService()

	index := mvc.New(b.Party("/"))
	index.Register(superstarService)
	index.Handle(new(controllers.IndexController))

	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(superstarService)
	admin.Handle(new(controllers.AdminController))
}

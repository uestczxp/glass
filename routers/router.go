package routers

import (
	"github.com/astaxie/beego"
	"glass/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/glass", &controllers.IndexController{})
}

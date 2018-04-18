package routers

import (
	"weixin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/wx", &controllers.MainController{})
}

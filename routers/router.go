package routers

import (
	"weixin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/ququ", &controllers.QuquController{})
}

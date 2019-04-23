package routers

import (
	"firstbee/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{}) // 因为user的controller中已经有了/路径的路由，所以此处注释掉
    //beego.Include(&controllers.UserController{}) // 引入user controller
}

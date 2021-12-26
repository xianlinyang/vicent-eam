package routers

import (
	"Eam_Server/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    //用户管理
	beego.Router("/User/UserLogin", &controllers.UserController{},"post:UserLogin")
	beego.Router("/User/browse", &controllers.UserController{},"post:UserQuery")
	beego.Router("/User/create", &controllers.UserController{},"post:UserAdd")
	beego.Router("/User/update", &controllers.UserController{},"put:UserUpdate")
	beego.Router("/User/delete", &controllers.UserController{},"delete:UserDel")
}

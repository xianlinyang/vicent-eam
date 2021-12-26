package routers

import (
	"Eam_Server/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//用户管理
	beego.Router("/api/User/UserLogin", &controllers.UserController{},"post:UserLogin")
	beego.Router("/api/User/browse", &controllers.UserController{},"get:UserQuery")
	beego.Router("/api/User/create", &controllers.UserController{},"post:UserAdd")
	beego.Router("/api/User/update", &controllers.UserController{},"put:UserUpdate")
	beego.Router("/api/User/delete", &controllers.UserController{},"delete:UserDel")

	beego.Router("/api/User/SetStatus", &controllers.UserController{},"post:SetStatus")  //启用，停用接口
	beego.Router("/api/User/SetPwd", &controllers.UserController{},"post:SetPwd")        //修改密码

	//微信登陆检测
	beego.Router("/api/User/wechat/UserCheck", &controllers.UserController{},"post:Wechat_UserCkAndLogin")

}

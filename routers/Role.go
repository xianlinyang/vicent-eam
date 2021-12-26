package routers

import (
	"Eam_Server/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//角色管理
	beego.Router("/api/Role/browse", &controllers.RoleController{},"get:RoleQuery")
	beego.Router("/api/Role/create", &controllers.RoleController{},"post:RoleAdd")
	beego.Router("/api/Role/update", &controllers.RoleController{},"put:RoleUpdate")
	beego.Router("/api/Role/delete", &controllers.RoleController{},"delete:RoleDel")
	beego.Router("/api/Role/Qxbrowse", &controllers.RoleController{},"get:Qxbrowse")
	beego.Router("/api/Role/RoleDeptbrowse", &controllers.RoleController{},"get:RoleDeptbrowse")

}


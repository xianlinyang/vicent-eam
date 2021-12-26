package routers

import (
	"Eam_Server/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//日志查询
	beego.Router("/api/SysLog/browse", &controllers.OptionsController{},"get:SysLog_Query")
    //单据录入 -商品选择
	beego.Router("/api/public/Input_Query", &controllers.OptionsController{},"get:Public_Input_Query")
    //单据录入 -部门选择
	beego.Router("/api/public/Input_Dept_Query", &controllers.OptionsController{},"get:Public_Input_Dept_Query")
	//单据录入 -存放位置选择
	beego.Router("/api/public/Input_Postion_Query", &controllers.OptionsController{},"get:Public_Input_Postion_Query")
	//单据录入 -人员选择
	beego.Router("/api/public/Input_People_Query", &controllers.OptionsController{},"get:Public_Input_People_Query")
	//单据录入 -部门，存放位置 选择
	beego.Router("/api/public/Input_Bill_BasicAll", &controllers.OptionsController{},"get:Public_Input_Bill_BasicAll")
	//资产流水查看
	beego.Router("/api/public/AssetsFlowing", &controllers.OptionsController{},"get:Public_AssetsFlowing")


}
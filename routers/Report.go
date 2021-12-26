package routers

import (
"Eam_Server/controllers"
"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ReportController{})

	//报表管理

	//资产类别分析
	beego.Router("/api/Report_As_Type/browse", &controllers.ReportController{},"get:Report_As_Type_Browse")
	//部门汇总报表
	beego.Router("/api/Report_As_DeptCount/browse", &controllers.ReportController{},"get:Report_As_DeptCount_Browse")
	//资产变更分析
	beego.Router("/api/Report_As_Change/browse", &controllers.ReportController{},"get:Report_As_Change_Browse")



}

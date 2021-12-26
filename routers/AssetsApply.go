package routers


import (
	"Eam_Server/controllers"
	"github.com/astaxie/beego"
)
func init() {
	beego.Router("/", &controllers.MainController{})


	//领用申请
	beego.Router("/api/ReceiveApply/browse/main", &controllers.AssetsApplyController{},"get:ReceiveApply_MainQuery")
	beego.Router("/api/ReceiveApply/browse/detail", &controllers.AssetsApplyController{},"get:ReceiveApply_DetailQuery")
	beego.Router("/api/ReceiveApply/create", &controllers.AssetsApplyController{},"post:ReceiveApplyAdd")
	beego.Router("/api/ReceiveApply/update", &controllers.AssetsApplyController{},"put:ReceiveApplyUpdate")
	beego.Router("/api/ReceiveApply/delete", &controllers.AssetsApplyController{},"delete:ReceiveApplyDel")
	beego.Router("/api/ReceiveApply/submit/back", &controllers.AssetsApplyController{},"post:ReceiveApplySubmit_Back")  //辙消提交
	beego.Router("/api/ReceiveApply/ManagerApproval", &controllers.AssetsApplyController{},"post:ReceiveApplyManagerApproval")  //主管审批
	beego.Router("/api/ReceiveApply/AsseetApproval", &controllers.AssetsApplyController{},"post:ReceiveApplyAsseetApproval")   //仓库审批
	beego.Router("/api/ReceiveApply/ManagerApproval/back", &controllers.AssetsApplyController{},"post:ReceiveApplyManagerApproval_Back")  //辙消主管审批
	beego.Router("/api/ReceiveApply/AsseetApproval/back", &controllers.AssetsApplyController{},"post:ReceiveApplyAsseetApproval_Back")   //辙消仓库审批

	//领用
	beego.Router("/api/AssetReceive/browse/main", &controllers.AssetsApplyController{},"get:Receive_MainQuery")
	beego.Router("/api/AssetReceive/browse/detail", &controllers.AssetsApplyController{},"get:Receive_DetailQuery")
	beego.Router("/api/AssetReceive/create", &controllers.AssetsApplyController{},"post:ReceiveAdd")
	beego.Router("/api/AssetReceive/update", &controllers.AssetsApplyController{},"put:ReceiveUpdate")
	beego.Router("/api/AssetReceive/delete", &controllers.AssetsApplyController{},"delete:ReceiveDel")
	beego.Router("/api/AssetReceive/submit/back", &controllers.AssetsApplyController{},"post:ReceiveSubmit_Back")  //辙消提交

	//领用退库
	beego.Router("/api/AssetsBack/browse/main", &controllers.AssetsApplyController{},"get:AssetsBack_MainQuery")
	beego.Router("/api/AssetsBack/browse/detail", &controllers.AssetsApplyController{},"get:AssetsBack_DetailQuery")
	beego.Router("/api/AssetsBack/create", &controllers.AssetsApplyController{},"post:AssetsBackAdd")
	beego.Router("/api/AssetsBack/update", &controllers.AssetsApplyController{},"put:AssetsBackUpdate")
	beego.Router("/api/AssetsBack/delete", &controllers.AssetsApplyController{},"delete:AssetsBackDel")
	beego.Router("/api/AssetsBack/submit/back", &controllers.AssetsApplyController{},"post:AssetsBackSubmit_Back")  //辙消提交


	//借用
	beego.Router("/api/AssetsLend/browse/main", &controllers.AssetsApplyController{},"get:AssetsLend_MainQuery")
	beego.Router("/api/AssetsLend/browse/detail", &controllers.AssetsApplyController{},"get:AssetsLend_DetailQuery")
	beego.Router("/api/AssetsLend/create", &controllers.AssetsApplyController{},"post:AssetsLendAdd")
	beego.Router("/api/AssetsLend/update", &controllers.AssetsApplyController{},"put:AssetsLendUpdate")
	beego.Router("/api/AssetsLend/delete", &controllers.AssetsApplyController{},"delete:AssetsLendDel")
	beego.Router("/api/AssetsLend/submit/back", &controllers.AssetsApplyController{},"post:AssetsLendSubmit_Back")  //辙消提交

    //借用归还
	beego.Router("/api/AssetsLendIn/browse/main", &controllers.AssetsApplyController{},"get:AssetsLendIn_MainQuery")
	beego.Router("/api/AssetsLendIn/browse/detail", &controllers.AssetsApplyController{},"get:AssetsLendIn_DetailQuery")
	beego.Router("/api/AssetsLendIn/create", &controllers.AssetsApplyController{},"post:AssetsLendInAdd")
	beego.Router("/api/AssetsLendIn/update", &controllers.AssetsApplyController{},"put:AssetsLendInUpdate")
	beego.Router("/api/AssetsLendIn/delete", &controllers.AssetsApplyController{},"delete:AssetsLendInDel")
	beego.Router("/api/AssetsLendIn/submit/back", &controllers.AssetsApplyController{},"post:AssetsLendInSubmit_Back")  //辙消提交

	//调拨
	beego.Router("/api/AssetsAllot/browse/main", &controllers.AssetsApplyController{},"get:AssetsAllot_MainQuery")
	beego.Router("/api/AssetsAllot/browse/detail", &controllers.AssetsApplyController{},"get:AssetsAllot_DetailQuery")
	beego.Router("/api/AssetsAllot/create", &controllers.AssetsApplyController{},"post:AssetsAllotAdd")
	beego.Router("/api/AssetsAllot/update", &controllers.AssetsApplyController{},"put:AssetsAllotUpdate")
	beego.Router("/api/AssetsAllot/delete", &controllers.AssetsApplyController{},"delete:AssetsAllotDel")
	beego.Router("/api/AssetsAllot/Audit", &controllers.AssetsApplyController{},"post:AssetsAllot_Audit")  //审批
	beego.Router("/api/AssetsAllot/StockOut", &controllers.AssetsApplyController{},"post:AssetsAllot_StockOut")  //调出
	beego.Router("/api/AssetsAllot/StockIn", &controllers.AssetsApplyController{},"post:AssetsAllot_StockIn")  //调入
	beego.Router("/api/AssetsAllot/Back", &controllers.AssetsApplyController{},"post:AssetsAllotSubmit_Back") //辙消

	//资产变更
	beego.Router("/api/AssetsChange/browse/main", &controllers.AssetsApplyController{},"get:AssetsChange_MainQuery")
	beego.Router("/api/AssetsChange/browse/detail", &controllers.AssetsApplyController{},"get:AssetsChange_DetailQuery")
	beego.Router("/api/AssetsChange/create", &controllers.AssetsApplyController{},"post:AssetsChangeAdd")
	beego.Router("/api/AssetsChange/update", &controllers.AssetsApplyController{},"put:AssetsChangeUpdate")
	beego.Router("/api/AssetsChange/delete", &controllers.AssetsApplyController{},"delete:AssetsChangeDel")
	beego.Router("/api/AssetsChange/submit", &controllers.AssetsApplyController{},"post:AssetsChangeSubmit")
	beego.Router("/api/AssetsChange/apply/back", &controllers.AssetsApplyController{},"post:AssetsChangeSubmit_Back") //辙消申请
}



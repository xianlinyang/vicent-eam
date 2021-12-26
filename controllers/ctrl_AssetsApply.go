package controllers

import (
	"Eam_Server/models"
	"fmt"
	"github.com/astaxie/beego"
	"Eam_Server/utils"
)

type AssetsApplyController struct {
	beego.Controller
}


//领用申请 主单查询
func (ctl *AssetsApplyController) ReceiveApply_MainQuery(){
	response := utils.ResponsModel{};
	var err error
	//fmt.Println(ctl.Ctx.Request.RequestURI)
	//fmt.Println(ctl.Ctx.Request.URL.Query())//--这个有用
	//fmt.Println(ctl.Ctx.Input.URL())
	//fmt.Println(ctl.Ctx.Request.Form)


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
			response.Success = false
		}
	}()

	amode := models.AssetsApplyModel{};
	response.Content,err =  amode.AssetsApply_Main_GetList(ctl.Controller);

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="查询成功！"
		response.Success = true
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="查询失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//领用申请 从单查询
func (ctl *AssetsApplyController) ReceiveApply_DetailQuery(){
	response := utils.ResponsModel{};
	var err error
	//fmt.Println(ctl.Ctx.Request.RequestURI)
	//fmt.Println(ctl.Ctx.Request.URL.Query())//--这个有用
	//fmt.Println(ctl.Ctx.Input.URL())
	//fmt.Println(ctl.Ctx.Request.Form)


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
			response.Success = false
		}
	}()

	amode := models.AssetsApplyModel{};
	response.Content,err =  amode.AssetsApply_DetailQuery(ctl.Controller);

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="查询成功！"
		response.Success = true
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="查询失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//领用申请新增
func (ctl *AssetsApplyController) ReceiveApplyAdd(){
	response := utils.ResponsModel{};


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode
			response.Message = err.(string)
			response.Success = false

			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};


	aid,err := amode.AssetsApply_Create(ctl.Controller)

	if err ==nil{
		rdata := make(map[string]interface{})
		rdata["id"] = aid
		response.Code = utils.SuccessCode
		response.Message = "新增成功"
		response.Success = true
		response.Content.Result = append(response.Content.Result, rdata)
	} else {
		response.Code = utils.FailedCode
		response.Message = "新增出错:"+err.Error()
		response.Success = false
	}


	ctl.Data["json"] = response
	ctl.ServeJSON();
}

//领用申请修改
func (ctl *AssetsApplyController) ReceiveApplyUpdate(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode
			response.Message = err.(string)

			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	_,err := amode.AssetsApply_Update(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode
		response.Message = "修改成功！";
		response.Success = true
	} else {
		response.Code = utils.FailedCode
		response.Message = "修改失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}

//领用申请删除
func (ctl *AssetsApplyController) ReceiveApplyDel(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode;
			response.Message = err.(string)
			response.Success = false
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	err := amode.AssetsApply_Delete(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode;
		response.Message = "删除成功!"
		response.Success = true
	} else {
		response.Code = utils.FailedCode;
		response.Message = "删除失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}

//领用申请提交
func (ctl *AssetsApplyController) ReceiveApplySubmit_Back(){
	response := utils.ResponsModel{};
	var err error

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
			response.Success = false
		}
	}()

	amode := models.AssetsApplyModel{};
	err =  amode.AssetsApply_SetStatus(ctl.Controller,"-1");

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="辙消成功！"
		response.Success = false
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="辙消失败,"+err.Error()
		response.Success = true
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}

//领用申请主管审批
func (ctl *AssetsApplyController) ReceiveApplyManagerApproval(){
	response := utils.ResponsModel{};
	var err error

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
			response.Success = false
		}
	}()

	amode := models.AssetsApplyModel{};
	err =  amode.AssetsApply_SetStatus(ctl.Controller,"2");

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="审批成功！"
		response.Success = true
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="审批失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}

//领用申请仓库审批
func (ctl *AssetsApplyController) ReceiveApplyAsseetApproval(){
	response := utils.ResponsModel{};
	var err error

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
			response.Success = false
		}
	}()

	amode := models.AssetsApplyModel{};
	err =  amode.AssetsApply_SetStatus(ctl.Controller,"3");

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="审批成功！"
		response.Success = true
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="审批失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}

//领用申请主管审批-辙消
func (ctl *AssetsApplyController) ReceiveApplyManagerApproval_Back(){
	response := utils.ResponsModel{};
	var err error

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
			response.Success = false
		}
	}()

	amode := models.AssetsApplyModel{};
	err =  amode.AssetsApply_SetStatus(ctl.Controller,"-2");

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="辙消成功！"
		response.Success = true
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="辙消失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}

//仓库审批-辙消
func (ctl *AssetsApplyController) ReceiveApplyAsseetApproval_Back(){
	response := utils.ResponsModel{};
	var err error

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
			response.Success = false
		}
	}()

	amode := models.AssetsApplyModel{};
	err =  amode.AssetsApply_SetStatus(ctl.Controller,"-3");

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="辙消成功！"
		response.Success = true
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="辙消失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}



//领用 主单查询
func (ctl *AssetsApplyController) Receive_MainQuery(){
	response := utils.ResponsModel{};
	var err error
	//fmt.Println(ctl.Ctx.Request.RequestURI)
	//fmt.Println(ctl.Ctx.Request.URL.Query())//--这个有用
	//fmt.Println(ctl.Ctx.Input.URL())
	//fmt.Println(ctl.Ctx.Request.Form)


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
			response.Success = false
		}
	}()

	amode := models.AssetsApplyModel{};
	response.Content,err =  amode.Receive_Main_GetList(ctl.Controller);

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="查询成功！"
		response.Success = true
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="查询失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//领用 从单查询
func (ctl *AssetsApplyController) Receive_DetailQuery(){
	response := utils.ResponsModel{};
	var err error
	//fmt.Println(ctl.Ctx.Request.RequestURI)
	//fmt.Println(ctl.Ctx.Request.URL.Query())//--这个有用
	//fmt.Println(ctl.Ctx.Input.URL())
	//fmt.Println(ctl.Ctx.Request.Form)


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
			response.Success = false
		}
	}()

	amode := models.AssetsApplyModel{};
	response.Content,err =  amode.Receive_DetailQuery(ctl.Controller);

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="查询成功！"
		response.Success = true
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="查询失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//领用新增
func (ctl *AssetsApplyController) ReceiveAdd(){
	response := utils.ResponsModel{};


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode
			response.Message = err.(string)
			response.Success = false

			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};


	aid,err := amode.Receive_Create(ctl.Controller)

	if err ==nil{
		rdata := make(map[string]interface{})
		rdata["id"] = aid
		response.Code = utils.SuccessCode
		response.Message = "新增成功"
		response.Success = true
		response.Content.Result = append(response.Content.Result, rdata)
	} else {
		response.Code = utils.FailedCode
		response.Message = "新增出错:"+err.Error()
		response.Success = false
	}


	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//领用修改
func (ctl *AssetsApplyController) ReceiveUpdate(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode
			response.Message = err.(string)

			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	_,err := amode.Receive_Update(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode
		response.Message = "修改成功！";
		response.Success = true
	} else {
		response.Code = utils.FailedCode
		response.Message = "修改失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//领用删除
func (ctl *AssetsApplyController) ReceiveDel(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode;
			response.Message = err.(string)
			response.Success = false
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	err := amode.Receive_Delete(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode;
		response.Message = "删除成功!"
		response.Success = true
	} else {
		response.Code = utils.FailedCode;
		response.Message = "删除失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//领用辙消到草稿
func (ctl *AssetsApplyController) ReceiveSubmit_Back(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode;
			response.Message = err.(string)
			response.Success = false
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	err := amode.Receive_SetStatus(ctl.Controller,"-1")

	if err == nil {
		response.Code = utils.SuccessCode;
		response.Message = "辙消成功!"
		response.Success = true
	} else {
		response.Code = utils.FailedCode;
		response.Message = "辙消失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}


//退还 主单查询
func (ctl *AssetsApplyController) AssetsBack_MainQuery(){
	response := utils.ResponsModel{};
	var err error
	//fmt.Println(ctl.Ctx.Request.RequestURI)
	//fmt.Println(ctl.Ctx.Request.URL.Query())//--这个有用
	//fmt.Println(ctl.Ctx.Input.URL())
	//fmt.Println(ctl.Ctx.Request.Form)


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
			response.Success = false
		}
	}()

	amode := models.AssetsApplyModel{};
	response.Content,err =  amode.AssetsBack_MainQuery(ctl.Controller);

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="查询成功！"
		response.Success = true
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="查询失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//退还 从单查询
func (ctl *AssetsApplyController) AssetsBack_DetailQuery(){
	response := utils.ResponsModel{};
	var err error
	//fmt.Println(ctl.Ctx.Request.RequestURI)
	//fmt.Println(ctl.Ctx.Request.URL.Query())//--这个有用
	//fmt.Println(ctl.Ctx.Input.URL())
	//fmt.Println(ctl.Ctx.Request.Form)


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
			response.Success = false
		}
	}()

	amode := models.AssetsApplyModel{};
	response.Content,err =  amode.AssetsBack_DetailQuery(ctl.Controller);

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="查询成功！"
		response.Success = true
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="查询失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//退还新增
func (ctl *AssetsApplyController) AssetsBackAdd(){
	response := utils.ResponsModel{};


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode
			response.Message = err.(string)
			response.Success = false

			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};


	aid,err := amode.AssetsBackAdd(ctl.Controller)

	if err ==nil{
		rdata := make(map[string]interface{})
		rdata["id"] = aid
		response.Code = utils.SuccessCode
		response.Message = "新增成功"
		response.Success = true
		response.Content.Result = append(response.Content.Result, rdata)
	} else {
		response.Code = utils.FailedCode
		response.Message = "新增出错:"+err.Error()
		response.Success = false
	}


	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//退还修改
func (ctl *AssetsApplyController) AssetsBackUpdate(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode
			response.Message = err.(string)

			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	_,err := amode.AssetsBackUpdate(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode
		response.Message = "修改成功！";
		response.Success = true
	} else {
		response.Code = utils.FailedCode
		response.Message = "修改失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//退还删除
func (ctl *AssetsApplyController) AssetsBackDel(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode;
			response.Message = err.(string)
			response.Success = false
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	err := amode.AssetsBackDel(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode;
		response.Message = "删除成功!"
		response.Success = true
	} else {
		response.Code = utils.FailedCode;
		response.Message = "删除失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//退还辙消到草稿
func (ctl *AssetsApplyController) AssetsBackSubmit_Back(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode;
			response.Message = err.(string)
			response.Success = false
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	err := amode.AssetsBackSubmit_Back(ctl.Controller,"-1")

	if err == nil {
		response.Code = utils.SuccessCode;
		response.Message = "辙消成功!"
		response.Success = true
	} else {
		response.Code = utils.FailedCode;
		response.Message = "辙消失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}


//借用 主单查询
func (ctl *AssetsApplyController) AssetsLend_MainQuery(){
	response := utils.ResponsModel{};
	var err error
	//fmt.Println(ctl.Ctx.Request.RequestURI)
	//fmt.Println(ctl.Ctx.Request.URL.Query())//--这个有用
	//fmt.Println(ctl.Ctx.Input.URL())
	//fmt.Println(ctl.Ctx.Request.Form)


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
			response.Success = false
		}
	}()

	amode := models.AssetsApplyModel{};
	response.Content,err =  amode.AssetsLend_MainQuery(ctl.Controller);

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="查询成功！"
		response.Success = true
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="查询失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//借用 从单查询
func (ctl *AssetsApplyController) AssetsLend_DetailQuery(){
	response := utils.ResponsModel{};
	var err error
	//fmt.Println(ctl.Ctx.Request.RequestURI)
	//fmt.Println(ctl.Ctx.Request.URL.Query())//--这个有用
	//fmt.Println(ctl.Ctx.Input.URL())
	//fmt.Println(ctl.Ctx.Request.Form)


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
			response.Success = false
		}
	}()

	amode := models.AssetsApplyModel{};
	response.Content,err =  amode.AssetsLend_DetailQuery(ctl.Controller);

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="查询成功！"
		response.Success = true
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="查询失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//借用新增
func (ctl *AssetsApplyController) AssetsLendAdd(){
	response := utils.ResponsModel{};


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode
			response.Message = err.(string)
			response.Success = false

			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};


	aid,err := amode.AssetsLendAdd(ctl.Controller)

	if err ==nil{
		rdata := make(map[string]interface{})
		rdata["id"] = aid
		response.Code = utils.SuccessCode
		response.Message = "新增成功"
		response.Success = true
		response.Content.Result = append(response.Content.Result, rdata)
	} else {
		response.Code = utils.FailedCode
		response.Message = "新增出错:"+err.Error()
		response.Success = false
	}


	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//借用修改
func (ctl *AssetsApplyController) AssetsLendUpdate(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode
			response.Message = err.(string)

			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	_,err := amode.AssetsLendUpdate(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode
		response.Message = "修改成功！";
		response.Success = true
	} else {
		response.Code = utils.FailedCode
		response.Message = "修改失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//借用删除
func (ctl *AssetsApplyController) AssetsLendDel(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode;
			response.Message = err.(string)
			response.Success = false
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	err := amode.AssetsLendDel(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode;
		response.Message = "删除成功!"
		response.Success = true
	} else {
		response.Code = utils.FailedCode;
		response.Message = "删除失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//借用辙消到草稿
func (ctl *AssetsApplyController) AssetsLendSubmit_Back(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode;
			response.Message = err.(string)
			response.Success = false
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	err := amode.AssetsLendSubmit_Back(ctl.Controller,"-1")

	if err == nil {
		response.Code = utils.SuccessCode;
		response.Message = "辙消成功!"
		response.Success = true
	} else {
		response.Code = utils.FailedCode;
		response.Message = "辙消失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}


//借用归还 主单查询
func (ctl *AssetsApplyController) AssetsLendIn_MainQuery(){
	response := utils.ResponsModel{};
	var err error
	//fmt.Println(ctl.Ctx.Request.RequestURI)
	//fmt.Println(ctl.Ctx.Request.URL.Query())//--这个有用
	//fmt.Println(ctl.Ctx.Input.URL())
	//fmt.Println(ctl.Ctx.Request.Form)


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
			response.Success = false
		}
	}()

	amode := models.AssetsApplyModel{};
	response.Content,err =  amode.AssetsLendIn_MainQuery(ctl.Controller);

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="查询成功！"
		response.Success = true
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="查询失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//借用归还 从单查询
func (ctl *AssetsApplyController) AssetsLendIn_DetailQuery(){
	response := utils.ResponsModel{};
	var err error
	//fmt.Println(ctl.Ctx.Request.RequestURI)
	//fmt.Println(ctl.Ctx.Request.URL.Query())//--这个有用
	//fmt.Println(ctl.Ctx.Input.URL())
	//fmt.Println(ctl.Ctx.Request.Form)


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
			response.Success = false
		}
	}()

	amode := models.AssetsApplyModel{};
	response.Content,err =  amode.AssetsLendIn_DetailQuery(ctl.Controller);

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="查询成功！"
		response.Success = true
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="查询失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//借用归还新增
func (ctl *AssetsApplyController) AssetsLendInAdd(){
	response := utils.ResponsModel{};


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode
			response.Message = err.(string)
			response.Success = false

			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};


	aid,err := amode.AssetsLendInAdd(ctl.Controller)

	if err ==nil{
		rdata := make(map[string]interface{})
		rdata["id"] = aid
		response.Code = utils.SuccessCode
		response.Message = "新增成功"
		response.Success = true
		response.Content.Result = append(response.Content.Result, rdata)
	} else {
		response.Code = utils.FailedCode
		response.Message = "新增出错:"+err.Error()
		response.Success = false
	}


	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//借用归还修改
func (ctl *AssetsApplyController) AssetsLendInUpdate(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode
			response.Message = err.(string)

			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	_,err := amode.AssetsLendInUpdate(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode
		response.Message = "修改成功！";
		response.Success = true
	} else {
		response.Code = utils.FailedCode
		response.Message = "修改失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//借用归还删除
func (ctl *AssetsApplyController) AssetsLendInDel(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode;
			response.Message = err.(string)
			response.Success = false
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	err := amode.AssetsLendInDel(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode;
		response.Message = "删除成功!"
		response.Success = true
	} else {
		response.Code = utils.FailedCode;
		response.Message = "删除失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//借用归还辙消到草稿
func (ctl *AssetsApplyController) AssetsLendInSubmit_Back(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode;
			response.Message = err.(string)
			response.Success = false
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	err := amode.AssetsLendInSubmit_Back(ctl.Controller,"-1")

	if err == nil {
		response.Code = utils.SuccessCode;
		response.Message = "辙消成功!"
		response.Success = true
	} else {
		response.Code = utils.FailedCode;
		response.Message = "辙消失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}



//调拨 主单查询
func (ctl *AssetsApplyController) AssetsAllot_MainQuery(){
	response := utils.ResponsModel{};
	var err error
	//fmt.Println(ctl.Ctx.Request.RequestURI)
	//fmt.Println(ctl.Ctx.Request.URL.Query())//--这个有用
	//fmt.Println(ctl.Ctx.Input.URL())
	//fmt.Println(ctl.Ctx.Request.Form)


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
			response.Success = false
		}
	}()

	amode := models.AssetsApplyModel{};
	response.Content,err =  amode.AssetsAllot_MainQuery(ctl.Controller);

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="查询成功！"
		response.Success = true
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="查询失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//调拨 从单查询
func (ctl *AssetsApplyController) AssetsAllot_DetailQuery(){
	response := utils.ResponsModel{};
	var err error
	//fmt.Println(ctl.Ctx.Request.RequestURI)
	//fmt.Println(ctl.Ctx.Request.URL.Query())//--这个有用
	//fmt.Println(ctl.Ctx.Input.URL())
	//fmt.Println(ctl.Ctx.Request.Form)


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
			response.Success = false
		}
	}()

	amode := models.AssetsApplyModel{};
	response.Content,err =  amode.AssetsAllot_DetailQuery(ctl.Controller);

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="查询成功！"
		response.Success = true
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="查询失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//调拨 新增
func (ctl *AssetsApplyController) AssetsAllotAdd(){
	response := utils.ResponsModel{};


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode
			response.Message = err.(string)
			response.Success = false

			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};


	aid,err := amode.AssetsAllotAdd(ctl.Controller)

	if err ==nil{
		rdata := make(map[string]interface{})
		rdata["id"] = aid
		response.Code = utils.SuccessCode
		response.Message = "新增成功"
		response.Success = true
		response.Content.Result = append(response.Content.Result, rdata)
	} else {
		response.Code = utils.FailedCode
		response.Message = "新增出错:"+err.Error()
		response.Success = false
	}


	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//调拨 修改
func (ctl *AssetsApplyController) AssetsAllotUpdate(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode
			response.Message = err.(string)

			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	_,err := amode.AssetsAllotUpdate(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode
		response.Message = "修改成功！";
		response.Success = true
	} else {
		response.Code = utils.FailedCode
		response.Message = "修改失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//调拨 删除
func (ctl *AssetsApplyController) AssetsAllotDel(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode;
			response.Message = err.(string)
			response.Success = false
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	err := amode.AssetsAllotDel(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode;
		response.Message = "删除成功!"
		response.Success = true
	} else {
		response.Code = utils.FailedCode;
		response.Message = "删除失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//调拨 审批
func (ctl *AssetsApplyController) AssetsAllot_Audit(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode;
			response.Message = err.(string)
			response.Success = false
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	err := amode.AssetsAllot_Audit(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode;
		response.Message = "审批成功!"
		response.Success = true
	} else {
		response.Code = utils.FailedCode;
		response.Message = "审批失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//调拨 调出
func (ctl *AssetsApplyController) AssetsAllot_StockOut(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode;
			response.Message = err.(string)
			response.Success = false
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	err := amode.AssetsAllot_StockOut(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode;
		response.Message = "调出成功!"
		response.Success = true
	} else {
		response.Code = utils.FailedCode;
		response.Message = "调出失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//调拨 调入
func (ctl *AssetsApplyController) AssetsAllot_StockIn(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode;
			response.Message = err.(string)
			response.Success = false
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	err := amode.AssetsAllot_StockIn(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode;
		response.Message = "调入成功!"
		response.Success = true
	} else {
		response.Code = utils.FailedCode;
		response.Message = "调入失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//调拨 辙消
func (ctl *AssetsApplyController) AssetsAllotSubmit_Back(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode;
			response.Message = err.(string)
			response.Success = false
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	err := amode.AssetsAllotSubmit_Back(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode;
		response.Message = "辙消成功!"
		response.Success = true
	} else {
		response.Code = utils.FailedCode;
		response.Message = "辙消失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}


//资产变更 主单查询
func (ctl *AssetsApplyController) AssetsChange_MainQuery(){
	response := utils.ResponsModel{};
	var err error
	//fmt.Println(ctl.Ctx.Request.RequestURI)
	//fmt.Println(ctl.Ctx.Request.URL.Query())//--这个有用
	//fmt.Println(ctl.Ctx.Input.URL())
	//fmt.Println(ctl.Ctx.Request.Form)


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
			response.Success = false
		}
	}()

	amode := models.AssetsApplyModel{};
	response.Content,err =  amode.AssetsChange_MainQuery(ctl.Controller);

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="查询成功！"
		response.Success = true
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="查询失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//资产变更 从单查询
func (ctl *AssetsApplyController) AssetsChange_DetailQuery(){
	response := utils.ResponsModel{};
	var err error
	//fmt.Println(ctl.Ctx.Request.RequestURI)
	//fmt.Println(ctl.Ctx.Request.URL.Query())//--这个有用
	//fmt.Println(ctl.Ctx.Input.URL())
	//fmt.Println(ctl.Ctx.Request.Form)


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
			response.Success = false
		}
	}()

	amode := models.AssetsApplyModel{};
	response.Content,err =  amode.AssetsChange_DetailQuery(ctl.Controller);

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="查询成功！"
		response.Success = true
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="查询失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//资产变更 新增
func (ctl *AssetsApplyController) AssetsChangeAdd(){
	response := utils.ResponsModel{};


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode
			response.Message = err.(string)
			response.Success = false

			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};


	aid,err := amode.AssetsChangeAdd(ctl.Controller)

	if err ==nil{
		rdata := make(map[string]interface{})
		rdata["id"] = aid
		response.Code = utils.SuccessCode
		response.Message = "新增成功"
		response.Success = true
		response.Content.Result = append(response.Content.Result, rdata)
	} else {
		response.Code = utils.FailedCode
		response.Message = "新增出错:"+err.Error()
		response.Success = false
	}


	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//资产变更 修改
func (ctl *AssetsApplyController) AssetsChangeUpdate(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode
			response.Message = err.(string)

			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	_,err := amode.AssetsChangeUpdate(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode
		response.Message = "修改成功！";
		response.Success = true
	} else {
		response.Code = utils.FailedCode
		response.Message = "修改失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//资产变更 删除
func (ctl *AssetsApplyController) AssetsChangeDel(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode;
			response.Message = err.(string)
			response.Success = false
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	err := amode.AssetsChangeDel(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode;
		response.Message = "删除成功!"
		response.Success = true
	} else {
		response.Code = utils.FailedCode;
		response.Message = "删除失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//资产变更 确认变更
func (ctl *AssetsApplyController) AssetsChangeSubmit(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode;
			response.Message = err.(string)
			response.Success = false
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	err := amode.AssetsChangeSubmit(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode;
		response.Message = "调整成功!"
		response.Success = true
	} else {
		response.Code = utils.FailedCode;
		response.Message = "调整失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}
//资产变更 辙消
func (ctl *AssetsApplyController) AssetsChangeSubmit_Back(){
	response := utils.ResponsModel{};

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode;
			response.Message = err.(string)
			response.Success = false
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()
	amode := models.AssetsApplyModel{};
	err := amode.AssetsChangeSubmit_Back(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode;
		response.Message = "辙消成功!"
		response.Success = true
	} else {
		response.Code = utils.FailedCode;
		response.Message = "辙消失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();
}
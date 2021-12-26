package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"Eam_Server/utils"
	"Eam_Server/models"
)

type ReportController struct {
	beego.Controller
}

//类别分析表
func (ctl *ReportController) Report_As_Type_Browse(){
	response := utils.ResponsModel{};
	var err error
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
	amode := models.ReportModel{};
	response.Content,err = amode.Report_As_Type_Browse(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode;
		response.Message = "查询成功!"
		response.Success = true

	} else {
		response.Code = utils.FailedCode;
		response.Message = "查询失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();

}

//部门分析表
func (ctl *ReportController) Report_As_DeptCount_Browse(){
	response := utils.ResponsModel{};
	var err error
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
	amode := models.ReportModel{};
	response.Content,err = amode.Report_As_DeptCount_Browse(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode;
		response.Message = "查询成功!"
		response.Success = true

	} else {
		response.Code = utils.FailedCode;
		response.Message = "查询失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();

}


//变更分析表
func (ctl *ReportController) Report_As_Change_Browse(){
	response := utils.ResponsModel{};
	var err error
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
	amode := models.ReportModel{};
	response.Content,err = amode.Report_As_Change_Browse(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode;
		response.Message = "查询成功!"
		response.Success = true

	} else {
		response.Code = utils.FailedCode;
		response.Message = "查询失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();

}

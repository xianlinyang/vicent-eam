package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"Eam_Server/utils"
	"Eam_Server/models"
)


type OptionsController struct {
	beego.Controller
}

//角色查询
func (ctl*OptionsController) SysLog_Query(){
	response := utils.ResponsModel{};
	var err error
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Success = false
			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()

	amode := models.OptionsModel{};
	response.Content,err =  amode.SysLog_GetList(ctl.Controller);

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
	fmt.Println("查询成功")
	//fmt.Println(response)
	ctl.Data["json"] = response
	ctl.ServeJSON();
}

//商品录入查询
func (ctl*OptionsController) Public_Input_Query(){
	response := utils.ResponsModel{};
	var err error
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Success = false
			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()

	amode := models.OptionsModel{};
	response.Content,err =  amode.Public_Input_Query(ctl.Controller);

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
	fmt.Println("查询成功")
	//fmt.Println(response)
	ctl.Data["json"] = response
	ctl.ServeJSON();
}

//商品录入人员查询
func (ctl*OptionsController) Public_Input_User_Query(){
	response := utils.ResponsModel{};
	var err error
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Success = false
			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()

	amode := models.OptionsModel{};
	response.Content,err =  amode.Public_Input_Query(ctl.Controller);

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
	fmt.Println("查询成功")
	//fmt.Println(response)
	ctl.Data["json"] = response
	ctl.ServeJSON();
}

//单据录入查询部门
func (ctl*OptionsController) Public_Input_Dept_Query(){
	response := utils.ResponsModel{};
	var err error
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Success = false
			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()

	amode := models.OptionsModel{};
	response.Content,err =  amode.Public_Input_Dept_Query(ctl.Controller);

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
	fmt.Println("查询成功")
	//fmt.Println(response)
	ctl.Data["json"] = response
	ctl.ServeJSON();
}

//单据录入--存放位置选择
func (ctl*OptionsController) Public_Input_Postion_Query(){
	response := utils.ResponsModel{};
	var err error
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Success = false
			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()

	amode := models.OptionsModel{};
	response.Content,err =  amode.Public_Input_Postion_Query(ctl.Controller);

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
	fmt.Println("查询成功")
	//fmt.Println(response)
	ctl.Data["json"] = response
	ctl.ServeJSON();
}

//单据录入--查询人员
func (ctl*OptionsController) Public_Input_People_Query(){
	response := utils.ResponsModel{};
	var err error
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Success = false
			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()

	amode := models.OptionsModel{};
	response.Content,err =  amode.Public_Input_People_Query(ctl.Controller);

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
	fmt.Println("查询成功")
	//fmt.Println(response)
	ctl.Data["json"] = response
	ctl.ServeJSON();
}

//单据录入--所有基础资料选择
func (ctl*OptionsController) Public_Input_Bill_BasicAll(){
	response := utils.ResponsModel{};
	var err error
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Success = false
			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()

	amode := models.OptionsModel{};
	response.Content,err =  amode.Public_Input_Bill_BasicAll(ctl.Controller);

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
	fmt.Println("查询成功")
	//fmt.Println(response)
	ctl.Data["json"] = response
	ctl.ServeJSON();
}

//单据录入--资产流水
func (ctl*OptionsController) Public_AssetsFlowing(){
	response := utils.ResponsModel{};
	var err error
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			response.Success = false
			response.Code = utils.FailedCode
			response.Message = err.(string)
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()

	amode := models.OptionsModel{};
	response.Content,err =  amode.Public_AssetsFlowing(ctl.Controller);

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
	fmt.Println("查询成功")
	//fmt.Println(response)
	ctl.Data["json"] = response
	ctl.ServeJSON();
}

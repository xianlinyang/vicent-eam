package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"Eam_Server/utils"
	"Eam_Server/models"
)

type RoleController struct {
	beego.Controller
}



//用户新增
func (ctl *RoleController) RoleAdd(){
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
	amode := models.RoleModel{};
	aid,err := amode.Role_Create(ctl.Controller)

	if err ==nil{
		rdata := make(map[string]interface{})
		rdata["id"] = aid
		response.Code = utils.SuccessCode
		response.Message = "添加成功"
		response.Content.Result = append(response.Content.Result, rdata)
		response.Success = true
	} else {
		response.Code = utils.FailedCode
		response.Message = "新增出错:"+err.Error()
		response.Success = false
	}


	ctl.Data["json"] = response
	ctl.ServeJSON();

}

//用户修改
func (ctl *RoleController) RoleUpdate(){

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
	amode := models.RoleModel{};
	_,err := amode.Role_Update(ctl.Controller)

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

//用户删除
func (ctl *RoleController) RoleDel(){
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
	amode := models.RoleModel{};
	_,err := amode.Role_Delete(ctl.Controller)

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

//角色查询
func (ctl*RoleController) RoleQuery(){
	response := utils.ResponsModel{};
	var err error

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

	//fmt.Println( ctl.GetString("svrport", "none"))

	amode := models.RoleModel{};
	response.Content,err =  amode.GetList(ctl.Controller);

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

func (ctl*RoleController) Qxbrowse(){
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

	amode := models.RoleModel{};
	response.Content,err =  amode.GetQxList(ctl.Controller);

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

func (ctl*RoleController) RoleDeptbrowse(){
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

	amode := models.RoleModel{};
	response.Content,err =  amode.Get_RoleDeptList(ctl.Controller);

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
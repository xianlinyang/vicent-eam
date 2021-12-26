package controllers

import (
	"Eam_Server/models"
	"Eam_Server/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type BasicController struct {
	beego.Controller
}

type AA struct {
	a string
}
type B struct {
	AA
}
//部门新增
func (ctl *BasicController) DeptnoAdd(){
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
	amode := models.BasicInfoModel{};

	//var a = beego.Controller{}
	//var b = BasicController{}
	//a  = b.Controller

	aid,err := amode.Deptno_Create(ctl.Controller)

	if err ==nil{
		rdata := make(map[string]interface{})
		rdata["id"] = aid
		response.Code = utils.SuccessCode
		response.Message = "添加成功"
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

//部门修改
func (ctl *BasicController) DeptnoUpdate(){

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
	amode := models.BasicInfoModel{};
	_,err := amode.Deptno_Update(ctl.Controller)

	if err == nil {
		response.Code = utils.SuccessCode
		response.Message = "修改成功！";
		response.Success = true
	} else {
		response.Code = utils.FailedCode
		response.Message = "修改失败！";
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();

}

//部门删除
func (ctl *BasicController) DeptnoDel(){
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
	amode := models.BasicInfoModel{};
	_,err := amode.Deptno_Delete(ctl.Controller)

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

//部门树查询
func (ctl*BasicController) DeptnoTrelistQuery(){
	response := utils.ResponsModel{};
	var err error
	fmt.Println(ctl.Ctx.Request.RequestURI)
	fmt.Println(ctl.Ctx.Request.URL.Query())//--这个有用
	fmt.Println(ctl.Ctx.Input.URL())
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

	amode := models.BasicInfoModel{};
	response.Content,err =  amode.Deptno_GetTreeList(ctl.Controller);

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

//部门查询
func (ctl*BasicController) DeptnoQuery(){
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

	amode := models.BasicInfoModel{};
	response.Content,err =  amode.Deptno_GetList(ctl.Controller);

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

//启用停用
func (ctl *BasicController) DeptnoSetStatus(){
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



	amode := models.BasicInfoModel{};
	err =  amode.DeptnoSetStatus(ctl.Controller);

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="设置成功！"
		response.Success = true
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="设置失败,"+err.Error()
		response.Success = false
	}
	if utils.ISLog {
		fmt.Println("设置成功")
	}

	//fmt.Println(response)
	ctl.Data["json"] = response
	ctl.ServeJSON();

}


////资产类型
func (ctl *BasicController) AssetsTypeTreeList(){
	response := utils.ResponsModel{};
	var err error

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容


		}
	}()

	amode := models.BasicInfoModel{};
	response.Content,err =  amode.Basic_AssetsType_GetTreeList(ctl.Controller,"1");

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
func (ctl *BasicController) AssetsTypeAdd(){
	Basic_Add(ctl,"1")
}
func (ctl *BasicController) AssetsTypeUpdate(){
	Basic_Update(ctl,"1")
}
func (ctl *BasicController) AssetsTypeDel(){
	Basic_Del(ctl,"1")
}
func (ctl *BasicController) AssetsTypeQuery(){
	Basic_Query(ctl,"1")
}

//供应商类型
func (ctl *BasicController)SupplierAdd_type(){
	Basic_Add(ctl,"2")
}
func (ctl *BasicController) SupplierUpdate_type(){
	Basic_Update(ctl,"2")
}
func (ctl *BasicController) SupplierDel_type(){
	Basic_Del(ctl,"2")
}
func (ctl *BasicController) SupplierQuery_type(){
	Basic_Query(ctl,"2")
}
func (ctl *BasicController) SupplierTreeList_type(){
	response := utils.ResponsModel{};
	var err error

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容


		}
	}()

	amode := models.BasicInfoModel{};
	response.Content,err =  amode.Basic_AssetsType_GetTreeList(ctl.Controller,"2");

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

//供应商
func (ctl *BasicController)SupplierAdd(){
	Basic_Add(ctl,"3")
}
func (ctl *BasicController) SupplierUpdate(){
	Basic_Update(ctl,"3")
}
func (ctl *BasicController) SupplierDel(){
	Basic_Del(ctl,"3")
}
func (ctl *BasicController) SupplierQuery(){
	Basic_Query(ctl,"3")
}

//存放地点
func (ctl *BasicController) StoreAddressAdd(){
	Basic_Add(ctl,"4")
}
func (ctl *BasicController) StoreAddressUpdate(){
	Basic_Update(ctl,"4")
}
func (ctl *BasicController) StoreAddressDel(){
	Basic_Del(ctl,"4")
}
func (ctl *BasicController) StoreAddressQuery(){
	Basic_Query(ctl,"4")
}
func (ctl *BasicController) StoreAddressTreeList(){
	response := utils.ResponsModel{};
	var err error

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容


		}
	}()

	amode := models.BasicInfoModel{};
	response.Content,err =  amode.Basic_AssetsType_GetTreeList(ctl.Controller,"4");

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

//基础资料通用查询
func Basic_Query(ctl *BasicController,atype string) {

	response := utils.ResponsModel{};
	var err error

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容


		}
	}()

	amode := models.BasicInfoModel{};
	response.Content,err =  amode.Basic_GetList(ctl.Controller,atype);

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

//基础资料通用新增
func Basic_Add(ctl *BasicController,atype string) {
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
	amode := models.BasicInfoModel{};


	aid,err := amode.Basic_Create(ctl.Controller,atype)

	if err ==nil{
		rdata := make(map[string]interface{})
		rdata["id"] = aid
		response.Code = utils.SuccessCode
		response.Message = "添加成功"
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

//基础资料通用修改
func Basic_Update(ctl *BasicController,atype string) {
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
	amode := models.BasicInfoModel{};


	aid,err := amode.Basic_Update(ctl.Controller,atype)

	if err ==nil{
		rdata := make(map[string]interface{})
		rdata["id"] = aid
		response.Code = utils.SuccessCode
		response.Message = "修改成功"
		response.Success = true
		response.Content.Result = append(response.Content.Result, rdata)
	} else {
		response.Code = utils.FailedCode
		response.Message = "修改出错:"+err.Error()
		response.Success = false
	}


	ctl.Data["json"] = response
	ctl.ServeJSON();

}

//基础资料通用删除
func Basic_Del(ctl *BasicController,atype string){
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
	amode := models.BasicInfoModel{};
	err := amode.Basic_Del(ctl.Controller,atype)

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

//基础资料停用/启用
  /// astatus 0停用 1启用
func Basic_StartOrStop(ctl *BasicController,atype string ){
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
	amode := models.BasicInfoModel{};
	err := amode.Basic_StartOrStop(ctl.Controller,atype)

	if err == nil {
		response.Code = utils.SuccessCode;
		response.Message = "操作成功!"
		response.Success = true
	} else {
		response.Code = utils.FailedCode;
		response.Message = "操作失败,"+err.Error();
		response.Success = false
	}

	ctl.Data["json"] = response
	ctl.ServeJSON();

}


//商品资料查询
func (ctl *BasicController) AssetsListQuery(){
	response := utils.ResponsModel{};
	var err error

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容


		}
	}()

	amode := models.BasicInfoModel{};
	response.Content,err =  amode.AssetsList_GetList(ctl.Controller);

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="查询成功！"
		response.Success = true

	} else {
		response.Code = utils.FailedCode
		response.Message ="查询失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}
func (ctl *BasicController) AssetsListAdd(){
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
	amode := models.BasicInfoModel{};


	aid,err := amode.AssetsList_Create(ctl.Controller)

	if err ==nil{
		rdata := make(map[string]interface{})
		rdata["id"] = aid
		response.Code = utils.SuccessCode
		response.Message = "添加成功"
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
func (ctl *BasicController) AssetsListUpdate(){
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
	amode := models.BasicInfoModel{};

	err := amode.AssetsList_Update(ctl.Controller)

	if err ==nil{
		rdata := make(map[string]interface{})

		response.Code = utils.SuccessCode
		response.Message = "修改成功"
		response.Success = true
		response.Content.Result = append(response.Content.Result, rdata)
	} else {
		response.Code = utils.FailedCode
		response.Message = "修改出错:"+err.Error()
		response.Success = false
	}


	ctl.Data["json"] = response
	ctl.ServeJSON();
}
func (ctl *BasicController) AssetsListDel(){
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
	amode := models.BasicInfoModel{};
	err := amode.AssetsList_Del(ctl.Controller)

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
func (ctl *BasicController) AssetsListQuery_AllList(){
	response := utils.ResponsModel{};
	var err error


	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容


		}
	}()

	amode := models.BasicInfoModel{};
	//资产类型
	AssetsTypeList,err :=  amode.Basic_AssetsType_GetTreeList(ctl.Controller,"1");
	//供应商树
	SupplierList,err :=  amode.Supplier_GetTreeList(ctl.Controller);
	//存放地点树
	StoreList,err :=  amode.Basic_AssetsType_GetTreeList(ctl.Controller,"4");
	//部门树
	DeptList,err :=  amode.Deptno_GetTreeList(ctl.Controller);
	//单位
	SPUnit,err :=amode.SPUnitQuery(ctl.Controller)
	//用户信息
	amodeuser := models.UserModel{};
	UserList,err :=  amodeuser.GetList(ctl.Controller);

	rdata := make( map[string]interface{})
	rdata["AssetsTypeList"] = AssetsTypeList.Result
	rdata["SupplierList"] = SupplierList.Result
	rdata["StoreList"] = StoreList.Result
	rdata["DeptList"] = DeptList.Result
	rdata["UserList"] = UserList.Result
	rdata["SPUnit"] = SPUnit.Result

	response.Content.Result = append(response.Content.Result, rdata)
	response.Content.CurPage =1
	response.Content.PageSize = 0

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="查询成功！"
		response.Success = true

	} else {
		response.Code = utils.FailedCode
		response.Message ="查询失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}

func (ctl *BasicController) AssetsImport(){
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
	amode := models.BasicInfoModel{};


	response.Content,err = amode.AssetsImport(ctl.Controller)

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message = "导入成功"
		response.Success = true
		//response.Content.Result = append(response.Content.Result, rdata)
	} else {
		response.Code = utils.FailedCode
		response.Message = "导入失败:"+err.Error()
		response.Success = false
	}


	ctl.Data["json"] = response
	ctl.ServeJSON();
}


//许可信息查询
func (ctl *BasicController) LicenseQuery(){
	response := utils.ResponsModel{};
	var err error

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容


		}
	}()

	amode := models.BasicInfoModel{};
	response.Content,err =  amode.License_GetList(ctl.Controller);

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="查询成功！"
		response.Success = true

	} else {
		response.Code = utils.FailedCode
		response.Message ="查询失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}


//数据字典类型
func (ctl *BasicController) Dictionary_Type_GetTreeList(){
	response := utils.ResponsModel{};
	var err error

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容


		}
	}()

	amode := models.BasicInfoModel{};
	response.Content,err =  amode.Dictionary_Type_GetTreeList(ctl.Controller);

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="查询成功！"
		response.Success = true

	} else {
		response.Code = utils.FailedCode
		response.Message ="查询失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}
func (ctl *BasicController) Dictionary_TypeAdd(){
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
	amode := models.BasicInfoModel{};


	aid,err := amode.Dictionary_Type_Create(ctl.Controller)

	if err ==nil{
		rdata := make(map[string]interface{})
		rdata["id"] = aid
		response.Code = utils.SuccessCode
		response.Message = "添加成功"
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
func (ctl *BasicController) Dictionary_TypeUpdate(){
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
	amode := models.BasicInfoModel{};

	err := amode.Dictionary_Type_Update(ctl.Controller)

	if err ==nil{
		rdata := make(map[string]interface{})

		response.Code = utils.SuccessCode
		response.Message = "修改成功"
		response.Success = true
		response.Content.Result = append(response.Content.Result, rdata)
	} else {
		response.Code = utils.FailedCode
		response.Message = "修改出错:"+err.Error()
		response.Success = false
	}


	ctl.Data["json"] = response
	ctl.ServeJSON();
}
func (ctl *BasicController) Dictionary_TypeDel(){
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
	amode := models.BasicInfoModel{};
	err := amode.Dictionary_Type_Del(ctl.Controller)

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

//数据字典
func (ctl *BasicController) DictionaryQuery(){
	response := utils.ResponsModel{};
	var err error

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容


		}
	}()

	amode := models.BasicInfoModel{};
	response.Content,err =  amode.Dictionary_GetList(ctl.Controller);

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="查询成功！"
		response.Success = true

	} else {
		response.Code = utils.FailedCode
		response.Message ="查询失败,"+err.Error()
		response.Success = false
	}
	ctl.Data["json"] = response
	ctl.ServeJSON();
}
func (ctl *BasicController) DictionaryAdd(){
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
	amode := models.BasicInfoModel{};


	aid,err := amode.Dictionary_Create(ctl.Controller)

	if err ==nil{
		rdata := make(map[string]interface{})
		rdata["id"] = aid
		response.Code = utils.SuccessCode
		response.Message = "添加成功"
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
func (ctl *BasicController) DictionaryUpdate(){
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
	amode := models.BasicInfoModel{};

	err := amode.Dictionary_Update(ctl.Controller)

	if err ==nil{
		rdata := make(map[string]interface{})

		response.Code = utils.SuccessCode
		response.Message = "修改成功"
		response.Success = true
		response.Content.Result = append(response.Content.Result, rdata)
	} else {
		response.Code = utils.FailedCode
		response.Message = "修改出错:"+err.Error()
		response.Success = false
	}


	ctl.Data["json"] = response
	ctl.ServeJSON();
}
func (ctl *BasicController) DictionaryDel(){
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
	amode := models.BasicInfoModel{};
	err := amode.Dictionary_Del(ctl.Controller)

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

//外部联系人
func (ctl *BasicController)ContactOuter_Browse(){
	Basic_Query(ctl,"5")
}
func (ctl *BasicController) ContactOuter_Create(){
	Basic_Add(ctl,"5")
}
func (ctl *BasicController) ContactOuter_Update(){
	Basic_Update(ctl,"5")
}
func (ctl *BasicController) ContactOuter_Delete(){
	Basic_Del(ctl,"5")
}
func (ctl *BasicController) ContactOuter_Start(){
	Basic_StartOrStop(ctl,"5")
}
func (ctl *BasicController) ContactOuter_Stop(){
	Basic_StartOrStop(ctl,"5")
}
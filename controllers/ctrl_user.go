package controllers

import (
	"Eam_Server/models"
	"Eam_Server/public"
	"Eam_Server/services"
	"Eam_Server/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type UserController struct {
	beego.Controller
}

//用户新增
func (ctl *UserController) UserAdd(){
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
	amode := models.UserModel{};
	aid,err := amode.User_Create(ctl.Controller)

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

//用户修改
func (ctl *UserController) UserUpdate(){

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
	amode := models.UserModel{};
	_,err := amode.User_Update(ctl.Controller)

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

	fmt.Println("调用成功，返回：",response)

}

//用户删除
func (ctl *UserController) UserDel(){
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
	amode := models.UserModel{};
	_,err := amode.User_Delete(ctl.Controller)

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
//用户查询
func (ctl *UserController) UserQuery(){
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



	amode := models.UserModel{};
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
	fmt.Println("查询成功")
	//fmt.Println(response)
	ctl.Data["json"] = response
	ctl.ServeJSON();

}

//用户登陆
func (ctl *UserController) UserLogin() {
	fmt.Println("进入")
	var req map[string]interface{};

	//var usrmap map[string]string;
	var err error;
	var sid string
	var conn = services.NewConnection(utils.GetAppid(ctl.Controller))
	err = json.Unmarshal(ctl.Ctx.Input.RequestBody, &req);
	response := utils.ResponsModel{};


	//获取到的req 转为json
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("登陆失败：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode

			response.Message =  "登陆失败,"+err.(string)//err.(error).Error()
			response.Success = false
			utils.LogOut("info","用户登陆出错",response)
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()


	fmt.Println(req)
	if err == nil {
		res, err :=  services.FindToList(utils.GetAppid(ctl.Controller),conn,"select  a.*,b.f_name as f_deptname from tbl_user a "+
			" left join tbl_dept b on a.f_deptno = b.id where a.f_no =? and a.f_customid = ? ", req["userno"],utils.GetAppid(ctl.Controller));
		if (err != nil){
			panic(err.Error())
		}

		if len(res) ==0{  // 如果用户编号没找到再用手机号找一下
			res, err =  services.FindToList(utils.GetAppid(ctl.Controller),conn,"select  a.*,b.f_name as f_deptname from tbl_user a "+
				" left join tbl_dept b on a.f_deptno = b.id where a.f_no =? and a.f_customid = ? ", req["f_phone"],utils.GetAppid(ctl.Controller));
		}
		if (len(res)>0) {
			if res[0]["f_password"] != req["password"] {
				panic("密码错误！")
			}

			if (res[0]["f_status"].(int64) != 1) {
				panic("当前账号已被停用！")
			}

		} else {
			panic("用户不存在！")
		}


		usrmap := make(map[string]string, 7)
		//for k, v := range res[0] {
		//	usrmap[k] = fmt.Sprint(v)
		//}
		usrmap["id"] = fmt.Sprint(res[0]["id"] )
		usrmap["f_no"] = fmt.Sprint(res[0]["f_no"] )
		usrmap["f_name"] = fmt.Sprint(res[0]["f_name"] )
		usrmap["f_password"] = fmt.Sprint(res[0]["f_password"] )
		usrmap["f_roleid"] = fmt.Sprint(res[0]["f_roleid"] )
		usrmap["f_status"] = fmt.Sprint(res[0]["f_status"] )
		usrmap["f_flag"] = fmt.Sprint(res[0]["f_flag"] )


		//usr,_ := json.Marshal(res[0])
		//_ = json.Unmarshal(usr,&usrmap);
		//fmt.Println(usrmap)
		svrdate := time.Now();
		usrmap["expire_time"] = svrdate.Format("2006-01-02")
		token := make( map[string]string)
		stoken,err := public.CreateToken(usrmap)
		if err != nil{
			panic("Token生成出错,"+err.Error())
		}
		token["token"] = stoken

		token["expire_time"] =svrdate.Format("2006-01-02"); //Token当天有效
		err = public.SetToken(ctl.Controller,usrmap,token)

		if err != nil{
			panic(err.Error())
		}

		sid = strconv.FormatInt(res[0]["id"].(int64),10)   //strconv.Itoa(res[0]["id"].(int64))
		//获取最新Token
		stoken,err = services.Get_FieldByValue(utils.GetAppid(ctl.Controller),conn,"SELECT f_token FROM tbl_token WHERE f_userid = "+sid,"f_token")
		if err !=nil{
			panic("Token获取出错,"+err.Error())
		}
		//delete(res[0],"f_password") //密码不返回给前台

		remap := make(map[string]interface{})
		var  remaparry []map[string] interface{}


		remap["id"] = fmt.Sprint(res[0]["id"] )
		remap["f_no"] = fmt.Sprint(res[0]["f_no"] )
		remap["f_name"] = fmt.Sprint(res[0]["f_name"] )
		remap["f_deptid"] = fmt.Sprint(res[0]["f_deptno"] )
		remap["f_deptname"] = fmt.Sprint(res[0]["f_deptname"] )
		remaparry = append(remaparry,remap)

		response.Content.Result =remaparry
		response.Token = stoken
		response.Code = utils.SuccessCode
		response.Message = "登陆成功！";
		response.Success = true
	} else {
		panic("登陆失败，"+err.Error())
	}


	_= services.DB_Log(utils.GetAppid(ctl.Controller),conn,"用户登陆","用户登陆",sid,sid,"IP:"+ctl.Ctx.Request.RemoteAddr+" 登陆成功",err)
	ctl.Data["json"] = response
	ctl.ServeJSON();

}

//微信登陆
func (ctl *UserController)Wechat_UserCkAndLogin(){
	fmt.Println("进入")
	var req map[string]interface{};

	//var usrmap map[string]string;
	var err error;
	var sid string
	var conn = services.NewConnection(utils.GetAppid(ctl.Controller))
	err = json.Unmarshal(ctl.Ctx.Input.RequestBody, &req);
	response := utils.ResponsModel{};


	//获取到的req 转为json
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("登陆失败：", err) // 这里的err其实就是panic传入的内容
			response.Code = utils.FailedCode
			switch err.(type){
			case error:response.Message =  "登陆失败,"+err.(error).Error()
			default:
				response.Message = "登陆失败，"+err.(string)
			}
			//err.(error).Error()
			response.Success = false
			utils.LogOut("info","用户登陆出错",response)
			ctl.Data["json"] = response
			ctl.ServeJSON();
		}
	}()

	//fmt.Println(req)

	res, err :=  services.FindToList(utils.GetAppid(ctl.Controller),conn,"select  a.*,b.f_name as f_deptname from tbl_user a "+
		" left join tbl_dept b on a.f_deptno = b.id where a.f_flag = 0 and a.f_wxopenid =?", req["wxopenid"]);
	if (err != nil){
		panic(err.Error())
	}

	//先找微信OPENID,再找手机号，没有OPEID有手机号直接绑上
	if (len(res) <= 0) {
		res, err =  services.FindToList(utils.GetAppid(ctl.Controller),conn,"select  a.*,b.f_name as f_deptname from tbl_user a "+
			" left join tbl_dept b on a.f_deptno = b.id where a.f_flag = 0 and a.f_phone =? limit 0,1 ", req["movephone"]);
		if (len(res) <= 0) {
			response.Code = utils.Code_User_Chk
			response.Message =  "不存在用户信息，请先注册！"
			response.Success = false

			ctl.Data["json"] = response
			ctl.ServeJSON();
		} else {
			fmt.Println(res)
			_,err = services.Exec_Sql(utils.GetAppid(ctl.Controller),nil,conn," update tbl_user set f_wxopenid = ?,f_wxpicurl=? "+
				",f_UniID = ? where id = ?", req["wxopenid"].(string),req["picurl"].(string),req["nicid"].(string),res[0]["id"])
		}

	}

	usrmap := make(map[string]string, 7)
	//for k, v := range res[0] {
	//	usrmap[k] = fmt.Sprint(v)
	//}
	usrmap["id"] = fmt.Sprint(res[0]["id"] )
	usrmap["f_no"] = fmt.Sprint(res[0]["f_no"] )
	usrmap["f_name"] = fmt.Sprint(res[0]["f_name"] )
	usrmap["f_password"] = fmt.Sprint(res[0]["f_password"] )
	usrmap["f_roleid"] = fmt.Sprint(res[0]["f_roleid"] )
	usrmap["f_status"] = fmt.Sprint(res[0]["f_status"] )
	usrmap["f_flag"] = fmt.Sprint(res[0]["f_flag"] )


	//usr,_ := json.Marshal(res[0])
	//_ = json.Unmarshal(usr,&usrmap);
	//fmt.Println(usrmap)
	svrdate := time.Now();
	usrmap["expire_time"] = svrdate.Format("2006-01-02")
	token := make( map[string]string)
	stoken,err := public.CreateToken(usrmap)
	if err != nil{
		panic("Token生成出错,"+err.Error())
	}
	token["token"] = stoken

	token["expire_time"] =svrdate.Format("2006-01-02"); //Token当天有效
	err = public.SetToken(ctl.Controller,usrmap,token)

	if err != nil{
		panic(err.Error())
	}

	sid = strconv.FormatInt(res[0]["id"].(int64),10)   //strconv.Itoa(res[0]["id"].(int64))
	//获取最新Token
	stoken,err = services.Get_FieldByValue(utils.GetAppid(ctl.Controller),conn,"SELECT f_token FROM tbl_token WHERE f_userid = "+sid,"f_token")
	if err !=nil{
		panic("Token获取出错,"+err.Error())
	}
	//delete(res[0],"f_password") //密码不返回给前台

	remap := make(map[string]interface{})
	var  remaparry []map[string] interface{}


	remap["id"] = fmt.Sprint(res[0]["id"] )
	remap["f_no"] = fmt.Sprint(res[0]["f_no"] )
	remap["f_name"] = fmt.Sprint(res[0]["f_name"] )
	remap["f_deptid"] = fmt.Sprint(res[0]["f_deptno"] )
	remap["f_deptname"] = fmt.Sprint(res[0]["f_deptname"] )
	remaparry = append(remaparry,remap)

	response.Content.Result =remaparry
	response.Token = stoken
	response.Code = utils.SuccessCode
	response.Message = "登陆成功！";
	response.Success = true


	_= services.DB_Log(utils.GetAppid(ctl.Controller),conn,"微信用户登陆","用户登陆",sid,sid,"IP:"+ctl.Ctx.Request.RemoteAddr+" 登陆成功",err)
	ctl.Data["json"] = response
	ctl.ServeJSON();
}

//启用停用
func (ctl *UserController) SetStatus(){
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



	amode := models.UserModel{};
	err =  amode.SetStatus(ctl.Controller);

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="操作成功！"
		response.Success = true
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="操作失败,"+err.Error()
		response.Success = false
	}
	if utils.ISLog {
		fmt.Println("操作成功")
	}

	//fmt.Println(response)
	ctl.Data["json"] = response
	ctl.ServeJSON();

}

//修改密码
func (ctl *UserController) SetPwd(){
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



	amode := models.UserModel{};
	err =  amode.SetPwd(ctl.Controller);

	if err ==nil{
		response.Code = utils.SuccessCode
		response.Message ="修改成功！"
		response.Success = true
		//fmt.Println(response.Content)
	} else {
		response.Code = utils.FailedCode
		response.Message ="修改失败,"+err.Error()
		response.Success = false
	}
	if utils.ISLog{
		fmt.Println("查询成功")
	}
	//fmt.Println(response)
	ctl.Data["json"] = response
	ctl.ServeJSON();

}

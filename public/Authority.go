package public

import (
	"Eam_Server/services"
	"Eam_Server/utils"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"reflect"
	"strings"
	"time"
)

//JWT加密
func CreateToken(user map[string]string)(tokenss string,err error){
	//var usr map[string]interface{}
	//自定义claim
	user = user
	//claim := jwt.MapClaims(usr)
	claim := jwt.MapClaims{
		"id":      	 	user["id"],
		"f_no":     	user["f_no"],
		"f_name":	 	user["f_name"],
		"f_password": 	user["f_password"],
		"f_roleid": 	user["f_roleid"],
		"f_status": 	user["f_status"],
		"f_flag": 		user["f_flag"],
		"expire_time": 	user["expire_time"],
		//"nbf":      	"",//time.Now().Unix(),
		//"iat":      	"",//time.Now().Unix(),
	}
	//id,f_no,f_name,f_password,f_roleid,f_status,f_flag
	fmt.Println("加密前",claim)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)

	tokenss,err  = token.SignedString([]byte(utils.TokenSecret))
	//fmt.Println("token加客",tokenss)
	return
}

func secret()jwt.Keyfunc{
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.TokenSecret),nil
	}
}

//JWT解密
func ParseToken(tokenss string)(user map[string]string,err error){

	token,err := jwt.Parse(tokenss,secret())
	//fmt.Println(token)
	//fmt.Println("TOKEN",tokenss)
	if err != nil{
		return
	}
	claim,ok := token.Claims.(jwt.MapClaims)

	if !ok{
		err = errors.New("cannot convert claim to mapclaim")
		return
	}
	//验证token，如果token被修改过则为false
	if  !token.Valid{
		err = errors.New("token is invalid")
		return
	}
	//claim_json := json.Marshal([]byte(claim))
	ret := make(map[string]string, len(claim))
	for k, v := range claim {
		ret[k] = fmt.Sprint(v)
	}
	user = ret

	//fmt.Println("解析MAP",user)

	return
}

//登陆时向数据库插入TOKEN
func SetToken(req beego.Controller,usermap,tokenmap map[string]string)(aerr error){

	//先查找这个用户TOKEN表中有没有TOKEN
	var conn = services.NewConnection(utils.GetAppid(req))
	var sql string

	defer func() {
		if err := recover(); err != nil {
			//fmt.Println(err.(string))
			//fmt.Println(fmt.Errorf("%v", err))
			aerr =  fmt.Errorf("%v", err)
		}
		conn.Close()
	}()
	//usrjson,err := json.Marshal(usermap)

	//if err != nil{
	//	panic("用户信息解析出错，"+err.Error())
	//}
	userid := usermap["id"];

	now := time.Now()
	svrdate :=now.Format("2006-01-02 15:04:05");

    //fmt.Println("userid ="+userid)
   sdate,err := services.Get_FieldByValue(utils.GetAppid(req),conn,"SELECT  DATE_FORMAT(f_expire_time,'%Y-%m-%d')  as sdate FROM tbl_token WHERE f_customid = "+utils.GetAppid(req)+" and f_userid ="+userid,"sdate")


   if err != nil{
		panic("Token查询失败,"+err.Error())
	}
	fmt.Println(tokenmap)
   if sdate == ""{
		//不存在Token就直接插入一条新Token到表
		sql = "INSERT INTO tbl_token(f_customid,f_userid,f_token,f_create_time,f_expire_time)values(?,?,?,?,?) "
		_,err= conn.Exec(sql,utils.GetAppid(req),userid,tokenmap["token"],svrdate,tokenmap["expire_time"])
		if err != nil{
			panic("和成TOKEN出错,"+err.Error())
		}
	} else {
		//从Token中解析出用户信息
		token, err := ParseToken(tokenmap["token"])


		if err !=nil{
			panic("Token解析失败,"+err.Error())
		}
	    //delete(token,"expire_time")

		//str_token,err := json.Marshal(token)
		//如果解析出来的用户信息不相等就把新数据库中用户信息生成的Token重新赋值
		//if string(usrjson) != string(str_token)
	    if !(reflect.DeepEqual(token, usermap)) { //对比两个map是否相等
			sql = "UPDATE tbl_token SET f_token = ? WHERE f_userid = ? "
			_,err= conn.Exec(sql,tokenmap["token"],userid)
			if err != nil{
				panic("Token设置失败")
			}
		}else {
			//如果相等判断是否过期，过期了也要重新生成Token
			//var expire_time time.Time
			//expire_time,_ = time.Parse("2006-01-02", token["expire_time"])//token["expire_time"]
		   now := time.Now().Format("2006-01-02")

		   //if now ==token["expire_time"]{
		   //	fmt.Println("相等")
		   //}
			if now !=sdate {
				sql = "UPDATE tbl_token SET f_token = ?,f_expire_time=? WHERE f_userid = ? "
				_,err= conn.Exec(sql,tokenmap["token"],now,userid)
				if err != nil{
					panic("Token设置失败")
				}
			}
		}
	}
	return nil
}
//调用接口时检测TOKEN
func CheckToken(appid,atoken string) (result int,aerr string){
	var conn = services.NewConnection(appid)
	defer func() {
		if err := recover(); err != nil {
			//err = fmt.Errorf("%v", err)
			aerr = err.(string)//err.(error).Error()
			result = utils.Code_Chk_TokenValid
			return
		}
		conn.Close()
	}()

	if atoken ==""{
		panic("Token为空")
	}
	//fmt.Println("SELECT * FROM tbl_token WHERE f_customid ="+appid+" and  f_token ="+"\""+atoken+"\"")
	//先找数据库存Token是否存在
	tokenmap,err := services.Get_OneData(appid,conn,"SELECT * FROM tbl_token WHERE f_customid ="+appid+" and  f_token ="+"\""+atoken+"\"" )
	if err != nil{
		panic("Token获取出错")
	}

	if tokenmap == nil {
		panic("当前Token无访问权限")
	}

	//判断是否过期
	now := time.Now()
	svrdate :=now.Format("2006-01-02 00:00:00");

	if svrdate != tokenmap["f_expire_time"]{
		panic("token过期")
	}

	result = 1
	return result,""

}
//权限检测
func Check_Permissions(appid,api,stoken string)(result int,aerr string){
	var conn = services.NewConnection(appid)
	defer func() {
		if err := recover(); err != nil {
			//err = fmt.Errorf("%v", err)
			aerr = err.(string)//err.(error).Error()
			result = utils.Code_Chk_Permis
			return
		}
		conn.Close()
	}()

	var apichk string

	apichk = Get_Public_Permiss(api)

	//先解密出用户ID
	token, err := ParseToken(stoken)
	if err != nil{
		panic("用户解析出错,"+err.Error())
	}

	userid := token["id"]
	//超级管理员只检测有没有菜单项不检测权限
	if userid == "1"{
		//SELECT id FROM tbl_perms

		sid,err := services.Get_FieldByValue(appid,conn,"SELECT id FROM tbl_perms where   f_api ="+"\"" +apichk+"\"","id")
		if err !=nil{
			panic("权限检测出错,"+err.Error())
		}
		if sid ==""{
			panic("无操作权限！")
		}
	} else {
		ssql := "   SELECT a.id FROM tbl_role_data_perms a  "+
				" 	INNER JOIN tbl_role_user tru ON A.f_roleid = TRU.f_roleid "+
				" 	INNER JOIN tbl_user b ON TRU.f_userid = b.ID "+
				" 	INNER JOIN tbl_perms c ON a.f_qxid=c.f_qxid "+
				" 	WHERE b.f_customid = "+appid+" and b.id = "+userid+" AND c.f_api ="+"\"" +api+"\""

		sid,err := services.Get_FieldByValue(appid,conn,ssql,"id")
		if err !=nil{
			panic("权限检测出错,"+err.Error())
		}
		if sid ==""{
			panic("无操作权限！")
		}
	}

	return 1,""
}

//检查接口是否需要检测
func Api_Check(api string)bool{
	var result bool
	sapi := strings.ToLower(api)
	switch sapi {
		case "/api/user/userlogin":{result = false}
		case "/api/user/wechat/usercheck":{result = false}
		case "/api/user/create":{result = false}
		case "/api/user/update":{result = false}
		default:
			result = true
	}
	//if sapi !="/api/user/userlogin"{
	//	result = true
	//} else {
	//	result = false
	//}
	return result
}

//检查是否有公用权限接口
func Get_Public_Permiss(api string)(str string){
	var result string
	switch api {
		case "Supplier/create/type": { result  = "Supplier/create"}
		case "Supplier/update/type": { result  = "Supplier/update"}
		case "Supplier/delete/type": { result  = "Supplier/delete"}
		case "Dictionary/create/type": { result  = "Dictionary/create"}
		case "Dictionary/update/type": { result  = "Dictionary/update"}
		case "Dictionary/delete/type": { result  = "Dictionary/delete"}
		case "ReceiveApply/submit/back": { result  = "ReceiveApply/create"} //领用申请 辙消提交 只要可以新增就可以辙消提交
		case "ReceiveApply/ManagerApproval/back": { result  = "ReceiveApply/ManagerApproval"} //领用申请 主管审批
		case "ReceiveApply/AsseetApproval/back": { result  = "ReceiveApply/AsseetApproval"} //领用申请 仓库审批
		case "AssetReceive/submit/back": { result  = "AssetReceive/submit"} //领用提交
		case "AssetsBack/submit/back": { result  = "AssetsBack/submit"} //退库提交
		case "AssetsLend/submit/back": { result  = "AssetsLend/create"}
		case "AssetsLendIn/submit/back": { result  = "AssetsLendIn/create"}
		case "AssetsChange/apply/back": {result  = "AssetsChange/apply"}
		case "AssetsList/AssetsImport":{result ="AssetsList/create"}
	default:
		result = api
	}

	return result
}
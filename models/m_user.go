package models

import (
	"Eam_Server/services"
	"Eam_Server/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"time"
)

type UserModel struct {
	Base_Model
}


func (ctl *UserModel) User_Create(req beego.Controller)(aid string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))  ////通过客户信息获取数据库连接
	var roles string
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			aid = ""
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}

			return
		}
	}()

	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);
	fmt.Println("用户参数",reqs)

	if err !=nil{
		panic("参数转换出错："+err.Error())
	}
    //结构1 要检测的字段 结构2 要检测的赋加条件不传的话按默认字段加传入值判断，加了后按字段加AND
	icheck,serro := services.CheckRepeat(utils.GetAppid(req),conn,reqs,[]byte(`[{"FName":"f_no","PName":"编号"},{"FName":"f_name","PName":"名称"},{"FName":"f_phone","PName":"手机号码"}]`),"tbl_user",
		[]byte(`[{"FName":"f_customid","Ptype":"=","Chcktype":1,"Fvalue":"`+utils.GetAppid(req)+`"},{"FName":"f_flag","Ptype":"=","Chcktype":1,"Fvalue":"0"}]`))
	if icheck != 1{
		panic(serro)
	}

	now := time.Now()
	reqs["f_status"] ="1";
	reqs["f_type"] ="1";
	reqs["f_create_time"] =now.Format("2006-01-02 15:04:05");
	reqs["f_flag"] ="0";
	reqs["f_customid"] =utils.GetAppid(req)
	if (reqs["f_password"] ==nil){
		reqs["f_password"] =""
	}

	roles = reqs["rolelist"].(string)
	delete(reqs,"rolelist")
	tx,err := conn.Begin()
	defer func(){
		if err := recover();err != nil{
			_= tx.Rollback()
			panic(err)
		}
	}()

	aid,err = services.One_ADD_Table(utils.GetAppid(req),tx,conn,"tbl_user",reqs)

	if (err==nil){
		s :=strings.Split(roles,",")
		for _,v := range s {
			_,err  = services.Exec_Sql(utils.GetAppid(req),tx,conn,"INSERT INTO tbl_role_user(f_customid,f_userid,f_roleid,f_accountid) value(?,?,?,1)",utils.GetAppid(req),aid,v)
			if err != nil{
				break
			}
		}
	}

	if (err != nil)  {
		aid =""
		rerr = err
		_ = tx.Rollback()
		_= services.DB_Log(utils.GetAppid(req),conn,"用户管理","新增","",reqs["f_create_user"].(string),"新增失败",err)
	} else {
		_ = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"用户管理","新增",aid,reqs["f_create_user"].(string),"新增成功 ID ="+aid,err)
	}
	return
}

func (ctl *UserModel) User_Update(req beego.Controller)(alen string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var d_update = make(map[string]string);
	var conn = services.NewConnection(utils.GetAppid(req))
	var roles string

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			alen = ""
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
		}
	}()

	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	if err !=nil{
		panic("参数转换出错："+err.Error())
	}
	//结构1 要检测的字段 结构2 要检测的赋加条件不传的话按默认字段加传入值判断，加了后按字段加AND,checktype 0取结构值，1取传入值
	icheck,serro := services.CheckRepeat(utils.GetAppid(req),conn,reqs,[]byte(`[{"FName":"f_no","PName":"编号"},{"FName":"f_name","PName":"名称"},{"FName":"f_phone","PName":"手机号码"}]`),"tbl_user",
		[]byte(`[{"FName":"f_flag","Ptype":"=","Chcktype":1,"Fvalue":"0"},{"FName":"id","Ptype":"<>","Chcktype":1,"Fvalue":"`+reqs["id"].(string)+`"}]`))
	if icheck != 1{
		panic(serro)
	}
	roles = reqs["rolelist"].(string)
	delete(reqs,"rolelist")
	for key,value := range reqs {
		if key !="id"{
			if value != nil{
				d_update[key] = value.(string)
			}
		}
	}

	d_update["f_update_time"] = time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(d_update)

	sqlstr :=  services.Sql_JionSet(d_update)

	if sqlstr !="" {
		sqlstr = " update tbl_user set "+sqlstr +" where id ="+reqs["id"].(string)
	}

	tx,err := conn.Begin()

	defer func(){
		if err := recover();err!=nil{
			_= tx.Rollback()
			panic(err)
		}
	}()
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr,)

	if err ==nil{
		_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,"delete from tbl_role_user where f_customid = ? and f_userid = ?   ",utils.GetAppid(req),reqs["id"].(string))
	}

	if err ==nil{
		s := strings.Split(roles,",")
		for _,v := range s{
			_,err  = services.Exec_Sql(utils.GetAppid(req),tx,conn,"INSERT INTO tbl_role_user(f_customid,f_userid,f_roleid,f_accountid) value(?,?,?,1)",utils.GetAppid(req),reqs["id"].(string),v)
			if err != nil{
				break;
			}
		}
	}

	if err !=nil{
		alen =""
		rerr = err
		_ =tx.Rollback()
		_= services.DB_Log(utils.GetAppid(req),conn,"用户管理","修改",reqs["id"].(string),reqs["f_update_user"].(string),"修改失败",err)
	} else{
		alen ="0"
		rerr = nil
		sreq,_ := json.Marshal(reqs)
		slog := "修改成功，内容："+string(sreq)
		_= tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"用户管理","修改",reqs["id"].(string),reqs["f_update_user"].(string),slog,err)
	}
	return
}

func (ctl *UserModel) User_Delete(req beego.Controller)(alen string,rerr error){
	var reqs map[string]interface{};
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var conn = services.NewConnection(utils.GetAppid(req))
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			alen = ""
			rerr = fmt.Errorf("%s",err)
		}
	}()
	//fmt.Println(req.Ctx.Input.RequestBody)

	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	sqlstr := "update tbl_user  set f_flag = 1,f_delete_user = ?,f_delete_time = NOW() where id = ? "

	_,err = services.Exec_Sql(utils.GetAppid(req),nil,conn,sqlstr,reqs["f_delete_user"].(string),reqs["id"].(string))

	if err !=nil{
		alen =""
		rerr = err
		_= services.DB_Log(utils.GetAppid(req),conn,"用户管理","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除失败",err)
	} else{
		alen ="0"
		rerr = nil
		_= services.DB_Log(utils.GetAppid(req),conn,"用户管理","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除成功",err)
	}
	return
}

func (ctl *UserModel) GetList(req beego.Controller) (result utils.ResData, rerr error) {
	//var reqs map[string]interface{};
	var sortmap map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))
	var acount string
	var icount int
	var sqlstr_ord string
	//response := utils.ResponsModel{};

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("GetList 失败，",err)
			rerr = fmt.Errorf("%s",err)
		}
	}()

	var sqlstr = " SELECT a.*,b.f_name AS createname,c.f_name AS updatename,d.f_name AS delname,e.f_name as f_deptname "+
		" FROM tbl_user a "+
		" LEFT JOIN tbl_user b ON a.f_create_user = b.id "+
		" LEFT JOIN tbl_user c ON a.f_update_user = c.id "+
		" LEFT JOIN tbl_user d ON a.f_delete_user = d.id  "+
		" left join tbl_dept e on a.f_deptno = e.id "+
		" where a.f_customid = "+utils.GetAppid(req)+" and a.f_flag = 0 and a.f_type = 1 "

	deptno := req.GetString("deptno","")//req["f_deptno"]
	wvalue := req.GetString("w_value","")//req["w_value"]
	pageindex := req.GetString("curPage","")//req["pageindex"]curPage,pageSize
	pagesize := req.GetString("pageSize","")//req["pagesize"]
	sort := req.GetString("sort");
	//fmt.Println(pageindex)
	//fmt.Println(deptno)
	if pageindex ==""{
		pageindex ="1"
	}
	if pagesize ==""{
		pagesize ="100"
	}
	if deptno != "" {
		sqlstr += " AND a.f_deptno = \"" + deptno +"\""
	}
	if wvalue != ""{
		sqlstr += " and  (a.f_name like \"%s\" or a.f_no like \"%s\" or a.f_phone like \"%s\")"
		sqlstr = fmt.Sprintf(sqlstr,wvalue,wvalue,wvalue)
	}

	sqlstr_count := "select count(*) as icount from ("+sqlstr+") a "

	acount,err := services.Get_FieldByValue(utils.GetAppid(req),conn,sqlstr_count,"icount")
	if err != nil{
		panic("查询合计数出错，"+err.Error())
	}

	//sqlstr_ord := sqlstr+ " order by a.f_create_time desc "
	if sort != ""{
		_ = json.Unmarshal([]byte(sort),&sortmap)

		if (sortmap["sortprop"] !="") &&(sortmap["order"] !="") {
			fmt.Println(sortmap["order"])
			var ord string
			if sortmap["order"].(string) =="descending"{
				ord = " desc"
			} else {
				ord =" asc"
			}
			sqlstr_ord = " order by a."+sortmap["sortprop"].(string)+" "+ord
		}
	}

	icount,err = strconv.Atoi(acount)
	if err != nil{
		panic("获取合计数出错 ，"+err.Error())
	}
	result,err = services.FindToListByPage(utils.GetAppid(req),conn,sqlstr+sqlstr_ord,utils.InterToInt(pageindex) ,utils.InterToInt(pagesize) )

	if err == nil{
		for _,v:=range result.Result{
			sname, err:= services.Get_FieldByValue(utils.GetAppid(req),conn,"SELECT GROUP_CONCAT(b.f_name) AS f_name FROM tbl_role_user a INNER JOIN tbl_role b ON a.f_roleid = b.id WHERE a.f_userid = "+v["id"].(string),"f_name")
			if err ==nil{
				v["rolename"] = sname
			} else {
				v["rolename"] = ""
			}

			rolelist,err := services.FindToList(utils.GetAppid(req),conn,"SELECT f_roleid FROM tbl_role_user  WHERE f_userid = ? ",v["id"].(string))
			if err == nil{
				v["rolelist"] = rolelist
			}else {
				break
			}
		}
	}

	if err!=nil{
		panic("查询失败,"+err.Error())
	}else{
		result.PageSize = utils.InterToInt(pagesize)
		result.CurPage =utils.InterToInt(pageindex)
		result.Totals = icount
		if sort != ""{
			result.Sort = sortmap
		}
	}
	return result,nil
}

//启用，停用
func (ctl *UserModel) SetStatus(req beego.Controller) ( rerr error) {
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))


	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("SetStatus 失败，",err)
			rerr = fmt.Errorf("%s",err)
			return
		}
	}()

	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	var sqlstr = " UPDATE tbl_user set f_status = ?,f_update_user=?,f_update_time =? WHERE ID = ? "

	status := reqs["f_status"]
	updateuser := reqs["f_update_user"]
	f_update_time := time.Now().Format("2006-01-02 15:04:05")

	sid := reqs["id"]

 	_,err = services.Exec_Sql(utils.GetAppid(req),nil,conn,sqlstr,status,updateuser,f_update_time,sid)

	if err != nil{
		panic("修改出错，"+err.Error())
	}
 	var slog string
 	if status =="0"{
 		slog ="停用"
	}else {
		slog ="启用"
	}
	_= services.DB_Log(utils.GetAppid(req),conn,"用户管理",slog,reqs["id"].(string),reqs["f_update_user"].(string),slog,err)

	return nil
}

//修改密码
func (ctl *UserModel) SetPwd(req beego.Controller) (rerr error) {
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))
    var soldpwd string
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("SetPwd 失败，",err)
			rerr = fmt.Errorf("%s",err)

			return
		}
	}()

	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	if err !=nil{
		panic("参数转换出错："+err.Error())
	}
	if reqs["new_password"] ==nil{
		panic("新密码参数不存在")
	}

	var sqlstr = " UPDATE tbl_user set f_password = ?,f_update_user=?,f_update_time =? WHERE ID = ? "
	if reqs["old_password"] ==nil{
		soldpwd =""
	} else {
		soldpwd = reqs["old_password"].(string)
	}


	spwd := reqs["new_password"]
	updateuser := reqs["f_update_user"]
	sid := reqs["id"]
	f_update_time := time.Now().Format("2006-01-02 15:04:05")

	rdata,err := services.Get_OneData(utils.GetAppid(req),conn,"select * from tbl_user where id ="+sid.(string))
	if err != nil{
		panic("用户信息获取出错："+err.Error())
	}

	if soldpwd != rdata["f_password"].(string){
		panic("原密码与录入不一至！")
	}

	_,err = services.Exec_Sql(utils.GetAppid(req),nil,conn,sqlstr,spwd,updateuser,f_update_time,sid)

	if err != nil{
		panic("修改出错，"+err.Error())
	}

	_= services.DB_Log(utils.GetAppid(req),conn,"用户管理","修改密码",reqs["id"].(string),reqs["f_update_user"].(string),"修改密码成功",err)

	return nil
}
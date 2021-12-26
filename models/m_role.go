package models

import (
	"Eam_Server/services"
	"Eam_Server/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	"strconv"
	"time"
)

type RoleModel struct {
	Base_Model
}


type Menu_3 struct {
	Qxid  		int
	Name  		string
	Pid 		int
	Achk     	int
	Visible     int
}
type Menu_2 struct {
	Qxid  		int
	Name  		string
	Pid 		int
	Achk     	int
	Visible     int
	Subs        [] Menu_3
}
//权限菜单返回结构
type MyJsonMemu struct {
	Qxid  		int
	Name  		string
	Pid     	int
	Achk     	int
	Visible     int
	Subs []     Menu_2
}

func (ctl *RoleModel) Role_Create(req beego.Controller)(aid string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接
	var s string
	fmt.Println("ss")
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			aid = ""
			rerr = fmt.Errorf("%s",err)
			return
		}
	}()
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	json_str,err:= json.Marshal(reqs)
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}
	rjson, err := simplejson.NewJson([]byte(json_str));
	//开始事务
	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	role := make(map[string] interface{});

	role["f_name"] = rjson.Get("f_name").MustString()
	role["f_ordid"] = rjson.Get("f_ordid").MustString()
	role["f_note"] = rjson.Get("f_note").MustString()
	role["f_create_user"] = rjson.Get("f_create_user").MustString()
	role["f_customid"] = utils.GetAppid(req)
	role["f_DeptType"] = rjson.Get("f_DeptType").MustString()
	if role["f_DeptType"] ==""{
		role["f_DeptType"] = "1"
	}
	now := time.Now()
	role["f_create_time"]= now.Format("2006-01-02 15:04:05");


	s,err = services.Get_FieldByValue(utils.GetAppid(req),conn,"select id from tbl_role where f_customid ="+utils.GetAppid(req)+" and f_name = \""+role["f_name"].(string)+"\"","id")
	if err ==nil{

		if s != ""{
			panic("角色【"+role["f_name"].(string)+"】已存在，不能重复添加！")
		}
	} else{
		panic("角色检测出错,"+err.Error())
	}
	//插入角色
	aid,err = services.One_ADD_Table(utils.GetAppid(req),tx,conn,"tbl_role",role)
	//fmt.Println(role)
	//插入权限
	if err ==nil{
		rows, _ := rjson.Get("qxlist").Array();
		qxsql :="INSERT INTO tbl_role_data_perms(f_roleid,f_qxid,f_qx_pid,f_customid) SELECT ?,?,f_pid,"+utils.GetAppid(req)+" FROM view_Menu WHERE f_qxid = ?"
		for _, row := range rows {
			if each_map, ok := row.(map[string] interface{}); ok {
				_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,qxsql,aid,each_map["f_qxid"].(string),each_map["f_qxid"].(string))
				if err !=nil{
					break;
				}
			}
		}
	}

	//插入角色可操作部门  -- f_DeptType=1时表示可以操作所有部门
	if (err==nil)&&(role["f_DeptType"]=="2"){
		rows_dept, _ := rjson.Get("deptlist").Array();
		deptsql :="INSERT INTO tbl_role_deptid(f_customid,f_roleid,f_deptid,f_create_time) VALUES(?,?,?,now())"
		for _, row := range rows_dept {
			if each_map, ok := row.(map[string] interface{}); ok {
				_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,deptsql,utils.GetAppid(req),aid,each_map["f_deptid"].(string))
				if err !=nil{
					break;
				}
			}
		}
	}

	if (err != nil)  {
		aid =""
		rerr = err
		err = tx.Rollback()
		if err == nil{
			fmt.Println("回滚完成")
		} else {
			fmt.Println(err)
		}
		_= services.DB_Log(utils.GetAppid(req),conn,"角色管理","新增","",reqs["f_create_user"].(string),"新增失败",err)
	} else {
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"角色管理","新增",reqs["f_create_user"].(string),reqs["f_create_user"].(string),"新增成功 ID ="+aid,err)
	}
	return

}

func (ctl *RoleModel) Role_Update(req beego.Controller)(alen string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var d_update = make(map[string]string);
	var conn = services.NewConnection(utils.GetAppid(req))

	defer func() {

		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			alen = ""
			rerr = fmt.Errorf("%s",err)
		}
	}()
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	if err !=nil{
		panic("参数转换出错："+err.Error())
	}
	json_str,err:= json.Marshal(reqs)
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}
	rjson, err := simplejson.NewJson([]byte(json_str));

	d_update["id"] = rjson.Get("id").MustString()
	d_update["f_name"] = rjson.Get("f_name").MustString()
	d_update["f_ordid"] = rjson.Get("f_ordid").MustString()
	d_update["f_note"] = rjson.Get("f_note").MustString()
	d_update["f_update_user"] = rjson.Get("f_update_user").MustString()
	d_update["f_update_time"] = time.Now().Format("2006-01-02 15:04:05")
	d_update["f_DeptType"] = rjson.Get("f_DeptType").MustString()

	//开始事务
	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()

	sqlstr :=  services.Sql_JionSet(d_update)
	if sqlstr !="" {
		sqlstr = " update tbl_role set "+sqlstr +" where id ="+d_update["id"]
	}

    //修改角色
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)

	//if err != nil{
	//	tx.Rollback();
	//	panic("修改失败，"+err.Error())
	//}
	//删除权限
	if err == nil{
		sqlstr = " DELETE from tbl_role_data_perms WHERE f_roleid ="+reqs["id"].(string)
		_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr,)
	}
	//重新插入权限
	if err ==nil{
		rows, _ := rjson.Get("qxlist").Array();
		qxsql :="INSERT INTO tbl_role_data_perms(f_roleid,f_qxid,f_qx_pid,f_customid) SELECT ?,?,f_pid,"+utils.GetAppid(req)+" FROM view_Menu WHERE f_qxid = ?"
		for _, row := range rows {
			if each_map, ok := row.(map[string] interface{}); ok {
				_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,qxsql,d_update["id"] ,each_map["f_qxid"].(string),each_map["f_qxid"].(string))
				if err !=nil{
					break;
				}
			}
		}
	}

	//删除可操作部门
	if err ==nil{
		sqlstr = " DELETE from tbl_role_deptid WHERE f_roleid ="+reqs["id"].(string)
		_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr,)

		if (err ==nil) && (d_update["f_DeptType"] =="2"){
			rows_dept, _ := rjson.Get("deptlist").Array();
			deptsql :="INSERT INTO tbl_role_deptid(f_customid,f_roleid,f_deptid,f_create_time) VALUES(?,?,?,now())"
			for _, row := range rows_dept {
				if each_map, ok := row.(map[string] interface{}); ok {
					_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,deptsql,utils.GetAppid(req),d_update["id"],each_map["f_deptid"].(string))
					if err !=nil{
						break;
					}
				}
			}
		}
	}

	if err !=nil{
		alen =""
		rerr = err
		rerr=tx.Rollback()
		_= services.DB_Log(utils.GetAppid(req),conn,"角色管理","修改",reqs["id"].(string),reqs["f_update_user"].(string),"修改失败",err)
		panic("修改失败,"+err.Error())
	} else{
		alen ="0"
		rerr = nil
		rerr=tx.Commit()
		sreq,_ := json.Marshal(reqs)
		slog := "修改成功，内容："+string(sreq)
		_= services.DB_Log(utils.GetAppid(req),conn,"角色管理","修改",reqs["id"].(string),reqs["f_update_user"].(string),slog,err)
	}
	return
}

func (ctl *RoleModel) Role_Delete(req beego.Controller)(alen string,rerr error){
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
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	//检查单据状态
	err = services.Check_BillState(utils.GetAppid(req),conn,"TBL_ROLE","F_DEFAULT",reqs["id"].(string),"0")
	if err != nil {
		panic("不能删除系统默认角色！")
	}

	tx,_ := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	sqlstr := "delete from tbl_role where id = ? "

	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr,reqs["id"].(string))
	//删除权限
	if err == nil{
		sqlstr = " DELETE from tbl_role_data_perms WHERE f_roleid ="+reqs["id"].(string)
		_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr,)
	}

	if err !=nil{
		alen =""
		rerr = err
		rerr = tx.Rollback()
		_= services.DB_Log(utils.GetAppid(req),conn,"角色管理","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除失败",err)
	} else{
		alen ="0"
		rerr = nil
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"角色管理","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除成功",err)
	}
	return
}

func (ctl *RoleModel) GetList(req beego.Controller) (result utils.ResData, rerr error) {
	var sortmap map[string]interface{};
	var acount string
	var icount int
	var err error
	var sqlstr_ord string
	var sql_getqx string
	var qxname    string
	//response := utils.ResponsModel{};
	var conn = services.NewConnection(utils.GetAppid(req))

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("GetList 失败，",err)
			rerr = fmt.Errorf("%s",err)
		}
	}()

	//err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	//if err !=nil{
	//	panic("参数转换出错："+err.Error())
	//}

	var sqlstr = " SELECT a.id,a.f_name,a.f_note,a.f_create_time,a.f_ordid,a.f_create_user,b.f_name AS f_create_username "+
				//"        ,(SELECT GROUP_CONCAT(f_name SEPARATOR ',')   FROM ( SELECT f_name FROM tbl_role_data_perms role INNER JOIN tbl_menu qx ON role.f_qxid = qx.f_qxid WHERE role.f_roleid= a.id and  qx.f_level = 1 GROUP BY qx.f_name)a) AS qxname "+
				 " FROM tbl_role a LEFT JOIN tbl_user   b ON a.f_create_user = b.id  where a.f_customid = "+utils.GetAppid(req)

	wvalue := req.GetString("w_value","")//reqs["w_value"]
	pageindex := req.GetString("curPage","")//reqs["pageindex"]
	pagesize := req.GetString("pageSize","")//reqs["pagesize"]
	sort := req.GetString("sort");
	sqlstr_ord = ""
	fmt.Println(sort)
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

	if sqlstr_ord ==""{
		sqlstr_ord = " ORDER BY a.f_ordid  "
	}

	//fmt.Println(deptno)


	if wvalue != ""{
		sqlstr += " and  (a.f_name like \"%s\" )"
		sqlstr = fmt.Sprintf(sqlstr,"%"+wvalue+"%")
	}

	sqlstr_count := "select count(*) as icount from ("+sqlstr+") a "
	acount,err = services.Get_FieldByValue(utils.GetAppid(req),conn,sqlstr_count,"icount")
	if err != nil{
		panic("查询合计数出错，"+err.Error())
	}

	icount,err = strconv.Atoi(acount)
	if err != nil{
		panic("获取合计数出错 ，"+err.Error())
	}
	fmt.Println(sqlstr+sqlstr_ord)
	result,err = services.FindToListByPage(utils.GetAppid(req),conn,sqlstr+sqlstr_ord,utils.InterToInt(pageindex) ,utils.InterToInt(pagesize) )
	fmt.Println(result.Result)
	if err!=nil{
		panic("查询失败,"+err.Error())
	}else{
		for _,v := range result.Result{
			sql_getqx = "  SELECT GROUP_CONCAT(f_name SEPARATOR ',') as qxname  "+
				         "      FROM ( SELECT   f_name FROM tbl_role_data_perms role INNER JOIN tbl_menu qx ON role.f_qxid = qx.f_qxid "+
				         "             WHERE role.f_roleid= "+v["id"].(string)+" and  qx.f_level = 1 GROUP BY qx.f_name LIMIT 0,5)a "
			qxname ,err = services.Get_FieldByValue(utils.GetAppid(req),conn,sql_getqx,"qxname")
			if err ==nil{
				v["qxname"] = qxname
			} else {
				v["qxname"] = ""
			}
		}
		fmt.Println(result.Result)
		result.PageSize = utils.InterToInt(pagesize)
		result.CurPage =utils.InterToInt(pageindex)
		result.Totals = icount
		if sort != ""{
			result.Sort = sortmap
		}
	}
	return result,nil
}

func (ctl *RoleModel) GetQxList(req beego.Controller) (result utils.ResData, rerr error) {

	//var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))


	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("GetQxList 失败，",err)
			rerr = fmt.Errorf("%s",err)
		}
	}()

	//err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	//if err !=nil{
	//	panic("参数转换出错："+err.Error())
	//}
	userid := req.GetString("userid","")

	var QxList []MyJsonMemu
	sqls := "	SELECT a.f_qxid,a.f_alias AS f_name,a.f_pid,(CASE WHEN ISNULL(b.f_qxid) =1 THEN 0 ELSE 1 END) AS chk,f_visible  FROM tbl_menu a "+
			"	LEFT JOIN ( SELECT rp.f_qxid FROM tbl_role_data_perms rp  "+
			"               WHERE rp.f_customid = "+utils.GetAppid(req) +" and f_roleid= ? AND rp.f_qx_pid =?  GROUP BY rp.f_qxid) b ON a.f_qxid = b.f_qxid "+
			"	WHERE f_flag = 0 AND f_pid = ? ORDER BY f_orderid  "
	sqlsthree := "SELECT a.f_qxid,a.f_name,a.f_menu_id AS f_pid,(CASE WHEN ISNULL(b.f_qxid) =1 THEN 0 ELSE 1 END) AS chk,1 as f_visible  FROM tbl_perms a "+
				"	LEFT JOIN ( SELECT rp.f_qxid FROM tbl_role_data_perms rp  "+
				"               WHERE rp.f_customid = "+utils.GetAppid(req) +" and f_roleid= ? AND rp.f_qx_pid =?  GROUP BY rp.f_qxid) b ON a.f_qxid = b.f_qxid "+
				" WHERE f_flag = 0 AND f_menu_id = ? ORDER BY f_orderid "
	//拉取第一层
	rows, err := conn.Query( sqls,userid,"0","0");
	if err != nil{
		panic("查询出错,"+err.Error())
	}

	for rows.Next() {
		QxMenu1 := MyJsonMemu{}
		_ = rows.Scan(&QxMenu1.Qxid, &QxMenu1.Name,&QxMenu1.Pid,&QxMenu1.Achk,&QxMenu1.Visible)

		//查二层
		rows1, err1 := conn.Query(sqls,userid,QxMenu1.Qxid,QxMenu1.Qxid);
		if err1 ==nil{
			for rows1.Next(){
				QxMenu2 := Menu_2{}
				//var s = ""
				//s = strings.Replace(sqlsthree,"?","%v",-1)
				//s = fmt.Sprintf("%v",userid,QxMenu2.Qxid,QxMenu2.Qxid)
				//
				//if utils.SqlLog {
				//
				//	utils.LogOut("info","svrport = "+utils.GetAppid(req)+" 三级类查询失败 exec sql ",s)
				//}
				_ = rows1.Scan(&QxMenu2.Qxid, &QxMenu2.Name,&QxMenu2.Pid,&QxMenu2.Achk,&QxMenu2.Visible)
				//查三层
				rows2, err2 := conn.Query(sqlsthree,userid,QxMenu2.Qxid,QxMenu2.Qxid);
				if err2 ==nil{
					for rows2.Next(){
						QxMenu3 := Menu_3{}
						_ = rows2.Scan(&QxMenu3.Qxid, &QxMenu3.Name,&QxMenu3.Pid,&QxMenu3.Achk,&QxMenu3.Visible)
						QxMenu2.Subs = append(QxMenu2.Subs,QxMenu3)
					}
				}else {
					fmt.Println(err2.Error())
				}
				_ =rows2.Close()
				QxMenu1.Subs = append(QxMenu1.Subs,QxMenu2)
			}
		}
		_= rows1.Close()
		QxList = append(QxList,QxMenu1)
	}
	_ = rows.Close()
	rjso,_ := json.Marshal(QxList)

	_=json.Unmarshal([]byte(rjso), &result.Result)

	return result,nil

}

//查询角色可操作部门
func (ctl *RoleModel) Get_RoleDeptList(req beego.Controller) (result utils.ResData, rerr error) {

	//var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))


	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("Get_RoleDeptList 失败，",err)
			rerr = fmt.Errorf("%s",err)
		}
	}()

	//err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	//if err !=nil{
	//	panic("参数转换出错："+err.Error())
	//}
	userid := req.GetString("userid","")
	roleid := req.GetString("roleid","")

	var QxList []MyJsonMemu
	sqls := " CALL Msp_Role_DeptQx (?,?,?) "
	//拉取第一层
	rows, err := conn.Query( sqls,roleid,"0",userid);
	if err != nil{
		panic("查询出错,"+err.Error())
	}

	for rows.Next() {
		QxMenu1 := MyJsonMemu{}
		_ = rows.Scan(&QxMenu1.Qxid, &QxMenu1.Name,&QxMenu1.Pid,&QxMenu1.Achk)
		QxMenu1.Visible = 1
		//查二层
		rows1, err1 := conn.Query(sqls,roleid,QxMenu1.Qxid,userid);
		if err1 ==nil{
			for rows1.Next(){
				QxMenu2 := Menu_2{}
				_ = rows1.Scan(&QxMenu2.Qxid, &QxMenu2.Name,&QxMenu2.Pid,&QxMenu2.Achk )
				QxMenu2.Visible =1
				//查三层
				rows2, err2 := conn.Query(sqls,roleid,QxMenu2.Qxid,userid);
				if err2 ==nil{
					for rows2.Next(){
						QxMenu3 := Menu_3{}
						_ = rows2.Scan(&QxMenu3.Qxid, &QxMenu3.Name,&QxMenu3.Pid,&QxMenu3.Achk)
						QxMenu3.Visible = 1
						QxMenu2.Subs = append(QxMenu2.Subs,QxMenu3)
					}
				}else {
					fmt.Println(err2.Error())
				}
				_ =rows2.Close()
				QxMenu1.Subs = append(QxMenu1.Subs,QxMenu2)
			}
		}
		_= rows1.Close()
		QxList = append(QxList,QxMenu1)
	}
	_ = rows.Close()
	rjso,_ := json.Marshal(QxList)

	_=json.Unmarshal([]byte(rjso), &result.Result)

	return result,nil

}

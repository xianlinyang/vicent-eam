package models

import (
	"Eam_Server/services"
	"Eam_Server/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	"os"
	"strconv"
	"strings"
	"time"
)

type BasicInfoModel struct {
	Base_Model
}

//树形菜单返回结构
type Basic_TreeMenu struct {
	ID  		int
	F_No  		string
	F_Name    	string
	F_PID     	int
	F_Level     int
	Subs []     Basic_TreeMenu_2
}

type Basic_TreeMenu_2 struct {
	ID  		int
	F_No  		string
	F_Name    	string
	F_PID     	int
	F_Level     int
	Subs []     Basic_TreeMenu_3
}

type Basic_TreeMenu_3 struct {
	ID  		int
	F_No  		string
	F_Name    	string
	F_PID     	int
	F_Level     int
}


//资产类型树结构
type Basic_AssetsTypeMenu struct {
	ID  		int
	F_No  		string
	F_Name    	string
	F_PID     	int
	F_Type      int
	F_Order     int
	F_Note      string
	Subs []     Basic_AssetsTypeMenu2
}

type Basic_AssetsTypeMenu2 struct {
	ID  		int
	F_No  		string
	F_Name    	string
	F_PID     	int
	F_Type      int
	F_Order     int
	F_Note      string
	Subs []     Basic_AssetsTypeMenu3
}

type Basic_AssetsTypeMenu3 struct {
	ID  		int
	F_No  		string
	F_Name    	string
	F_PID     	int
	F_Type      int
	F_Order     int
	F_Note      string
	Subs []     Basic_AssetsTypeMenu4
}
type Basic_AssetsTypeMenu4 struct {
	ID  		int
	F_No  		string
	F_Name    	string
	F_PID     	int
	F_Type      int
	F_Order     int
	F_Note      string

}
//*context.BeegoInput
//部门新增
func (ctl *BasicInfoModel) Deptno_Create(req beego.Controller)(aid string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))  //通过客户信息获取数据库连接

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
	now := time.Now()
	reqs["f_status"] ="1";
	reqs["f_create_time"] =now.Format("2006-01-02 15:04:05");
	reqs["f_flag"] ="0";
	reqs["f_level"] = "0"
	reqs["f_customid"] =utils.GetAppid(req)


	aid,err = services.One_ADD_Table(utils.GetAppid(req),nil,conn,"tbl_dept",reqs)
	if (err != nil)  {
		aid =""
		rerr = err
		_= services.DB_Log(utils.GetAppid(req),conn,"部门管理","新增",aid,reqs["f_create_user"].(string),"新增失败 ",err)
	} else {
		_= services.DB_Log(utils.GetAppid(req),conn,"部门管理","新增",aid,reqs["f_create_user"].(string),"新增成功 ID ="+aid,err)
	}

	return
}
//部门修改
func (ctl *BasicInfoModel) Deptno_Update(req beego.Controller)(alen string,rerr error){
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
		sqlstr = " update tbl_dept set "+sqlstr +" where id ="+reqs["id"].(string)
	}


	_,err = services.Exec_Sql(utils.GetAppid(req),nil,conn,sqlstr)
	if err !=nil{
		alen =""
		rerr = err
		_= services.DB_Log(utils.GetAppid(req),conn,"部门管理","修改",reqs["id"].(string),reqs["f_update_user"].(string),"修改失败 ",err)
	} else{
		alen ="0"
		rerr = nil
		sreq,_ := json.Marshal(reqs)
		slog := "修改成功，内容："+string(sreq)
		_= services.DB_Log(utils.GetAppid(req),conn,"部门管理","修改",reqs["id"].(string),reqs["f_update_user"].(string),slog,err)
	}

	return
}
//部门删除
func (ctl *BasicInfoModel) Deptno_Delete(req beego.Controller)(alen string,rerr error){
	var reqs map[string]interface{};
	var d_update = make(map[string]string);
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

	//sqlstr := "update tbl_dept set f_flag = 1 where id = ? "

	d_update["f_delete_time"] = time.Now().Format("2006-01-02 15:04:05")
	d_update["f_delete_user"] = reqs["f_delete_user"].(string)
	sqlstr :=  services.Sql_JionSet(d_update)

	sqlstr = "update tbl_dept set f_flag = 1,"+sqlstr+" where id = ? "

	_,err = services.Exec_Sql(utils.GetAppid(req),nil,conn,sqlstr,reqs["id"].(string))
	if err !=nil{
		alen =""
		rerr = err
		_= services.DB_Log(utils.GetAppid(req),conn,"部门管理","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除失败",err)
	} else{
		alen ="0"
		rerr = nil
		_= services.DB_Log(utils.GetAppid(req),conn,"部门管理","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除成功",err)
	}

	return
}
//部门树查询
func (ctl *BasicInfoModel) Deptno_GetTreeList(req beego.Controller) (result utils.ResData, rerr error) {
	//var reqs map[string]interface{};
	//var reqarry [] map[string] interface{}
	var conn = services.NewConnection(utils.GetAppid(req))
	//fmt.Println( req.Ctx.Request.Method)
	
	defer func() {
		_ = conn.Close()

		if err := recover(); err != nil {
			fmt.Println("Deptno_GetTreeList 失败，",err)
			rerr = fmt.Errorf("%s",err)
		}

	}()

	var DeptList []Basic_TreeMenu;
	scustomid := utils.GetAppid(req)
	//拉取第一层
	rows, rerr := conn.Query("SELECT id,f_no,f_name,f_parentid FROM tbl_dept WHERE f_customid = "+scustomid+" and f_flag = 0 and f_parentid = 0 ORDER BY f_ordid");
	for rows.Next() {
		DeptMenu1 := Basic_TreeMenu{}
		_ = rows.Scan(&DeptMenu1.ID, &DeptMenu1.F_No,&DeptMenu1.F_Name,&DeptMenu1.F_PID)
		DeptMenu1.F_Level = 1
		//查二层
		rows1, err1 := conn.Query("SELECT id,f_no,f_name,f_parentid FROM tbl_dept WHERE f_customid = "+scustomid+" and  f_flag = 0 AND f_parentid =  ? ORDER BY f_ordid ",DeptMenu1.ID);
		if err1 ==nil{
			for rows1.Next(){
				DeptMenu2 := Basic_TreeMenu_2{}
				_ = rows1.Scan(&DeptMenu2.ID, &DeptMenu2.F_No,&DeptMenu2.F_Name,&DeptMenu2.F_PID)
				//查三层
				rows2, err2 := conn.Query("SELECT id,f_no,f_name,f_parentid FROM tbl_dept WHERE f_customid = "+scustomid+" and f_flag = 0 AND f_parentid =  ? ORDER BY f_ordid ",DeptMenu2.ID);
				DeptMenu2.F_Level = 2
				if err2 ==nil{
					for rows2.Next(){
						DeptMenu3 := Basic_TreeMenu_3{}
						_ = rows2.Scan(&DeptMenu3.ID, &DeptMenu3.F_No,&DeptMenu3.F_Name,&DeptMenu3.F_PID)
						DeptMenu3.F_Level = 3
						DeptMenu2.Subs = append(DeptMenu2.Subs,DeptMenu3)
					}
				}
				_ =rows2.Close()
				DeptMenu1.Subs = append(DeptMenu1.Subs,DeptMenu2)

			}
			_= rows1.Close()
		}

		DeptList = append(DeptList,DeptMenu1)
	}
	_= rows.Close()
	rjso,_ := json.Marshal(DeptList)


	_ = json.Unmarshal([]byte(rjso), &result.Result)

	return result,nil
}
//部门查询
func (ctl *BasicInfoModel) Deptno_GetList(req beego.Controller) (result utils.ResData, rerr error) {
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))

	var acount string
	var icount int
	//response := utils.ResponsModel{};

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
	//
	//}
	scustomid := utils.GetAppid(req)
	var sqlstr = " SELECT a.*,b.f_name AS f_create_username,c.f_name AS f_update_username,(CASE a.f_status WHEN 0 THEN '已停用'  WHEN  1 THEN  '正常' ELSE '' END ) AS f_statusname "+
				 " FROM tbl_dept a LEFT JOIN tbl_user b ON a.f_create_user = b.id LEFT JOIN tbl_user c ON a.f_update_user = c.id "+
				 " WHERE a.f_customid = "+scustomid+" and  a.f_flag = 0   "


	wvalue := req.GetString("w_value","")
	pageindex := req.GetString("curPage","")//reqs["pageindex"]
	pagesize := req.GetString("pageSize","")//reqs["pagesize"]
	pid := req.GetString("parentid","")//reqs["parentid"]curPage,pageSize
	sid := req.GetString("id","")//reqs["parentid"]curPage,pageSize
	//fmt.Println(deptno)

	if pid != ""{
		sqlstr += " and  a.f_parentid = %s "
		sqlstr = fmt.Sprintf(sqlstr,pid)
	}
	if sid != ""{
		sqlstr += " and  a.id =  "+sid
	}

	if wvalue != ""{
		sqlstr += " and  (a.f_name like \"%s\" or a.f_no like \"%s\"  )"
		sqlstr = fmt.Sprintf(sqlstr,"%"+wvalue+"%","%"+wvalue+"%")
		fmt.Println(sqlstr)
	}

	sqlstr_ord := sqlstr+ " order by a.f_ordid  "
	sqlstr_count := "select count(*) as icount from ("+sqlstr+") a "

	acount,rerr = services.Get_FieldByValue(utils.GetAppid(req),conn,sqlstr_count,"icount")
	if rerr != nil{
		panic("查询合计数出错，"+rerr.Error())
	}

	icount,rerr = strconv.Atoi(acount)
	if rerr != nil{
		panic("获取合计数出错 ，"+rerr.Error())
	}
	result,rerr = services.FindToListByPage(utils.GetAppid(req),conn,sqlstr_ord,utils.InterToInt(pageindex) ,utils.InterToInt(pagesize) )
	if rerr!=nil{
		panic("查询失败,"+rerr.Error())
	}else{
		result.PageSize = utils.InterToInt(pagesize)
		result.CurPage =utils.InterToInt(pageindex)
		result.Totals = icount
		if reqs["sort"] != nil{
			result.Sort = reqs["sort"].(map[string]interface{})
		}
	}
	//fmt.Println(result)
	return result,nil
}
//启用，停用
func (ctl *BasicInfoModel) DeptnoSetStatus(req beego.Controller) ( rerr error) {
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))


	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("部门 SetStatus 失败，",err)
			rerr = fmt.Errorf("%s",err)
			return
		}
	}()

	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	var sqlstr = " UPDATE tbl_dept set f_status = ?,f_update_user=?,f_update_time =? WHERE ID = ? "

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
	_= services.DB_Log(utils.GetAppid(req),conn,"部门管理",slog,reqs["id"].(string),reqs["f_update_user"].(string),slog,err)

	return nil
}

//单位列表查询
func (ctl *BasicInfoModel) SPUnitQuery(req beego.Controller)(result utils.ResData, rerr error) {

	var conn = services.NewConnection(utils.GetAppid(req))
	var err error

	result = utils.ResData{};
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("单位列表获取失败，",err)
			rerr = fmt.Errorf("%s",err)
			return
		}
	}()


	var sqlstr = " SELECT  ts.f_Unit FROM tbl_spinfo ts where ts.f_Unit IS not null GROUP BY ts.f_Unit LIMIT 0,100"


	result.Result,err = services.FindToList(utils.GetAppid(req),conn,sqlstr)

	if err != nil{
		panic("查询出错，"+err.Error())
	}
	if rerr!=nil{
		panic("查询失败,"+rerr.Error())
	}else{
		result.PageSize = utils.InterToInt(100)
		result.CurPage =utils.InterToInt(1)
		result.Totals = 0
	}
	//fmt.Println(result)
	return result,nil
}

//基础资料树形查询通用方法 1类型  4存放地点
func (ctl *BasicInfoModel) Basic_AssetsType_GetTreeList(req beego.Controller,atype string) (result utils.ResData, rerr error) {
	var conn = services.NewConnection(utils.GetAppid(req))

	defer func() {
		_ = conn.Close()

		if err := recover(); err != nil {
			fmt.Println("Basic_AssetsType_GetTreeList 失败，",err)
			rerr = fmt.Errorf("%s",err)
		}

	}()
	scustomid := utils.GetAppid(req)
	var AssetsTypeList []Basic_AssetsTypeMenu;
	//拉取第一层
	rows, rerr := conn.Query("SELECT id,f_no,f_name,f_parentid,f_order,f_note FROM tbl_basic WHERE f_customid="+scustomid+" and f_type ="+atype+" and f_flag = 0 and f_parentid = 0 ORDER BY f_order");
	for rows.Next() {
		Basic_AssetsTypeMenu1 := Basic_AssetsTypeMenu{}
		_ = rows.Scan(&Basic_AssetsTypeMenu1.ID, &Basic_AssetsTypeMenu1.F_No,&Basic_AssetsTypeMenu1.F_Name,&Basic_AssetsTypeMenu1.F_PID,&Basic_AssetsTypeMenu1.F_Order,&Basic_AssetsTypeMenu1.F_Note)

		//查二层
		rows1, err1 := conn.Query("SELECT id,f_no,f_name,f_parentid,f_order,f_note  FROM tbl_basic WHERE f_customid="+scustomid+" and f_type ="+atype+" and f_flag = 0 AND f_parentid =  ? ORDER BY f_order ",Basic_AssetsTypeMenu1.ID);
		if err1 ==nil{
			for rows1.Next(){
				Basic_AssetsTypeMenu2 := Basic_AssetsTypeMenu2{}
				_ = rows1.Scan(&Basic_AssetsTypeMenu2.ID, &Basic_AssetsTypeMenu2.F_No,&Basic_AssetsTypeMenu2.F_Name,&Basic_AssetsTypeMenu2.F_PID,&Basic_AssetsTypeMenu1.F_Order,&Basic_AssetsTypeMenu1.F_Note)
				//查三层
				rows2, err2 := conn.Query("SELECT id,f_no,f_name,f_parentid,f_order,f_note  FROM tbl_basic WHERE f_customid="+scustomid+" and f_type ="+atype+" and f_flag = 0 AND f_parentid =  ? ORDER BY f_order ",Basic_AssetsTypeMenu2.ID);

				if err2 ==nil{
					for rows2.Next(){
						Basic_AssetsTypeMenu3 := Basic_AssetsTypeMenu3{}
						_ = rows2.Scan(&Basic_AssetsTypeMenu3.ID, &Basic_AssetsTypeMenu3.F_No,&Basic_AssetsTypeMenu3.F_Name,&Basic_AssetsTypeMenu3.F_PID,&Basic_AssetsTypeMenu1.F_Order,&Basic_AssetsTypeMenu1.F_Note)

						Basic_AssetsTypeMenu2.Subs = append(Basic_AssetsTypeMenu2.Subs,Basic_AssetsTypeMenu3)
					}
				}
				_= rows2.Close()
				Basic_AssetsTypeMenu1.Subs = append(Basic_AssetsTypeMenu1.Subs,Basic_AssetsTypeMenu2)

			}
		}
		_ = rows1.Close()
		AssetsTypeList = append(AssetsTypeList,Basic_AssetsTypeMenu1)
	}
	_ = rows.Close()
	rjso,_ := json.Marshal(AssetsTypeList)


	_ = json.Unmarshal([]byte(rjso), &result.Result)

	return result,nil
}

//基础资料公用查询方法 atype   1 资产类型  2供应商类型 3供应商 4存放位置
func (ctl *BasicInfoModel) Basic_GetList(req beego.Controller,atype string) (result utils.ResData, rerr error) {
	var conn = services.NewConnection(utils.GetAppid(req))
	var sortmap map[string]interface{};
	var acount string
	var sqlstr_ord string
	var icount int
	var swhere string
	var sqlstr string
	//response := utils.ResponsModel{};

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("GetList 失败，",err)
			rerr = fmt.Errorf("%s",err)
		}
	}()
	scustomid := utils.GetAppid(req)

	//var sqlstr = " SELECT a.*,b.f_name AS f_create_username,c.f_name AS f_update_username "+
	//			"  FROM `tbl_Basic` a LEFT JOIN tbl_user b ON a.f_create_user = b.id LEFT JOIN tbl_user c ON a.f_update_user = c.id "+
	//			"  WHERE a.f_customid="+scustomid+" and a.f_flag = 0 and a.f_type = "+atype

	wvalue := req.GetString("w_value","")
	pageindex := req.GetString("curPage","")//reqs["pageindex"]
	pagesize := req.GetString("pageSize","")//reqs["pagesize"]
	sort := req.GetString("sort");
	parentid := req.GetString("parentid")
	userid := req.GetString("userid")
	sid :=  req.GetString("id")

	swhere = ""
	//类别查询有层级
	if ((atype=="1"||(atype=="4"))) &&(parentid != ""){
		swhere += " and  a.f_parentid ="+parentid
	}

	//供应商按类别查询
	if (atype=="3") &&(parentid != ""){
		swhere += " and   a.id in(SELECT f_proid FROM tbl_provider_type WHERE f_customid = "+scustomid+" AND f_typeid = "+parentid+") "
	}

	if ((atype=="1")||(atype=="3")||(atype=="4")) &&(sid != ""){
		swhere += " and  a.id ="+sid
	}


	if wvalue != ""{
		swhere += " and  (a.f_name like \"%s\" or a.f_no like \"%s\"  )"
		swhere = fmt.Sprintf(swhere,"%"+wvalue+"%","%"+wvalue+"%")
		//fmt.Println(sqlstr)
	}


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
	} else {
		sqlstr_ord = " order by a.id "
	}
	sqlstr = " CALL msp_Basic_Query(0,"+atype+","+userid+","+scustomid+","+pageindex+","+pagesize+",'"+swhere+"','"+sqlstr_ord+"') "
	//acount,rerr = services.Get_FieldByValue(utils.GetAppid(req),conn,sqlstr_count,"icount")
	acount,rerr = services.Get_Procedure_Data(utils.GetAppid(req),conn,sqlstr,"icount")
	if rerr != nil{
		panic("查询合计数出错，"+rerr.Error())
	}

	icount,rerr = strconv.Atoi(acount)
	if rerr != nil{
		panic("获取合计数出错 ，"+rerr.Error())
	}

	sqlstr = " CALL msp_Basic_Query(1,"+atype+","+userid+","+scustomid+","+pageindex+","+pagesize+",'"+swhere+"','"+sqlstr_ord+"') "
	result.Result,rerr = services.FindToList(utils.GetAppid(req),conn,sqlstr )

	if (rerr ==nil)&&(atype == "3"){  // 查询关联的类别
		for _,v:=range result.Result{
			typelist,err := services.FindToList(utils.GetAppid(req),conn,"  SELECT tpt.f_typeid FROM tbl_provider_type tpt WHERE tpt.f_proid = ? ",v["id"].(string))
			if err ==nil{
				v["typelist"] = typelist
			}else {
				break;
			}
		}
	}

	if rerr!=nil{
		panic("查询失败,"+rerr.Error())
	}else{
		result.PageSize = utils.InterToInt(pagesize)
		result.CurPage =utils.InterToInt(pageindex)
		result.Totals = icount
		if sort != ""{
			result.Sort = sortmap
		}
	}
//fmt.Println(result)
	return result,nil
}

//基础资料公用添加方法 atype 类型    1 资产类型  2供应商类型 3供应商 4存放位置
func (ctl *BasicInfoModel) Basic_Create(req beego.Controller,atype string)(aid string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))  //通过客户信息获取数据库连接
	var slog string
	var slogval string
	var typelist string //供应商类型

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
	slogval =string(req.Ctx.Input.RequestBody)

	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	if (atype=="3")&&(reqs["f_typelist"]==""){
		panic("供应商类型不能为空！")
	}
	if (atype=="3") {
		typelist = reqs["f_typelist"].(string)
		delete(reqs,"f_typelist")
	}

	now := time.Now()
	reqs["f_status"] ="1";
	reqs["f_create_time"] =now.Format("2006-01-02 15:04:05");
	reqs["f_flag"] ="0";
	reqs["f_type"] =atype;
	reqs["f_customid"] =utils.GetAppid(req);


	icheck,serro := services.CheckRepeat(utils.GetAppid(req),conn,reqs,[]byte(`[{"FName":"f_no","PName":"编号"},{"FName":"f_name","PName":"名称"}]`),"tbl_basic",
		                        []byte(`[{"FName":"f_type","Ptype":"=","Chcktype":0,"Fvalue":""},{"FName":"f_flag","Ptype":"=","Chcktype":1,"Fvalue":"0"},{"FName":"f_customid","Ptype":"=","Chcktype":0,"Fvalue":""}]`))
	if icheck != 1{
		panic(serro)
	}

	aid,err = services.One_ADD_Table(utils.GetAppid(req),nil,conn,"tbl_Basic",reqs)
	//供应商增加类别添加，供应商有多个类别
	if (atype == "3") &&(err==nil){
		str1 := strings.Split(typelist, ",")
		for _,v := range str1{
			_,_ = services.Exec_Sql(utils.GetAppid(req),nil,conn," INSERT INTO tbl_provider_type(f_customid,f_proid,f_typeid) values(?,?,?) ",utils.GetAppid(req),aid,v)
		}
	}

	switch  atype{
	case "1":slog ="资产类型"
	case "2":slog ="资产品牌"
	case "3":slog="供应商"
	case "4":slog="资产单位"
	case "5":slog="外部联系人"

	default:
		slog =""
	}
	if (err != nil)  {
		aid =""
		rerr = err
		_= services.DB_Log(utils.GetAppid(req),conn,slog,"新增",aid,reqs["f_create_user"].(string),"新增失败 ",err)
	} else {
		slogval +="新增成功，内容："+slogval
		_= services.DB_Log(utils.GetAppid(req),conn,slog,"新增",aid,reqs["f_create_user"].(string),slogval+" ID ="+aid,err)
	}

	return
}

//基础资料公用修改 类型    1 资产类型  2供应商类型 3供应商 4存放位置
func (ctl *BasicInfoModel) Basic_Update(req beego.Controller,atype string)(alen string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var d_update = make(map[string]string);
	var conn = services.NewConnection(utils.GetAppid(req))
	var slog string
	var slogval string
	var typelist string

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			alen = ""
			rerr = fmt.Errorf("%s",err)

		}
	}()
	fmt.Println("进入接口",req.Ctx.Input.RequestBody)
//{"f_no":"004","f_name":"全系供应商","f_order":"2","f_note":"789","f_typelist":[115,126,94,115,126,94],"f_typeid":"-1","id":"127","f_update_user":"1"}
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);
	slogval =string(req.Ctx.Input.RequestBody)
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	if (atype=="3")&&(reqs["f_typelist"]==""){
		panic("供应商类型不能为空！")
	}
	if (atype=="3"){
		typelist = reqs["f_typelist"].(string)
		delete(reqs,"f_typelist")
	}

	for key,value := range reqs {
		if key !="id"{
			if value != nil{
				d_update[key] = value.(string)
			}
		}
	}
	d_update["f_update_time"] = time.Now().Format("2006-01-02 15:04:05")


	sqlstr :=  services.Sql_JionSet(d_update)

	if sqlstr !="" {
		sqlstr = " update tbl_Basic set "+sqlstr +" where id ="+reqs["id"].(string)
	}

	//str, _ := os.Getwd()
	//
	//fmt.Println(str)

	reqs["f_type"] =atype  //必须放在这里，要不然没有f_type值，不能加在前面因为又多出一个修改字段值，怕前端传错所以不改这个值
	fmt.Println("开始检测")
	icheck,serro := services.CheckRepeat(utils.GetAppid(req),conn,reqs,[]byte(`[{"FName":"f_no","PName":"编号"},{"FName":"f_name","PName":"名称"}]`),"tbl_basic",
		[]byte(`[{"FName":"id","Ptype":"<>","Chcktype":0,"Fvalue":""},{"FName":"f_type","Ptype":"=","Chcktype":0,"Fvalue":""},{"FName":"f_flag","Ptype":"=","Chcktype":1,"Fvalue":"0"},{"FName":"f_customid","Ptype":"=","Chcktype":1,"Fvalue":"`+utils.GetAppid(req)+`"}]`)) //[]string{"f_type"}

	if icheck != 1{
		panic(serro)
	}

	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)

	//供应商增加类别添加，供应商有多个类别
	if (atype == "3") &&(err==nil){
		_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn," delete from  tbl_provider_type  where f_customid =? and f_proid = ?",utils.GetAppid(req),reqs["id"].(string))
		if err ==nil{
			str1 := strings.Split(typelist, ",")
			for _,v := range str1{
				if err ==nil{
					_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn," INSERT INTO tbl_provider_type(f_customid,f_proid,f_typeid) values(?,?,?) ",utils.GetAppid(req),reqs["id"].(string),v)
				}
			}
		}
	}

	switch  atype{
	case "1":slog ="资产类型"
	case "2":slog ="资产品牌"
	case "3":slog="供应商"
	case "4":slog="资产单位"
	case "5":slog="外部联系人"

	default:
		slog =""
	}

	if err !=nil{
		alen =""
		rerr = err
		_ = tx.Rollback()
		_= services.DB_Log(utils.GetAppid(req),conn,slog,"修改",reqs["id"].(string),reqs["f_update_user"].(string),"修改失败 ",err)
	} else{
		alen ="0"
		rerr = nil
		_ = tx.Commit()
		slogval := "修改成功，内容："+string(slogval)
		_= services.DB_Log(utils.GetAppid(req),conn,slog,"修改",reqs["id"].(string),reqs["f_update_user"].(string),slogval,err)
	}

	return
}

//基础资料公用修改
func (ctl *BasicInfoModel) Basic_Del(req beego.Controller,atype string)(rerr error){
	var reqs map[string]interface{};
	var d_update = make(map[string]string);
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var conn = services.NewConnection(utils.GetAppid(req))
	var slog string
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			rerr = fmt.Errorf("%s",err)
		}

	}()
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	d_update["f_delete_time"] = time.Now().Format("2006-01-02 15:04:05")
	d_update["f_delete_user"] = reqs["f_delete_user"].(string)
	sqlstr :=  services.Sql_JionSet(d_update)

	sqlstr = "update tbl_Basic set f_flag = 1,"+sqlstr+" where id = ? "
	//fmt.Println(sqlstr)
	_,err = services.Exec_Sql(utils.GetAppid(req),nil,conn,sqlstr,reqs["id"].(string))
	switch  atype{
	case "1":slog ="资产类型"
	case "2":slog ="资产品牌"
	case "3":slog="供应商"
	case "4":slog="资产单位"
	case "5":slog="外部联系人"

	default:
		slog =""
	}
	if err !=nil{
		rerr = err
		_= services.DB_Log(utils.GetAppid(req),conn,slog,"删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除失败",err)
	} else{
		rerr = nil
		_= services.DB_Log(utils.GetAppid(req),conn,slog,"删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除成功",err)
	}

	return rerr
}

//基础资料公用停用启用
  // 0停用 1启用
func (ctl *BasicInfoModel) Basic_StartOrStop(req beego.Controller,atype string)(rerr error){
	var reqs map[string]interface{};
	var d_update = make(map[string]string);
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var conn = services.NewConnection(utils.GetAppid(req))
	var slog string
	var slogtype string
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			rerr = fmt.Errorf("%s",err)
		}

	}()
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	if (reqs["f_status"] ==nil) || (reqs["f_status"].(string) ==""){
		panic("请传入启用停用参数！")
	}
	if (reqs["f_update_user"] ==nil)|| (reqs["f_update_user"].(string) ==""){
		panic("请传入操作人参数！")
	}
	if (reqs["id"] ==nil)|| (reqs["id"].(string) ==""){
		panic("操作ID！")
	}

	if reqs["f_status"]=="0"{
		slogtype = "停用"
	}else{
		slogtype = "启用"
	}


	d_update["f_update_time"] = time.Now().Format("2006-01-02 15:04:05")
	d_update["f_update_user"] = reqs["f_update_user"].(string)
	sqlstr :=  services.Sql_JionSet(d_update)

	sqlstr = "update tbl_Basic set f_status = "+reqs["f_status"].(string)+","+sqlstr+" where id = ? "
	//fmt.Println(sqlstr)
	_,err = services.Exec_Sql(utils.GetAppid(req),nil,conn,sqlstr,reqs["id"].(string))
	switch  atype{
	case "1":slog ="资产类型"
	case "2":slog ="资产品牌"
	case "3":slog="供应商"
	case "4":slog="资产单位"
	case "5":slog="外部联系人"

	default:
		slog =""
	}
	if err !=nil{
		rerr = err
		_= services.DB_Log(utils.GetAppid(req),conn,slog,slogtype,reqs["id"].(string),reqs["f_update_user"].(string),"操作失败",err)
	} else{
		rerr = nil
		_= services.DB_Log(utils.GetAppid(req),conn,slog,slogtype,reqs["id"].(string),reqs["f_update_user"].(string),"操作成功",err)
	}

	return rerr
}

//供应商按父级查询供应商
func (ctl *BasicInfoModel) Supplier_ByType(req beego.Controller) (result utils.ResData, rerr error) {
	//var conn = services.NewConnection(utils.GetAppid(req))
	//var sortmap map[string]interface{};
	//var acount string
	//var sqlstr_ord string
	//var icount int
	////response := utils.ResponsModel{};
	//
	//defer func() {
	//	_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
	//	if err := recover(); err != nil {
	//		fmt.Println("GetList 失败，",err)
	//		rerr = fmt.Errorf("%s",err)
	//	}
	//}()
	//scustomid := utils.GetAppid(req)
	//
	//var sqlstr = " SELECT a.*,b.f_name AS f_create_username,c.f_name AS f_update_username "+
	//	"  FROM `tbl_Basic` a "+
	//	"  left join tbl_provider_type p on a.id = p.f_proid "+
	//    "  LEFT JOIN tbl_user b ON a.f_create_user = b.id "+
	//    "  LEFT JOIN tbl_user c ON a.f_update_user = c.id "+
	//	"  WHERE a.f_customid="+scustomid+" and a.f_flag = 0 and a.f_type = 3"
	//
	//
	//wvalue := req.GetString("w_value","")
	//pageindex := req.GetString("curPage","")//reqs["pageindex"]
	//pagesize := req.GetString("pageSize","")//reqs["pagesize"]
	//sort := req.GetString("sort");
	//parentid := req.GetString("parentid")
	//sid :=  req.GetString("id")
	//
	//
	////类别查询有层级
	//if ((atype=="1"||(atype=="4"))) &&(parentid != ""){
	//	sqlstr += " and  a.f_parentid ="+parentid
	//}
	//
	////供应商按类别查询
	//if (atype=="3") &&(parentid != ""){
	//	sqlstr += " and  a.f_typeid ="+parentid
	//}
	//
	//if ((atype=="1")||(atype=="3")||(atype=="4")) &&(sid != ""){
	//	sqlstr += " and  a.id ="+sid
	//}
	//fmt.Println(sqlstr)
	//
	//
	//if wvalue != ""{
	//	sqlstr += " and  (a.f_name like \"%s\" or a.f_no like \"%s\"  )"
	//	sqlstr = fmt.Sprintf(sqlstr,"%"+wvalue+"%","%"+wvalue+"%")
	//	//fmt.Println(sqlstr)
	//}
	//
	//
	//sqlstr_count := "select count(*) as icount from ("+sqlstr+") a "
	//
	//if sort != ""{
	//	_ = json.Unmarshal([]byte(sort),&sortmap)
	//
	//	if (sortmap["sortprop"] !="") &&(sortmap["order"] !="") {
	//		fmt.Println(sortmap["order"])
	//		var ord string
	//		if sortmap["order"].(string) =="descending"{
	//			ord = " desc"
	//		} else {
	//			ord =" asc"
	//		}
	//		sqlstr_ord = " order by a."+sortmap["sortprop"].(string)+" "+ord
	//	}
	//} else {
	//	sqlstr_ord = " order by a.id "
	//}
	//
	//acount,rerr = services.Get_FieldByValue(utils.GetAppid(req),conn,sqlstr_count,"icount")
	//if rerr != nil{
	//	panic("查询合计数出错，"+rerr.Error())
	//}
	//
	//icount,rerr = strconv.Atoi(acount)
	//if rerr != nil{
	//	panic("获取合计数出错 ，"+rerr.Error())
	//}
	//result,rerr = services.FindToListByPage(utils.GetAppid(req),conn,sqlstr+sqlstr_ord,utils.InterToInt(pageindex) ,utils.InterToInt(pagesize) )
	//if rerr!=nil{
	//	panic("查询失败,"+rerr.Error())
	//}else{
	//	result.PageSize = utils.InterToInt(pagesize)
	//	result.CurPage =utils.InterToInt(pageindex)
	//	result.Totals = icount
	//	if sort != ""{
	//		result.Sort = sortmap
	//	}
	//}

	return result,nil
}

//供应商树形查询
func (ctl *BasicInfoModel) Supplier_GetTreeList(req beego.Controller) (result utils.ResData, rerr error) {
	var conn = services.NewConnection(utils.GetAppid(req))

	defer func() {
		_ = conn.Close()

		if err := recover(); err != nil {
			fmt.Println("Supplier_GetTreeList 失败，",err)
			rerr = fmt.Errorf("%s",err)
		}

	}()

	var AssetsTypeList []Basic_AssetsTypeMenu;
	scustomid := utils.GetAppid(req)
	//拉取第一层
	rows, rerr := conn.Query("CALL Msp_Data_Treelist (5,"+scustomid+", \" and f_parentid =0 \",'')");

	for rows.Next() {
		Basic_AssetsTypeMenu1 := Basic_AssetsTypeMenu{}
		_ = rows.Scan(&Basic_AssetsTypeMenu1.ID, &Basic_AssetsTypeMenu1.F_No,&Basic_AssetsTypeMenu1.F_Name,&Basic_AssetsTypeMenu1.F_PID)
		Basic_AssetsTypeMenu1.F_Type = 0
		//查二层
		sid := strconv.Itoa(Basic_AssetsTypeMenu1.ID)
		rows1, err1 := conn.Query("CALL Msp_Data_Treelist (5,"+scustomid+", \" and  f_parentid ="+sid+" \",'')");
		//fmt.Println("CALL Msp_Data_Treelist (5, \" where  f_parentid ="+sid+" \")")

		if err1 ==nil{
			for rows1.Next(){
				Basic_AssetsTypeMenu2 := Basic_AssetsTypeMenu2{}
				_ = rows1.Scan(&Basic_AssetsTypeMenu2.ID, &Basic_AssetsTypeMenu2.F_No,&Basic_AssetsTypeMenu2.F_Name,&Basic_AssetsTypeMenu2.F_PID)
				Basic_AssetsTypeMenu2.F_Type = 0
				//查三层
				sid := strconv.Itoa(Basic_AssetsTypeMenu2.ID)
				rows2, err2 := conn.Query("CALL Msp_Data_Treelist (5, "+scustomid+" ,\" and  f_parentid ="+sid+" \",'')" );

				if err2 ==nil{
					for rows2.Next(){
						Basic_AssetsTypeMenu3 := Basic_AssetsTypeMenu3{}
						_ = rows2.Scan(&Basic_AssetsTypeMenu3.ID, &Basic_AssetsTypeMenu3.F_No,&Basic_AssetsTypeMenu3.F_Name,&Basic_AssetsTypeMenu3.F_PID)
						Basic_AssetsTypeMenu3.F_Type =0
						//查四层
						//sid := strconv.Itoa(Basic_AssetsTypeMenu3.ID)
						//rows3, err3 := conn.Query("select ID,f_no,f_name,f_typeid from tbl_basic  where f_type =3 AND f_flag = 0 AND f_typeid = "+sid );
						//
						//if err3 ==nil{
						//	for rows3.Next() {
						//		Basic_AssetsTypeMenu4 := Basic_AssetsTypeMenu4{}
						//		_ = rows3.Scan(&Basic_AssetsTypeMenu4.ID, &Basic_AssetsTypeMenu4.F_No, &Basic_AssetsTypeMenu4.F_Name, &Basic_AssetsTypeMenu4.F_PID)
						//		Basic_AssetsTypeMenu4.F_Type = 1
						//		Basic_AssetsTypeMenu3.Subs = append(Basic_AssetsTypeMenu3.Subs, Basic_AssetsTypeMenu4)
						//	}
						//}
						Basic_AssetsTypeMenu2.Subs = append(Basic_AssetsTypeMenu2.Subs,Basic_AssetsTypeMenu3)
						//_ =rows3.Close()
					}
					_ =rows2.Close();
				}

				Basic_AssetsTypeMenu1.Subs = append(Basic_AssetsTypeMenu1.Subs,Basic_AssetsTypeMenu2)
			}
			_ =rows1.Close();
		} else {
			//fmt.Println(err1.Error())
		}

		AssetsTypeList = append(AssetsTypeList,Basic_AssetsTypeMenu1)
	}
	_ =rows.Close();
	rjso,_ := json.Marshal(AssetsTypeList)


	_ = json.Unmarshal([]byte(rjso), &result.Result)

	return result,nil
}
//商品列表新增
func (ctl *BasicInfoModel) AssetsList_Create(req beego.Controller)(aid string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var serro string
	var conn = services.NewConnection(utils.GetAppid(req))  //通过客户信息获取数据库连接

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			aid = ""
			rerr = fmt.Errorf(serro+"，%s",err)

			return
		}
	}()
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	if err !=nil{
		panic("参数转换出错："+err.Error())
	}
	icheck,serro := services.CheckRepeat(utils.GetAppid(req),conn,reqs,[]byte(`[{"f_Code":"f_no","PName":"商品编号"}]`),"tbl_Spinfo",nil)
	if icheck != 1{
		panic(serro)
	}

	serro = ""
	reqsp := make( map[string]interface{})
	reqsp["f_accountid"] = "1"
	reqsp["f_Code"] = reqs["f_Code"]
	reqsp["f_Name"] = reqs["f_Name"]
	reqsp["f_Status"] = "0"  //状态 0闲置、1领用、 2借用  3处置 4维修中  -1报废
	reqsp["f_Typeid"] = reqs["f_Typeid"]
	//if reqs["f_PPID"].(string) !=""{   //品牌ID
	//	reqsp["f_PPID"] = reqs["f_PPID"]
	//} else {
	//	reqsp["f_PPID"] = "-1"
	//}

	// {"f_Name":"耐克运动服展示样品","f_Typeid":"144","f_Dept":"82","f_UseDept":"82","f_PositionDept":"140","f_IsPD":"1","f_Specification":"XXL","f_Price":"600","f_Unit":"件","f_BuyDate":"2020-10-30","f_Code":"NK001","f_UseWay":"展示","f_Providerid":"98","f_create_user":"1"}
	if reqs["f_Providerid"].(string) !=""{  //供应商ID
		reqsp["f_Providerid"] = reqs["f_Providerid"]
	} else {
		reqsp["f_Providerid"] = "-1"
	}
 	if reqs["f_Specification"] != nil{
		reqsp["f_Specification"] = reqs["f_Specification"] //规格
	}
	if reqs["f_Model"] != nil{
		reqsp["f_Model"] = reqs["f_Model"] //规格
	}
	if reqs["f_Unit"] != nil{
		reqsp["f_Unit"] = reqs["f_Unit"] //规格
	}
	if reqs["f_Price"] != nil{
		if reqs["f_Price"].(string) !=""{
			reqsp["f_Price"] = reqs["f_Price"]  //价值
		} else {
			reqsp["f_Price"] = "0"
		}
	}

	//if reqs["f_Company"] != nil{
	//	reqsp["f_Company"] = reqs["f_Company"]  //所属公司
	//} else {
	//	reqsp["f_Company"] = "-1"
	//}
	reqsp["f_Company"] = "-1" //所属公司暂时全部是 -1 没什么用
	if reqs["f_Dept"] != nil{  //所属部门
		reqsp["f_Dept"] = reqs["f_Dept"]
	} else {
		reqsp["f_Dept"] = "-1"
	}
	if (reqs["f_CustodyPeople"] != nil)&&(reqs["f_CustodyPeople"].(string) != ""){  //保管人
		reqsp["f_CustodyPeople"] = reqs["f_CustodyPeople"]
		reqsp["f_Default_CustodyPeople"] = reqs["f_CustodyPeople"]
	} else {
		reqsp["f_CustodyPeople"] = "-1"
		reqsp["f_Default_CustodyPeople"] ="-1"
	}
	if (reqs["f_UseCompany"] != nil) &&(reqs["f_UseCompany"].(string) != ""){  //使用公司
		reqsp["f_UseCompany"] = reqs["f_UseCompany"]
	} else {
		reqsp["f_UseCompany"] = "-1"
	}
	if (reqs["f_UseDept"] != nil) &&(reqs["f_UseDept"].(string) != ""){
		reqsp["f_UseDept"] = reqs["f_UseDept"]  //使用部门
	} else {
		reqsp["f_UseDept"] = "-1"
	}
	if (reqs["f_UsePeople"] != nil) &&(reqs["f_UsePeople"].(string) != ""){
		reqsp["f_UsePeople"] = reqs["f_UsePeople"]  //使用人
	} else {
		reqsp["f_UsePeople"] = "-1"
	}
	if (reqs["f_PositionDept"] != nil)&&(reqs["f_PositionDept"].(string) !=""){
		reqsp["f_PositionDept"] = reqs["f_PositionDept"]  //存放位置
		reqsp["f_Default_PositionDept"] = reqs["f_PositionDept"]  //存放位置
	} else {
		reqsp["f_PositionDept"] = "-1"
		reqsp["f_Default_PositionDept"] = "-1"  //存放位置
	}
	if reqs["f_Invoice"] != nil{
		reqsp["f_Invoice"] = reqs["f_Invoice"]     //发票号
	}

	if (reqs["f_InvoiceDate"] != nil) &&(reqs["f_InvoiceDate"].(string) != ""){
		reqsp["f_InvoiceDate"] = reqs["f_InvoiceDate"]  //开票日期
	}
	if (reqs["f_BuyDate"] != nil)&&(reqs["f_BuyDate"].(string) != ""){
		reqsp["f_BuyDate"] = reqs["f_BuyDate"]  //购买日期
	}
	if (reqs["f_ScrappedDate"] != nil) &&(reqs["f_ScrappedDate"].(string) != ""){
		reqsp["f_ScrappedDate"] = reqs["f_ScrappedDate"]  //报废日期
	}
	if (reqs["f_MaintenanceDate"] != nil)&&(reqs["f_MaintenanceDate"].(string) != ""){
		reqsp["f_MaintenanceDate"] = reqs["f_MaintenanceDate"]  //过保时间
	}
	if reqs["f_IsPD"] != nil{
		reqsp["f_IsPD"] = reqs["f_IsPD"]  //是否盘点
	} else {
		reqsp["f_IsPD"] = "0"
	}

	reqsp["f_UseWay"] = reqs["f_UseWay"]  //用途
	if reqs["f_note"] != nil{
		reqsp["f_note"] = reqs["f_note"]
	}
	reqsp["f_Qty"] = "1"
	reqsp["f_create_user"] = reqs["f_create_user"]
	now := time.Now()
	reqsp["f_create_time"] =now.Format("2006-01-02 15:04:05");
	reqsp["f_flag"] ="0";
	reqsp["f_customid"] =utils.GetAppid(req);

	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	aid,err = services.One_ADD_Table(utils.GetAppid(req),tx,conn,"tbl_Spinfo",reqsp)

	//保存图片
	if (reqs["pic"] != nil)&&(reqs["pic"].(string) !=""){
		if (err ==nil) || (reqs["pic"].(string) != ""){
			sguid := utils.GetGUID().Hex()+".jpg"

			str, _ := os.Getwd()  //获取程序路径
			str = str+`\static\img\SP_Pic`
			b_con,err := utils.PathExists(str) //判断文件是否存在
			//目录不存在就创建
			if b_con == false{
				err = os.MkdirAll(str, os.ModePerm) //os.modeperm可以创建多级目录
				if err != nil{
					panic("图片存放目录创建失败，"+err.Error())
				}
			}
			//去除图片前缀
			base64str := reqs["pic"].(string)
			base64str = base64str[strings.Index(base64str, ",")+1 : len(base64str)]
			fmt.Println(base64str)
			//filpath := beego.AppConfig.String("pic_path")+`\sp_pic`
			err =utils.Base64ToFile(str+`\`+utils.GetAppid(req),sguid,base64str);


			if err ==nil{
				picsql := " INSERT INTO tbl_Spinfo_Pic(f_customid,f_accountid,f_spid,f_PicName,f_flag,f_create_time) VALUES(?,1,?,?,0,now()) "
				_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,picsql,utils.GetAppid(req),aid,sguid)
				if err !=nil{
					serro ="图片存储失败"
				}
			}else {
				serro ="图片转换出错"
			}
		}else {
			serro ="商品添加出错"
		}
	}

	if err ==nil{
		sqlwater :=  " INSERT INTO tbl_spinfo_water(f_customid,f_type,f_spid,f_billid,f_billno,f_opruserid,f_memo,f_create_time) "+
					" VALUES(?,0,?,-1,'',?,'资产登记',NOW());"
		_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlwater,utils.GetAppid(req),aid,reqs["f_create_user"].(string))
	}

	if (err != nil)  {
		aid =""
		rerr = err
		_= tx.Rollback()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产清单","新增",aid,reqs["f_create_user"].(string),"新增失败 ",err)
	} else {
		if reqs["pic"]!= nil{
			delete(reqs,"pic")
		}
		sreq,_ := json.Marshal(reqs)
		slog := " 内容："+string(sreq)
		_= services.DB_Log(utils.GetAppid(req),conn,"资产清单","新增",aid,reqs["f_create_user"].(string),"新增成功 ID ="+aid+slog,err)
		_= tx.Commit()
	}

	return
}
//商品列表修改
func (ctl *BasicInfoModel) AssetsList_Update(req beego.Controller)(rerr error) {
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var d_update = make(map[string]string);
	var conn = services.NewConnection(utils.GetAppid(req))
	var serro string
	var picsql string
	var picpath string

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			rerr = fmt.Errorf(serro+"，%s",err)

		}
	}()
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	if err !=nil{
		panic("参数转换出错："+err.Error())
	}
	icheck,serro := services.CheckRepeat(utils.GetAppid(req),conn,reqs,[]byte(`[{"FName":"f_no","PName":"商品编号"}]`),"tbl_Spinfo",
		   []byte(`[{"FName":"id","Ptype":"<>","Chcktype":0,"Fvalue":""},{"FName":"f_flag","Ptype":"=","Chcktype":1,"Fvalue":"0"}]`))
	if icheck != 1{
		panic(serro)
	}

	if (reqs["f_Default_CustodyPeople"] != nil){
		if reqs["f_Default_CustodyPeople"].(string) != reqs["f_CustodyPeople"].(string){
			//panic("当前商品已存在库存变动单据，不能再修改保管人员！")
			reqs["f_Default_CustodyPeople"] = reqs["f_CustodyPeople"]
			delete(reqs,"f_CustodyPeople")
			delete(reqs,"f_Default_CustodyPeople")
		}
	}

	if (reqs["f_Default_PositionDept"] != nil){
		if reqs["f_Default_PositionDept"].(string) != reqs["f_PositionDept"].(string){
			//panic("当前商品已存在库存变动单据，不能再修改存放位置！")
			reqs["f_Default_PositionDept"] = reqs["f_PositionDept"]
			delete(reqs,"f_Default_PositionDept")
			delete(reqs,"f_PositionDept")
		}

	}

	if reqs["f_InvoiceDate"] != nil{
		delete(reqs,"f_InvoiceDate")
	}
	if reqs["f_BuyDate"]!= nil{
		delete(reqs,"f_BuyDate")
	}
	if reqs["f_ScrappedDate"] != nil{
		delete(reqs,"f_ScrappedDate")
	}
	if reqs["f_MaintenanceDate"] != nil{
		delete(reqs,"f_MaintenanceDate")
	}
	if reqs["f_UseDate"]!= nil{
		delete(reqs,"f_UseDate")
	}



	for key,value := range reqs {
		if (key !="id")&&(key !="pic"){
			if value != nil{
				d_update[key] = value.(string)
			}
		}
	}
	d_update["f_update_time"] = time.Now().Format("2006-01-02 15:04:05")
	//d_update["f_Default_CustodyPeople"] = d_update["f_CustodyPeople"]
	//d_update["f_Default_PositionDept"] = d_update["f_PositionDept"]
	sqlstr :=  services.Sql_JionSet(d_update)

	if sqlstr !="" {
		sqlstr = " update tbl_spinfo set "+sqlstr +" where id ="+reqs["id"].(string)
	}


	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)



	//删除图片
	if err ==nil{
		picpath, _ = os.Getwd()  //获取程序路径
		picpath = picpath+`\static\img\SP_Pic`
		b_con,err := utils.PathExists(picpath) //判断文件是否存在
		//目录不存在就创建
		if b_con == false{
			err = os.MkdirAll(picpath, os.ModePerm) //os.modeperm可以创建多级目录
			if err != nil{
				panic("图片存放目录创建失败，"+err.Error())
			}
		}

		sfilename,err := services.Get_FieldByValue(utils.GetAppid(req),conn," SELECT tsp.f_PicName FROM tbl_spinfo_pic tsp WHERE tsp.f_spid = "+reqs["id"].(string),"f_PicName")
		if err ==nil{
			os.Remove(picpath+`\`+utils.GetAppid(req)+`\`+sfilename)
		}
		picsql = " delete from tbl_Spinfo_Pic where f_spid  = "+reqs["id"].(string)
		_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,picsql)
		if err !=nil{
			serro ="删除图片出错,"+err.Error()
		}
	}

	//保存图片
	if (err ==nil) && (reqs["pic"].(string) != ""){
		sguid := utils.GetGUID().Hex()+".jpg"
		base64str := reqs["pic"].(string)

		//filpath := beego.AppConfig.String("pic_path")+`\sp_pic`
		//去除图片前缀
		base64str = base64str[strings.Index(base64str, ",")+1 : len(base64str)]
		fmt.Println(base64str)
		err =utils.Base64ToFile(picpath+`\`+utils.GetAppid(req),sguid,base64str);

		if err ==nil{
			picsql = " INSERT INTO tbl_Spinfo_Pic(f_customid,f_accountid,f_spid,f_PicName,f_flag,f_create_time) VALUES(?,1,?,?,0,now()) "
			_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,picsql,utils.GetAppid(req),reqs["id"].(string),sguid)
			if err !=nil{
				serro ="图片存储失败"
			}
		}else {
			serro ="图片转换出错"
		}
	}else {
		serro ="商品添加出错"
	}

	if err !=nil{
		_ =tx.Rollback()
		rerr = err
		_= services.DB_Log(utils.GetAppid(req),conn,"资产清单","修改",reqs["id"].(string),reqs["f_update_user"].(string),"修改失败 ",err)
	} else{
		_ =tx.Commit()
		rerr = nil
		if reqs["pic"]!= nil{
			delete(reqs,"pic")
		}
		sreq,_ := json.Marshal(reqs)
		slog := "修改成功，内容："+string(sreq)
		_= services.DB_Log(utils.GetAppid(req),conn,"资产清单","修改",reqs["id"].(string),reqs["f_update_user"].(string),slog,err)
	}

	return
}
//商品列表删除
func (ctl *BasicInfoModel) AssetsList_Del(req beego.Controller)( rerr error) {
	var reqs map[string]interface{};
	var d_update = make(map[string]string);
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var conn = services.NewConnection(utils.GetAppid(req))
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			rerr = fmt.Errorf("%s",err)
		}

	}()
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	//sqlstr := "update tbl_dept set f_flag = 1 where id = ? "

	d_update["f_delete_time"] = time.Now().Format("2006-01-02 15:04:05")
	d_update["f_delete_user"] = reqs["f_delete_user"].(string)
	sqlstr :=  services.Sql_JionSet(d_update)

	sqlstr = "update tbl_spinfo set f_flag = 1,"+sqlstr+" where id = ? "

	_,err = services.Exec_Sql(utils.GetAppid(req),nil,conn,sqlstr,reqs["id"].(string))
	if err !=nil{

		rerr = err
		_= services.DB_Log(utils.GetAppid(req),conn,"资产清单","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除失败",err)
	} else{

		rerr = nil
		_= services.DB_Log(utils.GetAppid(req),conn,"资产清单","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除成功",err)
	}

	return
}
//商品列表查询
func (ctl *BasicInfoModel) AssetsList_GetList(req beego.Controller)( result utils.ResData,rerr error) { //result utils.ResData
	var sortmap map[string]interface{};
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))
	var sqlstr_ord string
	var suserid string
	var gjwhere string


	var icount int

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("GetList 失败，",err)
			rerr = fmt.Errorf("%s",err)
		}
	}()

	stoken := req.Ctx.Input.Header("authorization")

	//查询USERID

	suserid =req.GetString("userid","")
	if suserid ==""{
		suserid,rerr = services.Get_FieldByValue(utils.GetAppid(req),conn,"SELECT tt.f_userid FROM tbl_token tt WHERE tt.f_token = "+"\""+stoken+"\"","f_userid")

		if (rerr != nil)||(suserid ==""){
			panic("用户信息获取出错，请检查TOKEN是否正确！")
		}
	}



	fmt.Println(suserid)
	wvalue := "\""+req.GetString("w_value","")+"\""
	pageindex := req.GetString("curPage","")
	pagesize := req.GetString("pageSize","")
	//SpreadWhere := "\""+req.GetString("SpreadWhere","")+"\""
	sort := req.GetString("sort");
	SpreadWhere := req.GetString("SpreadWhere");
	gjwhere = ""
	if SpreadWhere != ""{ //高级筛选处理

		rjson,err := simplejson.NewJson([]byte(SpreadWhere));
		if err == nil{
			if rjson.Get("f_status").MustString() != ""{
				gjwhere = gjwhere+" and ts.f_status ="+rjson.Get("f_status").MustString()
			}
			//类别
			if rjson.Get("f_typeId").MustString() != ""{
				gjwhere = gjwhere+" and ts.f_typeid ="+rjson.Get("f_typeId").MustString()
			}
			if rjson.Get("f_Specification").MustString() != ""{ //规格
				gjwhere = gjwhere+" and ts.f_Specification ="+rjson.Get("f_Specification").MustString()
			}
			if rjson.Get("f_Model").MustString() != ""{ //型号
				gjwhere = gjwhere+" and ts.f_Model ="+rjson.Get("f_Model").MustString()
			}
			if rjson.Get("Providername ").MustString() != ""{ //供应商
				gjwhere = gjwhere+" and ts.f_Providerid in (SELECT id FROM tbl_basic tb WHERE tb.f_customid = "+utils.GetAppid(req)+" AND tb.f_type = 3 AND tb.f_flag = 0 AND f_name LIKE '%"+rjson.Get("Providername").MustString()+"%')"


			}
			if rjson.Get("f_deptId").MustString() != ""{ //管理部门
				gjwhere = gjwhere+" and ts.f_Dept ="+rjson.Get("f_deptId").MustString()
			}
			if rjson.Get("f_Price").MustString() != ""{ //价值
				gjwhere = gjwhere+" and ts.f_Price ="+rjson.Get("f_Price").MustString()
			}
			if rjson.Get("f_UseDeptId").MustString() != ""{ //使用部门
				gjwhere = gjwhere+" and ts.f_UseDept ="+rjson.Get("f_UseDeptId").MustString()
			}
			if rjson.Get("f_CustodyPeopleName") .MustString() != ""{ //管理人员
				gjwhere = gjwhere+" and  ts.f_CustodyPeople in (SELECT id FROM tbl_user tu WHERE  tu.f_customid = "+utils.GetAppid(req)+" AND tu.f_flag = 0 and f_name LIKE '%"+rjson.Get("f_CustodyPeopleName").MustString()+"%')"
			}
			if rjson.Get("f_UsePeopleName") .MustString() != ""{ //使用人员
				gjwhere = gjwhere+" and  ts.f_UsePeople in (SELECT id FROM tbl_user tu WHERE tu.f_customid = "+utils.GetAppid(req)+" AND tu.f_flag = 0 and f_name LIKE '%"+rjson.Get("f_UsePeopleName").MustString()+"%')"
			}
			if rjson.Get("f_PositionId").MustString() != ""{ //存放位置
				gjwhere = gjwhere+" and ts.f_PositionDept ="+rjson.Get("f_PositionId").MustString()
			}
			if (rjson.Get("f_bBuyDate") .MustString() != "")&&(rjson.Get("f_eBuyDate") .MustString() != ""){
				gjwhere = gjwhere+" and ts.f_BuyDate >= '"+rjson.Get("f_bBuyDate").MustString()+" 00:00:00"+
					"' and ts.f_BuyDate <= '"+rjson.Get("f_eBuyDate").MustString()+" 23:59:59"+"'"
			}
			if (rjson.Get("f_bdate") .MustString() != "")&&(rjson.Get("f_edate") .MustString() != ""){
				gjwhere = gjwhere+" and ts.f_create_time >= '"+rjson.Get("f_bdate").MustString()+" 00:00:00"+
					"' and ts.f_create_time <= '"+rjson.Get("f_edate").MustString()+" 23:59:59"+"'"
			}
			gjwhere = gjwhere +" "
		}
	}

	//资产列表高级搜索： ：资产状态   ：资产类别  规格：  型号：  供应商：   所属部门：
	//管理人员：  使用部门：  使用人员：  存放位置：   资产价值：   购买日期起
	//：f_bBuyDate  购买日期止：f_eBuyDate  登记日期起：f_bdate  登记日期止：f_edate
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
			sqlstr_ord = "\""+" order by a."+sortmap["sortprop"].(string)+" "+ord+"\""
		}
	}else {
		sqlstr_ord ="\""+"\""
	}


	var sqlstr = "   CALL msp_spinfo_query("+suserid+","+wvalue+","+pageindex+","+pagesize+",\""+gjwhere+"\","+sqlstr_ord+")  "
	fmt.Println(sqlstr)
	result.Result,rerr = services.FindToList(utils.GetAppid(req),conn,sqlstr )

	//循环装入图片
	if rerr == nil{
		filpath := beego.AppConfig.String("pic_path")+`/images/SP_Pic`
		for _,value := range result.Result{

			if (value["f_PicName"] != nil) &&(value["f_PicName"].(string) != "") {
				fmt.Println("aa")
				filestr 	:= filpath+`/`+utils.GetAppid(req)+`/`+ value["f_PicName"].(string)//utils.FileToBase64(utils.SvrPicPath+`\`+utils.GetAppid(req)+`\`+ value["f_PicName"].(string))
				value["pic_value"] = filestr
				//if err ==nil{
				//	value["pic_value"] = filestr
				//}
			} else {
				value["pic_value"] = ""
			}
		}
	}

	if rerr != nil{
		panic("查询出错 ，"+rerr.Error())
	}
	if len(result.Result) > 0{
		icount,rerr = strconv.Atoi(result.Result[0]["codecount"].(string))
	}

	if rerr != nil{
		panic("获取合计数出错 ，"+rerr.Error())
	}

	if rerr!=nil{
		panic("查询失败,"+rerr.Error())
	}else{
		result.PageSize = utils.InterToInt(pagesize)
		result.CurPage =utils.InterToInt(pageindex)
		result.Totals = icount
		if reqs["sort"] != nil{
			result.Sort = reqs["sort"].(map[string]interface{})
		}
	}

	return result,nil
}
//商品资料导入
func (ctl *BasicInfoModel) AssetsImport(req beego.Controller)( result utils.ResData,rerr error) { //result utils.ResData
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))
	var suserid string

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("AssetsImport 失败，",err)
			rerr = fmt.Errorf("%s",err)
		}
	}()

    err := json.Unmarshal(req.Ctx.Input.RequestBody,&reqs)
    if err !=nil{
    	panic("参数转换出错,"+err.Error())
	}
	//stoken := req.Ctx.Input.Header("authorization")

	jsonstr,err := json.Marshal(reqs)
	if err != nil{
		panic("数据转换出错，"+err.Error())
	}

	rjson,err := simplejson.NewJson([]byte(jsonstr))
	if err != nil{
		panic("JSON装载出错，"+err.Error())
	}
	suserid = rjson.Get("userid").MustString();
	rows,err := rjson.Get("splist").Array()

	_,err = services.Exec_Sql(utils.GetAppid(req),nil,conn," CREATE TEMPORARY TABLE IF NOT EXISTS temp_spinfo_import(f_code varchar(100),f_name varchar(100),f_typename varchar(100),f_prodname varchar(100),f_Specification varchar(100),f_model varchar(100),f_unit varchar(100),f_Price varchar(100),f_source varchar(100), "+
	                                                                "   f_Companyname varchar(100),f_deptname varchar(100),f_useway varchar(100),f_note varchar(200),f_invoice varchar(100),f_buydate varchar(100), "+
																	"   f_invoicedate varchar(100),f_scrappeddate varchar(100),f_maintenancedate varchar(100),f_PositionDept varchar(100),f_CustodyPeople varchar(100) );");
	if err != nil{
		panic("临时表创建失败，"+err.Error())
	}
	sql_insert :="  INSERT INTO temp_spinfo_import(f_code,f_name,f_typename,f_prodname,f_Specification,f_model,f_unit,f_Price,f_source,f_deptname,f_useway,"+
		         "                                 f_note,f_invoice,f_buydate,f_invoicedate,"+
		         "                                 f_scrappeddate,f_maintenancedate,f_PositionDept,f_CustodyPeople)"+
				 " VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	for _,row := range rows  {
		if each_map, ok := row.(map[string] interface{}); ok {
			_,err = services.Exec_Sql(utils.GetAppid(req),nil,conn,sql_insert,
				each_map["f_code"].(string),
				each_map["f_name"].(string),
				each_map["f_typename"].(string),
				each_map["f_prodname"].(string),
				each_map["f_Specification"].(string), //规格
				each_map["f_model"].(string),//型号
				each_map["f_unit"].(string),
				each_map["f_Price"].(string),
				each_map["f_source"].(string),
				each_map["f_deptname"].(string),
				each_map["f_useway"].(string),
				each_map["f_note"].(string),
				each_map["f_invoice"].(string),
				each_map["f_buydate"].(string),
				each_map["f_invoicedate"].(string),
				each_map["f_scrappeddate"].(string),
				each_map["f_maintenancedate"].(string),
				each_map["f_PositionDept"].(string),
				each_map["f_CustodyPeople"].(string))

			if err != nil{
				break;
			}
		}
	}

	if err==nil{
		sqlstr := "CALL msp_spinfo_import ("+utils.GetAppid(req)+","+suserid+")"

		result.Result,rerr = services.FindToList(utils.GetAppid(req),conn,sqlstr)
	}

	if err ==nil{
		_,err = services.Exec_Sql(utils.GetAppid(req),nil,conn,"DROP TEMPORARY TABLE temp_sptype")
	}

	if rerr!=nil{
		panic("导入失败,"+rerr.Error())
	}else{
		if result.Result[0]["f_state"].(string) !="1"{
			panic(result.Result[0]["f_errors"].(string))
		}else{
			result.PageSize = 100
			result.CurPage = 1
			result.Totals = 1
		}
	}

	return result,nil
}

//许可信息查询
func (ctl *BasicInfoModel) License_GetList(req beego.Controller)( result utils.ResData,rerr error) { //result utils.ResData
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("GetList 失败，",err)
			rerr = fmt.Errorf("%s",err)
		}
	}()

	var sqlstr = " SELECT   * FROM tbl_System where f_customid ="+utils.GetAppid(req)+" LIMIT 0,1 "

	result,rerr = services.FindToListByPage(utils.GetAppid(req),conn,sqlstr,1 ,1 )
	if rerr!=nil{
		panic("查询失败,"+rerr.Error())
	}else{
		result.PageSize = 10
		result.CurPage =1
		result.Totals = 1
		if reqs["sort"] != nil{
			result.Sort = reqs["sort"].(map[string]interface{})
		}
	}
	//fmt.Println(result)
	return result,nil
}


//数据字典类型查询
func (ctl *BasicInfoModel) Dictionary_Type_GetTreeList(req beego.Controller)( result utils.ResData,rerr error) { //result utils.ResData

	var conn = services.NewConnection(utils.GetAppid(req))

	defer func() {
		_ = conn.Close()

		if err := recover(); err != nil {
			fmt.Println("Dictionary_Type_GetTreeList 失败，",err)
			rerr = fmt.Errorf("%s",err)
		}

	}()

	scustomid := utils.GetAppid(req)
	var DeptList []Basic_TreeMenu;
	//拉取第一层
	rows, rerr := conn.Query("SELECT id,f_Ch_name,f_parentid FROM tbl_dictionary WHERE f_customid ="+scustomid+" and  f_level = 0 AND f_parentid = 0 AND f_status = 1 ORDER BY f_ordid");
	for rows.Next() {
		DeptMenu1 := Basic_TreeMenu{}
		_ = rows.Scan(&DeptMenu1.ID,&DeptMenu1.F_Name,&DeptMenu1.F_PID)
		//查二层
		rows1, err1 := conn.Query(" SELECT id,f_Ch_name,f_parentid FROM tbl_dictionary WHERE f_customid ="+scustomid+" and f_level = 0   AND f_status = 1 and f_parentid =  ? ORDER BY f_ordid  ",DeptMenu1.ID);
		if err1 ==nil{
			for rows1.Next(){
				DeptMenu2 := Basic_TreeMenu_2{}
				_ = rows1.Scan(&DeptMenu2.ID,&DeptMenu2.F_Name,&DeptMenu2.F_PID)

				DeptMenu1.Subs = append(DeptMenu1.Subs,DeptMenu2)

			}
			_= rows1.Close()
		}
		DeptList = append(DeptList,DeptMenu1)
	}
	_= rows.Close()
	rjso,_ := json.Marshal(DeptList)

	_ = json.Unmarshal([]byte(rjso), &result.Result)

	return result,nil
}
//新增
func (ctl *BasicInfoModel) Dictionary_Type_Create(req beego.Controller)(aid string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var serro string
	var conn = services.NewConnection(utils.GetAppid(req))  //通过客户信息获取数据库连接

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			aid = ""
			rerr = fmt.Errorf(serro+"，%s",err)

			return
		}
	}()
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	if err !=nil{
		panic("参数转换出错："+err.Error())
	}
	icheck,serro := services.CheckRepeat(utils.GetAppid(req),conn,reqs,[]byte(`[{"FName":"f_Ch_name","PName":"名称"}]`),"tbl_dictionary",
		[]byte(`[{"FName":"f_type","Ptype":"=","Chcktype":1,"Fvalue":"1"},{"FName":"f_level","Ptype":"=","Chcktype":1,"Fvalue":"0"}]`))
	if icheck != 1{
		panic(serro)
	}

	serro = ""
	reqsp := make( map[string]interface{})
	reqsp["f_accountid"] = "1"
	reqsp["f_no"] = ""
	reqsp["f_Ch_name"] = reqs["f_Ch_name"]
	reqsp["f_Status"] = "1"  //1在用、0停用
	reqsp["f_Eh_name"] = reqs["f_Eh_name"]
	reqsp["f_default"] = "0"//reqs["f_default"]
	reqsp["f_issys"] = "0"
	reqsp["f_ordid"] = reqs["f_ordid"]
	reqsp["f_type"] = "1"
	reqsp["f_note"] = ""
	reqsp["f_level"] = "0"
	reqsp["f_parentid"] = reqs["f_parentid"]
	reqsp["f_customid"] = utils.GetAppid(req)


	reqsp["f_create_user"] = reqs["f_create_user"]
	now := time.Now()
	reqsp["f_create_time"] =now.Format("2006-01-02 15:04:05");


	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	aid,err = services.One_ADD_Table(utils.GetAppid(req),tx,conn,"tbl_dictionary",reqsp)


	if (err != nil)  {
		aid =""
		rerr = err
		_= tx.Rollback()
		_= services.DB_Log(utils.GetAppid(req),conn,"数据字典","新增",aid,reqs["f_create_user"].(string),"新增失败 ",err)
	} else {
		sreq,_ := json.Marshal(reqs)
		slog := " 内容："+string(sreq)
		_= services.DB_Log(utils.GetAppid(req),conn,"数据字典","新增",aid,reqs["f_create_user"].(string),"新增成功 ID ="+aid+slog,err)
		_= tx.Commit()
	}

	return
}
//修改
func (ctl *BasicInfoModel) Dictionary_Type_Update(req beego.Controller)(rerr error) {
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var d_update = make(map[string]string);
	var conn = services.NewConnection(utils.GetAppid(req))
	var serro string

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			rerr = fmt.Errorf(serro+"，%s",err)

		}
	}()
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	if err !=nil{
		panic("参数转换出错："+err.Error())
	}
	icheck,serro := services.CheckRepeat(utils.GetAppid(req),conn,reqs,[]byte(`[{"f_Code":"f_Ch_name","PName":"名称"}]`),"tbl_dictionary",
		[]byte(`[{"FName":"id","Ptype":"<>","Chcktype":0,"Fvalue":""},{"FName":"f_type","Ptype":"=","Chcktype":1,"Fvalue":"1"},{"FName":"f_level","Ptype":"=","Chcktype":1,"Fvalue":"0"}]`))
	if icheck != 1{
		panic(serro)
	}


	for key,value := range reqs {
		if (key !="id"){
			if value != nil{
				d_update[key] = value.(string)
			}
		}
	}
	d_update["f_update_time"] = time.Now().Format("2006-01-02 15:04:05")

	sqlstr :=  services.Sql_JionSet(d_update)

	if sqlstr !="" {
		sqlstr = " update tbl_dictionary set "+sqlstr +" where id ="+reqs["id"].(string)
	}


	_,err = services.Exec_Sql(utils.GetAppid(req),nil,conn,sqlstr)


	if err !=nil{
		rerr = err
		_= services.DB_Log(utils.GetAppid(req),conn,"数据字典","修改",reqs["id"].(string),reqs["f_update_user"].(string),"修改失败 ",err)
	} else{
		rerr = nil
		sreq,_ := json.Marshal(reqs)
		slog := "修改成功，内容："+string(sreq)
		_= services.DB_Log(utils.GetAppid(req),conn,"数据字典","修改",reqs["id"].(string),reqs["f_update_user"].(string),slog,err)
	}

	return
}
//删除
func (ctl *BasicInfoModel) Dictionary_Type_Del(req beego.Controller)( rerr error) {
	var reqs map[string]interface{};
	var sqlstr_count string

	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var conn = services.NewConnection(utils.GetAppid(req))
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			rerr = fmt.Errorf("%s",err)
		}

	}()
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	if err !=nil{
		panic("参数转换出错："+err.Error())
	}


	sqlstr_count = "select id from tbl_dictionary where f_parentid = "+reqs["id"].(string)

	acount,rerr := services.Get_FieldByValue(utils.GetAppid(req),conn,sqlstr_count,"id")

	if rerr != nil{
		panic("数据检测出错："+err.Error())
	}
	if acount != ""{
		panic("当前类型存在下级数据，不允许删除!")
	}

	sqlstr_count = "select id from tbl_dictionary where id = "+reqs["id"].(string)+" and f_issys = 1"

	acount,rerr = services.Get_FieldByValue(utils.GetAppid(req),conn,sqlstr_count,"id")

	if rerr != nil{
		panic("数据检测出错："+err.Error())
	}
	if acount != ""{
		panic("不能删除系统内置参数!")
	}

	sqlstr := "delete  from tbl_dictionary  where id = ? "

	_,err = services.Exec_Sql(utils.GetAppid(req),nil,conn,sqlstr,reqs["id"].(string))
	if err !=nil{

		rerr = err
		_= services.DB_Log(utils.GetAppid(req),conn,"数据字典","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除失败",err)
	} else{

		rerr = nil
		_= services.DB_Log(utils.GetAppid(req),conn,"数据字典","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除成功",err)
	}

	return
}

func (ctl *BasicInfoModel) Dictionary_GetList(req beego.Controller)( result utils.ResData,rerr error) { //result utils.ResData

	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))

	var acount string
	var icount int
	//response := utils.ResponsModel{};

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("GetList 失败，",err)
			rerr = fmt.Errorf("%s",err)
		}
	}()


	icheck,serro := services.CheckRepeat(utils.GetAppid(req),conn,reqs,[]byte(`[{"f_Code":"f_Ch_name","PName":"名称"},{"f_Code":"f_no","PName":"编码"}]`),"tbl_dictionary",nil)
	if icheck != 1{
		panic(serro)
	}

	var sqlstr = " SELECT  (@i:=@i+1) as ord_id,a.*,b.f_name AS f_create_username,c.f_name AS f_update_username,(CASE a.f_status WHEN 0 THEN '已停用'  WHEN  1 THEN  '正常' ELSE '' END ) AS f_statusname "+
	" FROM tbl_dictionary   a LEFT JOIN tbl_user b ON a.f_create_user = b.id LEFT JOIN tbl_user c ON a.f_update_user = c.id INNER JOIN (select   @i:=0)  t2 ON 1=1 "+
	" WHERE a.f_customid = "+utils.GetAppid(req)+" and  a.f_parentid =   "+req.GetString("parentid","")


	wvalue := req.GetString("w_value","")
	pageindex := req.GetString("curPage","")//reqs["pageindex"]
	pagesize := req.GetString("pageSize","")//reqs["pagesize"]



	if wvalue != ""{
		sqlstr += " and  (a.f_ch_name like \"%s\" or a.f_no like \"%s\"  )"
		sqlstr = fmt.Sprintf(sqlstr,"%"+wvalue+"%","%"+wvalue+"%")
		fmt.Println(sqlstr)
	}

	sqlstr_ord := sqlstr+ " order by a.f_ordid  "
	sqlstr_count := "select count(*) as icount from ("+sqlstr+") a "
	fmt.Println(sqlstr_count)
	acount,rerr = services.Get_FieldByValue(utils.GetAppid(req),conn,sqlstr_count,"icount")
	if rerr != nil{
		panic("查询合计数出错，"+rerr.Error())
	}

	icount,rerr = strconv.Atoi(acount)
	if rerr != nil{
		panic("获取合计数出错 ，"+rerr.Error())
	}
	fmt.Println(sqlstr_ord)
	result,rerr = services.FindToListByPage(utils.GetAppid(req),conn,sqlstr_ord,utils.InterToInt(pageindex) ,utils.InterToInt(pagesize) )

	if rerr!=nil{
		panic("查询失败,"+rerr.Error())
	}else{
		result.PageSize = utils.InterToInt(pagesize)
		result.CurPage =utils.InterToInt(pageindex)
		result.Totals = icount
		if reqs["sort"] != nil{
			result.Sort = reqs["sort"].(map[string]interface{})
		}
	}
	//fmt.Println(result)
	return result,nil
}
//数据字典列表新增
func (ctl *BasicInfoModel) Dictionary_Create(req beego.Controller)(aid string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var serro string
	var conn = services.NewConnection(utils.GetAppid(req))  //通过客户信息获取数据库连接

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			aid = ""
			rerr = fmt.Errorf(serro+"，%s",err)

			return
		}
	}()
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	if err !=nil{
		panic("参数转换出错："+err.Error())
	}
	icheck,serro := services.CheckRepeat(utils.GetAppid(req),conn,reqs,[]byte(`[{"f_Code":"f_Ch_name","PName":"名称"}]`),"tbl_dictionary",
		[]byte(`[{"FName":"f_type","Ptype":"=","Chcktype":1,"Fvalue":"1"},{"FName":"f_level","Ptype":"=","Chcktype":1,"Fvalue":"1"}]`))
	if icheck != 1{
		panic(serro)
	}

	serro = ""
	reqsp := make( map[string]interface{})
	reqsp["f_accountid"] = "1"
	reqsp["f_no"] = reqs["f_no"]
	reqsp["f_Ch_name"] = reqs["f_Ch_name"]
	reqsp["f_Status"] = "1"  //1在用、0停用
	reqsp["f_Eh_name"] = reqs["f_Eh_name"]
	reqsp["f_default"] =  reqs["f_default"]
	reqsp["f_issys"] = "0"
	reqsp["f_ordid"] = reqs["f_ordid"]
	reqsp["f_type"] = "1"
	reqsp["f_note"] = reqs["f_note"]
	reqsp["f_level"] = "1"
	reqsp["f_parentid"] = reqs["f_parentid"]
	reqsp["f_customid"] = utils.GetAppid(req)
	reqsp["f_create_user"] = reqs["f_create_user"]
	now := time.Now()
	reqsp["f_create_time"] =now.Format("2006-01-02 15:04:05");


	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	aid,err = services.One_ADD_Table(utils.GetAppid(req),tx,conn,"tbl_dictionary",reqsp)


	if (err != nil)  {
		aid =""
		rerr = err
		_= tx.Rollback()
		_= services.DB_Log(utils.GetAppid(req),conn,"数据字典","新增",aid,reqs["f_create_user"].(string),"新增失败 ",err)
	} else {
		sreq,_ := json.Marshal(reqs)
		slog := " 内容："+string(sreq)
		_= services.DB_Log(utils.GetAppid(req),conn,"数据字典","新增",aid,reqs["f_create_user"].(string),"新增成功 ID ="+aid+slog,err)
		_= tx.Commit()
	}

	return
}
//数据字典列表修改
func (ctl *BasicInfoModel) Dictionary_Update(req beego.Controller)(rerr error) {
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var d_update = make(map[string]string);
	var conn = services.NewConnection(utils.GetAppid(req))
	var serro string

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

			rerr = fmt.Errorf(serro+"，%s",err)

		}
	}()
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	if err !=nil{
		panic("参数转换出错："+err.Error())
	}
	icheck,serro := services.CheckRepeat(utils.GetAppid(req),conn,reqs,[]byte(`[{"f_Code":"f_Ch_name","PName":"名称"}]`),"tbl_dictionary",
		[]byte(`[{"FName":"id","Ptype":"<>","Chcktype":0,"Fvalue":""},{"FName":"f_type","Ptype":"=","Chcktype":1,"Fvalue":"1"},{"FName":"f_level","Ptype":"=","Chcktype":1,"Fvalue":"1"}]`))
	if icheck != 1{
		panic(serro)
	}


	for key,value := range reqs {
		if (key !="id"){
			if value != nil{
				d_update[key] = value.(string)
			}
		}
	}
	d_update["f_update_time"] = time.Now().Format("2006-01-02 15:04:05")

	sqlstr :=  services.Sql_JionSet(d_update)

	if sqlstr !="" {
		sqlstr = " update tbl_dictionary set "+sqlstr +" where id ="+reqs["id"].(string)
	}


	_,err = services.Exec_Sql(utils.GetAppid(req),nil,conn,sqlstr)


	if err !=nil{
		rerr = err
		_= services.DB_Log(utils.GetAppid(req),conn,"数据字典","修改",reqs["id"].(string),reqs["f_update_user"].(string),"修改失败 ",err)
	} else{
		rerr = nil
		sreq,_ := json.Marshal(reqs)
		slog := "修改成功，内容："+string(sreq)
		_= services.DB_Log(utils.GetAppid(req),conn,"数据字典","修改",reqs["id"].(string),reqs["f_update_user"].(string),slog,err)
	}

	return
}
//数据字典列表删除
func (ctl *BasicInfoModel) Dictionary_Del(req beego.Controller)( rerr error) {
	var reqs map[string]interface{};

	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var conn = services.NewConnection(utils.GetAppid(req))
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			rerr = fmt.Errorf("%s",err)
		}

	}()
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	if err !=nil{
		panic("参数转换出错："+err.Error())
	}


	sqlstr_count := "select id from tbl_dictionary where id = "+reqs["id"].(string)+" and f_issys = 1"

	acount,rerr := services.Get_FieldByValue(utils.GetAppid(req),conn,sqlstr_count,"id")

	if rerr != nil{
		panic("数据检测出错："+err.Error())
	}
	if acount != ""{
		panic("不能删除系统内置参数!")
	}

	sqlstr := "delete  from tbl_dictionary  where id = ? "

	_,err = services.Exec_Sql(utils.GetAppid(req),nil,conn,sqlstr,reqs["id"].(string))
	if err !=nil{

		rerr = err
		_= services.DB_Log(utils.GetAppid(req),conn,"数据字典","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除失败",err)
	} else{

		rerr = nil
		_= services.DB_Log(utils.GetAppid(req),conn,"数据字典","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除成功",err)
	}

	return
}
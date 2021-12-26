package models

import (
	"Eam_Server/services"
	"Eam_Server/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	"strconv"
)

type OptionsModel struct {
	Base_Model
}


func (ctl *OptionsModel) SysLog_GetList(req beego.Controller) (result utils.ResData, rerr error) {
	var reqs map[string]interface{};
	var sortmap map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))
	var acount string
	var icount int
	var swhere string
	var sqlstr_ord string
	var gjwhere string
	//response := utils.ResponsModel{};

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("SysLog_GetList 失败，",err)
			rerr = fmt.Errorf("%s",err)
		}
	}()



	var sqlstr = " SELECT *FROM ( "+
				" 		SELECT a.id,a.f_modal,a.f_keywords,a.f_oprtime,b.f_name AS usrname,a.f_oprtype,a.f_note,(CASE a.f_typelevel WHEN 0 THEN '普通日志' ELSE '' END) AS levelname FROM sys_log a "+
				"		LEFT JOIN tbl_user b ON a.f_userid = b.id  where a.f_customid = "+utils.GetAppid(req)+
				"  )a "
	LikeEqual := req.GetString("LikeEqual","")  // 0全部 1按模块 2按备注 3按类型 4按操作人
	LikeEqualValue := req.GetString("LikeEqualValue","")
	sbdate := req.GetString("b_date","")+" 00:00:00"
	sedate := req.GetString("e_date","") +" 23:59:59"
	pageindex := req.GetString("curPage","")//reqs["pageindex"]
	pagesize := req.GetString("pageSize","")//reqs["pagesize"]
	SpreadWhere := req.GetString("SpreadWhere");
	if SpreadWhere != ""{ //高级筛选处理
		gjwhere = ""
		rjson,err := simplejson.NewJson([]byte(SpreadWhere));
		if err == nil{
			if rjson.Get("optName").MustString() != ""{
				gjwhere = gjwhere+" and m.f_userid in (SELECT id FROM tbl_user tu WHERE tu.f_customid = "+utils.GetAppid(req)+" AND tu.f_flag = 0 and f_name LIKE '%"+rjson.Get("optName").MustString()+"%')"
			}
			if rjson.Get("f_modal").MustString() != ""{ //变更后所属部门
				gjwhere = gjwhere+" and m.f_modal = '"+rjson.Get("f_modal").MustString()+"'"
			}
			if rjson.Get("f_oprtype") .MustString() != ""{
				gjwhere = gjwhere+" and  m.f_oprtype = '"+rjson.Get("f_oprtype") .MustString()+"'"
			}
			if rjson.Get("f_note") .MustString() != ""{
				gjwhere = gjwhere+" and  m.f_note like '%"+rjson.Get("f_note") .MustString()+"%'"
			}
			if (rjson.Get("b_date") .MustString() != "")&&(rjson.Get("e_date") .MustString() != ""){
				gjwhere = gjwhere+" and m.f_oprtime >= '"+rjson.Get("b_date").MustString()+" 00:00:00"+
					"' and m.f_oprtime <= '"+rjson.Get("e_date").MustString()+" 23:59:59"+"'"
			}
			gjwhere = gjwhere +" "
		}
	}

	fmt.Println(LikeEqualValue)
	fmt.Println(utils.GetParamStr(LikeEqualValue))
	//fmt.Println(deptno)
	swhere = " where a.f_oprtime >= \""+sbdate+"\""+ " and a.f_oprtime <= \""+sedate+"\""


	equal :=utils.InterToInt(LikeEqual)
	if LikeEqualValue != "" {
		switch equal {
		case 0: swhere =swhere+ " and a.f_note like \"%"+LikeEqualValue+"%\""
		case 1: swhere =swhere+ " and a.f_modal = \""+LikeEqualValue+"\""
		case 2: {
			swhere =swhere+ " and a.f_note like \"%"+LikeEqualValue+"%\""
		}
		case 3: swhere =swhere+ " and a.f_oprtype = \""+LikeEqualValue+"\""
		case 4: swhere =swhere+ " and a.usrname = \""+LikeEqualValue+"\""
		}
	}

	if gjwhere != ""{
		swhere = swhere + gjwhere
	}

	fmt.Println(swhere)
	sqlstr_ord =  " order by a.f_oprtime desc "
	if reqs["sort"] != nil{
		sortmap = reqs["sort"].(map[string]interface{})
		if (sortmap["sortprop"] !=nil) &&(sortmap["order"] !=nil) {
			var ord string
			if sortmap["order"].(string) =="descending"{
				ord = " desc"
			} else {
				ord =" asc"
			}
			sqlstr_ord = " order by a."+sortmap["sortprop"].(string)+" "+ord
		}
	}
	sqlstr = sqlstr+swhere+sqlstr_ord
	sqlstr_count := "select count(*) as icount from ("+sqlstr+") a "
	acount,err := services.Get_FieldByValue(utils.GetAppid(req),conn,sqlstr_count,"icount")
	if err != nil{
		panic("查询合计数出错，"+err.Error())
	}

	icount,err = strconv.Atoi(acount)
	if err != nil{
		panic("获取合计数出错 ，"+err.Error())
	}

	result,err = services.FindToListByPage(utils.GetAppid(req),conn,sqlstr,utils.InterToInt(pageindex) ,utils.InterToInt(pagesize) )
	if err!=nil{
		panic("查询失败,"+err.Error())
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

//商品录入查询
func (ctl *OptionsModel) Public_Input_Query(req beego.Controller) (result utils.ResData, rerr error) {
	var reqs map[string]interface{};

	var conn = services.NewConnection(utils.GetAppid(req))
	var acount string
	var icount int
	var swhere string
	result = utils.ResData{};
	//response := utils.ResponsModel{};

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("SysLog_GetList 失败，",err)
			rerr = fmt.Errorf("%s",err)
		}
	}()



	atype := req.GetString("atype","")  // 0资产申领 1领用  2退库
	s_value := req.GetString("s_value","")
	s_userid := req.GetString("s_userid","")
	curPage := req.GetString("curPage","")
	pageSize := req.GetString("pageSize","")//reqs["pageindex"]
	swhere = ""
	var sqlstr = " CALL Msp_Spinfo_Input_Query("+atype+",0,\""+s_value+"\","+s_userid+",'',"+curPage+","+pageSize+") "


	acount,err := services.Get_Procedure_Data(utils.GetAppid(req),conn,sqlstr,"f_count")
	if err != nil{
		panic("查询合计数出错，"+err.Error())
	}

	icount,err = strconv.Atoi(acount)
	if err != nil{
		panic("获取合计数出错 ，"+err.Error())
	}
	sqlstr = " CALL Msp_Spinfo_Input_Query("+atype+",1,\""+s_value+"\","+s_userid+",\""+swhere+"\","+curPage+","+pageSize+") "

	result.Result,err = services.FindToList(utils.GetAppid(req),conn,sqlstr )
	if err!=nil{
		panic("查询失败,"+err.Error())
	}else{
		result.PageSize = utils.InterToInt(pageSize)
		result.CurPage =utils.InterToInt(curPage)
		result.Totals = icount
		if reqs["sort"] != nil{
			result.Sort = reqs["sort"].(map[string]interface{})
		}
	}
	return result,nil
}

//部门查询
func (ctl *OptionsModel) Public_Input_Dept_Query(req beego.Controller) (result utils.ResData, rerr error) {
	var conn = services.NewConnection(utils.GetAppid(req))

	defer func() {
		_ = conn.Close()

		if err := recover(); err != nil {
			fmt.Println("Public_Input_Dept_Query 失败，",err)
			rerr = fmt.Errorf("%s",err)
		}

	}()
	//customid := utils.GetAppid(req)
	suserid := req.GetString("userid","")

	var DeptList []Basic_TreeMenu;
	//拉取第一层
	rows, rerr := conn.Query("call Msp_Get_OprDept_Level ("+suserid+",0)");
	for rows.Next() {
		DeptMenu1 := Basic_TreeMenu{}
		_ = rows.Scan(&DeptMenu1.ID,&DeptMenu1.F_PID, &DeptMenu1.F_No,&DeptMenu1.F_Name)
		DeptMenu1.F_Level = 1
		//查二层
		rows1, err1 := conn.Query("call Msp_Get_OprDept_Level ("+suserid+",?)  ",DeptMenu1.ID);
		if err1 ==nil{
			for rows1.Next(){
				DeptMenu2 := Basic_TreeMenu_2{}
				_ = rows1.Scan(&DeptMenu2.ID,&DeptMenu2.F_PID, &DeptMenu2.F_No,&DeptMenu2.F_Name)
				//查三层
				rows2, err2 := conn.Query("call Msp_Get_OprDept_Level ("+suserid+",?) ",DeptMenu2.ID);
				DeptMenu2.F_Level = 2
				if err2 ==nil{
					for rows2.Next(){
						DeptMenu3 := Basic_TreeMenu_3{}
						_ = rows2.Scan(&DeptMenu3.ID,&DeptMenu3.F_PID, &DeptMenu3.F_No,&DeptMenu3.F_Name)
						DeptMenu3.F_Level = 3
						DeptMenu2.Subs = append(DeptMenu2.Subs,DeptMenu3)
					}
				}
				_= rows2.Close()
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
func (ctl *OptionsModel) Public_Input_Postion_Query(req beego.Controller) (result utils.ResData, rerr error) {
	var conn = services.NewConnection(utils.GetAppid(req))

	defer func() {
		_ = conn.Close()

		if err := recover(); err != nil {
			fmt.Println("Public_Input_Dept_Query 失败，",err)
			rerr = fmt.Errorf("%s",err)
		}

	}()

    scustomid := utils.GetAppid(req)
	var DeptList []Basic_TreeMenu;
	//拉取第一层
	rows, rerr := conn.Query("SELECT id,f_no,f_name,tb.f_parentid  FROM tbl_basic tb WHERE tb.f_customid = "+scustomid+" AND  tb.f_type = 4 AND tb.f_flag = 0 AND tb.f_parentid = 0");
	for rows.Next() {
		DeptMenu1 := Basic_TreeMenu{}
		_ = rows.Scan(&DeptMenu1.ID, &DeptMenu1.F_No,&DeptMenu1.F_Name,&DeptMenu1.F_PID)
		DeptMenu1.F_Level = 1
		//查二层
		rows1, err1 := conn.Query("SELECT id,f_no,f_name,tb.f_parentid  FROM tbl_basic tb WHERE tb.f_customid = "+scustomid+" AND  tb.f_type = 4 AND tb.f_flag = 0 AND tb.f_parentid = ? ",DeptMenu1.ID);
		if err1 ==nil{
			for rows1.Next(){
				DeptMenu2 := Basic_TreeMenu_2{}
				_ = rows1.Scan(&DeptMenu2.ID, &DeptMenu2.F_No,&DeptMenu2.F_Name,&DeptMenu2.F_PID)
				//查三层
				rows2, err2 := conn.Query("SELECT id,f_no,f_name,tb.f_parentid  FROM tbl_basic tb WHERE tb.f_customid = "+scustomid+" AND  tb.f_type = 4 AND tb.f_flag = 0 AND tb.f_parentid = ?  ",DeptMenu2.ID);
				DeptMenu2.F_Level = 2
				if err2 ==nil{
					for rows2.Next(){
						DeptMenu3 := Basic_TreeMenu_3{}
						_ = rows2.Scan(&DeptMenu3.ID, &DeptMenu3.F_No,&DeptMenu3.F_Name,&DeptMenu3.F_PID)
						DeptMenu3.F_Level = 3
						DeptMenu2.Subs = append(DeptMenu2.Subs,DeptMenu3)
					}
				}
				_= rows2.Close()
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
//单据录入人员查询
func (ctl *OptionsModel) Public_Input_People_Query(req beego.Controller) (result utils.ResData, rerr error) {
	var conn = services.NewConnection(utils.GetAppid(req))
	var sql_pepole string

	defer func() {
		_ = conn.Close()

		if err := recover(); err != nil {
			fmt.Println("Public_Input_People_Query 失败，",err)
			rerr = fmt.Errorf("%s",err)
		}

	}()
	//customid := utils.GetAppid(req)
	suserid := req.GetString("userid","")
	sdeptid := req.GetString("deptid","")
	scurPage := req.GetString("curPage","")
	spageSize := req.GetString("pageSize","")
	svalue := req.GetString("w_value","")

	var DeptList []Basic_TreeMenu;
	//拉取第一层
	rows, rerr := conn.Query("call Msp_Get_OprDept_Level ("+suserid+",0)");   //按用记获取可操作部门
	for rows.Next() {
		DeptMenu1 := Basic_TreeMenu{}
		_ = rows.Scan(&DeptMenu1.ID,&DeptMenu1.F_PID, &DeptMenu1.F_No,&DeptMenu1.F_Name)
		DeptMenu1.F_Level = 1
		//查二层
		rows1, err1 := conn.Query("call Msp_Get_OprDept_Level ("+suserid+",?)  ",DeptMenu1.ID);
		if err1 ==nil{
			for rows1.Next(){
				DeptMenu2 := Basic_TreeMenu_2{}
				_ = rows1.Scan(&DeptMenu2.ID,&DeptMenu2.F_PID, &DeptMenu2.F_No,&DeptMenu2.F_Name)
				//查三层
				rows2, err2 := conn.Query("call Msp_Get_OprDept_Level ("+suserid+",?) ",DeptMenu2.ID);
				DeptMenu2.F_Level = 2
				if err2 ==nil{
					for rows2.Next(){
						DeptMenu3 := Basic_TreeMenu_3{}
						_ = rows2.Scan(&DeptMenu3.ID,&DeptMenu3.F_PID, &DeptMenu3.F_No,&DeptMenu3.F_Name)
						DeptMenu3.F_Level = 3
						DeptMenu2.Subs = append(DeptMenu2.Subs,DeptMenu3)
					}
				}
				_= rows2.Close()
				DeptMenu1.Subs = append(DeptMenu1.Subs,DeptMenu2)

			}
			_= rows1.Close()
		}

		DeptList = append(DeptList,DeptMenu1)
	}
	_= rows.Close()

	r_data :=  make(map[string]interface{});
	r_data["dept_list"] = DeptList
	//先查人员总数
	sql_pepole =" CALL Msp_Get_OprUser ("+suserid+",0,\""+sdeptid+"\","+scurPage+","+spageSize+",\""+svalue+"\") "
	scount,err := services.Get_Procedure_Data(utils.GetAppid(req),conn,sql_pepole,"icount")
	if err != nil{
		panic("获取合计数据出错,"+err.Error())
	}

	sql_pepole =" CALL Msp_Get_OprUser ("+suserid+",1,\""+sdeptid+"\","+scurPage+","+spageSize+",\""+svalue+"\") "
	fmt.Println(sql_pepole)
	user,err := services.FindToList(utils.GetAppid(req),conn,sql_pepole)
	if err != nil{
		panic("用户查询失败,"+err.Error())
	}

	if err ==nil{
		r_data["user_list"] = user
	}

	result.Result = append(result.Result,r_data)

	icount ,err:= strconv.Atoi(scount)
	result.Totals = icount

	return result,nil
}

//单据录-所有基础信息选择
func (ctl *OptionsModel) Public_Input_Bill_BasicAll(req beego.Controller) (result utils.ResData, rerr error) {
	var conn = services.NewConnection(utils.GetAppid(req))


	defer func() {
		_ = conn.Close()

		if err := recover(); err != nil {
			fmt.Println("Public_Input_Bill_BasicAll 失败，",err)
			rerr = fmt.Errorf("%s",err)
		}

	}()
	//customid := utils.GetAppid(req)
	//suserid := req.GetString("userid","")
	//scurPage := req.GetString("curPage","")
	//spageSize := req.GetString("pageSize","")

	r_allmap := make(map[string]interface{})

	// 查部门
	r_dept,err := ctl.Public_Input_Dept_Query(req)
	if err != nil{
		panic("部门查询失败,"+err.Error())
	}
	r_allmap["dept_list"]= r_dept.Result

	//查存放地点
	r_postion,err := ctl.Public_Input_Postion_Query(req)
	if err != nil{
		panic("存放地点查询失败,"+err.Error())
	}

	r_allmap["postion_list"]= r_postion.Result

	result.Result = append(result.Result,r_allmap)

	return result,nil
}

//单据录-所有基础信息选择
func (ctl *OptionsModel) Public_AssetsFlowing(req beego.Controller) (result utils.ResData, rerr error) {
	var reqs map[string]interface{};

	var conn = services.NewConnection(utils.GetAppid(req))
	var acount string
	var icount int

	result = utils.ResData{};
	//response := utils.ResponsModel{};

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("SysLog_GetList 失败，",err)
			rerr = fmt.Errorf("%s",err)
		}
	}()


	spid := req.GetString("spid","")  // 0资产申领 1领用  2退库
	s_userid := req.GetString("userid","")
	curPage := req.GetString("curPage","")
	pageSize := req.GetString("pageSize","")//reqs["pageindex"]

	var sqlstr = " CALL Msp_Spinfo_StockList("+spid+","+s_userid+",0,"+curPage+","+pageSize+") "


	acount,err := services.Get_Procedure_Data(utils.GetAppid(req),conn,sqlstr,"icount")
	if err != nil{
		panic("查询合计数出错，"+err.Error())
	}

	icount,err = strconv.Atoi(acount)
	if err != nil{
		panic("获取合计数出错 ，"+err.Error())
	}
	sqlstr = " CALL Msp_Spinfo_StockList("+spid+","+s_userid+",1,"+curPage+","+pageSize+") "

	result.Result,err = services.FindToList(utils.GetAppid(req),conn,sqlstr )
	if err!=nil{
		panic("查询失败,"+err.Error())
	}else{
		result.PageSize = utils.InterToInt(pageSize)
		result.CurPage =utils.InterToInt(curPage)
		result.Totals = icount
		if reqs["sort"] != nil{
			result.Sort = reqs["sort"].(map[string]interface{})
		}
	}
	return result,nil
}

package models

import (
	"Eam_Server/services"
	"Eam_Server/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)


type ReportModel struct {
	Base_Model
}
type Type_List struct {
	ID  		string
	No          string
	ParentID    string
	Name 		string
	SPCount     string
	SPTotal     string
	UseQty		string
	IdleQty		string
	ScrappedQty	string
	ServiceQty	string
	DealQty		string
	AcquireQty	string
	BorrowQty	string
	Subs        [] Type_List2
}
type Type_List2 struct {
	ID  		string
	No          string
	ParentID    string
	Name 		string
	SPCount     string
	SPTotal     string
	UseQty		string
	IdleQty		string
	ScrappedQty	string
	ServiceQty	string
	DealQty		string
	AcquireQty	string
	BorrowQty	string
	Subs        [] Type_List3
}
type Type_List3 struct {
	ID  		string
	No          string
	ParentID    string
	Name 		string
	SPCount     string
	SPTotal     string
	UseQty		string
	IdleQty		string
	ScrappedQty	string
	ServiceQty	string
	DealQty		string
	AcquireQty	string
	BorrowQty	string
	Subs        [] Type_List2
}

//类别分析
func (ctl *ReportModel)Report_As_Type_Browse(req beego.Controller)(result utils.ResData, rerr error){
	var conn = services.NewConnection(utils.GetAppid(req))
	var err error
	var map_count []map[string]interface{}
	var sqlstr string
	defer func(){
		_ = conn.Close()

		if err := recover();err!=nil{
			fmt.Println("Report_As_Type_Browse查询失败",err)
			rerr = fmt.Errorf("%s",err)
		}
	}()

	sdeptid := req.GetString("deptid","") //管理部门ID
	suserid := req.GetString("userid","")
	sshowtype := req.GetString("sshowtype","")
	//curPage := req.GetString("curPage","")
	//pageSize := req.GetString("pageSize","")
	if sshowtype == ""{
		sshowtype ="1"
	}

	sqlstr = "call  MSP_Report_SPType (0,"+utils.GetAppid(req)+",0,"+suserid+",\""+sdeptid+"\",'',"+sshowtype+",1,10000)"
	//查询合计数据
	map_count,err = services.FindToList(utils.GetAppid(req),conn,sqlstr)
	if err !=nil{
		panic("查询合计数据出错，"+err.Error())
	}
	//查询父级
	sqlstr = "call  MSP_Report_SPType (1,"+utils.GetAppid(req)+",?,"+suserid+",\""+sdeptid+"\",'',"+sshowtype+",1,10000)"
	rows1, err1 := conn.Query(sqlstr,0);

	var relist  []Type_List
	if err1 == nil{
		for rows1.Next(){
			TypeList := Type_List{}
			_ = rows1.Scan(&TypeList.ID,&TypeList.No, &TypeList.ParentID,&TypeList.Name,&TypeList.SPCount,&TypeList.SPTotal,&TypeList.UseQty,
				           &TypeList.IdleQty,&TypeList.ScrappedQty,&TypeList.ServiceQty,&TypeList.DealQty,&TypeList.AcquireQty,&TypeList.BorrowQty)
			//查二层
			rows2, err2 := conn.Query(sqlstr,TypeList.ID);
			if err2 ==nil{
				for rows2.Next(){
					TypeList2 := Type_List2{}
					_ = rows2.Scan(&TypeList2.ID,&TypeList2.No, &TypeList2.ParentID,&TypeList2.Name,&TypeList2.SPCount,&TypeList2.SPTotal,&TypeList2.UseQty,
						&TypeList2.IdleQty,&TypeList2.ScrappedQty,&TypeList2.ServiceQty,&TypeList2.DealQty,&TypeList2.AcquireQty,&TypeList2.BorrowQty)
					//查三层
					rows3, err3 := conn.Query(sqlstr,TypeList2.ID);
					if err3 ==nil{
						for rows3.Next() {
							TypeList3 := Type_List3{}
							_ = rows3.Scan(&TypeList3.ID,&TypeList3.No, &TypeList3.ParentID, &TypeList3.Name, &TypeList3.SPCount, &TypeList3.SPTotal, &TypeList3.UseQty,
								&TypeList3.IdleQty, &TypeList3.ScrappedQty, &TypeList3.ServiceQty, &TypeList3.DealQty, &TypeList3.AcquireQty, &TypeList3.BorrowQty)
							TypeList2.Subs = append(TypeList2.Subs,TypeList3)
						}
					}else {
						panic("三级类查询失败"+err3.Error())
					}
					_ = rows3.Close()

					TypeList.Subs = append(TypeList.Subs,TypeList2)
				}
			} else{
				panic("二级类查询失败"+err2.Error())
			}
			_ = rows2.Close()
			relist = append(relist,TypeList)
		}
	}else{
		panic("一级类查询失败"+err1.Error())
	}
	_ = rows1.Close()

	rjson,err := json.Marshal(relist)
	if err != nil{
		panic("数据装载出错，"+err.Error())
	}
	Rdata := utils.ResData{}
	Rdata.Footer = map_count[0]
	_= json.Unmarshal([]byte(rjson),&Rdata.Result)


	return Rdata,nil

}

//部门分析
func (ctl *ReportModel)Report_As_DeptCount_Browse(req beego.Controller)(result utils.ResData, rerr error){
	var conn = services.NewConnection(utils.GetAppid(req))
	var err error
	var map_count []map[string]interface{}
	var sqlstr string
	defer func(){
		_ = conn.Close()

		if err := recover();err!=nil{
			fmt.Println("Report_As_DeptCount_Browse查询失败",err)
			rerr = fmt.Errorf("%s",err)
		}
	}()

	sdeptid := req.GetString("deptid","") //管理部门ID
	suserid := req.GetString("userid","")
	sshowtype := req.GetString("sshowtype","")
	slbid := req.GetString("lbid","")
	//curPage := req.GetString("curPage","")
	//pageSize := req.GetString("pageSize","")
	if sshowtype == ""{
		sshowtype ="1"
	}

	sqlstr = "call  MSP_Report_Dept_Fx (0,"+utils.GetAppid(req)+",0,"+suserid+",\""+sdeptid+"\",\""+slbid+"\",'',"+sshowtype+",1,10000)"
	//查询合计数据
	map_count,err = services.FindToList(utils.GetAppid(req),conn,sqlstr)
	if err !=nil{
		panic("查询合计数据出错，"+err.Error())
	}
	//查询父级
	sqlstr = "call  MSP_Report_Dept_Fx (1,"+utils.GetAppid(req)+",?,"+suserid+",\""+sdeptid+"\",\""+slbid+"\",'',"+sshowtype+",1,10000)"
	rows1, err1 := conn.Query(sqlstr,0);
	//fmt.Println(strings.Replace(sqlstr, "?", "0", -1 ))
	var relist  []Type_List
	if err1 == nil{
		for rows1.Next(){
			TypeList := Type_List{}
			_ = rows1.Scan(&TypeList.ID,&TypeList.No, &TypeList.ParentID,&TypeList.Name,&TypeList.SPCount,&TypeList.SPTotal,&TypeList.UseQty,
				&TypeList.IdleQty,&TypeList.ScrappedQty,&TypeList.ServiceQty,&TypeList.DealQty,&TypeList.AcquireQty,&TypeList.BorrowQty)
			//查二层
			//fmt.Println("输出",TypeList)
			rows2, err2 := conn.Query(sqlstr,TypeList.ID);
			//fmt.Println(strings.Replace(sqlstr, "?", TypeList.ID, -1 ))
			if err2 ==nil{
				for rows2.Next(){
					TypeList2 := Type_List2{}
					_ = rows2.Scan(&TypeList2.ID,&TypeList2.No, &TypeList2.ParentID,&TypeList2.Name,&TypeList2.SPCount,&TypeList2.SPTotal,&TypeList2.UseQty,
						&TypeList2.IdleQty,&TypeList2.ScrappedQty,&TypeList2.ServiceQty,&TypeList2.DealQty,&TypeList2.AcquireQty,&TypeList2.BorrowQty)
					//查三层
					rows3, err3 := conn.Query(sqlstr,TypeList2.ID);
					//fmt.Println(strings.Replace(sqlstr, "?", TypeList2.ID, -1 ))
					if err3 ==nil{
						for rows3.Next() {
							TypeList3 := Type_List3{}
							_ = rows3.Scan(&TypeList3.ID,&TypeList3.No, &TypeList3.ParentID, &TypeList3.Name, &TypeList3.SPCount, &TypeList3.SPTotal, &TypeList3.UseQty,
								&TypeList3.IdleQty, &TypeList3.ScrappedQty, &TypeList3.ServiceQty, &TypeList3.DealQty, &TypeList3.AcquireQty, &TypeList3.BorrowQty)
							TypeList2.Subs = append(TypeList2.Subs,TypeList3)
						}
					}else {
						if utils.SqlLog {
							utils.LogOut("info","svrport = "+utils.GetAppid(req)+" 三级类查询失败 exec sql ",strings.Replace(sqlstr, "?", TypeList2.ID, -1 ))
						}
						panic("三级类查询失败"+err3.Error())
					}
					_ = rows3.Close()

					TypeList.Subs = append(TypeList.Subs,TypeList2)
				}
			} else{
				if utils.SqlLog {
					utils.LogOut("info","svrport = "+utils.GetAppid(req)+" 二级类查询失败 exec sql ",strings.Replace(sqlstr, "?", TypeList.ID, -1 ))
				}
				panic("二级类查询失败"+err2.Error())
			}
			_ = rows2.Close()
			relist = append(relist,TypeList)
		}
	}else{
		if utils.SqlLog {
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" 一级类查询失败 exec sql ",strings.Replace(sqlstr, "?", "0", -1 ))
		}
		panic("一级类查询失败"+err1.Error())
	}
	_ = rows1.Close()

	rjson,err := json.Marshal(relist)
	if err != nil{
		panic("数据装载出错，"+err.Error())
	}
	Rdata := utils.ResData{}
	Rdata.Footer = map_count[0]
	_= json.Unmarshal([]byte(rjson),&Rdata.Result)


	return Rdata,nil

}

//资产流水分析
func (ctl *ReportModel) Report_As_Change_Browse(req beego.Controller) (result utils.ResData, rerr error) {
	var reqs map[string]interface{};
	var sortmap map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))
	var acount string
	var icount int
	var swhere string
	var sqlstr string
	var sqlstr_ord string
	//var gjwhere string
	//response := utils.ResponsModel{};

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("Report_As_Change_Browse 失败，",err)
			rerr = fmt.Errorf("%s",err)
		}
	}()




	stype := req.GetString("atype","")  // 变更类型 0	登记,1	修改,2	领用,3	借用,4	归还,5	退库,6	维修,7	报废,8	处置,9	批量变更
	sbdate := req.GetString("bdate","")//+" 00:00:00"
	sedate := req.GetString("e_date","") //+" 23:59:59"
	suerid := req.GetString("userid","")
	svalue := req.GetString("avalue","")
	sloginid := req.GetString("loginid","")
	pageindex := req.GetString("curPage","")//reqs["pageindex"]
	pagesize := req.GetString("pageSize","")//reqs["pagesize"]

	if (sbdate =="") ||(sedate ==""){
		panic("请正确输入查询日期！")
	}

	sbdate = sbdate+" 00:00:00"
	sedate = sedate +" 23:59:59"
	swhere = " and f_create_time >= '"+sbdate+ "' and f_create_time <= '"+sedate+"' "
	if stype != ""{
		swhere = swhere + " and f_type = "+stype
	}
	if suerid ==""{
		swhere = swhere + " and f_opruserid = "+suerid
	}
	if svalue != ""{
		swhere =swhere+" and (f_code ='"+svalue+"' or f_name like '%"+svalue+"%')"
	}

	//SpreadWhere := req.GetString("SpreadWhere");
	//if SpreadWhere != ""{ //高级筛选处理
	//	gjwhere = ""
	//	rjson,err := simplejson.NewJson([]byte(SpreadWhere));
	//	if err == nil{
	//		if rjson.Get("optName").MustString() != ""{
	//			gjwhere = gjwhere+" and m.f_userid in (SELECT id FROM tbl_user tu WHERE tu.f_customid = "+utils.GetAppid(req)+" AND tu.f_flag = 0 and f_name LIKE '%"+rjson.Get("optName").MustString()+"%')"
	//		}
	//		if rjson.Get("f_modal").MustString() != ""{ //变更后所属部门
	//			gjwhere = gjwhere+" and m.f_modal = '"+rjson.Get("f_modal").MustString()+"'"
	//		}
	//		if rjson.Get("f_oprtype") .MustString() != ""{
	//			gjwhere = gjwhere+" and  m.f_oprtype = '"+rjson.Get("f_oprtype") .MustString()+"'"
	//		}
	//		if rjson.Get("f_note") .MustString() != ""{
	//			gjwhere = gjwhere+" and  m.f_note like '%"+rjson.Get("f_note") .MustString()+"%'"
	//		}
	//		if (rjson.Get("b_date") .MustString() != "")&&(rjson.Get("e_date") .MustString() != ""){
	//			gjwhere = gjwhere+" and m.f_oprtime >= '"+rjson.Get("b_date").MustString()+" 00:00:00"+
	//				"' and m.f_oprtime <= '"+rjson.Get("e_date").MustString()+" 23:59:59"+"'"
	//		}
	//		gjwhere = gjwhere +" "
	//	}
	//}


	//fmt.Println(deptno)


	fmt.Println(swhere)
	sqlstr_ord =  " order by id desc "
	if reqs["sort"] != nil{
		sortmap = reqs["sort"].(map[string]interface{})
		if (sortmap["sortprop"] !=nil) &&(sortmap["order"] !=nil) {
			var ord string
			if sortmap["order"].(string) =="descending"{
				ord = " desc"
			} else {
				ord =" asc"
			}
			sqlstr_ord = " order by "+sortmap["id"].(string)+" "+ord
		}
	}
	sqlstr =" call MSP_Report_SPWater(0,"+sloginid+","+utils.GetAppid(req)+",\""+swhere+"\","+pageindex+","+pagesize+",\""+sqlstr_ord+"\")"

	acount,rerr = services.Get_Procedure_Data(utils.GetAppid(req),conn,sqlstr,"icount")

	if rerr != nil{
		panic("查询合计数出错，"+rerr.Error())
	}

	icount,rerr = strconv.Atoi(acount)
	if rerr != nil{
		panic("获取合计数出错 ，"+rerr.Error())
	}
	sqlstr =" call MSP_Report_SPWater(1,"+sloginid+","+utils.GetAppid(req)+",\""+swhere+"\","+pageindex+","+pagesize+",\""+sqlstr_ord+"\")"
	result.Result,rerr = services.FindToList(utils.GetAppid(req),conn,sqlstr )
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
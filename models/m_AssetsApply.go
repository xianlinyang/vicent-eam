package models

import (
	"Eam_Server/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	"strconv"
	"time"
	"Eam_Server/services"
)

type AssetsApplyModel struct {
	Base_Model
}


//领用申请查询
func (ctl *AssetsApplyModel) AssetsApply_Main_GetList(req beego.Controller) (result utils.ResData, rerr error) {
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))
	var sortmap map[string]interface{};
	var acount string
	var icount int
	var sqlstr_ord string
	var sqlstr string
	var gjwhere string //高级条件
	result = utils.ResData{};

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("AssetsApply_Main_GetList 失败，",err)
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
			//rerr = fmt.Errorf("%s",err)
			return
		}
	}()

	wvalue := req.GetString("w_value","")
	pageindex := req.GetString("curPage","")//reqs["pageindex"]
	pagesize := req.GetString("pageSize","")//reqs["pagesize"]
	suserid := req.GetString("userid","")
	sstatus := req.GetString("status","")// 0:全部 0	草稿，1	待审批，2	待确认，3	已确认，	8	已完结
	sort := req.GetString("sort");
	SpreadWhere := req.GetString("SpreadWhere"); //高级筛选条件

	if SpreadWhere != ""{ //高级筛选处理
		gjwhere = ""
		rjson,err := simplejson.NewJson([]byte(SpreadWhere));
		if err == nil{
			if rjson.Get("f_status").MustString() != ""{
				gjwhere = gjwhere+" and m.f_status ="+rjson.Get("f_status").MustString()
			}
			if rjson.Get("f_apply_deptid") .MustString() != ""{
				gjwhere = gjwhere+" and m.f_re_deptid ="+rjson.Get("f_apply_deptid").MustString()
			}
			if rjson.Get("f_apply_username") .MustString() != ""{
				gjwhere = gjwhere+" and m.f_re_userid in (SELECT id FROM tbl_user tu WHERE tu.f_customid = "+utils.GetAppid(req)+" AND tu.f_flag = 0 and f_name LIKE '%"+rjson.Get("f_apply_username").MustString()+"%')"
			}
			if (rjson.Get("f_apply_bdate") .MustString() != "") && (rjson.Get("f_apply_edate") .MustString() != ""){
				gjwhere = gjwhere+" and m.f_re_date >= '"+rjson.Get("f_apply_bdate").MustString()+" 00:00:00"+
					"' and m.f_re_date <= '"+rjson.Get("f_apply_edate").MustString()+" 23:59:59"+"'"
			}
			gjwhere = gjwhere +" "
		}
	}else {
		gjwhere =""
		if sstatus != ""{
			gjwhere = " m.f_status = "+sstatus
		}
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
			sqlstr_ord = " order by m."+sortmap["sortprop"].(string)+" "+ord
		}
	} else {
		sqlstr_ord = " order by m.id desc "
	}




	sqlstr = "CALL Msp_BillMain_Query (0,0,\""+wvalue+"\","+suserid+",\""+gjwhere+"\","+pageindex+","+pagesize+",\""+sqlstr_ord+"\")"


	acount,rerr = services.Get_Procedure_Data(utils.GetAppid(req),conn,sqlstr,"icount")
	if rerr != nil{
		panic("查询合计数出错，"+rerr.Error())
	}

	icount,rerr = strconv.Atoi(acount)
	if rerr != nil{
		panic("获取合计数出错 ，"+rerr.Error())
	}

	sqlstr = "CALL Msp_BillMain_Query (0,1,\""+wvalue+"\","+suserid+",\""+gjwhere+"\","+pageindex+","+pagesize+",\""+sqlstr_ord+"\")"

	result.Result,rerr = services.FindToList(utils.GetAppid(req),conn,sqlstr)
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
//领用申请查询
func (ctl *AssetsApplyModel) AssetsApply_DetailQuery(req beego.Controller) (result utils.ResData, rerr error) {
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))

	var acount string
	var icount int

	var sqlstr string
	result = utils.ResData{};

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("AssetsApply_DetailQuery 失败，",err)
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
			return
		}
	}()

	smid := req.GetString("mid","")
	userid := req.GetString("userid","")//reqs["pageindex"]


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
	//		sqlstr_ord = " order by m."+sortmap["sortprop"].(string)+" "+ord
	//	}
	//} else {
	//	sqlstr_ord = " order by m.id desc "
	//}



	sqlstr = "CALL Msp_BillDetail_Query (0,"+smid+","+userid+",0)"


	acount,rerr = services.Get_Procedure_Data(utils.GetAppid(req),conn,sqlstr,"icount")
	if rerr != nil{
		panic("查询合计数出错，"+rerr.Error())
	}

	icount,rerr = strconv.Atoi(acount)
	if rerr != nil{
		panic("获取合计数出错 ，"+rerr.Error())
	}

	sqlstr = "CALL Msp_BillDetail_Query (0,"+smid+","+userid+",1)"
	result.Result,rerr = services.FindToList(utils.GetAppid(req),conn,sqlstr)
	if rerr!=nil{
		panic("查询失败,"+rerr.Error())
	}else{
		result.PageSize = icount
		result.CurPage =1
		result.Totals = icount
		result.PageSize = 0
		result.CurPage = 0
		result.Totals = icount
		if reqs["sort"] != nil{
			result.Sort = reqs["sort"].(map[string]interface{})
		}
	}
	//fmt.Println(result)
	return result,nil

}
//领用申请新增
func (ctl *AssetsApplyModel) AssetsApply_Create(req beego.Controller)(aid string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接
	var s string
	var d_total float64

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
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	json_str,err:= json.Marshal(reqs)
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}


	rjson, err := simplejson.NewJson([]byte(json_str));

	if err != nil{
		panic("传入参数错误，JSON转换出错："+err.Error())
	}

	//if (reqs["f_re_date"] != nil) &&(reqs["f_re_date"].(string) != ""){
	//	dt, err1 := time.Parse("2006-01-02",reqs["f_re_date"].(string))
	//	reqs["f_re_date"] =dt
	//	fmt.Println(err1)
	//}
	//先做数据检测
	rows_chk, err := rjson.Get("splist").Array();
	if err == nil{
		for _, row := range rows_chk {
			if each_map, ok := row.(map[string] interface{}); ok {
				if each_map["f_qty"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 数量不能为空！")
				}
				if each_map["f_price"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 价值不能为空！")
				}
				if each_map["f_public"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 公用属性不能为空！")
				}
			}
		}
	}
	if (rows_chk==nil) || (len(rows_chk)==0){
		panic("不能保存空单据！")
	}

	b_main := make(map[string] interface{});
	//获了单号
	b_main["f_accountid"] = "1"
	s,err = services.Get_Procedure_Data(utils.GetAppid(req),conn," CALL Msp_Get_Serno ("+utils.GetAppid(req)+",'tbl_receive_apply_main','f_serno','SL') ","f_serno")
	if err != nil{
		panic("单号获取出错,"+err.Error())
	}
	b_main["f_serno"] =s
	if rjson.Get("f_status").MustString() ==""{
		b_main["f_status"] = "0"
	} else {
		b_main["f_status"] = rjson.Get("f_status").MustString()
	}

	if rjson.Get("f_re_userid").MustString() !=""{
		b_main["f_re_userid"] = rjson.Get("f_re_userid").MustString()
	}
	if rjson.Get("f_apply_userid").MustString() !=""{
		b_main["f_apply_userid"] = rjson.Get("f_apply_userid").MustString()
	}
	if rjson.Get("f_return_date").MustString() !=""{
		b_main["f_return_date"] = rjson.Get("f_return_date").MustString()
	}
	b_main["f_re_deptid"] = rjson.Get("f_re_deptid").MustString()
	if rjson.Get("f_re_postion") != nil{
		b_main["f_re_postion"] = rjson.Get("f_re_postion").MustString()
	} else{
		b_main["f_re_postion"] = "-1";
	}


	b_main["f_apply_reason"] = rjson.Get("f_apply_reason").MustString()
	b_main["f_note"] = rjson.Get("f_note").MustString()
	b_main["f_create_user"] = rjson.Get("f_create_user").MustString()
	now := time.Now()
	b_main["f_create_time"]= now.Format("2006-01-02 15:04:05");
	b_main["f_customid"] = utils.GetAppid(req)
	if rjson.Get("f_re_date").MustString() != ""{
		b_main["f_re_date"] = rjson.Get("f_re_date").MustString()
	}

	if rjson.Get("f_nextaudit_user").MustString() == ""{
		b_main["f_nextaudit_user"]= "-1";
	} else {
		b_main["f_nextaudit_user"]= rjson.Get("f_nextaudit_user").MustString();
	}

	if b_main["f_status"].(string) =="1"{
		if b_main["f_nextaudit_user"].(string) =="-1"{
			panic("提交申请时必须指定下一审批人！")
		}
	}


	//开始事务
	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	//插入主单
	aid,err = services.One_ADD_Table(utils.GetAppid(req),tx,conn,"tbl_Receive_Apply_Main",b_main)

	//插入从单
	if err ==nil{
		rows, _ := rjson.Get("splist").Array();
		qxsql :=" INSERT INTO tbl_Receive_Apply_Detail(f_mid,f_spid,f_qty,f_price,f_total,f_custody_deptid,f_custody_userid,f_postion,f_use_deptid,f_public,f_use_userid,"+
		        " 									   f_note,f_create_user,f_create_time,f_customid) "+
		        " VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,now(),?)"
		for _, row := range rows {
			if each_map, ok := row.(map[string] interface{}); ok {

				if each_map["f_price"].(string) == ""{
					each_map["f_price"] = "0"
				}
				if each_map["f_custody_deptid"].(string) == ""{
					each_map["f_custody_deptid"] = "-1"
				}
				if each_map["f_custody_userid"].(string) == ""{
					each_map["f_custody_userid"] = "-1"
				}
				if each_map["f_postion"].(string) == ""{
					each_map["f_postion"] = "-1"
				}
				if each_map["f_use_deptid"].(string) == ""{
					each_map["f_use_deptid"] = "-1"
				}
				if each_map["f_public"].(string) == ""{
					each_map["f_public"] = "0"
				}
				if each_map["f_use_userid"].(string) == ""{
					each_map["f_use_userid"] = "0"
				}
				d_qty,_ := strconv.ParseFloat(each_map["f_qty"].(string),64)
				d_price,_ := strconv.ParseFloat(each_map["f_price"].(string),64)
				d_total = d_qty * d_price
				//f_mid,f_spid,f_qty,f_price,f_total,f_custody_deptid,f_custody_userid,f_postion,f_use_deptid,f_public,f_use_userid,"+
				//" 									   f_note,f_create_user
				_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,qxsql,aid,each_map["f_spid"].(string),each_map["f_qty"].(string),each_map["f_price"].(string),
						d_total,each_map["f_custody_deptid"].(string),each_map["f_custody_userid"].(string),each_map["f_postion"].(string),each_map["f_use_deptid"].(string),
						each_map["f_public"].(string),each_map["f_use_userid"].(string),each_map["f_note"].(string),rjson.Get("f_create_user").MustString(),utils.GetAppid(req))
				if err !=nil{
					break;
				}
			}
		}
	}

	if err ==nil{
		if (rjson.Get("f_status").MustString() =="0")||(rjson.Get("f_status").MustString() ==""){
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产申请",b_main["f_create_user"].(string),aid,"新增草稿单","")
			//atypename string,auserid int,aid int,aopr string,amemo string
		}else {
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产申请",b_main["f_create_user"].(string),aid,"新增草稿单","")
			if err ==nil {
				err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产申请",b_main["f_create_user"].(string),aid,"提交申请","")
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
		_= services.DB_Log(utils.GetAppid(req),conn,"领用申请","新增","",reqs["f_create_user"].(string),"新增失败",err)
		if utils.ISLog{
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI,"新增失败，"+rerr.Error())
		}
	} else {
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"领用申请","新增",reqs["f_create_user"].(string),reqs["f_create_user"].(string),"新增成功 ID ="+aid,err)
	}
	return

}
//领用申请修改
func (ctl *AssetsApplyModel) AssetsApply_Update(req beego.Controller)(alen string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接
	var d_total float64
	fmt.Println("ss")
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
			alen  =""
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

	if err != nil{
		panic("传入参数错误，JSON转换出错："+err.Error())
	}
	//先做数据检测
	rows_chk, err := rjson.Get("splist").Array();
	if err == nil{
		for _, row := range rows_chk {
			if each_map, ok := row.(map[string] interface{}); ok {
				if each_map["f_qty"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 数量不能为空！")
				}
				if each_map["f_price"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 价值不能为空！")
				}
				if each_map["f_public"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 公用属性不能为空！")
				}
			}
		}
	}


	d_update_main := make(map[string] string);
	//获了单号
	d_update_main["id"] = rjson.Get("id").MustString()

	//检查单据状态
	err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_Receive_Apply_Main","f_status",d_update_main["id"],"0")
	if err != nil {
		fmt.Println(err)
		panic("单据已生效不允许修改！")
	}

	d_update_main["f_accountid"] = rjson.Get("f_accountid").MustString()
	d_update_main["f_status"]  = rjson.Get("f_status").MustString()
	d_update_main["f_serno"] =  rjson.Get("f_serno").MustString()
	//d_update_main["f_status"] =  rjson.Get("f_status").MustString()
	d_update_main["f_re_deptid"] =  rjson.Get("f_re_deptid").MustString()
	if rjson.Get("f_re_postion") != nil{
		d_update_main["f_re_postion"] = rjson.Get("f_re_postion").MustString()
	} else{
		d_update_main["f_re_postion"] = "-1";
	}
	d_update_main["f_re_userid"] =  rjson.Get("f_re_userid").MustString()
	d_update_main["f_apply_userid"] =  rjson.Get("f_apply_userid").MustString()
	if rjson.Get("f_return_date").MustString() ==""{
		d_update_main["f_return_date"] = utils.Data_null //设置为NULL 值
	}else{
		d_update_main["f_return_date"] =  rjson.Get("f_return_date").MustString()
	}

	if rjson.Get("f_re_date").MustString() ==""{
		d_update_main["f_re_date"] = utils.Data_null //设置为NULL 值
	}else{
		d_update_main["f_re_date"] =  rjson.Get("f_re_date").MustString()
	}
	//d_update_main["f_audit_date"] =  rjson.Get("f_audit_date").MustString()
	d_update_main["f_note"] =  rjson.Get("f_note").MustString()
	//d_update_main["f_nextaudit_user"] =  rjson.Get("f_nextaudit_user").MustString()
	//d_update_main["f_create_user"] =  rjson.Get("f_create_user").MustString()
	//d_update_main["f_create_time"] =  rjson.Get("f_create_time").MustString()
	d_update_main["f_update_user"] =  rjson.Get("f_update_user").MustString()
	d_update_main["f_update_time"] =  time.Now().Format("2006-01-02 15:04:05")
	d_update_main["f_apply_reason"] =  rjson.Get("f_apply_reason").MustString()
	if rjson.Get("f_nextaudit_user").MustString() != ""{
		d_update_main["f_nextaudit_user"] =  rjson.Get("f_nextaudit_user").MustString()
	}

	if (d_update_main["f_status"] =="1") && ( (d_update_main["f_nextaudit_user"]=="") || (d_update_main["f_nextaudit_user"]=="-1") ){
		snextid,err1 := services.Get_FieldByValue(utils.GetAppid(req),conn,"SELECT f_nextaudit_user FROM tbl_receive_apply_main where   id ="+d_update_main["id"],"f_nextaudit_user")
		if err1 ==nil{
			if (snextid =="-1") || (snextid ==""){
				panic("提交申请时必须指定下一审批人！")
			}
		}else {
			panic("下一审批人获取失败,"+err1.Error())
		}
	}

	sqlstr :=  services.Sql_JionSet(d_update_main)
	if sqlstr !="" {
		sqlstr = " update tbl_Receive_Apply_Main set "+sqlstr +" where id ="+d_update_main["id"]
	}

	//开始事务
	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	//修改角色
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)

	if err == nil{
		//先删除从单
		sqlstr = " delete from tbl_Receive_Apply_Detail where f_mid =  "+d_update_main["id"]
		_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)
	}

	//插入从单
	if err ==nil{
		rows, _ := rjson.Get("splist").Array();
		qxsql :=" INSERT INTO tbl_Receive_Apply_Detail(f_mid,f_spid,f_qty,f_price,f_total,f_custody_deptid,f_custody_userid,f_postion,f_use_deptid,f_public,f_use_userid,"+
			" 									   f_note,f_create_user,f_create_time,f_customid) "+
			" VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,now(),?)"
		for _, row := range rows {
			if each_map, ok := row.(map[string] interface{}); ok {

				if each_map["f_price"].(string) == ""{
					each_map["f_price"] = "0"
				}
				if each_map["f_custody_deptid"].(string) == ""{
					each_map["f_custody_deptid"] = "-1"
				}
				if each_map["f_custody_userid"].(string) == ""{
					each_map["f_custody_userid"] = "-1"
				}
				if each_map["f_postion"].(string) == ""{
					each_map["f_postion"] = "-1"
				}
				if each_map["f_use_deptid"].(string) == ""{
					each_map["f_use_deptid"] = "-1"
				}
				if each_map["f_public"].(string) == ""{
					each_map["f_public"] = "0"
				}
				if each_map["f_use_userid"].(string) == ""{
					each_map["f_use_userid"] = "0"
				}
				d_qty,_ := strconv.ParseFloat(each_map["f_qty"].(string),64)
				d_price,_ := strconv.ParseFloat(each_map["f_price"].(string),64)
				d_total = d_qty * d_price
				//f_mid,f_spid,f_qty,f_price,f_total,f_custody_deptid,f_custody_userid,f_postion,f_use_deptid,f_public,f_use_userid,"+
				//" 									   f_note,f_create_user
				_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,qxsql,d_update_main["id"],each_map["f_spid"].(string),each_map["f_qty"].(string),each_map["f_price"].(string),
					d_total,each_map["f_custody_deptid"].(string),each_map["f_custody_userid"].(string),each_map["f_postion"].(string),each_map["f_use_deptid"].(string),
					each_map["f_public"].(string),each_map["f_use_userid"].(string),each_map["f_note"].(string),rjson.Get("f_update_user").MustString(),utils.GetAppid(req))
				if err !=nil{
					break;
				}
			}
		}
	}

	if err==nil{
		if (d_update_main["f_status"] =="0")||(d_update_main["f_status"] ==""){
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产申请",d_update_main["f_update_user"],d_update_main["id"],"草稿单修改","")
			//atypename string,auserid int,aid int,aopr string,amemo string
		}else {
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产申请",d_update_main["f_update_user"],d_update_main["id"],"草稿单修改","")
			if err ==nil {
				err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产申请",d_update_main["f_update_user"],d_update_main["id"],"提交申请","")
			}
		}

	}

	if (err != nil)  {
		rerr = err
		err = tx.Rollback()
		if err == nil{
			fmt.Println("回滚完成")
		} else {
			fmt.Println(err)
		}
		_= services.DB_Log(utils.GetAppid(req),conn,"领用申请","修改","",reqs["f_update_user"].(string),"修改失败",err)

		if utils.ISLog{
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI,"修改失败，"+rerr.Error())
		}
	} else {
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"领用申请","修改",reqs["f_update_user"].(string),reqs["f_update_user"].(string),"修改成功",err)
	}
	return
}
//领用申请删除
func (ctl *AssetsApplyModel) AssetsApply_Delete(req beego.Controller)(rerr error){
	var reqs map[string]interface{};
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var conn = services.NewConnection(utils.GetAppid(req))
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

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

	//检查单据状态
	err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_receive_apply_main","f_status",reqs["id"].(string),"0")
	if err != nil {
		fmt.Println(err)
		panic("单据已生效不允许删除！")
	}

	tx,_ := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	sqlstr := "delete from tbl_Receive_Apply_Main where id = ? "

	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr,reqs["id"].(string))
	//删除从单
	if err == nil{
		sqlstr = " DELETE from tbl_Receive_Apply_Detail WHERE f_mid ="+reqs["id"].(string)
		_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr,)
	}

	if err !=nil{
		rerr = err
		rerr = tx.Rollback()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产申请","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除失败",err)
	} else{

		rerr = nil
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产申请","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除成功 单号 ="+reqs["f_serno"].(string),err)
	}
	return
}

//领用申请 状态操作
func (ctl *AssetsApplyModel) AssetsApply_SetStatus(req beego.Controller,atype string)(rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接
	var newstate string
	var billMemo string
	var spyj     string

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
			return
		}
	}()
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	stype := atype  //  2主管审批  3仓库审批   -3仓库辙消审批   -2主管取消审批  -1辙消提交


	if (stype=="2") &&((reqs["f_nextaudit_user"] ==nil) ||(reqs["f_nextaudit_user"].(string) =="")){
		panic("下一审批人必须提定！")
	}

	if reqs["spyj"] != nil{
		spyj = reqs["spyj"].(string)  //1审批通过 0审批通过，驳回
	} else {
		spyj ="0"
	}




	//状态检测
	if stype == "2"{  //主管审批
		err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_Receive_Apply_Main","f_status",reqs["id"].(string),"1")
		if err != nil {
			fmt.Println(err)
			panic("单据不在待审批状态不能审批！")
		}
	}
	if stype == "3"{  //仓库审批
		err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_Receive_Apply_Main","f_status",reqs["id"].(string),"2")
		if err != nil {
			fmt.Println(err)
			panic("单据不在待确认状态不能审批！")
		}
	}
	if stype == "-1"{  //辙消提交
		err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_Receive_Apply_Main","f_status",reqs["id"].(string),"1")
		if err != nil {
			fmt.Println(err)
			panic("单据不在待审批状态不能辙消！")
		}
	}
	if stype == "-2"{  //主管取消审批
		err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_Receive_Apply_Main","f_status",reqs["id"].(string),"2")
		if err != nil {
			fmt.Println(err)
			panic("单据不在待确认状态不能辙消！")
		}
	}
	if stype == "-3"{ //仓库辙消审批
		err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_Receive_Apply_Main","f_status",reqs["id"].(string),"3")
		if err != nil {
			fmt.Println(err)
			panic("单据不在确认状态不能辙消！")
		}
	}
	switch stype{  //2主管审批  3仓库审批   -3仓库辙消审批   -2主管取消审批  -1辙消提交
	case "2":{ if spyj =="1"{
					newstate = "2"
					billMemo	= "主管审批"
				} else { newstate ="-1"
				         billMemo ="主管驳回"}
	            }
	case "3":{ if spyj == "1"{
					newstate = "3"
					billMemo	= "仓库审批"
			   }else{
					newstate = "-1"
					billMemo	= "仓库驳回"
				} }
	case "-3":{ newstate = "2"
			    billMemo = "仓为辙消审批"}
	case "-2":{newstate = "1"
		       billMemo = "主管辙消审批"}
	case "-1":{newstate ="0"
		       billMemo = "辙消提交"}
	default:
		newstate =""
	}

	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,"update tbl_Receive_Apply_Main set f_status = ? where id =? ",newstate,reqs["id"].(string))
	if (err ==nil) &&(stype =="2"){
		_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,"update tbl_Receive_Apply_Main set f_nextaudit_user = ? where id =? ",reqs["f_nextaudit_user"].(string),reqs["id"].(string))
	}
	if err == nil{
		err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产申请",reqs["userid"].(string),reqs["id"].(string),billMemo,reqs["memo"].(string))
	}
	if err == nil{
		err= services.DB_Log(utils.GetAppid(req),conn,"领用申请",billMemo,"",reqs["userid"].(string),"操作成功",err)
	}

	if err ==nil{
		_ = tx.Commit()
	}else {
		_ = tx.Rollback();
		rerr = err
		if utils.ISLog{
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI,"操作失败，"+rerr.Error())
		}
	}

	return rerr
}


//领用查询
func (ctl *AssetsApplyModel) Receive_Main_GetList(req beego.Controller) (result utils.ResData, rerr error) {
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))
	var sortmap map[string]interface{};
	var acount string
	var icount int
	var sqlstr_ord string
	var sqlstr string
	var gjwhere string
	result = utils.ResData{};

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("Receive_Main_GetList 失败，",err)
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
		}
	}()

	wvalue := req.GetString("w_value","")
	pageindex := req.GetString("curPage","")//reqs["pageindex"]
	pagesize := req.GetString("pageSize","")//reqs["pagesize"]
	userid := req.GetString("userid","")//reqs["pagesize"]
	sort := req.GetString("sort");
	status := req.GetString("status");
	SpreadWhere := req.GetString("SpreadWhere"); //高级筛选条件

	if SpreadWhere != ""{ //高级筛选处理
		gjwhere = ""
		rjson,err := simplejson.NewJson([]byte(SpreadWhere));
		if err == nil{
			if rjson.Get("f_status").MustString() != ""{
				gjwhere = gjwhere+" and m.f_status ="+rjson.Get("f_status").MustString()
			}
			if rjson.Get("f_re_deptid") .MustString() != ""{
				gjwhere = gjwhere+" and m.f_re_deptid ="+rjson.Get("f_re_deptid").MustString()
			}
			if rjson.Get("f_re_username") .MustString() != ""{
				gjwhere = gjwhere+" and m.f_re_userid in (SELECT id FROM tbl_user tu WHERE tu.f_customid = "+utils.GetAppid(req)+" AND tu.f_flag = 0 and f_name LIKE '%"+rjson.Get("f_re_username").MustString()+"%')"
			}
			if (rjson.Get("f_re_bdate") .MustString() != "")&&(rjson.Get("f_re_edate") .MustString() != ""){
				gjwhere = gjwhere+" and m.f_re_date >= '"+rjson.Get("f_re_bdate").MustString()+" 00:00:00"+
					"' and m.f_re_date <= '"+rjson.Get("f_re_edate").MustString()+" 23:59:59"+"'"
			}
			gjwhere = gjwhere +" "
		}
	}else {
		gjwhere =""
		if status != ""{
			gjwhere = " m.f_status = "+status
		}
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
			sqlstr_ord = " order by m."+sortmap["sortprop"].(string)+" "+ord
		}
	} else {
		sqlstr_ord = " order by m.id desc "
	}


	sqlstr = "CALL Msp_BillMain_Query (1,0,\""+wvalue+"\","+userid+",\""+gjwhere+"\","+pageindex+","+pagesize+",\""+sqlstr_ord+"\")"


	acount,rerr = services.Get_Procedure_Data(utils.GetAppid(req),conn,sqlstr,"icount")
	if rerr != nil{
		panic("查询合计数出错，"+rerr.Error())
	}

	icount,rerr = strconv.Atoi(acount)
	if rerr != nil{
		panic("获取合计数出错 ，"+rerr.Error())
	}

	sqlstr = "CALL Msp_BillMain_Query (1,1,\""+wvalue+"\","+userid+",\""+gjwhere+"\","+pageindex+","+pagesize+",\""+sqlstr_ord+"\")"

	result.Result,rerr = services.FindToList(utils.GetAppid(req),conn,sqlstr)
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
//领用查询
func (ctl *AssetsApplyModel) Receive_DetailQuery(req beego.Controller) (result utils.ResData, rerr error) {
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))

	var acount string
	var icount int

	var sqlstr string
	result = utils.ResData{};

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("Receive_DetailQuery 失败，",err)
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
		}
	}()

	smid := req.GetString("mid","")
	userid := req.GetString("userid","")//reqs["pageindex"]


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
	//		sqlstr_ord = " order by m."+sortmap["sortprop"].(string)+" "+ord
	//	}
	//} else {
	//	sqlstr_ord = " order by m.id desc "
	//}



	sqlstr = "CALL Msp_BillDetail_Query (1,"+smid+","+userid+",0)"


	acount,rerr = services.Get_Procedure_Data(utils.GetAppid(req),conn,sqlstr,"icount")
	if rerr != nil{
		panic("查询合计数出错，"+rerr.Error())
	}

	icount,rerr = strconv.Atoi(acount)
	if rerr != nil{
		panic("获取合计数出错 ，"+rerr.Error())
	}

	sqlstr = "CALL Msp_BillDetail_Query (1,"+smid+","+userid+",1)"
	result.Result,rerr = services.FindToList(utils.GetAppid(req),conn,sqlstr)
	if rerr!=nil{
		panic("查询失败,"+rerr.Error())
	}else{
		result.PageSize = icount
		result.CurPage =1
		result.Totals = icount
		result.PageSize = 0
		result.CurPage = 0
		result.Totals = icount
		if reqs["sort"] != nil{
			result.Sort = reqs["sort"].(map[string]interface{})
		}
	}
	//fmt.Println(result)
	return result,nil

}
//领用新增
func (ctl *AssetsApplyModel) Receive_Create(req beego.Controller)(aid string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接
	var s string
	var d_total float64
	fmt.Println("ss")
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
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	json_str,err:= json.Marshal(reqs)
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}
	rjson, err := simplejson.NewJson([]byte(json_str));

	if err != nil{
		panic("传入参数错误，JSON转换出错："+err.Error())
	}
	//先做数据检测
	rows_chk, err := rjson.Get("splist").Array();
	if err == nil{
		for _, row := range rows_chk {
			if each_map, ok := row.(map[string] interface{}); ok {
				if each_map["f_qty"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 数量不能为空！")
				}
				if each_map["f_price"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 价值不能为空！")
				}
				if each_map["f_public"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 公用属性不能为空！")
				}
			}
		}
	}

	if (rows_chk==nil) || (len(rows_chk)==0){
		panic("不能保存空单据！")
	}


	b_main := make(map[string] interface{});
	//获了单号
	b_main["f_accountid"] = "1"
	s,err = services.Get_Procedure_Data(utils.GetAppid(req),conn," CALL Msp_Get_Serno ("+utils.GetAppid(req)+",'tbl_receive_main','f_serno','LY') ","f_serno")
	if err != nil{
		panic("单号获取出错,"+err.Error())
	}
	b_main["f_serno"] =s
	if rjson.Get("f_status").MustString() == ""{
		b_main["f_status"] = "0"
		b_main["f_audit"] = "0"
	} else {
		b_main["f_status"] = rjson.Get("f_status").MustString()
		if rjson.Get("f_status").MustString()=="1"{
			b_main["f_audit"] = "1"
			now := time.Now()
			b_main["f_audit_date"]= now.Format("2006-01-02 15:04:05");
		} else {
			b_main["f_audit"] = "0"
		}

	}

	if rjson.Get("f_re_userid").MustString() !=""{
		b_main["f_re_userid"] = rjson.Get("f_re_userid").MustString()
	}
	b_main["f_apply_billid"] = rjson.Get("f_apply_billid").MustString()

	if rjson.Get("f_re_date").MustString() !=""{
		b_main["f_re_date"] = rjson.Get("f_re_date").MustString()
	}
	b_main["f_re_deptid"] = rjson.Get("f_re_deptid").MustString()
	b_main["f_re_postion"] = rjson.Get("f_re_postion").MustString()
	b_main["f_note"] = rjson.Get("f_note").MustString()
	b_main["f_create_user"] = rjson.Get("f_create_user").MustString()
	now := time.Now()
	b_main["f_create_time"]= now.Format("2006-01-02 15:04:05");
	b_main["f_customid"]= utils.GetAppid(req)





	//开始事务
	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	//插入主单
	aid,err = services.One_ADD_Table(utils.GetAppid(req),tx,conn,"tbl_Receive_Main",b_main)

	//插入从单
	if err ==nil{
		rows, _ := rjson.Get("splist").Array();
		qxsql :=" INSERT INTO tbl_Receive_Detail(f_mid,f_spid,f_qty,f_price,f_total,f_custody_deptid,f_custody_userid,f_postion,f_use_deptid,f_public,f_use_userid,"+
			" 									   f_note,f_create_user,f_create_time,f_customid) "+
			" VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,now(),?)"
		for _, row := range rows {
			if each_map, ok := row.(map[string] interface{}); ok {

				if each_map["f_price"].(string) == ""{
					each_map["f_price"] = "0"
				}
				if each_map["f_custody_deptid"].(string) == ""{
					each_map["f_custody_deptid"] = "-1"
				}
				if each_map["f_custody_userid"].(string) == ""{
					each_map["f_custody_userid"] = "-1"
				}
				if each_map["f_postion"].(string) == ""{
					each_map["f_postion"] = "-1"
				}
				if each_map["f_use_deptid"].(string) == ""{
					each_map["f_use_deptid"] = "-1"
				}
				if each_map["f_public"].(string) == ""{
					each_map["f_public"] = "0"
				}
				if each_map["f_use_userid"].(string) == ""{
					each_map["f_use_userid"] = "0"
				}
				d_qty,_ := strconv.ParseFloat(each_map["f_qty"].(string),64)
				d_price,_ := strconv.ParseFloat(each_map["f_price"].(string),64)
				d_total = d_qty * d_price
				//f_mid,f_spid,f_qty,f_price,f_total,f_custody_deptid,f_custody_userid,f_postion,f_use_deptid,f_public,f_use_userid,"+
				//" 									   f_note,f_create_user
				_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,qxsql,aid,each_map["f_spid"].(string),each_map["f_qty"].(string),each_map["f_price"].(string),
					d_total,each_map["f_custody_deptid"].(string),each_map["f_custody_userid"].(string),each_map["f_postion"].(string),each_map["f_use_deptid"].(string),
					each_map["f_public"].(string),each_map["f_use_userid"].(string),each_map["f_note"].(string),rjson.Get("f_create_user").MustString(),utils.GetAppid(req))
				if err !=nil{
					break;
				}
			}
		}
	}

	// 新增成功后 资产位置要更改到领用部门
	if (err == nil) && (b_main["f_status"] == "1"){
		var aparam = []string{"1", aid,b_main["f_create_user"].(string),utils.GetAppid(req)}
		var rdata map[string]interface{}

		if err ==nil{
			fmt.Println("调用")
			rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
			fmt.Println(rdata)
			if (err !=nil) || (rdata["f_state"] !="1") {
				err = fmt.Errorf("%s","库存变动失败,"+err.Error())
			}
		}
	}

	if err == nil{
		if (rjson.Get("f_status").MustString() =="0")||(rjson.Get("f_status").MustString() ==""){
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产领用",b_main["f_create_user"].(string),aid,"新增草稿单","")
			//atypename string,auserid int,aid int,aopr string,amemo string
		}else {
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产领用",b_main["f_create_user"].(string),aid,"新增草稿单","")
			if err ==nil {
				err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产领用",b_main["f_create_user"].(string),aid,"提交","")
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
		_= services.DB_Log(utils.GetAppid(req),conn,"资产领用","新增","",reqs["f_create_user"].(string),"新增失败",err)
		if utils.ISLog{
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI,"新增失败，"+rerr.Error())
		}
	} else {
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产领用","新增",reqs["f_create_user"].(string),reqs["f_create_user"].(string),"新增成功 ID ="+aid,err)
	}
	return

}
//领用修改
func (ctl *AssetsApplyModel) Receive_Update(req beego.Controller)(alen string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接
	var d_total float64
	fmt.Println("ss")
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
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

	if err != nil{
		panic("传入参数错误，JSON转换出错："+err.Error())
	}
	//先做数据检测
	rows_chk, err := rjson.Get("splist").Array();
	if err == nil{
		for _, row := range rows_chk {
			if each_map, ok := row.(map[string] interface{}); ok {
				if each_map["f_qty"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 数量不能为空！")
				}
				if each_map["f_price"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 价值不能为空！")
				}
				if each_map["f_public"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 公用属性不能为空！")
				}
			}
		}
	}


	d_update_main := make(map[string] string);
	//获了单号
	d_update_main["id"] = rjson.Get("id").MustString()

	//检查单据状态
	err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_Receive_Main","f_status",d_update_main["id"],"0")
	if err != nil {
		fmt.Println(err)
		panic("单据已生效不允许修改！")
	}

	d_update_main["f_accountid"] = rjson.Get("f_accountid").MustString()
	d_update_main["f_status"]  = rjson.Get("f_status").MustString()
	if (d_update_main["f_status"] =="1"){
		d_update_main["f_audit"] ="1"
		now := time.Now()
		d_update_main["f_audit_date"] = now.Format("2006-01-02 15:04:05");
	}
	d_update_main["f_serno"] =  rjson.Get("f_serno").MustString()
	d_update_main["f_re_postion"] =  rjson.Get("f_re_postion").MustString()
	d_update_main["f_re_deptid"] =  rjson.Get("f_re_deptid").MustString()
	d_update_main["f_re_userid"] =  rjson.Get("f_re_userid").MustString()
	d_update_main["f_re_date"] =  rjson.Get("f_re_date").MustString()

	d_update_main["f_note"] =  rjson.Get("f_note").MustString()

	d_update_main["f_update_user"] =  rjson.Get("f_update_user").MustString()
	d_update_main["f_update_time"] =  time.Now().Format("2006-01-02 15:04:05")


	sqlstr :=  services.Sql_JionSet(d_update_main)
	if sqlstr !="" {
		sqlstr = " update tbl_Receive_Main set "+sqlstr +" where id ="+d_update_main["id"]
	}

	//开始事务
	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	//修改角色
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)

	if err != nil{
		panic("主单修改失败,"+err.Error())
	}
	//先删除从单
	sqlstr = " delete from tbl_Receive_Detail where f_mid =  "+d_update_main["id"]
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)
	//插入从单
	if err ==nil{
		rows, _ := rjson.Get("splist").Array();
		qxsql :=" INSERT INTO tbl_Receive_Detail(f_mid,f_spid,f_qty,f_price,f_total,f_custody_deptid,f_custody_userid,f_postion,f_use_deptid,f_public,f_use_userid,"+
			" 									   f_note,f_create_user,f_create_time,f_customid) "+
			" VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,now(),?)"
		for _, row := range rows {
			if each_map, ok := row.(map[string] interface{}); ok {

				if each_map["f_price"].(string) == ""{
					each_map["f_price"] = "0"
				}
				if each_map["f_custody_deptid"].(string) == ""{
					each_map["f_custody_deptid"] = "-1"
				}
				if each_map["f_custody_userid"].(string) == ""{
					each_map["f_custody_userid"] = "-1"
				}
				if each_map["f_postion"].(string) == ""{
					each_map["f_postion"] = "-1"
				}
				if each_map["f_use_deptid"].(string) == ""{
					each_map["f_use_deptid"] = "-1"
				}
				if each_map["f_public"].(string) == ""{
					each_map["f_public"] = "0"
				}
				if each_map["f_use_userid"].(string) == ""{
					each_map["f_use_userid"] = "0"
				}
				d_qty,_ := strconv.ParseFloat(each_map["f_qty"].(string),64)
				d_price,_ := strconv.ParseFloat(each_map["f_price"].(string),64)
				d_total = d_qty * d_price
				//f_mid,f_spid,f_qty,f_price,f_total,f_custody_deptid,f_custody_userid,f_postion,f_use_deptid,f_public,f_use_userid,"+
				//" 									   f_note,f_create_user
				_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,qxsql,d_update_main["id"],each_map["f_spid"].(string),each_map["f_qty"].(string),each_map["f_price"].(string),
					d_total,each_map["f_custody_deptid"].(string),each_map["f_custody_userid"].(string),each_map["f_postion"].(string),each_map["f_use_deptid"].(string),
					each_map["f_public"].(string),each_map["f_use_userid"].(string),each_map["f_note"].(string),rjson.Get("f_update_user").MustString(),utils.GetAppid(req))
				if err !=nil{
					break;
				}
			}
		}
	}

	// 修改成功后 资产位置要更改到领用部门
	if (err == nil) && (d_update_main["f_status"] == "1"){
		var aparam = []string{"1", d_update_main["id"],d_update_main["f_update_user"] ,utils.GetAppid(req)}
		var rdata map[string]interface{}

		if err ==nil{
			rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
			if (err !=nil) || (rdata["f_state"] !="1") {
				err = fmt.Errorf("%s","库存变动失败！")
			}
		}
	}

	if err == nil{
		if (d_update_main["f_status"] =="0")||(d_update_main["f_status"] ==""){
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产领用",d_update_main["f_update_user"],d_update_main["id"],"草稿单修改","")
			//atypename string,auserid int,aid int,aopr string,amemo string
		}else {
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产领用",d_update_main["f_update_user"],d_update_main["id"],"草稿单修改","")
			if err ==nil {
				err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产领用",d_update_main["f_update_user"],d_update_main["id"],"提交","")
			}
		}
	}


	if (err != nil)  {
		rerr = err
		err = tx.Rollback()
		if err == nil{
			fmt.Println("回滚完成")
		} else {
			fmt.Println(err)
		}
		_= services.DB_Log(utils.GetAppid(req),conn,"资产领用","修改","",reqs["f_update_user"].(string),"修改失败",err)
		if utils.ISLog{
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI,"修改失败，"+rerr.Error())
		}
	} else {
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产领用","修改",reqs["f_update_user"].(string),reqs["f_update_user"].(string),"修改成功",err)
	}
	return
}
//领用删除
func (ctl *AssetsApplyModel) Receive_Delete(req beego.Controller)(rerr error){
	var reqs map[string]interface{};
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var conn = services.NewConnection(utils.GetAppid(req))
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

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
	tx,_ := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()

	//检查单据状态
	err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_Receive_Main","f_status",reqs["id"].(string),"0")
	if err != nil {
		fmt.Println(err)
		panic("单据已生效不允许删除！")
	}

	sqlstr := "delete from tbl_Receive_Main where id = ? "

	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr,reqs["id"].(string))
	//删除从单
	if err == nil{
		sqlstr = " DELETE from tbl_Receive_Detail WHERE f_mid ="+reqs["id"].(string)
		_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr,)
	}

	if err !=nil{
		rerr = err
		rerr = tx.Rollback()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产领用","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除失败",err)
	} else{

		rerr = nil
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产领用","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除成功 单号 ="+reqs["f_serno"].(string),err)
	}
	return
}
//领用 状态操作
func (ctl *AssetsApplyModel) Receive_SetStatus(req beego.Controller,atype string)(rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接
	var newstate string
	var rkstate  string
	var billMemo string


	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
			return
		}
	}()
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);

	//if (reqs["f_re_date"] != nil) &&(reqs["f_re_date"].(string) != ""){
	//	dt, err1 := time.Parse("2006-01-02",reqs["f_re_date"].(string))
	//	reqs["f_re_date"] =dt.Format("2006-01-02")
	//	reqs["f_re_date"] =reqs["f_re_date"].(string)+" "+time.Now().Format("15:04:05");
	//	//now.Format("2006-01-02 15:04:05");
	//	fmt.Println(err1)
	//}
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	stype := atype  //    -1辙消提交
	fmt.Println("状态检测")
	//状态检测
	if stype == "-1"{  //辙消提交
		err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_Receive_Main","f_status",reqs["id"].(string),"1")
		if err != nil {
			fmt.Println(err)
			panic("单据不在已提交状态不能辙消！")
		}
		err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_Receive_Main","f_audit",reqs["id"].(string),"1")
		if err != nil {
			fmt.Println(err)
			panic("单据不在领用状态不能辙消！")
		}
	}

	switch stype{  // -1辙消提交
	case "-1":{
		newstate ="0"
		rkstate = "0"
		billMemo = "辙消领用"}
	default:
		newstate =""
	}

	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()

	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,"update tbl_Receive_Main set f_status = ?,f_audit=? where id =? ",newstate,rkstate,reqs["id"].(string))
	var aparam = []string{"1", reqs["id"].(string),reqs["userid"].(string),utils.GetAppid(req)}
	var rdata map[string]interface{}

	if err ==nil{
		rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
		if (err !=nil) || (rdata["f_state"] !="1") {
			err = fmt.Errorf("%s","库存变动失败！")
		}
	}
	if err == nil{
		err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产领用",reqs["userid"].(string),reqs["id"].(string),billMemo,reqs["memo"].(string))
	}
	if err == nil{
		err= services.DB_Log(utils.GetAppid(req),conn,"资产领用",billMemo,"",reqs["userid"].(string),"操作成功",err)
	}

	if err ==nil{
		_ = tx.Commit()
	}else {
		_ = tx.Rollback();
		rerr = err
		if utils.ISLog{
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI,"操作失败，"+rerr.Error())
		}
	}

	return rerr
}


//退库查询
func (ctl *AssetsApplyModel) AssetsBack_MainQuery(req beego.Controller) (result utils.ResData, rerr error) {
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))
	var sortmap map[string]interface{};
	var acount string
	var icount int
	var sqlstr_ord string
	var sqlstr string
	var gjwhere string
	result = utils.ResData{};

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("Receive_Main_GetList 失败，",err)
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
		}
	}()

	wvalue := req.GetString("w_value","")
	pageindex := req.GetString("curPage","")//reqs["pageindex"]
	pagesize := req.GetString("pageSize","")//reqs["pagesize"]
	userid := req.GetString("userid","")//reqs["pagesize"]
	sort := req.GetString("sort");
	SpreadWhere := req.GetString("SpreadWhere");
	if SpreadWhere != ""{ //高级筛选处理
		gjwhere = ""
		rjson,err := simplejson.NewJson([]byte(SpreadWhere));
		if err == nil{
			if rjson.Get("f_status").MustString() != ""{
				gjwhere = gjwhere+" and m.f_status ="+rjson.Get("f_status").MustString()
			}
			if rjson.Get("f_return_deptid") .MustString() != ""{
				gjwhere = gjwhere+" and m.f_return_deptid ="+rjson.Get("f_return_deptid").MustString()
			}
			if rjson.Get("f_return_username") .MustString() != ""{
				gjwhere = gjwhere+" and m.f_return_userid in (SELECT id FROM tbl_user tu WHERE tu.f_customid = "+utils.GetAppid(req)+" AND tu.f_flag = 0 and f_name LIKE '%"+rjson.Get("f_return_username").MustString()+"%')"
			}
			if (rjson.Get("f_return_bdate") .MustString() != "")&&(rjson.Get("f_return_edate") .MustString() != ""){
				gjwhere = gjwhere+" and m.f_return_date >= '"+rjson.Get("f_return_bdate").MustString()+" 00:00:00"+
					"' and m.f_return_date <= '"+rjson.Get("f_return_edate").MustString()+" 23:59:59"+"'"
			}
			gjwhere = gjwhere +" "
		}
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
			sqlstr_ord = " order by m."+sortmap["sortprop"].(string)+" "+ord
		}
	} else {
		sqlstr_ord = " order by m.id desc "
	}



	sqlstr = "CALL Msp_BillMain_Query (2,0,\""+wvalue+"\","+userid+",\""+gjwhere+"\","+pageindex+","+pagesize+",\""+sqlstr_ord+"\")"


	acount,rerr = services.Get_Procedure_Data(utils.GetAppid(req),conn,sqlstr,"icount")
	if rerr != nil{
		panic("查询合计数出错，"+rerr.Error())
	}

	icount,rerr = strconv.Atoi(acount)
	if rerr != nil{
		panic("获取合计数出错 ，"+rerr.Error())
	}

	sqlstr = "CALL Msp_BillMain_Query (2,1,\""+wvalue+"\",1,\""+gjwhere+"\","+pageindex+","+pagesize+",\""+sqlstr_ord+"\")"

	result.Result,rerr = services.FindToList(utils.GetAppid(req),conn,sqlstr)
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
//退库从单查询
func (ctl *AssetsApplyModel) AssetsBack_DetailQuery(req beego.Controller) (result utils.ResData, rerr error) {
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))

	var acount string
	var icount int

	var sqlstr string
	result = utils.ResData{};

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("Receive_DetailQuery 失败，",err)
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
		}
	}()

	smid := req.GetString("mid","")
	userid := req.GetString("userid","")//reqs["pageindex"]


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
	//		sqlstr_ord = " order by m."+sortmap["sortprop"].(string)+" "+ord
	//	}
	//} else {
	//	sqlstr_ord = " order by m.id desc "
	//}



	sqlstr = "CALL Msp_BillDetail_Query (2,"+smid+","+userid+",0)"


	acount,rerr = services.Get_Procedure_Data(utils.GetAppid(req),conn,sqlstr,"icount")
	if rerr != nil{
		panic("查询合计数出错，"+rerr.Error())
	}

	icount,rerr = strconv.Atoi(acount)
	if rerr != nil{
		panic("获取合计数出错 ，"+rerr.Error())
	}

	sqlstr = "CALL Msp_BillDetail_Query (2,"+smid+","+userid+",1)"
	result.Result,rerr = services.FindToList(utils.GetAppid(req),conn,sqlstr)
	if rerr!=nil{
		panic("查询失败,"+rerr.Error())
	}else{
		result.PageSize = icount
		result.CurPage =1
		result.Totals = icount
		result.PageSize = 0
		result.CurPage = 0
		result.Totals = icount
		if reqs["sort"] != nil{
			result.Sort = reqs["sort"].(map[string]interface{})
		}
	}
	//fmt.Println(result)
	return result,nil

}
//退库新增
func (ctl *AssetsApplyModel) AssetsBackAdd(req beego.Controller)(aid string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接
	var s string
	var d_total float64
	fmt.Println("ss")
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
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	json_str,err:= json.Marshal(reqs)
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}
	rjson, err := simplejson.NewJson([]byte(json_str));

	if err != nil{
		panic("传入参数错误，JSON转换出错："+err.Error())
	}
	//先做数据检测
	rows_chk, err := rjson.Get("splist").Array();
	if err == nil{
		for _, row := range rows_chk {
			if each_map, ok := row.(map[string] interface{}); ok {
				if each_map["f_qty"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 数量不能为空！")
				}
				if each_map["f_price"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 价值不能为空！")
				}
			}
		}
	}

	if (rows_chk==nil) || (len(rows_chk)==0){
		panic("不能保存空单据！")
	}


	b_main := make(map[string] interface{});
	//获了单号
	b_main["f_accountid"] = "1"
	s,err = services.Get_Procedure_Data(utils.GetAppid(req),conn," CALL Msp_Get_Serno ("+utils.GetAppid(req)+",'tbl_Return_Main','f_serno','TK') ","f_serno")
	if err != nil{
		panic("单号获取出错,"+err.Error())
	}
	b_main["f_serno"] =s
	if rjson.Get("f_status").MustString() ==""{
		b_main["f_status"] = "0"
		b_main["f_audit"] ="0"
	} else {
		b_main["f_status"] = rjson.Get("f_status").MustString()
		if (b_main["f_status"] =="1"){
			b_main["f_audit"] ="1"
			now := time.Now()
			b_main["f_audit_date"] = now.Format("2006-01-02 15:04:05");
		}else {
			b_main["f_audit"] ="0"
		}
	}

	b_main["f_deptid"] = rjson.Get("f_deptid").MustString() //制单部门
	b_main["f_return_date"] = rjson.Get("f_return_date").MustString()
	b_main["f_return_deptid"] = rjson.Get("f_return_deptid").MustString()
	b_main["f_postion"] = rjson.Get("f_postion").MustString()
	b_main["f_return_userid"] = rjson.Get("f_return_userid").MustString()
	b_main["f_note"] = rjson.Get("f_note").MustString()
	b_main["f_create_user"] = rjson.Get("f_create_user").MustString()
	now := time.Now()
	b_main["f_create_time"]= now.Format("2006-01-02 15:04:05");
	b_main["f_customid"]= utils.GetAppid(req)
	b_main["f_flag"]= "0"

	//开始事务
	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	//插入主单
	aid,err = services.One_ADD_Table(utils.GetAppid(req),tx,conn,"tbl_Return_Main",b_main)


	//插入从单
	if err ==nil{
		rows, _ := rjson.Get("splist").Array();
		//`f_befor_deptid`		int  NOT  NULL  COMMENT '退库前管理部门',
		//`f_befor_userid`		int  NOT  NULL  COMMENT '退库前管理人员',
		//`f_beforuse_deptid`		int  NOT  NULL  COMMENT '退库前使用部门',
		//`f_beforuse_userid`		int  NOT  NULL  COMMENT '退库前使用人员',
		//`f_beforuse_option`		int  NOT  NULL  COMMENT '退库前所在位置',
		//`f_custody_deptid`		int  NOT  NULL  COMMENT '归还后管理部门',
		//`f_custody_userid`		int  NOT  NULL  COMMENT '归还后管理人员',
		//`f_postion`		    int    NULL  COMMENT '归还后存放位置',

		qxsql :=" INSERT INTO tbl_Return_Detail(f_mid,f_spid,f_qty,f_price,f_total,f_custody_deptid,f_custody_userid,f_postion,"+
			    " 									   f_note,f_create_user,f_create_time,f_customid ) "+
		 	    " values(?,?,?,?,?,?,?,?,?,?,now(),?) "
		for _, row := range rows {
			if each_map, ok := row.(map[string] interface{}); ok {

				if each_map["f_price"].(string) == ""{
					each_map["f_price"] = "0"
				}
				if each_map["f_custody_deptid"].(string) == ""{
					each_map["f_custody_deptid"] = "-1"
				}
				if each_map["f_custody_userid"].(string) == ""{
					each_map["f_custody_userid"] = "-1"
				}
				if each_map["f_postion"].(string) == ""{
					each_map["f_postion"] = "-1"
				}

				d_qty,_ := strconv.ParseFloat(each_map["f_qty"].(string),64)
				d_price,_ := strconv.ParseFloat(each_map["f_price"].(string),64)
				d_total = d_qty * d_price
				//f_mid,f_spid,f_qty,f_price,f_total,f_custody_deptid,f_custody_userid,f_postion,f_use_deptid,f_public,f_use_userid,"+
				//" 									   f_note,f_create_user
				_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,qxsql,aid,each_map["f_spid"].(string),each_map["f_qty"].(string),each_map["f_price"].(string),
					d_total,each_map["f_custody_deptid"].(string),each_map["f_custody_userid"].(string),each_map["f_postion"].(string) ,// 归还后管理部门，归还后管理人，归还后存放位置
					"\""+each_map["f_note"].(string)+"\"",rjson.Get("f_create_user").MustString(),utils.GetAppid(req))
				if err !=nil{
					break;
				}
			}
		}
	}

	// 新增成功后 资产位置要更改到领用部门
	if (err == nil) && (b_main["f_status"] == "1"){
		var aparam = []string{"2", aid,b_main["f_create_user"].(string),utils.GetAppid(req)}
		var rdata map[string]interface{}

		if err ==nil{
			rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
			if (err !=nil) || (rdata["f_state"] !="1") {
				err = fmt.Errorf("%s","库存变动失败！")
			}
		}
	}

	if err == nil{
		if (rjson.Get("f_status").MustString() =="0")||(rjson.Get("f_status").MustString() ==""){
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产退库",b_main["f_create_user"].(string),aid,"新增草稿单","")
			//atypename string,auserid int,aid int,aopr string,amemo string
		}else {
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产退库",b_main["f_create_user"].(string),aid,"新增草稿单","")
			if err ==nil {
				err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产退库",b_main["f_create_user"].(string),aid,"提交","")
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
		_= services.DB_Log(utils.GetAppid(req),conn,"资产退库","新增","",reqs["f_create_user"].(string),"新增失败",err)
		if utils.ISLog{
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI,"新增失败，"+rerr.Error())
		}
	} else {
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产退库","新增",reqs["f_create_user"].(string),reqs["f_create_user"].(string),"新增成功 ID ="+aid,err)
	}
	return

}
//退库修改
func (ctl *AssetsApplyModel) AssetsBackUpdate(req beego.Controller)(alen string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接
	var d_total float64
	fmt.Println("ss")
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
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

	if err != nil{
		panic("传入参数错误，JSON转换出错："+err.Error())
	}
	//先做数据检测
	rows_chk, err := rjson.Get("splist").Array();
	if err == nil{
		for _, row := range rows_chk {
			if each_map, ok := row.(map[string] interface{}); ok {
				if each_map["f_qty"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 数量不能为空！")
				}
				if each_map["f_price"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 价值不能为空！")
				}
			}
		}
	}


	d_update_main := make(map[string] string);
	//获了单号
	d_update_main["id"] = rjson.Get("id").MustString()

	//检查单据状态
	err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_Return_Main","f_status",d_update_main["id"],"0")
	if err != nil {
		fmt.Println(err)
		panic("单据已生效不允许修改！")
	}


	d_update_main["f_status"]  = rjson.Get("f_status").MustString()
	if (d_update_main["f_status"] =="1"){
		d_update_main["f_audit"] ="1"
		now := time.Now()
		d_update_main["f_audit_date"] = now.Format("2006-01-02 15:04:05");
	}
	d_update_main["f_serno"] =  rjson.Get("f_serno").MustString()
	//d_update_main["f_status"] =  rjson.Get("f_status").MustString()
	d_update_main["f_return_date"] =  rjson.Get("f_return_date").MustString()
	d_update_main["f_return_deptid"] =  rjson.Get("f_return_deptid").MustString()
	d_update_main["f_return_userid"] =  rjson.Get("f_return_userid").MustString()
	d_update_main["f_postion"] = rjson.Get("f_postion").MustString()
	d_update_main["f_note"] =  rjson.Get("f_note").MustString()

	d_update_main["f_update_user"] =  rjson.Get("f_update_user").MustString()
	d_update_main["f_update_time"] =  time.Now().Format("2006-01-02 15:04:05")


	sqlstr :=  services.Sql_JionSet(d_update_main)
	if sqlstr !="" {
		sqlstr = " update tbl_Return_Main set "+sqlstr +" where id ="+d_update_main["id"]
	}

	//开始事务
	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	//修改角色
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)

	if err != nil{
		panic("主单修改失败,"+err.Error())
	}
	//先删除从单
	sqlstr = " delete from tbl_Return_Detail where f_mid =  "+d_update_main["id"]
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)
	//插入从单
	if err ==nil{
		rows, _ := rjson.Get("splist").Array();
		qxsql :=" INSERT INTO tbl_Return_Detail(f_mid,f_spid,f_qty,f_price,f_total,f_custody_deptid,f_custody_userid,f_postion,"+
			" 									   f_note,f_create_user,f_create_time,f_customid ) "+
			" values(?,?,?,?,?,?,?,?,?,?,now(),?) "
		for _, row := range rows {
			if each_map, ok := row.(map[string] interface{}); ok {

				if each_map["f_price"].(string) == ""{
					each_map["f_price"] = "0"
				}
				if each_map["f_custody_deptid"].(string) == ""{
					each_map["f_custody_deptid"] = "-1"
				}
				if each_map["f_custody_userid"].(string) == ""{
					each_map["f_custody_userid"] = "-1"
				}
				if each_map["f_postion"].(string) == ""{
					each_map["f_postion"] = "-1"
				}

				d_qty,_ := strconv.ParseFloat(each_map["f_qty"].(string),64)
				d_price,_ := strconv.ParseFloat(each_map["f_price"].(string),64)
				d_total = d_qty * d_price
				//f_mid,f_spid,f_qty,f_price,f_total,f_custody_deptid,f_custody_userid,f_postion,f_use_deptid,f_public,f_use_userid,"+
				//" 									   f_note,f_create_user
				_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,qxsql,rjson.Get("id").MustString(),each_map["f_spid"].(string),each_map["f_qty"].(string),each_map["f_price"].(string),
					d_total,each_map["f_custody_deptid"].(string),each_map["f_custody_userid"].(string),each_map["f_postion"].(string) ,// 归还后管理部门，归还后管理人，归还后存放位置
					"\""+each_map["f_note"].(string)+"\"",rjson.Get("f_update_user").MustString(),utils.GetAppid(req))
				if err !=nil{
					break;
				}
			}
		}
	}
	// 修改成功后 资产位置要更改到领用部门
	if (err == nil) && (d_update_main["f_status"] == "1"){
		var aparam = []string{"2", d_update_main["id"],d_update_main["f_update_user"],utils.GetAppid(req)}
		var rdata map[string]interface{}

		if err ==nil{
			rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
			if (err !=nil) || (rdata["f_state"] !="1") {
				err = fmt.Errorf("%s","库存变动失败！")
			}
		}
	}

	if err == nil{
		if (d_update_main["f_status"] =="0")||(d_update_main["f_status"] ==""){
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产退库",d_update_main["f_update_user"],d_update_main["id"],"草稿单修改","")
			//atypename string,auserid int,aid int,aopr string,amemo string
		}else {
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产退库",d_update_main["f_update_user"],d_update_main["id"],"草稿单修改","")
			if err ==nil {
				err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产退库",d_update_main["f_update_user"],d_update_main["id"],"提交","")
			}
		}
	}


	if (err != nil)  {
		rerr = err
		err = tx.Rollback()
		if err == nil{
			fmt.Println("回滚完成")
		} else {
			fmt.Println(err)
		}
		_= services.DB_Log(utils.GetAppid(req),conn,"资产退库","修改","",reqs["f_update_user"].(string),"修改失败",err)
		if utils.ISLog{
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI,"修改失败，"+rerr.Error())
		}
	} else {
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产退库","修改",reqs["f_update_user"].(string),reqs["f_update_user"].(string),"修改成功",err)
	}
	return
}
//退库删除
func (ctl *AssetsApplyModel) AssetsBackDel(req beego.Controller)(rerr error){
	var reqs map[string]interface{};
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var conn = services.NewConnection(utils.GetAppid(req))
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

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
	//检查单据状态
	err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_return_main","f_status",reqs["id"].(string),"0")
	if err != nil {
		fmt.Println(err)
		panic("单据已生效不允许删除！")
	}


	tx,_ := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	sqlstr := "update tbl_Return_Main set  f_flag = 1, f_delete_user = "+reqs["f_delete_user"].(string)+",f_delete_time = now() where id ="+reqs["id"].(string)

	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)
	//删除从单
	if err == nil{
		sqlstr = " update tbl_Return_Detail set  f_flag = 1, f_delete_user = "+reqs["f_delete_user"].(string)+",f_delete_time = now()  WHERE f_mid ="+reqs["id"].(string)
		_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr,)
	}

	if err !=nil{
		rerr = err
		rerr = tx.Rollback()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产退库","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除失败",err)
	} else{

		rerr = nil
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产退库","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除成功 单号 ="+reqs["f_serno"].(string),err)
	}
	return
}
//退库 状态操作
func (ctl *AssetsApplyModel) AssetsBackSubmit_Back(req beego.Controller,atype string)(rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接
	var newstate string
	var rkstate  string
	var billMemo string

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
			return
		}
	}()
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	stype := atype  //    -1辙消提交

	//状态检测
	if stype == "-1"{  //辙消提交
		err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_return_main","f_status",reqs["id"].(string),"1")
		if err != nil {
			fmt.Println(err)
			panic("单据不在已提交状态不能辙消！")
		}
	}

	switch stype{  // -1辙消提交
	case "-1":{
		newstate ="0"
		rkstate = "0"
		billMemo = "辙消提交"}
	default:
		newstate =""
	}

	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,"update tbl_return_main set f_status = ?,f_audit =? where id =? ",newstate,rkstate,reqs["id"].(string))

	// 修改成功后 资产位置要更改到领用部门
	if (err == nil) {
		var aparam = []string{"2", reqs["id"].(string),reqs["userid"].(string),utils.GetAppid(req)}
		var rdata map[string]interface{}

		if err ==nil{
			rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
			if (err !=nil) || (rdata["f_state"] !="1") {
				err = fmt.Errorf("%s","库存变动失败！")
			}
		}
	}

	if err == nil{
		err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产退库",reqs["userid"].(string),reqs["id"].(string),billMemo,reqs["memo"].(string))
	}
	if err == nil{
		err= services.DB_Log(utils.GetAppid(req),conn,"辙消提交",billMemo,"",reqs["userid"].(string),"辙消成功",err)
	}

	if err ==nil{
		_ = tx.Commit()
	}else {
		_ = tx.Rollback();
		rerr = err
		if utils.ISLog{
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI,"操作失败，"+rerr.Error())
		}
	}

	return rerr
}

//借用查询
func (ctl *AssetsApplyModel) AssetsLend_MainQuery(req beego.Controller) (result utils.ResData, rerr error) {
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))
	var sortmap map[string]interface{};
	var acount string
	var icount int
	var sqlstr_ord string
	var sqlstr string
	var gjwhere string
	result = utils.ResData{};

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("AssetsLend_MainQuery 失败，",err)
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
		}
	}()

	wvalue := req.GetString("w_value","")
	pageindex := req.GetString("curPage","")//reqs["pageindex"]
	pagesize := req.GetString("pageSize","")//reqs["pagesize"]
	userid := req.GetString("userid","")//reqs["pagesize"]
	sort := req.GetString("sort");
	SpreadWhere := req.GetString("SpreadWhere");
	if SpreadWhere != ""{ //高级筛选处理
		gjwhere = ""
		rjson,err := simplejson.NewJson([]byte(SpreadWhere));
		if err == nil{
			if rjson.Get("f_status").MustString() != ""{
				gjwhere = gjwhere+" and m.f_status ="+rjson.Get("f_status").MustString()
			}
			if rjson.Get("f_type").MustString() != ""{
				gjwhere = gjwhere+" and m.f_Type ="+rjson.Get("f_type").MustString()
			}
			if rjson.Get("f_borrow_deptid") .MustString() != ""{
				gjwhere = gjwhere+" and m.f_Lend_Deptid ="+rjson.Get("f_borrow_deptid").MustString()
			}
			if rjson.Get("f_borrow_username") .MustString() != ""{
				gjwhere = gjwhere+" and ( m.f_Lend_In_people in (SELECT id FROM tbl_user tu WHERE tu.f_customid = "+utils.GetAppid(req)+" AND tu.f_flag = 0 and f_name LIKE '%"+rjson.Get("f_borrow_username").MustString()+"%')"+
					" or m.f_Lend_Out_people in (SELECT id FROM tbl_basic tb WHERE tb.f_customid = "+utils.GetAppid(req)+" and tb.f_type = 5 AND tb.f_name LIKE '%"+rjson.Get("f_borrow_username").MustString()+"%')"+
					" or m.f_Lend_PID in (SELECT id FROM tbl_basic tb WHERE tb.f_customid = "+utils.GetAppid(req)+" and tb.f_type = 3 AND tb.f_name LIKE '%"+rjson.Get("f_borrow_username").MustString()+"%') ) "
			}
			if (rjson.Get("f_borrow_bdate") .MustString() != "")&&(rjson.Get("f_borrow_edate") .MustString() != ""){
				gjwhere = gjwhere+" and m.f_Lend_date >= '"+rjson.Get("f_borrow_bdate").MustString()+" 00:00:00"+
					"' and m.f_Lend_date <= '"+rjson.Get("f_borrow_edate").MustString()+" 23:59:59"+"'"
			}
			if (rjson.Get("f_return_bdate") .MustString() != "") {
				gjwhere = gjwhere+" and m.f_Return_date >= '"+rjson.Get("f_return_bdate").MustString()+" 00:00:00' "
			}
			if (rjson.Get("f_return_edate") .MustString() != "") {
				gjwhere = gjwhere+" and m.f_Return_date <= '"+rjson.Get("f_return_edate").MustString()+" 23:59:59' "
			}
			gjwhere = gjwhere +" "
		}
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
			sqlstr_ord = " order by m."+sortmap["sortprop"].(string)+" "+ord
		}
	} else {
		sqlstr_ord = " order by m.id desc "
	}



	sqlstr = "CALL Msp_BillMain_Query (3,0,\""+wvalue+"\","+userid+",\""+gjwhere+"\","+pageindex+","+pagesize+",\""+sqlstr_ord+"\")"


	acount,rerr = services.Get_Procedure_Data(utils.GetAppid(req),conn,sqlstr,"icount")
	if rerr != nil{
		panic("查询合计数出错，"+rerr.Error())
	}

	icount,rerr = strconv.Atoi(acount)
	if rerr != nil{
		panic("获取合计数出错 ，"+rerr.Error())
	}

	sqlstr = "CALL Msp_BillMain_Query (3,1,\""+wvalue+"\",1,\""+gjwhere+"\","+pageindex+","+pagesize+",\""+sqlstr_ord+"\")"

	result.Result,rerr = services.FindToList(utils.GetAppid(req),conn,sqlstr)
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
//借用从单查询
func (ctl *AssetsApplyModel) AssetsLend_DetailQuery(req beego.Controller) (result utils.ResData, rerr error) {
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))

	var acount string
	var icount int

	var sqlstr string
	result = utils.ResData{};

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("AssetsLend_DetailQuery 失败，",err)
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
		}
	}()

	smid := req.GetString("mid","")
	userid := req.GetString("userid","")//reqs["pageindex"]


	sqlstr = "CALL Msp_BillDetail_Query (3,"+smid+","+userid+",0)"


	acount,rerr = services.Get_Procedure_Data(utils.GetAppid(req),conn,sqlstr,"icount")
	if rerr != nil{
		panic("查询合计数出错，"+rerr.Error())
	}

	icount,rerr = strconv.Atoi(acount)
	if rerr != nil{
		panic("获取合计数出错 ，"+rerr.Error())
	}

	sqlstr = "CALL Msp_BillDetail_Query (3,"+smid+","+userid+",1)"
	result.Result,rerr = services.FindToList(utils.GetAppid(req),conn,sqlstr)
	if rerr!=nil{
		panic("查询失败,"+rerr.Error())
	}else{
		result.PageSize = icount
		result.CurPage =1
		result.Totals = icount
		result.PageSize = 0
		result.CurPage = 0
		result.Totals = icount
		if reqs["sort"] != nil{
			result.Sort = reqs["sort"].(map[string]interface{})
		}
	}
	//fmt.Println(result)
	return result,nil

}
//借用新增
func (ctl *AssetsApplyModel) AssetsLendAdd(req beego.Controller)(aid string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接
	var s string
	var d_total float64
	var f_lend_in_people string
	var f_len_out_people string
	var f_lend_pid string

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
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	json_str,err:= json.Marshal(reqs)
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}
	rjson, err := simplejson.NewJson([]byte(json_str));

	if err != nil{
		panic("传入参数错误，JSON转换出错："+err.Error())
	}
	//先做数据检测
	rows_chk, err := rjson.Get("splist").Array();
	if err == nil{
		for _, row := range rows_chk {
			if each_map, ok := row.(map[string] interface{}); ok {
				if each_map["f_qty"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 数量不能为空！")
				}
				if each_map["f_price"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 价值不能为空！")
				}
			}
		}
	}

	if (rows_chk==nil) || (len(rows_chk)==0){
		panic("不能保存空单据！")
	}


	b_main := make(map[string] interface{});
	//获了单号
	b_main["f_accountid"] = "1"
	s,err = services.Get_Procedure_Data(utils.GetAppid(req),conn," CALL Msp_Get_Serno ("+utils.GetAppid(req)+",'tbl_lend_out_main','f_serno','JY') ","f_serno")
	if err != nil{
		panic("单号获取出错,"+err.Error())
	}
	b_main["f_serno"] =s
	if rjson.Get("f_status").MustString() ==""{
		b_main["f_status"] = "0"
		b_main["f_audit"] ="0"
	} else {
		b_main["f_status"] = rjson.Get("f_status").MustString()
		if (b_main["f_status"] =="1"){
			b_main["f_audit"] ="1"
			now := time.Now()
			b_main["f_audit_date"] = now.Format("2006-01-02 15:04:05");
			b_main["f_audit_userid"] = rjson.Get("f_create_user").MustString();
		}else {
			b_main["f_audit"] ="0"
		}
	}

	b_main["f_type"] = rjson.Get("f_type").MustString()  // 0内部借用 1外部联系人借用 2供应商借用
	b_main["f_deptid"] = rjson.Get("f_deptid").MustString() //制单部门
	if b_main["f_type"].(string) =="0"{  // 内部借用
		b_main["f_lend_in_people"] = rjson.Get("f_lend_id").MustString() //内部借用人
		b_main["f_lend_out_people"] = "-1" //外部借用人
		b_main["f_lend_pid"] = "-1" //借用供应商
	}
	if b_main["f_type"].(string) =="1"{  // 外部联系人借用
		b_main["f_lend_in_people"] ="-1" //内部借用人
		b_main["f_lend_out_people"] = rjson.Get("f_lend_id").MustString() //外部借用人
		b_main["f_lend_pid"] = "-1" //借用供应商
	}
	if b_main["f_type"].(string) =="2"{  // 供应商借用
		b_main["f_lend_in_people"] = "-1" //内部借用人
		b_main["f_lend_out_people"] = "-1" //外部借用人
		b_main["f_lend_pid"] = rjson.Get("f_lend_id").MustString()//借用供应商
	}

	b_main["f_lend_date"] = rjson.Get("f_lend_date").MustString() //借用日期
	b_main["f_return_date"] = rjson.Get("f_return_date").MustString()
	b_main["f_Lend_Deptid"] = rjson.Get("f_Lend_Deptid").MustString()
	b_main["f_Lend_In_postion"] = rjson.Get("f_Lend_In_postion").MustString()

	b_main["f_note"] = rjson.Get("f_note").MustString()
	b_main["f_create_user"] = rjson.Get("f_create_user").MustString()
	now := time.Now()
	b_main["f_create_time"]= now.Format("2006-01-02 15:04:05");
	b_main["f_customid"]= utils.GetAppid(req)
	b_main["f_flag"]= "0"

	//开始事务
	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	//插入主单
	aid,err = services.One_ADD_Table(utils.GetAppid(req),tx,conn,"tbl_lend_out_main",b_main)


	//插入从单
	if err ==nil{
		rows, _ := rjson.Get("splist").Array();

		qxsql :=" INSERT INTO tbl_lend_out_detail(f_mid,f_spid,f_qty,f_price,f_total,f_type,f_Lend_In_people,f_Lend_Out_people,f_Lend_PID,"+
			" 									  f_lend_deptid,f_postion,f_Return_date,f_note,f_create_user,f_create_time,f_customid) "+
			" values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,now(),?) "
		for _, row := range rows {
			if each_map, ok := row.(map[string] interface{}); ok {

				if each_map["f_price"].(string) == ""{
					each_map["f_price"] = "0"
				}

				d_qty,_ := strconv.ParseFloat(each_map["f_qty"].(string),64)
				d_price,_ := strconv.ParseFloat(each_map["f_price"].(string),64)
				d_total = d_qty * d_price
				if b_main["f_type"].(string) =="0"{
					f_lend_in_people = each_map["f_lend_id"].(string)
					f_len_out_people ="-1"
					f_lend_pid ="-1"
				}
				if b_main["f_type"].(string) =="1"{
					f_lend_in_people = "-1"
					f_len_out_people =each_map["f_lend_id"].(string)
					f_lend_pid ="-1"
				}
				if b_main["f_type"].(string) =="2"{
					f_lend_in_people = "-1"
					f_len_out_people ="-1"
					f_lend_pid =each_map["f_lend_id"].(string)
				}

				_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,qxsql,aid,each_map["f_spid"].(string),each_map["f_qty"].(string),each_map["f_price"].(string),d_total,
					b_main["f_type"].(string),f_lend_in_people,f_len_out_people ,f_lend_pid,each_map["f_lend_deptid"].(string),each_map["f_postion"].(string),
					each_map["f_Return_date"].(string),each_map["f_note"].(string),rjson.Get("f_create_user").MustString(),utils.GetAppid(req))
				if err !=nil{
					break;
				}
			}
		}
	}

	// 新增成功后 资产位置要更改到领用部门
	if (err == nil) && (b_main["f_status"] == "1"){
		var aparam = []string{"3", aid,b_main["f_create_user"].(string),utils.GetAppid(req)}
		var rdata map[string]interface{}

		if err ==nil{
			rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
			if (err !=nil) || (rdata["f_state"] !="1") {
				err = fmt.Errorf("%s","库存变动失败！")
			}
		}
	}

	if err == nil{
		if (rjson.Get("f_status").MustString() =="0")||(rjson.Get("f_status").MustString() ==""){
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产借用",b_main["f_create_user"].(string),aid,"新增草稿单","")
			//atypename string,auserid int,aid int,aopr string,amemo string
		}else {
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产借用",b_main["f_create_user"].(string),aid,"新增草稿单","")
			if err ==nil {
				err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产借用",b_main["f_create_user"].(string),aid,"提交","")
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
		_= services.DB_Log(utils.GetAppid(req),conn,"资产借用","新增","",reqs["f_create_user"].(string),"新增失败",err)
		if utils.ISLog{
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI,"新增失败，"+rerr.Error())
		}
	} else {
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产借用","新增",reqs["f_create_user"].(string),reqs["f_create_user"].(string),"新增成功 ID ="+aid,err)
	}
	return

}
//借用修改
func (ctl *AssetsApplyModel) AssetsLendUpdate(req beego.Controller)(alen string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接
	var f_lend_in_people string
	var f_len_out_people string
	var f_lend_pid string
	var d_total float64

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
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

	if err != nil{
		panic("传入参数错误，JSON转换出错："+err.Error())
	}
	//先做数据检测
	rows_chk, err := rjson.Get("splist").Array();
	if err == nil{
		for _, row := range rows_chk {
			if each_map, ok := row.(map[string] interface{}); ok {
				if each_map["f_qty"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 数量不能为空！")
				}
				if each_map["f_price"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 价值不能为空！")
				}
			}
		}
	}

	d_update_main := make(map[string] string);
	//获了单号
	d_update_main["id"] = rjson.Get("id").MustString()

	//检查单据状态
	err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_lend_out_main","f_status",d_update_main["id"],"0")
	if err != nil {
		fmt.Println(err)
		panic("单据已生效不允许修改！")
	}


	d_update_main["f_status"]  = rjson.Get("f_status").MustString()
	if (d_update_main["f_status"] =="1"){
		d_update_main["f_audit"] ="1"
		now := time.Now()
		d_update_main["f_audit_date"] = now.Format("2006-01-02 15:04:05");
		d_update_main["f_audit_userid"] = rjson.Get("f_update_user").MustString();
	}

	d_update_main["f_serno"] =  rjson.Get("f_serno").MustString()
	d_update_main["f_type"] = rjson.Get("f_type").MustString()
	d_update_main["f_deptid"] = rjson.Get("f_deptid").MustString() //制单部门
	if d_update_main["f_type"] =="0"{  // 内部借用
		d_update_main["f_lend_in_people"] = rjson.Get("f_lend_id").MustString() //内部借用人
		d_update_main["f_lend_out_people"] = "-1" //外部借用人
		d_update_main["f_lend_pid"] = "-1" //借用供应商
	}
	if d_update_main["f_type"] =="1"{  // 外部联系人借用
		d_update_main["f_lend_in_people"] ="-1" //内部借用人
		d_update_main["f_lend_out_people"] = rjson.Get("f_lend_id").MustString() //外部借用人
		d_update_main["f_lend_pid"] = "-1" //借用供应商
	}
	if d_update_main["f_type"] =="2"{  // 供应商借用
		d_update_main["f_lend_in_people"] = "-1" //内部借用人
		d_update_main["f_lend_out_people"] = "-1" //外部借用人
		d_update_main["f_lend_pid"] = rjson.Get("f_lend_id").MustString()//借用供应商
	}

	d_update_main["f_Lend_Deptid"] = rjson.Get("f_Lend_Deptid").MustString()
	d_update_main["f_lend_date"] = rjson.Get("f_lend_date").MustString() //借用日期
	d_update_main["f_return_date"] = rjson.Get("f_return_date").MustString()
	d_update_main["f_note"] =  rjson.Get("f_note").MustString()
	d_update_main["f_update_user"] =  rjson.Get("f_update_user").MustString()
	d_update_main["f_update_time"] =  time.Now().Format("2006-01-02 15:04:05")

	d_update_main["f_Lend_In_postion"] = rjson.Get("f_Lend_In_postion").MustString()
	sqlstr :=  services.Sql_JionSet(d_update_main)
	if sqlstr !="" {
		sqlstr = " update tbl_lend_out_main set "+sqlstr +" where id ="+d_update_main["id"]
	}

	//开始事务
	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	//修改角色
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)

	if err != nil{
		panic("主单修改失败,"+err.Error())
	}
	//先删除从单
	sqlstr = " delete from tbl_lend_out_detail where f_mid =  "+d_update_main["id"]
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)

	//插入从单
	if err ==nil{
		rows, _ := rjson.Get("splist").Array();

		qxsql :=" INSERT INTO tbl_lend_out_detail(f_mid,f_spid,f_qty,f_price,f_total,f_type,f_Lend_In_people,f_Lend_Out_people,f_Lend_PID,"+
			" 									  f_lend_deptid,f_postion,f_Return_date,f_note,f_create_user,f_create_time,f_customid) "+
			" values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,now(),?) "
		for _, row := range rows {
			if each_map, ok := row.(map[string] interface{}); ok {

				if each_map["f_price"].(string) == ""{
					each_map["f_price"] = "0"
				}

				d_qty,_ := strconv.ParseFloat(each_map["f_qty"].(string),64)
				d_price,_ := strconv.ParseFloat(each_map["f_price"].(string),64)
				d_total = d_qty * d_price
				if d_update_main["f_type"] =="0"{
					f_lend_in_people = each_map["f_lend_id"].(string)
					f_len_out_people ="-1"
					f_lend_pid ="-1"
				}
				if d_update_main["f_type"] =="1"{
					f_lend_in_people = "-1"
					f_len_out_people =each_map["f_lend_id"].(string)
					f_lend_pid ="-1"
				}
				if d_update_main["f_type"] =="2"{
					f_lend_in_people = "-1"
					f_len_out_people ="-1"
					f_lend_pid =each_map["f_lend_id"].(string)
				}

				_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,qxsql,d_update_main["id"],each_map["f_spid"].(string),each_map["f_qty"].(string),each_map["f_price"].(string),d_total,
					d_update_main["f_type"],f_lend_in_people,f_len_out_people ,f_lend_pid,each_map["f_lend_deptid"].(string),each_map["f_postion"].(string),
					each_map["f_Return_date"].(string),each_map["f_note"].(string),rjson.Get("f_update_user").MustString(),utils.GetAppid(req))
				if err !=nil{
					break;
				}
			}
		}
	}

	// 修改成功后 资产位置要更改到领用部门
	if (err == nil) && (d_update_main["f_status"] == "1"){
		var aparam = []string{"3", d_update_main["id"],d_update_main["f_update_user"],utils.GetAppid(req)}
		var rdata map[string]interface{}

		if err ==nil{
			rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
			if (err !=nil) || (rdata["f_state"] !="1") {
				err = fmt.Errorf("%s","库存变动失败！")
			}
		}
	}

	if err == nil{
		if (d_update_main["f_status"] =="0")||(d_update_main["f_status"] ==""){
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产借用",d_update_main["f_update_user"],d_update_main["id"],"草稿单修改","")
			//atypename string,auserid int,aid int,aopr string,amemo string
		}else {
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产借用",d_update_main["f_update_user"],d_update_main["id"],"草稿单修改","")
			if err ==nil {
				err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产借用",d_update_main["f_update_user"],d_update_main["id"],"提交","")
			}
		}
	}


	if (err != nil)  {
		rerr = err
		err = tx.Rollback()
		if err == nil{
			fmt.Println("回滚完成")
		} else {
			fmt.Println(err)
		}
		_= services.DB_Log(utils.GetAppid(req),conn,"资产借用","修改","",reqs["f_update_user"].(string),"修改失败",err)
		if utils.ISLog{
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI,"修改失败，"+rerr.Error())
		}
	} else {
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产借用","修改",reqs["f_update_user"].(string),reqs["f_update_user"].(string),"修改成功",err)
	}
	return
}
//借用删除
func (ctl *AssetsApplyModel) AssetsLendDel(req beego.Controller)(rerr error){
	var reqs map[string]interface{};
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var conn = services.NewConnection(utils.GetAppid(req))
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

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
	//检查单据状态
	err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_lend_out_main","f_status",reqs["id"].(string),"0")
	if err != nil {
		fmt.Println(err)
		panic("单据已生效不允许删除！")
	}


	tx,_ := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	sqlstr := "update tbl_lend_out_main set  f_flag = 1, f_delete_user = "+reqs["f_delete_user"].(string)+",f_delete_time = now() where id ="+reqs["id"].(string)

	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)
	//删除从单
	if err == nil{
		sqlstr = " update tbl_lend_out_detail set  f_flag = 1, f_delete_user = "+reqs["f_delete_user"].(string)+",f_delete_time = now()  WHERE f_mid ="+reqs["id"].(string)
		_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr,)
	}

	if err !=nil{
		rerr = err
		_ = tx.Rollback()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产借用","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除失败",err)
	} else{

		rerr = nil
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产借用","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除成功 单号 ="+reqs["f_serno"].(string),err)
	}
	return
}
//借用 状态操作
func (ctl *AssetsApplyModel) AssetsLendSubmit_Back(req beego.Controller,atype string)(rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接
	var newstate string
	var rkstate  string
	var billMemo string

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
			return
		}
	}()
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	stype := atype  //    -1辙消提交

	//状态检测
	if stype == "-1"{  //辙消提交
		err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_lend_out_main","f_status",reqs["id"].(string),"1")
		if err != nil {
			fmt.Println(err)
			panic("单据不在已提交状态不能辙消！")
		}
	}

	switch stype{  // -1辙消提交
	case "-1":{
		newstate ="0"
		rkstate = "0"
		billMemo = "辙消提交"}
	default:
		newstate =""
	}

	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,"update tbl_lend_out_main set f_status = ?,f_audit =? where id =? ",newstate,rkstate,reqs["id"].(string))

	// 修改成功后 资产位置要更改到领用部门
	if (err == nil) {
		var aparam = []string{"3", reqs["id"].(string),reqs["userid"].(string),utils.GetAppid(req)}
		var rdata map[string]interface{}

		if err ==nil{
			rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
			if (err !=nil) || (rdata["f_state"] !="1") {
				err = fmt.Errorf("%s","库存变动失败！")
			}
		}
	}

	if err == nil{
		err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产借用",reqs["userid"].(string),reqs["id"].(string),billMemo,reqs["memo"].(string))
	}
	if err == nil{
		err= services.DB_Log(utils.GetAppid(req),conn,"辙消提交",billMemo,"",reqs["userid"].(string),"辙消成功",err)
	}

	if err ==nil{
		_ = tx.Commit()
	}else {
		_ = tx.Rollback();
		rerr = err
		if utils.ISLog{
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI,"操作失败，"+rerr.Error())
		}
	}

	return rerr
}

//借用归还查询
func (ctl *AssetsApplyModel) AssetsLendIn_MainQuery(req beego.Controller) (result utils.ResData, rerr error) {
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))
	var sortmap map[string]interface{};
	var acount string
	var icount int
	var sqlstr_ord string
	var sqlstr string
	var gjwhere string
	result = utils.ResData{};

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("AssetsLendIn_MainQuery 失败，",err)
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
		}
	}()

	wvalue := req.GetString("w_value","")
	pageindex := req.GetString("curPage","")//reqs["pageindex"]
	pagesize := req.GetString("pageSize","")//reqs["pagesize"]
	userid := req.GetString("userid","")//reqs["pagesize"]
	sort := req.GetString("sort");
	SpreadWhere := req.GetString("SpreadWhere");
	if SpreadWhere != ""{ //高级筛选处理
		gjwhere = ""
		rjson,err := simplejson.NewJson([]byte(SpreadWhere));
		if err == nil{
			if rjson.Get("f_status").MustString() != ""{
				gjwhere = gjwhere+" and m.f_status ="+rjson.Get("f_status").MustString()
			}
			if rjson.Get("f_type").MustString() != ""{
				gjwhere = gjwhere+" and m.f_Type ="+rjson.Get("f_type").MustString()
			}

			if rjson.Get("f_R_username") .MustString() != ""{
				gjwhere = gjwhere+" and ( m.f_R_In_people in (SELECT id FROM tbl_user tu WHERE tu.f_customid = "+utils.GetAppid(req)+" AND tu.f_flag = 0 and f_name LIKE '%"+rjson.Get("f_R_username").MustString()+"%')"+
					" or m.f_R_Out_people in (SELECT id FROM tbl_basic tb WHERE tb.f_customid = "+utils.GetAppid(req)+" and tb.f_type = 5 AND tb.f_name LIKE '%"+rjson.Get("f_R_username").MustString()+"%')"+
					" or m.f_R_PID in (SELECT id FROM tbl_basic tb WHERE tb.f_customid = "+utils.GetAppid(req)+" and tb.f_type = 3 AND tb.f_name LIKE '%"+rjson.Get("f_R_username").MustString()+"%') ) "
			}
			if (rjson.Get("f_return_bdate") .MustString() != "")&&(rjson.Get("f_return_edate") .MustString() != ""){
				gjwhere = gjwhere+" and m.f_R_date >= '"+rjson.Get("f_return_bdate").MustString()+" 00:00:00"+
					"' and m.f_R_date <= '"+rjson.Get("f_return_edate").MustString()+" 23:59:59"+"'"
			}
			gjwhere = gjwhere +" "
		}
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
			sqlstr_ord = " order by m."+sortmap["sortprop"].(string)+" "+ord
		}
	} else {
		sqlstr_ord = " order by m.id desc "
	}



	sqlstr = "CALL Msp_BillMain_Query (4,0,\""+wvalue+"\","+userid+",\""+gjwhere+"\","+pageindex+","+pagesize+",\""+sqlstr_ord+"\")"


	acount,rerr = services.Get_Procedure_Data(utils.GetAppid(req),conn,sqlstr,"icount")
	if rerr != nil{
		panic("查询合计数出错，"+rerr.Error())
	}

	icount,rerr = strconv.Atoi(acount)
	if rerr != nil{
		panic("获取合计数出错 ，"+rerr.Error())
	}

	sqlstr = "CALL Msp_BillMain_Query (4,1,\""+wvalue+"\","+userid+",\""+gjwhere+"\","+pageindex+","+pagesize+",\""+sqlstr_ord+"\")"

	result.Result,rerr = services.FindToList(utils.GetAppid(req),conn,sqlstr)
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
//借用归还从单查询
func (ctl *AssetsApplyModel) AssetsLendIn_DetailQuery(req beego.Controller) (result utils.ResData, rerr error) {
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))

	var acount string
	var icount int

	var sqlstr string
	result = utils.ResData{};

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("Receive_DetailQuery 失败，",err)
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
		}
	}()

	smid := req.GetString("mid","")
	userid := req.GetString("userid","")//reqs["pageindex"]

	sqlstr = "CALL Msp_BillDetail_Query (4,"+smid+","+userid+",0)"


	acount,rerr = services.Get_Procedure_Data(utils.GetAppid(req),conn,sqlstr,"icount")
	if rerr != nil{
		panic("查询合计数出错，"+rerr.Error())
	}

	icount,rerr = strconv.Atoi(acount)
	if rerr != nil{
		panic("获取合计数出错 ，"+rerr.Error())
	}

	sqlstr = "CALL Msp_BillDetail_Query (4,"+smid+","+userid+",1)"
	result.Result,rerr = services.FindToList(utils.GetAppid(req),conn,sqlstr)
	if rerr!=nil{
		panic("查询失败,"+rerr.Error())
	}else{
		result.PageSize = icount
		result.CurPage =1
		result.Totals = icount
		result.PageSize = 0
		result.CurPage = 0
		result.Totals = icount
		if reqs["sort"] != nil{
			result.Sort = reqs["sort"].(map[string]interface{})
		}
	}
	//fmt.Println(result)
	return result,nil

}
//借用归还新增
func (ctl *AssetsApplyModel) AssetsLendInAdd(req beego.Controller)(aid string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接
	var s string
	var d_total float64
	var f_lend_in_people string
	var f_len_out_people string
	var f_lend_pid string

	fmt.Println("ss")
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
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	json_str,err:= json.Marshal(reqs)
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}
	rjson, err := simplejson.NewJson([]byte(json_str));

	if err != nil{
		panic("传入参数错误，JSON转换出错："+err.Error())
	}
	//先做数据检测
	rows_chk, err := rjson.Get("splist").Array();
	if err == nil{
		for _, row := range rows_chk {
			if each_map, ok := row.(map[string] interface{}); ok {
				if each_map["f_qty"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 数量不能为空！")
				}
				if each_map["f_price"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 价值不能为空！")
				}
			}
		}
	}

	if (rows_chk==nil) || (len(rows_chk)==0){
		panic("不能保存空单据！")
	}


	b_main := make(map[string] interface{});
	//获了单号
	b_main["f_accountid"] = "1"
	s,err = services.Get_Procedure_Data(utils.GetAppid(req),conn," CALL Msp_Get_Serno ("+utils.GetAppid(req)+",'tbl_lend_in_main','f_serno','JYGH') ","f_serno")
	if err != nil{
		panic("单号获取出错,"+err.Error())
	}
	b_main["f_serno"] =s
	if rjson.Get("f_status").MustString() ==""{
		b_main["f_status"] = "0"
		b_main["f_audit"] ="0"
	} else {
		b_main["f_status"] = rjson.Get("f_status").MustString()
		if (b_main["f_status"] =="1"){
			b_main["f_audit"] ="1"
			now := time.Now()
			b_main["f_audit_date"] = now.Format("2006-01-02 15:04:05");
		}else {
			b_main["f_audit"] ="0"
		}
	}

	b_main["f_type"] = rjson.Get("f_type").MustString() // 0内部归还  1外部联系人归还 2供应商归还
	if b_main["f_type"].(string) =="0"{  // 内部借用
		b_main["f_R_In_people"] = rjson.Get("f_lend_id").MustString() //内部借用人
		b_main["f_R_Out_people"] = "-1" //外部借用人
		b_main["f_R_PID"] = "-1" //借用供应商
	}
	if b_main["f_type"].(string) =="1"{  // 外部联系人借用
		b_main["f_R_In_people"] ="-1" //内部借用人
		b_main["f_R_Out_people"] = rjson.Get("f_lend_id").MustString() //外部借用人
		b_main["f_R_PID"] = "-1" //借用供应商
	}
	if b_main["f_type"].(string) =="2"{  // 供应商借用
		b_main["f_R_In_people"] = "-1" //内部借用人
		b_main["f_R_Out_people"] = "-1" //外部借用人
		b_main["f_R_PID"] = rjson.Get("f_lend_id").MustString()//借用供应商
	}
	b_main["f_deptid"] = rjson.Get("f_deptid").MustString() //制单部门
	//b_main["f_R_In_people"] = rjson.Get("f_R_In_people").MustString()
	//b_main["f_R_Out_people"] = rjson.Get("f_R_Out_people").MustString()
	//b_main["f_R_PID"] = rjson.Get("f_R_PID").MustString()
	b_main["f_R_date"] = rjson.Get("f_R_date").MustString()
	b_main["f_note"] = rjson.Get("f_note").MustString()
	b_main["f_create_user"] = rjson.Get("f_create_user").MustString()
	now := time.Now()
	b_main["f_create_time"]= now.Format("2006-01-02 15:04:05");
	b_main["f_customid"]= utils.GetAppid(req)
	b_main["f_flag"]= "0"
	b_main["f_postion"] = rjson.Get("f_R_postion").MustString()
	//开始事务
	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	//插入主单
	aid,err = services.One_ADD_Table(utils.GetAppid(req),tx,conn,"tbl_lend_in_main",b_main)


	//插入从单
	if err ==nil{
		rows, _ := rjson.Get("splist").Array();
		qxsql :=" INSERT INTO tbl_lend_in_detail(f_mid,f_spid,f_qty,f_price,f_total,f_Type,f_R_In_people,f_R_Out_people,f_R_PID,f_custody_deptid,"+
			" 									 f_custody_userid,f_postion,f_spstate,f_note,f_create_user,f_create_time,f_customid ) "+
			" values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,now(),?) "
		for _, row := range rows {
			if each_map, ok := row.(map[string] interface{}); ok {

				if each_map["f_price"].(string) == ""{
					each_map["f_price"] = "0"
				}
				if each_map["f_custody_deptid"].(string) == ""{
					each_map["f_custody_deptid"] = "-1"
				}
				if each_map["f_custody_userid"].(string) == ""{
					each_map["f_custody_userid"] = "-1"
				}
				if each_map["f_postion"].(string) == ""{
					each_map["f_postion"] = "-1"
				}

				if b_main["f_type"].(string) =="0"{
					f_lend_in_people = each_map["f_lend_id"].(string)
					f_len_out_people ="-1"
					f_lend_pid ="-1"
				}
				if b_main["f_type"].(string) =="1"{
					f_lend_in_people = "-1"
					f_len_out_people =each_map["f_lend_id"].(string)
					f_lend_pid ="-1"
				}
				if b_main["f_type"].(string) =="2"{
					f_lend_in_people = "-1"
					f_len_out_people ="-1"
					f_lend_pid =each_map["f_lend_id"].(string)
				}

				d_qty,_ := strconv.ParseFloat(each_map["f_qty"].(string),64)
				d_price,_ := strconv.ParseFloat(each_map["f_price"].(string),64)
				d_total = d_qty * d_price
				//f_mid,f_spid,f_qty,f_price,f_total,f_custody_deptid,f_custody_userid,f_postion,f_use_deptid,f_public,f_use_userid,"+
				//" 									   f_note,f_create_user
				_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,qxsql,aid,each_map["f_spid"].(string),each_map["f_qty"].(string),each_map["f_price"].(string),
					d_total,b_main["f_type"].(string),f_lend_in_people,f_len_out_people,f_lend_pid,each_map["f_custody_deptid"].(string),
					each_map["f_custody_userid"].(string),each_map["f_postion"].(string),each_map["f_spstate"].(string),each_map["f_note"].(string),rjson.Get("f_create_user").MustString(),utils.GetAppid(req))
				if err !=nil{
					break;
				}
			}
		}
	}

	// 新增成功后 资产位置要更改到领用部门
	if (err == nil) && (b_main["f_status"] == "1"){
		var aparam = []string{"4", aid,b_main["f_create_user"].(string),utils.GetAppid(req)}
		var rdata map[string]interface{}

		if err ==nil{
			rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
			if (err !=nil) || (rdata["f_state"] !="1") {
				err = fmt.Errorf("%s","库存变动失败！")
			}
		}
	}

	if err == nil{
		if (rjson.Get("f_status").MustString() =="0")||(rjson.Get("f_status").MustString() ==""){
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"借出归还",b_main["f_create_user"].(string),aid,"新增草稿单","")
			//atypename string,auserid int,aid int,aopr string,amemo string
		}else {
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"借出归还",b_main["f_create_user"].(string),aid,"新增草稿单","")
			if err ==nil {
				err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"借出归还",b_main["f_create_user"].(string),aid,"提交","")
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
		_= services.DB_Log(utils.GetAppid(req),conn,"借出归还","新增","",reqs["f_create_user"].(string),"新增失败",err)
		if utils.ISLog{
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI,"新增失败，"+rerr.Error())
		}
	} else {
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"借出归还","新增",reqs["f_create_user"].(string),reqs["f_create_user"].(string),"新增成功 ID ="+aid,err)
	}
	return

}
//借用归还修改
func (ctl *AssetsApplyModel) AssetsLendInUpdate(req beego.Controller)(alen string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接
	var d_total float64
	var f_lend_in_people string
	var f_len_out_people string
	var f_lend_pid string


	fmt.Println("ss")
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
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

	if err != nil{
		panic("传入参数错误，JSON转换出错："+err.Error())
	}
	//先做数据检测
	rows_chk, err := rjson.Get("splist").Array();
	if err == nil{
		for _, row := range rows_chk {
			if each_map, ok := row.(map[string] interface{}); ok {
				if each_map["f_qty"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 数量不能为空！")
				}
				if each_map["f_price"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 价值不能为空！")
				}
			}
		}
	}


	d_update_main := make(map[string] string);
	//获了单号
	d_update_main["id"] = rjson.Get("id").MustString()

	//检查单据状态
	err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_lend_in_main","f_status",d_update_main["id"],"0")
	if err != nil {
		fmt.Println(err)
		panic("单据已生效不允许修改！")
	}


	d_update_main["f_status"]  = rjson.Get("f_status").MustString()
	if (d_update_main["f_status"] =="1"){
		d_update_main["f_audit"] ="1"
		now := time.Now()
		d_update_main["f_audit_date"] = now.Format("2006-01-02 15:04:05");
	}
	d_update_main["f_serno"] =  rjson.Get("f_serno").MustString()
	d_update_main["f_type"] =  rjson.Get("f_type").MustString()
	if d_update_main["f_type"] =="0"{  // 内部借用
		d_update_main["f_R_In_people"] = rjson.Get("f_lend_id").MustString() //内部借用人
		d_update_main["f_R_Out_people"] = "-1" //外部借用人
		d_update_main["f_R_PID"] = "-1" //借用供应商
	}
	if d_update_main["f_type"] =="1"{  // 外部联系人借用
		d_update_main["f_R_In_people"] ="-1" //内部借用人
		d_update_main["f_R_Out_people"] = rjson.Get("f_lend_id").MustString() //外部借用人
		d_update_main["f_R_PID"] = "-1" //借用供应商
	}
	if d_update_main["f_type"] =="2"{  // 供应商借用
		d_update_main["f_R_In_people"] = "-1" //内部借用人
		d_update_main["f_R_Out_people"] = "-1" //外部借用人
		d_update_main["f_R_PID"] = rjson.Get("f_lend_id").MustString()//借用供应商
	}

	d_update_main["f_deptid"] =  rjson.Get("f_deptid").MustString()
	d_update_main["f_R_date"] =  rjson.Get("f_R_date").MustString()
	d_update_main["f_note"] =  rjson.Get("f_note").MustString()
	d_update_main["f_update_user"] =  rjson.Get("f_update_user").MustString()
	d_update_main["f_update_time"] =  time.Now().Format("2006-01-02 15:04:05")
	d_update_main["f_postion"] = rjson.Get("f_R_postion").MustString()

	sqlstr :=  services.Sql_JionSet(d_update_main)
	if sqlstr !="" {
		sqlstr = " update tbl_lend_in_main set "+sqlstr +" where id ="+d_update_main["id"]
	}

	//开始事务
	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	//修改角色
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)

	if err != nil{
		panic("主单修改失败,"+err.Error())
	}
	//先删除从单
	sqlstr = " delete from tbl_lend_in_detail where f_mid =  "+d_update_main["id"]
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)

	//插入从单
	if err ==nil{
		rows, _ := rjson.Get("splist").Array();
		qxsql :=" INSERT INTO tbl_lend_in_detail(f_mid,f_spid,f_qty,f_price,f_total,f_Type,f_R_In_people,f_R_Out_people,f_R_PID,f_custody_deptid,"+
			" 									 f_custody_userid,f_postion,f_spstate,f_note,f_create_user,f_create_time,f_customid ) "+
			" values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,now(),?) "
		for _, row := range rows {
			if each_map, ok := row.(map[string] interface{}); ok {

				if each_map["f_price"].(string) == ""{
					each_map["f_price"] = "0"
				}
				if each_map["f_custody_deptid"].(string) == ""{
					each_map["f_custody_deptid"] = "-1"
				}
				if each_map["f_custody_userid"].(string) == ""{
					each_map["f_custody_userid"] = "-1"
				}
				if each_map["f_postion"].(string) == ""{
					each_map["f_postion"] = "-1"
				}
				if d_update_main["f_type"]  =="0"{
					f_lend_in_people = each_map["f_lend_id"].(string)
					f_len_out_people ="-1"
					f_lend_pid ="-1"
				}
				if d_update_main["f_type"]  =="1"{
					f_lend_in_people = "-1"
					f_len_out_people =each_map["f_lend_id"].(string)
					f_lend_pid ="-1"
				}
				if d_update_main["f_type"]  =="2"{
					f_lend_in_people = "-1"
					f_len_out_people ="-1"
					f_lend_pid =each_map["f_lend_id"].(string)
				}
				d_qty,_ := strconv.ParseFloat(each_map["f_qty"].(string),64)
				d_price,_ := strconv.ParseFloat(each_map["f_price"].(string),64)
				d_total = d_qty * d_price
				//f_mid,f_spid,f_qty,f_price,f_total,f_custody_deptid,f_custody_userid,f_postion,f_use_deptid,f_public,f_use_userid,"+
				//" 									   f_note,f_create_user
				_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,qxsql,d_update_main["id"],each_map["f_spid"].(string),each_map["f_qty"].(string),each_map["f_price"].(string),
					d_total,d_update_main["f_type"],f_lend_in_people,f_len_out_people,f_lend_pid ,each_map["f_custody_deptid"].(string),
					each_map["f_custody_userid"].(string),each_map["f_postion"].(string),each_map["f_spstate"].(string),each_map["f_note"].(string),rjson.Get("f_update_user").MustString(),utils.GetAppid(req))
				if err !=nil{
					break;
				}
			}
		}
	}
	// 修改成功后 资产位置要更改到领用部门
	if (err == nil) && (d_update_main["f_status"] == "1"){
		var aparam = []string{"4", d_update_main["id"],d_update_main["f_update_user"],utils.GetAppid(req)}
		var rdata map[string]interface{}

		if err ==nil{
			rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
			if (err !=nil) || (rdata["f_state"] !="1") {
				err = fmt.Errorf("%s","库存变动失败！")
			}
		}
	}

	if err == nil{
		if (d_update_main["f_status"] =="0")||(d_update_main["f_status"] ==""){
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"借用归还",d_update_main["f_update_user"],d_update_main["id"],"草稿单修改","")
			//atypename string,auserid int,aid int,aopr string,amemo string
		}else {
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"借用归还",d_update_main["f_update_user"],d_update_main["id"],"草稿单修改","")
			if err ==nil {
				err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"借用归还",d_update_main["f_update_user"],d_update_main["id"],"提交","")
			}
		}
	}


	if (err != nil)  {
		rerr = err
		err = tx.Rollback()
		if err == nil{
			fmt.Println("回滚完成")
		} else {
			fmt.Println(err)
		}
		_= services.DB_Log(utils.GetAppid(req),conn,"借用归还","修改","",reqs["f_update_user"].(string),"修改失败",err)
		if utils.ISLog{
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI,"修改失败，"+rerr.Error())
		}
	} else {
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"借用归还","修改",reqs["f_update_user"].(string),reqs["f_update_user"].(string),"修改成功",err)
	}
	return
}
//借用归还删除
func (ctl *AssetsApplyModel) AssetsLendInDel(req beego.Controller)(rerr error){
	var reqs map[string]interface{};
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var conn = services.NewConnection(utils.GetAppid(req))
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

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
	//检查单据状态
	err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_lend_in_main","f_status",reqs["id"].(string),"0")
	if err != nil {
		fmt.Println(err)
		panic("单据已生效不允许删除！")
	}


	tx,_ := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	sqlstr := "update tbl_lend_in_main set  f_flag = 1, f_delete_user = "+reqs["f_delete_user"].(string)+",f_delete_time = now() where id ="+reqs["id"].(string)

	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)
	//删除从单
	if err == nil{
		sqlstr = " update tbl_lend_in_detail set  f_flag = 1, f_delete_user = "+reqs["f_delete_user"].(string)+",f_delete_time = now()  WHERE f_mid ="+reqs["id"].(string)
		_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr,)
	}

	if err !=nil{
		rerr = err
		rerr = tx.Rollback()
		_= services.DB_Log(utils.GetAppid(req),conn,"借用归还","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除失败",err)
	} else{

		rerr = nil
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"借用归还","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除成功 单号 ="+reqs["f_serno"].(string),err)
	}
	return
}
//借用归还 状态操作
func (ctl *AssetsApplyModel) AssetsLendInSubmit_Back(req beego.Controller,atype string)(rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接
	var newstate string
	var rkstate  string
	var billMemo string

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
			return
		}
	}()
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	stype := atype  //    -1辙消提交

	//状态检测
	if stype == "-1"{  //辙消提交
		err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_lend_in_main","f_status",reqs["id"].(string),"1")
		if err != nil {
			fmt.Println(err)
			panic("单据不在已提交状态不能辙消！")
		}
	}

	switch stype{  // -1辙消提交
	case "-1":{
		newstate ="0"
		rkstate = "0"
		billMemo = "辙消提交"}
	default:
		newstate =""
	}

	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,"update tbl_lend_in_main set f_status = ?,f_audit =? where id =? ",newstate,rkstate,reqs["id"].(string))

	// 修改成功后 资产位置要更改到领用部门
	if (err == nil) {
		var aparam = []string{"4", reqs["id"].(string),reqs["userid"].(string),utils.GetAppid(req)}
		var rdata map[string]interface{}

		if err ==nil{
			rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
			if (err !=nil) || (rdata["f_state"] !="1") {
				err = fmt.Errorf("%s","库存变动失败！")
			}
		}
	}

	if err == nil{
		err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产退库",reqs["userid"].(string),reqs["id"].(string),billMemo,reqs["memo"].(string))
	}
	if err == nil{
		err= services.DB_Log(utils.GetAppid(req),conn,"辙消提交",billMemo,"",reqs["userid"].(string),"辙消成功",err)
	}

	if err ==nil{
		_ = tx.Commit()
	}else {
		_ = tx.Rollback();
		rerr = err
		if utils.ISLog{
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI,"操作失败，"+rerr.Error())
		}
	}

	return rerr
}



//调拨主单查询
func (ctl *AssetsApplyModel) AssetsAllot_MainQuery(req beego.Controller) (result utils.ResData, rerr error) {
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))
	var sortmap map[string]interface{};
	var acount string
	var icount int
	var sqlstr_ord string
	var sqlstr string
	var gjwhere string
	result = utils.ResData{};

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("AssetsLendIn_MainQuery 失败，",err)
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
		}
	}()

	wvalue := req.GetString("w_value","")
	pageindex := req.GetString("curPage","")//reqs["pageindex"]
	pagesize := req.GetString("pageSize","")//reqs["pagesize"]
	userid := req.GetString("userid","")//reqs["pagesize"]
	sort := req.GetString("sort");
	SpreadWhere := req.GetString("SpreadWhere");
	if SpreadWhere != ""{ //高级筛选处理
		gjwhere = ""
		rjson,err := simplejson.NewJson([]byte(SpreadWhere));
		if err == nil{
			if rjson.Get("f_status").MustString() != ""{
				gjwhere = gjwhere+" and m.f_status ="+rjson.Get("f_status").MustString()
			}
			//调出部门
			if rjson.Get("f_Out_Deptid").MustString() != ""{ //调入部门
				gjwhere = gjwhere+" and m.f_Out_Deptid ="+rjson.Get("f_Out_Deptid").MustString()
			}
			if rjson.Get("f_custody_deptid").MustString() != ""{ //调入部门
				gjwhere = gjwhere+" and m.f_IN_Deptid ="+rjson.Get("f_custody_deptid").MustString()
			}

			if rjson.Get("f_create_username") .MustString() != ""{
				gjwhere = gjwhere+" and  m.f_create_user in (SELECT id FROM tbl_user tu WHERE tu.f_customid = "+utils.GetAppid(req)+" AND tu.f_flag = 0 and f_name LIKE '%"+rjson.Get("f_create_username").MustString()+"%')"
			}

			if (rjson.Get("f_DB_bdate") .MustString() != "")&&(rjson.Get("f_DB_edate") .MustString() != ""){
				gjwhere = gjwhere+" and m.f_date >= '"+rjson.Get("f_DB_bdate").MustString()+" 00:00:00"+
					"' and m.f_date <= '"+rjson.Get("f_DB_edate").MustString()+" 23:59:59"+"'"
			}
			gjwhere = gjwhere +" "
		}
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
			sqlstr_ord = " order by m."+sortmap["sortprop"].(string)+" "+ord
		}
	} else {
		sqlstr_ord = " order by m.id desc "
	}



	sqlstr = "CALL Msp_BillMain_Query (6,0,\""+wvalue+"\","+userid+",\""+gjwhere+"\","+pageindex+","+pagesize+",\""+sqlstr_ord+"\")"


	acount,rerr = services.Get_Procedure_Data(utils.GetAppid(req),conn,sqlstr,"icount")
	if rerr != nil{
		panic("查询合计数出错，"+rerr.Error())
	}

	icount,rerr = strconv.Atoi(acount)
	if rerr != nil{
		panic("获取合计数出错 ，"+rerr.Error())
	}

	sqlstr = "CALL Msp_BillMain_Query (6,1,\""+wvalue+"\","+userid+",\""+gjwhere+"\","+pageindex+","+pagesize+",\""+sqlstr_ord+"\")"

	result.Result,rerr = services.FindToList(utils.GetAppid(req),conn,sqlstr)
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
//调拨从单查询
func (ctl *AssetsApplyModel) AssetsAllot_DetailQuery(req beego.Controller) (result utils.ResData, rerr error) {
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))

	var acount string
	var icount int

	var sqlstr string
	result = utils.ResData{};

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("AssetsAllot_DetailQuery 失败，",err)
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
		}
	}()

	smid := req.GetString("mid","")
	userid := req.GetString("userid","")//reqs["pageindex"]

	sqlstr = "CALL Msp_BillDetail_Query (6,"+smid+","+userid+",0)"


	acount,rerr = services.Get_Procedure_Data(utils.GetAppid(req),conn,sqlstr,"icount")
	if rerr != nil{
		panic("查询合计数出错，"+rerr.Error())
	}

	icount,rerr = strconv.Atoi(acount)
	if rerr != nil{
		panic("获取合计数出错 ，"+rerr.Error())
	}

	sqlstr = "CALL Msp_BillDetail_Query (6,"+smid+","+userid+",1)"
	result.Result,rerr = services.FindToList(utils.GetAppid(req),conn,sqlstr)
	if rerr!=nil{
		panic("查询失败,"+rerr.Error())
	}else{
		result.PageSize = icount
		result.CurPage =1
		result.Totals = icount
		result.PageSize = 0
		result.CurPage = 0
		result.Totals = icount
		if reqs["sort"] != nil{
			result.Sort = reqs["sort"].(map[string]interface{})
		}
	}
	//fmt.Println(result)
	return result,nil

}
//调拨新增
func (ctl *AssetsApplyModel) AssetsAllotAdd(req beego.Controller)(aid string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接
	var s string
	var d_total float64


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
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	json_str,err:= json.Marshal(reqs)
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}
	rjson, err := simplejson.NewJson([]byte(json_str));

	if err != nil{
		panic("传入参数错误，JSON转换出错："+err.Error())
	}
	//先做数据检测
	rows_chk, err := rjson.Get("splist").Array();
	if err == nil{
		for _, row := range rows_chk {
			if each_map, ok := row.(map[string] interface{}); ok {
				if each_map["f_qty"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 数量不能为空！")
				}
				if each_map["f_price"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 价值不能为空！")
				}
			}
		}
	}

	if (rows_chk==nil) || (len(rows_chk)==0){
		panic("不能保存空单据！")
	}


	b_main := make(map[string] interface{});
	//获了单号
	b_main["f_accountid"] = "1"
	s,err = services.Get_Procedure_Data(utils.GetAppid(req),conn," CALL Msp_Get_Serno ("+utils.GetAppid(req)+",'tbl_db_main','f_serno','DB') ","f_serno")
	if err != nil{
		panic("单号获取出错,"+err.Error())
	}
    if rjson.Get("f_status").MustString() != ""{
    	if rjson.Get("f_status").MustInt() > 1 {
    		panic("单据状态错误！")
		}
	}

	b_main["f_serno"] =s
	b_main["f_status"] = rjson.Get("f_status").MustString() // 0状态
	b_main["f_type"] = rjson.Get("f_type").MustString() // 0普通调拨
	b_main["f_deptid"] = rjson.Get("f_deptid").MustString() //制单部门
	b_main["f_Out_Deptid"] = rjson.Get("f_Out_Deptid").MustString() //调出部门
	b_main["f_IN_Deptid"] = rjson.Get("f_IN_Deptid").MustString() //调入部门
	b_main["f_date"] = rjson.Get("f_date").MustString() //单据日期
	b_main["f_note"] = rjson.Get("f_note").MustString()
	b_main["f_create_user"] = rjson.Get("f_create_user").MustString()
	now := time.Now()
	b_main["f_create_time"]= now.Format("2006-01-02 15:04:05");
	b_main["f_customid"]= utils.GetAppid(req)
	b_main["f_flag"]= "0"
	b_main["f_In_postion"] = rjson.Get("f_In_postion").MustString()

	//开始事务
	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	//插入主单
	aid,err = services.One_ADD_Table(utils.GetAppid(req),tx,conn,"tbl_db_main",b_main)


	//插入从单
	if err ==nil{
		rows, _ := rjson.Get("splist").Array();
		qxsql :=" INSERT INTO tbl_db_detail(f_mid,f_spid,f_qty,f_price,f_total,f_Out_Deptid,f_custody_deptid,"+
			" 							    f_custody_userid,f_postion,f_note,f_create_user,f_create_time,f_customid ) "+
			" values(?,?,?,?,?,?,?,?,?,?,?,now(),?) "

		for _, row := range rows {
			if each_map, ok := row.(map[string] interface{}); ok {

				if each_map["f_price"].(string) == ""{
					each_map["f_price"] = "0"
				}
				if each_map["f_custody_deptid"].(string) == ""{
					each_map["f_custody_deptid"] = "-1"
				}
				if each_map["f_custody_userid"].(string) == ""{
					each_map["f_custody_userid"] = "-1"
				}
				if each_map["f_postion"].(string) == ""{
					each_map["f_postion"] = "-1"
				}

				d_qty,_ := strconv.ParseFloat(each_map["f_qty"].(string),64)
				d_price,_ := strconv.ParseFloat(each_map["f_price"].(string),64)
				d_total = d_qty * d_price
				_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,qxsql,aid,each_map["f_spid"].(string),each_map["f_qty"].(string),each_map["f_price"].(string),
					d_total,each_map["f_Out_Deptid"].(string),each_map["f_custody_deptid"].(string),
					each_map["f_custody_userid"].(string),each_map["f_postion"].(string),each_map["f_note"].(string),rjson.Get("f_create_user").MustString(),utils.GetAppid(req))
				if err !=nil{
					break;
				}
			}
		}
	}

	//// 新增成功后 资产位置要更改到领用部门
	//if (err == nil) && (b_main["f_status"] == "1"){
	//	var aparam = []string{"6", aid,b_main["f_create_user"].(string),utils.GetAppid(req)}
	//	var rdata map[string]interface{}
	//
	//	if err ==nil{
	//		rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
	//		if (err !=nil) || (rdata["f_state"] !="1") {
	//			err = fmt.Errorf("%s","库存变动失败！")
	//		}
	//	}
	//}

	if err == nil{
		if (rjson.Get("f_status").MustString() =="0")||(rjson.Get("f_status").MustString() ==""){
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产调拨",b_main["f_create_user"].(string),aid,"新增草稿单","")
			//atypename string,auserid int,aid int,aopr string,amemo string
		}else {
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产调拨",b_main["f_create_user"].(string),aid,"新增草稿单","")
			if err ==nil {
				err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产调拨",b_main["f_create_user"].(string),aid,"提交","")
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
		_= services.DB_Log(utils.GetAppid(req),conn,"资产调拨","新增","",reqs["f_create_user"].(string),"新增失败",err)
		if utils.ISLog{
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI,"新增失败，"+rerr.Error())
		}
	} else {
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产调拨","新增",reqs["f_create_user"].(string),reqs["f_create_user"].(string),"新增成功 ID ="+aid,err)
	}
	return

}
//调拨修改
func (ctl *AssetsApplyModel) AssetsAllotUpdate(req beego.Controller)(alen string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接
	var d_total float64

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
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

	if err != nil{
		panic("传入参数错误，JSON转换出错："+err.Error())
	}
	//先做数据检测
	rows_chk, err := rjson.Get("splist").Array();
	if err == nil{
		for _, row := range rows_chk {
			if each_map, ok := row.(map[string] interface{}); ok {
				if each_map["f_qty"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 数量不能为空！")
				}
				if each_map["f_price"].(string) ==""{
					panic("资产 "+each_map["f_code"].(string)+" 价值不能为空！")
				}
			}
		}
	}

	d_update_main := make(map[string] string);
	//获了单号
	d_update_main["id"] = rjson.Get("id").MustString()

	//检查单据状态
	err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_db_main","f_status",d_update_main["id"],"0")
	if err != nil {
		fmt.Println(err)
		panic("单据已生效不允许修改！")
	}
	if rjson.Get("f_status").MustString() != ""{
		if rjson.Get("f_status").MustInt() > 1 {
			panic("单据状态错误！")
		}
	}

	d_update_main["f_status"]  = rjson.Get("f_status").MustString()
	//if (d_update_main["f_status"] =="1"){
	//	d_update_main["f_audit"] ="1"
	//	now := time.Now()
	//	d_update_main["f_audit_date"] = now.Format("2006-01-02 15:04:05");
	//	d_update_main["f_status"] ="8"
	//}
	d_update_main["f_serno"] =  rjson.Get("f_serno").MustString()
	d_update_main["f_type"] = rjson.Get("f_type").MustString() // 0普通调拨
	d_update_main["f_deptid"] = rjson.Get("f_deptid").MustString() //制单部门
	d_update_main["f_Out_Deptid"] = rjson.Get("f_Out_Deptid").MustString() //调出部门
	d_update_main["f_IN_Deptid"] = rjson.Get("f_IN_Deptid").MustString() //调入部门
	d_update_main["f_date"] = rjson.Get("f_date").MustString() //单据日期
	d_update_main["f_note"] =  rjson.Get("f_note").MustString()
	d_update_main["f_update_user"] =  rjson.Get("f_update_user").MustString()
	d_update_main["f_update_time"] =  time.Now().Format("2006-01-02 15:04:05")
	d_update_main["f_In_postion"] =  rjson.Get("f_In_postion").MustString()

	sqlstr :=  services.Sql_JionSet(d_update_main)
	if sqlstr !="" {
		sqlstr = " update tbl_db_main set "+sqlstr +" where id ="+d_update_main["id"]
	}

	//开始事务
	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	//修改角色
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)

	if err != nil{
		panic("主单修改失败,"+err.Error())
	}
	//先删除从单
	sqlstr = " delete from tbl_db_detail where f_mid =  "+d_update_main["id"]
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)

	//插入从单
	if err ==nil{
		rows, _ := rjson.Get("splist").Array();
		qxsql :=" INSERT INTO tbl_db_detail(f_mid,f_spid,f_qty,f_price,f_total,f_Out_Deptid,f_custody_deptid,"+
			" 							    f_custody_userid,f_postion,f_note,f_create_user,f_create_time,f_customid ) "+
			" values(?,?,?,?,?,?,?,?,?,?,?,now(),?) "

		for _, row := range rows {
			if each_map, ok := row.(map[string] interface{}); ok {

				if each_map["f_price"].(string) == ""{
					each_map["f_price"] = "0"
				}
				if each_map["f_custody_deptid"].(string) == ""{
					each_map["f_custody_deptid"] = "-1"
				}
				if each_map["f_custody_userid"].(string) == ""{
					each_map["f_custody_userid"] = "-1"
				}
				if each_map["f_postion"].(string) == ""{
					each_map["f_postion"] = "-1"
				}

				d_qty,_ := strconv.ParseFloat(each_map["f_qty"].(string),64)
				d_price,_ := strconv.ParseFloat(each_map["f_price"].(string),64)
				d_total = d_qty * d_price
				_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,qxsql,d_update_main["id"],each_map["f_spid"].(string),each_map["f_qty"].(string),each_map["f_price"].(string),
					d_total,each_map["f_Out_Deptid"].(string),each_map["f_custody_deptid"].(string),
					each_map["f_custody_userid"].(string),each_map["f_postion"].(string),each_map["f_note"].(string),d_update_main["f_update_user"] ,utils.GetAppid(req))
				if err !=nil{
					break;
				}
			}
		}
	}

	//// 修改成功后 资产位置要更改到领用部门
	//if (err == nil) && (d_update_main["f_status"] == "1"){
	//	var aparam = []string{"6", d_update_main["id"],d_update_main["f_update_user"],utils.GetAppid(req)}
	//	var rdata map[string]interface{}
	//
	//	if err ==nil{
	//		rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
	//		if (err !=nil) || (rdata["f_state"] !="1") {
	//			err = fmt.Errorf("%s","库存变动失败！")
	//		}
	//	}
	//}

	if err == nil{
		if (d_update_main["f_status"] =="0")||(d_update_main["f_status"] ==""){
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产调拨",d_update_main["f_update_user"],d_update_main["id"],"草稿单修改","")
			//atypename string,auserid int,aid int,aopr string,amemo string
		}else {
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产调拨",d_update_main["f_update_user"],d_update_main["id"],"草稿单修改","")
			if err ==nil {
				err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产调拨",d_update_main["f_update_user"],d_update_main["id"],"提交","")
			}
		}
	}


	if (err != nil)  {
		rerr = err
		err = tx.Rollback()
		if err == nil{
			fmt.Println("回滚完成")
		} else {
			fmt.Println(err)
		}
		_= services.DB_Log(utils.GetAppid(req),conn,"资产调拨","修改","",reqs["f_update_user"].(string),"修改失败",err)
		if utils.ISLog{
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI,"修改失败，"+rerr.Error())
		}
	} else {
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产调拨","修改",reqs["f_update_user"].(string),reqs["f_update_user"].(string),"修改成功",err)
	}
	return
}
//调拨删除
func (ctl *AssetsApplyModel) AssetsAllotDel(req beego.Controller)(rerr error){
	var reqs map[string]interface{};
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var conn = services.NewConnection(utils.GetAppid(req))
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

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
	//检查单据状态
	err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_db_main","f_status",reqs["id"].(string),"0")
	if err != nil {
		fmt.Println(err)
		panic("单据已生效不允许删除！")
	}


	tx,_ := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	sqlstr := "update tbl_db_main set  f_flag = 1, f_delete_user = "+reqs["f_delete_user"].(string)+",f_delete_time = now() where id ="+reqs["id"].(string)

	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)
	//删除从单
	if err == nil{
		sqlstr = " update tbl_db_detail set  f_flag = 1, f_delete_user = "+reqs["f_delete_user"].(string)+",f_delete_time = now()  WHERE f_mid ="+reqs["id"].(string)
		_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr,)
	}

	if err !=nil{
		rerr = err
		rerr = tx.Rollback()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产调拨","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除失败",err)
	} else{

		rerr = nil
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产调拨","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除成功 单号 ="+reqs["f_serno"].(string),err)
	}
	return
}
//调拨 审批
func (ctl *AssetsApplyModel) AssetsAllot_Audit(req beego.Controller)(rerr error){
	var reqs map[string]interface{};
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var conn = services.NewConnection(utils.GetAppid(req))
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

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
	//检查单据状态
	err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_lend_in_main","f_status",reqs["id"].(string),"1")
	if err != nil {
		fmt.Println(err)
		panic("单据不在已提交状态不能审批！")
	}


	tx,_ := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	sqlstr := "update tbl_db_main set  f_status = 2  where id ="+reqs["id"].(string)

	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)

	if err !=nil{
		rerr = err
		rerr = tx.Rollback()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产调拨","审批",reqs["id"].(string),reqs["f_update_user"].(string),"审批失败",err)
	} else{

		rerr = nil
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产调拨","审批",reqs["id"].(string),reqs["f_update_user"].(string),"审批成功",err)
	}
	return
}
//调拨 调出
func (ctl *AssetsApplyModel) AssetsAllot_StockOut(req beego.Controller)(rerr error){
	var reqs map[string]interface{};
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var conn = services.NewConnection(utils.GetAppid(req))
	var serrs string
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

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
	avalue,verr := services.Get_FieldByValue(utils.GetAppid(req),conn,"SELECT tsd.f_status FROM tbl_sys_dictionary tsd WHERE tsd.f_customid = "+utils.GetAppid(req)+" AND tsd.f_type = 6 AND f_id = 2","f_status")
	if verr != nil{
		panic("单据状态获取出错！")
	}
	//检查单据状态
	if avalue =="1"{ //启用了审批功能
		err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_db_main","f_status",reqs["id"].(string),"2")
		serrs = "审批"
	}else{
		err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_db_main","f_status",reqs["id"].(string),"1") //没启用审批，判断是否提交
		serrs = "提交"
	}
	if err != nil {
		fmt.Println(err)
		panic("单据不在"+serrs+"状态不能调出！")
	}

	tx,_ := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	sqlstr := "update tbl_db_main set  f_status = 8,f_audit = 1,f_audit_date =now(),f_audit_userid ="+reqs["f_update_user"].(string)+" where id ="+reqs["id"].(string)

	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)

	// 修改成功后 资产位置要更改到领用部门
	var aparam = []string{"6", reqs["id"].(string),reqs["f_update_user"].(string),utils.GetAppid(req)}
	var rdata map[string]interface{}

	if err ==nil{
		rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
		if (err !=nil) || (rdata["f_state"] !="1") {
			err = fmt.Errorf("%s","库存变动失败！")
		}
	}

	if err !=nil{
		rerr = err
		rerr = tx.Rollback()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产调拨","确认调出",reqs["id"].(string),reqs["f_update_user"].(string),"调出失败",err)
	} else{

		rerr = nil
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产调拨","确认调出",reqs["id"].(string),reqs["f_update_user"].(string),"调出成功 单号 ="+reqs["f_serno"].(string),err)
	}
	return
}
//调拨 调入
func (ctl *AssetsApplyModel) AssetsAllot_StockIn(req beego.Controller)(rerr error){

	var reqs map[string]interface{};
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var conn = services.NewConnection(utils.GetAppid(req))
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

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
	//检查单据状态
	err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_db_main","f_audit",reqs["id"].(string),"1")
	if err != nil {
		fmt.Println(err)
		panic("单据不在已调出状态不能调入！")
	}


	tx,_ := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	sqlstr := "update tbl_db_main set  f_status = 8,f_audit = 2,f_audit_date =now(),f_audit_userid ="+reqs["f_update_user"].(string)+" where id ="+reqs["id"].(string)

	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)


	var aparam = []string{"6", reqs["id"].(string),reqs["f_update_user"].(string),utils.GetAppid(req)}
	var rdata map[string]interface{}

	if err ==nil{
		rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
		if (err !=nil) || (rdata["f_state"] !="1") {
			err = fmt.Errorf("%s","库存变动失败！")
		}
	}


	if err !=nil{
		rerr = err
		rerr = tx.Rollback()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产调拨","确认调入",reqs["id"].(string),reqs["f_update_user"].(string),"调入失败",err)
	} else{

		rerr = nil
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产调拨","确认调入",reqs["id"].(string),reqs["f_update_user"].(string),"调入成功 单号 ="+reqs["f_serno"].(string),err)
	}
	return
}
//调拨 状态操作
func (ctl *AssetsApplyModel) AssetsAllotSubmit_Back(req beego.Controller)(rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接
	var newstate string
	var rkstate  string
	var billMemo string
	var oldstate string
	var oldaudit string
	var old_int_state int
	var old_int_audit int

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
			return
		}
	}()
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	stype := reqs["newstate"].(string)  //    0辙消提交

	oldstate,err = services.Get_FieldByValue(utils.GetAppid(req),conn,"SELECT tdm.f_status FROM tbl_DB_Main tdm WHERE id =  "+reqs["id"].(string),"f_status");
	if err != nil{
		panic("获取单据状态出错！")
	}
	oldaudit,err = services.Get_FieldByValue(utils.GetAppid(req),conn,"SELECT tdm.f_audit FROM tbl_DB_Main tdm WHERE id =  "+reqs["id"].(string),"f_audit");
	if err != nil{
		panic("获取单据出库状态出错！")
	}else{
		old_int_state,err = strconv.Atoi(oldstate)
		if err != nil{
			panic("数据类型转换失败，"+err.Error())
		}
		old_int_audit,err =  strconv.Atoi(oldaudit)
		if err != nil{
			panic("数据类型转换失败，"+err.Error())
		}
	}


	//状态检测
	if stype == "0"{  //辙消到草稿
		if oldstate =="0"{  //可以一次性辙消到草稿
			panic("当前单据状态不允许辙消到草稿！")
		}
	}
	if stype == "1"{  //辙消到已提交
		if old_int_state <= 1 {
			panic("当前单据状态不允许辙消到已提交！")
		}
	}
	if stype == "2"{  //辙消到已审批
		if old_int_state <= 2 {
			panic("当前单据状态不允许辙消到已审批！")
		}
	}
	if stype == "2"{  //辙消到已调出
		if old_int_audit != 2 {
			panic("当前单据未调入，不能辙消到已调出！")
		}
	}
	switch stype{  // 0辙消提交
	case "0":{
		newstate ="0"
		rkstate = "0"
		billMemo = "辙消到草稿"}
	case "1":{
		newstate ="1"
		rkstate = "0"
		billMemo = "辙消到已提交"}
	case "2":{
		newstate ="2"
		rkstate = "0"
		billMemo = "辙消到已审批"}
	case "3":{
		newstate ="8"
		rkstate = "1"
		billMemo = "辙消到已调出"}
	default:
		newstate =""
	}

	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()

	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,"update tbl_db_main set f_status = ?,f_audit =? where id =? ",newstate,rkstate,reqs["id"].(string))

	// 修改成功后 资产位置要更改到领用部门
	if (err == nil) {
		var aparam = []string{"6", reqs["id"].(string),reqs["userid"].(string),utils.GetAppid(req)}
		var rdata map[string]interface{}

		if err ==nil{
			rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
			if (err !=nil) || (rdata["f_state"] !="1") {
				err = fmt.Errorf("%s","库存变动失败！")
			}
		}
	}

	if (err == nil)&&(stype=="3"){ // 辙消到已调出要变动库存
		var aparam = []string{"6", reqs["id"].(string),reqs["f_update_user"].(string),utils.GetAppid(req)}
		var rdata map[string]interface{}

		if err ==nil{
			rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
			if (err !=nil) || (rdata["f_state"] !="1") {
				err = fmt.Errorf("%s","库存变动失败！")
			}
		}
	}

	if err == nil{
		err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产调拨",reqs["userid"].(string),reqs["id"].(string),billMemo,reqs["memo"].(string))
	}
	if err == nil{
		err= services.DB_Log(utils.GetAppid(req),conn,"资产调拨",billMemo,"",reqs["userid"].(string),"辙消成功",err)
	}

	if err ==nil{
		_ = tx.Commit()
	}else {
		_ = tx.Rollback();
		rerr = err
		if utils.ISLog{
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI,"操作失败，"+rerr.Error())
		}
	}

	return rerr
}



//资产变动 主单查询
func (ctl *AssetsApplyModel) AssetsChange_MainQuery(req beego.Controller) (result utils.ResData, rerr error) {
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))
	var sortmap map[string]interface{};
	var acount string
	var icount int
	var sqlstr_ord string
	var sqlstr string
	var gjwhere string

	result = utils.ResData{};

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("AssetsChange_MainQuery 失败，",err)
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
		}
	}()

	wvalue := req.GetString("w_value","")
	pageindex := req.GetString("curPage","")//reqs["pageindex"]
	pagesize := req.GetString("pageSize","")//reqs["pagesize"]
	userid := req.GetString("userid","")//reqs["pagesize"]
	sort := req.GetString("sort");
	SpreadWhere := req.GetString("SpreadWhere");
	if SpreadWhere != ""{ //高级筛选处理
		gjwhere = ""
		rjson,err := simplejson.NewJson([]byte(SpreadWhere));
		if err == nil{
			if rjson.Get("f_status").MustString() != ""{
				gjwhere = gjwhere+" and m.f_status ="+rjson.Get("f_status").MustString()
			}
			if rjson.Get("f_To_Deptid").MustString() != ""{ //变更后所属部门
				gjwhere = gjwhere+" and m.f_To_Deptid ="+rjson.Get("f_To_Deptid").MustString()
			}
			if rjson.Get("f_create_username") .MustString() != ""{
				gjwhere = gjwhere+" and  m.f_create_user in (SELECT id FROM tbl_user tu WHERE tu.f_customid = "+utils.GetAppid(req)+" AND tu.f_flag = 0 and f_name LIKE '%"+rjson.Get("f_create_username").MustString()+"%')"
			}
			if (rjson.Get("f_change_bdate") .MustString() != "")&&(rjson.Get("f_change_edate") .MustString() != ""){
				gjwhere = gjwhere+" and m.f_create_time >= '"+rjson.Get("f_change_bdate").MustString()+" 00:00:00"+
					"' and m.f_create_time <= '"+rjson.Get("f_change_edate").MustString()+" 23:59:59"+"'"
			}
			gjwhere = gjwhere +" "
		}
	}
	//f_status 状态   f_change_bdate 变更日期起  f_change_edate 变更日期止
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
			sqlstr_ord = " order by m."+sortmap["sortprop"].(string)+" "+ord
		}
	} else {
		sqlstr_ord = " order by m.id desc "
	}



	sqlstr = "CALL Msp_BillMain_Query (5,0,\""+wvalue+"\","+userid+",\""+gjwhere+"\","+pageindex+","+pagesize+",\""+sqlstr_ord+"\")"


	acount,rerr = services.Get_Procedure_Data(utils.GetAppid(req),conn,sqlstr,"icount")
	if rerr != nil{
		panic("查询合计数出错，"+rerr.Error())
	}

	icount,rerr = strconv.Atoi(acount)
	if rerr != nil{
		panic("获取合计数出错 ，"+rerr.Error())
	}

	sqlstr = "CALL Msp_BillMain_Query (5,1,\""+wvalue+"\","+userid+",\""+gjwhere+"\","+pageindex+","+pagesize+",\""+sqlstr_ord+"\")"

	result.Result,rerr = services.FindToList(utils.GetAppid(req),conn,sqlstr)
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
//资产变动 从单查询
func (ctl *AssetsApplyModel) AssetsChange_DetailQuery(req beego.Controller) (result utils.ResData, rerr error) {
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req))

	var acount string
	var icount int

	var sqlstr string
	result = utils.ResData{};

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("AssetsChange_DetailQuery 失败，",err)
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
		}
	}()

	smid := req.GetString("mid","")
	userid := req.GetString("userid","")//reqs["pageindex"]

	sqlstr = "CALL Msp_BillDetail_Query (5,"+smid+","+userid+",0)"


	acount,rerr = services.Get_Procedure_Data(utils.GetAppid(req),conn,sqlstr,"icount")
	if rerr != nil{
		panic("查询合计数出错，"+rerr.Error())
	}

	icount,rerr = strconv.Atoi(acount)
	if rerr != nil{
		panic("获取合计数出错 ，"+rerr.Error())
	}

	sqlstr = "CALL Msp_BillDetail_Query (5,"+smid+","+userid+",1)"
	result.Result,rerr = services.FindToList(utils.GetAppid(req),conn,sqlstr)
	if rerr!=nil{
		panic("查询失败,"+rerr.Error())
	}else{
		result.PageSize = icount
		result.CurPage =1
		result.Totals = icount
		result.PageSize = 0
		result.CurPage = 0
		result.Totals = icount
		if reqs["sort"] != nil{
			result.Sort = reqs["sort"].(map[string]interface{})
		}
	}
	//fmt.Println(result)
	return result,nil

}
//资产变动 新增
func (ctl *AssetsApplyModel) AssetsChangeAdd(req beego.Controller)(aid string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接
	var s string

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
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	json_str,err:= json.Marshal(reqs)
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}
	rjson, err := simplejson.NewJson([]byte(json_str));

	if err != nil{
		panic("传入参数错误，JSON转换出错："+err.Error())
	}
	//先做数据检测
	rows_chk, err := rjson.Get("splist").Array();

	if (rows_chk==nil) || (len(rows_chk)==0){
		panic("不能保存空单据！")
	}


	b_main := make(map[string] interface{});
	//获了单号
	b_main["f_accountid"] = "1"
	s,err = services.Get_Procedure_Data(utils.GetAppid(req),conn," CALL Msp_Get_Serno ("+utils.GetAppid(req)+",'tbl_change_main','f_serno','BG') ","f_serno")
	if err != nil{
		panic("单号获取出错,"+err.Error())
	}
	if rjson.Get("f_status").MustString() != ""{
		if rjson.Get("f_status").MustInt() > 1 {
			panic("单据状态错误！")
		}
	}
	//if rjson.Get("f_status").MustString() ==""{
	//	b_main["f_status"] = "0"
	//	b_main["f_audit"] ="0"
	//} else {
	//	b_main["f_status"] = rjson.Get("f_status").MustString()
	//	if (b_main["f_status"] =="1"){
	//		b_main["f_audit"] ="1"
	//		now := time.Now()
	//		b_main["f_audit_date"] = now.Format("2006-01-02 15:04:05");
	//		b_main["f_audit_userid"] = rjson.Get("f_create_user").MustString()
	//	}else {
	//		b_main["f_audit"] ="0"
	//	}
	//}
	b_main["f_serno"] =s
	b_main["f_deptid"] = rjson.Get("f_deptid").MustString()
	b_main["f_status"] = rjson.Get("f_status").MustString()
	b_main["f_To_name"] = rjson.Get("f_To_Name").MustString()
	b_main["f_To_lbid"] = rjson.Get("f_To_Lbid").MustString()
	b_main["f_To_Deptid"] = rjson.Get("f_To_Deptid").MustString()
	b_main["f_To_UseDeptid"] = rjson.Get("f_To_UseDeptid").MustString()
	b_main["f_To_UserID"] = rjson.Get("f_To_UserID").MustString()
	b_main["f_To_UseUserID"] = rjson.Get("f_To_UseUserID").MustString()
	b_main["f_To_Postion"] = rjson.Get("f_To_Postion").MustString()
	b_main["f_note"] = rjson.Get("f_note").MustString()
	b_main["f_create_user"] = rjson.Get("f_create_user").MustString()
	now := time.Now()
	b_main["f_create_time"]= now.Format("2006-01-02 15:04:05");
	b_main["f_customid"]= utils.GetAppid(req)
	b_main["f_flag"]= "0"

	//开始事务
	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	//插入主单
	aid,err = services.One_ADD_Table(utils.GetAppid(req),tx,conn,"tbl_change_main",b_main)

	//插入从单
	if err ==nil{
		rows, _ := rjson.Get("splist").Array();
		qxsql :=" INSERT INTO tbl_change_detail(f_mid,f_spid,f_customid,f_to_name,f_to_lbid,f_to_deptid,f_to_usedeptid,f_to_userid,f_to_useuserid,"+
				" f_to_postion,f_note,f_create_user,f_create_time,f_flag) "+
				" values(?,?,?,?,?,?,?,?,?,?,?,?,now(),0) "

		for _, row := range rows {
			if each_map, ok := row.(map[string] interface{}); ok {

				if each_map["f_To_Lbid"].(string) == ""{
					each_map["f_To_Lbid"] = "-1"
				}
				if each_map["f_To_Deptid"].(string) == ""{
					each_map["f_To_Deptid"] = "-1"
				}
				if each_map["f_To_UseDeptid"].(string) == ""{
					each_map["f_To_UseDeptid"] = "-1"
				}
				if each_map["f_To_UserID"].(string) == ""{
					each_map["f_To_UserID"] = "-1"
				}
				if each_map["f_To_UseUserID"].(string) == ""{
					each_map["f_To_UseUserID"] = "-1"
				}
				if each_map["f_To_Postion"].(string) == ""{
					each_map["f_To_Postion"] = "-1"
				}
				_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,qxsql,aid,each_map["f_spid"].(string),utils.GetAppid(req),each_map["f_To_Name"].(string),each_map["f_To_Lbid"].(string),
					each_map["f_To_Deptid"].(string),each_map["f_To_UseDeptid"].(string),each_map["f_To_UserID"].(string),each_map["f_To_UseUserID"].(string),
					each_map["f_To_Postion"].(string),"",b_main["f_create_user"].(string) )
				if err !=nil{
					break;
				}
			}
		}
	}

	// 修改成功后 资产位置要更改到领用部门
	if (err == nil) && (b_main["f_status"].(string) == "1"){
		var aparam = []string{"5", b_main["id"].(string),b_main["f_create_user"].(string),utils.GetAppid(req)}
		var rdata map[string]interface{}

		if err ==nil{
			rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
			if (err !=nil) || (rdata["f_state"] !="1") {
				err = fmt.Errorf("%s","库存变动失败！")
			}
		}

		//修改资产列表对应属性
		if err ==nil{
			sqlup := "UPDATE  tbl_spinfo ts "+
			" INNER JOIN tbl_change_detail tcd ON ts.id = tcd.f_spid "+
			" INNER JOIN tbl_change_main tcm ON tcd.f_mid = tcm.id "+
			" SET ts.f_Name = tcd.f_to_name,ts.f_Typeid = tcd.f_To_lbid "+
			" WHERE tcm.id = "+b_main["id"].(string)+" AND tcm.f_audit = 1;"
			_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlup)
			if err !=nil{
				panic("资产属性更新失败，"+err.Error())
			}
		}
	}

	if err == nil{
		if (rjson.Get("f_status").MustString() =="0")||(rjson.Get("f_status").MustString() ==""){
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产变更",b_main["f_create_user"].(string),aid,"新增草稿单","")
			//atypename string,auserid int,aid int,aopr string,amemo string
		}else {
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产变更",b_main["f_create_user"].(string),aid,"新增草稿单","")
			if err ==nil {
				err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产变更",b_main["f_create_user"].(string),aid,"提交申请","")
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
		_= services.DB_Log(utils.GetAppid(req),conn,"资产变更","新增","",reqs["f_create_user"].(string),"新增失败",err)
		if utils.ISLog{
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI,"新增失败，"+rerr.Error())
		}
	} else {
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产变更","新增",reqs["f_create_user"].(string),reqs["f_create_user"].(string),"新增成功 ID ="+aid,err)
	}
	return

}
//资产变动 修改
func (ctl *AssetsApplyModel) AssetsChangeUpdate(req beego.Controller)(alen string,rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接


	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
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

	if err != nil{
		panic("传入参数错误，JSON转换出错："+err.Error())
	}
	//先做数据检测
	rows_chk, err := rjson.Get("splist").Array();
	if (rows_chk==nil) || (len(rows_chk)==0){
		panic("不能保存空单据！")
	}


	d_update_main := make(map[string] string);
	//获了单号
	d_update_main["id"] = rjson.Get("id").MustString()

	//检查单据状态
	err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_change_main","f_status",d_update_main["id"],"0")
	if err != nil {
		fmt.Println(err)
		panic("单据已生效不允许修改！")
	}
	if rjson.Get("f_status").MustString() != ""{
		if rjson.Get("f_status").MustInt() > 1 {
			panic("单据状态错误！")
		}
	}

	//if rjson.Get("f_status").MustString() ==""{
	//	d_update_main["f_status"] = "0"
	//	d_update_main["f_audit"] ="0"
	//} else {
	//	d_update_main["f_status"] = rjson.Get("f_status").MustString()
	//	if (d_update_main["f_status"] =="1"){
	//		d_update_main["f_audit"] ="1"
	//		now := time.Now()
	//		d_update_main["f_audit_date"] = now.Format("2006-01-02 15:04:05");
	//		d_update_main["f_audit_userid"] = rjson.Get("f_update_user").MustString()
	//	}else {
	//		d_update_main["f_audit"] ="0"
	//	}
	//}

	d_update_main["f_serno"] =  rjson.Get("f_serno").MustString()
	d_update_main["f_status"] = rjson.Get("f_status").MustString()
	d_update_main["f_deptid"] = rjson.Get("f_deptid").MustString()
	d_update_main["f_To_name"] = rjson.Get("f_To_Name").MustString()
	d_update_main["f_To_lbid"] = rjson.Get("f_To_Lbid").MustString()
	d_update_main["f_To_Deptid"] = rjson.Get("f_To_Deptid").MustString()
	d_update_main["f_To_UseDeptid"] = rjson.Get("f_To_UseDeptid").MustString()
	d_update_main["f_To_UserID"] = rjson.Get("f_To_UserID").MustString()
	d_update_main["f_To_UseUserID"] = rjson.Get("f_To_UseUserID").MustString()
	d_update_main["f_To_Postion"] = rjson.Get("f_To_Postion").MustString()
	d_update_main["f_note"] =  rjson.Get("f_note").MustString()
	d_update_main["f_update_user"] =  rjson.Get("f_update_user").MustString()
	d_update_main["f_update_time"] =  time.Now().Format("2006-01-02 15:04:05")


	sqlstr :=  services.Sql_JionSet(d_update_main)
	if sqlstr !="" {
		sqlstr = " update tbl_change_main set "+sqlstr +" where id ="+d_update_main["id"]
	}

	//开始事务
	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	//修改角色
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)

	if err != nil{
		panic("主单修改失败,"+err.Error())
	}
	//先删除从单
	sqlstr = " delete from tbl_change_detail where f_mid =  "+d_update_main["id"]
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)

	//插入从单
	if err ==nil{
		rows, _ := rjson.Get("splist").Array();
		qxsql :=" INSERT INTO tbl_change_detail(f_mid,f_spid,f_customid,f_to_name,f_to_lbid,f_to_deptid,f_to_usedeptid,f_to_userid,f_to_useuserid,"+
				" f_to_postion,f_note,f_create_user,f_create_time,f_flag) "+
				" values(?,?,?,?,?,?,?,?,?,?,?,?,now(),0) "

		for _, row := range rows {
			if each_map, ok := row.(map[string] interface{}); ok {

				if each_map["f_To_Lbid"].(string) == ""{
					each_map["f_To_Lbid"] = "-1"
				}
				if each_map["f_To_Deptid"].(string) == ""{
					each_map["f_To_Deptid"] = "-1"
				}
				if each_map["f_To_UseDeptid"].(string) == ""{
					each_map["f_To_UseDeptid"] = "-1"
				}
				if each_map["f_To_UserID"].(string) == ""{
					each_map["f_To_UserID"] = "-1"
				}
				if each_map["f_To_UseUserID"].(string) == ""{
					each_map["f_To_UseUserID"] = "-1"
				}
				if each_map["f_To_Postion"].(string) == ""{
					each_map["f_To_Postion"] = "-1"
				}
				_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,qxsql,d_update_main["id"],each_map["f_spid"].(string),utils.GetAppid(req),each_map["f_To_Name"].(string),each_map["f_To_Lbid"].(string),
					each_map["f_To_Deptid"].(string),each_map["f_To_UseDeptid"].(string),each_map["f_To_UserID"].(string),each_map["f_To_UseUserID"].(string),
					each_map["f_To_Postion"].(string),"",d_update_main["f_update_user"] )
				if err !=nil{
					break;
				}
			}
		}
	}

	//// 修改成功后 资产位置要更改到领用部门
	//if (err == nil) && (d_update_main["f_status"] == "1"){
	//	var aparam = []string{"5", d_update_main["id"],d_update_main["f_update_user"],utils.GetAppid(req)}
	//	var rdata map[string]interface{}
	//
	//	if err ==nil{
	//		rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
	//		if (err !=nil) || (rdata["f_state"] !="1") {
	//			err = fmt.Errorf("%s","库存变动失败！")
	//		}
	//	}
	//
	//	//修改资产列表对应属性
	//	if err ==nil{
	//		sqlup := "UPDATE  tbl_spinfo ts "+
	//			" INNER JOIN tbl_change_detail tcd ON ts.id = tcd.f_spid "+
	//			" INNER JOIN tbl_change_main tcm ON tcd.f_mid = tcm.id "+
	//			" SET ts.f_Name = tcd.f_to_name,ts.f_Typeid = tcd.f_To_lbid "+
	//			" WHERE tcm.id = "+d_update_main["id"]+" AND tcm.f_audit = 1;"
	//		_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlup)
	//		if err !=nil{
	//			panic("资产属性更新失败，"+err.Error())
	//		}
	//	}
	//}

	if err == nil{
		if (d_update_main["f_status"] =="0")||(d_update_main["f_status"] ==""){
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产变更",d_update_main["f_update_user"],d_update_main["id"],"草稿单修改","")
			//atypename string,auserid int,aid int,aopr string,amemo string
		}else {
			err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产变更",d_update_main["f_update_user"],d_update_main["id"],"草稿单修改","")
			if err ==nil {
				err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产变更",d_update_main["f_update_user"],d_update_main["id"],"提交申请","")
			}
		}
	}


	if (err != nil)  {
		rerr = err
		err = tx.Rollback()
		if err == nil{
			fmt.Println("回滚完成")
		} else {
			fmt.Println(err)
		}
		_= services.DB_Log(utils.GetAppid(req),conn,"资产变更","修改","",reqs["f_update_user"].(string),"修改失败",err)
		if utils.ISLog{
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI,"修改失败，"+rerr.Error())
		}
	} else {
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产变更","修改",reqs["f_update_user"].(string),reqs["f_update_user"].(string),"修改成功",err)
	}
	return
}
//资产变动 删除
func (ctl *AssetsApplyModel) AssetsChangeDel(req beego.Controller)(rerr error){
	var reqs map[string]interface{};
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var conn = services.NewConnection(utils.GetAppid(req))
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

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
	//检查单据状态
	err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_change_main","f_status",reqs["id"].(string),"0")
	if err != nil {
		fmt.Println(err)
		panic("单据已生效不允许删除！")
	}


	tx,_ := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	sqlstr := "update tbl_change_main set  f_flag = 1, f_delete_user = "+reqs["f_delete_user"].(string)+",f_delete_time = now() where id ="+reqs["id"].(string)

	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)
	//删除从单
	if err == nil{
		sqlstr = " update tbl_change_detail set  f_flag = 1, f_delete_user = "+reqs["f_delete_user"].(string)+",f_delete_time = now()  WHERE f_mid ="+reqs["id"].(string)
		_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr,)
	}

	if err !=nil{
		rerr = err
		rerr = tx.Rollback()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产变更","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除失败",err)
	} else{

		rerr = nil
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产变更","删除",reqs["id"].(string),reqs["f_delete_user"].(string),"删除成功 单号 ="+reqs["f_serno"].(string),err)
	}
	return
}
//资产变动 确认调整
func (ctl *AssetsApplyModel) AssetsChangeSubmit(req beego.Controller)(rerr error){
	var reqs map[string]interface{};
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var conn = services.NewConnection(utils.GetAppid(req))
	var sqlup string
	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容

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
	//检查单据状态
	err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_change_main","f_status",reqs["id"].(string),"1")
	if err != nil {
		fmt.Println(err)
		panic("单据不在申请状态不能调整！")
	}


	tx,_ := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()

	sqlstr := "update tbl_change_main set  f_audit = 1,f_audit_date = now(), f_audit_userid = "+reqs["userid"].(string)+" where id ="+reqs["id"].(string)

	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlstr)
	// 修改成功后 资产位置要更改到领用部门
	if (err == nil)  {
		var aparam = []string{"5", reqs["id"].(string),reqs["userid"].(string),utils.GetAppid(req)}
		var rdata map[string]interface{}

		if err ==nil{
			rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
			if (err !=nil) || (rdata["f_state"] !="1") {
				err = fmt.Errorf("%s","库存变动失败！")
			}
		}

		//修改资产列表对应属性
		if err ==nil{
			sqlup = " UPDATE  tbl_spinfo ts "+
				     " INNER JOIN tbl_change_detail tcd ON ts.id = tcd.f_spid "+
				     " INNER JOIN tbl_change_main tcm ON tcd.f_mid = tcm.id "+
				     " SET ts.f_Name = tcd.f_to_name "+
				     " WHERE tcm.id = "+reqs["id"].(string)+" AND tcm.f_audit = 1 and tcd.f_to_name <> '';"
			_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlup)
			if err !=nil{
				panic("资产名称修改失败，"+err.Error())
			}
			sqlup = " UPDATE  tbl_spinfo ts "+
				" INNER JOIN tbl_change_detail tcd ON ts.id = tcd.f_spid "+
				" INNER JOIN tbl_change_main tcm ON tcd.f_mid = tcm.id "+
				" SET  ts.f_Typeid = tcd.f_To_lbid "+
				" WHERE tcm.id = "+reqs["id"].(string)+" AND tcm.f_audit = 1 and tcd.f_To_lbid > -1;"
			_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlup)
			if err !=nil{
				panic("资产类别修改失败，"+err.Error())
			}

		}
	}

	if err !=nil{
		rerr = err
		rerr = tx.Rollback()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产变更","确认调整",reqs["id"].(string),reqs["userid"].(string),"调整失败",err)
	} else{

		rerr = nil
		rerr = tx.Commit()
		_= services.DB_Log(utils.GetAppid(req),conn,"资产变更","确认调整",reqs["id"].(string),reqs["userid"].(string),"调整成功 单号 ="+reqs["f_serno"].(string),err)
	}
	return
}
//资产变动 --辙消提交，不能辙消调整
func (ctl *AssetsApplyModel) AssetsChangeSubmit_Back(req beego.Controller)(rerr error){
	ctl.Base_Model.Add_BeginLog(req)  //调用基类方法
	var reqs map[string]interface{};
	var conn = services.NewConnection(utils.GetAppid(req)) ////通过客户信息获取数据库连接
	var newstate string
	var rkstate  string
	var billMemo string

	defer func() {
		_= conn.Close() //释放连接，要不然MYSQL连接会疯涨
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			switch err.(type){
			case error:rerr = err.(error)  //fmt.Errorf("%s",err)
			default:
				rerr = fmt.Errorf("%s",err)
			}
			return
		}
	}()
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &reqs);
	if err !=nil{
		panic("参数转换出错："+err.Error())
	}

	stype := "0"  //    0辙消申请

	//状态检测
	if stype == "0"{  //辙消提交
		err = services.Check_BillState(utils.GetAppid(req),conn,"tbl_change_main","f_status",reqs["id"].(string),"1")
		if err != nil {
			fmt.Println(err)
			panic("单据不在已提交申请状态不能辙消！")
		}
	}

	switch stype{  // 0辙消提交
	case "0":{
		newstate ="0"
		rkstate = "0"
		billMemo = "辙消提交"}
	default:
		newstate =""
	}

	tx,err := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			_=tx.Rollback()
			panic(err)
		}
	}()
	_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,"update tbl_change_main set f_status = ?,f_audit =? where id =? ",newstate,rkstate,reqs["id"].(string))

	//// 修改成功后 资产位置要更改到领用部门
	//if (err == nil) {
	//	var aparam = []string{"5", reqs["id"].(string),reqs["userid"].(string),utils.GetAppid(req)}
	//	var rdata map[string]interface{}
	//
	//	if err ==nil{
	//		rdata,err = services.Exec_Procedure(utils.GetAppid(req),tx,conn,"msp_Postion_Chang",aparam)
	//		if (err !=nil) || (rdata["f_state"] !="1") {
	//			err = fmt.Errorf("%s","库存变动失败！")
	//		}
	//	}
	//
	//	//修改资产列表对应属性
	//	if err ==nil{
	//		sqlup := "UPDATE  tbl_spinfo ts "+
	//			" INNER JOIN tbl_change_detail tcd ON ts.id = tcd.f_spid "+
	//			" INNER JOIN tbl_change_main tcm ON tcd.f_mid = tcm.id "+
	//			" SET ts.f_Name = tcd.f_to_name,ts.f_Typeid = tcd.f_To_lbid "+
	//			" WHERE tcm.id = "+reqs["id"].(string)+" AND tcm.f_audit = 1;"
	//		_,err = services.Exec_Sql(utils.GetAppid(req),tx,conn,sqlup)
	//		if err !=nil{
	//			panic("资产属性更新失败，"+err.Error())
	//		}
	//	}
	//
	//}

	if err == nil{
		err = services.Bill_AddStream(utils.GetAppid(req),tx,conn,"资产变更",reqs["userid"].(string),reqs["id"].(string),billMemo,reqs["memo"].(string))
	}
	if err == nil{
		err= services.DB_Log(utils.GetAppid(req),conn,"资产变更",billMemo,"",reqs["userid"].(string),"辙消申请成功",err)
	}

	if err ==nil{
		_ = tx.Commit()
	}else {
		_ = tx.Rollback();
		rerr = err
		if utils.ISLog{
			utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI,"操作失败，"+rerr.Error())
		}
	}

	return rerr
}
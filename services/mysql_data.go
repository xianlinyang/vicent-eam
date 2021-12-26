package services
// 数据库操作底层，所有数据库操作统一调用本单元函数进行
import (
	"Eam_Server/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"log"

	"strconv"
	"strings"
	"time"

)


func  init()  {

}

type Check_Repeat struct {
	FName    string   //字段名称
	PName    string   //提示中文名称
}

type Check_Type struct {
	FName    string
	Ptype    string
	Chcktype int
	Fvalue   string
}
//创建连接
func NewConnection(appid string) *sql.DB {
	var acustom = utils.CustomConfig{}
	var db_host string
	var dbName string
	var dbUser string
	var dbPwd string
	var dbType string

	acustom = utils.GetDBName(appid)
	fmt.Println(acustom.DbName)//使用数据库名，用于动态配置

	//"gopkg.in/ini.v1"
	//str, _ := os.Getwd()  //获取程序路径
	//b_con,err := utils.PathExists(str) //判断文件是否存在
	//if  (err== nil) &&(b_con){  //如果存在配置文件则取配置文件的数据库配置
	//	cfg, err_wj := ini.Load(str+`\DBConfig.ini`)
	//	if err_wj ==nil{
	//		dbType = "mysql"
	//		db_host = cfg.Section("mysql").Key("db_host").String()
	//		dbName = cfg.Section("mysql").Key("db_name").String()
	//		dbUser = cfg.Section("mysql").Key("db_user").String()
	//		dbPwd = cfg.Section("mysql").Key("db_pwd").String()
	//	}
	//}else{
		dbType = beego.AppConfig.String("db_type")
		db_host = beego.AppConfig.String(utils.StringsJoin(dbType, "::db_host"))
		dbName = beego.AppConfig.String(utils.StringsJoin(dbType, "::db_name")) // acustom.DbName//
		dbUser = beego.AppConfig.String(utils.StringsJoin(dbType, "::db_user"))
		dbPwd = beego.AppConfig.String(utils.StringsJoin(dbType, "::db_pwd"))
	//}


	dbconstr := dbUser+":"+dbPwd+"@("+db_host+")/"+dbName+"?charset=utf8"
	fmt.Println(dbconstr)

	db, err := sql.Open("mysql",  dbconstr)//"root:1234@(localhost:3306)/EsaleEam_DB?charset=utf8"
	if err != nil {
		log.Fatal(err)
		fmt.Println("数据库连接出错")
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	//db.SetMaxOpenConns(4) 设置mysql最大连接数
	//db.SetMaxIdleConns(10) 设置mysql最大闲置连接数
	return db

	//return GetConnecttion(req.GetString(utils.Custom_SvrPort,""))
}

func GetConnecttion(asvrport string) *sql.DB{
	var acustom = utils.CustomConfig{}
	acustom = utils.GetDBName(asvrport)

	dbType := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String(utils.StringsJoin(dbType, "::db_host"))
	dbName := acustom.DbName//beego.AppConfig.String(utils.StringsJoin(dbType, "::db_name"))
	dbUser := beego.AppConfig.String(utils.StringsJoin(dbType, "::db_user"))
	dbPwd := beego.AppConfig.String(utils.StringsJoin(dbType, "::db_pwd"))
	dbconstr := dbUser+":"+dbPwd+"@("+db_host+")/"+dbName+"?charset=utf8"
	//fmt.Println(dbconstr)

	db, err := sql.Open("mysql",  dbconstr)//"root:1234@(localhost:3306)/EsaleEam_DB?charset=utf8"
	if err != nil {
		log.Fatal(err)
		fmt.Println("数据库连接出错")
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	//db.SetMaxOpenConns(4) 设置mysql最大连接数
	//db.SetMaxIdleConns(10) 设置mysql最大闲置连接数
	return db
}

func Exec_Sql(appid string,aconn1 *sql.Tx,aconn *sql.DB,query string, args ...interface{}) (n int64, aerr error){
	var v string
	//准备打日志
	asvrport := appid


	if len(args) > 0{
		v = "exec sql "+utils.Get_LogSql(query,args...)
	} else {
		v =  "exec sql "+query
	}
	if utils.SqlLog {
		utils.LogOut("info","svrport = "+asvrport+" exec sql ",v)
	}
	if utils.ISLog{
		fmt.Println(v)
	}
	if aconn1 !=nil{  //有事务连接优生使用事务连接，保证事务统一
		_,aerr = aconn1.Exec(query,args...)
		utils.LogOut("error","svrport = "+asvrport+" exec sql ",v)
	}else{
		_,aerr = aconn.Exec(query,args...)
	}

	return 0,aerr
}
//查询多条数据，返回JSON格式
func  FindToList(appid string,aconn *sql.DB,asql string, args ...interface{}) (result []map[string]interface{}, rerr error) {
	var v string
	//准备打日志
	asvrport := appid
	if len(args) > 0{
		v = "exec sql "+utils.Get_LogSql(asql,args...)
	} else {
		v =  "exec sql "+asql
	}
	//fmt.Println(v)
	if utils.SqlLog {
		fmt.Println(asql)
		utils.LogOut("info","svrport = "+asvrport+" exec sql ",v)
	}
	rows, err := aconn.Query(asql,args ...);

	if err == nil{

		columns,err := rows.Columns()
		if err != nil {
			panic(err.Error())
		}
		count := len(columns)
		tableData := make([]map[string]interface{}, 0)  // yxl map key-value结构  这里的key 是string value是interface
		values := make([]interface{}, count)
		valuePtrs := make([]interface{}, count)

		for rows.Next() {
			for i := 0; i < count; i++ {
				valuePtrs[i] = &values[i]
			}
			_ = rows.Scan(valuePtrs...)
			entry := make(map[string]interface{})
			for i, col := range columns {
				var v interface{}
				val := values[i]  // yxl values为数据值
				b, ok := val.([]byte)
				//fmt.Println(string(b))
				if ok {
					v = string(b)
				} else {
					v = val
				}
				entry[col] = v
				//fmt.Println(entry)
			}
			tableData = append(tableData, entry)
		}
		_ = rows.Close()
		//utils.LogOut("info","svrport = "+asvrport+" query sql ",v)
		return tableData,nil

	}else {
		fmt.Println("查询失败",err)
		utils.LogOut("error","svrport = "+asvrport+" query sql ",v)
		return nil,err
	}
}

//带事务-查询多条数据，返回JSON格式
func  FindToList_tax(appid string,aconn *sql.Tx,asql string, args ...interface{}) (result []map[string]interface{}, rerr error) {
	var v string
	//准备打日志
	asvrport := appid
	if len(args) > 0{
		v = "exec sql "+utils.Get_LogSql(asql,args...)
	} else {
		v =  "exec sql "+asql
	}
	//fmt.Println(v)
	if utils.SqlLog {
		fmt.Println(asql)
		utils.LogOut("info","svrport = "+asvrport+" exec sql ",v)
	}
	rows, err := aconn.Query(asql,args ...);

	if err == nil{

		columns,err := rows.Columns()
		if err != nil {
			panic(err.Error())
		}
		count := len(columns)
		tableData := make([]map[string]interface{}, 0)  // yxl map key-value结构  这里的key 是string value是interface
		values := make([]interface{}, count)
		valuePtrs := make([]interface{}, count)

		for rows.Next() {
			for i := 0; i < count; i++ {
				valuePtrs[i] = &values[i]
			}
			_ = rows.Scan(valuePtrs...)
			entry := make(map[string]interface{})
			for i, col := range columns {
				var v interface{}
				val := values[i]  // yxl values为数据值
				b, ok := val.([]byte)
				//fmt.Println(string(b))
				if ok {
					v = string(b)
				} else {
					v = val
				}
				entry[col] = v
				//fmt.Println(entry)
			}
			tableData = append(tableData, entry)
		}
		_ = rows.Close()
		//utils.LogOut("info","svrport = "+asvrport+" query sql ",v)
		return tableData,nil

	}else {
		fmt.Println("查询失败",err)
		utils.LogOut("error","svrport = "+asvrport+" query sql ",v)
		return nil,err
	}
}
//按分页查询数据，SQL中必须要有ID字段
func FindToListByPage(appid string,aconn *sql.DB,asql string,aPageIndex int,aPageSize int) (result utils.ResData,rerr error){
	//fmt.Println("进入分页")
	var err error
	var rdata = utils.ResData{};

	var apage_bid string
	var limrecode string

	apage_bid = strconv.Itoa((aPageIndex-1)*aPageSize)
	limrecode = strconv.Itoa(aPageSize )
	//var abid string
	////var aeid int64
	//rdata.CurPage = aPageIndex
	//rdata.PageSize = aPageSize
	////得到开始ID
	//err := F_Connection_db.QueryRow("SELECT id from ("+asql+") a LIMIT "+apage_bid+",1").Scan(&abid)
	//if err != nil{
	//	fmt.Println(err.Error())
	//	return  rdata,err
	//}
	//
	//////得到结束ID
	////err = F_Connection_db.QueryRow("SELECT id from ("+asql+") a LIMIT "+apage_eid+",1").Scan(&aeid)
	////if err != nil{
	////	return  rdata,err
	////}
	////fmt.Println(aeid)
	////得到合计条数
	//err = F_Connection_db.QueryRow("SELECT count(*) as id from ("+asql+") a ").Scan(&rdata.Totals)
	//if err != nil{
	//	return  rdata,err
	//}
	////fmt.Println(abid)
	//fmt.Println(apage_bid)
	////得到数查询分页数据
	//fmt.Println("SELECT id from ("+asql+") a  where a.id >  "+abid +" limit 0,"+string(limrecode))
	////result.Data
	//rdata.Data,err = FindToList("SELECT * from ("+asql+") a  where a.id >  "+abid +" limit 0,"+limrecode)

	rdata.Result,err = FindToList(appid,aconn,"SELECT * from ("+asql+") a   limit "+apage_bid+","+limrecode)
	if err !=nil{
		fmt.Println("FindToListByPage 出错,",err)
		return  rdata,err
	}
    //s,err := json.Marshal(result)
    //if err ==nil{
    //	fmt.Println(string(s))
	//}
	//fmt.Println(json.Marshal(result.Data))
	return rdata,nil
}


//批量添加到数据表
func Batch_ADD_Table(appid string,aconn1 *sql.Tx,aconn *sql.DB,atable string,afile []string,adata []map[string]interface{})(result,autoid string,err error){
	var sinsert =" insert into "+atable;
	var	sfile = "";
	//找出需要插入的字段
	for i := 0;i<len(afile);i++{
		if sfile ==""{
			sfile = afile[i]
		}else {
			sfile = sfile +","+afile[i]
		}
	}
	sinsert =sinsert+"("+sfile+") values "

	var stm =""
	//设置(?,?,?)
	for i:=0;i<len(afile);i++{
		if stm ==""{
			stm ="?"
		} else {
			stm =stm+",?"
		}
	}
	stm ="("+stm+")"
	valueStrings := make([]string, 0, len(adata))
	valueArgs := make([]interface{}, 0, len(adata) * len(afile))

	//设置(?,?,?)
	for i :=0;i<len(adata);i++{
		valueStrings = append(valueStrings,stm)
	}
	//设置VALUE值
	for i := 0;i<len(adata);i++{
		for j:=0;j<len(afile);j++{
			valueArgs = append(valueArgs,adata[i][afile[j]])
		}
	}

	smt :=sinsert+strings.Join(valueStrings,",")

	//fmt.Println(smt)
	//fmt.Println(valueArgs)


	if aconn1 !=nil{  //有事务连接优生使用事务连接，保证事务统一
		_, err = aconn1.Exec(smt,valueArgs...)
	} else{
		_, err = aconn.Exec(smt,valueArgs...)
	}

	if err ==nil{

		return "1","-1",nil
	} else {

		return "-1","-1",err
	}
}

//单条数据手插入
func One_ADD_Table(appid string,aconn1 *sql.Tx,aconn *sql.DB,atable string,adata map[string]interface{})(autoid string,aerr error){
	var sinsert =" insert into "+atable+" set ";
	var	sqlstr = "";
	fmt.Println(adata)
	filedlist := make(map[string]string)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误 One_ADD_Table ：", err) // 这里的err其实就是panic传入的内容
			autoid ="-1"
			aerr = fmt.Errorf("%s",err)  //Errorf 根据于格式说明符进行格式化，并将字符串作为满足 error 的值返回，其返回类型是error．
			return
		}
	}()

	//赋值需要插入的字段值
	for key,value := range  adata {
		filedlist[key] = value.(string)
	}

	sqlstr =  sinsert+Sql_JionSet(filedlist);
	if utils.ISLog{
		fmt.Println(sqlstr)
	}

	if aconn1 != nil{  //有事务连接优生使用事务连接，保证事务统一
		_, aerr = aconn1.Exec(sqlstr)
	} else {
		_, aerr = aconn.Exec(sqlstr)
	}

	asvrport := appid
	v := "exec sql "+sqlstr

	if utils.SqlLog {
		utils.LogOut("info","svrport = "+asvrport+" exec sql ",v)
	}

	if aerr !=nil{
		panic("插入失败："+aerr.Error())
	}


	var aid int
	if aconn1 != nil{ //有事务连接优生使用事务连接，保证事务统一
		aerr = aconn1.QueryRow("SELECT LAST_INSERT_ID();").Scan(&aid)

	} else {
		aerr = aconn.QueryRow("SELECT LAST_INSERT_ID();").Scan(&aid)

	}

	if (aerr == nil) {
		autoid = strconv.Itoa(aid)
		aerr = nil
 		fmt.Println(autoid)
	}else {
		utils.LogOut("error","svrport = "+asvrport+" exec sql ",v)
	}
	return autoid,aerr
}

//得到一条SQL某个字段值
func Get_FieldByValue(appid string,aconn *sql.DB,asql,afilename string)(avalue string,aerr error){
	var e_sql string
	e_sql = "select "+afilename+" from ("+asql+") a LIMIT 0,1"
	aerr = nil
	if utils.ISLog{
		fmt.Println(e_sql)
	}

	asvrport := appid
	v := "query sql "+e_sql

	if utils.SqlLog {
		utils.LogOut("info","svrport = "+asvrport+" exec sql ",v)
	}
	//fmt.Println(v)
	err := aconn.QueryRow(e_sql).Scan(&avalue)
	if err != nil{
		if err == sql.ErrNoRows{  //空数据返回
			avalue = ""
			fmt.Println("没数据"+err.Error())
		} else {
			avalue = ""
			aerr = err
		}
		utils.LogOut("error","svrport = "+asvrport+" query sql ",v)
	}
	return
}

//得到一行数据
func Get_OneData(appid string,aconn *sql.DB,asql string )(rdata map[string]interface{},aerr error){
	var dataarry []map[string]interface{}
	var e_sql string

	e_sql = " SELECT *FROM ("+asql+")a  limit 0,1"

	aerr = nil


	dataarry,aerr = FindToList(appid,aconn,e_sql)
	if len(dataarry) > 0{
		rdata = dataarry[0]
	} else {
		rdata = nil
	}

	//as,_ := json.Marshal(rdata)
	//
    //s := string(as)
 	//fmt.Println(s)
	return
}

//得到存储过程返回第一行数据一个字段值
func Get_Procedure_Data(appid string,aconn *sql.DB,asql,afilename string)(avalue string,aerr error){
	var e_sql string
	e_sql = asql
	aerr = nil
	if utils.ISLog{
		fmt.Println(e_sql)
	}

	asvrport := appid
	v := "query sql "+e_sql

	if utils.SqlLog {
		utils.LogOut("info","svrport = "+asvrport+" exec sql ",v)
	}
	//fmt.Println(v)
	err := aconn.QueryRow(e_sql).Scan(&avalue)
	if err != nil{
		if err == sql.ErrNoRows{  //空数据返回
			avalue = ""
			fmt.Println("没数据"+err.Error())
		} else {
			avalue = ""
			aerr = err
		}
		utils.LogOut("error","svrport = "+asvrport+" query sql ",v)
	}
	return
}
//系统
func DB_Log(appid string,aconn *sql.DB,amodal,atype,akeyword,auserid,slog string,aerr error)error{
	var sql_log = "  insert into sys_log(f_customid,f_typelevel,f_modal,f_keywords,f_oprtime,f_userid,f_note,f_oprtype) value(?,?,?,?,?,?,?,?) "
	var err error
	svrdate := time.Now().Format("2006-01-02 15:04:05")
	if aerr == nil{
		_,err = aconn.Exec(sql_log,appid,"0",amodal,akeyword,svrdate,auserid,slog,atype)
	} else {
		serr := aerr.Error()
		if len(serr) > 300{
			slog = slog+serr[:300] //取前300位字符
		} else {
			slog = slog+serr
		}


		_,err = aconn.Exec(sql_log,appid,"0",amodal,akeyword,svrdate,auserid,slog,atype)
	}

	return err
}

//检测数据是否存在表中  achk --要检测的重复字段  awhere--检测的副加条件
func CheckRepeat(appid string,aconn *sql.DB,adata map[string]interface{},achk []byte,atabl string,awhere []byte)(result int,serror string){
	var sql string
	var swhere string
	var achkmodel = []Check_Repeat{}
	var achktype =[]Check_Type{}

	defer func(){
		if err := recover(); err != nil {
			result = -1
			serror = err.(string)
		}
	}()
    fmt.Println("进入检测")
	//jsonstr := `[{"FName":"f_no","PName":"编号"},{"FName":"f_name","PName":"名称"}]`
	result = 1
	err := json.Unmarshal(achk,&achkmodel)
	if err != nil{
		panic("数据检测出错，"+err.Error())
	}
	if awhere != nil{
		err = json.Unmarshal(awhere,&achktype)
		if err != nil{
			panic("数据检测出错，"+err.Error())
		}
	}


	swhere =" "

	for _,v := range achkmodel{
		if result != 1{
			break
		}
		swhere =" where f_customid ="+appid+" "
		if adata[v.FName] != nil{
			//for I := 0; I < len(awhere); I++ {
			if len(achktype) > 0{
				for _,v1 := range achktype{
					if v1.Chcktype == 0{  //取用结构值
						if swhere ==""{
							swhere +=  v1.FName+v1.Ptype+adata[v1.FName].(string)  //{"FName":"id","Ptype":"<>","Chcktype":0,"Fvalue":""}
						}else {
							swhere += " and "+v1.FName+v1.Ptype+adata[v1.FName].(string)
						}
					} else {  //取用传入值
						if swhere ==""{
							swhere +=  v1.FName+v1.Ptype+v1.Fvalue
						}else {
							swhere += " and "+v1.FName+v1.Ptype+v1.Fvalue
						}
					}

				}
			}
			//}
			swhere =swhere+" and "+v.FName+" =\""+adata[v.FName].(string)+"\""
			//if swhere !=""{
			//
			//} else {
			//	swhere =" where "+v.FName+" =\""+adata[v.FName].(string)+"\""
			//}

			sql = "select "+v.FName+" from "+atabl+swhere
			if utils.ISLog{
				fmt.Println(sql)
			}

			svalue,err := Get_FieldByValue(appid,aconn,sql,v.FName)
			if err == nil{
				if svalue !=""{
					result = -1
					serror =v.PName+"【"+adata[v.FName].(string)+"】系统已存在，不能重复添加！"
					break
				} else {
					result = 1
					serror =""
				}
			} else {
				result = -1
				serror = "数据检测失败，"+err.Error()
				break
			}
		}
	}

	return
}

//检查单据状态
func Check_BillState(appid string,aconn *sql.DB,atable string,afilestatename string,aid string,astatevalue string)(serror error){
	var e_sql string
	var avalue string
	e_sql = "select "+afilestatename+" from "+atable+" where id = "+aid+"   LIMIT 0,1"

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出现错误：", err) // 这里的err其实就是panic传入的内容
			aid = ""
			serror = fmt.Errorf("%s",err)

			return
		}
	}()

	if utils.ISLog{
		fmt.Println(e_sql)
	}

	asvrport := appid
	v := "query sql "+e_sql

	if utils.SqlLog {
		utils.LogOut("info","svrport = "+asvrport+" exec sql ",v)
	}

	err := aconn.QueryRow(e_sql).Scan(&avalue)
	if err != nil{
		if err == sql.ErrNoRows{  //空数据返回
			panic("单据不存在")
		} else {
			panic("单据状态检测出错"+err.Error())
		}
	}

	if avalue != astatevalue{
		panic("单据状态与服务器不匹配！")
	}
	return nil
}

//插入单据流水表
  // atype 0资产领用申请 1资产领用
func Bill_AddStream(appid string,aconn1 *sql.Tx,aconn *sql.DB,atypename string,auserid string,aid string,aopr string,amemo string)(err error){
	sql := " INSERT INTO tbl_billStream(f_customid,f_type,f_opruserid,f_oprtime,f_opr,f_memo,f_busid) "+
	       " VALUES(?,?,?,now(),?,?,?) "


	if aconn1 != nil{
		_,err = aconn1.Exec(sql,appid,utils.GetBus_TypeID(atypename),auserid,aopr,amemo,aid)
	} else {
		_,err = aconn.Exec(sql,appid,utils.GetBus_TypeID(atypename),auserid,aopr,amemo,aid)
	}
	return err
}

//执行业务存储过程
func Exec_Procedure(appid string,aconn1 *sql.Tx,aconn *sql.DB,PName string,adata []string) (rdata map[string]interface{},err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("存储过程执行出错：", err) // 这里的err其实就是panic传入的内容
			switch err.(type){
			case error:err = err.(error)  //fmt.Errorf("%s",err)
			default:
				err = fmt.Errorf("%s",err)
			}
			return
		}
	}()
	var callstr string
	//var handle *sql.Stmt
	if (len(adata) == 0) {
		callstr = "call "+PName
	} else {
		callstr = "call "+PName+" ( "
		for i,v := range adata{
			if (i ==0){
				callstr = callstr+"\""+v+"\""
			}else {
				callstr = callstr+",\""+v+"\""
			}

		}
		callstr =callstr+" )"
	}

	//fmt.Println("进入存储过程调用")
	//if utils.SqlLog{
	//	fmt.Println(callstr)
	//}
	var rdatalist []map[string]interface{}
	rdatalist,err = FindToList_tax(appid,aconn1,callstr)

	if err ==nil{
		if (len(rdatalist)>0){
			rdata = rdatalist[0]

		}
	}

	//if (aconn1 !=nil){
	//	//handle, err := db.Prepare("CALL dsp_settle.settle_balance_deduction(?, ?, ?, @out_status)")
	//	handle, err = aconn1.Prepare(callstr)
	//} else {
	//	handle, err = aconn.Prepare(callstr)
	//}
	//
	//if err != nil {
	//	panic(err.Error())
	//}
	//defer handle.Close()

	//var sql string = "SELECT @out_status as ret_status"
	//selectInstance, err := db.Prepare(sql)
	//if err != nil {
	//	panic(err.Error())
	//}
	//defer selectInstance.Close()
	//
	//var ret_status int
	//err = selectInstance.QueryRow().Scan(&ret_status)
	//if err != nil {
	//	panic(err.Error())
	//}
	//fmt.Println(ret_status)
	//return ret_status

	//var result sql.Result
	////result, err = handle.Exec(userID, planID, charge)  // 如果 callstr 中带有 ? 参数，这里可以增加参数值
	//result,err = handle.Exec()
	//if (err !=nil){
	//	panic(err)
	//}
	//
	//fmt.Println("存储过程调用返回", result )
	return rdata,err


}
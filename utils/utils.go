package utils
//公共方法，公共结构单元
import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

var runmode string
// fileLogs 生产环境下日志
var fileLogs *logs.BeeLogger
// consoleLogs开发模式下日志
var consoleLogs *logs.BeeLogger
var svrdate string

func init() {
	//创建日志
	Logs_Create()
	runmode = strings.TrimSpace(strings.ToLower(beego.AppConfig.String("runmode")))
	if runmode == "" {
		runmode = "dev"
	}
}

func Logs_Create(){
	consoleLogs = logs.NewLogger(1)
	_=consoleLogs.SetLogger(logs.AdapterConsole)
	fileLogs = logs.NewLogger(10000)
	//fileLogs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/goERP.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"],"level":7,"daily":true,"maxdays":10}`)
	now := time.Now()
	svrdate = now.Format("2006-01-02");
	_=fileLogs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/`+svrdate+`.log","separate":["info", "error"],"level":7,"daily":true,"maxdays":30}`)
}
// StringsJoin 字符串拼接
func StringsJoin(strs ...string) string {
	var str string
	var b bytes.Buffer
	strsLen := len(strs)
	if strsLen == 0 {
		return str
	}
	for i := 0; i < strsLen; i++ {
		b.WriteString(strs[i])
	}
	str = b.String()
	return str

}

// interface 转 int
func InterToInt(a interface{}) (ra int){
	var rint int
	//fmt.Println("进入转换")
	//fmt.Println("要转值",a)
	switch  a.(type) {
		case string:{
			str :=strings.Replace(a.(string),"\"","",-1)
			astr := str
			//fmt.Println(astr)
			aint64,err := strconv.Atoi(astr)
			if err == nil{
				rint = aint64
			} else {
				rint = 0
				fmt.Println(err)
			}
		}
		case int:{
			rint = a.(int)
		}
		case int64:{
			aint64 := a.(int64)
			rint = int(aint64)
		}
		case float32:{
			aint64 := int(a.(float32))
			rint = aint64
		}
		case float64:{
			aint64 := int(a.(float64))
			rint = aint64
		}
	}
	return rint
}

func GetParamStr(aparstr string)string{
	var str string
	str =strings.Replace(aparstr,"\"","",-1)
	str =strings.Replace(str,"”","",-1)
	return str
}
func Get_LogSql(asql string,args ...interface{})string{
	var str string
	str = asql
	for _,v := range  args {
		str = strings.Replace(str, "?", fmt.Sprintf("%v", v), 1) //字符串替换就可以了
	}
	return str
}

//BASE64转文件
func Base64ToFile(path, fileName, base64Str string) (err error) {

	err = os.MkdirAll(path, 0777) //创建目录
	if err != nil {
		return err
	}

	ddd, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path+`\`+fileName, ddd, 0666) //buffer输出文件中（不做处理，直接写到文件）
	if err != nil {
		return err
	}
	return nil
}

//文件转成 base64
func FileToBase64(apath string)(filstr string,err error) {
	data, err := ioutil.ReadFile(apath)
	if err != nil {
		panic(err)
	}
	base64Str := base64.StdEncoding.EncodeToString(data)

	return base64Str,err

}
//添加日志
func Set_LocalLog(req beego.Controller,slog string,err error){
	asvrport := req.GetString(Custom_SvrPort,"")
	if err == nil{
		LogOut("info","svrport = "+asvrport+" ",slog)
	} else{
		LogOut("error","svrport = "+asvrport+" ",slog+" error:"+err.Error())
	}
}

//LogOut 输出日志
// @Title LogOut
// @Param	body		body 	models.AccountAccountTag	true		"body for AccountAccountTag content"
func LogOut(level, title, v interface{}) {
	format := "%s"
	if level == "" {
		level = "debug"
	}
	now := time.Now()
	if now.Format("2006-01-02") != svrdate {
		Logs_Create() //跨天后重新创建日志文件
	}

	if runmode == "dev" {
		switch level {
		//case "emergency":
		//	fileLogs.Emergency(format, title, v)
		//case "alert":
		//	fileLogs.Alert(format, title, v)
		//case "critical":
		//	fileLogs.Critical(format, title, v)
		case "error":
			fileLogs.Info(format, title,v)  // 出错也要打普通日志，普通日志为所有日志方便以后查流水
			fileLogs.Error(format, title, v)
		//case "warning":
		//	fileLogs.Warning(format, title, v)
		//case "notice":
		//	fileLogs.Notice(format, title, v)
		//case "informational":
		//	fileLogs.Informational(format, title, v)
		//case "debug":
		//	fileLogs.Debug(format, title, v)
		//case "warn":
		//	fileLogs.Warn(format, title, v)
		case "info":
			fileLogs.Info(format, title, v)
		//case "trace":
		//	fileLogs.Trace(format, title, v)
		default:
			fileLogs.Debug(format, title, v)
		}
	} else {
		switch level {
		case "emergency":
			fileLogs.Emergency(format, title, v)
		case "alert":
			fileLogs.Alert(format, title, v)
		case "critical":
			fileLogs.Critical(format, title, v)
		case "error":
			fileLogs.Error(format, title, v)
		case "warning":
			fileLogs.Warning(format, title, v)
		case "notice":
			fileLogs.Notice(format, title, v)
		case "informational":
			fileLogs.Informational(format, title, v)
		case "debug":
			fileLogs.Debug(format, title, v)
		case "warn":
			fileLogs.Warn(format, title, v)
		case "info":
			fileLogs.Info(format, title, v)
			//consoleLogs.Info(format,title, v)
		case "trace":
			fileLogs.Trace(format, title, v)
		default:
			fileLogs.Debug(format, title, v)
		}
	}
}

//获取各业务对应的ID
func GetBus_TypeID(BusName string)(sid string){
	bus_arr:=map[string]string{
		"资产申请":"1",
		"资产领用":"2",
		"资产退库":"3",
	}

	s := bus_arr[BusName]
	if s ==""{
		sid = "-1"
	}else {
		sid = s
	}
	return sid
}



//func main() {
//	path := "d:/test.txt"
//	b, err := PathExists(path)
//	if err != nil {
//		fmt.Printf("PathExists(%s),err(%v)\n", path, err)
//	}
//	if b {
//		fmt.Printf("path %s 存在\n", path)
//	} else {
//		fmt.Printf("path %s 不存在\n", path)
//		err := os.Mkdir(path, os.ModePerm)
//		if err != nil {
//			fmt.Printf("mkdir failed![%v]\n", err)
//		} else {
//			fmt.Printf("mkdir success!\n")
//		}
//	}
//}

/*
   判断文件或文件夹是否存在
   如果返回的错误为nil,说明文件或文件夹存在
   如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
   如果返回的错误为其它类型,则不确定是否在存在
*/
func PathExists(path string) (bool, error) {

	_, err := os.Stat(path)  // 文件不存在会返回 <> nil
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) { // 判断错误类型是否是文件不存在
		return false, nil
	}
	return false, err
}
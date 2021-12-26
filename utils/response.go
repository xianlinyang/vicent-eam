package utils

//接口返回参数单元

//SuccessCode 请求成功
const SuccessCode = 1
//FailedCode 请求失败
const FailedCode = -1

//SuccessMsg 成功提示信息
const SuccessMsg = "操作成功"
//FailedMsg 失败提示信息
const FailedMsg = "操作失败"
//是否打操作日志
const ISLog = true
const SqlLog = true

//获取客户唯一ID 字估
const Custom_SvrPort = "customid"
const TokenSecret ="esale20200819-Eam_Server"

const Code_ChkValid = 20301  //参数有效性检测不过
const Code_Chk_TokenValid = 20201 //TOKEN检测不过
const Code_Chk_Permis  = 20001
const Code_User_Chk  = 20002
const SvrPicPath =   `http://192.168.1.42:8080/images`
//const SvrPicPath =   `http://114.116.74.95:8090/images`
////返回字段名
//const Response_Code = "msgid"
//const Response_Message ="msgstr"
//const Response_Data ="data"

//其他系统全局变量
const Data_null = "<NULL>"

type ResData struct {
	CurPage    int
	PageSize   int
	Totals     int
	Footer     map[string] interface{}
	Sort       map[string]interface{}
	Result       []map[string]interface{}
}

type ResponsModel struct {
	Code    	int
	Success     bool
	Message     string
	Content  	ResData
	Token       string
}


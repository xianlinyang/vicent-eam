package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"runtime"
	"Eam_Server/utils"
)

type Base_Model struct {

}


//日志方法
func (ctl *Base_Model) Add_BeginLog(req beego.Controller) {
	fmt.Println("开始输出")
	pc := make([]uintptr,1)
	runtime.Callers(2,pc)
	f := runtime.FuncForPC(pc[0])
	fmt.Println(f.Name())
	//Request.RequestURI
	//" "+req.Ctx.Request.URL
	if utils.ISLog{
		utils.LogOut("info","svrport = "+utils.GetAppid(req)+" "+req.Ctx.Request.RequestURI+"传入参数：",string(req.Ctx.Input.RequestBody))
	}
}
////所有修改基类方法
//func (ctl *Base_Model) All_Update(reqbody []byte) {
//
//}
////所有删除基类方法
//func (ctl *Base_Model) All_Delete(reqbody []byte) {
//
//}
////所有查询基类方法
//func (ctl *Base_Model) All_GetList(reqbody []byte) {
//
//}
////所有查询基类方法
//func (ctl *Base_Model) All_GetRow(reqbody []byte) {
//
//}
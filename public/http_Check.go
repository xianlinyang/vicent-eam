package public
//  HTTP 防问中间件，用于检测防问参数的有效情和防问日志记录
import (
	"Eam_Server/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/url"
	"strings"
)


func init() {
	var filterLogin = func(ctx *context.Context) {
		var amodel = Check_API{};
		var v string
		//var url string
		var reqbody   map[string]interface{}
		appid := ctx.Input.Query(utils.Custom_SvrPort)
		if appid ==""{
			appid ="0"
		}
		//url = ctx.Request.RequestURI
		//if(strings.Contains(url,"?")) {
		//	url=_string.Substr(url,0,strings.Index(url,"?"))
		//};
		//url = strings.Replace(url,"/api/", "", -1)
		v =  "params = "+ctx.Request.RequestURI
		//escapeUrl := url.QueryEscape(v)
		enEscapeUrl, _ := url.QueryUnescape(v)

		//output,_ := iconv.ConvertString(v, "utf-8","gbk")
		fmt.Println(enEscapeUrl)
		_= json.Unmarshal(ctx.Input.RequestBody, &reqbody);
		obdys , _ := json.Marshal(reqbody)
		v=v+"; body = "+string(obdys)
		//fmt.Println(v)

		fmt.Println(ctx.Input.Method())
		if utils.ISLog{
			utils.LogOut("info","接口访问 appid="+appid+"; 接口 ="+ctx.Input.URL()+"; ",v)
		}

		var Rstr = utils.ResponsModel{}
		var result int
		var serr string
		Rstr.Code = 1
		apis := strings.Replace(ctx.Input.URL(), "/api/", "", -1)

		if (ctx.Input.Method() =="POST") ||(ctx.Input.Method() =="PUT") || (ctx.Input.Method() =="DELETE"){
			if utils.ISLog{
				//bodymap := make(map[string]interface{})
				//_=json.Unmarshal(ctx.Input.RequestBody,&bodymap)
				fmt.Println(ctx.Request.RequestURI+"传入参数："+string(ctx.Input.RequestBody))

				//fmt.Println(bodymap)
			}
			//Token检测
			stoken := ctx.Input.Header("authorization")
			fmt.Println("token = "+stoken)


			ischk := Api_Check(ctx.Input.URL())

			if ischk{
				result,serr = CheckToken(appid ,stoken)
			} else{
				result = 1
				serr =""
			}

			//权限检测
			if (result == 1) &&(ischk){
				//apis := strings.Replace(ctx.Input.URL(), "/api/", "", -1)
				result,serr = Check_Permissions(appid,apis,stoken)
			}

			if (result == 1) {
				result,serr = amodel.Chk_API_validate(ctx.Input.RequestBody,ctx.Request.RequestURI); //传入值有效性检测
			}

			if result != 1{
				Rstr.Success = false
				Rstr.Code = result
				Rstr.Message = serr
			}
		}

		if Rstr.Code != 1{
			rstr,err := json.Marshal(Rstr)
			if err ==nil{
				ctx.ResponseWriter.Header().Set("Content-type", "application/text; charset=utf-8")
				ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin","*")
				//ctx.ResponseWriter.Header().Set("ccess-Control-Allow-Origin", "application/text")
				//ctx.ResponseWriter.WriteHeader(403)
				_,_ =ctx.ResponseWriter.Write([]byte(string(rstr)));  //检测不通过跳出
				//ctx.ResponseWriter.Status =200

				//ctx.Request.Response.Body.Read([]byte(string(rstr)))

				if utils.ISLog{
					fmt.Println("检测不通过",Rstr)
				}

				return
			}
		}
	}
	beego.InsertFilter("/*",beego.BeforeRouter,filterLogin)

	// 因为maingo 加载了 _ "Eam_Server/routers" 所以会运行到这里来，加载一个包会执行这个包下所有go下面的的init方法，依此类推
}
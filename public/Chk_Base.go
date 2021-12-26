package public
//接口参数检测基类
import (
	"Eam_Server/utils"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

type Check_Base struct {

}

type Check_Model struct {
	Label    string
	PName    string
	Rules    []string
	conditions [] int  //大于 0位 小于1位
}

func (ctl *Check_Base) validate(res []byte,achk []Check_Model)(result int,serr string)  {
	var req map[string]interface{};
	var bol bool

	var slabel string
	var sinfo string
	var aleng int
	_ = json.Unmarshal(res,&req)
	result = 1
	serr = ""

	I:
	for _,v := range achk{
		slabel = v.Label
 		for _,role := range v.Rules{
 			if role == "notnull"{
				sinfo = "不能为空！"
 				bol = ctl.chknull(req[v.PName])
			} else if role =="length"{
				aleng = len(ctl.InterToStr(req[v.PName]))
				//bol =  (aleng >= v.conditions[0] && aleng<= v.conditions[1])
				bol = aleng != v.conditions[0]
				sinfo = "长度必须是"+strconv.Itoa(v.conditions[0])+"位！"
			} else if role =="number"{
				bol,_ = regexp.MatchString(`^[0-9]*\.?[0-9]*$`, ctl.InterToStr(req[v.PName])) //[0-9]*\.?[0-9]*$
				bol = !bol
				sinfo = "必须为数字！"
			} else if role =="string"{
				bol,_ =regexp.MatchString("^[a-zA-Z0-9]*$", ctl.InterToStr(req[v.PName]))
				bol = !bol
				sinfo = "必须为字符！"
			} else if role =="moblie"{
				//regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
				//reg := regexp.MustCompile(regular)
				//bol = reg.MatchString(ctl.InterToStr(req[v.PName]))
				bol,_ =regexp.MatchString("^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$", ctl.InterToStr(req[v.PName]))  //`/^1[3,5,8,7]{1}[\d]{9}$/`
				bol = !bol
				sinfo = "必须是有效手机号！"
			} else if role =="phone"{
				fmt.Println(ctl.InterToStr(req[v.PName]))
				bol,_ =regexp.MatchString("", ctl.InterToStr(req[v.PName]))  //`/^1[3,5,8,7]{1}[\d]{9}$/`
				bol = !bol
				sinfo = "必须是有效电话号码！"
			} else if role =="isChinese"{
				var a = regexp.MustCompile("^[\u4e00-\u9fa5]$")
				//接受正则表达式的范围
				for i, v := range  ctl.InterToStr(req[v.PName]) {
					//golang中string的底层是byte类型，所以单纯的for输出中文会出现乱码，这里选择for-range来输出
					if a.MatchString(string(v)) {
						//判断是否为中文，如果是返回一个true，不是返回false。这俩面MatchString的参数要求是string
						//但是 for-range 返回的 value 是 rune 类型，所以需要做一个 string() 转换
						fmt.Printf("str 字符串第 %v 个字符是中文。是“%v”字\n", i+1, string(v))
						bol = true
					}
				}

				sinfo = "不能包含中文！"
			} else if role =="url"{
				bol,_ =regexp.MatchString(`^([hH][tT]{2}[pP]:\/\/|[hH][tT]{2}[pP][sS]:\/\/)(([A-Za-z0-9-~]+)\.)+([A-Za-z0-9-~\/])+$`, ctl.InterToStr(req[v.PName]))
				bol = !bol
				sinfo = "URL格式错误！"
			} else if role =="email"{
				bol,_ =regexp.MatchString(`^([-_A-Za-z0-9\.]+)@([_A-Za-z0-9]+\.)+[A-Za-z0-9]{2,3}$`, ctl.InterToStr(req[v.PName]))
				bol = !bol
				sinfo = "邮箱格式错误！"
			} else if role =="max"{
				aleng = len(ctl.InterToStr(req[v.PName]))
				bol =  (aleng >= v.conditions[0])

				sinfo = "长度不能超出 "+strconv.Itoa(v.conditions[0])+"位!"
			} else if role =="min"{
				aleng = len(ctl.InterToStr(req[v.PName]))
				bol =  (aleng < v.conditions[0])

				sinfo = "长度不能低于 "+strconv.Itoa(v.conditions[0])+"位!"
			}

			if bol{
				result = utils.Code_ChkValid
				break I
			}
		}
	}
 	if result == utils.Code_ChkValid{
		serr = slabel+sinfo
	}
	return
}
//检查空值
func (ctl *Check_Base) chknull(aval interface{}) bool{
	var result bool
	result = false
	if aval !=nil{
		switch  aval.(type){
		case string:{
			if len(aval.(string)) ==0{
				result =true
			}
		}
		//case int,int32,int64:{
		//
		//}
		//case float32,float64:

		default:
			result = false
		}
	} else{
		result = true
	}

	return result
}
//interface转string
func (ctl *Check_Base) InterToStr(aval interface{}) string{
	var result string
	result = ""
	if aval !=nil{
		switch  aval.(type){
		case string:{
			result = aval.(string)
		}
		case int,int32,int64:{
			result = strconv.FormatInt(int64(aval.(int)), 10)
		}
		case float32,float64:{
			result = strconv.FormatFloat(aval.(float64), 'E', -1, 64)
		}
		default:
			result = ""
		}
	} else{
		result = ""
	}

	return result
}
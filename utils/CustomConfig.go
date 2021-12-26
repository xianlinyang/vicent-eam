package utils

import (
	"github.com/astaxie/beego"
)

// 客户信息记录单元

type CustomConfig struct {
	Svrport  	string
	DbName  	string
	Secret      string
	Bdate 	    string
	EDate     	string
}
//获取appid
func GetAppid(req beego.Controller)string{
	svrport := req.GetString(Custom_SvrPort,"")

	if svrport ==""{
		svrport = "0"
	}
	return svrport
}
//获取客户对应数据库
func GetDBName(asvrport string) CustomConfig{
	 requests := []CustomConfig{
	         CustomConfig{
	         	"8060",
	         	"EsaleEam_DB",
	         	"",
	         	"",
	         	"",
			 },
			 CustomConfig{
				 "8061",
				 "EsaleEam_DB_8061",
				 "",
				 "",
				 "",
			 },
	}

	var result = CustomConfig{};
	for _,v :=range requests{
		if v.Svrport == asvrport {
			result =v
		}
	}
	if result.DbName == ""{  //空为测试，打死数据库
		result.DbName ="EsaleEam_DB"
	}

	return result
};



package services

import "Eam_Server/utils"

// 数据库操作 公共方法，公共结构单元
//公共struct
type MSql_Get struct {
	W_and         map[string] string //and 字段内容
	W_or          map[string] string //or 字段内容
	W_like        map[string] string //like 字段内容
	W_larger      map[string] string // 大于 字段内容
	W_lessThan    map[string] string // 小于 字段内容
	Set_field  map[string] string //要插入的数据
}

func Sql_JionSet(amap map[string]string)string{
	var sql =""
	for key,value := range  amap {
		if sql ==""{
			if value ==utils.Data_null{
				sql = key+"= NULL"
			}else{
				sql = key+"= \""+value+"\""
			}

		} else{
			if value ==utils.Data_null{
				sql += ","+key+"= NULL"
			} else{
				sql += ","+key+"= \""+value+"\""
			}

		}
	}
	return sql
}


//拼接where条件
  //atype 0拼条件SQL 1拼插入或是修改SQL
func Sql_Joining(atype int64,amodel MSql_Get) string {
	var sql = " "
	if atype == 0{  //条件SQL
		sql = " 1=1 "
		if len(amodel.W_and) >0 {
			for key,value := range amodel.W_and {
				sql += " and "+key+"= \""+value+"\""
			}
		}
		if len(amodel.W_or) > 0{
			for key,value := range amodel.W_or {
				sql += " or "+key+"= \""+value+"\""
			}
			sql =" ( "+sql+" ) "
		}
		if len(amodel.W_like) > 0{
			for key,value := range amodel.W_or {
				sql += " and "+key+"like \"%"+value+"%\""
			}
		}
		if len(amodel.W_larger) > 0{
			for key,value := range amodel.W_or {
				sql += " and "+key+"> "+value
			}
		}
		if len(amodel.W_lessThan) > 0{
			for key,value := range amodel.W_or {
				sql += " and "+key+"< "+value
			}
		}
	}
	if atype == 1{ //插入SQL
		for key,value := range  amodel.Set_field {
			if sql ==""{
				sql = key+"= \""+value+"\""
			} else{
				sql += ","+key+"= \""+value+"\""
			}
		}
	}

	return sql
}

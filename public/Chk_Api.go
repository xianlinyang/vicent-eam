package public
// 接口参数有效性检测单元
import (
	"encoding/json"
)

type Check_API struct {
	Check_Base
}

func (ctl *Check_API) Chk_API_validate(res []byte,api string)(result int,serr string)  {
	var achk = []Check_Model{};

	result = 1
	switch  api{
		case "/api/User/create":{
			jsonstr := `[
                         {"Label": "用户编号", "PName":"f_no", "Rules":["notnull"],"conditions":[]},{"Label": "用户名称", "PName":"f_name", "Rules":["notnull"],"conditions":[]},
						 {"Label": "手机", "PName":"f_phone", "Rules":["isChinese"],"conditions":[]},{"Label": "创建人", "PName":"f_create_user", "Rules":["notnull"],"conditions":[]},
						 {"Label": "角色", "PName":"f_roleid", "Rules":["notnull"],"conditions":[]}]`; //f_create_user

			_= json.Unmarshal([]byte(jsonstr), &achk)

			result,serr = ctl.validate(res,achk)

		}
		case "/api/User/update":{
			jsonstr := `[{"Label":"用户ID","PName":"id","Rules":["notnull"],"conditions":[]},{"Label": "用户编号", "PName":"f_no", "Rules":["notnull"],"conditions":[]},
						 {"Label": "用户名称", "PName":"f_name", "Rules":["notnull"],"conditions":[]},{"Label": "手机", "PName":"f_phone", "Rules":["isChinese"],"conditions":[]},
						 {"Label": "最后修改人", "PName":"f_update_user", "Rules":["notnull"],"conditions":[]}]`;

			_= json.Unmarshal([]byte(jsonstr), &achk)
			result,serr = ctl.validate(res,achk)
		}
		case "/api/User/delete":{
			jsonstr := `[{"Label":"用户ID","PName":"id","Rules":["notnull"],"conditions":[]},{"Label": "删除人", "PName":"f_delete_user", "Rules":["notnull"],"conditions":[]}]`;
			_= json.Unmarshal([]byte(jsonstr), &achk)
			result,serr = ctl.validate(res,achk)
		}
		case "/api/Role/create":{
			jsonstr := `[{"Label":"角色名称","PName":"f_name","Rules":["notnull"],"conditions":[]},{"Label": "创建人", "PName":"f_create_user", "Rules":["notnull"],"conditions":[]} ]`;
			_= json.Unmarshal([]byte(jsonstr), &achk)
			result,serr = ctl.validate(res,achk)
		}
		case "/api/Role/update":{
			jsonstr := `[{"Label":"角色ID","PName":"id","Rules":["notnull"],"conditions":[]},{"Label":"角色名称","PName":"f_name","Rules":["notnull"],"conditions":[]} ]`;
			_= json.Unmarshal([]byte(jsonstr), &achk)
			result,serr = ctl.validate(res,achk)
		}
		case "/api/Role/delete":{
			jsonstr := `[{"Label":"角色ID","PName":"id","Rules":["notnull"],"conditions":[]},{"Label": "删除人", "PName":"f_delete_user", "Rules":["notnull"],"conditions":[]} ]`;
			_= json.Unmarshal([]byte(jsonstr), &achk)
			result,serr = ctl.validate(res,achk)
		}
		case "/api/Deptno/create","/api/AssetsType/create","/api/Supplier/create/type":{
			jsonstr := `[{"Label":"名称","PName":"f_name","Rules":["notnull"],"conditions":[]},{"Label": "创建人", "PName":"f_create_user", "Rules":["notnull"],"conditions":[]},
						 {"Label":"编号","PName":"f_no","Rules":["notnull"],"conditions":[]},{"Label": "父级ID", "PName":"f_parentid", "Rules":["notnull"],"conditions":[]}]`;
			_= json.Unmarshal([]byte(jsonstr), &achk)
			result,serr = ctl.validate(res,achk)
		}
		case "/api/Deptno/update","/api/AssetsType/update","/api/Supplier/update/type":{
			jsonstr := `[{"Label":"操作ID","PName":"id","Rules":["notnull"],"conditions":[]},{"Label":"名称","PName":"f_name","Rules":["notnull"],"conditions":[]},
						 {"Label": "最后修改人", "PName":"f_update_user", "Rules":["notnull"],"conditions":[]},{"Label": "父级ID", "PName":"f_parentid", "Rules":["notnull"],"conditions":[]}]`;
			_= json.Unmarshal([]byte(jsonstr), &achk)
			result,serr = ctl.validate(res,achk)
		}
		case "/api/Deptno/delete","/api/AssetsUnit/delete","/api/Supplier/delete","/api/Brands/delete","/api/AssetsType/delete","/api/AssetsList/delete","/api/Supplier/delete/type":{
			jsonstr := `[{"Label":"删除ID","PName":"id","Rules":["notnull"],"conditions":[]},{"Label": "删除人", "PName":"f_delete_user", "Rules":["notnull"],"conditions":[]} ]`;
			_= json.Unmarshal([]byte(jsonstr), &achk)
			result,serr = ctl.validate(res,achk)
		}
		case "/api/User/SetStatus":{
			jsonstr := `[{"Label":"用户ID","PName":"id","Rules":["notnull"],"conditions":[]},{"Label": "用户状态", "PName":"f_status", "Rules":["notnull"],"conditions":[]},
						{"Label": "修改人", "PName":"f_update_user", "Rules":["notnull"],"conditions":[]}]`;
			_= json.Unmarshal([]byte(jsonstr), &achk)
			result,serr = ctl.validate(res,achk)
		}
		case "/api/User/SetPwd":{
			jsonstr := `[{"Label":"用户ID","PName":"id","Rules":["notnull"],"conditions":[]},
						{"Label": "修改人", "PName":"f_update_user", "Rules":["notnull"],"conditions":[]}]`;
			_= json.Unmarshal([]byte(jsonstr), &achk)
			result,serr = ctl.validate(res,achk)
		}
		case "/api/Deptno/SetStatus":{
			jsonstr := `[{"Label":"部门ID","PName":"id","Rules":["notnull"],"conditions":[]},{"Label": "部门状态", "PName":"f_status", "Rules":["notnull"],"conditions":[]},
						{"Label": "修改人", "PName":"f_update_user", "Rules":["notnull"],"conditions":[]}]`;
			_= json.Unmarshal([]byte(jsonstr), &achk)
			result,serr = ctl.validate(res,achk)
		}
		case "/api/AssetsUnit/create", "/api/Brands/create":{
			jsonstr := `[
                         {"Label": "编号", "PName":"f_no", "Rules":["notnull"],"conditions":[]},{"Label": "名称", "PName":"f_name", "Rules":["notnull"],"conditions":[]},
						 {"Label": "创建人", "PName":"f_create_user", "Rules":["notnull"],"conditions":[]}
						  ]`; //f_create_user

			_= json.Unmarshal([]byte(jsonstr), &achk)

			result,serr = ctl.validate(res,achk)
		}
		case "/api/AssetsUnit/update","/api/Brands/update":{
			jsonstr := `[
                         {"Label": "编号", "PName":"f_no", "Rules":["notnull"],"conditions":[]},{"Label": "名称", "PName":"f_name", "Rules":["notnull"],"conditions":[]},
						 {"Label": "ID", "PName":"id",  "Rules":["notnull"],"conditions":[]},{"Label": "修改人", "PName":"f_update_user", "Rules":["notnull"],"conditions":[]}
						  ]`; //f_create_user

			_= json.Unmarshal([]byte(jsonstr), &achk)

			result,serr = ctl.validate(res,achk)
		}
		case "/api/Supplier/create":{
			jsonstr := `[
                         {"Label": "编号", "PName":"f_no", "Rules":["notnull"],"conditions":[]},{"Label": "名称", "PName":"f_name", "Rules":["notnull"],"conditions":[]},
						 {"Label": "创建人", "PName":"f_create_user", "Rules":["notnull"],"conditions":[]},{"Label": "类型ID", "PName":"f_typeid", "Rules":["notnull"],"conditions":[]}
						  ]`; //f_create_user

			_= json.Unmarshal([]byte(jsonstr), &achk)

			result,serr = ctl.validate(res,achk)
		}
		case "/api/Supplier/update":{
			jsonstr := `[
                         {"Label": "编号", "PName":"f_no", "Rules":["notnull"],"conditions":[]},{"Label": "名称", "PName":"f_name", "Rules":["notnull"],"conditions":[]},
						 {"Label": "ID", "PName":"id",  "Rules":["notnull"],"conditions":[]},{"Label": "创建人", "PName":"f_update_user", "Rules":["notnull"],"conditions":[]},{"Label": "类型ID", "PName":"f_typeid", "Rules":["notnull"],"conditions":[]}
						  ]`; //f_create_user

			_= json.Unmarshal([]byte(jsonstr), &achk)

			result,serr = ctl.validate(res,achk)
		}
		case "/api/AssetsList/create":{
			jsonstr := `[
                         {"Label": "资产编号", "PName":"f_Code", "Rules":["notnull"],"conditions":[]},{"Label": "资产名称", "PName":"f_Name", "Rules":["notnull"],"conditions":[]},
						 {"Label": "创建人", "PName":"f_create_user", "Rules":["notnull"],"conditions":[]},{"Label": "资产类别", "PName":"f_Typeid", "Rules":["notnull"],"conditions":[]}
						  ]`; //f_create_user

			_= json.Unmarshal([]byte(jsonstr), &achk)

			result,serr = ctl.validate(res,achk)
		}
		case "/api/AssetsList/update":{
			jsonstr := `[
                         {"Label": "资产编号", "PName":"f_Code", "Rules":["notnull"],"conditions":[]},{"Label": "资产名称", "PName":"f_Name", "Rules":["notnull"],"conditions":[]},
						 {"Label": "ID", "PName":"id",  "Rules":["notnull"],"conditions":[]},{"Label": "修改人", "PName":"f_update_user", "Rules":["notnull"],"conditions":[]},{"Label": "资产类别", "PName":"f_Typeid", "Rules":["notnull"],"conditions":[]}
						  ]`; //f_create_user

			_= json.Unmarshal([]byte(jsonstr), &achk)

			result,serr = ctl.validate(res,achk)
		}
		case "/api/Dictionary/create":{
			jsonstr := `[
                         {"Label": "编码", "PName":"f_no", "Rules":["notnull"],"conditions":[]},{"Label": "名称", "PName":"f_Ch_name", "Rules":["notnull"],"conditions":[]},
						 {"Label": "创建人", "PName":"f_create_user", "Rules":["notnull"],"conditions":[]},{"Label": "父级", "PName":"f_parentid", "Rules":["notnull"],"conditions":[]}
						  ]`; //f_create_user

			_= json.Unmarshal([]byte(jsonstr), &achk)

			result,serr = ctl.validate(res,achk)
		}
		case "/api/Dictionary/update":{
			jsonstr := `[{"Label": "ID", "PName":"id", "Rules":["notnull"],"conditions":[]},{"Label": "编码", "PName":"f_no", "Rules":["notnull"],"conditions":[]},{"Label": "名称", "PName":"f_Ch_name", "Rules":["notnull"],"conditions":[]},
						 {"Label": "修改人", "PName":"f_update_user", "Rules":["notnull"],"conditions":[]},{"Label": "父级", "PName":"f_parentid", "Rules":["notnull"],"conditions":[]}]`; //f_create_user

			_ = json.Unmarshal([]byte(jsonstr), &achk)
			result,serr = ctl.validate(res,achk)

		}
		case "/api/Dictionary/create/type":{
			jsonstr := `[
                         {"Label": "名称", "PName":"f_Ch_name", "Rules":["notnull"],"conditions":[]},
						 {"Label": "创建人", "PName":"f_create_user", "Rules":["notnull"],"conditions":[]},{"Label": "父级", "PName":"f_parentid", "Rules":["notnull"],"conditions":[]}
						  ]`; //f_create_user

			_= json.Unmarshal([]byte(jsonstr), &achk)

			result,serr = ctl.validate(res,achk)
		}
		case "/api/Dictionary/update/type":{
			jsonstr := `[
                         {"Label": "ID", "PName":"id", "Rules":["notnull"],"conditions":[]},{"Label": "名称", "PName":"f_Ch_name", "Rules":["notnull"],"conditions":[]},
						 {"Label": "修改人", "PName":"f_update_user", "Rules":["notnull"],"conditions":[]},{"Label": "父级", "PName":"f_parentid", "Rules":["notnull"],"conditions":[]}
						  ]`; //f_create_user

			_= json.Unmarshal([]byte(jsonstr), &achk)

			result,serr = ctl.validate(res,achk)
		}

		////
		//"/api/Deptno/update","/api/AssetsType/update","/api/BuyModel/update","/api/StoreAddress/update","/api/SavePeople/update","/api/AssetsUnit/update","/api/AssetsUse/update","/api/AssetsState/update","/api/Supplier/update"
		//"/api/Deptno/delete","/api/AssetsType/delete","/api/BuyModel/delete","/api/StoreAddress/delete","/api/SavePeople/delete","/api/AssetsUnit/delete","/api/AssetsUse/delete","/api/AssetsState/delete","/api/Supplier/delete"
	}
	return
}


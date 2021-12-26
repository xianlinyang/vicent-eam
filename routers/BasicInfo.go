package routers

import (
"Eam_Server/controllers"
"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//许可信息查询
	beego.Router("/api/License/browse", &controllers.BasicController{},"get:LicenseQuery")
	//部门管理
	beego.Router("/api/Deptno/DeptnoTrelistQuery", &controllers.BasicController{},"get:DeptnoTrelistQuery")
	beego.Router("/api/Deptno/browse", &controllers.BasicController{},"get:DeptnoQuery")
	beego.Router("/api/Deptno/create", &controllers.BasicController{},"post:DeptnoAdd")
	beego.Router("/api/Deptno/update", &controllers.BasicController{},"put:DeptnoUpdate")
	beego.Router("/api/Deptno/delete", &controllers.BasicController{},"delete:DeptnoDel")
	beego.Router("/api/Deptno/SetStatus", &controllers.BasicController{},"post:DeptnoSetStatus")  //启用，停用接口

	//资产类型
	beego.Router("/api/AssetsType/TreeList", &controllers.BasicController{},"get:AssetsTypeTreeList")
	beego.Router("/api/AssetsType/browse", &controllers.BasicController{},"get:AssetsTypeQuery")
	beego.Router("/api/AssetsType/create", &controllers.BasicController{},"post:AssetsTypeAdd")
	beego.Router("/api/AssetsType/update", &controllers.BasicController{},"put:AssetsTypeUpdate")
	beego.Router("/api/AssetsType/delete", &controllers.BasicController{},"delete:AssetsTypeDel")

	////品牌  --暂时不用
	//beego.Router("/api/Brands/browse", &controllers.BasicController{},"get:BrandsQuery")
	//beego.Router("/api/Brands/create", &controllers.BasicController{},"post:BrandsAdd")
	//beego.Router("/api/Brands/update", &controllers.BasicController{},"put:BrandsUpdate")
	//beego.Router("/api/Brands/delete", &controllers.BasicController{},"delete:BrandsDel")
	//存放地点
	beego.Router("/api/StoreAddress/TreeList", &controllers.BasicController{},"get:StoreAddressTreeList")
	beego.Router("/api/StoreAddress/browse", &controllers.BasicController{},"get:StoreAddressQuery")
	beego.Router("/api/StoreAddress/create", &controllers.BasicController{},"post:StoreAddressAdd")
	beego.Router("/api/StoreAddress/update", &controllers.BasicController{},"put:StoreAddressUpdate")
	beego.Router("/api/StoreAddress/delete", &controllers.BasicController{},"delete:StoreAddressDel")

	//供应商类型
	beego.Router("/api/Supplier/Type_TreeList", &controllers.BasicController{},"get:SupplierTreeList_type")
	beego.Router("/api/Supplier/browse/type", &controllers.BasicController{},"get:SupplierQuery_type")
	beego.Router("/api/Supplier/create/type", &controllers.BasicController{},"post:SupplierAdd_type")
	beego.Router("/api/Supplier/update/type", &controllers.BasicController{},"put:SupplierUpdate_type")
	beego.Router("/api/Supplier/delete/type", &controllers.BasicController{},"delete:SupplierDel_type")
	//供应商
	beego.Router("/api/Supplier/browse", &controllers.BasicController{},"get:SupplierQuery")
	beego.Router("/api/Supplier/create", &controllers.BasicController{},"post:SupplierAdd")
	beego.Router("/api/Supplier/update", &controllers.BasicController{},"put:SupplierUpdate")
	beego.Router("/api/Supplier/delete", &controllers.BasicController{},"delete:SupplierDel")

	////资产单位 --暂时不用
	//beego.Router("/api/AssetsUnit/browse", &controllers.BasicController{},"get:AssetsUnitQuery")
	//beego.Router("/api/AssetsUnit/create", &controllers.BasicController{},"post:AssetsUnitAdd")
	//beego.Router("/api/AssetsUnit/update", &controllers.BasicController{},"put:AssetsUnitUpdate")
	//beego.Router("/api/AssetsUnit/delete", &controllers.BasicController{},"delete:AssetsUnitDel")

	//商品清单
	beego.Router("/api/AssetsList/brows/AllList", &controllers.BasicController{},"get:AssetsListQuery_AllList")  //新增商品时查询所有列表信息
	beego.Router("/api/AssetsList/browse", &controllers.BasicController{},"get:AssetsListQuery")
	beego.Router("/api/AssetsList/create", &controllers.BasicController{},"post:AssetsListAdd")
	beego.Router("/api/AssetsList/update", &controllers.BasicController{},"put:AssetsListUpdate")
	beego.Router("/api/AssetsList/delete", &controllers.BasicController{},"delete:AssetsListDel")
	beego.Router("/api/AssetsList/AssetsImport", &controllers.BasicController{},"post:AssetsImport")


    //数据字典
	beego.Router("/api/Dictionary/brows/type", &controllers.BasicController{},"get:Dictionary_Type_GetTreeList")
	beego.Router("/api/Dictionary/create/type", &controllers.BasicController{},"post:Dictionary_TypeAdd")
	beego.Router("/api/Dictionary/update/type", &controllers.BasicController{},"put:Dictionary_TypeUpdate")
	beego.Router("/api/Dictionary/delete/type", &controllers.BasicController{},"delete:Dictionary_TypeDel")


	beego.Router("/api/Dictionary/browse", &controllers.BasicController{},"get:DictionaryQuery")
	beego.Router("/api/Dictionary/create", &controllers.BasicController{},"post:DictionaryAdd")
	beego.Router("/api/Dictionary/update", &controllers.BasicController{},"put:DictionaryUpdate")
	beego.Router("/api/Dictionary/delete", &controllers.BasicController{},"delete:DictionaryDel")


	// 外部联系人
	beego.Router("/api/ContactOuter/browse", &controllers.BasicController{},"get:ContactOuter_Browse")
	beego.Router("/api/ContactOuter/create", &controllers.BasicController{},"post:ContactOuter_Create")
	beego.Router("/api/ContactOuter/update", &controllers.BasicController{},"put:ContactOuter_Update")
	beego.Router("/api/ContactOuter/delete", &controllers.BasicController{},"delete:ContactOuter_Delete")
	beego.Router("/api/ContactOuter/start", &controllers.BasicController{},"put:ContactOuter_Start")
	beego.Router("/api/ContactOuter/stop", &controllers.BasicController{},"put:ContactOuter_Stop")



	////保管人，使用人
	//beego.Router("/api/SavePeople/browse", &controllers.BasicController{},"get:SavePeopleQuery")
	//beego.Router("/api/SavePeople/create", &controllers.BasicController{},"post:SavePeopleAdd")
	//beego.Router("/api/SavePeople/update", &controllers.BasicController{},"put:SavePeopleUpdate")
	//beego.Router("/api/SavePeople/delete", &controllers.BasicController{},"post:SavePeopleDel")
	//beego.Router("/api/SavePeople/SetStatus", &controllers.BasicController{},"post:SavePeopleSetStatus")



	////资产用途
	//beego.Router("/api/AssetsUse/browse", &controllers.BasicController{},"get:AssetsUseQuery")
	//beego.Router("/api/AssetsUse/create", &controllers.BasicController{},"post:AssetsUseAdd")
	//beego.Router("/api/AssetsUse/update", &controllers.BasicController{},"put:AssetsUseUpdate")
	//beego.Router("/api/AssetsUse/delete", &controllers.BasicController{},"post:AssetsUseDel")

	////资产状态
	//beego.Router("/api/AssetsState/browse", &controllers.BasicController{},"get:AssetsStateQuery")
	//beego.Router("/api/AssetsState/create", &controllers.BasicController{},"post:AssetsStateAdd")
	//beego.Router("/api/AssetsState/update", &controllers.BasicController{},"put:AssetsStateUpdate")
	//beego.Router("/api/AssetsState/delete", &controllers.BasicController{},"post:AssetsStateDel")



}


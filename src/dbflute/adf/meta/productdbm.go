package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type ProductDbm_T struct {
	df.BaseDBMeta
	ColumnProductId *df.ColumnInfo
	ColumnProductName *df.ColumnInfo
	ColumnProductHandleCode *df.ColumnInfo
	ColumnProductCategoryCode *df.ColumnInfo
	ColumnProductStatusCode *df.ColumnInfo
	ColumnRegularPrice *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnRegisterProcess *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnUpdateProcess *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
}

func (b *ProductDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *ProductDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var ProductDbm *ProductDbm_T

func Create_ProductDbm() {
	ProductDbm = new(ProductDbm_T)
	ProductDbm.TableDbName = "PRODUCT"
	ProductDbm.TableDispName = "PRODUCT"
	ProductDbm.TablePropertyName = "product"
	ProductDbm.TableSqlName = new(df.TableSqlName)
	ProductDbm.TableSqlName.TableSqlName = "exampledb.dbo.PRODUCT"
	ProductDbm.TableSqlName.CorrespondingDbName = ProductDbm.TableDbName
	ProductDbm.Identity=true

	var product df.DBMeta
	product = ProductDbm
	ProductDbm.DBMeta=&product
	productIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo PRODUCT_ID
	productIdSqlName.ColumnSqlName = "PRODUCT_ID"
	productIdSqlName.IrregularChar = false
	ProductDbm.ColumnProductId = df.CCI(&product, "PRODUCT_ID", productIdSqlName, "", "", "Integer.class", "productId", "", true, true,true, "int identity", 10, 0, "",false,"","", "","purchaseList","",false,"int64")
	productNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo PRODUCT_NAME
	productNameSqlName.ColumnSqlName = "PRODUCT_NAME"
	productNameSqlName.IrregularChar = false
	ProductDbm.ColumnProductName = df.CCI(&product, "PRODUCT_NAME", productNameSqlName, "", "", "String.class", "productName", "", false, false,true, "nvarchar", 50, 0, "",false,"","", "","","",false,"string")
	productHandleCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo PRODUCT_HANDLE_CODE
	productHandleCodeSqlName.ColumnSqlName = "PRODUCT_HANDLE_CODE"
	productHandleCodeSqlName.IrregularChar = false
	ProductDbm.ColumnProductHandleCode = df.CCI(&product, "PRODUCT_HANDLE_CODE", productHandleCodeSqlName, "", "", "String.class", "productHandleCode", "", false, false,true, "nvarchar", 100, 0, "",false,"","", "","","",false,"string")
	productCategoryCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo PRODUCT_CATEGORY_CODE
	productCategoryCodeSqlName.ColumnSqlName = "PRODUCT_CATEGORY_CODE"
	productCategoryCodeSqlName.IrregularChar = false
	ProductDbm.ColumnProductCategoryCode = df.CCI(&product, "PRODUCT_CATEGORY_CODE", productCategoryCodeSqlName, "", "", "String.class", "productCategoryCode", "", false, false,true, "char", 3, 0, "",false,"","", "productCategory","","",false,"string")
	productStatusCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo PRODUCT_STATUS_CODE
	productStatusCodeSqlName.ColumnSqlName = "PRODUCT_STATUS_CODE"
	productStatusCodeSqlName.IrregularChar = false
	ProductDbm.ColumnProductStatusCode = df.CCI(&product, "PRODUCT_STATUS_CODE", productStatusCodeSqlName, "", "", "String.class", "productStatusCode", "", false, false,true, "char", 3, 0, "",false,"","", "productStatus","","",false,"string")
	regularPriceSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo REGULAR_PRICE
	regularPriceSqlName.ColumnSqlName = "REGULAR_PRICE"
	regularPriceSqlName.IrregularChar = false
	ProductDbm.ColumnRegularPrice = df.CCI(&product, "REGULAR_PRICE", regularPriceSqlName, "", "", "Integer.class", "regularPrice", "", false, false,true, "int", 10, 0, "",false,"","", "","","",false,"int64")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo REGISTER_DATETIME
	registerDatetimeSqlName.ColumnSqlName = "REGISTER_DATETIME"
	registerDatetimeSqlName.IrregularChar = false
	ProductDbm.ColumnRegisterDatetime = df.CCI(&product, "REGISTER_DATETIME", registerDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "datetime", 23, 3, "",false,"","", "","","",false,"df.Timestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo REGISTER_USER
	registerUserSqlName.ColumnSqlName = "REGISTER_USER"
	registerUserSqlName.IrregularChar = false
	ProductDbm.ColumnRegisterUser = df.CCI(&product, "REGISTER_USER", registerUserSqlName, "", "", "String.class", "registerUser", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	registerProcessSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo REGISTER_PROCESS
	registerProcessSqlName.ColumnSqlName = "REGISTER_PROCESS"
	registerProcessSqlName.IrregularChar = false
	ProductDbm.ColumnRegisterProcess = df.CCI(&product, "REGISTER_PROCESS", registerProcessSqlName, "", "", "String.class", "registerProcess", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo UPDATE_DATETIME
	updateDatetimeSqlName.ColumnSqlName = "UPDATE_DATETIME"
	updateDatetimeSqlName.IrregularChar = false
	ProductDbm.ColumnUpdateDatetime = df.CCI(&product, "UPDATE_DATETIME", updateDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "datetime", 23, 3, "",false,"","", "","","",false,"df.Timestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo UPDATE_USER
	updateUserSqlName.ColumnSqlName = "UPDATE_USER"
	updateUserSqlName.IrregularChar = false
	ProductDbm.ColumnUpdateUser = df.CCI(&product, "UPDATE_USER", updateUserSqlName, "", "", "String.class", "updateUser", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	updateProcessSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo UPDATE_PROCESS
	updateProcessSqlName.ColumnSqlName = "UPDATE_PROCESS"
	updateProcessSqlName.IrregularChar = false
	ProductDbm.ColumnUpdateProcess = df.CCI(&product, "UPDATE_PROCESS", updateProcessSqlName, "", "", "String.class", "updateProcess", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	versionNoSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo VERSION_NO
	versionNoSqlName.ColumnSqlName = "VERSION_NO"
	versionNoSqlName.IrregularChar = false
	ProductDbm.ColumnVersionNo = df.CCI(&product, "VERSION_NO", versionNoSqlName, "", "", "Long.class", "versionNo", "", false, false,true, "bigint", 19, 0, "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")

	ProductDbm.ColumnInfoList = new(df.List)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnProductId)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnProductName)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnProductHandleCode)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnProductCategoryCode)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnProductStatusCode)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnRegularPrice)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnRegisterDatetime)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnRegisterUser)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnRegisterProcess)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnUpdateDatetime)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnUpdateUser)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnUpdateProcess)
	ProductDbm.ColumnInfoList.Add(ProductDbm.ColumnVersionNo)


	ProductDbm.ColumnInfoMap=make(map[string]int)
	ProductDbm.ColumnInfoMap["productId"]=0
		ProductDbm.ColumnInfoMap["productName"]=1
		ProductDbm.ColumnInfoMap["productHandleCode"]=2
		ProductDbm.ColumnInfoMap["productCategoryCode"]=3
		ProductDbm.ColumnInfoMap["productStatusCode"]=4
		ProductDbm.ColumnInfoMap["regularPrice"]=5
		ProductDbm.ColumnInfoMap["registerDatetime"]=6
		ProductDbm.ColumnInfoMap["registerUser"]=7
		ProductDbm.ColumnInfoMap["registerProcess"]=8
		ProductDbm.ColumnInfoMap["updateDatetime"]=9
		ProductDbm.ColumnInfoMap["updateUser"]=10
		ProductDbm.ColumnInfoMap["updateProcess"]=11
		ProductDbm.ColumnInfoMap["versionNo"]=12
	    ProductDbm.PrimaryKey = true
    ProductDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &product
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(ProductDbm.ColumnProductId)

	ProductDbm.PrimaryInfo = new(df.PrimaryInfo)
	ProductDbm.PrimaryInfo.UniqueInfo = ui
	
	ProductDbm.VersionNoFlag = true
	ProductDbm.VersionNoColumnInfo = ProductDbm.ColumnVersionNo
	
	var productMeta df.DBMeta = ProductDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["Product"] = &productMeta
}

package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type PurchaseDbm_T struct {
	df.BaseDBMeta
	ColumnPurchaseId *df.ColumnInfo
	ColumnMemberId *df.ColumnInfo
	ColumnProductId *df.ColumnInfo
	ColumnPurchaseDatetime *df.ColumnInfo
	ColumnPurchaseCount *df.ColumnInfo
	ColumnPurchasePrice *df.ColumnInfo
	ColumnPaymentCompleteFlg *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnRegisterProcess *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnUpdateProcess *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
}

func (b *PurchaseDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *PurchaseDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var PurchaseDbm *PurchaseDbm_T

func Create_PurchaseDbm() {
	PurchaseDbm = new(PurchaseDbm_T)
	PurchaseDbm.TableDbName = "PURCHASE"
	PurchaseDbm.TableDispName = "PURCHASE"
	PurchaseDbm.TablePropertyName = "purchase"
	PurchaseDbm.TableSqlName = new(df.TableSqlName)
	PurchaseDbm.TableSqlName.TableSqlName = "exampledb.dbo.PURCHASE"
	PurchaseDbm.TableSqlName.CorrespondingDbName = PurchaseDbm.TableDbName
	PurchaseDbm.Identity=true

	var purchase df.DBMeta
	purchase = PurchaseDbm
	PurchaseDbm.DBMeta=&purchase
	purchaseIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo PURCHASE_ID
	purchaseIdSqlName.ColumnSqlName = "PURCHASE_ID"
	purchaseIdSqlName.IrregularChar = false
	PurchaseDbm.ColumnPurchaseId = df.CCI(&purchase, "PURCHASE_ID", purchaseIdSqlName, "", "", "Long.class", "purchaseId", "", true, true,true, "bigint identity", 19, 0, "",false,"","", "","","",false,"int64")
	memberIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo MEMBER_ID
	memberIdSqlName.ColumnSqlName = "MEMBER_ID"
	memberIdSqlName.IrregularChar = false
	PurchaseDbm.ColumnMemberId = df.CCI(&purchase, "MEMBER_ID", memberIdSqlName, "", "", "Integer.class", "memberId", "", false, false,true, "int", 10, 0, "",false,"","", "member","","",false,"int64")
	productIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo PRODUCT_ID
	productIdSqlName.ColumnSqlName = "PRODUCT_ID"
	productIdSqlName.IrregularChar = false
	PurchaseDbm.ColumnProductId = df.CCI(&purchase, "PRODUCT_ID", productIdSqlName, "", "", "Integer.class", "productId", "", false, false,true, "int", 10, 0, "",false,"","", "product","","",false,"int64")
	purchaseDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo PURCHASE_DATETIME
	purchaseDatetimeSqlName.ColumnSqlName = "PURCHASE_DATETIME"
	purchaseDatetimeSqlName.IrregularChar = false
	PurchaseDbm.ColumnPurchaseDatetime = df.CCI(&purchase, "PURCHASE_DATETIME", purchaseDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "purchaseDatetime", "", false, false,true, "datetime", 23, 3, "",false,"","", "","","",false,"df.Timestamp")
	purchaseCountSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo PURCHASE_COUNT
	purchaseCountSqlName.ColumnSqlName = "PURCHASE_COUNT"
	purchaseCountSqlName.IrregularChar = false
	PurchaseDbm.ColumnPurchaseCount = df.CCI(&purchase, "PURCHASE_COUNT", purchaseCountSqlName, "", "", "Integer.class", "purchaseCount", "", false, false,true, "int", 10, 0, "",false,"","", "","","",false,"int64")
	purchasePriceSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo PURCHASE_PRICE
	purchasePriceSqlName.ColumnSqlName = "PURCHASE_PRICE"
	purchasePriceSqlName.IrregularChar = false
	PurchaseDbm.ColumnPurchasePrice = df.CCI(&purchase, "PURCHASE_PRICE", purchasePriceSqlName, "", "", "Integer.class", "purchasePrice", "", false, false,true, "int", 10, 0, "",false,"","", "","","",false,"int64")
	paymentCompleteFlgSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo PAYMENT_COMPLETE_FLG
	paymentCompleteFlgSqlName.ColumnSqlName = "PAYMENT_COMPLETE_FLG"
	paymentCompleteFlgSqlName.IrregularChar = false
	PurchaseDbm.ColumnPaymentCompleteFlg = df.CCI(&purchase, "PAYMENT_COMPLETE_FLG", paymentCompleteFlgSqlName, "", "", "Integer.class", "paymentCompleteFlg", "", false, false,true, "int", 10, 0, "",false,"","", "","","",false,"int64")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo REGISTER_DATETIME
	registerDatetimeSqlName.ColumnSqlName = "REGISTER_DATETIME"
	registerDatetimeSqlName.IrregularChar = false
	PurchaseDbm.ColumnRegisterDatetime = df.CCI(&purchase, "REGISTER_DATETIME", registerDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "datetime", 23, 3, "",false,"","", "","","",false,"df.Timestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo REGISTER_USER
	registerUserSqlName.ColumnSqlName = "REGISTER_USER"
	registerUserSqlName.IrregularChar = false
	PurchaseDbm.ColumnRegisterUser = df.CCI(&purchase, "REGISTER_USER", registerUserSqlName, "", "", "String.class", "registerUser", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	registerProcessSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo REGISTER_PROCESS
	registerProcessSqlName.ColumnSqlName = "REGISTER_PROCESS"
	registerProcessSqlName.IrregularChar = false
	PurchaseDbm.ColumnRegisterProcess = df.CCI(&purchase, "REGISTER_PROCESS", registerProcessSqlName, "", "", "String.class", "registerProcess", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo UPDATE_DATETIME
	updateDatetimeSqlName.ColumnSqlName = "UPDATE_DATETIME"
	updateDatetimeSqlName.IrregularChar = false
	PurchaseDbm.ColumnUpdateDatetime = df.CCI(&purchase, "UPDATE_DATETIME", updateDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "datetime", 23, 3, "",false,"","", "","","",false,"df.Timestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo UPDATE_USER
	updateUserSqlName.ColumnSqlName = "UPDATE_USER"
	updateUserSqlName.IrregularChar = false
	PurchaseDbm.ColumnUpdateUser = df.CCI(&purchase, "UPDATE_USER", updateUserSqlName, "", "", "String.class", "updateUser", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	updateProcessSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo UPDATE_PROCESS
	updateProcessSqlName.ColumnSqlName = "UPDATE_PROCESS"
	updateProcessSqlName.IrregularChar = false
	PurchaseDbm.ColumnUpdateProcess = df.CCI(&purchase, "UPDATE_PROCESS", updateProcessSqlName, "", "", "String.class", "updateProcess", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	versionNoSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo VERSION_NO
	versionNoSqlName.ColumnSqlName = "VERSION_NO"
	versionNoSqlName.IrregularChar = false
	PurchaseDbm.ColumnVersionNo = df.CCI(&purchase, "VERSION_NO", versionNoSqlName, "", "", "Long.class", "versionNo", "", false, false,true, "bigint", 19, 0, "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")

	PurchaseDbm.ColumnInfoList = new(df.List)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnPurchaseId)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnMemberId)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnProductId)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnPurchaseDatetime)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnPurchaseCount)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnPurchasePrice)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnPaymentCompleteFlg)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnRegisterDatetime)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnRegisterUser)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnRegisterProcess)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnUpdateDatetime)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnUpdateUser)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnUpdateProcess)
	PurchaseDbm.ColumnInfoList.Add(PurchaseDbm.ColumnVersionNo)


	PurchaseDbm.ColumnInfoMap=make(map[string]int)
	PurchaseDbm.ColumnInfoMap["purchaseId"]=0
		PurchaseDbm.ColumnInfoMap["memberId"]=1
		PurchaseDbm.ColumnInfoMap["productId"]=2
		PurchaseDbm.ColumnInfoMap["purchaseDatetime"]=3
		PurchaseDbm.ColumnInfoMap["purchaseCount"]=4
		PurchaseDbm.ColumnInfoMap["purchasePrice"]=5
		PurchaseDbm.ColumnInfoMap["paymentCompleteFlg"]=6
		PurchaseDbm.ColumnInfoMap["registerDatetime"]=7
		PurchaseDbm.ColumnInfoMap["registerUser"]=8
		PurchaseDbm.ColumnInfoMap["registerProcess"]=9
		PurchaseDbm.ColumnInfoMap["updateDatetime"]=10
		PurchaseDbm.ColumnInfoMap["updateUser"]=11
		PurchaseDbm.ColumnInfoMap["updateProcess"]=12
		PurchaseDbm.ColumnInfoMap["versionNo"]=13
	    PurchaseDbm.PrimaryKey = true
    PurchaseDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &purchase
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(PurchaseDbm.ColumnPurchaseId)

	PurchaseDbm.PrimaryInfo = new(df.PrimaryInfo)
	PurchaseDbm.PrimaryInfo.UniqueInfo = ui
	
	PurchaseDbm.VersionNoFlag = true
	PurchaseDbm.VersionNoColumnInfo = PurchaseDbm.ColumnVersionNo
	
	var purchaseMeta df.DBMeta = PurchaseDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["Purchase"] = &purchaseMeta
}

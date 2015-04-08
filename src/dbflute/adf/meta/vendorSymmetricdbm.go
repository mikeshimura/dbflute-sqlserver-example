package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorSymmetricDbm_T struct {
	df.BaseDBMeta
	ColumnVendorSymmetricId *df.ColumnInfo
	ColumnPlainText *df.ColumnInfo
	ColumnEncryptedData *df.ColumnInfo
}

func (b *VendorSymmetricDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *VendorSymmetricDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var VendorSymmetricDbm *VendorSymmetricDbm_T

func Create_VendorSymmetricDbm() {
	VendorSymmetricDbm = new(VendorSymmetricDbm_T)
	VendorSymmetricDbm.TableDbName = "VENDOR_SYMMETRIC"
	VendorSymmetricDbm.TableDispName = "VENDOR_SYMMETRIC"
	VendorSymmetricDbm.TablePropertyName = "vendorSymmetric"
	VendorSymmetricDbm.TableSqlName = new(df.TableSqlName)
	VendorSymmetricDbm.TableSqlName.TableSqlName = "exampledb.dbo.VENDOR_SYMMETRIC"
	VendorSymmetricDbm.TableSqlName.CorrespondingDbName = VendorSymmetricDbm.TableDbName

	var vendorSymmetric df.DBMeta
	vendorSymmetric = VendorSymmetricDbm
	VendorSymmetricDbm.DBMeta=&vendorSymmetric
	vendorSymmetricIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo VENDOR_SYMMETRIC_ID
	vendorSymmetricIdSqlName.ColumnSqlName = "VENDOR_SYMMETRIC_ID"
	vendorSymmetricIdSqlName.IrregularChar = false
	VendorSymmetricDbm.ColumnVendorSymmetricId = df.CCI(&vendorSymmetric, "VENDOR_SYMMETRIC_ID", vendorSymmetricIdSqlName, "", "", "Long.class", "vendorSymmetricId", "", true, false,true, "numeric", 16, 0, "",false,"","", "","","",false,"df.Numeric")
	plainTextSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo PLAIN_TEXT
	plainTextSqlName.ColumnSqlName = "PLAIN_TEXT"
	plainTextSqlName.IrregularChar = false
	VendorSymmetricDbm.ColumnPlainText = df.CCI(&vendorSymmetric, "PLAIN_TEXT", plainTextSqlName, "", "", "String.class", "plainText", "", false, false,false, "nvarchar", 100, 0, "",false,"","", "","","",false,"df.NullString")
	encryptedDataSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo ENCRYPTED_DATA
	encryptedDataSqlName.ColumnSqlName = "ENCRYPTED_DATA"
	encryptedDataSqlName.IrregularChar = false
	VendorSymmetricDbm.ColumnEncryptedData = df.CCI(&vendorSymmetric, "ENCRYPTED_DATA", encryptedDataSqlName, "", "", "byte[].class", "encryptedData", "", false, false,false, "varbinary", 2147483647, 0, "",false,"","", "","","",false,"[]byte")

	VendorSymmetricDbm.ColumnInfoList = new(df.List)
	VendorSymmetricDbm.ColumnInfoList.Add(VendorSymmetricDbm.ColumnVendorSymmetricId)
	VendorSymmetricDbm.ColumnInfoList.Add(VendorSymmetricDbm.ColumnPlainText)
	VendorSymmetricDbm.ColumnInfoList.Add(VendorSymmetricDbm.ColumnEncryptedData)


	VendorSymmetricDbm.ColumnInfoMap=make(map[string]int)
	VendorSymmetricDbm.ColumnInfoMap["vendorSymmetricId"]=0
		VendorSymmetricDbm.ColumnInfoMap["plainText"]=1
		VendorSymmetricDbm.ColumnInfoMap["encryptedData"]=2
	    VendorSymmetricDbm.PrimaryKey = true
    VendorSymmetricDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &vendorSymmetric
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(VendorSymmetricDbm.ColumnVendorSymmetricId)

	VendorSymmetricDbm.PrimaryInfo = new(df.PrimaryInfo)
	VendorSymmetricDbm.PrimaryInfo.UniqueInfo = ui
	
	var vendorSymmetricMeta df.DBMeta = VendorSymmetricDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["VendorSymmetric"] = &vendorSymmetricMeta
}

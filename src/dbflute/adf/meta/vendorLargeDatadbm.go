package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorLargeDataDbm_T struct {
	df.BaseDBMeta
	ColumnLargeDataId *df.ColumnInfo
	ColumnStringIndex *df.ColumnInfo
	ColumnStringNoIndex *df.ColumnInfo
	ColumnStringUniqueIndex *df.ColumnInfo
	ColumnIntflgIndex *df.ColumnInfo
	ColumnNumericIntegerIndex *df.ColumnInfo
	ColumnNumericIntegerNoIndex *df.ColumnInfo
}

func (b *VendorLargeDataDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *VendorLargeDataDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var VendorLargeDataDbm *VendorLargeDataDbm_T

func Create_VendorLargeDataDbm() {
	VendorLargeDataDbm = new(VendorLargeDataDbm_T)
	VendorLargeDataDbm.TableDbName = "vendor_large_data"
	VendorLargeDataDbm.TableDispName = "vendor_large_data"
	VendorLargeDataDbm.TablePropertyName = "vendorLargeData"
	VendorLargeDataDbm.TableSqlName = new(df.TableSqlName)
	VendorLargeDataDbm.TableSqlName.TableSqlName = "vendor_large_data"
	VendorLargeDataDbm.TableSqlName.CorrespondingDbName = VendorLargeDataDbm.TableDbName

	var vendorLargeData df.DBMeta
	vendorLargeData = VendorLargeDataDbm
	VendorLargeDataDbm.DBMeta=&vendorLargeData
	largeDataIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo LARGE_DATA_ID
	largeDataIdSqlName.ColumnSqlName = "LARGE_DATA_ID"
	largeDataIdSqlName.IrregularChar = false
	VendorLargeDataDbm.ColumnLargeDataId = df.CCI(&vendorLargeData, "LARGE_DATA_ID", largeDataIdSqlName, "", "", "Long.class", "largeDataId", "", true, false,true, "BIGINT", 19, 0, "",false,"","", "","","",false,"int64")
	stringIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo STRING_INDEX
	stringIndexSqlName.ColumnSqlName = "STRING_INDEX"
	stringIndexSqlName.IrregularChar = false
	VendorLargeDataDbm.ColumnStringIndex = df.CCI(&vendorLargeData, "STRING_INDEX", stringIndexSqlName, "", "", "String.class", "stringIndex", "", false, false,true, "VARCHAR", 180, 0, "",false,"","", "","","",false,"string")
	stringNoIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo STRING_NO_INDEX
	stringNoIndexSqlName.ColumnSqlName = "STRING_NO_INDEX"
	stringNoIndexSqlName.IrregularChar = false
	VendorLargeDataDbm.ColumnStringNoIndex = df.CCI(&vendorLargeData, "STRING_NO_INDEX", stringNoIndexSqlName, "", "", "String.class", "stringNoIndex", "", false, false,true, "VARCHAR", 180, 0, "",false,"","", "","","",false,"string")
	stringUniqueIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo STRING_UNIQUE_INDEX
	stringUniqueIndexSqlName.ColumnSqlName = "STRING_UNIQUE_INDEX"
	stringUniqueIndexSqlName.IrregularChar = false
	VendorLargeDataDbm.ColumnStringUniqueIndex = df.CCI(&vendorLargeData, "STRING_UNIQUE_INDEX", stringUniqueIndexSqlName, "", "", "String.class", "stringUniqueIndex", "", false, false,true, "VARCHAR", 180, 0, "",false,"","", "","","",false,"string")
	intflgIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo INTFLG_INDEX
	intflgIndexSqlName.ColumnSqlName = "INTFLG_INDEX"
	intflgIndexSqlName.IrregularChar = false
	VendorLargeDataDbm.ColumnIntflgIndex = df.CCI(&vendorLargeData, "INTFLG_INDEX", intflgIndexSqlName, "", "", "Integer.class", "intflgIndex", "", false, false,true, "INT", 10, 0, "",false,"","", "","","",false,"int64")
	numericIntegerIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo NUMERIC_INTEGER_INDEX
	numericIntegerIndexSqlName.ColumnSqlName = "NUMERIC_INTEGER_INDEX"
	numericIntegerIndexSqlName.IrregularChar = false
	VendorLargeDataDbm.ColumnNumericIntegerIndex = df.CCI(&vendorLargeData, "NUMERIC_INTEGER_INDEX", numericIntegerIndexSqlName, "", "", "Integer.class", "numericIntegerIndex", "", false, false,true, "DECIMAL", 8, 0, "",false,"","", "","","",false,"df.Numeric")
	numericIntegerNoIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo NUMERIC_INTEGER_NO_INDEX
	numericIntegerNoIndexSqlName.ColumnSqlName = "NUMERIC_INTEGER_NO_INDEX"
	numericIntegerNoIndexSqlName.IrregularChar = false
	VendorLargeDataDbm.ColumnNumericIntegerNoIndex = df.CCI(&vendorLargeData, "NUMERIC_INTEGER_NO_INDEX", numericIntegerNoIndexSqlName, "", "", "Integer.class", "numericIntegerNoIndex", "", false, false,true, "DECIMAL", 8, 0, "",false,"","", "","","",false,"df.Numeric")

	VendorLargeDataDbm.ColumnInfoList = new(df.List)
	VendorLargeDataDbm.ColumnInfoList.Add(VendorLargeDataDbm.ColumnLargeDataId)
	VendorLargeDataDbm.ColumnInfoList.Add(VendorLargeDataDbm.ColumnStringIndex)
	VendorLargeDataDbm.ColumnInfoList.Add(VendorLargeDataDbm.ColumnStringNoIndex)
	VendorLargeDataDbm.ColumnInfoList.Add(VendorLargeDataDbm.ColumnStringUniqueIndex)
	VendorLargeDataDbm.ColumnInfoList.Add(VendorLargeDataDbm.ColumnIntflgIndex)
	VendorLargeDataDbm.ColumnInfoList.Add(VendorLargeDataDbm.ColumnNumericIntegerIndex)
	VendorLargeDataDbm.ColumnInfoList.Add(VendorLargeDataDbm.ColumnNumericIntegerNoIndex)


	VendorLargeDataDbm.ColumnInfoMap=make(map[string]int)
	VendorLargeDataDbm.ColumnInfoMap["largeDataId"]=0
		VendorLargeDataDbm.ColumnInfoMap["stringIndex"]=1
		VendorLargeDataDbm.ColumnInfoMap["stringNoIndex"]=2
		VendorLargeDataDbm.ColumnInfoMap["stringUniqueIndex"]=3
		VendorLargeDataDbm.ColumnInfoMap["intflgIndex"]=4
		VendorLargeDataDbm.ColumnInfoMap["numericIntegerIndex"]=5
		VendorLargeDataDbm.ColumnInfoMap["numericIntegerNoIndex"]=6
	    VendorLargeDataDbm.PrimaryKey = true
    VendorLargeDataDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &vendorLargeData
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(VendorLargeDataDbm.ColumnLargeDataId)

	VendorLargeDataDbm.PrimaryInfo = new(df.PrimaryInfo)
	VendorLargeDataDbm.PrimaryInfo.UniqueInfo = ui
	
	var vendorLargeDataMeta df.DBMeta = VendorLargeDataDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["VendorLargeData"] = &vendorLargeDataMeta
}

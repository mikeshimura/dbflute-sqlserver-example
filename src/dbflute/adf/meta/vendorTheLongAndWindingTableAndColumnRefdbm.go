package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorTheLongAndWindingTableAndColumnRefDbm_T struct {
	df.BaseDBMeta
	ColumnTheLongAndWindingTableAndColumnRefId *df.ColumnInfo
	ColumnTheLongAndWindingTableAndColumnId *df.ColumnInfo
	ColumnTheLongAndWindingTableAndColumnRefDate *df.ColumnInfo
	ColumnShortDate *df.ColumnInfo
}

func (b *VendorTheLongAndWindingTableAndColumnRefDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *VendorTheLongAndWindingTableAndColumnRefDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var VendorTheLongAndWindingTableAndColumnRefDbm *VendorTheLongAndWindingTableAndColumnRefDbm_T

func Create_VendorTheLongAndWindingTableAndColumnRefDbm() {
	VendorTheLongAndWindingTableAndColumnRefDbm = new(VendorTheLongAndWindingTableAndColumnRefDbm_T)
	VendorTheLongAndWindingTableAndColumnRefDbm.TableDbName = "vendor_the_long_and_winding_table_and_column_ref"
	VendorTheLongAndWindingTableAndColumnRefDbm.TableDispName = "vendor_the_long_and_winding_table_and_column_ref"
	VendorTheLongAndWindingTableAndColumnRefDbm.TablePropertyName = "vendorTheLongAndWindingTableAndColumnRef"
	VendorTheLongAndWindingTableAndColumnRefDbm.TableSqlName = new(df.TableSqlName)
	VendorTheLongAndWindingTableAndColumnRefDbm.TableSqlName.TableSqlName = "vendor_the_long_and_winding_table_and_column_ref"
	VendorTheLongAndWindingTableAndColumnRefDbm.TableSqlName.CorrespondingDbName = VendorTheLongAndWindingTableAndColumnRefDbm.TableDbName

	var vendorTheLongAndWindingTableAndColumnRef df.DBMeta
	vendorTheLongAndWindingTableAndColumnRef = VendorTheLongAndWindingTableAndColumnRefDbm
	VendorTheLongAndWindingTableAndColumnRefDbm.DBMeta=&vendorTheLongAndWindingTableAndColumnRef
	theLongAndWindingTableAndColumnRefIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo THE_LONG_AND_WINDING_TABLE_AND_COLUMN_REF_ID
	theLongAndWindingTableAndColumnRefIdSqlName.ColumnSqlName = "THE_LONG_AND_WINDING_TABLE_AND_COLUMN_REF_ID"
	theLongAndWindingTableAndColumnRefIdSqlName.IrregularChar = false
	VendorTheLongAndWindingTableAndColumnRefDbm.ColumnTheLongAndWindingTableAndColumnRefId = df.CCI(&vendorTheLongAndWindingTableAndColumnRef, "THE_LONG_AND_WINDING_TABLE_AND_COLUMN_REF_ID", theLongAndWindingTableAndColumnRefIdSqlName, "", "", "Long.class", "theLongAndWindingTableAndColumnRefId", "", true, false,true, "BIGINT", 19, 0, "",false,"","", "","","",false,"int64")
	theLongAndWindingTableAndColumnIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo THE_LONG_AND_WINDING_TABLE_AND_COLUMN_ID
	theLongAndWindingTableAndColumnIdSqlName.ColumnSqlName = "THE_LONG_AND_WINDING_TABLE_AND_COLUMN_ID"
	theLongAndWindingTableAndColumnIdSqlName.IrregularChar = false
	VendorTheLongAndWindingTableAndColumnRefDbm.ColumnTheLongAndWindingTableAndColumnId = df.CCI(&vendorTheLongAndWindingTableAndColumnRef, "THE_LONG_AND_WINDING_TABLE_AND_COLUMN_ID", theLongAndWindingTableAndColumnIdSqlName, "", "", "Long.class", "theLongAndWindingTableAndColumnId", "", false, false,true, "BIGINT", 19, 0, "",false,"","", "vendorTheLongAndWindingTableAndColumn","","",false,"int64")
	theLongAndWindingTableAndColumnRefDateSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo THE_LONG_AND_WINDING_TABLE_AND_COLUMN_REF_DATE
	theLongAndWindingTableAndColumnRefDateSqlName.ColumnSqlName = "THE_LONG_AND_WINDING_TABLE_AND_COLUMN_REF_DATE"
	theLongAndWindingTableAndColumnRefDateSqlName.IrregularChar = false
	VendorTheLongAndWindingTableAndColumnRefDbm.ColumnTheLongAndWindingTableAndColumnRefDate = df.CCI(&vendorTheLongAndWindingTableAndColumnRef, "THE_LONG_AND_WINDING_TABLE_AND_COLUMN_REF_DATE", theLongAndWindingTableAndColumnRefDateSqlName, "", "", "java.time.LocalDate.class", "theLongAndWindingTableAndColumnRefDate", "", false, false,true, "DATE", 10, 0, "",false,"","", "","","",false,"df.MysqlDate")
	shortDateSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo SHORT_DATE
	shortDateSqlName.ColumnSqlName = "SHORT_DATE"
	shortDateSqlName.IrregularChar = false
	VendorTheLongAndWindingTableAndColumnRefDbm.ColumnShortDate = df.CCI(&vendorTheLongAndWindingTableAndColumnRef, "SHORT_DATE", shortDateSqlName, "", "", "java.time.LocalDate.class", "shortDate", "", false, false,true, "DATE", 10, 0, "",false,"","", "","","",false,"df.MysqlDate")

	VendorTheLongAndWindingTableAndColumnRefDbm.ColumnInfoList = new(df.List)
	VendorTheLongAndWindingTableAndColumnRefDbm.ColumnInfoList.Add(VendorTheLongAndWindingTableAndColumnRefDbm.ColumnTheLongAndWindingTableAndColumnRefId)
	VendorTheLongAndWindingTableAndColumnRefDbm.ColumnInfoList.Add(VendorTheLongAndWindingTableAndColumnRefDbm.ColumnTheLongAndWindingTableAndColumnId)
	VendorTheLongAndWindingTableAndColumnRefDbm.ColumnInfoList.Add(VendorTheLongAndWindingTableAndColumnRefDbm.ColumnTheLongAndWindingTableAndColumnRefDate)
	VendorTheLongAndWindingTableAndColumnRefDbm.ColumnInfoList.Add(VendorTheLongAndWindingTableAndColumnRefDbm.ColumnShortDate)


	VendorTheLongAndWindingTableAndColumnRefDbm.ColumnInfoMap=make(map[string]int)
	VendorTheLongAndWindingTableAndColumnRefDbm.ColumnInfoMap["theLongAndWindingTableAndColumnRefId"]=0
		VendorTheLongAndWindingTableAndColumnRefDbm.ColumnInfoMap["theLongAndWindingTableAndColumnId"]=1
		VendorTheLongAndWindingTableAndColumnRefDbm.ColumnInfoMap["theLongAndWindingTableAndColumnRefDate"]=2
		VendorTheLongAndWindingTableAndColumnRefDbm.ColumnInfoMap["shortDate"]=3
	    VendorTheLongAndWindingTableAndColumnRefDbm.PrimaryKey = true
    VendorTheLongAndWindingTableAndColumnRefDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &vendorTheLongAndWindingTableAndColumnRef
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(VendorTheLongAndWindingTableAndColumnRefDbm.ColumnTheLongAndWindingTableAndColumnRefId)

	VendorTheLongAndWindingTableAndColumnRefDbm.PrimaryInfo = new(df.PrimaryInfo)
	VendorTheLongAndWindingTableAndColumnRefDbm.PrimaryInfo.UniqueInfo = ui
	
	var vendorTheLongAndWindingTableAndColumnRefMeta df.DBMeta = VendorTheLongAndWindingTableAndColumnRefDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["VendorTheLongAndWindingTableAndColumnRef"] = &vendorTheLongAndWindingTableAndColumnRefMeta
}

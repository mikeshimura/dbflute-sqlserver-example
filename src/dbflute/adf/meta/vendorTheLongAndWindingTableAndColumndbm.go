package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorTheLongAndWindingTableAndColumnDbm_T struct {
	df.BaseDBMeta
	ColumnTheLongAndWindingTableAndColumnId *df.ColumnInfo
	ColumnTheLongAndWindingTableAndColumnName *df.ColumnInfo
	ColumnShortName *df.ColumnInfo
	ColumnShortSize *df.ColumnInfo
}

func (b *VendorTheLongAndWindingTableAndColumnDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *VendorTheLongAndWindingTableAndColumnDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var VendorTheLongAndWindingTableAndColumnDbm *VendorTheLongAndWindingTableAndColumnDbm_T

func Create_VendorTheLongAndWindingTableAndColumnDbm() {
	VendorTheLongAndWindingTableAndColumnDbm = new(VendorTheLongAndWindingTableAndColumnDbm_T)
	VendorTheLongAndWindingTableAndColumnDbm.TableDbName = "vendor_the_long_and_winding_table_and_column"
	VendorTheLongAndWindingTableAndColumnDbm.TableDispName = "vendor_the_long_and_winding_table_and_column"
	VendorTheLongAndWindingTableAndColumnDbm.TablePropertyName = "vendorTheLongAndWindingTableAndColumn"
	VendorTheLongAndWindingTableAndColumnDbm.TableSqlName = new(df.TableSqlName)
	VendorTheLongAndWindingTableAndColumnDbm.TableSqlName.TableSqlName = "vendor_the_long_and_winding_table_and_column"
	VendorTheLongAndWindingTableAndColumnDbm.TableSqlName.CorrespondingDbName = VendorTheLongAndWindingTableAndColumnDbm.TableDbName

	var vendorTheLongAndWindingTableAndColumn df.DBMeta
	vendorTheLongAndWindingTableAndColumn = VendorTheLongAndWindingTableAndColumnDbm
	VendorTheLongAndWindingTableAndColumnDbm.DBMeta=&vendorTheLongAndWindingTableAndColumn
	theLongAndWindingTableAndColumnIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo THE_LONG_AND_WINDING_TABLE_AND_COLUMN_ID
	theLongAndWindingTableAndColumnIdSqlName.ColumnSqlName = "THE_LONG_AND_WINDING_TABLE_AND_COLUMN_ID"
	theLongAndWindingTableAndColumnIdSqlName.IrregularChar = false
	VendorTheLongAndWindingTableAndColumnDbm.ColumnTheLongAndWindingTableAndColumnId = df.CCI(&vendorTheLongAndWindingTableAndColumn, "THE_LONG_AND_WINDING_TABLE_AND_COLUMN_ID", theLongAndWindingTableAndColumnIdSqlName, "", "", "Long.class", "theLongAndWindingTableAndColumnId", "", true, false,true, "BIGINT", 19, 0, "",false,"","", "","vendorTheLongAndWindingTableAndColumnRefList","",false,"int64")
	theLongAndWindingTableAndColumnNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo THE_LONG_AND_WINDING_TABLE_AND_COLUMN_NAME
	theLongAndWindingTableAndColumnNameSqlName.ColumnSqlName = "THE_LONG_AND_WINDING_TABLE_AND_COLUMN_NAME"
	theLongAndWindingTableAndColumnNameSqlName.IrregularChar = false
	VendorTheLongAndWindingTableAndColumnDbm.ColumnTheLongAndWindingTableAndColumnName = df.CCI(&vendorTheLongAndWindingTableAndColumn, "THE_LONG_AND_WINDING_TABLE_AND_COLUMN_NAME", theLongAndWindingTableAndColumnNameSqlName, "", "", "String.class", "theLongAndWindingTableAndColumnName", "", false, false,true, "VARCHAR", 180, 0, "",false,"","", "","","",false,"string")
	shortNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo SHORT_NAME
	shortNameSqlName.ColumnSqlName = "SHORT_NAME"
	shortNameSqlName.IrregularChar = false
	VendorTheLongAndWindingTableAndColumnDbm.ColumnShortName = df.CCI(&vendorTheLongAndWindingTableAndColumn, "SHORT_NAME", shortNameSqlName, "", "", "String.class", "shortName", "", false, false,true, "VARCHAR", 200, 0, "",false,"","", "","","",false,"string")
	shortSizeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo SHORT_SIZE
	shortSizeSqlName.ColumnSqlName = "SHORT_SIZE"
	shortSizeSqlName.IrregularChar = false
	VendorTheLongAndWindingTableAndColumnDbm.ColumnShortSize = df.CCI(&vendorTheLongAndWindingTableAndColumn, "SHORT_SIZE", shortSizeSqlName, "", "", "Integer.class", "shortSize", "", false, false,true, "INT", 10, 0, "",false,"","", "","","",false,"int64")

	VendorTheLongAndWindingTableAndColumnDbm.ColumnInfoList = new(df.List)
	VendorTheLongAndWindingTableAndColumnDbm.ColumnInfoList.Add(VendorTheLongAndWindingTableAndColumnDbm.ColumnTheLongAndWindingTableAndColumnId)
	VendorTheLongAndWindingTableAndColumnDbm.ColumnInfoList.Add(VendorTheLongAndWindingTableAndColumnDbm.ColumnTheLongAndWindingTableAndColumnName)
	VendorTheLongAndWindingTableAndColumnDbm.ColumnInfoList.Add(VendorTheLongAndWindingTableAndColumnDbm.ColumnShortName)
	VendorTheLongAndWindingTableAndColumnDbm.ColumnInfoList.Add(VendorTheLongAndWindingTableAndColumnDbm.ColumnShortSize)


	VendorTheLongAndWindingTableAndColumnDbm.ColumnInfoMap=make(map[string]int)
	VendorTheLongAndWindingTableAndColumnDbm.ColumnInfoMap["theLongAndWindingTableAndColumnId"]=0
		VendorTheLongAndWindingTableAndColumnDbm.ColumnInfoMap["theLongAndWindingTableAndColumnName"]=1
		VendorTheLongAndWindingTableAndColumnDbm.ColumnInfoMap["shortName"]=2
		VendorTheLongAndWindingTableAndColumnDbm.ColumnInfoMap["shortSize"]=3
	    VendorTheLongAndWindingTableAndColumnDbm.PrimaryKey = true
    VendorTheLongAndWindingTableAndColumnDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &vendorTheLongAndWindingTableAndColumn
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(VendorTheLongAndWindingTableAndColumnDbm.ColumnTheLongAndWindingTableAndColumnId)

	VendorTheLongAndWindingTableAndColumnDbm.PrimaryInfo = new(df.PrimaryInfo)
	VendorTheLongAndWindingTableAndColumnDbm.PrimaryInfo.UniqueInfo = ui
	
	var vendorTheLongAndWindingTableAndColumnMeta df.DBMeta = VendorTheLongAndWindingTableAndColumnDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["VendorTheLongAndWindingTableAndColumn"] = &vendorTheLongAndWindingTableAndColumnMeta
}

package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorLargeDataRefDbm_T struct {
	df.BaseDBMeta
	ColumnLargeDataRefId *df.ColumnInfo
	ColumnLargeDataId *df.ColumnInfo
	ColumnDateIndex *df.ColumnInfo
	ColumnDateNoIndex *df.ColumnInfo
	ColumnTimestampIndex *df.ColumnInfo
	ColumnTimestampNoIndex *df.ColumnInfo
	ColumnNullableDecimalIndex *df.ColumnInfo
	ColumnNullableDecimalNoIndex *df.ColumnInfo
	ColumnSelfParentId *df.ColumnInfo
}

func (b *VendorLargeDataRefDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *VendorLargeDataRefDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var VendorLargeDataRefDbm *VendorLargeDataRefDbm_T

func Create_VendorLargeDataRefDbm() {
	VendorLargeDataRefDbm = new(VendorLargeDataRefDbm_T)
	VendorLargeDataRefDbm.TableDbName = "vendor_large_data_ref"
	VendorLargeDataRefDbm.TableDispName = "vendor_large_data_ref"
	VendorLargeDataRefDbm.TablePropertyName = "vendorLargeDataRef"
	VendorLargeDataRefDbm.TableSqlName = new(df.TableSqlName)
	VendorLargeDataRefDbm.TableSqlName.TableSqlName = "vendor_large_data_ref"
	VendorLargeDataRefDbm.TableSqlName.CorrespondingDbName = VendorLargeDataRefDbm.TableDbName

	var vendorLargeDataRef df.DBMeta
	vendorLargeDataRef = VendorLargeDataRefDbm
	VendorLargeDataRefDbm.DBMeta=&vendorLargeDataRef
	largeDataRefIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo LARGE_DATA_REF_ID
	largeDataRefIdSqlName.ColumnSqlName = "LARGE_DATA_REF_ID"
	largeDataRefIdSqlName.IrregularChar = false
	VendorLargeDataRefDbm.ColumnLargeDataRefId = df.CCI(&vendorLargeDataRef, "LARGE_DATA_REF_ID", largeDataRefIdSqlName, "", "", "Long.class", "largeDataRefId", "", true, false,true, "BIGINT", 19, 0, "",false,"","", "","","",false,"int64")
	largeDataIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo LARGE_DATA_ID
	largeDataIdSqlName.ColumnSqlName = "LARGE_DATA_ID"
	largeDataIdSqlName.IrregularChar = false
	VendorLargeDataRefDbm.ColumnLargeDataId = df.CCI(&vendorLargeDataRef, "LARGE_DATA_ID", largeDataIdSqlName, "", "", "Long.class", "largeDataId", "", false, false,true, "BIGINT", 19, 0, "",false,"","", "","","",false,"int64")
	dateIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo DATE_INDEX
	dateIndexSqlName.ColumnSqlName = "DATE_INDEX"
	dateIndexSqlName.IrregularChar = false
	VendorLargeDataRefDbm.ColumnDateIndex = df.CCI(&vendorLargeDataRef, "DATE_INDEX", dateIndexSqlName, "", "", "java.time.LocalDate.class", "dateIndex", "", false, false,true, "DATE", 10, 0, "",false,"","", "","","",false,"df.MysqlDate")
	dateNoIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo DATE_NO_INDEX
	dateNoIndexSqlName.ColumnSqlName = "DATE_NO_INDEX"
	dateNoIndexSqlName.IrregularChar = false
	VendorLargeDataRefDbm.ColumnDateNoIndex = df.CCI(&vendorLargeDataRef, "DATE_NO_INDEX", dateNoIndexSqlName, "", "", "java.time.LocalDate.class", "dateNoIndex", "", false, false,true, "DATE", 10, 0, "",false,"","", "","","",false,"df.MysqlDate")
	timestampIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo TIMESTAMP_INDEX
	timestampIndexSqlName.ColumnSqlName = "TIMESTAMP_INDEX"
	timestampIndexSqlName.IrregularChar = false
	VendorLargeDataRefDbm.ColumnTimestampIndex = df.CCI(&vendorLargeDataRef, "TIMESTAMP_INDEX", timestampIndexSqlName, "", "", "java.time.LocalDateTime.class", "timestampIndex", "", false, false,true, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlTimestamp")
	timestampNoIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo TIMESTAMP_NO_INDEX
	timestampNoIndexSqlName.ColumnSqlName = "TIMESTAMP_NO_INDEX"
	timestampNoIndexSqlName.IrregularChar = false
	VendorLargeDataRefDbm.ColumnTimestampNoIndex = df.CCI(&vendorLargeDataRef, "TIMESTAMP_NO_INDEX", timestampNoIndexSqlName, "", "", "java.time.LocalDateTime.class", "timestampNoIndex", "", false, false,true, "DATETIME", 19, 0, "",false,"","", "","","",false,"df.MysqlTimestamp")
	nullableDecimalIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo NULLABLE_DECIMAL_INDEX
	nullableDecimalIndexSqlName.ColumnSqlName = "NULLABLE_DECIMAL_INDEX"
	nullableDecimalIndexSqlName.IrregularChar = false
	VendorLargeDataRefDbm.ColumnNullableDecimalIndex = df.CCI(&vendorLargeDataRef, "NULLABLE_DECIMAL_INDEX", nullableDecimalIndexSqlName, "", "", "java.math.BigDecimal.class", "nullableDecimalIndex", "", false, false,false, "DECIMAL", 12, 3, "",false,"","", "","","",false,"df.NullNumeric")
	nullableDecimalNoIndexSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo NULLABLE_DECIMAL_NO_INDEX
	nullableDecimalNoIndexSqlName.ColumnSqlName = "NULLABLE_DECIMAL_NO_INDEX"
	nullableDecimalNoIndexSqlName.IrregularChar = false
	VendorLargeDataRefDbm.ColumnNullableDecimalNoIndex = df.CCI(&vendorLargeDataRef, "NULLABLE_DECIMAL_NO_INDEX", nullableDecimalNoIndexSqlName, "", "", "java.math.BigDecimal.class", "nullableDecimalNoIndex", "", false, false,false, "DECIMAL", 12, 3, "",false,"","", "","","",false,"df.NullNumeric")
	selfParentIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo SELF_PARENT_ID
	selfParentIdSqlName.ColumnSqlName = "SELF_PARENT_ID"
	selfParentIdSqlName.IrregularChar = false
	VendorLargeDataRefDbm.ColumnSelfParentId = df.CCI(&vendorLargeDataRef, "SELF_PARENT_ID", selfParentIdSqlName, "", "", "Long.class", "selfParentId", "", false, false,false, "BIGINT", 19, 0, "",false,"","", "","","",false,"sql.NullInt64")

	VendorLargeDataRefDbm.ColumnInfoList = new(df.List)
	VendorLargeDataRefDbm.ColumnInfoList.Add(VendorLargeDataRefDbm.ColumnLargeDataRefId)
	VendorLargeDataRefDbm.ColumnInfoList.Add(VendorLargeDataRefDbm.ColumnLargeDataId)
	VendorLargeDataRefDbm.ColumnInfoList.Add(VendorLargeDataRefDbm.ColumnDateIndex)
	VendorLargeDataRefDbm.ColumnInfoList.Add(VendorLargeDataRefDbm.ColumnDateNoIndex)
	VendorLargeDataRefDbm.ColumnInfoList.Add(VendorLargeDataRefDbm.ColumnTimestampIndex)
	VendorLargeDataRefDbm.ColumnInfoList.Add(VendorLargeDataRefDbm.ColumnTimestampNoIndex)
	VendorLargeDataRefDbm.ColumnInfoList.Add(VendorLargeDataRefDbm.ColumnNullableDecimalIndex)
	VendorLargeDataRefDbm.ColumnInfoList.Add(VendorLargeDataRefDbm.ColumnNullableDecimalNoIndex)
	VendorLargeDataRefDbm.ColumnInfoList.Add(VendorLargeDataRefDbm.ColumnSelfParentId)


	VendorLargeDataRefDbm.ColumnInfoMap=make(map[string]int)
	VendorLargeDataRefDbm.ColumnInfoMap["largeDataRefId"]=0
		VendorLargeDataRefDbm.ColumnInfoMap["largeDataId"]=1
		VendorLargeDataRefDbm.ColumnInfoMap["dateIndex"]=2
		VendorLargeDataRefDbm.ColumnInfoMap["dateNoIndex"]=3
		VendorLargeDataRefDbm.ColumnInfoMap["timestampIndex"]=4
		VendorLargeDataRefDbm.ColumnInfoMap["timestampNoIndex"]=5
		VendorLargeDataRefDbm.ColumnInfoMap["nullableDecimalIndex"]=6
		VendorLargeDataRefDbm.ColumnInfoMap["nullableDecimalNoIndex"]=7
		VendorLargeDataRefDbm.ColumnInfoMap["selfParentId"]=8
	    VendorLargeDataRefDbm.PrimaryKey = true
    VendorLargeDataRefDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &vendorLargeDataRef
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(VendorLargeDataRefDbm.ColumnLargeDataRefId)

	VendorLargeDataRefDbm.PrimaryInfo = new(df.PrimaryInfo)
	VendorLargeDataRefDbm.PrimaryInfo.UniqueInfo = ui
	
	var vendorLargeDataRefMeta df.DBMeta = VendorLargeDataRefDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["VendorLargeDataRef"] = &vendorLargeDataRefMeta
}

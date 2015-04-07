package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorConstraintNameAutoBarDbm_T struct {
	df.BaseDBMeta
	ColumnConstraintNameAutoBarId *df.ColumnInfo
	ColumnConstraintNameAutoBarName *df.ColumnInfo
}

func (b *VendorConstraintNameAutoBarDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *VendorConstraintNameAutoBarDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var VendorConstraintNameAutoBarDbm *VendorConstraintNameAutoBarDbm_T

func Create_VendorConstraintNameAutoBarDbm() {
	VendorConstraintNameAutoBarDbm = new(VendorConstraintNameAutoBarDbm_T)
	VendorConstraintNameAutoBarDbm.TableDbName = "vendor_constraint_name_auto_bar"
	VendorConstraintNameAutoBarDbm.TableDispName = "vendor_constraint_name_auto_bar"
	VendorConstraintNameAutoBarDbm.TablePropertyName = "vendorConstraintNameAutoBar"
	VendorConstraintNameAutoBarDbm.TableSqlName = new(df.TableSqlName)
	VendorConstraintNameAutoBarDbm.TableSqlName.TableSqlName = "vendor_constraint_name_auto_bar"
	VendorConstraintNameAutoBarDbm.TableSqlName.CorrespondingDbName = VendorConstraintNameAutoBarDbm.TableDbName

	var vendorConstraintNameAutoBar df.DBMeta
	vendorConstraintNameAutoBar = VendorConstraintNameAutoBarDbm
	VendorConstraintNameAutoBarDbm.DBMeta=&vendorConstraintNameAutoBar
	constraintNameAutoBarIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo CONSTRAINT_NAME_AUTO_BAR_ID
	constraintNameAutoBarIdSqlName.ColumnSqlName = "CONSTRAINT_NAME_AUTO_BAR_ID"
	constraintNameAutoBarIdSqlName.IrregularChar = false
	VendorConstraintNameAutoBarDbm.ColumnConstraintNameAutoBarId = df.CCI(&vendorConstraintNameAutoBar, "CONSTRAINT_NAME_AUTO_BAR_ID", constraintNameAutoBarIdSqlName, "", "", "Long.class", "constraintNameAutoBarId", "", true, false,true, "DECIMAL", 16, 0, "",false,"","", "","vendorConstraintNameAutoRefList","",false,"df.Numeric")
	constraintNameAutoBarNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo CONSTRAINT_NAME_AUTO_BAR_NAME
	constraintNameAutoBarNameSqlName.ColumnSqlName = "CONSTRAINT_NAME_AUTO_BAR_NAME"
	constraintNameAutoBarNameSqlName.IrregularChar = false
	VendorConstraintNameAutoBarDbm.ColumnConstraintNameAutoBarName = df.CCI(&vendorConstraintNameAutoBar, "CONSTRAINT_NAME_AUTO_BAR_NAME", constraintNameAutoBarNameSqlName, "", "", "String.class", "constraintNameAutoBarName", "", false, false,true, "VARCHAR", 50, 0, "",false,"","", "","","",false,"string")

	VendorConstraintNameAutoBarDbm.ColumnInfoList = new(df.List)
	VendorConstraintNameAutoBarDbm.ColumnInfoList.Add(VendorConstraintNameAutoBarDbm.ColumnConstraintNameAutoBarId)
	VendorConstraintNameAutoBarDbm.ColumnInfoList.Add(VendorConstraintNameAutoBarDbm.ColumnConstraintNameAutoBarName)


	VendorConstraintNameAutoBarDbm.ColumnInfoMap=make(map[string]int)
	VendorConstraintNameAutoBarDbm.ColumnInfoMap["constraintNameAutoBarId"]=0
		VendorConstraintNameAutoBarDbm.ColumnInfoMap["constraintNameAutoBarName"]=1
	    VendorConstraintNameAutoBarDbm.PrimaryKey = true
    VendorConstraintNameAutoBarDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &vendorConstraintNameAutoBar
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(VendorConstraintNameAutoBarDbm.ColumnConstraintNameAutoBarId)

	VendorConstraintNameAutoBarDbm.PrimaryInfo = new(df.PrimaryInfo)
	VendorConstraintNameAutoBarDbm.PrimaryInfo.UniqueInfo = ui
	
	var vendorConstraintNameAutoBarMeta df.DBMeta = VendorConstraintNameAutoBarDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["VendorConstraintNameAutoBar"] = &vendorConstraintNameAutoBarMeta
}

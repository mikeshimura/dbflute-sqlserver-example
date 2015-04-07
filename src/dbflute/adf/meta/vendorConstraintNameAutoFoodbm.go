package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorConstraintNameAutoFooDbm_T struct {
	df.BaseDBMeta
	ColumnConstraintNameAutoFooId *df.ColumnInfo
	ColumnConstraintNameAutoFooName *df.ColumnInfo
}

func (b *VendorConstraintNameAutoFooDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *VendorConstraintNameAutoFooDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var VendorConstraintNameAutoFooDbm *VendorConstraintNameAutoFooDbm_T

func Create_VendorConstraintNameAutoFooDbm() {
	VendorConstraintNameAutoFooDbm = new(VendorConstraintNameAutoFooDbm_T)
	VendorConstraintNameAutoFooDbm.TableDbName = "vendor_constraint_name_auto_foo"
	VendorConstraintNameAutoFooDbm.TableDispName = "vendor_constraint_name_auto_foo"
	VendorConstraintNameAutoFooDbm.TablePropertyName = "vendorConstraintNameAutoFoo"
	VendorConstraintNameAutoFooDbm.TableSqlName = new(df.TableSqlName)
	VendorConstraintNameAutoFooDbm.TableSqlName.TableSqlName = "vendor_constraint_name_auto_foo"
	VendorConstraintNameAutoFooDbm.TableSqlName.CorrespondingDbName = VendorConstraintNameAutoFooDbm.TableDbName

	var vendorConstraintNameAutoFoo df.DBMeta
	vendorConstraintNameAutoFoo = VendorConstraintNameAutoFooDbm
	VendorConstraintNameAutoFooDbm.DBMeta=&vendorConstraintNameAutoFoo
	constraintNameAutoFooIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo CONSTRAINT_NAME_AUTO_FOO_ID
	constraintNameAutoFooIdSqlName.ColumnSqlName = "CONSTRAINT_NAME_AUTO_FOO_ID"
	constraintNameAutoFooIdSqlName.IrregularChar = false
	VendorConstraintNameAutoFooDbm.ColumnConstraintNameAutoFooId = df.CCI(&vendorConstraintNameAutoFoo, "CONSTRAINT_NAME_AUTO_FOO_ID", constraintNameAutoFooIdSqlName, "", "", "Long.class", "constraintNameAutoFooId", "", true, false,true, "DECIMAL", 16, 0, "",false,"","", "","vendorConstraintNameAutoRefList","",false,"df.Numeric")
	constraintNameAutoFooNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo CONSTRAINT_NAME_AUTO_FOO_NAME
	constraintNameAutoFooNameSqlName.ColumnSqlName = "CONSTRAINT_NAME_AUTO_FOO_NAME"
	constraintNameAutoFooNameSqlName.IrregularChar = false
	VendorConstraintNameAutoFooDbm.ColumnConstraintNameAutoFooName = df.CCI(&vendorConstraintNameAutoFoo, "CONSTRAINT_NAME_AUTO_FOO_NAME", constraintNameAutoFooNameSqlName, "", "", "String.class", "constraintNameAutoFooName", "", false, false,true, "VARCHAR", 50, 0, "",false,"","", "","","",false,"string")

	VendorConstraintNameAutoFooDbm.ColumnInfoList = new(df.List)
	VendorConstraintNameAutoFooDbm.ColumnInfoList.Add(VendorConstraintNameAutoFooDbm.ColumnConstraintNameAutoFooId)
	VendorConstraintNameAutoFooDbm.ColumnInfoList.Add(VendorConstraintNameAutoFooDbm.ColumnConstraintNameAutoFooName)


	VendorConstraintNameAutoFooDbm.ColumnInfoMap=make(map[string]int)
	VendorConstraintNameAutoFooDbm.ColumnInfoMap["constraintNameAutoFooId"]=0
		VendorConstraintNameAutoFooDbm.ColumnInfoMap["constraintNameAutoFooName"]=1
	    VendorConstraintNameAutoFooDbm.PrimaryKey = true
    VendorConstraintNameAutoFooDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &vendorConstraintNameAutoFoo
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(VendorConstraintNameAutoFooDbm.ColumnConstraintNameAutoFooId)

	VendorConstraintNameAutoFooDbm.PrimaryInfo = new(df.PrimaryInfo)
	VendorConstraintNameAutoFooDbm.PrimaryInfo.UniqueInfo = ui
	
	var vendorConstraintNameAutoFooMeta df.DBMeta = VendorConstraintNameAutoFooDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["VendorConstraintNameAutoFoo"] = &vendorConstraintNameAutoFooMeta
}

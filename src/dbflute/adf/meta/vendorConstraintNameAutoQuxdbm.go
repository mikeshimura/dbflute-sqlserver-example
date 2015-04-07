package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorConstraintNameAutoQuxDbm_T struct {
	df.BaseDBMeta
	ColumnConstraintNameAutoQuxId *df.ColumnInfo
	ColumnConstraintNameAutoQuxName *df.ColumnInfo
}

func (b *VendorConstraintNameAutoQuxDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *VendorConstraintNameAutoQuxDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var VendorConstraintNameAutoQuxDbm *VendorConstraintNameAutoQuxDbm_T

func Create_VendorConstraintNameAutoQuxDbm() {
	VendorConstraintNameAutoQuxDbm = new(VendorConstraintNameAutoQuxDbm_T)
	VendorConstraintNameAutoQuxDbm.TableDbName = "vendor_constraint_name_auto_qux"
	VendorConstraintNameAutoQuxDbm.TableDispName = "vendor_constraint_name_auto_qux"
	VendorConstraintNameAutoQuxDbm.TablePropertyName = "vendorConstraintNameAutoQux"
	VendorConstraintNameAutoQuxDbm.TableSqlName = new(df.TableSqlName)
	VendorConstraintNameAutoQuxDbm.TableSqlName.TableSqlName = "vendor_constraint_name_auto_qux"
	VendorConstraintNameAutoQuxDbm.TableSqlName.CorrespondingDbName = VendorConstraintNameAutoQuxDbm.TableDbName

	var vendorConstraintNameAutoQux df.DBMeta
	vendorConstraintNameAutoQux = VendorConstraintNameAutoQuxDbm
	VendorConstraintNameAutoQuxDbm.DBMeta=&vendorConstraintNameAutoQux
	constraintNameAutoQuxIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo CONSTRAINT_NAME_AUTO_QUX_ID
	constraintNameAutoQuxIdSqlName.ColumnSqlName = "CONSTRAINT_NAME_AUTO_QUX_ID"
	constraintNameAutoQuxIdSqlName.IrregularChar = false
	VendorConstraintNameAutoQuxDbm.ColumnConstraintNameAutoQuxId = df.CCI(&vendorConstraintNameAutoQux, "CONSTRAINT_NAME_AUTO_QUX_ID", constraintNameAutoQuxIdSqlName, "", "", "Long.class", "constraintNameAutoQuxId", "", true, false,true, "DECIMAL", 16, 0, "",false,"","", "","vendorConstraintNameAutoRefList","",false,"df.Numeric")
	constraintNameAutoQuxNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo CONSTRAINT_NAME_AUTO_QUX_NAME
	constraintNameAutoQuxNameSqlName.ColumnSqlName = "CONSTRAINT_NAME_AUTO_QUX_NAME"
	constraintNameAutoQuxNameSqlName.IrregularChar = false
	VendorConstraintNameAutoQuxDbm.ColumnConstraintNameAutoQuxName = df.CCI(&vendorConstraintNameAutoQux, "CONSTRAINT_NAME_AUTO_QUX_NAME", constraintNameAutoQuxNameSqlName, "", "", "String.class", "constraintNameAutoQuxName", "", false, false,true, "VARCHAR", 50, 0, "",false,"","", "","","",false,"string")

	VendorConstraintNameAutoQuxDbm.ColumnInfoList = new(df.List)
	VendorConstraintNameAutoQuxDbm.ColumnInfoList.Add(VendorConstraintNameAutoQuxDbm.ColumnConstraintNameAutoQuxId)
	VendorConstraintNameAutoQuxDbm.ColumnInfoList.Add(VendorConstraintNameAutoQuxDbm.ColumnConstraintNameAutoQuxName)


	VendorConstraintNameAutoQuxDbm.ColumnInfoMap=make(map[string]int)
	VendorConstraintNameAutoQuxDbm.ColumnInfoMap["constraintNameAutoQuxId"]=0
		VendorConstraintNameAutoQuxDbm.ColumnInfoMap["constraintNameAutoQuxName"]=1
	    VendorConstraintNameAutoQuxDbm.PrimaryKey = true
    VendorConstraintNameAutoQuxDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &vendorConstraintNameAutoQux
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(VendorConstraintNameAutoQuxDbm.ColumnConstraintNameAutoQuxId)

	VendorConstraintNameAutoQuxDbm.PrimaryInfo = new(df.PrimaryInfo)
	VendorConstraintNameAutoQuxDbm.PrimaryInfo.UniqueInfo = ui
	
	var vendorConstraintNameAutoQuxMeta df.DBMeta = VendorConstraintNameAutoQuxDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["VendorConstraintNameAutoQux"] = &vendorConstraintNameAutoQuxMeta
}

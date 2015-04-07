package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorConstraintNameAutoRefDbm_T struct {
	df.BaseDBMeta
	ColumnConstraintNameAutoRefId *df.ColumnInfo
	ColumnConstraintNameAutoFooId *df.ColumnInfo
	ColumnConstraintNameAutoBarId *df.ColumnInfo
	ColumnConstraintNameAutoQuxId *df.ColumnInfo
	ColumnConstraintNameAutoCorgeId *df.ColumnInfo
	ColumnConstraintNameAutoUnique *df.ColumnInfo
}

func (b *VendorConstraintNameAutoRefDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *VendorConstraintNameAutoRefDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var VendorConstraintNameAutoRefDbm *VendorConstraintNameAutoRefDbm_T

func Create_VendorConstraintNameAutoRefDbm() {
	VendorConstraintNameAutoRefDbm = new(VendorConstraintNameAutoRefDbm_T)
	VendorConstraintNameAutoRefDbm.TableDbName = "vendor_constraint_name_auto_ref"
	VendorConstraintNameAutoRefDbm.TableDispName = "vendor_constraint_name_auto_ref"
	VendorConstraintNameAutoRefDbm.TablePropertyName = "vendorConstraintNameAutoRef"
	VendorConstraintNameAutoRefDbm.TableSqlName = new(df.TableSqlName)
	VendorConstraintNameAutoRefDbm.TableSqlName.TableSqlName = "vendor_constraint_name_auto_ref"
	VendorConstraintNameAutoRefDbm.TableSqlName.CorrespondingDbName = VendorConstraintNameAutoRefDbm.TableDbName

	var vendorConstraintNameAutoRef df.DBMeta
	vendorConstraintNameAutoRef = VendorConstraintNameAutoRefDbm
	VendorConstraintNameAutoRefDbm.DBMeta=&vendorConstraintNameAutoRef
	constraintNameAutoRefIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo CONSTRAINT_NAME_AUTO_REF_ID
	constraintNameAutoRefIdSqlName.ColumnSqlName = "CONSTRAINT_NAME_AUTO_REF_ID"
	constraintNameAutoRefIdSqlName.IrregularChar = false
	VendorConstraintNameAutoRefDbm.ColumnConstraintNameAutoRefId = df.CCI(&vendorConstraintNameAutoRef, "CONSTRAINT_NAME_AUTO_REF_ID", constraintNameAutoRefIdSqlName, "", "", "Long.class", "constraintNameAutoRefId", "", true, false,true, "DECIMAL", 16, 0, "",false,"","", "","","",false,"df.Numeric")
	constraintNameAutoFooIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo CONSTRAINT_NAME_AUTO_FOO_ID
	constraintNameAutoFooIdSqlName.ColumnSqlName = "CONSTRAINT_NAME_AUTO_FOO_ID"
	constraintNameAutoFooIdSqlName.IrregularChar = false
	VendorConstraintNameAutoRefDbm.ColumnConstraintNameAutoFooId = df.CCI(&vendorConstraintNameAutoRef, "CONSTRAINT_NAME_AUTO_FOO_ID", constraintNameAutoFooIdSqlName, "", "", "Long.class", "constraintNameAutoFooId", "", false, false,true, "DECIMAL", 16, 0, "",false,"","", "vendorConstraintNameAutoFoo","","",false,"df.Numeric")
	constraintNameAutoBarIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo CONSTRAINT_NAME_AUTO_BAR_ID
	constraintNameAutoBarIdSqlName.ColumnSqlName = "CONSTRAINT_NAME_AUTO_BAR_ID"
	constraintNameAutoBarIdSqlName.IrregularChar = false
	VendorConstraintNameAutoRefDbm.ColumnConstraintNameAutoBarId = df.CCI(&vendorConstraintNameAutoRef, "CONSTRAINT_NAME_AUTO_BAR_ID", constraintNameAutoBarIdSqlName, "", "", "Long.class", "constraintNameAutoBarId", "", false, false,true, "DECIMAL", 16, 0, "",false,"","", "vendorConstraintNameAutoBar","","",false,"df.Numeric")
	constraintNameAutoQuxIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo CONSTRAINT_NAME_AUTO_QUX_ID
	constraintNameAutoQuxIdSqlName.ColumnSqlName = "CONSTRAINT_NAME_AUTO_QUX_ID"
	constraintNameAutoQuxIdSqlName.IrregularChar = false
	VendorConstraintNameAutoRefDbm.ColumnConstraintNameAutoQuxId = df.CCI(&vendorConstraintNameAutoRef, "CONSTRAINT_NAME_AUTO_QUX_ID", constraintNameAutoQuxIdSqlName, "", "", "Long.class", "constraintNameAutoQuxId", "", false, false,true, "DECIMAL", 16, 0, "",false,"","", "vendorConstraintNameAutoQux","","",false,"df.Numeric")
	constraintNameAutoCorgeIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo CONSTRAINT_NAME_AUTO_CORGE_ID
	constraintNameAutoCorgeIdSqlName.ColumnSqlName = "CONSTRAINT_NAME_AUTO_CORGE_ID"
	constraintNameAutoCorgeIdSqlName.IrregularChar = false
	VendorConstraintNameAutoRefDbm.ColumnConstraintNameAutoCorgeId = df.CCI(&vendorConstraintNameAutoRef, "CONSTRAINT_NAME_AUTO_CORGE_ID", constraintNameAutoCorgeIdSqlName, "", "", "Long.class", "constraintNameAutoCorgeId", "", false, false,true, "DECIMAL", 16, 0, "",false,"","", "","","",false,"df.Numeric")
	constraintNameAutoUniqueSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo CONSTRAINT_NAME_AUTO_UNIQUE
	constraintNameAutoUniqueSqlName.ColumnSqlName = "CONSTRAINT_NAME_AUTO_UNIQUE"
	constraintNameAutoUniqueSqlName.IrregularChar = false
	VendorConstraintNameAutoRefDbm.ColumnConstraintNameAutoUnique = df.CCI(&vendorConstraintNameAutoRef, "CONSTRAINT_NAME_AUTO_UNIQUE", constraintNameAutoUniqueSqlName, "", "", "String.class", "constraintNameAutoUnique", "", false, false,true, "VARCHAR", 50, 0, "",false,"","", "","","",false,"string")

	VendorConstraintNameAutoRefDbm.ColumnInfoList = new(df.List)
	VendorConstraintNameAutoRefDbm.ColumnInfoList.Add(VendorConstraintNameAutoRefDbm.ColumnConstraintNameAutoRefId)
	VendorConstraintNameAutoRefDbm.ColumnInfoList.Add(VendorConstraintNameAutoRefDbm.ColumnConstraintNameAutoFooId)
	VendorConstraintNameAutoRefDbm.ColumnInfoList.Add(VendorConstraintNameAutoRefDbm.ColumnConstraintNameAutoBarId)
	VendorConstraintNameAutoRefDbm.ColumnInfoList.Add(VendorConstraintNameAutoRefDbm.ColumnConstraintNameAutoQuxId)
	VendorConstraintNameAutoRefDbm.ColumnInfoList.Add(VendorConstraintNameAutoRefDbm.ColumnConstraintNameAutoCorgeId)
	VendorConstraintNameAutoRefDbm.ColumnInfoList.Add(VendorConstraintNameAutoRefDbm.ColumnConstraintNameAutoUnique)


	VendorConstraintNameAutoRefDbm.ColumnInfoMap=make(map[string]int)
	VendorConstraintNameAutoRefDbm.ColumnInfoMap["constraintNameAutoRefId"]=0
		VendorConstraintNameAutoRefDbm.ColumnInfoMap["constraintNameAutoFooId"]=1
		VendorConstraintNameAutoRefDbm.ColumnInfoMap["constraintNameAutoBarId"]=2
		VendorConstraintNameAutoRefDbm.ColumnInfoMap["constraintNameAutoQuxId"]=3
		VendorConstraintNameAutoRefDbm.ColumnInfoMap["constraintNameAutoCorgeId"]=4
		VendorConstraintNameAutoRefDbm.ColumnInfoMap["constraintNameAutoUnique"]=5
	    VendorConstraintNameAutoRefDbm.PrimaryKey = true
    VendorConstraintNameAutoRefDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &vendorConstraintNameAutoRef
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(VendorConstraintNameAutoRefDbm.ColumnConstraintNameAutoRefId)

	VendorConstraintNameAutoRefDbm.PrimaryInfo = new(df.PrimaryInfo)
	VendorConstraintNameAutoRefDbm.PrimaryInfo.UniqueInfo = ui
	
	var vendorConstraintNameAutoRefMeta df.DBMeta = VendorConstraintNameAutoRefDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["VendorConstraintNameAutoRef"] = &vendorConstraintNameAutoRefMeta
}

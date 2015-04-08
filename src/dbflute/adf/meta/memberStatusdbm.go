package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberStatusDbm_T struct {
	df.BaseDBMeta
	ColumnMemberStatusCode *df.ColumnInfo
	ColumnMemberStatusName *df.ColumnInfo
	ColumnDescription *df.ColumnInfo
	ColumnDisplayOrder *df.ColumnInfo
}

func (b *MemberStatusDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *MemberStatusDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var MemberStatusDbm *MemberStatusDbm_T

func Create_MemberStatusDbm() {
	MemberStatusDbm = new(MemberStatusDbm_T)
	MemberStatusDbm.TableDbName = "MEMBER_STATUS"
	MemberStatusDbm.TableDispName = "MEMBER_STATUS"
	MemberStatusDbm.TablePropertyName = "memberStatus"
	MemberStatusDbm.TableSqlName = new(df.TableSqlName)
	MemberStatusDbm.TableSqlName.TableSqlName = "exampledb.dbo.MEMBER_STATUS"
	MemberStatusDbm.TableSqlName.CorrespondingDbName = MemberStatusDbm.TableDbName

	var memberStatus df.DBMeta
	memberStatus = MemberStatusDbm
	MemberStatusDbm.DBMeta=&memberStatus
	memberStatusCodeSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo MEMBER_STATUS_CODE
	memberStatusCodeSqlName.ColumnSqlName = "MEMBER_STATUS_CODE"
	memberStatusCodeSqlName.IrregularChar = false
	MemberStatusDbm.ColumnMemberStatusCode = df.CCI(&memberStatus, "MEMBER_STATUS_CODE", memberStatusCodeSqlName, "", "", "String.class", "memberStatusCode", "", true, false,true, "char", 3, 0, "",false,"","", "","memberList,memberLoginList","",false,"string")
	memberStatusNameSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo MEMBER_STATUS_NAME
	memberStatusNameSqlName.ColumnSqlName = "MEMBER_STATUS_NAME"
	memberStatusNameSqlName.IrregularChar = false
	MemberStatusDbm.ColumnMemberStatusName = df.CCI(&memberStatus, "MEMBER_STATUS_NAME", memberStatusNameSqlName, "", "", "String.class", "memberStatusName", "", false, false,true, "nvarchar", 50, 0, "",false,"","", "","","",false,"string")
	descriptionSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo DESCRIPTION
	descriptionSqlName.ColumnSqlName = "DESCRIPTION"
	descriptionSqlName.IrregularChar = false
	MemberStatusDbm.ColumnDescription = df.CCI(&memberStatus, "DESCRIPTION", descriptionSqlName, "", "", "String.class", "description", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	displayOrderSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo DISPLAY_ORDER
	displayOrderSqlName.ColumnSqlName = "DISPLAY_ORDER"
	displayOrderSqlName.IrregularChar = false
	MemberStatusDbm.ColumnDisplayOrder = df.CCI(&memberStatus, "DISPLAY_ORDER", displayOrderSqlName, "", "", "Integer.class", "displayOrder", "", false, false,true, "int", 10, 0, "",false,"","", "","","",false,"int64")

	MemberStatusDbm.ColumnInfoList = new(df.List)
	MemberStatusDbm.ColumnInfoList.Add(MemberStatusDbm.ColumnMemberStatusCode)
	MemberStatusDbm.ColumnInfoList.Add(MemberStatusDbm.ColumnMemberStatusName)
	MemberStatusDbm.ColumnInfoList.Add(MemberStatusDbm.ColumnDescription)
	MemberStatusDbm.ColumnInfoList.Add(MemberStatusDbm.ColumnDisplayOrder)


	MemberStatusDbm.ColumnInfoMap=make(map[string]int)
	MemberStatusDbm.ColumnInfoMap["memberStatusCode"]=0
		MemberStatusDbm.ColumnInfoMap["memberStatusName"]=1
		MemberStatusDbm.ColumnInfoMap["description"]=2
		MemberStatusDbm.ColumnInfoMap["displayOrder"]=3
	    MemberStatusDbm.PrimaryKey = true
    MemberStatusDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &memberStatus
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(MemberStatusDbm.ColumnMemberStatusCode)

	MemberStatusDbm.PrimaryInfo = new(df.PrimaryInfo)
	MemberStatusDbm.PrimaryInfo.UniqueInfo = ui
	
	var memberStatusMeta df.DBMeta = MemberStatusDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["MemberStatus"] = &memberStatusMeta
}

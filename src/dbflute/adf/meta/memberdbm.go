package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberDbm_T struct {
	df.BaseDBMeta
	ColumnMemberId *df.ColumnInfo
	ColumnMemberName *df.ColumnInfo
	ColumnMemberAccount *df.ColumnInfo
	ColumnMemberStatusCode *df.ColumnInfo
	ColumnFormalizedDatetime *df.ColumnInfo
	ColumnBirthdate *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnRegisterProcess *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnUpdateProcess *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
}

func (b *MemberDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}
func (b *MemberDbm_T) foreignMemberStatus() *df.ForeignInfo {
	columns := []*df.ColumnInfo{
		MemberDbm.GetColumnInfoByPropertyName("memberStatusCode"),
		MemberStatusDbm.GetColumnInfoByPropertyName("memberStatusCode"),
	}

	return b.BaseDBMeta.Cfi("FK_MEMBER_MEMBER_STATUS", "MemberStatus",
		columns, 0, false, false, false, false,
		"", nil, false, "memberList")
}	
func (b *MemberDbm_T) CreateForeignInfoMap() {
	b.ForeignInfoMap = make(map[string]*df.ForeignInfo)
	b.ForeignInfoMap["MemberStatus"] = b.foreignMemberStatus()
}

func (b *MemberDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var MemberDbm *MemberDbm_T

func Create_MemberDbm() {
	MemberDbm = new(MemberDbm_T)
	MemberDbm.TableDbName = "MEMBER"
	MemberDbm.TableDispName = "MEMBER"
	MemberDbm.TablePropertyName = "member"
	MemberDbm.TableSqlName = new(df.TableSqlName)
	MemberDbm.TableSqlName.TableSqlName = "exampledb.dbo.MEMBER"
	MemberDbm.TableSqlName.CorrespondingDbName = MemberDbm.TableDbName
	MemberDbm.Identity=true

	var member df.DBMeta
	member = MemberDbm
	MemberDbm.DBMeta=&member
	memberIdSqlName := new(df.ColumnSqlName)
	memberIdSqlName.ColumnSqlName = "MEMBER_ID"
	memberIdSqlName.IrregularChar = false
	MemberDbm.ColumnMemberId = df.CCI(&member, "MEMBER_ID", memberIdSqlName, "", "", "Integer.class", "memberId", "", true, true,true, "int identity", 10, 0, "",false,"","", "","memberAddressList,memberLoginList,purchaseList","",false,"int64")
	memberNameSqlName := new(df.ColumnSqlName)
	memberNameSqlName.ColumnSqlName = "MEMBER_NAME"
	memberNameSqlName.IrregularChar = false
	MemberDbm.ColumnMemberName = df.CCI(&member, "MEMBER_NAME", memberNameSqlName, "", "", "String.class", "memberName", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	memberAccountSqlName := new(df.ColumnSqlName)
	memberAccountSqlName.ColumnSqlName = "MEMBER_ACCOUNT"
	memberAccountSqlName.IrregularChar = false
	MemberDbm.ColumnMemberAccount = df.CCI(&member, "MEMBER_ACCOUNT", memberAccountSqlName, "", "", "String.class", "memberAccount", "", false, false,true, "nvarchar", 50, 0, "",false,"","", "","","",false,"string")
	memberStatusCodeSqlName := new(df.ColumnSqlName)
	memberStatusCodeSqlName.ColumnSqlName = "MEMBER_STATUS_CODE"
	memberStatusCodeSqlName.IrregularChar = false
	MemberDbm.ColumnMemberStatusCode = df.CCI(&member, "MEMBER_STATUS_CODE", memberStatusCodeSqlName, "", "", "String.class", "memberStatusCode", "", false, false,true, "char", 3, 0, "",false,"","", "memberStatus","","",false,"string")
	formalizedDatetimeSqlName := new(df.ColumnSqlName)
	formalizedDatetimeSqlName.ColumnSqlName = "FORMALIZED_DATETIME"
	formalizedDatetimeSqlName.IrregularChar = false
	MemberDbm.ColumnFormalizedDatetime = df.CCI(&member, "FORMALIZED_DATETIME", formalizedDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "formalizedDatetime", "", false, false,false, "datetime", 23, 3, "",false,"","", "","","",false,"df.NullTimestamp")
	birthdateSqlName := new(df.ColumnSqlName)
	birthdateSqlName.ColumnSqlName = "BIRTHDATE"
	birthdateSqlName.IrregularChar = false
	MemberDbm.ColumnBirthdate = df.CCI(&member, "BIRTHDATE", birthdateSqlName, "", "", "String.class", "birthdate", "", false, false,false, "date", 10, 0, "",false,"","", "","","",false,"df.NullDate")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	registerDatetimeSqlName.ColumnSqlName = "REGISTER_DATETIME"
	registerDatetimeSqlName.IrregularChar = false
	MemberDbm.ColumnRegisterDatetime = df.CCI(&member, "REGISTER_DATETIME", registerDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "datetime", 23, 3, "",false,"","", "","","",false,"df.Timestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	registerUserSqlName.ColumnSqlName = "REGISTER_USER"
	registerUserSqlName.IrregularChar = false
	MemberDbm.ColumnRegisterUser = df.CCI(&member, "REGISTER_USER", registerUserSqlName, "", "", "String.class", "registerUser", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	registerProcessSqlName := new(df.ColumnSqlName)
	registerProcessSqlName.ColumnSqlName = "REGISTER_PROCESS"
	registerProcessSqlName.IrregularChar = false
	MemberDbm.ColumnRegisterProcess = df.CCI(&member, "REGISTER_PROCESS", registerProcessSqlName, "", "", "String.class", "registerProcess", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	updateDatetimeSqlName.ColumnSqlName = "UPDATE_DATETIME"
	updateDatetimeSqlName.IrregularChar = false
	MemberDbm.ColumnUpdateDatetime = df.CCI(&member, "UPDATE_DATETIME", updateDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "datetime", 23, 3, "",false,"","", "","","",false,"df.Timestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	updateUserSqlName.ColumnSqlName = "UPDATE_USER"
	updateUserSqlName.IrregularChar = false
	MemberDbm.ColumnUpdateUser = df.CCI(&member, "UPDATE_USER", updateUserSqlName, "", "", "String.class", "updateUser", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	updateProcessSqlName := new(df.ColumnSqlName)
	updateProcessSqlName.ColumnSqlName = "UPDATE_PROCESS"
	updateProcessSqlName.IrregularChar = false
	MemberDbm.ColumnUpdateProcess = df.CCI(&member, "UPDATE_PROCESS", updateProcessSqlName, "", "", "String.class", "updateProcess", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	versionNoSqlName := new(df.ColumnSqlName)
	versionNoSqlName.ColumnSqlName = "VERSION_NO"
	versionNoSqlName.IrregularChar = false
	MemberDbm.ColumnVersionNo = df.CCI(&member, "VERSION_NO", versionNoSqlName, "", "", "Long.class", "versionNo", "", false, false,true, "bigint", 19, 0, "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")

	MemberDbm.ColumnInfoList = new(df.List)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnMemberId)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnMemberName)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnMemberAccount)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnMemberStatusCode)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnFormalizedDatetime)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnBirthdate)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnRegisterDatetime)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnRegisterUser)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnRegisterProcess)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnUpdateDatetime)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnUpdateUser)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnUpdateProcess)
	MemberDbm.ColumnInfoList.Add(MemberDbm.ColumnVersionNo)


	MemberDbm.ColumnInfoMap=make(map[string]int)
	MemberDbm.ColumnInfoMap["memberId"]=0
		MemberDbm.ColumnInfoMap["memberName"]=1
		MemberDbm.ColumnInfoMap["memberAccount"]=2
		MemberDbm.ColumnInfoMap["memberStatusCode"]=3
		MemberDbm.ColumnInfoMap["formalizedDatetime"]=4
		MemberDbm.ColumnInfoMap["birthdate"]=5
		MemberDbm.ColumnInfoMap["registerDatetime"]=6
		MemberDbm.ColumnInfoMap["registerUser"]=7
		MemberDbm.ColumnInfoMap["registerProcess"]=8
		MemberDbm.ColumnInfoMap["updateDatetime"]=9
		MemberDbm.ColumnInfoMap["updateUser"]=10
		MemberDbm.ColumnInfoMap["updateProcess"]=11
		MemberDbm.ColumnInfoMap["versionNo"]=12
	    MemberDbm.PrimaryKey = true
    MemberDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &member
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(MemberDbm.ColumnMemberId)

	MemberDbm.PrimaryInfo = new(df.PrimaryInfo)
	MemberDbm.PrimaryInfo.UniqueInfo = ui
	
	MemberDbm.VersionNoFlag = true
	MemberDbm.VersionNoColumnInfo = MemberDbm.ColumnVersionNo
	
	var memberMeta df.DBMeta = MemberDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["Member"] = &memberMeta
}

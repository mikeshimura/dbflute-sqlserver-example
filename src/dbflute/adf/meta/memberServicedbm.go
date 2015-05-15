package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberServiceDbm_T struct {
	df.BaseDBMeta
	ColumnMemberServiceId *df.ColumnInfo
	ColumnMemberId *df.ColumnInfo
	ColumnServicePointCount *df.ColumnInfo
	ColumnServiceRankCode *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnRegisterProcess *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnUpdateProcess *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
}

func (b *MemberServiceDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}
func (b *MemberServiceDbm_T) foreignMember() *df.ForeignInfo {
	columns := []*df.ColumnInfo{
		MemberServiceDbm.GetColumnInfoByPropertyName("memberId"),
		MemberDbm.GetColumnInfoByPropertyName("memberId"),
	}

	return b.BaseDBMeta.Cfi("FK_MEMBER_SERVICE_MEMBER", "Member",
		columns, 0, true, false, false, false,
		"", nil, false, "memberServiceAsOne")
}	
func (b *MemberServiceDbm_T) foreignServiceRank() *df.ForeignInfo {
	columns := []*df.ColumnInfo{
		MemberServiceDbm.GetColumnInfoByPropertyName("serviceRankCode"),
		ServiceRankDbm.GetColumnInfoByPropertyName("serviceRankCode"),
	}

	return b.BaseDBMeta.Cfi("FK_MEMBER_SERVICE_SERVICE_RANK_CODE", "ServiceRank",
		columns, 1, false, false, false, false,
		"", nil, false, "memberServiceList")
}	
func (b *MemberServiceDbm_T) CreateForeignInfoMap() {
	b.ForeignInfoMap = make(map[string]*df.ForeignInfo)
	b.ForeignInfoMap["Member"] = b.foreignMember()
	b.ForeignInfoMap["ServiceRank"] = b.foreignServiceRank()
}

func (b *MemberServiceDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var MemberServiceDbm *MemberServiceDbm_T

func Create_MemberServiceDbm() {
	MemberServiceDbm = new(MemberServiceDbm_T)
	MemberServiceDbm.TableDbName = "MEMBER_SERVICE"
	MemberServiceDbm.TableDispName = "MEMBER_SERVICE"
	MemberServiceDbm.TablePropertyName = "memberService"
	MemberServiceDbm.TableSqlName = new(df.TableSqlName)
	MemberServiceDbm.TableSqlName.TableSqlName = "exampledb.dbo.MEMBER_SERVICE"
	MemberServiceDbm.TableSqlName.CorrespondingDbName = MemberServiceDbm.TableDbName
	MemberServiceDbm.Identity=true

	var memberService df.DBMeta
	memberService = MemberServiceDbm
	MemberServiceDbm.DBMeta=&memberService
	memberServiceIdSqlName := new(df.ColumnSqlName)
	memberServiceIdSqlName.ColumnSqlName = "MEMBER_SERVICE_ID"
	memberServiceIdSqlName.IrregularChar = false
	MemberServiceDbm.ColumnMemberServiceId = df.CCI(&memberService, "MEMBER_SERVICE_ID", memberServiceIdSqlName, "", "", "Integer.class", "memberServiceId", "", true, true,true, "int identity", 10, 0, "",false,"","", "","","",false,"int64")
	memberIdSqlName := new(df.ColumnSqlName)
	memberIdSqlName.ColumnSqlName = "MEMBER_ID"
	memberIdSqlName.IrregularChar = false
	MemberServiceDbm.ColumnMemberId = df.CCI(&memberService, "MEMBER_ID", memberIdSqlName, "", "", "Integer.class", "memberId", "", false, false,true, "int", 10, 0, "",false,"","", "member","","",false,"int64")
	servicePointCountSqlName := new(df.ColumnSqlName)
	servicePointCountSqlName.ColumnSqlName = "SERVICE_POINT_COUNT"
	servicePointCountSqlName.IrregularChar = false
	MemberServiceDbm.ColumnServicePointCount = df.CCI(&memberService, "SERVICE_POINT_COUNT", servicePointCountSqlName, "", "", "Integer.class", "servicePointCount", "", false, false,true, "int", 10, 0, "",false,"","", "","","",false,"int64")
	serviceRankCodeSqlName := new(df.ColumnSqlName)
	serviceRankCodeSqlName.ColumnSqlName = "SERVICE_RANK_CODE"
	serviceRankCodeSqlName.IrregularChar = false
	MemberServiceDbm.ColumnServiceRankCode = df.CCI(&memberService, "SERVICE_RANK_CODE", serviceRankCodeSqlName, "", "", "String.class", "serviceRankCode", "", false, false,true, "char", 3, 0, "",false,"","", "serviceRank","","",false,"string")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	registerDatetimeSqlName.ColumnSqlName = "REGISTER_DATETIME"
	registerDatetimeSqlName.IrregularChar = false
	MemberServiceDbm.ColumnRegisterDatetime = df.CCI(&memberService, "REGISTER_DATETIME", registerDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "datetime", 23, 3, "",false,"","", "","","",false,"df.Timestamp")
	registerUserSqlName := new(df.ColumnSqlName)
	registerUserSqlName.ColumnSqlName = "REGISTER_USER"
	registerUserSqlName.IrregularChar = false
	MemberServiceDbm.ColumnRegisterUser = df.CCI(&memberService, "REGISTER_USER", registerUserSqlName, "", "", "String.class", "registerUser", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	registerProcessSqlName := new(df.ColumnSqlName)
	registerProcessSqlName.ColumnSqlName = "REGISTER_PROCESS"
	registerProcessSqlName.IrregularChar = false
	MemberServiceDbm.ColumnRegisterProcess = df.CCI(&memberService, "REGISTER_PROCESS", registerProcessSqlName, "", "", "String.class", "registerProcess", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	updateDatetimeSqlName.ColumnSqlName = "UPDATE_DATETIME"
	updateDatetimeSqlName.IrregularChar = false
	MemberServiceDbm.ColumnUpdateDatetime = df.CCI(&memberService, "UPDATE_DATETIME", updateDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "datetime", 23, 3, "",false,"","", "","","",false,"df.Timestamp")
	updateUserSqlName := new(df.ColumnSqlName)
	updateUserSqlName.ColumnSqlName = "UPDATE_USER"
	updateUserSqlName.IrregularChar = false
	MemberServiceDbm.ColumnUpdateUser = df.CCI(&memberService, "UPDATE_USER", updateUserSqlName, "", "", "String.class", "updateUser", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	updateProcessSqlName := new(df.ColumnSqlName)
	updateProcessSqlName.ColumnSqlName = "UPDATE_PROCESS"
	updateProcessSqlName.IrregularChar = false
	MemberServiceDbm.ColumnUpdateProcess = df.CCI(&memberService, "UPDATE_PROCESS", updateProcessSqlName, "", "", "String.class", "updateProcess", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	versionNoSqlName := new(df.ColumnSqlName)
	versionNoSqlName.ColumnSqlName = "VERSION_NO"
	versionNoSqlName.IrregularChar = false
	MemberServiceDbm.ColumnVersionNo = df.CCI(&memberService, "VERSION_NO", versionNoSqlName, "", "", "Long.class", "versionNo", "", false, false,true, "bigint", 19, 0, "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")

	MemberServiceDbm.ColumnInfoList = new(df.List)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnMemberServiceId)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnMemberId)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnServicePointCount)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnServiceRankCode)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnRegisterDatetime)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnRegisterUser)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnRegisterProcess)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnUpdateDatetime)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnUpdateUser)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnUpdateProcess)
	MemberServiceDbm.ColumnInfoList.Add(MemberServiceDbm.ColumnVersionNo)


	MemberServiceDbm.ColumnInfoMap=make(map[string]int)
	MemberServiceDbm.ColumnInfoMap["memberServiceId"]=0
		MemberServiceDbm.ColumnInfoMap["memberId"]=1
		MemberServiceDbm.ColumnInfoMap["servicePointCount"]=2
		MemberServiceDbm.ColumnInfoMap["serviceRankCode"]=3
		MemberServiceDbm.ColumnInfoMap["registerDatetime"]=4
		MemberServiceDbm.ColumnInfoMap["registerUser"]=5
		MemberServiceDbm.ColumnInfoMap["registerProcess"]=6
		MemberServiceDbm.ColumnInfoMap["updateDatetime"]=7
		MemberServiceDbm.ColumnInfoMap["updateUser"]=8
		MemberServiceDbm.ColumnInfoMap["updateProcess"]=9
		MemberServiceDbm.ColumnInfoMap["versionNo"]=10
	    MemberServiceDbm.PrimaryKey = true
    MemberServiceDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &memberService
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(MemberServiceDbm.ColumnMemberServiceId)

	MemberServiceDbm.PrimaryInfo = new(df.PrimaryInfo)
	MemberServiceDbm.PrimaryInfo.UniqueInfo = ui
	
	MemberServiceDbm.VersionNoFlag = true
	MemberServiceDbm.VersionNoColumnInfo = MemberServiceDbm.ColumnVersionNo
	
	var memberServiceMeta df.DBMeta = MemberServiceDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["MemberService"] = &memberServiceMeta
}

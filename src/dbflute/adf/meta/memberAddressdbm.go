package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberAddressDbm_T struct {
	df.BaseDBMeta
	ColumnMemberAddressId *df.ColumnInfo
	ColumnMemberId *df.ColumnInfo
	ColumnValidBeginDate *df.ColumnInfo
	ColumnValidEndDate *df.ColumnInfo
	ColumnAddress *df.ColumnInfo
	ColumnRegionId *df.ColumnInfo
	ColumnRegisterDatetime *df.ColumnInfo
	ColumnRegisterProcess *df.ColumnInfo
	ColumnRegisterUser *df.ColumnInfo
	ColumnUpdateDatetime *df.ColumnInfo
	ColumnUpdateProcess *df.ColumnInfo
	ColumnUpdateUser *df.ColumnInfo
	ColumnVersionNo *df.ColumnInfo
}

func (b *MemberAddressDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}
func (b *MemberAddressDbm_T) foreignMember() *df.ForeignInfo {
	columns := []*df.ColumnInfo{
		MemberAddressDbm.GetColumnInfoByPropertyName("memberId"),
		MemberDbm.GetColumnInfoByPropertyName("memberId"),
	}

	return b.BaseDBMeta.Cfi("FK_MEMBER_ADDRESS_MEMBER", "Member",
		columns, 0, false, false, false, false,
		"", nil, false, "memberAddressList")
}	
func (b *MemberAddressDbm_T) foreignRegion() *df.ForeignInfo {
	columns := []*df.ColumnInfo{
		MemberAddressDbm.GetColumnInfoByPropertyName("regionId"),
		RegionDbm.GetColumnInfoByPropertyName("regionId"),
	}

	return b.BaseDBMeta.Cfi("FK_MEMBER_ADDRESS_REGION", "Region",
		columns, 1, false, false, false, false,
		"", nil, false, "memberAddressList")
}	
func (b *MemberAddressDbm_T) CreateForeignInfoMap() {
	b.ForeignInfoMap = make(map[string]*df.ForeignInfo)
	b.ForeignInfoMap["Member"] = b.foreignMember()
	b.ForeignInfoMap["Region"] = b.foreignRegion()
}

func (b *MemberAddressDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var MemberAddressDbm *MemberAddressDbm_T

func Create_MemberAddressDbm() {
	MemberAddressDbm = new(MemberAddressDbm_T)
	MemberAddressDbm.TableDbName = "MEMBER_ADDRESS"
	MemberAddressDbm.TableDispName = "MEMBER_ADDRESS"
	MemberAddressDbm.TablePropertyName = "memberAddress"
	MemberAddressDbm.TableSqlName = new(df.TableSqlName)
	MemberAddressDbm.TableSqlName.TableSqlName = "exampledb.dbo.MEMBER_ADDRESS"
	MemberAddressDbm.TableSqlName.CorrespondingDbName = MemberAddressDbm.TableDbName

	var memberAddress df.DBMeta
	memberAddress = MemberAddressDbm
	MemberAddressDbm.DBMeta=&memberAddress
	memberAddressIdSqlName := new(df.ColumnSqlName)
	memberAddressIdSqlName.ColumnSqlName = "MEMBER_ADDRESS_ID"
	memberAddressIdSqlName.IrregularChar = false
	MemberAddressDbm.ColumnMemberAddressId = df.CCI(&memberAddress, "MEMBER_ADDRESS_ID", memberAddressIdSqlName, "", "", "Integer.class", "memberAddressId", "", true, false,true, "int", 10, 0, "",false,"","", "","","",false,"int64")
	memberIdSqlName := new(df.ColumnSqlName)
	memberIdSqlName.ColumnSqlName = "MEMBER_ID"
	memberIdSqlName.IrregularChar = false
	MemberAddressDbm.ColumnMemberId = df.CCI(&memberAddress, "MEMBER_ID", memberIdSqlName, "", "", "Integer.class", "memberId", "", false, false,true, "int", 10, 0, "",false,"","", "member","","",false,"int64")
	validBeginDateSqlName := new(df.ColumnSqlName)
	validBeginDateSqlName.ColumnSqlName = "VALID_BEGIN_DATE"
	validBeginDateSqlName.IrregularChar = false
	MemberAddressDbm.ColumnValidBeginDate = df.CCI(&memberAddress, "VALID_BEGIN_DATE", validBeginDateSqlName, "", "", "String.class", "validBeginDate", "", false, false,true, "date", 10, 0, "",false,"","", "","","",false,"df.Date")
	validEndDateSqlName := new(df.ColumnSqlName)
	validEndDateSqlName.ColumnSqlName = "VALID_END_DATE"
	validEndDateSqlName.IrregularChar = false
	MemberAddressDbm.ColumnValidEndDate = df.CCI(&memberAddress, "VALID_END_DATE", validEndDateSqlName, "", "", "String.class", "validEndDate", "", false, false,true, "date", 10, 0, "",false,"","", "","","",false,"df.Date")
	addressSqlName := new(df.ColumnSqlName)
	addressSqlName.ColumnSqlName = "ADDRESS"
	addressSqlName.IrregularChar = false
	MemberAddressDbm.ColumnAddress = df.CCI(&memberAddress, "ADDRESS", addressSqlName, "", "", "String.class", "address", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	regionIdSqlName := new(df.ColumnSqlName)
	regionIdSqlName.ColumnSqlName = "REGION_ID"
	regionIdSqlName.IrregularChar = false
	MemberAddressDbm.ColumnRegionId = df.CCI(&memberAddress, "REGION_ID", regionIdSqlName, "", "", "Integer.class", "regionId", "", false, false,true, "int", 10, 0, "",false,"","", "region","","",false,"int64")
	registerDatetimeSqlName := new(df.ColumnSqlName)
	registerDatetimeSqlName.ColumnSqlName = "REGISTER_DATETIME"
	registerDatetimeSqlName.IrregularChar = false
	MemberAddressDbm.ColumnRegisterDatetime = df.CCI(&memberAddress, "REGISTER_DATETIME", registerDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "registerDatetime", "", false, false,true, "datetime", 23, 3, "",false,"","", "","","",false,"df.Timestamp")
	registerProcessSqlName := new(df.ColumnSqlName)
	registerProcessSqlName.ColumnSqlName = "REGISTER_PROCESS"
	registerProcessSqlName.IrregularChar = false
	MemberAddressDbm.ColumnRegisterProcess = df.CCI(&memberAddress, "REGISTER_PROCESS", registerProcessSqlName, "", "", "String.class", "registerProcess", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	registerUserSqlName := new(df.ColumnSqlName)
	registerUserSqlName.ColumnSqlName = "REGISTER_USER"
	registerUserSqlName.IrregularChar = false
	MemberAddressDbm.ColumnRegisterUser = df.CCI(&memberAddress, "REGISTER_USER", registerUserSqlName, "", "", "String.class", "registerUser", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	updateDatetimeSqlName := new(df.ColumnSqlName)
	updateDatetimeSqlName.ColumnSqlName = "UPDATE_DATETIME"
	updateDatetimeSqlName.IrregularChar = false
	MemberAddressDbm.ColumnUpdateDatetime = df.CCI(&memberAddress, "UPDATE_DATETIME", updateDatetimeSqlName, "", "", "java.time.LocalDateTime.class", "updateDatetime", "", false, false,true, "datetime", 23, 3, "",false,"","", "","","",false,"df.Timestamp")
	updateProcessSqlName := new(df.ColumnSqlName)
	updateProcessSqlName.ColumnSqlName = "UPDATE_PROCESS"
	updateProcessSqlName.IrregularChar = false
	MemberAddressDbm.ColumnUpdateProcess = df.CCI(&memberAddress, "UPDATE_PROCESS", updateProcessSqlName, "", "", "String.class", "updateProcess", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	updateUserSqlName := new(df.ColumnSqlName)
	updateUserSqlName.ColumnSqlName = "UPDATE_USER"
	updateUserSqlName.IrregularChar = false
	MemberAddressDbm.ColumnUpdateUser = df.CCI(&memberAddress, "UPDATE_USER", updateUserSqlName, "", "", "String.class", "updateUser", "", false, false,true, "nvarchar", 200, 0, "",false,"","", "","","",false,"string")
	versionNoSqlName := new(df.ColumnSqlName)
	versionNoSqlName.ColumnSqlName = "VERSION_NO"
	versionNoSqlName.IrregularChar = false
	MemberAddressDbm.ColumnVersionNo = df.CCI(&memberAddress, "VERSION_NO", versionNoSqlName, "", "", "Long.class", "versionNo", "", false, false,true, "bigint", 19, 0, "",false,"OptimisticLockType.VERSION_NO","", "","","",false,"int64")

	MemberAddressDbm.ColumnInfoList = new(df.List)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnMemberAddressId)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnMemberId)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnValidBeginDate)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnValidEndDate)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnAddress)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnRegionId)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnRegisterDatetime)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnRegisterProcess)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnRegisterUser)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnUpdateDatetime)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnUpdateProcess)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnUpdateUser)
	MemberAddressDbm.ColumnInfoList.Add(MemberAddressDbm.ColumnVersionNo)


	MemberAddressDbm.ColumnInfoMap=make(map[string]int)
	MemberAddressDbm.ColumnInfoMap["memberAddressId"]=0
		MemberAddressDbm.ColumnInfoMap["memberId"]=1
		MemberAddressDbm.ColumnInfoMap["validBeginDate"]=2
		MemberAddressDbm.ColumnInfoMap["validEndDate"]=3
		MemberAddressDbm.ColumnInfoMap["address"]=4
		MemberAddressDbm.ColumnInfoMap["regionId"]=5
		MemberAddressDbm.ColumnInfoMap["registerDatetime"]=6
		MemberAddressDbm.ColumnInfoMap["registerProcess"]=7
		MemberAddressDbm.ColumnInfoMap["registerUser"]=8
		MemberAddressDbm.ColumnInfoMap["updateDatetime"]=9
		MemberAddressDbm.ColumnInfoMap["updateProcess"]=10
		MemberAddressDbm.ColumnInfoMap["updateUser"]=11
		MemberAddressDbm.ColumnInfoMap["versionNo"]=12
	    MemberAddressDbm.PrimaryKey = true
    MemberAddressDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &memberAddress
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(MemberAddressDbm.ColumnMemberAddressId)

	MemberAddressDbm.PrimaryInfo = new(df.PrimaryInfo)
	MemberAddressDbm.PrimaryInfo.UniqueInfo = ui
	
	MemberAddressDbm.VersionNoFlag = true
	MemberAddressDbm.VersionNoColumnInfo = MemberAddressDbm.ColumnVersionNo
	
	var memberAddressMeta df.DBMeta = MemberAddressDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["MemberAddress"] = &memberAddressMeta
}

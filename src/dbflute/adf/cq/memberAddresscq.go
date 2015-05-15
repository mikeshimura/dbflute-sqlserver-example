package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type MemberAddressCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	MemberAddressId *df.ConditionValue
	MemberId *df.ConditionValue
	ValidBeginDate *df.ConditionValue
	ValidEndDate *df.ConditionValue
	Address *df.ConditionValue
	RegionId *df.ConditionValue
	RegisterDatetime *df.ConditionValue
	RegisterProcess *df.ConditionValue
	RegisterUser *df.ConditionValue
	UpdateDatetime *df.ConditionValue
	UpdateProcess *df.ConditionValue
	UpdateUser *df.ConditionValue
	VersionNo *df.ConditionValue
    conditionQueryMember *MemberCQ
    conditionQueryRegion *RegionCQ
}

func (q *MemberAddressCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *MemberAddressCQ) getCValueMemberAddressId() *df.ConditionValue {
	if q.MemberAddressId == nil {
		q.MemberAddressId = new(df.ConditionValue)
	}
	return q.MemberAddressId
}



func (q *MemberAddressCQ) SetMemberAddressId_Equal(value int64) *MemberAddressCQ {
	q.regMemberAddressId(df.CK_EQ_C, value)
	return q
}
func (q *MemberAddressCQ) SetMemberAddressId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueMemberAddressId(), "memberAddressId")
}
func (q *MemberAddressCQ) SetMemberAddressId_NotEqual(value int64) *MemberAddressCQ {
	q.regMemberAddressId(df.CK_NE_C, value)
	return q
}

func (q *MemberAddressCQ) SetMemberAddressId_GreaterThan(value int64) *MemberAddressCQ {
	q.regMemberAddressId(df.CK_GT_C, value)
	return q
}

func (q *MemberAddressCQ) SetMemberAddressId_LessThan(value int64) *MemberAddressCQ {
	q.regMemberAddressId(df.CK_LT_C, value)
	return q
}

func (q *MemberAddressCQ) SetMemberAddressId_GreaterEqual(value int64) *MemberAddressCQ {
	q.regMemberAddressId(df.CK_GE_C, value)
	return q
}

func (q *MemberAddressCQ) SetMemberAddressId_LessEqual(value int64) *MemberAddressCQ {
	q.regMemberAddressId(df.CK_LE_C, value)
	return q
}
func (q *MemberAddressCQ) SetMemberAddressId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueMemberAddressId(),"memberAddressId",rangeOfOption)
}	


func (q *MemberAddressCQ) SetMemberAddressId_IsNull() *MemberAddressCQ {
	q.regMemberAddressId(df.CK_ISN_C, 0)
	return q
}
func (q *MemberAddressCQ) SetMemberAddressId_IsNotNull() *MemberAddressCQ {
	q.regMemberAddressId(df.CK_ISNN_C, 0)
	return q
}
func (q *MemberAddressCQ) AddOrderBy_MemberAddressId_Asc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBA("memberAddressId")
	return q
}
func (q *MemberAddressCQ) AddOrderBy_MemberAddressId_Desc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBD("memberAddressId")
	return q
}
func (q *MemberAddressCQ) regMemberAddressId(key *df.ConditionKey, value interface{}) {
	if q.MemberAddressId == nil {
		q.MemberAddressId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MemberAddressId, "memberAddressId")
}

func (q *MemberAddressCQ) getCValueMemberId() *df.ConditionValue {
	if q.MemberId == nil {
		q.MemberId = new(df.ConditionValue)
	}
	return q.MemberId
}



func (q *MemberAddressCQ) SetMemberId_Equal(value int64) *MemberAddressCQ {
	q.regMemberId(df.CK_EQ_C, value)
	return q
}
func (q *MemberAddressCQ) SetMemberId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueMemberId(), "memberId")
}
func (q *MemberAddressCQ) SetMemberId_NotEqual(value int64) *MemberAddressCQ {
	q.regMemberId(df.CK_NE_C, value)
	return q
}

func (q *MemberAddressCQ) SetMemberId_GreaterThan(value int64) *MemberAddressCQ {
	q.regMemberId(df.CK_GT_C, value)
	return q
}

func (q *MemberAddressCQ) SetMemberId_LessThan(value int64) *MemberAddressCQ {
	q.regMemberId(df.CK_LT_C, value)
	return q
}

func (q *MemberAddressCQ) SetMemberId_GreaterEqual(value int64) *MemberAddressCQ {
	q.regMemberId(df.CK_GE_C, value)
	return q
}

func (q *MemberAddressCQ) SetMemberId_LessEqual(value int64) *MemberAddressCQ {
	q.regMemberId(df.CK_LE_C, value)
	return q
}
func (q *MemberAddressCQ) SetMemberId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueMemberId(),"memberId",rangeOfOption)
}	


func (q *MemberAddressCQ) AddOrderBy_MemberId_Asc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBA("memberId")
	return q
}
func (q *MemberAddressCQ) AddOrderBy_MemberId_Desc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBD("memberId")
	return q
}
func (q *MemberAddressCQ) regMemberId(key *df.ConditionKey, value interface{}) {
	if q.MemberId == nil {
		q.MemberId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.MemberId, "memberId")
}

func (q *MemberAddressCQ) getCValueValidBeginDate() *df.ConditionValue {
	if q.ValidBeginDate == nil {
		q.ValidBeginDate = new(df.ConditionValue)
	}
	return q.ValidBeginDate
}


func (q *MemberAddressCQ) SetValidBeginDate_Equal(value df.Date) *MemberAddressCQ {
	q.regValidBeginDate(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *MemberAddressCQ) SetValidBeginDate_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueValidBeginDate(), "validBeginDate")
}
func (q *MemberAddressCQ) SetValidBeginDate_NotEqual(value df.Date) *MemberAddressCQ {
	q.regValidBeginDate(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetValidBeginDate_GreaterThan(value df.Date) *MemberAddressCQ {
	q.regValidBeginDate(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetValidBeginDate_LessThan(value df.Date) *MemberAddressCQ {
	q.regValidBeginDate(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetValidBeginDate_GreaterEqualThan(value df.Date) *MemberAddressCQ {
	q.regValidBeginDate(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberAddressCQ) SetValidBeginDate_LessEqualThan(value df.Date) *MemberAddressCQ {
	q.regValidBeginDate(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetValidBeginDate_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueValidBeginDate(), "validBeginDate", option)
}

func (q *MemberAddressCQ) SetValidBeginDate_PrefixSearch(value string) error {
	return q.SetValidBeginDate_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberAddressCQ) SetValidBeginDate_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueValidBeginDate(), "validBeginDate", option)
}



func (q *MemberAddressCQ) AddOrderBy_ValidBeginDate_Asc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBA("validBeginDate")
	return q
}
func (q *MemberAddressCQ) AddOrderBy_ValidBeginDate_Desc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBD("validBeginDate")
	return q
}
func (q *MemberAddressCQ) regValidBeginDate(key *df.ConditionKey, value interface{}) {
	if q.ValidBeginDate == nil {
		q.ValidBeginDate = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ValidBeginDate, "validBeginDate")
}

func (q *MemberAddressCQ) getCValueValidEndDate() *df.ConditionValue {
	if q.ValidEndDate == nil {
		q.ValidEndDate = new(df.ConditionValue)
	}
	return q.ValidEndDate
}


func (q *MemberAddressCQ) SetValidEndDate_Equal(value df.Date) *MemberAddressCQ {
	q.regValidEndDate(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *MemberAddressCQ) SetValidEndDate_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueValidEndDate(), "validEndDate")
}
func (q *MemberAddressCQ) SetValidEndDate_NotEqual(value df.Date) *MemberAddressCQ {
	q.regValidEndDate(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetValidEndDate_GreaterThan(value df.Date) *MemberAddressCQ {
	q.regValidEndDate(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetValidEndDate_LessThan(value df.Date) *MemberAddressCQ {
	q.regValidEndDate(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetValidEndDate_GreaterEqualThan(value df.Date) *MemberAddressCQ {
	q.regValidEndDate(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberAddressCQ) SetValidEndDate_LessEqualThan(value df.Date) *MemberAddressCQ {
	q.regValidEndDate(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetValidEndDate_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueValidEndDate(), "validEndDate", option)
}

func (q *MemberAddressCQ) SetValidEndDate_PrefixSearch(value string) error {
	return q.SetValidEndDate_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberAddressCQ) SetValidEndDate_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueValidEndDate(), "validEndDate", option)
}



func (q *MemberAddressCQ) AddOrderBy_ValidEndDate_Asc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBA("validEndDate")
	return q
}
func (q *MemberAddressCQ) AddOrderBy_ValidEndDate_Desc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBD("validEndDate")
	return q
}
func (q *MemberAddressCQ) regValidEndDate(key *df.ConditionKey, value interface{}) {
	if q.ValidEndDate == nil {
		q.ValidEndDate = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ValidEndDate, "validEndDate")
}

func (q *MemberAddressCQ) getCValueAddress() *df.ConditionValue {
	if q.Address == nil {
		q.Address = new(df.ConditionValue)
	}
	return q.Address
}


func (q *MemberAddressCQ) SetAddress_Equal(value string) *MemberAddressCQ {
	q.regAddress(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *MemberAddressCQ) SetAddress_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueAddress(), "address")
}
func (q *MemberAddressCQ) SetAddress_NotEqual(value string) *MemberAddressCQ {
	q.regAddress(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetAddress_GreaterThan(value string) *MemberAddressCQ {
	q.regAddress(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetAddress_LessThan(value string) *MemberAddressCQ {
	q.regAddress(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetAddress_GreaterEqualThan(value string) *MemberAddressCQ {
	q.regAddress(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberAddressCQ) SetAddress_LessEqualThan(value string) *MemberAddressCQ {
	q.regAddress(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetAddress_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueAddress(), "address", option)
}

func (q *MemberAddressCQ) SetAddress_PrefixSearch(value string) error {
	return q.SetAddress_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberAddressCQ) SetAddress_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueAddress(), "address", option)
}



func (q *MemberAddressCQ) AddOrderBy_Address_Asc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBA("address")
	return q
}
func (q *MemberAddressCQ) AddOrderBy_Address_Desc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBD("address")
	return q
}
func (q *MemberAddressCQ) regAddress(key *df.ConditionKey, value interface{}) {
	if q.Address == nil {
		q.Address = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.Address, "address")
}

func (q *MemberAddressCQ) getCValueRegionId() *df.ConditionValue {
	if q.RegionId == nil {
		q.RegionId = new(df.ConditionValue)
	}
	return q.RegionId
}



func (q *MemberAddressCQ) SetRegionId_Equal(value int64) *MemberAddressCQ {
	q.regRegionId(df.CK_EQ_C, value)
	return q
}
func (q *MemberAddressCQ) SetRegionId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegionId(), "regionId")
}
func (q *MemberAddressCQ) SetRegionId_NotEqual(value int64) *MemberAddressCQ {
	q.regRegionId(df.CK_NE_C, value)
	return q
}

func (q *MemberAddressCQ) SetRegionId_GreaterThan(value int64) *MemberAddressCQ {
	q.regRegionId(df.CK_GT_C, value)
	return q
}

func (q *MemberAddressCQ) SetRegionId_LessThan(value int64) *MemberAddressCQ {
	q.regRegionId(df.CK_LT_C, value)
	return q
}

func (q *MemberAddressCQ) SetRegionId_GreaterEqual(value int64) *MemberAddressCQ {
	q.regRegionId(df.CK_GE_C, value)
	return q
}

func (q *MemberAddressCQ) SetRegionId_LessEqual(value int64) *MemberAddressCQ {
	q.regRegionId(df.CK_LE_C, value)
	return q
}
func (q *MemberAddressCQ) SetRegionId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueRegionId(),"regionId",rangeOfOption)
}	


func (q *MemberAddressCQ) AddOrderBy_RegionId_Asc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBA("regionId")
	return q
}
func (q *MemberAddressCQ) AddOrderBy_RegionId_Desc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBD("regionId")
	return q
}
func (q *MemberAddressCQ) regRegionId(key *df.ConditionKey, value interface{}) {
	if q.RegionId == nil {
		q.RegionId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegionId, "regionId")
}

func (q *MemberAddressCQ) getCValueRegisterDatetime() *df.ConditionValue {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	return q.RegisterDatetime
}




func (q *MemberAddressCQ) SetRegisterDatetime_Equal(value df.Timestamp) *MemberAddressCQ {
	q.regRegisterDatetime(df.CK_EQ_C, value)
	return q
}


func (q *MemberAddressCQ) SetRegisterDatetime_GreaterThan(value df.Timestamp) *MemberAddressCQ {
	q.regRegisterDatetime(df.CK_GT_C, value)
	return q
}

func (q *MemberAddressCQ) SetRegisterDatetime_LessThan(value df.Timestamp) *MemberAddressCQ {
	q.regRegisterDatetime(df.CK_LT_C, value)
	return q
}

func (q *MemberAddressCQ) SetRegisterDatetime_GreaterEqual(value df.Timestamp) *MemberAddressCQ {
	q.regRegisterDatetime(df.CK_GE_C, value)
	return q
}

func (q *MemberAddressCQ) SetRegisterDatetime_LessEqual(value df.Timestamp) *MemberAddressCQ {
	q.regRegisterDatetime(df.CK_LE_C, value)
	return q
}

func (q *MemberAddressCQ) AddOrderBy_RegisterDatetime_Asc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBA("registerDatetime")
	return q
}
func (q *MemberAddressCQ) AddOrderBy_RegisterDatetime_Desc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBD("registerDatetime")
	return q
}
func (q *MemberAddressCQ) regRegisterDatetime(key *df.ConditionKey, value interface{}) {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterDatetime, "registerDatetime")
}

func (q *MemberAddressCQ) getCValueRegisterProcess() *df.ConditionValue {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	return q.RegisterProcess
}


func (q *MemberAddressCQ) SetRegisterProcess_Equal(value string) *MemberAddressCQ {
	q.regRegisterProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *MemberAddressCQ) SetRegisterProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterProcess(), "registerProcess")
}
func (q *MemberAddressCQ) SetRegisterProcess_NotEqual(value string) *MemberAddressCQ {
	q.regRegisterProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetRegisterProcess_GreaterThan(value string) *MemberAddressCQ {
	q.regRegisterProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetRegisterProcess_LessThan(value string) *MemberAddressCQ {
	q.regRegisterProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetRegisterProcess_GreaterEqualThan(value string) *MemberAddressCQ {
	q.regRegisterProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberAddressCQ) SetRegisterProcess_LessEqualThan(value string) *MemberAddressCQ {
	q.regRegisterProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetRegisterProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}

func (q *MemberAddressCQ) SetRegisterProcess_PrefixSearch(value string) error {
	return q.SetRegisterProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberAddressCQ) SetRegisterProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterProcess(), "registerProcess", option)
}



func (q *MemberAddressCQ) AddOrderBy_RegisterProcess_Asc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBA("registerProcess")
	return q
}
func (q *MemberAddressCQ) AddOrderBy_RegisterProcess_Desc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBD("registerProcess")
	return q
}
func (q *MemberAddressCQ) regRegisterProcess(key *df.ConditionKey, value interface{}) {
	if q.RegisterProcess == nil {
		q.RegisterProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterProcess, "registerProcess")
}

func (q *MemberAddressCQ) getCValueRegisterUser() *df.ConditionValue {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	return q.RegisterUser
}


func (q *MemberAddressCQ) SetRegisterUser_Equal(value string) *MemberAddressCQ {
	q.regRegisterUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *MemberAddressCQ) SetRegisterUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueRegisterUser(), "registerUser")
}
func (q *MemberAddressCQ) SetRegisterUser_NotEqual(value string) *MemberAddressCQ {
	q.regRegisterUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetRegisterUser_GreaterThan(value string) *MemberAddressCQ {
	q.regRegisterUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetRegisterUser_LessThan(value string) *MemberAddressCQ {
	q.regRegisterUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetRegisterUser_GreaterEqualThan(value string) *MemberAddressCQ {
	q.regRegisterUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberAddressCQ) SetRegisterUser_LessEqualThan(value string) *MemberAddressCQ {
	q.regRegisterUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetRegisterUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}

func (q *MemberAddressCQ) SetRegisterUser_PrefixSearch(value string) error {
	return q.SetRegisterUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberAddressCQ) SetRegisterUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}



func (q *MemberAddressCQ) AddOrderBy_RegisterUser_Asc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBA("registerUser")
	return q
}
func (q *MemberAddressCQ) AddOrderBy_RegisterUser_Desc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBD("registerUser")
	return q
}
func (q *MemberAddressCQ) regRegisterUser(key *df.ConditionKey, value interface{}) {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterUser, "registerUser")
}

func (q *MemberAddressCQ) getCValueUpdateDatetime() *df.ConditionValue {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	return q.UpdateDatetime
}




func (q *MemberAddressCQ) SetUpdateDatetime_Equal(value df.Timestamp) *MemberAddressCQ {
	q.regUpdateDatetime(df.CK_EQ_C, value)
	return q
}


func (q *MemberAddressCQ) SetUpdateDatetime_GreaterThan(value df.Timestamp) *MemberAddressCQ {
	q.regUpdateDatetime(df.CK_GT_C, value)
	return q
}

func (q *MemberAddressCQ) SetUpdateDatetime_LessThan(value df.Timestamp) *MemberAddressCQ {
	q.regUpdateDatetime(df.CK_LT_C, value)
	return q
}

func (q *MemberAddressCQ) SetUpdateDatetime_GreaterEqual(value df.Timestamp) *MemberAddressCQ {
	q.regUpdateDatetime(df.CK_GE_C, value)
	return q
}

func (q *MemberAddressCQ) SetUpdateDatetime_LessEqual(value df.Timestamp) *MemberAddressCQ {
	q.regUpdateDatetime(df.CK_LE_C, value)
	return q
}

func (q *MemberAddressCQ) AddOrderBy_UpdateDatetime_Asc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBA("updateDatetime")
	return q
}
func (q *MemberAddressCQ) AddOrderBy_UpdateDatetime_Desc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBD("updateDatetime")
	return q
}
func (q *MemberAddressCQ) regUpdateDatetime(key *df.ConditionKey, value interface{}) {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateDatetime, "updateDatetime")
}

func (q *MemberAddressCQ) getCValueUpdateProcess() *df.ConditionValue {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	return q.UpdateProcess
}


func (q *MemberAddressCQ) SetUpdateProcess_Equal(value string) *MemberAddressCQ {
	q.regUpdateProcess(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *MemberAddressCQ) SetUpdateProcess_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateProcess(), "updateProcess")
}
func (q *MemberAddressCQ) SetUpdateProcess_NotEqual(value string) *MemberAddressCQ {
	q.regUpdateProcess(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetUpdateProcess_GreaterThan(value string) *MemberAddressCQ {
	q.regUpdateProcess(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetUpdateProcess_LessThan(value string) *MemberAddressCQ {
	q.regUpdateProcess(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetUpdateProcess_GreaterEqualThan(value string) *MemberAddressCQ {
	q.regUpdateProcess(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberAddressCQ) SetUpdateProcess_LessEqualThan(value string) *MemberAddressCQ {
	q.regUpdateProcess(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetUpdateProcess_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}

func (q *MemberAddressCQ) SetUpdateProcess_PrefixSearch(value string) error {
	return q.SetUpdateProcess_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberAddressCQ) SetUpdateProcess_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateProcess(), "updateProcess", option)
}



func (q *MemberAddressCQ) AddOrderBy_UpdateProcess_Asc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBA("updateProcess")
	return q
}
func (q *MemberAddressCQ) AddOrderBy_UpdateProcess_Desc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBD("updateProcess")
	return q
}
func (q *MemberAddressCQ) regUpdateProcess(key *df.ConditionKey, value interface{}) {
	if q.UpdateProcess == nil {
		q.UpdateProcess = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateProcess, "updateProcess")
}

func (q *MemberAddressCQ) getCValueUpdateUser() *df.ConditionValue {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	return q.UpdateUser
}


func (q *MemberAddressCQ) SetUpdateUser_Equal(value string) *MemberAddressCQ {
	q.regUpdateUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *MemberAddressCQ) SetUpdateUser_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueUpdateUser(), "updateUser")
}
func (q *MemberAddressCQ) SetUpdateUser_NotEqual(value string) *MemberAddressCQ {
	q.regUpdateUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetUpdateUser_GreaterThan(value string) *MemberAddressCQ {
	q.regUpdateUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetUpdateUser_LessThan(value string) *MemberAddressCQ {
	q.regUpdateUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetUpdateUser_GreaterEqualThan(value string) *MemberAddressCQ {
	q.regUpdateUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *MemberAddressCQ) SetUpdateUser_LessEqualThan(value string) *MemberAddressCQ {
	q.regUpdateUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *MemberAddressCQ) SetUpdateUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}

func (q *MemberAddressCQ) SetUpdateUser_PrefixSearch(value string) error {
	return q.SetUpdateUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *MemberAddressCQ) SetUpdateUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}



func (q *MemberAddressCQ) AddOrderBy_UpdateUser_Asc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBA("updateUser")
	return q
}
func (q *MemberAddressCQ) AddOrderBy_UpdateUser_Desc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBD("updateUser")
	return q
}
func (q *MemberAddressCQ) regUpdateUser(key *df.ConditionKey, value interface{}) {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateUser, "updateUser")
}

func (q *MemberAddressCQ) getCValueVersionNo() *df.ConditionValue {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	return q.VersionNo
}



func (q *MemberAddressCQ) SetVersionNo_Equal(value int64) *MemberAddressCQ {
	q.regVersionNo(df.CK_EQ_C, value)
	return q
}
func (q *MemberAddressCQ) SetVersionNo_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueVersionNo(), "versionNo")
}
func (q *MemberAddressCQ) SetVersionNo_NotEqual(value int64) *MemberAddressCQ {
	q.regVersionNo(df.CK_NE_C, value)
	return q
}

func (q *MemberAddressCQ) SetVersionNo_GreaterThan(value int64) *MemberAddressCQ {
	q.regVersionNo(df.CK_GT_C, value)
	return q
}

func (q *MemberAddressCQ) SetVersionNo_LessThan(value int64) *MemberAddressCQ {
	q.regVersionNo(df.CK_LT_C, value)
	return q
}

func (q *MemberAddressCQ) SetVersionNo_GreaterEqual(value int64) *MemberAddressCQ {
	q.regVersionNo(df.CK_GE_C, value)
	return q
}

func (q *MemberAddressCQ) SetVersionNo_LessEqual(value int64) *MemberAddressCQ {
	q.regVersionNo(df.CK_LE_C, value)
	return q
}
func (q *MemberAddressCQ) SetVersionNo_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueVersionNo(),"versionNo",rangeOfOption)
}	


func (q *MemberAddressCQ) AddOrderBy_VersionNo_Asc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBA("versionNo")
	return q
}
func (q *MemberAddressCQ) AddOrderBy_VersionNo_Desc() *MemberAddressCQ {
	q.BaseConditionQuery.RegOBD("versionNo")
	return q
}
func (q *MemberAddressCQ) regVersionNo(key *df.ConditionKey, value interface{}) {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.VersionNo, "versionNo")
}


func (q *MemberAddressCQ) QueryMember() *MemberCQ {
	if q.conditionQueryMember == nil {
		q.conditionQueryMember = q.xcreateQueryMember()
		q.xsetupOuterJoinMember()
	}
	return q.conditionQueryMember
}

func (q *MemberAddressCQ) xcreateQueryMember() *MemberCQ {
	nrp := q.BaseConditionQuery.ResolveNextRelationPath("MemberAddress", "Member")
	jan := q.BaseConditionQuery.ResolveJoinAliasName(nrp)
	var basecq df.ConditionQuery = q
	cq := CreateMemberCQ(&basecq, q.BaseConditionQuery.SqlClause, jan, q.BaseConditionQuery.NestLevel+1)
	cq.BaseConditionQuery.BaseCB = q.BaseConditionQuery.BaseCB
	cq.BaseConditionQuery.ForeignPropertyName = "Member"
	cq.BaseConditionQuery.RelationPath = nrp
	return cq
}
func (q *MemberAddressCQ) xsetupOuterJoinMember() {
	    cq := q.QueryMember()
        joinOnMap := make(map[string]string)
        joinOnMap["memberId"]="memberId"
        q.BaseConditionQuery.RegisterOuterJoin(
        	cq.BaseConditionQuery.ConditionQuery, joinOnMap, "Member");
}	
	
func (q *MemberAddressCQ) QueryRegion() *RegionCQ {
	if q.conditionQueryRegion == nil {
		q.conditionQueryRegion = q.xcreateQueryRegion()
		q.xsetupOuterJoinRegion()
	}
	return q.conditionQueryRegion
}

func (q *MemberAddressCQ) xcreateQueryRegion() *RegionCQ {
	nrp := q.BaseConditionQuery.ResolveNextRelationPath("MemberAddress", "Region")
	jan := q.BaseConditionQuery.ResolveJoinAliasName(nrp)
	var basecq df.ConditionQuery = q
	cq := CreateRegionCQ(&basecq, q.BaseConditionQuery.SqlClause, jan, q.BaseConditionQuery.NestLevel+1)
	cq.BaseConditionQuery.BaseCB = q.BaseConditionQuery.BaseCB
	cq.BaseConditionQuery.ForeignPropertyName = "Region"
	cq.BaseConditionQuery.RelationPath = nrp
	return cq
}
func (q *MemberAddressCQ) xsetupOuterJoinRegion() {
	    cq := q.QueryRegion()
        joinOnMap := make(map[string]string)
        joinOnMap["regionId"]="regionId"
        q.BaseConditionQuery.RegisterOuterJoin(
        	cq.BaseConditionQuery.ConditionQuery, joinOnMap, "Region");
}	
	
func CreateMemberAddressCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *MemberAddressCQ {
	cq := new(MemberAddressCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "MemberAddress"
	cq.BaseConditionQuery.ReferrerQuery = referrerQuery
	cq.BaseConditionQuery.SqlClause = sqlClause
	cq.BaseConditionQuery.AliasName = aliasName
	cq.BaseConditionQuery.NestLevel = nestlevel
	cq.BaseConditionQuery.DBMetaProvider = df.DBMetaProvider_I
	cq.BaseConditionQuery.CQ_PROPERTY = "Query"
	var cqi df.ConditionQuery = cq
	cq.BaseConditionQuery.ConditionQuery=&cqi
	return cq
}	
package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type ProductCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	ProductId *df.ConditionValue
	ProductName *df.ConditionValue
	ProductHandleCode *df.ConditionValue
	ProductCategoryCode *df.ConditionValue
	ProductStatusCode *df.ConditionValue
	RegularPrice *df.ConditionValue
	RegisterDatetime *df.ConditionValue
	RegisterUser *df.ConditionValue
	UpdateDatetime *df.ConditionValue
	UpdateUser *df.ConditionValue
	VersionNo *df.ConditionValue
}

func (q *ProductCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *ProductCQ) getCValueProductId() *df.ConditionValue {
	if q.ProductId == nil {
		q.ProductId = new(df.ConditionValue)
	}
	return q.ProductId
}



func (q *ProductCQ) SetProductId_Equal(value int64) *ProductCQ {
	q.regProductId(df.CK_EQ_C, value)
	return q
}

func (q *ProductCQ) SetProductId_NotEqual(value int64) *ProductCQ {
	q.regProductId(df.CK_NE_C, value)
	return q
}

func (q *ProductCQ) SetProductId_GreaterThan(value int64) *ProductCQ {
	q.regProductId(df.CK_GT_C, value)
	return q
}

func (q *ProductCQ) SetProductId_LessThan(value int64) *ProductCQ {
	q.regProductId(df.CK_LT_C, value)
	return q
}

func (q *ProductCQ) SetProductId_GreaterEqual(value int64) *ProductCQ {
	q.regProductId(df.CK_GE_C, value)
	return q
}

func (q *ProductCQ) SetProductId_LessEqual(value int64) *ProductCQ {
	q.regProductId(df.CK_LE_C, value)
	return q
}
func (q *ProductCQ) SetProductId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueProductId(),"productId",rangeOfOption)
}	


func (q *ProductCQ) SetProductId_IsNull() *ProductCQ {
	q.regProductId(df.CK_ISN_C, 0)
	return q
}
func (q *ProductCQ) SetProductId_IsNotNull() *ProductCQ {
	q.regProductId(df.CK_ISNN_C, 0)
	return q
}
func (q *ProductCQ) AddOrderBy_ProductId_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("productId")
	return q
}
func (q *ProductCQ) AddOrderBy_ProductId_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("productId")
	return q
}
func (q *ProductCQ) regProductId(key *df.ConditionKey, value interface{}) {
	if q.ProductId == nil {
		q.ProductId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ProductId, "productId")
}

func (q *ProductCQ) getCValueProductName() *df.ConditionValue {
	if q.ProductName == nil {
		q.ProductName = new(df.ConditionValue)
	}
	return q.ProductName
}


func (q *ProductCQ) SetProductName_Equal(value string) *ProductCQ {
	q.regProductName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *ProductCQ) SetProductName_NotEqual(value string) *ProductCQ {
	q.regProductName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetProductName_GreaterThan(value string) *ProductCQ {
	q.regProductName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetProductName_LessThan(value string) *ProductCQ {
	q.regProductName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetProductName_GreaterEqualThan(value string) *ProductCQ {
	q.regProductName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ProductCQ) SetProductName_LessEqualThan(value string) *ProductCQ {
	q.regProductName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetProductName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueProductName(), "productName", option)
}

func (q *ProductCQ) SetProductName_PrefixSearch(value string) error {
	return q.SetProductName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ProductCQ) SetProductName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueProductName(), "productName", option)
}



func (q *ProductCQ) AddOrderBy_ProductName_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("productName")
	return q
}
func (q *ProductCQ) AddOrderBy_ProductName_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("productName")
	return q
}
func (q *ProductCQ) regProductName(key *df.ConditionKey, value interface{}) {
	if q.ProductName == nil {
		q.ProductName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ProductName, "productName")
}

func (q *ProductCQ) getCValueProductHandleCode() *df.ConditionValue {
	if q.ProductHandleCode == nil {
		q.ProductHandleCode = new(df.ConditionValue)
	}
	return q.ProductHandleCode
}


func (q *ProductCQ) SetProductHandleCode_Equal(value string) *ProductCQ {
	q.regProductHandleCode(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *ProductCQ) SetProductHandleCode_NotEqual(value string) *ProductCQ {
	q.regProductHandleCode(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetProductHandleCode_GreaterThan(value string) *ProductCQ {
	q.regProductHandleCode(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetProductHandleCode_LessThan(value string) *ProductCQ {
	q.regProductHandleCode(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetProductHandleCode_GreaterEqualThan(value string) *ProductCQ {
	q.regProductHandleCode(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ProductCQ) SetProductHandleCode_LessEqualThan(value string) *ProductCQ {
	q.regProductHandleCode(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetProductHandleCode_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueProductHandleCode(), "productHandleCode", option)
}

func (q *ProductCQ) SetProductHandleCode_PrefixSearch(value string) error {
	return q.SetProductHandleCode_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ProductCQ) SetProductHandleCode_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueProductHandleCode(), "productHandleCode", option)
}



func (q *ProductCQ) AddOrderBy_ProductHandleCode_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("productHandleCode")
	return q
}
func (q *ProductCQ) AddOrderBy_ProductHandleCode_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("productHandleCode")
	return q
}
func (q *ProductCQ) regProductHandleCode(key *df.ConditionKey, value interface{}) {
	if q.ProductHandleCode == nil {
		q.ProductHandleCode = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ProductHandleCode, "productHandleCode")
}

func (q *ProductCQ) getCValueProductCategoryCode() *df.ConditionValue {
	if q.ProductCategoryCode == nil {
		q.ProductCategoryCode = new(df.ConditionValue)
	}
	return q.ProductCategoryCode
}


func (q *ProductCQ) SetProductCategoryCode_Equal(value string) *ProductCQ {
	q.regProductCategoryCode(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *ProductCQ) SetProductCategoryCode_NotEqual(value string) *ProductCQ {
	q.regProductCategoryCode(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetProductCategoryCode_GreaterThan(value string) *ProductCQ {
	q.regProductCategoryCode(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetProductCategoryCode_LessThan(value string) *ProductCQ {
	q.regProductCategoryCode(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetProductCategoryCode_GreaterEqualThan(value string) *ProductCQ {
	q.regProductCategoryCode(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ProductCQ) SetProductCategoryCode_LessEqualThan(value string) *ProductCQ {
	q.regProductCategoryCode(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetProductCategoryCode_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueProductCategoryCode(), "productCategoryCode", option)
}

func (q *ProductCQ) SetProductCategoryCode_PrefixSearch(value string) error {
	return q.SetProductCategoryCode_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ProductCQ) SetProductCategoryCode_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueProductCategoryCode(), "productCategoryCode", option)
}



func (q *ProductCQ) AddOrderBy_ProductCategoryCode_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("productCategoryCode")
	return q
}
func (q *ProductCQ) AddOrderBy_ProductCategoryCode_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("productCategoryCode")
	return q
}
func (q *ProductCQ) regProductCategoryCode(key *df.ConditionKey, value interface{}) {
	if q.ProductCategoryCode == nil {
		q.ProductCategoryCode = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ProductCategoryCode, "productCategoryCode")
}

func (q *ProductCQ) getCValueProductStatusCode() *df.ConditionValue {
	if q.ProductStatusCode == nil {
		q.ProductStatusCode = new(df.ConditionValue)
	}
	return q.ProductStatusCode
}


func (q *ProductCQ) SetProductStatusCode_Equal(value string) *ProductCQ {
	q.regProductStatusCode(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *ProductCQ) SetProductStatusCode_NotEqual(value string) *ProductCQ {
	q.regProductStatusCode(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetProductStatusCode_GreaterThan(value string) *ProductCQ {
	q.regProductStatusCode(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetProductStatusCode_LessThan(value string) *ProductCQ {
	q.regProductStatusCode(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetProductStatusCode_GreaterEqualThan(value string) *ProductCQ {
	q.regProductStatusCode(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ProductCQ) SetProductStatusCode_LessEqualThan(value string) *ProductCQ {
	q.regProductStatusCode(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetProductStatusCode_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueProductStatusCode(), "productStatusCode", option)
}

func (q *ProductCQ) SetProductStatusCode_PrefixSearch(value string) error {
	return q.SetProductStatusCode_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ProductCQ) SetProductStatusCode_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueProductStatusCode(), "productStatusCode", option)
}



func (q *ProductCQ) AddOrderBy_ProductStatusCode_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("productStatusCode")
	return q
}
func (q *ProductCQ) AddOrderBy_ProductStatusCode_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("productStatusCode")
	return q
}
func (q *ProductCQ) regProductStatusCode(key *df.ConditionKey, value interface{}) {
	if q.ProductStatusCode == nil {
		q.ProductStatusCode = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ProductStatusCode, "productStatusCode")
}

func (q *ProductCQ) getCValueRegularPrice() *df.ConditionValue {
	if q.RegularPrice == nil {
		q.RegularPrice = new(df.ConditionValue)
	}
	return q.RegularPrice
}



func (q *ProductCQ) SetRegularPrice_Equal(value int64) *ProductCQ {
	q.regRegularPrice(df.CK_EQ_C, value)
	return q
}

func (q *ProductCQ) SetRegularPrice_NotEqual(value int64) *ProductCQ {
	q.regRegularPrice(df.CK_NE_C, value)
	return q
}

func (q *ProductCQ) SetRegularPrice_GreaterThan(value int64) *ProductCQ {
	q.regRegularPrice(df.CK_GT_C, value)
	return q
}

func (q *ProductCQ) SetRegularPrice_LessThan(value int64) *ProductCQ {
	q.regRegularPrice(df.CK_LT_C, value)
	return q
}

func (q *ProductCQ) SetRegularPrice_GreaterEqual(value int64) *ProductCQ {
	q.regRegularPrice(df.CK_GE_C, value)
	return q
}

func (q *ProductCQ) SetRegularPrice_LessEqual(value int64) *ProductCQ {
	q.regRegularPrice(df.CK_LE_C, value)
	return q
}
func (q *ProductCQ) SetRegularPrice_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueRegularPrice(),"regularPrice",rangeOfOption)
}	


func (q *ProductCQ) AddOrderBy_RegularPrice_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("regularPrice")
	return q
}
func (q *ProductCQ) AddOrderBy_RegularPrice_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("regularPrice")
	return q
}
func (q *ProductCQ) regRegularPrice(key *df.ConditionKey, value interface{}) {
	if q.RegularPrice == nil {
		q.RegularPrice = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegularPrice, "regularPrice")
}

func (q *ProductCQ) getCValueRegisterDatetime() *df.ConditionValue {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	return q.RegisterDatetime
}




func (q *ProductCQ) SetRegisterDatetime_Equal(value df.MysqlTimestamp) *ProductCQ {
	q.regRegisterDatetime(df.CK_EQ_C, value)
	return q
}


func (q *ProductCQ) SetRegisterDatetime_GreaterThan(value df.MysqlTimestamp) *ProductCQ {
	q.regRegisterDatetime(df.CK_GT_C, value)
	return q
}

func (q *ProductCQ) SetRegisterDatetime_LessThan(value df.MysqlTimestamp) *ProductCQ {
	q.regRegisterDatetime(df.CK_LT_C, value)
	return q
}

func (q *ProductCQ) SetRegisterDatetime_GreaterEqual(value df.MysqlTimestamp) *ProductCQ {
	q.regRegisterDatetime(df.CK_GE_C, value)
	return q
}

func (q *ProductCQ) SetRegisterDatetime_LessEqual(value df.MysqlTimestamp) *ProductCQ {
	q.regRegisterDatetime(df.CK_LE_C, value)
	return q
}

func (q *ProductCQ) AddOrderBy_RegisterDatetime_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("registerDatetime")
	return q
}
func (q *ProductCQ) AddOrderBy_RegisterDatetime_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("registerDatetime")
	return q
}
func (q *ProductCQ) regRegisterDatetime(key *df.ConditionKey, value interface{}) {
	if q.RegisterDatetime == nil {
		q.RegisterDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterDatetime, "registerDatetime")
}

func (q *ProductCQ) getCValueRegisterUser() *df.ConditionValue {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	return q.RegisterUser
}


func (q *ProductCQ) SetRegisterUser_Equal(value string) *ProductCQ {
	q.regRegisterUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *ProductCQ) SetRegisterUser_NotEqual(value string) *ProductCQ {
	q.regRegisterUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetRegisterUser_GreaterThan(value string) *ProductCQ {
	q.regRegisterUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetRegisterUser_LessThan(value string) *ProductCQ {
	q.regRegisterUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetRegisterUser_GreaterEqualThan(value string) *ProductCQ {
	q.regRegisterUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ProductCQ) SetRegisterUser_LessEqualThan(value string) *ProductCQ {
	q.regRegisterUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetRegisterUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}

func (q *ProductCQ) SetRegisterUser_PrefixSearch(value string) error {
	return q.SetRegisterUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ProductCQ) SetRegisterUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueRegisterUser(), "registerUser", option)
}



func (q *ProductCQ) AddOrderBy_RegisterUser_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("registerUser")
	return q
}
func (q *ProductCQ) AddOrderBy_RegisterUser_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("registerUser")
	return q
}
func (q *ProductCQ) regRegisterUser(key *df.ConditionKey, value interface{}) {
	if q.RegisterUser == nil {
		q.RegisterUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.RegisterUser, "registerUser")
}

func (q *ProductCQ) getCValueUpdateDatetime() *df.ConditionValue {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	return q.UpdateDatetime
}




func (q *ProductCQ) SetUpdateDatetime_Equal(value df.MysqlTimestamp) *ProductCQ {
	q.regUpdateDatetime(df.CK_EQ_C, value)
	return q
}


func (q *ProductCQ) SetUpdateDatetime_GreaterThan(value df.MysqlTimestamp) *ProductCQ {
	q.regUpdateDatetime(df.CK_GT_C, value)
	return q
}

func (q *ProductCQ) SetUpdateDatetime_LessThan(value df.MysqlTimestamp) *ProductCQ {
	q.regUpdateDatetime(df.CK_LT_C, value)
	return q
}

func (q *ProductCQ) SetUpdateDatetime_GreaterEqual(value df.MysqlTimestamp) *ProductCQ {
	q.regUpdateDatetime(df.CK_GE_C, value)
	return q
}

func (q *ProductCQ) SetUpdateDatetime_LessEqual(value df.MysqlTimestamp) *ProductCQ {
	q.regUpdateDatetime(df.CK_LE_C, value)
	return q
}

func (q *ProductCQ) AddOrderBy_UpdateDatetime_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("updateDatetime")
	return q
}
func (q *ProductCQ) AddOrderBy_UpdateDatetime_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("updateDatetime")
	return q
}
func (q *ProductCQ) regUpdateDatetime(key *df.ConditionKey, value interface{}) {
	if q.UpdateDatetime == nil {
		q.UpdateDatetime = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateDatetime, "updateDatetime")
}

func (q *ProductCQ) getCValueUpdateUser() *df.ConditionValue {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	return q.UpdateUser
}


func (q *ProductCQ) SetUpdateUser_Equal(value string) *ProductCQ {
	q.regUpdateUser(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *ProductCQ) SetUpdateUser_NotEqual(value string) *ProductCQ {
	q.regUpdateUser(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetUpdateUser_GreaterThan(value string) *ProductCQ {
	q.regUpdateUser(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetUpdateUser_LessThan(value string) *ProductCQ {
	q.regUpdateUser(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetUpdateUser_GreaterEqualThan(value string) *ProductCQ {
	q.regUpdateUser(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ProductCQ) SetUpdateUser_LessEqualThan(value string) *ProductCQ {
	q.regUpdateUser(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCQ) SetUpdateUser_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}

func (q *ProductCQ) SetUpdateUser_PrefixSearch(value string) error {
	return q.SetUpdateUser_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ProductCQ) SetUpdateUser_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueUpdateUser(), "updateUser", option)
}



func (q *ProductCQ) AddOrderBy_UpdateUser_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("updateUser")
	return q
}
func (q *ProductCQ) AddOrderBy_UpdateUser_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("updateUser")
	return q
}
func (q *ProductCQ) regUpdateUser(key *df.ConditionKey, value interface{}) {
	if q.UpdateUser == nil {
		q.UpdateUser = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.UpdateUser, "updateUser")
}

func (q *ProductCQ) getCValueVersionNo() *df.ConditionValue {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	return q.VersionNo
}



func (q *ProductCQ) SetVersionNo_Equal(value int64) *ProductCQ {
	q.regVersionNo(df.CK_EQ_C, value)
	return q
}

func (q *ProductCQ) SetVersionNo_NotEqual(value int64) *ProductCQ {
	q.regVersionNo(df.CK_NE_C, value)
	return q
}

func (q *ProductCQ) SetVersionNo_GreaterThan(value int64) *ProductCQ {
	q.regVersionNo(df.CK_GT_C, value)
	return q
}

func (q *ProductCQ) SetVersionNo_LessThan(value int64) *ProductCQ {
	q.regVersionNo(df.CK_LT_C, value)
	return q
}

func (q *ProductCQ) SetVersionNo_GreaterEqual(value int64) *ProductCQ {
	q.regVersionNo(df.CK_GE_C, value)
	return q
}

func (q *ProductCQ) SetVersionNo_LessEqual(value int64) *ProductCQ {
	q.regVersionNo(df.CK_LE_C, value)
	return q
}
func (q *ProductCQ) SetVersionNo_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueVersionNo(),"versionNo",rangeOfOption)
}	


func (q *ProductCQ) AddOrderBy_VersionNo_Asc() *ProductCQ {
	q.BaseConditionQuery.RegOBA("versionNo")
	return q
}
func (q *ProductCQ) AddOrderBy_VersionNo_Desc() *ProductCQ {
	q.BaseConditionQuery.RegOBD("versionNo")
	return q
}
func (q *ProductCQ) regVersionNo(key *df.ConditionKey, value interface{}) {
	if q.VersionNo == nil {
		q.VersionNo = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.VersionNo, "versionNo")
}


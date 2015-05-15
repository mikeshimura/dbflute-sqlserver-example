package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type WhiteDelimiterCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	DelimiterId *df.ConditionValue
	NumberNullable *df.ConditionValue
	StringConverted *df.ConditionValue
	StringNonConverted *df.ConditionValue
	DateDefault *df.ConditionValue
}

func (q *WhiteDelimiterCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *WhiteDelimiterCQ) getCValueDelimiterId() *df.ConditionValue {
	if q.DelimiterId == nil {
		q.DelimiterId = new(df.ConditionValue)
	}
	return q.DelimiterId
}



func (q *WhiteDelimiterCQ) SetDelimiterId_Equal(value int64) *WhiteDelimiterCQ {
	q.regDelimiterId(df.CK_EQ_C, value)
	return q
}
func (q *WhiteDelimiterCQ) SetDelimiterId_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueDelimiterId(), "delimiterId")
}
func (q *WhiteDelimiterCQ) SetDelimiterId_NotEqual(value int64) *WhiteDelimiterCQ {
	q.regDelimiterId(df.CK_NE_C, value)
	return q
}

func (q *WhiteDelimiterCQ) SetDelimiterId_GreaterThan(value int64) *WhiteDelimiterCQ {
	q.regDelimiterId(df.CK_GT_C, value)
	return q
}

func (q *WhiteDelimiterCQ) SetDelimiterId_LessThan(value int64) *WhiteDelimiterCQ {
	q.regDelimiterId(df.CK_LT_C, value)
	return q
}

func (q *WhiteDelimiterCQ) SetDelimiterId_GreaterEqual(value int64) *WhiteDelimiterCQ {
	q.regDelimiterId(df.CK_GE_C, value)
	return q
}

func (q *WhiteDelimiterCQ) SetDelimiterId_LessEqual(value int64) *WhiteDelimiterCQ {
	q.regDelimiterId(df.CK_LE_C, value)
	return q
}
func (q *WhiteDelimiterCQ) SetDelimiterId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueDelimiterId(),"delimiterId",rangeOfOption)
}	


func (q *WhiteDelimiterCQ) SetDelimiterId_IsNull() *WhiteDelimiterCQ {
	q.regDelimiterId(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteDelimiterCQ) SetDelimiterId_IsNotNull() *WhiteDelimiterCQ {
	q.regDelimiterId(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteDelimiterCQ) AddOrderBy_DelimiterId_Asc() *WhiteDelimiterCQ {
	q.BaseConditionQuery.RegOBA("delimiterId")
	return q
}
func (q *WhiteDelimiterCQ) AddOrderBy_DelimiterId_Desc() *WhiteDelimiterCQ {
	q.BaseConditionQuery.RegOBD("delimiterId")
	return q
}
func (q *WhiteDelimiterCQ) regDelimiterId(key *df.ConditionKey, value interface{}) {
	if q.DelimiterId == nil {
		q.DelimiterId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DelimiterId, "delimiterId")
}

func (q *WhiteDelimiterCQ) getCValueNumberNullable() *df.ConditionValue {
	if q.NumberNullable == nil {
		q.NumberNullable = new(df.ConditionValue)
	}
	return q.NumberNullable
}



func (q *WhiteDelimiterCQ) SetNumberNullable_Equal(value int64) *WhiteDelimiterCQ {
	q.regNumberNullable(df.CK_EQ_C, value)
	return q
}
func (q *WhiteDelimiterCQ) SetNumberNullable_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueNumberNullable(), "numberNullable")
}
func (q *WhiteDelimiterCQ) SetNumberNullable_NotEqual(value int64) *WhiteDelimiterCQ {
	q.regNumberNullable(df.CK_NE_C, value)
	return q
}

func (q *WhiteDelimiterCQ) SetNumberNullable_GreaterThan(value int64) *WhiteDelimiterCQ {
	q.regNumberNullable(df.CK_GT_C, value)
	return q
}

func (q *WhiteDelimiterCQ) SetNumberNullable_LessThan(value int64) *WhiteDelimiterCQ {
	q.regNumberNullable(df.CK_LT_C, value)
	return q
}

func (q *WhiteDelimiterCQ) SetNumberNullable_GreaterEqual(value int64) *WhiteDelimiterCQ {
	q.regNumberNullable(df.CK_GE_C, value)
	return q
}

func (q *WhiteDelimiterCQ) SetNumberNullable_LessEqual(value int64) *WhiteDelimiterCQ {
	q.regNumberNullable(df.CK_LE_C, value)
	return q
}
func (q *WhiteDelimiterCQ) SetNumberNullable_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueNumberNullable(),"numberNullable",rangeOfOption)
}	


func (q *WhiteDelimiterCQ) SetNumberNullable_IsNull() *WhiteDelimiterCQ {
	q.regNumberNullable(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteDelimiterCQ) SetNumberNullable_IsNotNull() *WhiteDelimiterCQ {
	q.regNumberNullable(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteDelimiterCQ) AddOrderBy_NumberNullable_Asc() *WhiteDelimiterCQ {
	q.BaseConditionQuery.RegOBA("numberNullable")
	return q
}
func (q *WhiteDelimiterCQ) AddOrderBy_NumberNullable_Desc() *WhiteDelimiterCQ {
	q.BaseConditionQuery.RegOBD("numberNullable")
	return q
}
func (q *WhiteDelimiterCQ) regNumberNullable(key *df.ConditionKey, value interface{}) {
	if q.NumberNullable == nil {
		q.NumberNullable = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.NumberNullable, "numberNullable")
}

func (q *WhiteDelimiterCQ) getCValueStringConverted() *df.ConditionValue {
	if q.StringConverted == nil {
		q.StringConverted = new(df.ConditionValue)
	}
	return q.StringConverted
}


func (q *WhiteDelimiterCQ) SetStringConverted_Equal(value string) *WhiteDelimiterCQ {
	q.regStringConverted(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *WhiteDelimiterCQ) SetStringConverted_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueStringConverted(), "stringConverted")
}
func (q *WhiteDelimiterCQ) SetStringConverted_NotEqual(value string) *WhiteDelimiterCQ {
	q.regStringConverted(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteDelimiterCQ) SetStringConverted_GreaterThan(value string) *WhiteDelimiterCQ {
	q.regStringConverted(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteDelimiterCQ) SetStringConverted_LessThan(value string) *WhiteDelimiterCQ {
	q.regStringConverted(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteDelimiterCQ) SetStringConverted_GreaterEqualThan(value string) *WhiteDelimiterCQ {
	q.regStringConverted(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *WhiteDelimiterCQ) SetStringConverted_LessEqualThan(value string) *WhiteDelimiterCQ {
	q.regStringConverted(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteDelimiterCQ) SetStringConverted_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueStringConverted(), "stringConverted", option)
}

func (q *WhiteDelimiterCQ) SetStringConverted_PrefixSearch(value string) error {
	return q.SetStringConverted_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *WhiteDelimiterCQ) SetStringConverted_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueStringConverted(), "stringConverted", option)
}



func (q *WhiteDelimiterCQ) AddOrderBy_StringConverted_Asc() *WhiteDelimiterCQ {
	q.BaseConditionQuery.RegOBA("stringConverted")
	return q
}
func (q *WhiteDelimiterCQ) AddOrderBy_StringConverted_Desc() *WhiteDelimiterCQ {
	q.BaseConditionQuery.RegOBD("stringConverted")
	return q
}
func (q *WhiteDelimiterCQ) regStringConverted(key *df.ConditionKey, value interface{}) {
	if q.StringConverted == nil {
		q.StringConverted = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.StringConverted, "stringConverted")
}

func (q *WhiteDelimiterCQ) getCValueStringNonConverted() *df.ConditionValue {
	if q.StringNonConverted == nil {
		q.StringNonConverted = new(df.ConditionValue)
	}
	return q.StringNonConverted
}


func (q *WhiteDelimiterCQ) SetStringNonConverted_Equal(value string) *WhiteDelimiterCQ {
	q.regStringNonConverted(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *WhiteDelimiterCQ) SetStringNonConverted_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueStringNonConverted(), "stringNonConverted")
}
func (q *WhiteDelimiterCQ) SetStringNonConverted_NotEqual(value string) *WhiteDelimiterCQ {
	q.regStringNonConverted(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteDelimiterCQ) SetStringNonConverted_GreaterThan(value string) *WhiteDelimiterCQ {
	q.regStringNonConverted(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteDelimiterCQ) SetStringNonConverted_LessThan(value string) *WhiteDelimiterCQ {
	q.regStringNonConverted(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteDelimiterCQ) SetStringNonConverted_GreaterEqualThan(value string) *WhiteDelimiterCQ {
	q.regStringNonConverted(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *WhiteDelimiterCQ) SetStringNonConverted_LessEqualThan(value string) *WhiteDelimiterCQ {
	q.regStringNonConverted(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *WhiteDelimiterCQ) SetStringNonConverted_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueStringNonConverted(), "stringNonConverted", option)
}

func (q *WhiteDelimiterCQ) SetStringNonConverted_PrefixSearch(value string) error {
	return q.SetStringNonConverted_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *WhiteDelimiterCQ) SetStringNonConverted_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueStringNonConverted(), "stringNonConverted", option)
}



func (q *WhiteDelimiterCQ) SetStringNonConverted_IsNull() *WhiteDelimiterCQ {
	q.regStringNonConverted(df.CK_ISN_C, 0)
	return q
}
func (q *WhiteDelimiterCQ) SetStringNonConverted_IsNullOrEmpty() *WhiteDelimiterCQ {
	q.regStringNonConverted(df.CK_ISNOE_C, 0)
	return q
}
func (q *WhiteDelimiterCQ) SetStringNonConverted_IsNotNull() *WhiteDelimiterCQ {
	q.regStringNonConverted(df.CK_ISNN_C, 0)
	return q
}
func (q *WhiteDelimiterCQ) AddOrderBy_StringNonConverted_Asc() *WhiteDelimiterCQ {
	q.BaseConditionQuery.RegOBA("stringNonConverted")
	return q
}
func (q *WhiteDelimiterCQ) AddOrderBy_StringNonConverted_Desc() *WhiteDelimiterCQ {
	q.BaseConditionQuery.RegOBD("stringNonConverted")
	return q
}
func (q *WhiteDelimiterCQ) regStringNonConverted(key *df.ConditionKey, value interface{}) {
	if q.StringNonConverted == nil {
		q.StringNonConverted = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.StringNonConverted, "stringNonConverted")
}

func (q *WhiteDelimiterCQ) getCValueDateDefault() *df.ConditionValue {
	if q.DateDefault == nil {
		q.DateDefault = new(df.ConditionValue)
	}
	return q.DateDefault
}




func (q *WhiteDelimiterCQ) SetDateDefault_Equal(value df.Timestamp) *WhiteDelimiterCQ {
	q.regDateDefault(df.CK_EQ_C, value)
	return q
}


func (q *WhiteDelimiterCQ) SetDateDefault_GreaterThan(value df.Timestamp) *WhiteDelimiterCQ {
	q.regDateDefault(df.CK_GT_C, value)
	return q
}

func (q *WhiteDelimiterCQ) SetDateDefault_LessThan(value df.Timestamp) *WhiteDelimiterCQ {
	q.regDateDefault(df.CK_LT_C, value)
	return q
}

func (q *WhiteDelimiterCQ) SetDateDefault_GreaterEqual(value df.Timestamp) *WhiteDelimiterCQ {
	q.regDateDefault(df.CK_GE_C, value)
	return q
}

func (q *WhiteDelimiterCQ) SetDateDefault_LessEqual(value df.Timestamp) *WhiteDelimiterCQ {
	q.regDateDefault(df.CK_LE_C, value)
	return q
}

func (q *WhiteDelimiterCQ) AddOrderBy_DateDefault_Asc() *WhiteDelimiterCQ {
	q.BaseConditionQuery.RegOBA("dateDefault")
	return q
}
func (q *WhiteDelimiterCQ) AddOrderBy_DateDefault_Desc() *WhiteDelimiterCQ {
	q.BaseConditionQuery.RegOBD("dateDefault")
	return q
}
func (q *WhiteDelimiterCQ) regDateDefault(key *df.ConditionKey, value interface{}) {
	if q.DateDefault == nil {
		q.DateDefault = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.DateDefault, "dateDefault")
}


func CreateWhiteDelimiterCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *WhiteDelimiterCQ {
	cq := new(WhiteDelimiterCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "WhiteDelimiter"
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
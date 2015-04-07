package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorTheLongAndWindingTableAndColumnRefCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	TheLongAndWindingTableAndColumnRefId *df.ConditionValue
	TheLongAndWindingTableAndColumnId *df.ConditionValue
	TheLongAndWindingTableAndColumnRefDate *df.ConditionValue
	ShortDate *df.ConditionValue
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *VendorTheLongAndWindingTableAndColumnRefCQ) getCValueTheLongAndWindingTableAndColumnRefId() *df.ConditionValue {
	if q.TheLongAndWindingTableAndColumnRefId == nil {
		q.TheLongAndWindingTableAndColumnRefId = new(df.ConditionValue)
	}
	return q.TheLongAndWindingTableAndColumnRefId
}



func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetTheLongAndWindingTableAndColumnRefId_Equal(value int64) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regTheLongAndWindingTableAndColumnRefId(df.CK_EQ_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetTheLongAndWindingTableAndColumnRefId_NotEqual(value int64) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regTheLongAndWindingTableAndColumnRefId(df.CK_NE_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetTheLongAndWindingTableAndColumnRefId_GreaterThan(value int64) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regTheLongAndWindingTableAndColumnRefId(df.CK_GT_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetTheLongAndWindingTableAndColumnRefId_LessThan(value int64) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regTheLongAndWindingTableAndColumnRefId(df.CK_LT_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetTheLongAndWindingTableAndColumnRefId_GreaterEqual(value int64) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regTheLongAndWindingTableAndColumnRefId(df.CK_GE_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetTheLongAndWindingTableAndColumnRefId_LessEqual(value int64) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regTheLongAndWindingTableAndColumnRefId(df.CK_LE_C, value)
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetTheLongAndWindingTableAndColumnRefId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueTheLongAndWindingTableAndColumnRefId(),"theLongAndWindingTableAndColumnRefId",rangeOfOption)
}	


func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetTheLongAndWindingTableAndColumnRefId_IsNull() *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regTheLongAndWindingTableAndColumnRefId(df.CK_ISN_C, 0)
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetTheLongAndWindingTableAndColumnRefId_IsNotNull() *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regTheLongAndWindingTableAndColumnRefId(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnRefCQ) AddOrderBy_TheLongAndWindingTableAndColumnRefId_Asc() *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.BaseConditionQuery.RegOBA("theLongAndWindingTableAndColumnRefId")
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnRefCQ) AddOrderBy_TheLongAndWindingTableAndColumnRefId_Desc() *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.BaseConditionQuery.RegOBD("theLongAndWindingTableAndColumnRefId")
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnRefCQ) regTheLongAndWindingTableAndColumnRefId(key *df.ConditionKey, value interface{}) {
	if q.TheLongAndWindingTableAndColumnRefId == nil {
		q.TheLongAndWindingTableAndColumnRefId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.TheLongAndWindingTableAndColumnRefId, "theLongAndWindingTableAndColumnRefId")
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) getCValueTheLongAndWindingTableAndColumnId() *df.ConditionValue {
	if q.TheLongAndWindingTableAndColumnId == nil {
		q.TheLongAndWindingTableAndColumnId = new(df.ConditionValue)
	}
	return q.TheLongAndWindingTableAndColumnId
}



func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetTheLongAndWindingTableAndColumnId_Equal(value int64) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regTheLongAndWindingTableAndColumnId(df.CK_EQ_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetTheLongAndWindingTableAndColumnId_NotEqual(value int64) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regTheLongAndWindingTableAndColumnId(df.CK_NE_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetTheLongAndWindingTableAndColumnId_GreaterThan(value int64) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regTheLongAndWindingTableAndColumnId(df.CK_GT_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetTheLongAndWindingTableAndColumnId_LessThan(value int64) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regTheLongAndWindingTableAndColumnId(df.CK_LT_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetTheLongAndWindingTableAndColumnId_GreaterEqual(value int64) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regTheLongAndWindingTableAndColumnId(df.CK_GE_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetTheLongAndWindingTableAndColumnId_LessEqual(value int64) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regTheLongAndWindingTableAndColumnId(df.CK_LE_C, value)
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetTheLongAndWindingTableAndColumnId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueTheLongAndWindingTableAndColumnId(),"theLongAndWindingTableAndColumnId",rangeOfOption)
}	


func (q *VendorTheLongAndWindingTableAndColumnRefCQ) AddOrderBy_TheLongAndWindingTableAndColumnId_Asc() *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.BaseConditionQuery.RegOBA("theLongAndWindingTableAndColumnId")
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnRefCQ) AddOrderBy_TheLongAndWindingTableAndColumnId_Desc() *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.BaseConditionQuery.RegOBD("theLongAndWindingTableAndColumnId")
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnRefCQ) regTheLongAndWindingTableAndColumnId(key *df.ConditionKey, value interface{}) {
	if q.TheLongAndWindingTableAndColumnId == nil {
		q.TheLongAndWindingTableAndColumnId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.TheLongAndWindingTableAndColumnId, "theLongAndWindingTableAndColumnId")
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) getCValueTheLongAndWindingTableAndColumnRefDate() *df.ConditionValue {
	if q.TheLongAndWindingTableAndColumnRefDate == nil {
		q.TheLongAndWindingTableAndColumnRefDate = new(df.ConditionValue)
	}
	return q.TheLongAndWindingTableAndColumnRefDate
}




func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetTheLongAndWindingTableAndColumnRefDate_Equal(value df.MysqlDate) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regTheLongAndWindingTableAndColumnRefDate(df.CK_EQ_C, value)
	return q
}


func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetTheLongAndWindingTableAndColumnRefDate_GreaterThan(value df.MysqlDate) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regTheLongAndWindingTableAndColumnRefDate(df.CK_GT_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetTheLongAndWindingTableAndColumnRefDate_LessThan(value df.MysqlDate) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regTheLongAndWindingTableAndColumnRefDate(df.CK_LT_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetTheLongAndWindingTableAndColumnRefDate_GreaterEqual(value df.MysqlDate) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regTheLongAndWindingTableAndColumnRefDate(df.CK_GE_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetTheLongAndWindingTableAndColumnRefDate_LessEqual(value df.MysqlDate) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regTheLongAndWindingTableAndColumnRefDate(df.CK_LE_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) AddOrderBy_TheLongAndWindingTableAndColumnRefDate_Asc() *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.BaseConditionQuery.RegOBA("theLongAndWindingTableAndColumnRefDate")
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnRefCQ) AddOrderBy_TheLongAndWindingTableAndColumnRefDate_Desc() *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.BaseConditionQuery.RegOBD("theLongAndWindingTableAndColumnRefDate")
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnRefCQ) regTheLongAndWindingTableAndColumnRefDate(key *df.ConditionKey, value interface{}) {
	if q.TheLongAndWindingTableAndColumnRefDate == nil {
		q.TheLongAndWindingTableAndColumnRefDate = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.TheLongAndWindingTableAndColumnRefDate, "theLongAndWindingTableAndColumnRefDate")
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) getCValueShortDate() *df.ConditionValue {
	if q.ShortDate == nil {
		q.ShortDate = new(df.ConditionValue)
	}
	return q.ShortDate
}




func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetShortDate_Equal(value df.MysqlDate) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regShortDate(df.CK_EQ_C, value)
	return q
}


func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetShortDate_GreaterThan(value df.MysqlDate) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regShortDate(df.CK_GT_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetShortDate_LessThan(value df.MysqlDate) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regShortDate(df.CK_LT_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetShortDate_GreaterEqual(value df.MysqlDate) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regShortDate(df.CK_GE_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) SetShortDate_LessEqual(value df.MysqlDate) *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.regShortDate(df.CK_LE_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnRefCQ) AddOrderBy_ShortDate_Asc() *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.BaseConditionQuery.RegOBA("shortDate")
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnRefCQ) AddOrderBy_ShortDate_Desc() *VendorTheLongAndWindingTableAndColumnRefCQ {
	q.BaseConditionQuery.RegOBD("shortDate")
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnRefCQ) regShortDate(key *df.ConditionKey, value interface{}) {
	if q.ShortDate == nil {
		q.ShortDate = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ShortDate, "shortDate")
}


package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorTheLongAndWindingTableAndColumnCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	TheLongAndWindingTableAndColumnId *df.ConditionValue
	TheLongAndWindingTableAndColumnName *df.ConditionValue
	ShortName *df.ConditionValue
	ShortSize *df.ConditionValue
}

func (q *VendorTheLongAndWindingTableAndColumnCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *VendorTheLongAndWindingTableAndColumnCQ) getCValueTheLongAndWindingTableAndColumnId() *df.ConditionValue {
	if q.TheLongAndWindingTableAndColumnId == nil {
		q.TheLongAndWindingTableAndColumnId = new(df.ConditionValue)
	}
	return q.TheLongAndWindingTableAndColumnId
}



func (q *VendorTheLongAndWindingTableAndColumnCQ) SetTheLongAndWindingTableAndColumnId_Equal(value int64) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regTheLongAndWindingTableAndColumnId(df.CK_EQ_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetTheLongAndWindingTableAndColumnId_NotEqual(value int64) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regTheLongAndWindingTableAndColumnId(df.CK_NE_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetTheLongAndWindingTableAndColumnId_GreaterThan(value int64) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regTheLongAndWindingTableAndColumnId(df.CK_GT_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetTheLongAndWindingTableAndColumnId_LessThan(value int64) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regTheLongAndWindingTableAndColumnId(df.CK_LT_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetTheLongAndWindingTableAndColumnId_GreaterEqual(value int64) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regTheLongAndWindingTableAndColumnId(df.CK_GE_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetTheLongAndWindingTableAndColumnId_LessEqual(value int64) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regTheLongAndWindingTableAndColumnId(df.CK_LE_C, value)
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnCQ) SetTheLongAndWindingTableAndColumnId_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueTheLongAndWindingTableAndColumnId(),"theLongAndWindingTableAndColumnId",rangeOfOption)
}	


func (q *VendorTheLongAndWindingTableAndColumnCQ) SetTheLongAndWindingTableAndColumnId_IsNull() *VendorTheLongAndWindingTableAndColumnCQ {
	q.regTheLongAndWindingTableAndColumnId(df.CK_ISN_C, 0)
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnCQ) SetTheLongAndWindingTableAndColumnId_IsNotNull() *VendorTheLongAndWindingTableAndColumnCQ {
	q.regTheLongAndWindingTableAndColumnId(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnCQ) AddOrderBy_TheLongAndWindingTableAndColumnId_Asc() *VendorTheLongAndWindingTableAndColumnCQ {
	q.BaseConditionQuery.RegOBA("theLongAndWindingTableAndColumnId")
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnCQ) AddOrderBy_TheLongAndWindingTableAndColumnId_Desc() *VendorTheLongAndWindingTableAndColumnCQ {
	q.BaseConditionQuery.RegOBD("theLongAndWindingTableAndColumnId")
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnCQ) regTheLongAndWindingTableAndColumnId(key *df.ConditionKey, value interface{}) {
	if q.TheLongAndWindingTableAndColumnId == nil {
		q.TheLongAndWindingTableAndColumnId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.TheLongAndWindingTableAndColumnId, "theLongAndWindingTableAndColumnId")
}

func (q *VendorTheLongAndWindingTableAndColumnCQ) getCValueTheLongAndWindingTableAndColumnName() *df.ConditionValue {
	if q.TheLongAndWindingTableAndColumnName == nil {
		q.TheLongAndWindingTableAndColumnName = new(df.ConditionValue)
	}
	return q.TheLongAndWindingTableAndColumnName
}


func (q *VendorTheLongAndWindingTableAndColumnCQ) SetTheLongAndWindingTableAndColumnName_Equal(value string) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regTheLongAndWindingTableAndColumnName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetTheLongAndWindingTableAndColumnName_NotEqual(value string) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regTheLongAndWindingTableAndColumnName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetTheLongAndWindingTableAndColumnName_GreaterThan(value string) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regTheLongAndWindingTableAndColumnName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetTheLongAndWindingTableAndColumnName_LessThan(value string) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regTheLongAndWindingTableAndColumnName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetTheLongAndWindingTableAndColumnName_GreaterEqualThan(value string) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regTheLongAndWindingTableAndColumnName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *VendorTheLongAndWindingTableAndColumnCQ) SetTheLongAndWindingTableAndColumnName_LessEqualThan(value string) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regTheLongAndWindingTableAndColumnName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetTheLongAndWindingTableAndColumnName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueTheLongAndWindingTableAndColumnName(), "theLongAndWindingTableAndColumnName", option)
}

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetTheLongAndWindingTableAndColumnName_PrefixSearch(value string) error {
	return q.SetTheLongAndWindingTableAndColumnName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetTheLongAndWindingTableAndColumnName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueTheLongAndWindingTableAndColumnName(), "theLongAndWindingTableAndColumnName", option)
}



func (q *VendorTheLongAndWindingTableAndColumnCQ) AddOrderBy_TheLongAndWindingTableAndColumnName_Asc() *VendorTheLongAndWindingTableAndColumnCQ {
	q.BaseConditionQuery.RegOBA("theLongAndWindingTableAndColumnName")
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnCQ) AddOrderBy_TheLongAndWindingTableAndColumnName_Desc() *VendorTheLongAndWindingTableAndColumnCQ {
	q.BaseConditionQuery.RegOBD("theLongAndWindingTableAndColumnName")
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnCQ) regTheLongAndWindingTableAndColumnName(key *df.ConditionKey, value interface{}) {
	if q.TheLongAndWindingTableAndColumnName == nil {
		q.TheLongAndWindingTableAndColumnName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.TheLongAndWindingTableAndColumnName, "theLongAndWindingTableAndColumnName")
}

func (q *VendorTheLongAndWindingTableAndColumnCQ) getCValueShortName() *df.ConditionValue {
	if q.ShortName == nil {
		q.ShortName = new(df.ConditionValue)
	}
	return q.ShortName
}


func (q *VendorTheLongAndWindingTableAndColumnCQ) SetShortName_Equal(value string) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regShortName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetShortName_NotEqual(value string) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regShortName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetShortName_GreaterThan(value string) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regShortName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetShortName_LessThan(value string) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regShortName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetShortName_GreaterEqualThan(value string) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regShortName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *VendorTheLongAndWindingTableAndColumnCQ) SetShortName_LessEqualThan(value string) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regShortName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetShortName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueShortName(), "shortName", option)
}

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetShortName_PrefixSearch(value string) error {
	return q.SetShortName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetShortName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueShortName(), "shortName", option)
}



func (q *VendorTheLongAndWindingTableAndColumnCQ) AddOrderBy_ShortName_Asc() *VendorTheLongAndWindingTableAndColumnCQ {
	q.BaseConditionQuery.RegOBA("shortName")
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnCQ) AddOrderBy_ShortName_Desc() *VendorTheLongAndWindingTableAndColumnCQ {
	q.BaseConditionQuery.RegOBD("shortName")
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnCQ) regShortName(key *df.ConditionKey, value interface{}) {
	if q.ShortName == nil {
		q.ShortName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ShortName, "shortName")
}

func (q *VendorTheLongAndWindingTableAndColumnCQ) getCValueShortSize() *df.ConditionValue {
	if q.ShortSize == nil {
		q.ShortSize = new(df.ConditionValue)
	}
	return q.ShortSize
}



func (q *VendorTheLongAndWindingTableAndColumnCQ) SetShortSize_Equal(value int64) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regShortSize(df.CK_EQ_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetShortSize_NotEqual(value int64) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regShortSize(df.CK_NE_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetShortSize_GreaterThan(value int64) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regShortSize(df.CK_GT_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetShortSize_LessThan(value int64) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regShortSize(df.CK_LT_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetShortSize_GreaterEqual(value int64) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regShortSize(df.CK_GE_C, value)
	return q
}

func (q *VendorTheLongAndWindingTableAndColumnCQ) SetShortSize_LessEqual(value int64) *VendorTheLongAndWindingTableAndColumnCQ {
	q.regShortSize(df.CK_LE_C, value)
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnCQ) SetShortSize_RangeOf(minNumber int64, maxNumber int64, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueShortSize(),"shortSize",rangeOfOption)
}	


func (q *VendorTheLongAndWindingTableAndColumnCQ) AddOrderBy_ShortSize_Asc() *VendorTheLongAndWindingTableAndColumnCQ {
	q.BaseConditionQuery.RegOBA("shortSize")
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnCQ) AddOrderBy_ShortSize_Desc() *VendorTheLongAndWindingTableAndColumnCQ {
	q.BaseConditionQuery.RegOBD("shortSize")
	return q
}
func (q *VendorTheLongAndWindingTableAndColumnCQ) regShortSize(key *df.ConditionKey, value interface{}) {
	if q.ShortSize == nil {
		q.ShortSize = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ShortSize, "shortSize")
}


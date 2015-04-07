package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorConstraintNameAutoBarCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	ConstraintNameAutoBarId *df.ConditionValue
	ConstraintNameAutoBarName *df.ConditionValue
}

func (q *VendorConstraintNameAutoBarCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *VendorConstraintNameAutoBarCQ) getCValueConstraintNameAutoBarId() *df.ConditionValue {
	if q.ConstraintNameAutoBarId == nil {
		q.ConstraintNameAutoBarId = new(df.ConditionValue)
	}
	return q.ConstraintNameAutoBarId
}



func (q *VendorConstraintNameAutoBarCQ) SetConstraintNameAutoBarId_Equal(value df.Numeric) *VendorConstraintNameAutoBarCQ {
	q.regConstraintNameAutoBarId(df.CK_EQ_C, value)
	return q
}

func (q *VendorConstraintNameAutoBarCQ) SetConstraintNameAutoBarId_NotEqual(value df.Numeric) *VendorConstraintNameAutoBarCQ {
	q.regConstraintNameAutoBarId(df.CK_NE_C, value)
	return q
}

func (q *VendorConstraintNameAutoBarCQ) SetConstraintNameAutoBarId_GreaterThan(value df.Numeric) *VendorConstraintNameAutoBarCQ {
	q.regConstraintNameAutoBarId(df.CK_GT_C, value)
	return q
}

func (q *VendorConstraintNameAutoBarCQ) SetConstraintNameAutoBarId_LessThan(value df.Numeric) *VendorConstraintNameAutoBarCQ {
	q.regConstraintNameAutoBarId(df.CK_LT_C, value)
	return q
}

func (q *VendorConstraintNameAutoBarCQ) SetConstraintNameAutoBarId_GreaterEqual(value df.Numeric) *VendorConstraintNameAutoBarCQ {
	q.regConstraintNameAutoBarId(df.CK_GE_C, value)
	return q
}

func (q *VendorConstraintNameAutoBarCQ) SetConstraintNameAutoBarId_LessEqual(value df.Numeric) *VendorConstraintNameAutoBarCQ {
	q.regConstraintNameAutoBarId(df.CK_LE_C, value)
	return q
}
func (q *VendorConstraintNameAutoBarCQ) SetConstraintNameAutoBarId_RangeOf(minNumber df.Numeric, maxNumber df.Numeric, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueConstraintNameAutoBarId(),"constraintNameAutoBarId",rangeOfOption)
}	


func (q *VendorConstraintNameAutoBarCQ) SetConstraintNameAutoBarId_IsNull() *VendorConstraintNameAutoBarCQ {
	q.regConstraintNameAutoBarId(df.CK_ISN_C, 0)
	return q
}
func (q *VendorConstraintNameAutoBarCQ) SetConstraintNameAutoBarId_IsNotNull() *VendorConstraintNameAutoBarCQ {
	q.regConstraintNameAutoBarId(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorConstraintNameAutoBarCQ) AddOrderBy_ConstraintNameAutoBarId_Asc() *VendorConstraintNameAutoBarCQ {
	q.BaseConditionQuery.RegOBA("constraintNameAutoBarId")
	return q
}
func (q *VendorConstraintNameAutoBarCQ) AddOrderBy_ConstraintNameAutoBarId_Desc() *VendorConstraintNameAutoBarCQ {
	q.BaseConditionQuery.RegOBD("constraintNameAutoBarId")
	return q
}
func (q *VendorConstraintNameAutoBarCQ) regConstraintNameAutoBarId(key *df.ConditionKey, value interface{}) {
	if q.ConstraintNameAutoBarId == nil {
		q.ConstraintNameAutoBarId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ConstraintNameAutoBarId, "constraintNameAutoBarId")
}

func (q *VendorConstraintNameAutoBarCQ) getCValueConstraintNameAutoBarName() *df.ConditionValue {
	if q.ConstraintNameAutoBarName == nil {
		q.ConstraintNameAutoBarName = new(df.ConditionValue)
	}
	return q.ConstraintNameAutoBarName
}


func (q *VendorConstraintNameAutoBarCQ) SetConstraintNameAutoBarName_Equal(value string) *VendorConstraintNameAutoBarCQ {
	q.regConstraintNameAutoBarName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *VendorConstraintNameAutoBarCQ) SetConstraintNameAutoBarName_NotEqual(value string) *VendorConstraintNameAutoBarCQ {
	q.regConstraintNameAutoBarName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorConstraintNameAutoBarCQ) SetConstraintNameAutoBarName_GreaterThan(value string) *VendorConstraintNameAutoBarCQ {
	q.regConstraintNameAutoBarName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorConstraintNameAutoBarCQ) SetConstraintNameAutoBarName_LessThan(value string) *VendorConstraintNameAutoBarCQ {
	q.regConstraintNameAutoBarName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorConstraintNameAutoBarCQ) SetConstraintNameAutoBarName_GreaterEqualThan(value string) *VendorConstraintNameAutoBarCQ {
	q.regConstraintNameAutoBarName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *VendorConstraintNameAutoBarCQ) SetConstraintNameAutoBarName_LessEqualThan(value string) *VendorConstraintNameAutoBarCQ {
	q.regConstraintNameAutoBarName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorConstraintNameAutoBarCQ) SetConstraintNameAutoBarName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueConstraintNameAutoBarName(), "constraintNameAutoBarName", option)
}

func (q *VendorConstraintNameAutoBarCQ) SetConstraintNameAutoBarName_PrefixSearch(value string) error {
	return q.SetConstraintNameAutoBarName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *VendorConstraintNameAutoBarCQ) SetConstraintNameAutoBarName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueConstraintNameAutoBarName(), "constraintNameAutoBarName", option)
}



func (q *VendorConstraintNameAutoBarCQ) AddOrderBy_ConstraintNameAutoBarName_Asc() *VendorConstraintNameAutoBarCQ {
	q.BaseConditionQuery.RegOBA("constraintNameAutoBarName")
	return q
}
func (q *VendorConstraintNameAutoBarCQ) AddOrderBy_ConstraintNameAutoBarName_Desc() *VendorConstraintNameAutoBarCQ {
	q.BaseConditionQuery.RegOBD("constraintNameAutoBarName")
	return q
}
func (q *VendorConstraintNameAutoBarCQ) regConstraintNameAutoBarName(key *df.ConditionKey, value interface{}) {
	if q.ConstraintNameAutoBarName == nil {
		q.ConstraintNameAutoBarName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ConstraintNameAutoBarName, "constraintNameAutoBarName")
}


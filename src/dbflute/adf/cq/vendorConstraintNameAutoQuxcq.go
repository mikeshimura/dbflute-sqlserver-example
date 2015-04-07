package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorConstraintNameAutoQuxCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	ConstraintNameAutoQuxId *df.ConditionValue
	ConstraintNameAutoQuxName *df.ConditionValue
}

func (q *VendorConstraintNameAutoQuxCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *VendorConstraintNameAutoQuxCQ) getCValueConstraintNameAutoQuxId() *df.ConditionValue {
	if q.ConstraintNameAutoQuxId == nil {
		q.ConstraintNameAutoQuxId = new(df.ConditionValue)
	}
	return q.ConstraintNameAutoQuxId
}



func (q *VendorConstraintNameAutoQuxCQ) SetConstraintNameAutoQuxId_Equal(value df.Numeric) *VendorConstraintNameAutoQuxCQ {
	q.regConstraintNameAutoQuxId(df.CK_EQ_C, value)
	return q
}

func (q *VendorConstraintNameAutoQuxCQ) SetConstraintNameAutoQuxId_NotEqual(value df.Numeric) *VendorConstraintNameAutoQuxCQ {
	q.regConstraintNameAutoQuxId(df.CK_NE_C, value)
	return q
}

func (q *VendorConstraintNameAutoQuxCQ) SetConstraintNameAutoQuxId_GreaterThan(value df.Numeric) *VendorConstraintNameAutoQuxCQ {
	q.regConstraintNameAutoQuxId(df.CK_GT_C, value)
	return q
}

func (q *VendorConstraintNameAutoQuxCQ) SetConstraintNameAutoQuxId_LessThan(value df.Numeric) *VendorConstraintNameAutoQuxCQ {
	q.regConstraintNameAutoQuxId(df.CK_LT_C, value)
	return q
}

func (q *VendorConstraintNameAutoQuxCQ) SetConstraintNameAutoQuxId_GreaterEqual(value df.Numeric) *VendorConstraintNameAutoQuxCQ {
	q.regConstraintNameAutoQuxId(df.CK_GE_C, value)
	return q
}

func (q *VendorConstraintNameAutoQuxCQ) SetConstraintNameAutoQuxId_LessEqual(value df.Numeric) *VendorConstraintNameAutoQuxCQ {
	q.regConstraintNameAutoQuxId(df.CK_LE_C, value)
	return q
}
func (q *VendorConstraintNameAutoQuxCQ) SetConstraintNameAutoQuxId_RangeOf(minNumber df.Numeric, maxNumber df.Numeric, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueConstraintNameAutoQuxId(),"constraintNameAutoQuxId",rangeOfOption)
}	


func (q *VendorConstraintNameAutoQuxCQ) SetConstraintNameAutoQuxId_IsNull() *VendorConstraintNameAutoQuxCQ {
	q.regConstraintNameAutoQuxId(df.CK_ISN_C, 0)
	return q
}
func (q *VendorConstraintNameAutoQuxCQ) SetConstraintNameAutoQuxId_IsNotNull() *VendorConstraintNameAutoQuxCQ {
	q.regConstraintNameAutoQuxId(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorConstraintNameAutoQuxCQ) AddOrderBy_ConstraintNameAutoQuxId_Asc() *VendorConstraintNameAutoQuxCQ {
	q.BaseConditionQuery.RegOBA("constraintNameAutoQuxId")
	return q
}
func (q *VendorConstraintNameAutoQuxCQ) AddOrderBy_ConstraintNameAutoQuxId_Desc() *VendorConstraintNameAutoQuxCQ {
	q.BaseConditionQuery.RegOBD("constraintNameAutoQuxId")
	return q
}
func (q *VendorConstraintNameAutoQuxCQ) regConstraintNameAutoQuxId(key *df.ConditionKey, value interface{}) {
	if q.ConstraintNameAutoQuxId == nil {
		q.ConstraintNameAutoQuxId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ConstraintNameAutoQuxId, "constraintNameAutoQuxId")
}

func (q *VendorConstraintNameAutoQuxCQ) getCValueConstraintNameAutoQuxName() *df.ConditionValue {
	if q.ConstraintNameAutoQuxName == nil {
		q.ConstraintNameAutoQuxName = new(df.ConditionValue)
	}
	return q.ConstraintNameAutoQuxName
}


func (q *VendorConstraintNameAutoQuxCQ) SetConstraintNameAutoQuxName_Equal(value string) *VendorConstraintNameAutoQuxCQ {
	q.regConstraintNameAutoQuxName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *VendorConstraintNameAutoQuxCQ) SetConstraintNameAutoQuxName_NotEqual(value string) *VendorConstraintNameAutoQuxCQ {
	q.regConstraintNameAutoQuxName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorConstraintNameAutoQuxCQ) SetConstraintNameAutoQuxName_GreaterThan(value string) *VendorConstraintNameAutoQuxCQ {
	q.regConstraintNameAutoQuxName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorConstraintNameAutoQuxCQ) SetConstraintNameAutoQuxName_LessThan(value string) *VendorConstraintNameAutoQuxCQ {
	q.regConstraintNameAutoQuxName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorConstraintNameAutoQuxCQ) SetConstraintNameAutoQuxName_GreaterEqualThan(value string) *VendorConstraintNameAutoQuxCQ {
	q.regConstraintNameAutoQuxName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *VendorConstraintNameAutoQuxCQ) SetConstraintNameAutoQuxName_LessEqualThan(value string) *VendorConstraintNameAutoQuxCQ {
	q.regConstraintNameAutoQuxName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorConstraintNameAutoQuxCQ) SetConstraintNameAutoQuxName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueConstraintNameAutoQuxName(), "constraintNameAutoQuxName", option)
}

func (q *VendorConstraintNameAutoQuxCQ) SetConstraintNameAutoQuxName_PrefixSearch(value string) error {
	return q.SetConstraintNameAutoQuxName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *VendorConstraintNameAutoQuxCQ) SetConstraintNameAutoQuxName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueConstraintNameAutoQuxName(), "constraintNameAutoQuxName", option)
}



func (q *VendorConstraintNameAutoQuxCQ) AddOrderBy_ConstraintNameAutoQuxName_Asc() *VendorConstraintNameAutoQuxCQ {
	q.BaseConditionQuery.RegOBA("constraintNameAutoQuxName")
	return q
}
func (q *VendorConstraintNameAutoQuxCQ) AddOrderBy_ConstraintNameAutoQuxName_Desc() *VendorConstraintNameAutoQuxCQ {
	q.BaseConditionQuery.RegOBD("constraintNameAutoQuxName")
	return q
}
func (q *VendorConstraintNameAutoQuxCQ) regConstraintNameAutoQuxName(key *df.ConditionKey, value interface{}) {
	if q.ConstraintNameAutoQuxName == nil {
		q.ConstraintNameAutoQuxName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ConstraintNameAutoQuxName, "constraintNameAutoQuxName")
}


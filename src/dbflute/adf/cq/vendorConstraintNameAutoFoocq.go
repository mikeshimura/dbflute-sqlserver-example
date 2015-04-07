package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorConstraintNameAutoFooCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	ConstraintNameAutoFooId *df.ConditionValue
	ConstraintNameAutoFooName *df.ConditionValue
}

func (q *VendorConstraintNameAutoFooCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *VendorConstraintNameAutoFooCQ) getCValueConstraintNameAutoFooId() *df.ConditionValue {
	if q.ConstraintNameAutoFooId == nil {
		q.ConstraintNameAutoFooId = new(df.ConditionValue)
	}
	return q.ConstraintNameAutoFooId
}



func (q *VendorConstraintNameAutoFooCQ) SetConstraintNameAutoFooId_Equal(value df.Numeric) *VendorConstraintNameAutoFooCQ {
	q.regConstraintNameAutoFooId(df.CK_EQ_C, value)
	return q
}

func (q *VendorConstraintNameAutoFooCQ) SetConstraintNameAutoFooId_NotEqual(value df.Numeric) *VendorConstraintNameAutoFooCQ {
	q.regConstraintNameAutoFooId(df.CK_NE_C, value)
	return q
}

func (q *VendorConstraintNameAutoFooCQ) SetConstraintNameAutoFooId_GreaterThan(value df.Numeric) *VendorConstraintNameAutoFooCQ {
	q.regConstraintNameAutoFooId(df.CK_GT_C, value)
	return q
}

func (q *VendorConstraintNameAutoFooCQ) SetConstraintNameAutoFooId_LessThan(value df.Numeric) *VendorConstraintNameAutoFooCQ {
	q.regConstraintNameAutoFooId(df.CK_LT_C, value)
	return q
}

func (q *VendorConstraintNameAutoFooCQ) SetConstraintNameAutoFooId_GreaterEqual(value df.Numeric) *VendorConstraintNameAutoFooCQ {
	q.regConstraintNameAutoFooId(df.CK_GE_C, value)
	return q
}

func (q *VendorConstraintNameAutoFooCQ) SetConstraintNameAutoFooId_LessEqual(value df.Numeric) *VendorConstraintNameAutoFooCQ {
	q.regConstraintNameAutoFooId(df.CK_LE_C, value)
	return q
}
func (q *VendorConstraintNameAutoFooCQ) SetConstraintNameAutoFooId_RangeOf(minNumber df.Numeric, maxNumber df.Numeric, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueConstraintNameAutoFooId(),"constraintNameAutoFooId",rangeOfOption)
}	


func (q *VendorConstraintNameAutoFooCQ) SetConstraintNameAutoFooId_IsNull() *VendorConstraintNameAutoFooCQ {
	q.regConstraintNameAutoFooId(df.CK_ISN_C, 0)
	return q
}
func (q *VendorConstraintNameAutoFooCQ) SetConstraintNameAutoFooId_IsNotNull() *VendorConstraintNameAutoFooCQ {
	q.regConstraintNameAutoFooId(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorConstraintNameAutoFooCQ) AddOrderBy_ConstraintNameAutoFooId_Asc() *VendorConstraintNameAutoFooCQ {
	q.BaseConditionQuery.RegOBA("constraintNameAutoFooId")
	return q
}
func (q *VendorConstraintNameAutoFooCQ) AddOrderBy_ConstraintNameAutoFooId_Desc() *VendorConstraintNameAutoFooCQ {
	q.BaseConditionQuery.RegOBD("constraintNameAutoFooId")
	return q
}
func (q *VendorConstraintNameAutoFooCQ) regConstraintNameAutoFooId(key *df.ConditionKey, value interface{}) {
	if q.ConstraintNameAutoFooId == nil {
		q.ConstraintNameAutoFooId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ConstraintNameAutoFooId, "constraintNameAutoFooId")
}

func (q *VendorConstraintNameAutoFooCQ) getCValueConstraintNameAutoFooName() *df.ConditionValue {
	if q.ConstraintNameAutoFooName == nil {
		q.ConstraintNameAutoFooName = new(df.ConditionValue)
	}
	return q.ConstraintNameAutoFooName
}


func (q *VendorConstraintNameAutoFooCQ) SetConstraintNameAutoFooName_Equal(value string) *VendorConstraintNameAutoFooCQ {
	q.regConstraintNameAutoFooName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *VendorConstraintNameAutoFooCQ) SetConstraintNameAutoFooName_NotEqual(value string) *VendorConstraintNameAutoFooCQ {
	q.regConstraintNameAutoFooName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorConstraintNameAutoFooCQ) SetConstraintNameAutoFooName_GreaterThan(value string) *VendorConstraintNameAutoFooCQ {
	q.regConstraintNameAutoFooName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorConstraintNameAutoFooCQ) SetConstraintNameAutoFooName_LessThan(value string) *VendorConstraintNameAutoFooCQ {
	q.regConstraintNameAutoFooName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorConstraintNameAutoFooCQ) SetConstraintNameAutoFooName_GreaterEqualThan(value string) *VendorConstraintNameAutoFooCQ {
	q.regConstraintNameAutoFooName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *VendorConstraintNameAutoFooCQ) SetConstraintNameAutoFooName_LessEqualThan(value string) *VendorConstraintNameAutoFooCQ {
	q.regConstraintNameAutoFooName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorConstraintNameAutoFooCQ) SetConstraintNameAutoFooName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueConstraintNameAutoFooName(), "constraintNameAutoFooName", option)
}

func (q *VendorConstraintNameAutoFooCQ) SetConstraintNameAutoFooName_PrefixSearch(value string) error {
	return q.SetConstraintNameAutoFooName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *VendorConstraintNameAutoFooCQ) SetConstraintNameAutoFooName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueConstraintNameAutoFooName(), "constraintNameAutoFooName", option)
}



func (q *VendorConstraintNameAutoFooCQ) AddOrderBy_ConstraintNameAutoFooName_Asc() *VendorConstraintNameAutoFooCQ {
	q.BaseConditionQuery.RegOBA("constraintNameAutoFooName")
	return q
}
func (q *VendorConstraintNameAutoFooCQ) AddOrderBy_ConstraintNameAutoFooName_Desc() *VendorConstraintNameAutoFooCQ {
	q.BaseConditionQuery.RegOBD("constraintNameAutoFooName")
	return q
}
func (q *VendorConstraintNameAutoFooCQ) regConstraintNameAutoFooName(key *df.ConditionKey, value interface{}) {
	if q.ConstraintNameAutoFooName == nil {
		q.ConstraintNameAutoFooName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ConstraintNameAutoFooName, "constraintNameAutoFooName")
}


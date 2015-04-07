package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorConstraintNameAutoRefCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	ConstraintNameAutoRefId *df.ConditionValue
	ConstraintNameAutoFooId *df.ConditionValue
	ConstraintNameAutoBarId *df.ConditionValue
	ConstraintNameAutoQuxId *df.ConditionValue
	ConstraintNameAutoCorgeId *df.ConditionValue
	ConstraintNameAutoUnique *df.ConditionValue
}

func (q *VendorConstraintNameAutoRefCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *VendorConstraintNameAutoRefCQ) getCValueConstraintNameAutoRefId() *df.ConditionValue {
	if q.ConstraintNameAutoRefId == nil {
		q.ConstraintNameAutoRefId = new(df.ConditionValue)
	}
	return q.ConstraintNameAutoRefId
}



func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoRefId_Equal(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoRefId(df.CK_EQ_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoRefId_NotEqual(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoRefId(df.CK_NE_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoRefId_GreaterThan(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoRefId(df.CK_GT_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoRefId_LessThan(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoRefId(df.CK_LT_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoRefId_GreaterEqual(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoRefId(df.CK_GE_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoRefId_LessEqual(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoRefId(df.CK_LE_C, value)
	return q
}
func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoRefId_RangeOf(minNumber df.Numeric, maxNumber df.Numeric, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueConstraintNameAutoRefId(),"constraintNameAutoRefId",rangeOfOption)
}	


func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoRefId_IsNull() *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoRefId(df.CK_ISN_C, 0)
	return q
}
func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoRefId_IsNotNull() *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoRefId(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorConstraintNameAutoRefCQ) AddOrderBy_ConstraintNameAutoRefId_Asc() *VendorConstraintNameAutoRefCQ {
	q.BaseConditionQuery.RegOBA("constraintNameAutoRefId")
	return q
}
func (q *VendorConstraintNameAutoRefCQ) AddOrderBy_ConstraintNameAutoRefId_Desc() *VendorConstraintNameAutoRefCQ {
	q.BaseConditionQuery.RegOBD("constraintNameAutoRefId")
	return q
}
func (q *VendorConstraintNameAutoRefCQ) regConstraintNameAutoRefId(key *df.ConditionKey, value interface{}) {
	if q.ConstraintNameAutoRefId == nil {
		q.ConstraintNameAutoRefId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ConstraintNameAutoRefId, "constraintNameAutoRefId")
}

func (q *VendorConstraintNameAutoRefCQ) getCValueConstraintNameAutoFooId() *df.ConditionValue {
	if q.ConstraintNameAutoFooId == nil {
		q.ConstraintNameAutoFooId = new(df.ConditionValue)
	}
	return q.ConstraintNameAutoFooId
}



func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoFooId_Equal(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoFooId(df.CK_EQ_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoFooId_NotEqual(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoFooId(df.CK_NE_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoFooId_GreaterThan(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoFooId(df.CK_GT_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoFooId_LessThan(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoFooId(df.CK_LT_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoFooId_GreaterEqual(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoFooId(df.CK_GE_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoFooId_LessEqual(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoFooId(df.CK_LE_C, value)
	return q
}
func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoFooId_RangeOf(minNumber df.Numeric, maxNumber df.Numeric, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueConstraintNameAutoFooId(),"constraintNameAutoFooId",rangeOfOption)
}	


func (q *VendorConstraintNameAutoRefCQ) AddOrderBy_ConstraintNameAutoFooId_Asc() *VendorConstraintNameAutoRefCQ {
	q.BaseConditionQuery.RegOBA("constraintNameAutoFooId")
	return q
}
func (q *VendorConstraintNameAutoRefCQ) AddOrderBy_ConstraintNameAutoFooId_Desc() *VendorConstraintNameAutoRefCQ {
	q.BaseConditionQuery.RegOBD("constraintNameAutoFooId")
	return q
}
func (q *VendorConstraintNameAutoRefCQ) regConstraintNameAutoFooId(key *df.ConditionKey, value interface{}) {
	if q.ConstraintNameAutoFooId == nil {
		q.ConstraintNameAutoFooId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ConstraintNameAutoFooId, "constraintNameAutoFooId")
}

func (q *VendorConstraintNameAutoRefCQ) getCValueConstraintNameAutoBarId() *df.ConditionValue {
	if q.ConstraintNameAutoBarId == nil {
		q.ConstraintNameAutoBarId = new(df.ConditionValue)
	}
	return q.ConstraintNameAutoBarId
}



func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoBarId_Equal(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoBarId(df.CK_EQ_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoBarId_NotEqual(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoBarId(df.CK_NE_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoBarId_GreaterThan(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoBarId(df.CK_GT_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoBarId_LessThan(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoBarId(df.CK_LT_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoBarId_GreaterEqual(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoBarId(df.CK_GE_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoBarId_LessEqual(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoBarId(df.CK_LE_C, value)
	return q
}
func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoBarId_RangeOf(minNumber df.Numeric, maxNumber df.Numeric, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueConstraintNameAutoBarId(),"constraintNameAutoBarId",rangeOfOption)
}	


func (q *VendorConstraintNameAutoRefCQ) AddOrderBy_ConstraintNameAutoBarId_Asc() *VendorConstraintNameAutoRefCQ {
	q.BaseConditionQuery.RegOBA("constraintNameAutoBarId")
	return q
}
func (q *VendorConstraintNameAutoRefCQ) AddOrderBy_ConstraintNameAutoBarId_Desc() *VendorConstraintNameAutoRefCQ {
	q.BaseConditionQuery.RegOBD("constraintNameAutoBarId")
	return q
}
func (q *VendorConstraintNameAutoRefCQ) regConstraintNameAutoBarId(key *df.ConditionKey, value interface{}) {
	if q.ConstraintNameAutoBarId == nil {
		q.ConstraintNameAutoBarId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ConstraintNameAutoBarId, "constraintNameAutoBarId")
}

func (q *VendorConstraintNameAutoRefCQ) getCValueConstraintNameAutoQuxId() *df.ConditionValue {
	if q.ConstraintNameAutoQuxId == nil {
		q.ConstraintNameAutoQuxId = new(df.ConditionValue)
	}
	return q.ConstraintNameAutoQuxId
}



func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoQuxId_Equal(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoQuxId(df.CK_EQ_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoQuxId_NotEqual(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoQuxId(df.CK_NE_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoQuxId_GreaterThan(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoQuxId(df.CK_GT_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoQuxId_LessThan(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoQuxId(df.CK_LT_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoQuxId_GreaterEqual(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoQuxId(df.CK_GE_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoQuxId_LessEqual(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoQuxId(df.CK_LE_C, value)
	return q
}
func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoQuxId_RangeOf(minNumber df.Numeric, maxNumber df.Numeric, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueConstraintNameAutoQuxId(),"constraintNameAutoQuxId",rangeOfOption)
}	


func (q *VendorConstraintNameAutoRefCQ) AddOrderBy_ConstraintNameAutoQuxId_Asc() *VendorConstraintNameAutoRefCQ {
	q.BaseConditionQuery.RegOBA("constraintNameAutoQuxId")
	return q
}
func (q *VendorConstraintNameAutoRefCQ) AddOrderBy_ConstraintNameAutoQuxId_Desc() *VendorConstraintNameAutoRefCQ {
	q.BaseConditionQuery.RegOBD("constraintNameAutoQuxId")
	return q
}
func (q *VendorConstraintNameAutoRefCQ) regConstraintNameAutoQuxId(key *df.ConditionKey, value interface{}) {
	if q.ConstraintNameAutoQuxId == nil {
		q.ConstraintNameAutoQuxId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ConstraintNameAutoQuxId, "constraintNameAutoQuxId")
}

func (q *VendorConstraintNameAutoRefCQ) getCValueConstraintNameAutoCorgeId() *df.ConditionValue {
	if q.ConstraintNameAutoCorgeId == nil {
		q.ConstraintNameAutoCorgeId = new(df.ConditionValue)
	}
	return q.ConstraintNameAutoCorgeId
}



func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoCorgeId_Equal(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoCorgeId(df.CK_EQ_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoCorgeId_NotEqual(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoCorgeId(df.CK_NE_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoCorgeId_GreaterThan(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoCorgeId(df.CK_GT_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoCorgeId_LessThan(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoCorgeId(df.CK_LT_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoCorgeId_GreaterEqual(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoCorgeId(df.CK_GE_C, value)
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoCorgeId_LessEqual(value df.Numeric) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoCorgeId(df.CK_LE_C, value)
	return q
}
func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoCorgeId_RangeOf(minNumber df.Numeric, maxNumber df.Numeric, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueConstraintNameAutoCorgeId(),"constraintNameAutoCorgeId",rangeOfOption)
}	


func (q *VendorConstraintNameAutoRefCQ) AddOrderBy_ConstraintNameAutoCorgeId_Asc() *VendorConstraintNameAutoRefCQ {
	q.BaseConditionQuery.RegOBA("constraintNameAutoCorgeId")
	return q
}
func (q *VendorConstraintNameAutoRefCQ) AddOrderBy_ConstraintNameAutoCorgeId_Desc() *VendorConstraintNameAutoRefCQ {
	q.BaseConditionQuery.RegOBD("constraintNameAutoCorgeId")
	return q
}
func (q *VendorConstraintNameAutoRefCQ) regConstraintNameAutoCorgeId(key *df.ConditionKey, value interface{}) {
	if q.ConstraintNameAutoCorgeId == nil {
		q.ConstraintNameAutoCorgeId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ConstraintNameAutoCorgeId, "constraintNameAutoCorgeId")
}

func (q *VendorConstraintNameAutoRefCQ) getCValueConstraintNameAutoUnique() *df.ConditionValue {
	if q.ConstraintNameAutoUnique == nil {
		q.ConstraintNameAutoUnique = new(df.ConditionValue)
	}
	return q.ConstraintNameAutoUnique
}


func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoUnique_Equal(value string) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoUnique(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoUnique_NotEqual(value string) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoUnique(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoUnique_GreaterThan(value string) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoUnique(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoUnique_LessThan(value string) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoUnique(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoUnique_GreaterEqualThan(value string) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoUnique(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoUnique_LessEqualThan(value string) *VendorConstraintNameAutoRefCQ {
	q.regConstraintNameAutoUnique(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoUnique_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueConstraintNameAutoUnique(), "constraintNameAutoUnique", option)
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoUnique_PrefixSearch(value string) error {
	return q.SetConstraintNameAutoUnique_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *VendorConstraintNameAutoRefCQ) SetConstraintNameAutoUnique_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueConstraintNameAutoUnique(), "constraintNameAutoUnique", option)
}



func (q *VendorConstraintNameAutoRefCQ) AddOrderBy_ConstraintNameAutoUnique_Asc() *VendorConstraintNameAutoRefCQ {
	q.BaseConditionQuery.RegOBA("constraintNameAutoUnique")
	return q
}
func (q *VendorConstraintNameAutoRefCQ) AddOrderBy_ConstraintNameAutoUnique_Desc() *VendorConstraintNameAutoRefCQ {
	q.BaseConditionQuery.RegOBD("constraintNameAutoUnique")
	return q
}
func (q *VendorConstraintNameAutoRefCQ) regConstraintNameAutoUnique(key *df.ConditionKey, value interface{}) {
	if q.ConstraintNameAutoUnique == nil {
		q.ConstraintNameAutoUnique = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ConstraintNameAutoUnique, "constraintNameAutoUnique")
}


package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorSymmetricCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	VendorSymmetricId *df.ConditionValue
	PlainText *df.ConditionValue
	EncryptedData *df.ConditionValue
}

func (q *VendorSymmetricCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *VendorSymmetricCQ) getCValueVendorSymmetricId() *df.ConditionValue {
	if q.VendorSymmetricId == nil {
		q.VendorSymmetricId = new(df.ConditionValue)
	}
	return q.VendorSymmetricId
}



func (q *VendorSymmetricCQ) SetVendorSymmetricId_Equal(value df.Numeric) *VendorSymmetricCQ {
	q.regVendorSymmetricId(df.CK_EQ_C, value)
	return q
}

func (q *VendorSymmetricCQ) SetVendorSymmetricId_NotEqual(value df.Numeric) *VendorSymmetricCQ {
	q.regVendorSymmetricId(df.CK_NE_C, value)
	return q
}

func (q *VendorSymmetricCQ) SetVendorSymmetricId_GreaterThan(value df.Numeric) *VendorSymmetricCQ {
	q.regVendorSymmetricId(df.CK_GT_C, value)
	return q
}

func (q *VendorSymmetricCQ) SetVendorSymmetricId_LessThan(value df.Numeric) *VendorSymmetricCQ {
	q.regVendorSymmetricId(df.CK_LT_C, value)
	return q
}

func (q *VendorSymmetricCQ) SetVendorSymmetricId_GreaterEqual(value df.Numeric) *VendorSymmetricCQ {
	q.regVendorSymmetricId(df.CK_GE_C, value)
	return q
}

func (q *VendorSymmetricCQ) SetVendorSymmetricId_LessEqual(value df.Numeric) *VendorSymmetricCQ {
	q.regVendorSymmetricId(df.CK_LE_C, value)
	return q
}
func (q *VendorSymmetricCQ) SetVendorSymmetricId_RangeOf(minNumber df.Numeric, maxNumber df.Numeric, rangeOfOption *df.RangeOfOption) error {
	return q.BaseConditionQuery.RegROO(minNumber,maxNumber,q.getCValueVendorSymmetricId(),"vendorSymmetricId",rangeOfOption)
}	


func (q *VendorSymmetricCQ) SetVendorSymmetricId_IsNull() *VendorSymmetricCQ {
	q.regVendorSymmetricId(df.CK_ISN_C, 0)
	return q
}
func (q *VendorSymmetricCQ) SetVendorSymmetricId_IsNotNull() *VendorSymmetricCQ {
	q.regVendorSymmetricId(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorSymmetricCQ) AddOrderBy_VendorSymmetricId_Asc() *VendorSymmetricCQ {
	q.BaseConditionQuery.RegOBA("vendorSymmetricId")
	return q
}
func (q *VendorSymmetricCQ) AddOrderBy_VendorSymmetricId_Desc() *VendorSymmetricCQ {
	q.BaseConditionQuery.RegOBD("vendorSymmetricId")
	return q
}
func (q *VendorSymmetricCQ) regVendorSymmetricId(key *df.ConditionKey, value interface{}) {
	if q.VendorSymmetricId == nil {
		q.VendorSymmetricId = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.VendorSymmetricId, "vendorSymmetricId")
}

func (q *VendorSymmetricCQ) getCValuePlainText() *df.ConditionValue {
	if q.PlainText == nil {
		q.PlainText = new(df.ConditionValue)
	}
	return q.PlainText
}


func (q *VendorSymmetricCQ) SetPlainText_Equal(value string) *VendorSymmetricCQ {
	q.regPlainText(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}

func (q *VendorSymmetricCQ) SetPlainText_NotEqual(value string) *VendorSymmetricCQ {
	q.regPlainText(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorSymmetricCQ) SetPlainText_GreaterThan(value string) *VendorSymmetricCQ {
	q.regPlainText(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorSymmetricCQ) SetPlainText_LessThan(value string) *VendorSymmetricCQ {
	q.regPlainText(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorSymmetricCQ) SetPlainText_GreaterEqualThan(value string) *VendorSymmetricCQ {
	q.regPlainText(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *VendorSymmetricCQ) SetPlainText_LessEqualThan(value string) *VendorSymmetricCQ {
	q.regPlainText(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *VendorSymmetricCQ) SetPlainText_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValuePlainText(), "plainText", option)
}

func (q *VendorSymmetricCQ) SetPlainText_PrefixSearch(value string) error {
	return q.SetPlainText_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *VendorSymmetricCQ) SetPlainText_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValuePlainText(), "plainText", option)
}



func (q *VendorSymmetricCQ) SetPlainText_IsNull() *VendorSymmetricCQ {
	q.regPlainText(df.CK_ISN_C, 0)
	return q
}
func (q *VendorSymmetricCQ) SetPlainText_IsNullOrEmpty() *VendorSymmetricCQ {
	q.regPlainText(df.CK_ISNOE_C, 0)
	return q
}
func (q *VendorSymmetricCQ) SetPlainText_IsNotNull() *VendorSymmetricCQ {
	q.regPlainText(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorSymmetricCQ) AddOrderBy_PlainText_Asc() *VendorSymmetricCQ {
	q.BaseConditionQuery.RegOBA("plainText")
	return q
}
func (q *VendorSymmetricCQ) AddOrderBy_PlainText_Desc() *VendorSymmetricCQ {
	q.BaseConditionQuery.RegOBD("plainText")
	return q
}
func (q *VendorSymmetricCQ) regPlainText(key *df.ConditionKey, value interface{}) {
	if q.PlainText == nil {
		q.PlainText = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.PlainText, "plainText")
}

func (q *VendorSymmetricCQ) getCValueEncryptedData() *df.ConditionValue {
	if q.EncryptedData == nil {
		q.EncryptedData = new(df.ConditionValue)
	}
	return q.EncryptedData
}





func (q *VendorSymmetricCQ) SetEncryptedData_IsNull() *VendorSymmetricCQ {
	q.regEncryptedData(df.CK_ISN_C, 0)
	return q
}
func (q *VendorSymmetricCQ) SetEncryptedData_IsNotNull() *VendorSymmetricCQ {
	q.regEncryptedData(df.CK_ISNN_C, 0)
	return q
}
func (q *VendorSymmetricCQ) AddOrderBy_EncryptedData_Asc() *VendorSymmetricCQ {
	q.BaseConditionQuery.RegOBA("encryptedData")
	return q
}
func (q *VendorSymmetricCQ) AddOrderBy_EncryptedData_Desc() *VendorSymmetricCQ {
	q.BaseConditionQuery.RegOBD("encryptedData")
	return q
}
func (q *VendorSymmetricCQ) regEncryptedData(key *df.ConditionKey, value interface{}) {
	if q.EncryptedData == nil {
		q.EncryptedData = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.EncryptedData, "encryptedData")
}


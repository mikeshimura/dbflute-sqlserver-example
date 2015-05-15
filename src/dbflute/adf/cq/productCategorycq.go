package cq

import (
	"github.com/mikeshimura/dbflute/df"
)

type ProductCategoryCQ struct {
	BaseConditionQuery *df.BaseConditionQuery
	ProductCategoryCode *df.ConditionValue
	ProductCategoryName *df.ConditionValue
	ParentCategoryCode *df.ConditionValue
    conditionQueryProductCategorySelf *ProductCategoryCQ
}

func (q *ProductCategoryCQ) GetBaseConditionQuery() *df.BaseConditionQuery{
	return q.BaseConditionQuery
}


func (q *ProductCategoryCQ) getCValueProductCategoryCode() *df.ConditionValue {
	if q.ProductCategoryCode == nil {
		q.ProductCategoryCode = new(df.ConditionValue)
	}
	return q.ProductCategoryCode
}


func (q *ProductCategoryCQ) SetProductCategoryCode_Equal(value string) *ProductCategoryCQ {
	q.regProductCategoryCode(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *ProductCategoryCQ) SetProductCategoryCode_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueProductCategoryCode(), "productCategoryCode")
}
func (q *ProductCategoryCQ) SetProductCategoryCode_NotEqual(value string) *ProductCategoryCQ {
	q.regProductCategoryCode(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCategoryCQ) SetProductCategoryCode_GreaterThan(value string) *ProductCategoryCQ {
	q.regProductCategoryCode(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCategoryCQ) SetProductCategoryCode_LessThan(value string) *ProductCategoryCQ {
	q.regProductCategoryCode(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCategoryCQ) SetProductCategoryCode_GreaterEqualThan(value string) *ProductCategoryCQ {
	q.regProductCategoryCode(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ProductCategoryCQ) SetProductCategoryCode_LessEqualThan(value string) *ProductCategoryCQ {
	q.regProductCategoryCode(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCategoryCQ) SetProductCategoryCode_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueProductCategoryCode(), "productCategoryCode", option)
}

func (q *ProductCategoryCQ) SetProductCategoryCode_PrefixSearch(value string) error {
	return q.SetProductCategoryCode_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ProductCategoryCQ) SetProductCategoryCode_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueProductCategoryCode(), "productCategoryCode", option)
}



func (q *ProductCategoryCQ) SetProductCategoryCode_IsNull() *ProductCategoryCQ {
	q.regProductCategoryCode(df.CK_ISN_C, 0)
	return q
}
func (q *ProductCategoryCQ) SetProductCategoryCode_IsNotNull() *ProductCategoryCQ {
	q.regProductCategoryCode(df.CK_ISNN_C, 0)
	return q
}
func (q *ProductCategoryCQ) AddOrderBy_ProductCategoryCode_Asc() *ProductCategoryCQ {
	q.BaseConditionQuery.RegOBA("productCategoryCode")
	return q
}
func (q *ProductCategoryCQ) AddOrderBy_ProductCategoryCode_Desc() *ProductCategoryCQ {
	q.BaseConditionQuery.RegOBD("productCategoryCode")
	return q
}
func (q *ProductCategoryCQ) regProductCategoryCode(key *df.ConditionKey, value interface{}) {
	if q.ProductCategoryCode == nil {
		q.ProductCategoryCode = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ProductCategoryCode, "productCategoryCode")
}

func (q *ProductCategoryCQ) getCValueProductCategoryName() *df.ConditionValue {
	if q.ProductCategoryName == nil {
		q.ProductCategoryName = new(df.ConditionValue)
	}
	return q.ProductCategoryName
}


func (q *ProductCategoryCQ) SetProductCategoryName_Equal(value string) *ProductCategoryCQ {
	q.regProductCategoryName(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *ProductCategoryCQ) SetProductCategoryName_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueProductCategoryName(), "productCategoryName")
}
func (q *ProductCategoryCQ) SetProductCategoryName_NotEqual(value string) *ProductCategoryCQ {
	q.regProductCategoryName(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCategoryCQ) SetProductCategoryName_GreaterThan(value string) *ProductCategoryCQ {
	q.regProductCategoryName(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCategoryCQ) SetProductCategoryName_LessThan(value string) *ProductCategoryCQ {
	q.regProductCategoryName(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCategoryCQ) SetProductCategoryName_GreaterEqualThan(value string) *ProductCategoryCQ {
	q.regProductCategoryName(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ProductCategoryCQ) SetProductCategoryName_LessEqualThan(value string) *ProductCategoryCQ {
	q.regProductCategoryName(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCategoryCQ) SetProductCategoryName_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueProductCategoryName(), "productCategoryName", option)
}

func (q *ProductCategoryCQ) SetProductCategoryName_PrefixSearch(value string) error {
	return q.SetProductCategoryName_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ProductCategoryCQ) SetProductCategoryName_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueProductCategoryName(), "productCategoryName", option)
}



func (q *ProductCategoryCQ) AddOrderBy_ProductCategoryName_Asc() *ProductCategoryCQ {
	q.BaseConditionQuery.RegOBA("productCategoryName")
	return q
}
func (q *ProductCategoryCQ) AddOrderBy_ProductCategoryName_Desc() *ProductCategoryCQ {
	q.BaseConditionQuery.RegOBD("productCategoryName")
	return q
}
func (q *ProductCategoryCQ) regProductCategoryName(key *df.ConditionKey, value interface{}) {
	if q.ProductCategoryName == nil {
		q.ProductCategoryName = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ProductCategoryName, "productCategoryName")
}

func (q *ProductCategoryCQ) getCValueParentCategoryCode() *df.ConditionValue {
	if q.ParentCategoryCode == nil {
		q.ParentCategoryCode = new(df.ConditionValue)
	}
	return q.ParentCategoryCode
}


func (q *ProductCategoryCQ) SetParentCategoryCode_Equal(value string) *ProductCategoryCQ {
	q.regParentCategoryCode(df.CK_EQ_C, q.BaseConditionQuery.FRES(value))
	return q
}
func (q *ProductCategoryCQ) SetParentCategoryCode_InScope(list *df.List){
	q.BaseConditionQuery.RegINS(df.CK_INS_C, list,
		 q.getCValueParentCategoryCode(), "parentCategoryCode")
}
func (q *ProductCategoryCQ) SetParentCategoryCode_NotEqual(value string) *ProductCategoryCQ {
	q.regParentCategoryCode(df.CK_NE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCategoryCQ) SetParentCategoryCode_GreaterThan(value string) *ProductCategoryCQ {
	q.regParentCategoryCode(df.CK_GT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCategoryCQ) SetParentCategoryCode_LessThan(value string) *ProductCategoryCQ {
	q.regParentCategoryCode(df.CK_LT_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCategoryCQ) SetParentCategoryCode_GreaterEqualThan(value string) *ProductCategoryCQ {
	q.regParentCategoryCode(df.CK_GE_C, q.BaseConditionQuery.FRES(value))
	return q
}	
func (q *ProductCategoryCQ) SetParentCategoryCode_LessEqualThan(value string) *ProductCategoryCQ {
	q.regParentCategoryCode(df.CK_LE_C, q.BaseConditionQuery.FRES(value))
	return q
}	

func (q *ProductCategoryCQ) SetParentCategoryCode_LikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_LS_C, value, q.getCValueParentCategoryCode(), "parentCategoryCode", option)
}

func (q *ProductCategoryCQ) SetParentCategoryCode_PrefixSearch(value string) error {
	return q.SetParentCategoryCode_LikeSearch(value, q.BaseConditionQuery.CLSOP())
}

func (q *ProductCategoryCQ) SetParentCategoryCode_NotLikeSearch(value string, option *df.LikeSearchOption) error {
	return q.BaseConditionQuery.RegLSQ(df.CK_NLS_C, value, q.getCValueParentCategoryCode(), "parentCategoryCode", option)
}



func (q *ProductCategoryCQ) SetParentCategoryCode_IsNull() *ProductCategoryCQ {
	q.regParentCategoryCode(df.CK_ISN_C, 0)
	return q
}
func (q *ProductCategoryCQ) SetParentCategoryCode_IsNotNull() *ProductCategoryCQ {
	q.regParentCategoryCode(df.CK_ISNN_C, 0)
	return q
}
func (q *ProductCategoryCQ) AddOrderBy_ParentCategoryCode_Asc() *ProductCategoryCQ {
	q.BaseConditionQuery.RegOBA("parentCategoryCode")
	return q
}
func (q *ProductCategoryCQ) AddOrderBy_ParentCategoryCode_Desc() *ProductCategoryCQ {
	q.BaseConditionQuery.RegOBD("parentCategoryCode")
	return q
}
func (q *ProductCategoryCQ) regParentCategoryCode(key *df.ConditionKey, value interface{}) {
	if q.ParentCategoryCode == nil {
		q.ParentCategoryCode = new(df.ConditionValue)
	}
	q.BaseConditionQuery.RegQ(key, value, q.ParentCategoryCode, "parentCategoryCode")
}


func (q *ProductCategoryCQ) QueryProductCategorySelf() *ProductCategoryCQ {
	if q.conditionQueryProductCategorySelf == nil {
		q.conditionQueryProductCategorySelf = q.xcreateQueryProductCategorySelf()
		q.xsetupOuterJoinProductCategorySelf()
	}
	return q.conditionQueryProductCategorySelf
}

func (q *ProductCategoryCQ) xcreateQueryProductCategorySelf() *ProductCategoryCQ {
	nrp := q.BaseConditionQuery.ResolveNextRelationPath("ProductCategory", "ProductCategorySelf")
	jan := q.BaseConditionQuery.ResolveJoinAliasName(nrp)
	var basecq df.ConditionQuery = q
	cq := CreateProductCategoryCQ(&basecq, q.BaseConditionQuery.SqlClause, jan, q.BaseConditionQuery.NestLevel+1)
	cq.BaseConditionQuery.BaseCB = q.BaseConditionQuery.BaseCB
	cq.BaseConditionQuery.ForeignPropertyName = "ProductCategorySelf"
	cq.BaseConditionQuery.RelationPath = nrp
	return cq
}
func (q *ProductCategoryCQ) xsetupOuterJoinProductCategorySelf() {
	    cq := q.QueryProductCategorySelf()
        joinOnMap := make(map[string]string)
        joinOnMap["parentCategoryCode"]="productCategoryCode"
        q.BaseConditionQuery.RegisterOuterJoin(
        	cq.BaseConditionQuery.ConditionQuery, joinOnMap, "ProductCategorySelf");
}	
	
func CreateProductCategoryCQ(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *ProductCategoryCQ {
	cq := new(ProductCategoryCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = "ProductCategory"
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
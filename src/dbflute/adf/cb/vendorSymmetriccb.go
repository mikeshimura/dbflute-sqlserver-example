package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type VendorSymmetricCB struct {
	BaseConditionBean *df.BaseConditionBean
	Query             *cq.VendorSymmetricCQ
}

func CreateVendorSymmetricCB() *VendorSymmetricCB {
	cb := new(VendorSymmetricCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "VendorSymmetric"
	cb.BaseConditionBean.Name = "VendorSymmetricCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	//dm:=DBMetaProvider_I.TableDbNameInstanceMap["VendorSymmetric"]
	var dmx df.DBMeta = meta.VendorSymmetricDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	cb.Query = cb.createConditionQuery(nil, cb.BaseConditionBean.SqlClause, (*cb.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
	return cb
}
func (l *VendorSymmetricCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *VendorSymmetricCB) createConditionQuery(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *cq.VendorSymmetricCQ {
	cq := new(cq.VendorSymmetricCQ)
	cq.BaseConditionQuery = new(df.BaseConditionQuery)
	cq.BaseConditionQuery.TableDbName = l.BaseConditionBean.TableDbName
	cq.BaseConditionQuery.ReferrerQuery = referrerQuery
	cq.BaseConditionQuery.SqlClause = sqlClause
	cq.BaseConditionQuery.AliasName = aliasName
	cq.BaseConditionQuery.NestLevel = nestlevel
	cq.BaseConditionQuery.DBMetaProvider = l.BaseConditionBean.DBMetaProvider
	cq.BaseConditionQuery.CQ_PROPERTY = "Query"
	cq.BaseConditionQuery.ConditionQuery=cq
	return cq
}

func (l *VendorSymmetricCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}

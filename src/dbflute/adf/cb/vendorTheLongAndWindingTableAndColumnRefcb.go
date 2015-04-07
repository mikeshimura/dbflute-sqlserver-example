package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type VendorTheLongAndWindingTableAndColumnRefCB struct {
	BaseConditionBean *df.BaseConditionBean
	Query             *cq.VendorTheLongAndWindingTableAndColumnRefCQ
}

func CreateVendorTheLongAndWindingTableAndColumnRefCB() *VendorTheLongAndWindingTableAndColumnRefCB {
	cb := new(VendorTheLongAndWindingTableAndColumnRefCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "VendorTheLongAndWindingTableAndColumnRef"
	cb.BaseConditionBean.Name = "VendorTheLongAndWindingTableAndColumnRefCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	//dm:=DBMetaProvider_I.TableDbNameInstanceMap["VendorTheLongAndWindingTableAndColumnRef"]
	var dmx df.DBMeta = meta.VendorTheLongAndWindingTableAndColumnRefDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	cb.Query = cb.createConditionQuery(nil, cb.BaseConditionBean.SqlClause, (*cb.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
	return cb
}
func (l *VendorTheLongAndWindingTableAndColumnRefCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *VendorTheLongAndWindingTableAndColumnRefCB) createConditionQuery(referrerQuery *df.ConditionQuery, sqlClause *df.SqlClause, aliasName string, nestlevel int8) *cq.VendorTheLongAndWindingTableAndColumnRefCQ {
	cq := new(cq.VendorTheLongAndWindingTableAndColumnRefCQ)
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

func (l *VendorTheLongAndWindingTableAndColumnRefCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}

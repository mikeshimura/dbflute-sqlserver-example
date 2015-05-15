package cb

import (
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/cq"
	"dbflute/adf/meta"
)

type WhiteDelimiterCB struct {
	BaseConditionBean *df.BaseConditionBean
	query             *cq.WhiteDelimiterCQ
}

func CreateWhiteDelimiterCB() *WhiteDelimiterCB {
	cb := new(WhiteDelimiterCB)
	cb.BaseConditionBean = new(df.BaseConditionBean)
	cb.BaseConditionBean.DBMetaProvider = df.DBMetaProvider_I
	cb.BaseConditionBean.TableDbName = "WhiteDelimiter"
	cb.BaseConditionBean.Name = "WhiteDelimiterCB"
	cb.BaseConditionBean.SqlClause = df.CreateSqlClause(cb, df.DBCurrent_I)
	var dmx df.DBMeta = meta.WhiteDelimiterDbm
	(*cb.BaseConditionBean.SqlClause).SetDBMeta(&dmx)
	(*cb.BaseConditionBean.SqlClause).SetUseSelectIndex(true)
	return cb
}

func (l *WhiteDelimiterCB) Query() *cq.WhiteDelimiterCQ {
	if l.query == nil {
		l.query = cq.CreateWhiteDelimiterCQ(nil, l.BaseConditionBean.SqlClause, (*l.BaseConditionBean.SqlClause).GetBasePorintAliasName(), 0)
		var cb df.ConditionBean = l
		l.query.BaseConditionQuery.BaseCB = &cb	
	}
	return l.query
}
func (l *WhiteDelimiterCB) GetBaseConditionBean() *df.BaseConditionBean {
	return l.BaseConditionBean
}

func (l *WhiteDelimiterCB) AllowEmptyStringQuery() {
	l.BaseConditionBean.AllowEmptyStringQuery()
}


func (l *WhiteDelimiterCB) FetchFirst(fetchSize int){
	(*l.GetBaseConditionBean().SqlClause).FetchFirst(fetchSize)
}

func (l *WhiteDelimiterCB) OrScopeQuery(fquery func(*WhiteDelimiterCB)){
	(*l.BaseConditionBean.SqlClause).MakeOrScopeQueryEffective()
	fquery(l)
	(*l.BaseConditionBean.SqlClause).CloseOrScopeQuery()
}

type WhiteDelimiterNss struct {
	Query *cq.WhiteDelimiterCQ
}
func (p *WhiteDelimiterNss) hasConditionQuery() bool {
	return p.Query != nil
}
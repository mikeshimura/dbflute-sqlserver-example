package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type WhiteDelimiterBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *WhiteDelimiterBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.WhiteDelimiterDbm
	return &meta
}
func (l *WhiteDelimiterBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *WhiteDelimiterBhv) Update(whiteDelimiter *entity.WhiteDelimiter, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = whiteDelimiter
	return l.BaseBehavior.DoUpdate(&entity, nil, tx, ctx)
}
func (l *WhiteDelimiterBhv) Insert(whiteDelimiter *entity.WhiteDelimiter, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = whiteDelimiter
	return l.BaseBehavior.DoInsert(&entity, nil, tx, ctx)
}
func (l *WhiteDelimiterBhv) Delete(whiteDelimiter *entity.WhiteDelimiter, tx *sql.Tx, ctx *df.Context) (int64, error) {
	var entity df.Entity = whiteDelimiter
	return l.BaseBehavior.DoDelete(&entity, nil, tx, ctx)
}
func (l *WhiteDelimiterBhv) SelectList(cb *cb.WhiteDelimiterCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "WhiteDelimiter", tx)
}
func (l *WhiteDelimiterBhv) SelectCount(cb *cb.WhiteDelimiterCB, tx *sql.Tx) (int64, error) {

	return l.BaseBehavior.DoSelectCount(cb, tx)
}
func (l *WhiteDelimiterBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *WhiteDelimiterBhv) ReadNextVal(tx *sql.Tx) int64{
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var WhiteDelimiterBhv_I *WhiteDelimiterBhv

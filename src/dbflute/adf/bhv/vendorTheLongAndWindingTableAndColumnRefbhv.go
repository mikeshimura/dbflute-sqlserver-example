package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type VendorTheLongAndWindingTableAndColumnRefBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *VendorTheLongAndWindingTableAndColumnRefBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.VendorTheLongAndWindingTableAndColumnRefDbm
	return &meta
}
func (l *VendorTheLongAndWindingTableAndColumnRefBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *VendorTheLongAndWindingTableAndColumnRefBhv) Update(vendorTheLongAndWindingTableAndColumnRef *entity.VendorTheLongAndWindingTableAndColumnRef, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorTheLongAndWindingTableAndColumnRef
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *VendorTheLongAndWindingTableAndColumnRefBhv) Insert(vendorTheLongAndWindingTableAndColumnRef *entity.VendorTheLongAndWindingTableAndColumnRef, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorTheLongAndWindingTableAndColumnRef
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *VendorTheLongAndWindingTableAndColumnRefBhv) Delete(vendorTheLongAndWindingTableAndColumnRef *entity.VendorTheLongAndWindingTableAndColumnRef, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorTheLongAndWindingTableAndColumnRef
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *VendorTheLongAndWindingTableAndColumnRefBhv) SelectList(cb *cb.VendorTheLongAndWindingTableAndColumnRefCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "VendorTheLongAndWindingTableAndColumnRef", tx)
}
func (l *VendorTheLongAndWindingTableAndColumnRefBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *VendorTheLongAndWindingTableAndColumnRefBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var VendorTheLongAndWindingTableAndColumnRefBhv_I *VendorTheLongAndWindingTableAndColumnRefBhv

package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type VendorTheLongAndWindingTableAndColumnBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *VendorTheLongAndWindingTableAndColumnBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.VendorTheLongAndWindingTableAndColumnDbm
	return &meta
}
func (l *VendorTheLongAndWindingTableAndColumnBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *VendorTheLongAndWindingTableAndColumnBhv) Update(vendorTheLongAndWindingTableAndColumn *entity.VendorTheLongAndWindingTableAndColumn, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorTheLongAndWindingTableAndColumn
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *VendorTheLongAndWindingTableAndColumnBhv) Insert(vendorTheLongAndWindingTableAndColumn *entity.VendorTheLongAndWindingTableAndColumn, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorTheLongAndWindingTableAndColumn
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *VendorTheLongAndWindingTableAndColumnBhv) Delete(vendorTheLongAndWindingTableAndColumn *entity.VendorTheLongAndWindingTableAndColumn, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorTheLongAndWindingTableAndColumn
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *VendorTheLongAndWindingTableAndColumnBhv) SelectList(cb *cb.VendorTheLongAndWindingTableAndColumnCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "VendorTheLongAndWindingTableAndColumn", tx)
}
func (l *VendorTheLongAndWindingTableAndColumnBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *VendorTheLongAndWindingTableAndColumnBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var VendorTheLongAndWindingTableAndColumnBhv_I *VendorTheLongAndWindingTableAndColumnBhv

package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type VendorConstraintNameAutoBarBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *VendorConstraintNameAutoBarBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.VendorConstraintNameAutoBarDbm
	return &meta
}
func (l *VendorConstraintNameAutoBarBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *VendorConstraintNameAutoBarBhv) Update(vendorConstraintNameAutoBar *entity.VendorConstraintNameAutoBar, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorConstraintNameAutoBar
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *VendorConstraintNameAutoBarBhv) Insert(vendorConstraintNameAutoBar *entity.VendorConstraintNameAutoBar, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorConstraintNameAutoBar
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *VendorConstraintNameAutoBarBhv) Delete(vendorConstraintNameAutoBar *entity.VendorConstraintNameAutoBar, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorConstraintNameAutoBar
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *VendorConstraintNameAutoBarBhv) SelectList(cb *cb.VendorConstraintNameAutoBarCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "VendorConstraintNameAutoBar", tx)
}
func (l *VendorConstraintNameAutoBarBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *VendorConstraintNameAutoBarBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var VendorConstraintNameAutoBarBhv_I *VendorConstraintNameAutoBarBhv

package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type VendorConstraintNameAutoQuxBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *VendorConstraintNameAutoQuxBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.VendorConstraintNameAutoQuxDbm
	return &meta
}
func (l *VendorConstraintNameAutoQuxBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *VendorConstraintNameAutoQuxBhv) Update(vendorConstraintNameAutoQux *entity.VendorConstraintNameAutoQux, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorConstraintNameAutoQux
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *VendorConstraintNameAutoQuxBhv) Insert(vendorConstraintNameAutoQux *entity.VendorConstraintNameAutoQux, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorConstraintNameAutoQux
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *VendorConstraintNameAutoQuxBhv) Delete(vendorConstraintNameAutoQux *entity.VendorConstraintNameAutoQux, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorConstraintNameAutoQux
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *VendorConstraintNameAutoQuxBhv) SelectList(cb *cb.VendorConstraintNameAutoQuxCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "VendorConstraintNameAutoQux", tx)
}
func (l *VendorConstraintNameAutoQuxBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *VendorConstraintNameAutoQuxBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var VendorConstraintNameAutoQuxBhv_I *VendorConstraintNameAutoQuxBhv

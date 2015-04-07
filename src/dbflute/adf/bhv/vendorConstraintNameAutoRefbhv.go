package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type VendorConstraintNameAutoRefBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *VendorConstraintNameAutoRefBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.VendorConstraintNameAutoRefDbm
	return &meta
}
func (l *VendorConstraintNameAutoRefBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *VendorConstraintNameAutoRefBhv) Update(vendorConstraintNameAutoRef *entity.VendorConstraintNameAutoRef, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorConstraintNameAutoRef
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *VendorConstraintNameAutoRefBhv) Insert(vendorConstraintNameAutoRef *entity.VendorConstraintNameAutoRef, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorConstraintNameAutoRef
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *VendorConstraintNameAutoRefBhv) Delete(vendorConstraintNameAutoRef *entity.VendorConstraintNameAutoRef, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorConstraintNameAutoRef
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *VendorConstraintNameAutoRefBhv) SelectList(cb *cb.VendorConstraintNameAutoRefCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "VendorConstraintNameAutoRef", tx)
}
func (l *VendorConstraintNameAutoRefBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *VendorConstraintNameAutoRefBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var VendorConstraintNameAutoRefBhv_I *VendorConstraintNameAutoRefBhv

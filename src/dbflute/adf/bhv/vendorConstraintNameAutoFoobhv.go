package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type VendorConstraintNameAutoFooBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *VendorConstraintNameAutoFooBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.VendorConstraintNameAutoFooDbm
	return &meta
}
func (l *VendorConstraintNameAutoFooBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *VendorConstraintNameAutoFooBhv) Update(vendorConstraintNameAutoFoo *entity.VendorConstraintNameAutoFoo, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorConstraintNameAutoFoo
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *VendorConstraintNameAutoFooBhv) Insert(vendorConstraintNameAutoFoo *entity.VendorConstraintNameAutoFoo, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorConstraintNameAutoFoo
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *VendorConstraintNameAutoFooBhv) Delete(vendorConstraintNameAutoFoo *entity.VendorConstraintNameAutoFoo, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorConstraintNameAutoFoo
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *VendorConstraintNameAutoFooBhv) SelectList(cb *cb.VendorConstraintNameAutoFooCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "VendorConstraintNameAutoFoo", tx)
}
func (l *VendorConstraintNameAutoFooBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *VendorConstraintNameAutoFooBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var VendorConstraintNameAutoFooBhv_I *VendorConstraintNameAutoFooBhv

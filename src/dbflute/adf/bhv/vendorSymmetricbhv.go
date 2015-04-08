package bhv

import (
	"database/sql"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/cb"
	"dbflute/adf/meta"
)

type VendorSymmetricBhv struct {
	BaseBehavior *df.BaseBehavior
}
func (l *VendorSymmetricBhv) GetDBMeta() *df.DBMeta{
	var meta df.DBMeta=meta.VendorSymmetricDbm
	return &meta
}
func (l *VendorSymmetricBhv) GetBaseBehavior() *df.BaseBehavior {
	return l.BaseBehavior
}
func (l *VendorSymmetricBhv) Update(vendorSymmetric *entity.VendorSymmetric, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorSymmetric
	return l.BaseBehavior.DoUpdate(&entity, nil, tx)
}
func (l *VendorSymmetricBhv) Insert(vendorSymmetric *entity.VendorSymmetric, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorSymmetric
	return l.BaseBehavior.DoInsert(&entity, nil, tx)
}
func (l *VendorSymmetricBhv) Delete(vendorSymmetric *entity.VendorSymmetric, tx *sql.Tx) (int64, error) {
	var entity df.Entity = vendorSymmetric
	return l.BaseBehavior.DoDelete(&entity, nil, tx)
}
func (l *VendorSymmetricBhv) SelectList(cb *cb.VendorSymmetricCB, tx *sql.Tx) (*df.ListResultBean, error) {

	return l.BaseBehavior.DoSelectList(cb, "VendorSymmetric", tx)
}
func (l *VendorSymmetricBhv) OutsideSql() *df.OutsideSqlBasicExecutor {
	return l.BaseBehavior.DoOutsideSql()
}

func (l *VendorSymmetricBhv) ReadNextVal(tx *sql.Tx) (int64,error){
	return l.BaseBehavior.DoSelectNextVal(tx)
}

var VendorSymmetricBhv_I *VendorSymmetricBhv

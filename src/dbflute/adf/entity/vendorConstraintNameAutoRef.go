package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorConstraintNameAutoRef struct {
	constraintNameAutoRefId df.Numeric
	constraintNameAutoFooId df.Numeric
	constraintNameAutoBarId df.Numeric
	constraintNameAutoQuxId df.Numeric
	constraintNameAutoCorgeId df.Numeric
	constraintNameAutoUnique string
	df.BaseEntity
}

func CreateVendorConstraintNameAutoRef() *VendorConstraintNameAutoRef{
	vendorConstraintNameAutoRef :=new(VendorConstraintNameAutoRef)
	vendorConstraintNameAutoRef.SetUp()
	return vendorConstraintNameAutoRef 
}

func (l *VendorConstraintNameAutoRef) GetConstraintNameAutoRefId () df.Numeric {
	return l.constraintNameAutoRefId
}
func (l *VendorConstraintNameAutoRef) GetConstraintNameAutoFooId () df.Numeric {
	return l.constraintNameAutoFooId
}
func (l *VendorConstraintNameAutoRef) GetConstraintNameAutoBarId () df.Numeric {
	return l.constraintNameAutoBarId
}
func (l *VendorConstraintNameAutoRef) GetConstraintNameAutoQuxId () df.Numeric {
	return l.constraintNameAutoQuxId
}
func (l *VendorConstraintNameAutoRef) GetConstraintNameAutoCorgeId () df.Numeric {
	return l.constraintNameAutoCorgeId
}
func (l *VendorConstraintNameAutoRef) GetConstraintNameAutoUnique () string {
	return l.constraintNameAutoUnique
}

func (t *VendorConstraintNameAutoRef) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 6)
	i[0] = &(t.constraintNameAutoRefId)
	i[1] = &(t.constraintNameAutoFooId)
	i[2] = &(t.constraintNameAutoBarId)
	i[3] = &(t.constraintNameAutoQuxId)
	i[4] = &(t.constraintNameAutoCorgeId)
	i[5] = &(t.constraintNameAutoUnique)
	return i
}


func (t *VendorConstraintNameAutoRef) AsTableDbName() string {
	return "VendorConstraintNameAutoRef"
}

func (t *VendorConstraintNameAutoRef) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("constraintNameAutoRefId") == false {
            return false 
        }
        return true;
}
func (t *VendorConstraintNameAutoRef) SetConstraintNameAutoRefId(constraintNameAutoRefId df.Numeric) {
	t.AddPropertyName("constraintNameAutoRefId")
	t.constraintNameAutoRefId = constraintNameAutoRefId
}
func (t *VendorConstraintNameAutoRef) SetConstraintNameAutoFooId(constraintNameAutoFooId df.Numeric) {
	t.AddPropertyName("constraintNameAutoFooId")
	t.constraintNameAutoFooId = constraintNameAutoFooId
}
func (t *VendorConstraintNameAutoRef) SetConstraintNameAutoBarId(constraintNameAutoBarId df.Numeric) {
	t.AddPropertyName("constraintNameAutoBarId")
	t.constraintNameAutoBarId = constraintNameAutoBarId
}
func (t *VendorConstraintNameAutoRef) SetConstraintNameAutoQuxId(constraintNameAutoQuxId df.Numeric) {
	t.AddPropertyName("constraintNameAutoQuxId")
	t.constraintNameAutoQuxId = constraintNameAutoQuxId
}
func (t *VendorConstraintNameAutoRef) SetConstraintNameAutoCorgeId(constraintNameAutoCorgeId df.Numeric) {
	t.AddPropertyName("constraintNameAutoCorgeId")
	t.constraintNameAutoCorgeId = constraintNameAutoCorgeId
}
func (t *VendorConstraintNameAutoRef) SetConstraintNameAutoUnique(constraintNameAutoUnique string) {
	t.AddPropertyName("constraintNameAutoUnique")
	t.constraintNameAutoUnique = constraintNameAutoUnique
}

func (t *VendorConstraintNameAutoRef) SetUp(){
	t.constraintNameAutoRefId.DecPoint = 0
	t.constraintNameAutoFooId.DecPoint = 0
	t.constraintNameAutoBarId.DecPoint = 0
	t.constraintNameAutoQuxId.DecPoint = 0
	t.constraintNameAutoCorgeId.DecPoint = 0
	
}
func (t *VendorConstraintNameAutoRef)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}
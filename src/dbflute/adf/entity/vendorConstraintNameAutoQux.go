package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorConstraintNameAutoQux struct {
	constraintNameAutoQuxId df.Numeric
	constraintNameAutoQuxName string
	df.BaseEntity
}

func CreateVendorConstraintNameAutoQux() *VendorConstraintNameAutoQux{
	vendorConstraintNameAutoQux :=new(VendorConstraintNameAutoQux)
	vendorConstraintNameAutoQux.SetUp()
	return vendorConstraintNameAutoQux 
}

func (l *VendorConstraintNameAutoQux) GetConstraintNameAutoQuxId () df.Numeric {
	return l.constraintNameAutoQuxId
}
func (l *VendorConstraintNameAutoQux) GetConstraintNameAutoQuxName () string {
	return l.constraintNameAutoQuxName
}

func (t *VendorConstraintNameAutoQux) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 2)
	i[0] = &(t.constraintNameAutoQuxId)
	i[1] = &(t.constraintNameAutoQuxName)
	return i
}


func (t *VendorConstraintNameAutoQux) AsTableDbName() string {
	return "VendorConstraintNameAutoQux"
}

func (t *VendorConstraintNameAutoQux) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("constraintNameAutoQuxId") == false {
            return false 
        }
        return true;
}
func (t *VendorConstraintNameAutoQux) SetConstraintNameAutoQuxId(constraintNameAutoQuxId df.Numeric) {
	t.AddPropertyName("constraintNameAutoQuxId")
	t.constraintNameAutoQuxId = constraintNameAutoQuxId
}
func (t *VendorConstraintNameAutoQux) SetConstraintNameAutoQuxName(constraintNameAutoQuxName string) {
	t.AddPropertyName("constraintNameAutoQuxName")
	t.constraintNameAutoQuxName = constraintNameAutoQuxName
}

func (t *VendorConstraintNameAutoQux) SetUp(){
	t.constraintNameAutoQuxId.DecPoint = 0
	
}
func (t *VendorConstraintNameAutoQux)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}
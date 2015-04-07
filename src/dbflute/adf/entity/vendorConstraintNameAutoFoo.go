package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorConstraintNameAutoFoo struct {
	constraintNameAutoFooId df.Numeric
	constraintNameAutoFooName string
	df.BaseEntity
}

func CreateVendorConstraintNameAutoFoo() *VendorConstraintNameAutoFoo{
	vendorConstraintNameAutoFoo :=new(VendorConstraintNameAutoFoo)
	vendorConstraintNameAutoFoo.SetUp()
	return vendorConstraintNameAutoFoo 
}

func (l *VendorConstraintNameAutoFoo) GetConstraintNameAutoFooId () df.Numeric {
	return l.constraintNameAutoFooId
}
func (l *VendorConstraintNameAutoFoo) GetConstraintNameAutoFooName () string {
	return l.constraintNameAutoFooName
}

func (t *VendorConstraintNameAutoFoo) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 2)
	i[0] = &(t.constraintNameAutoFooId)
	i[1] = &(t.constraintNameAutoFooName)
	return i
}


func (t *VendorConstraintNameAutoFoo) AsTableDbName() string {
	return "VendorConstraintNameAutoFoo"
}

func (t *VendorConstraintNameAutoFoo) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("constraintNameAutoFooId") == false {
            return false 
        }
        return true;
}
func (t *VendorConstraintNameAutoFoo) SetConstraintNameAutoFooId(constraintNameAutoFooId df.Numeric) {
	t.AddPropertyName("constraintNameAutoFooId")
	t.constraintNameAutoFooId = constraintNameAutoFooId
}
func (t *VendorConstraintNameAutoFoo) SetConstraintNameAutoFooName(constraintNameAutoFooName string) {
	t.AddPropertyName("constraintNameAutoFooName")
	t.constraintNameAutoFooName = constraintNameAutoFooName
}

func (t *VendorConstraintNameAutoFoo) SetUp(){
	t.constraintNameAutoFooId.DecPoint = 0
	
}
func (t *VendorConstraintNameAutoFoo)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}
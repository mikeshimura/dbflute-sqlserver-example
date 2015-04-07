package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorConstraintNameAutoBar struct {
	constraintNameAutoBarId df.Numeric
	constraintNameAutoBarName string
	df.BaseEntity
}

func CreateVendorConstraintNameAutoBar() *VendorConstraintNameAutoBar{
	vendorConstraintNameAutoBar :=new(VendorConstraintNameAutoBar)
	vendorConstraintNameAutoBar.SetUp()
	return vendorConstraintNameAutoBar 
}

func (l *VendorConstraintNameAutoBar) GetConstraintNameAutoBarId () df.Numeric {
	return l.constraintNameAutoBarId
}
func (l *VendorConstraintNameAutoBar) GetConstraintNameAutoBarName () string {
	return l.constraintNameAutoBarName
}

func (t *VendorConstraintNameAutoBar) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 2)
	i[0] = &(t.constraintNameAutoBarId)
	i[1] = &(t.constraintNameAutoBarName)
	return i
}


func (t *VendorConstraintNameAutoBar) AsTableDbName() string {
	return "VendorConstraintNameAutoBar"
}

func (t *VendorConstraintNameAutoBar) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("constraintNameAutoBarId") == false {
            return false 
        }
        return true;
}
func (t *VendorConstraintNameAutoBar) SetConstraintNameAutoBarId(constraintNameAutoBarId df.Numeric) {
	t.AddPropertyName("constraintNameAutoBarId")
	t.constraintNameAutoBarId = constraintNameAutoBarId
}
func (t *VendorConstraintNameAutoBar) SetConstraintNameAutoBarName(constraintNameAutoBarName string) {
	t.AddPropertyName("constraintNameAutoBarName")
	t.constraintNameAutoBarName = constraintNameAutoBarName
}

func (t *VendorConstraintNameAutoBar) SetUp(){
	t.constraintNameAutoBarId.DecPoint = 0
	
}
func (t *VendorConstraintNameAutoBar)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}
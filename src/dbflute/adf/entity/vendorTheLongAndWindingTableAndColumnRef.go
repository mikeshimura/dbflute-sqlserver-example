package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorTheLongAndWindingTableAndColumnRef struct {
	theLongAndWindingTableAndColumnRefId int64
	theLongAndWindingTableAndColumnId int64
	theLongAndWindingTableAndColumnRefDate df.MysqlDate
	shortDate df.MysqlDate
	df.BaseEntity
}

func CreateVendorTheLongAndWindingTableAndColumnRef() *VendorTheLongAndWindingTableAndColumnRef{
	vendorTheLongAndWindingTableAndColumnRef :=new(VendorTheLongAndWindingTableAndColumnRef)
	vendorTheLongAndWindingTableAndColumnRef.SetUp()
	return vendorTheLongAndWindingTableAndColumnRef 
}

func (l *VendorTheLongAndWindingTableAndColumnRef) GetTheLongAndWindingTableAndColumnRefId () int64 {
	return l.theLongAndWindingTableAndColumnRefId
}
func (l *VendorTheLongAndWindingTableAndColumnRef) GetTheLongAndWindingTableAndColumnId () int64 {
	return l.theLongAndWindingTableAndColumnId
}
func (l *VendorTheLongAndWindingTableAndColumnRef) GetTheLongAndWindingTableAndColumnRefDate () df.MysqlDate {
	return l.theLongAndWindingTableAndColumnRefDate
}
func (l *VendorTheLongAndWindingTableAndColumnRef) GetShortDate () df.MysqlDate {
	return l.shortDate
}

func (t *VendorTheLongAndWindingTableAndColumnRef) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 4)
	i[0] = &(t.theLongAndWindingTableAndColumnRefId)
	i[1] = &(t.theLongAndWindingTableAndColumnId)
	i[2] = &(t.theLongAndWindingTableAndColumnRefDate)
	i[3] = &(t.shortDate)
	return i
}


func (t *VendorTheLongAndWindingTableAndColumnRef) AsTableDbName() string {
	return "VendorTheLongAndWindingTableAndColumnRef"
}

func (t *VendorTheLongAndWindingTableAndColumnRef) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("theLongAndWindingTableAndColumnRefId") == false {
            return false 
        }
        return true;
}
func (t *VendorTheLongAndWindingTableAndColumnRef) SetTheLongAndWindingTableAndColumnRefId(theLongAndWindingTableAndColumnRefId int64) {
	t.AddPropertyName("theLongAndWindingTableAndColumnRefId")
	t.theLongAndWindingTableAndColumnRefId = theLongAndWindingTableAndColumnRefId
}
func (t *VendorTheLongAndWindingTableAndColumnRef) SetTheLongAndWindingTableAndColumnId(theLongAndWindingTableAndColumnId int64) {
	t.AddPropertyName("theLongAndWindingTableAndColumnId")
	t.theLongAndWindingTableAndColumnId = theLongAndWindingTableAndColumnId
}
func (t *VendorTheLongAndWindingTableAndColumnRef) SetTheLongAndWindingTableAndColumnRefDate(theLongAndWindingTableAndColumnRefDate df.MysqlDate) {
	t.AddPropertyName("theLongAndWindingTableAndColumnRefDate")
	t.theLongAndWindingTableAndColumnRefDate = theLongAndWindingTableAndColumnRefDate
}
func (t *VendorTheLongAndWindingTableAndColumnRef) SetShortDate(shortDate df.MysqlDate) {
	t.AddPropertyName("shortDate")
	t.shortDate = shortDate
}

func (t *VendorTheLongAndWindingTableAndColumnRef) SetUp(){
	
}
func (t *VendorTheLongAndWindingTableAndColumnRef)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}
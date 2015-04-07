package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorTheLongAndWindingTableAndColumn struct {
	theLongAndWindingTableAndColumnId int64
	theLongAndWindingTableAndColumnName string
	shortName string
	shortSize int64
	df.BaseEntity
}

func CreateVendorTheLongAndWindingTableAndColumn() *VendorTheLongAndWindingTableAndColumn{
	vendorTheLongAndWindingTableAndColumn :=new(VendorTheLongAndWindingTableAndColumn)
	vendorTheLongAndWindingTableAndColumn.SetUp()
	return vendorTheLongAndWindingTableAndColumn 
}

func (l *VendorTheLongAndWindingTableAndColumn) GetTheLongAndWindingTableAndColumnId () int64 {
	return l.theLongAndWindingTableAndColumnId
}
func (l *VendorTheLongAndWindingTableAndColumn) GetTheLongAndWindingTableAndColumnName () string {
	return l.theLongAndWindingTableAndColumnName
}
func (l *VendorTheLongAndWindingTableAndColumn) GetShortName () string {
	return l.shortName
}
func (l *VendorTheLongAndWindingTableAndColumn) GetShortSize () int64 {
	return l.shortSize
}

func (t *VendorTheLongAndWindingTableAndColumn) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 4)
	i[0] = &(t.theLongAndWindingTableAndColumnId)
	i[1] = &(t.theLongAndWindingTableAndColumnName)
	i[2] = &(t.shortName)
	i[3] = &(t.shortSize)
	return i
}


func (t *VendorTheLongAndWindingTableAndColumn) AsTableDbName() string {
	return "VendorTheLongAndWindingTableAndColumn"
}

func (t *VendorTheLongAndWindingTableAndColumn) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("theLongAndWindingTableAndColumnId") == false {
            return false 
        }
        return true;
}
func (t *VendorTheLongAndWindingTableAndColumn) SetTheLongAndWindingTableAndColumnId(theLongAndWindingTableAndColumnId int64) {
	t.AddPropertyName("theLongAndWindingTableAndColumnId")
	t.theLongAndWindingTableAndColumnId = theLongAndWindingTableAndColumnId
}
func (t *VendorTheLongAndWindingTableAndColumn) SetTheLongAndWindingTableAndColumnName(theLongAndWindingTableAndColumnName string) {
	t.AddPropertyName("theLongAndWindingTableAndColumnName")
	t.theLongAndWindingTableAndColumnName = theLongAndWindingTableAndColumnName
}
func (t *VendorTheLongAndWindingTableAndColumn) SetShortName(shortName string) {
	t.AddPropertyName("shortName")
	t.shortName = shortName
}
func (t *VendorTheLongAndWindingTableAndColumn) SetShortSize(shortSize int64) {
	t.AddPropertyName("shortSize")
	t.shortSize = shortSize
}

func (t *VendorTheLongAndWindingTableAndColumn) SetUp(){
	
}
func (t *VendorTheLongAndWindingTableAndColumn)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}
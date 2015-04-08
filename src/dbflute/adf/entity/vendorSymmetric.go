package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

type VendorSymmetric struct {
	vendorSymmetricId df.Numeric
	plainText df.NullString
	encryptedData []byte
	df.BaseEntity
}

func CreateVendorSymmetric() *VendorSymmetric{
	vendorSymmetric :=new(VendorSymmetric)
	vendorSymmetric.SetUp()
	return vendorSymmetric 
}

func (l *VendorSymmetric) GetVendorSymmetricId () df.Numeric {
	return l.vendorSymmetricId
}
func (l *VendorSymmetric) GetPlainText () df.NullString {
	return l.plainText
}
func (l *VendorSymmetric) GetEncryptedData () []byte {
	return l.encryptedData
}

func (t *VendorSymmetric) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 3)
	i[0] = &(t.vendorSymmetricId)
	i[1] = &(t.plainText)
	i[2] = &(t.encryptedData)
	return i
}


func (t *VendorSymmetric) AsTableDbName() string {
	return "VendorSymmetric"
}

func (t *VendorSymmetric) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("vendorSymmetricId") == false {
            return false 
        }
        return true;
}
func (t *VendorSymmetric) SetVendorSymmetricId(vendorSymmetricId df.Numeric) {
	t.AddPropertyName("vendorSymmetricId")
	t.vendorSymmetricId = vendorSymmetricId
}
func (t *VendorSymmetric) SetPlainText(plainText df.NullString) {
	t.AddPropertyName("plainText")
	t.plainText = plainText
}
func (t *VendorSymmetric) SetEncryptedData(encryptedData []byte) {
	t.AddPropertyName("encryptedData")
	t.encryptedData = encryptedData
}

func (t *VendorSymmetric) SetUp(){
	t.vendorSymmetricId.DecPoint = 0
	
}
func (t *VendorSymmetric)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}
package entity

import (
	"github.com/mikeshimura/dbflute/df"
	"database/sql"
)

type WhiteDelimiter struct {
	delimiterId int64
	numberNullable sql.NullInt64
	stringConverted string
	stringNonConverted df.NullString
	dateDefault df.Timestamp
	df.BaseEntity
}

func CreateWhiteDelimiter() *WhiteDelimiter{
	whiteDelimiter :=new(WhiteDelimiter)
	whiteDelimiter.SetUp()
	return whiteDelimiter 
}

func (l *WhiteDelimiter) GetDelimiterId () int64 {
	return l.delimiterId
}
func (l *WhiteDelimiter) GetNumberNullable () sql.NullInt64 {
	return l.numberNullable
}
func (l *WhiteDelimiter) GetStringConverted () string {
	return l.stringConverted
}
func (l *WhiteDelimiter) GetStringNonConverted () df.NullString {
	return l.stringNonConverted
}
func (l *WhiteDelimiter) GetDateDefault () df.Timestamp {
	return l.dateDefault
}

func (t *WhiteDelimiter) GetAsInterfaceArray() []interface{} {
	i := make([]interface{}, 5)
	i[0] = &(t.delimiterId)
	i[1] = &(t.numberNullable)
	i[2] = &(t.stringConverted)
	i[3] = &(t.stringNonConverted)
	i[4] = &(t.dateDefault)
	return i
}


func (t *WhiteDelimiter) AsTableDbName() string {
	return "WhiteDelimiter"
}

func (t *WhiteDelimiter) HasPrimaryKeyValue() bool{
        if t.IsModifiedProperty("delimiterId") == false {
            return false 
        }
        return true;
}
func (t *WhiteDelimiter) SetDelimiterId(delimiterId int64) {
	t.AddPropertyName("delimiterId")
	t.delimiterId = delimiterId
}
func (t *WhiteDelimiter) SetNumberNullable(numberNullable sql.NullInt64) {
	t.AddPropertyName("numberNullable")
	t.numberNullable = numberNullable
}
func (t *WhiteDelimiter) SetStringConverted(stringConverted string) {
	t.AddPropertyName("stringConverted")
	t.stringConverted = stringConverted
}
func (t *WhiteDelimiter) SetStringNonConverted(stringNonConverted df.NullString) {
	t.AddPropertyName("stringNonConverted")
	t.stringNonConverted = stringNonConverted
}
func (t *WhiteDelimiter) SetDateDefault(dateDefault df.Timestamp) {
	t.AddPropertyName("dateDefault")
	t.dateDefault = dateDefault
}

func (t *WhiteDelimiter) SetUp(){
	
}
func (t *WhiteDelimiter)GetDBMeta() *df.DBMeta{
	return df.DBMetaInstanceHandler_I.TableDbNameInstanceMap[t.AsTableDbName()]
}
package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

type WhiteDelimiterDbm_T struct {
	df.BaseDBMeta
	ColumnDelimiterId *df.ColumnInfo
	ColumnNumberNullable *df.ColumnInfo
	ColumnStringConverted *df.ColumnInfo
	ColumnStringNonConverted *df.ColumnInfo
	ColumnDateDefault *df.ColumnInfo
}

func (b *WhiteDelimiterDbm_T) GetProjectName() string {
	return df.DBCurrent_I.ProjectName
}

func (b *WhiteDelimiterDbm_T) GetDbCurrent() *df.DBCurrent {
	return df.DBCurrent_I
}

var WhiteDelimiterDbm *WhiteDelimiterDbm_T

func Create_WhiteDelimiterDbm() {
	WhiteDelimiterDbm = new(WhiteDelimiterDbm_T)
	WhiteDelimiterDbm.TableDbName = "WHITE_DELIMITER"
	WhiteDelimiterDbm.TableDispName = "WHITE_DELIMITER"
	WhiteDelimiterDbm.TablePropertyName = "whiteDelimiter"
	WhiteDelimiterDbm.TableSqlName = new(df.TableSqlName)
	WhiteDelimiterDbm.TableSqlName.TableSqlName = "exampledb.dbo.WHITE_DELIMITER"
	WhiteDelimiterDbm.TableSqlName.CorrespondingDbName = WhiteDelimiterDbm.TableDbName
	WhiteDelimiterDbm.Identity=true

	var whiteDelimiter df.DBMeta
	whiteDelimiter = WhiteDelimiterDbm
	WhiteDelimiterDbm.DBMeta=&whiteDelimiter
	delimiterIdSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo DELIMITER_ID
	delimiterIdSqlName.ColumnSqlName = "DELIMITER_ID"
	delimiterIdSqlName.IrregularChar = false
	WhiteDelimiterDbm.ColumnDelimiterId = df.CCI(&whiteDelimiter, "DELIMITER_ID", delimiterIdSqlName, "", "", "Long.class", "delimiterId", "", true, true,true, "bigint identity", 19, 0, "",false,"","", "","","",false,"int64")
	numberNullableSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo NUMBER_NULLABLE
	numberNullableSqlName.ColumnSqlName = "NUMBER_NULLABLE"
	numberNullableSqlName.IrregularChar = false
	WhiteDelimiterDbm.ColumnNumberNullable = df.CCI(&whiteDelimiter, "NUMBER_NULLABLE", numberNullableSqlName, "", "", "Integer.class", "numberNullable", "", false, false,false, "int", 10, 0, "",false,"","", "","","",false,"sql.NullInt64")
	stringConvertedSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo STRING_CONVERTED
	stringConvertedSqlName.ColumnSqlName = "STRING_CONVERTED"
	stringConvertedSqlName.IrregularChar = false
	WhiteDelimiterDbm.ColumnStringConverted = df.CCI(&whiteDelimiter, "STRING_CONVERTED", stringConvertedSqlName, "", "", "String.class", "stringConverted", "", false, false,true, "varchar", 200, 0, "",false,"","", "","","",false,"string")
	stringNonConvertedSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo STRING_NON_CONVERTED
	stringNonConvertedSqlName.ColumnSqlName = "STRING_NON_CONVERTED"
	stringNonConvertedSqlName.IrregularChar = false
	WhiteDelimiterDbm.ColumnStringNonConverted = df.CCI(&whiteDelimiter, "STRING_NON_CONVERTED", stringNonConvertedSqlName, "", "", "String.class", "stringNonConverted", "", false, false,false, "varchar", 200, 0, "",false,"","", "","","",false,"df.NullString")
	dateDefaultSqlName := new(df.ColumnSqlName)
	//colsqlname dayoo DATE_DEFAULT
	dateDefaultSqlName.ColumnSqlName = "DATE_DEFAULT"
	dateDefaultSqlName.IrregularChar = false
	WhiteDelimiterDbm.ColumnDateDefault = df.CCI(&whiteDelimiter, "DATE_DEFAULT", dateDefaultSqlName, "", "", "java.time.LocalDateTime.class", "dateDefault", "", false, false,true, "datetime", 23, 3, "",false,"","", "","","",false,"df.Timestamp")

	WhiteDelimiterDbm.ColumnInfoList = new(df.List)
	WhiteDelimiterDbm.ColumnInfoList.Add(WhiteDelimiterDbm.ColumnDelimiterId)
	WhiteDelimiterDbm.ColumnInfoList.Add(WhiteDelimiterDbm.ColumnNumberNullable)
	WhiteDelimiterDbm.ColumnInfoList.Add(WhiteDelimiterDbm.ColumnStringConverted)
	WhiteDelimiterDbm.ColumnInfoList.Add(WhiteDelimiterDbm.ColumnStringNonConverted)
	WhiteDelimiterDbm.ColumnInfoList.Add(WhiteDelimiterDbm.ColumnDateDefault)


	WhiteDelimiterDbm.ColumnInfoMap=make(map[string]int)
	WhiteDelimiterDbm.ColumnInfoMap["delimiterId"]=0
		WhiteDelimiterDbm.ColumnInfoMap["numberNullable"]=1
		WhiteDelimiterDbm.ColumnInfoMap["stringConverted"]=2
		WhiteDelimiterDbm.ColumnInfoMap["stringNonConverted"]=3
		WhiteDelimiterDbm.ColumnInfoMap["dateDefault"]=4
	    WhiteDelimiterDbm.PrimaryKey = true
    WhiteDelimiterDbm.CompoundPrimaryKey = false
	ui := new(df.UniqueInfo)
	ui.DbMeta = &whiteDelimiter
	ui.Primary = true
	ui.UniqueColumnList = new(df.List)
	ui.UniqueColumnList.Add(WhiteDelimiterDbm.ColumnDelimiterId)

	WhiteDelimiterDbm.PrimaryInfo = new(df.PrimaryInfo)
	WhiteDelimiterDbm.PrimaryInfo.UniqueInfo = ui
	
	var whiteDelimiterMeta df.DBMeta = WhiteDelimiterDbm
	df.DBMetaInstanceHandler_I.TableDbNameInstanceMap["WhiteDelimiter"] = &whiteDelimiterMeta
}

@echo off

set ANT_OPTS=-Xmx512m

set DBFLUTE_HOME=..\mydbflute\dbflute-1.1.0-sp1

set MY_PROJECT_NAME=mssql

set MY_PROPERTIES_PATH=build.properties
set JAVA_HOME=d:\pg\Java\jdk1.8.0_20

if "%pause_at_end%"=="" set pause_at_end=y

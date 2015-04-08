package com.mssoftech.dbflute.cbean.cq.bs;

import java.util.*;

import org.dbflute.cbean.*;
import org.dbflute.cbean.chelper.*;
import org.dbflute.cbean.ckey.*;
import org.dbflute.cbean.coption.*;
import org.dbflute.cbean.cvalue.ConditionValue;
import org.dbflute.cbean.ordering.*;
import org.dbflute.cbean.scoping.*;
import org.dbflute.cbean.sqlclause.SqlClause;
import org.dbflute.dbmeta.DBMetaProvider;
import com.mssoftech.dbflute.allcommon.*;
import com.mssoftech.dbflute.cbean.*;
import com.mssoftech.dbflute.cbean.cq.*;

/**
 * The abstract condition-query of WHITE_DELIMITER.
 * @author DBFlute(AutoGenerator)
 */
public abstract class AbstractBsWhiteDelimiterCQ extends AbstractConditionQuery {

    // ===================================================================================
    //                                                                         Constructor
    //                                                                         ===========
    public AbstractBsWhiteDelimiterCQ(ConditionQuery referrerQuery, SqlClause sqlClause, String aliasName, int nestLevel) {
        super(referrerQuery, sqlClause, aliasName, nestLevel);
    }

    // ===================================================================================
    //                                                                             DB Meta
    //                                                                             =======
    @Override
    protected DBMetaProvider xgetDBMetaProvider() {
        return DBMetaInstanceHandler.getProvider();
    }

    public String asTableDbName() {
        return "WHITE_DELIMITER";
    }

    // ===================================================================================
    //                                                                               Query
    //                                                                               =====
    /**
     * Equal(=). And NullIgnored, OnlyOnceRegistered. <br>
     * DELIMITER_ID: {PK, ID, NotNull, bigint identity(19)}
     * @param delimiterId The value of delimiterId as equal. (NullAllowed: if null, no condition)
     */
    public void setDelimiterId_Equal(Long delimiterId) {
        doSetDelimiterId_Equal(delimiterId);
    }

    protected void doSetDelimiterId_Equal(Long delimiterId) {
        regDelimiterId(CK_EQ, delimiterId);
    }

    /**
     * NotEqual(&lt;&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * DELIMITER_ID: {PK, ID, NotNull, bigint identity(19)}
     * @param delimiterId The value of delimiterId as notEqual. (NullAllowed: if null, no condition)
     */
    public void setDelimiterId_NotEqual(Long delimiterId) {
        doSetDelimiterId_NotEqual(delimiterId);
    }

    protected void doSetDelimiterId_NotEqual(Long delimiterId) {
        regDelimiterId(CK_NES, delimiterId);
    }

    /**
     * GreaterThan(&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * DELIMITER_ID: {PK, ID, NotNull, bigint identity(19)}
     * @param delimiterId The value of delimiterId as greaterThan. (NullAllowed: if null, no condition)
     */
    public void setDelimiterId_GreaterThan(Long delimiterId) {
        regDelimiterId(CK_GT, delimiterId);
    }

    /**
     * LessThan(&lt;). And NullIgnored, OnlyOnceRegistered. <br>
     * DELIMITER_ID: {PK, ID, NotNull, bigint identity(19)}
     * @param delimiterId The value of delimiterId as lessThan. (NullAllowed: if null, no condition)
     */
    public void setDelimiterId_LessThan(Long delimiterId) {
        regDelimiterId(CK_LT, delimiterId);
    }

    /**
     * GreaterEqual(&gt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * DELIMITER_ID: {PK, ID, NotNull, bigint identity(19)}
     * @param delimiterId The value of delimiterId as greaterEqual. (NullAllowed: if null, no condition)
     */
    public void setDelimiterId_GreaterEqual(Long delimiterId) {
        regDelimiterId(CK_GE, delimiterId);
    }

    /**
     * LessEqual(&lt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * DELIMITER_ID: {PK, ID, NotNull, bigint identity(19)}
     * @param delimiterId The value of delimiterId as lessEqual. (NullAllowed: if null, no condition)
     */
    public void setDelimiterId_LessEqual(Long delimiterId) {
        regDelimiterId(CK_LE, delimiterId);
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * DELIMITER_ID: {PK, ID, NotNull, bigint identity(19)}
     * @param minNumber The min number of delimiterId. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of delimiterId. (NullAllowed: if null, no to-condition)
     * @param opLambda The callback for option of range-of. (NotNull)
     */
    public void setDelimiterId_RangeOf(Long minNumber, Long maxNumber, ConditionOptionCall<RangeOfOption> opLambda) {
        setDelimiterId_RangeOf(minNumber, maxNumber, xcROOP(opLambda));
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * DELIMITER_ID: {PK, ID, NotNull, bigint identity(19)}
     * @param minNumber The min number of delimiterId. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of delimiterId. (NullAllowed: if null, no to-condition)
     * @param rangeOfOption The option of range-of. (NotNull)
     */
    protected void setDelimiterId_RangeOf(Long minNumber, Long maxNumber, RangeOfOption rangeOfOption) {
        regROO(minNumber, maxNumber, xgetCValueDelimiterId(), "DELIMITER_ID", rangeOfOption);
    }

    /**
     * InScope {in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * DELIMITER_ID: {PK, ID, NotNull, bigint identity(19)}
     * @param delimiterIdList The collection of delimiterId as inScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setDelimiterId_InScope(Collection<Long> delimiterIdList) {
        doSetDelimiterId_InScope(delimiterIdList);
    }

    protected void doSetDelimiterId_InScope(Collection<Long> delimiterIdList) {
        regINS(CK_INS, cTL(delimiterIdList), xgetCValueDelimiterId(), "DELIMITER_ID");
    }

    /**
     * NotInScope {not in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * DELIMITER_ID: {PK, ID, NotNull, bigint identity(19)}
     * @param delimiterIdList The collection of delimiterId as notInScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setDelimiterId_NotInScope(Collection<Long> delimiterIdList) {
        doSetDelimiterId_NotInScope(delimiterIdList);
    }

    protected void doSetDelimiterId_NotInScope(Collection<Long> delimiterIdList) {
        regINS(CK_NINS, cTL(delimiterIdList), xgetCValueDelimiterId(), "DELIMITER_ID");
    }

    /**
     * IsNull {is null}. And OnlyOnceRegistered. <br>
     * DELIMITER_ID: {PK, ID, NotNull, bigint identity(19)}
     */
    public void setDelimiterId_IsNull() { regDelimiterId(CK_ISN, DOBJ); }

    /**
     * IsNotNull {is not null}. And OnlyOnceRegistered. <br>
     * DELIMITER_ID: {PK, ID, NotNull, bigint identity(19)}
     */
    public void setDelimiterId_IsNotNull() { regDelimiterId(CK_ISNN, DOBJ); }

    protected void regDelimiterId(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueDelimiterId(), "DELIMITER_ID"); }
    protected abstract ConditionValue xgetCValueDelimiterId();

    /**
     * Equal(=). And NullIgnored, OnlyOnceRegistered. <br>
     * NUMBER_NULLABLE: {int(10)}
     * @param numberNullable The value of numberNullable as equal. (NullAllowed: if null, no condition)
     */
    public void setNumberNullable_Equal(Integer numberNullable) {
        doSetNumberNullable_Equal(numberNullable);
    }

    protected void doSetNumberNullable_Equal(Integer numberNullable) {
        regNumberNullable(CK_EQ, numberNullable);
    }

    /**
     * NotEqual(&lt;&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * NUMBER_NULLABLE: {int(10)}
     * @param numberNullable The value of numberNullable as notEqual. (NullAllowed: if null, no condition)
     */
    public void setNumberNullable_NotEqual(Integer numberNullable) {
        doSetNumberNullable_NotEqual(numberNullable);
    }

    protected void doSetNumberNullable_NotEqual(Integer numberNullable) {
        regNumberNullable(CK_NES, numberNullable);
    }

    /**
     * GreaterThan(&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * NUMBER_NULLABLE: {int(10)}
     * @param numberNullable The value of numberNullable as greaterThan. (NullAllowed: if null, no condition)
     */
    public void setNumberNullable_GreaterThan(Integer numberNullable) {
        regNumberNullable(CK_GT, numberNullable);
    }

    /**
     * LessThan(&lt;). And NullIgnored, OnlyOnceRegistered. <br>
     * NUMBER_NULLABLE: {int(10)}
     * @param numberNullable The value of numberNullable as lessThan. (NullAllowed: if null, no condition)
     */
    public void setNumberNullable_LessThan(Integer numberNullable) {
        regNumberNullable(CK_LT, numberNullable);
    }

    /**
     * GreaterEqual(&gt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * NUMBER_NULLABLE: {int(10)}
     * @param numberNullable The value of numberNullable as greaterEqual. (NullAllowed: if null, no condition)
     */
    public void setNumberNullable_GreaterEqual(Integer numberNullable) {
        regNumberNullable(CK_GE, numberNullable);
    }

    /**
     * LessEqual(&lt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * NUMBER_NULLABLE: {int(10)}
     * @param numberNullable The value of numberNullable as lessEqual. (NullAllowed: if null, no condition)
     */
    public void setNumberNullable_LessEqual(Integer numberNullable) {
        regNumberNullable(CK_LE, numberNullable);
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * NUMBER_NULLABLE: {int(10)}
     * @param minNumber The min number of numberNullable. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of numberNullable. (NullAllowed: if null, no to-condition)
     * @param opLambda The callback for option of range-of. (NotNull)
     */
    public void setNumberNullable_RangeOf(Integer minNumber, Integer maxNumber, ConditionOptionCall<RangeOfOption> opLambda) {
        setNumberNullable_RangeOf(minNumber, maxNumber, xcROOP(opLambda));
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * NUMBER_NULLABLE: {int(10)}
     * @param minNumber The min number of numberNullable. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of numberNullable. (NullAllowed: if null, no to-condition)
     * @param rangeOfOption The option of range-of. (NotNull)
     */
    protected void setNumberNullable_RangeOf(Integer minNumber, Integer maxNumber, RangeOfOption rangeOfOption) {
        regROO(minNumber, maxNumber, xgetCValueNumberNullable(), "NUMBER_NULLABLE", rangeOfOption);
    }

    /**
     * InScope {in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * NUMBER_NULLABLE: {int(10)}
     * @param numberNullableList The collection of numberNullable as inScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setNumberNullable_InScope(Collection<Integer> numberNullableList) {
        doSetNumberNullable_InScope(numberNullableList);
    }

    protected void doSetNumberNullable_InScope(Collection<Integer> numberNullableList) {
        regINS(CK_INS, cTL(numberNullableList), xgetCValueNumberNullable(), "NUMBER_NULLABLE");
    }

    /**
     * NotInScope {not in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * NUMBER_NULLABLE: {int(10)}
     * @param numberNullableList The collection of numberNullable as notInScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setNumberNullable_NotInScope(Collection<Integer> numberNullableList) {
        doSetNumberNullable_NotInScope(numberNullableList);
    }

    protected void doSetNumberNullable_NotInScope(Collection<Integer> numberNullableList) {
        regINS(CK_NINS, cTL(numberNullableList), xgetCValueNumberNullable(), "NUMBER_NULLABLE");
    }

    /**
     * IsNull {is null}. And OnlyOnceRegistered. <br>
     * NUMBER_NULLABLE: {int(10)}
     */
    public void setNumberNullable_IsNull() { regNumberNullable(CK_ISN, DOBJ); }

    /**
     * IsNotNull {is not null}. And OnlyOnceRegistered. <br>
     * NUMBER_NULLABLE: {int(10)}
     */
    public void setNumberNullable_IsNotNull() { regNumberNullable(CK_ISNN, DOBJ); }

    protected void regNumberNullable(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueNumberNullable(), "NUMBER_NULLABLE"); }
    protected abstract ConditionValue xgetCValueNumberNullable();

    /**
     * Equal(=). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_CONVERTED: {NotNull, varchar(200)}
     * @param stringConverted The value of stringConverted as equal. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringConverted_Equal(String stringConverted) {
        doSetStringConverted_Equal(fRES(stringConverted));
    }

    protected void doSetStringConverted_Equal(String stringConverted) {
        regStringConverted(CK_EQ, stringConverted);
    }

    /**
     * NotEqual(&lt;&gt;). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_CONVERTED: {NotNull, varchar(200)}
     * @param stringConverted The value of stringConverted as notEqual. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringConverted_NotEqual(String stringConverted) {
        doSetStringConverted_NotEqual(fRES(stringConverted));
    }

    protected void doSetStringConverted_NotEqual(String stringConverted) {
        regStringConverted(CK_NES, stringConverted);
    }

    /**
     * GreaterThan(&gt;). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_CONVERTED: {NotNull, varchar(200)}
     * @param stringConverted The value of stringConverted as greaterThan. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringConverted_GreaterThan(String stringConverted) {
        regStringConverted(CK_GT, fRES(stringConverted));
    }

    /**
     * LessThan(&lt;). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_CONVERTED: {NotNull, varchar(200)}
     * @param stringConverted The value of stringConverted as lessThan. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringConverted_LessThan(String stringConverted) {
        regStringConverted(CK_LT, fRES(stringConverted));
    }

    /**
     * GreaterEqual(&gt;=). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_CONVERTED: {NotNull, varchar(200)}
     * @param stringConverted The value of stringConverted as greaterEqual. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringConverted_GreaterEqual(String stringConverted) {
        regStringConverted(CK_GE, fRES(stringConverted));
    }

    /**
     * LessEqual(&lt;=). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_CONVERTED: {NotNull, varchar(200)}
     * @param stringConverted The value of stringConverted as lessEqual. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringConverted_LessEqual(String stringConverted) {
        regStringConverted(CK_LE, fRES(stringConverted));
    }

    /**
     * InScope {in ('a', 'b')}. And NullOrEmptyIgnored, NullOrEmptyElementIgnored, SeveralRegistered. <br>
     * STRING_CONVERTED: {NotNull, varchar(200)}
     * @param stringConvertedList The collection of stringConverted as inScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringConverted_InScope(Collection<String> stringConvertedList) {
        doSetStringConverted_InScope(stringConvertedList);
    }

    protected void doSetStringConverted_InScope(Collection<String> stringConvertedList) {
        regINS(CK_INS, cTL(stringConvertedList), xgetCValueStringConverted(), "STRING_CONVERTED");
    }

    /**
     * NotInScope {not in ('a', 'b')}. And NullOrEmptyIgnored, NullOrEmptyElementIgnored, SeveralRegistered. <br>
     * STRING_CONVERTED: {NotNull, varchar(200)}
     * @param stringConvertedList The collection of stringConverted as notInScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringConverted_NotInScope(Collection<String> stringConvertedList) {
        doSetStringConverted_NotInScope(stringConvertedList);
    }

    protected void doSetStringConverted_NotInScope(Collection<String> stringConvertedList) {
        regINS(CK_NINS, cTL(stringConvertedList), xgetCValueStringConverted(), "STRING_CONVERTED");
    }

    /**
     * LikeSearch with various options. (versatile) {like '%xxx%' escape ...}. And NullOrEmptyIgnored, SeveralRegistered. <br>
     * STRING_CONVERTED: {NotNull, varchar(200)} <br>
     * <pre>e.g. setStringConverted_LikeSearch("xxx", op <span style="color: #90226C; font-weight: bold"><span style="font-size: 120%">-</span>&gt;</span> op.<span style="color: #CC4747">likeContain()</span>);</pre>
     * @param stringConverted The value of stringConverted as likeSearch. (NullAllowed: if null (or empty), no condition)
     * @param opLambda The callback for option of like-search. (NotNull)
     */
    public void setStringConverted_LikeSearch(String stringConverted, ConditionOptionCall<LikeSearchOption> opLambda) {
        setStringConverted_LikeSearch(stringConverted, xcLSOP(opLambda));
    }

    /**
     * LikeSearch with various options. (versatile) {like '%xxx%' escape ...}. And NullOrEmptyIgnored, SeveralRegistered. <br>
     * STRING_CONVERTED: {NotNull, varchar(200)} <br>
     * <pre>e.g. setStringConverted_LikeSearch("xxx", new <span style="color: #CC4747">LikeSearchOption</span>().likeContain());</pre>
     * @param stringConverted The value of stringConverted as likeSearch. (NullAllowed: if null (or empty), no condition)
     * @param likeSearchOption The option of like-search. (NotNull)
     */
    protected void setStringConverted_LikeSearch(String stringConverted, LikeSearchOption likeSearchOption) {
        regLSQ(CK_LS, fRES(stringConverted), xgetCValueStringConverted(), "STRING_CONVERTED", likeSearchOption);
    }

    /**
     * NotLikeSearch with various options. (versatile) {not like 'xxx%' escape ...} <br>
     * And NullOrEmptyIgnored, SeveralRegistered. <br>
     * STRING_CONVERTED: {NotNull, varchar(200)}
     * @param stringConverted The value of stringConverted as notLikeSearch. (NullAllowed: if null (or empty), no condition)
     * @param opLambda The callback for option of like-search. (NotNull)
     */
    public void setStringConverted_NotLikeSearch(String stringConverted, ConditionOptionCall<LikeSearchOption> opLambda) {
        setStringConverted_NotLikeSearch(stringConverted, xcLSOP(opLambda));
    }

    /**
     * NotLikeSearch with various options. (versatile) {not like 'xxx%' escape ...} <br>
     * And NullOrEmptyIgnored, SeveralRegistered. <br>
     * STRING_CONVERTED: {NotNull, varchar(200)}
     * @param stringConverted The value of stringConverted as notLikeSearch. (NullAllowed: if null (or empty), no condition)
     * @param likeSearchOption The option of not-like-search. (NotNull)
     */
    protected void setStringConverted_NotLikeSearch(String stringConverted, LikeSearchOption likeSearchOption) {
        regLSQ(CK_NLS, fRES(stringConverted), xgetCValueStringConverted(), "STRING_CONVERTED", likeSearchOption);
    }

    protected void regStringConverted(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueStringConverted(), "STRING_CONVERTED"); }
    protected abstract ConditionValue xgetCValueStringConverted();

    /**
     * Equal(=). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_NON_CONVERTED: {varchar(200)}
     * @param stringNonConverted The value of stringNonConverted as equal. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringNonConverted_Equal(String stringNonConverted) {
        doSetStringNonConverted_Equal(fRES(stringNonConverted));
    }

    protected void doSetStringNonConverted_Equal(String stringNonConverted) {
        regStringNonConverted(CK_EQ, stringNonConverted);
    }

    /**
     * NotEqual(&lt;&gt;). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_NON_CONVERTED: {varchar(200)}
     * @param stringNonConverted The value of stringNonConverted as notEqual. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringNonConverted_NotEqual(String stringNonConverted) {
        doSetStringNonConverted_NotEqual(fRES(stringNonConverted));
    }

    protected void doSetStringNonConverted_NotEqual(String stringNonConverted) {
        regStringNonConverted(CK_NES, stringNonConverted);
    }

    /**
     * GreaterThan(&gt;). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_NON_CONVERTED: {varchar(200)}
     * @param stringNonConverted The value of stringNonConverted as greaterThan. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringNonConverted_GreaterThan(String stringNonConverted) {
        regStringNonConverted(CK_GT, fRES(stringNonConverted));
    }

    /**
     * LessThan(&lt;). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_NON_CONVERTED: {varchar(200)}
     * @param stringNonConverted The value of stringNonConverted as lessThan. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringNonConverted_LessThan(String stringNonConverted) {
        regStringNonConverted(CK_LT, fRES(stringNonConverted));
    }

    /**
     * GreaterEqual(&gt;=). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_NON_CONVERTED: {varchar(200)}
     * @param stringNonConverted The value of stringNonConverted as greaterEqual. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringNonConverted_GreaterEqual(String stringNonConverted) {
        regStringNonConverted(CK_GE, fRES(stringNonConverted));
    }

    /**
     * LessEqual(&lt;=). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_NON_CONVERTED: {varchar(200)}
     * @param stringNonConverted The value of stringNonConverted as lessEqual. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringNonConverted_LessEqual(String stringNonConverted) {
        regStringNonConverted(CK_LE, fRES(stringNonConverted));
    }

    /**
     * InScope {in ('a', 'b')}. And NullOrEmptyIgnored, NullOrEmptyElementIgnored, SeveralRegistered. <br>
     * STRING_NON_CONVERTED: {varchar(200)}
     * @param stringNonConvertedList The collection of stringNonConverted as inScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringNonConverted_InScope(Collection<String> stringNonConvertedList) {
        doSetStringNonConverted_InScope(stringNonConvertedList);
    }

    protected void doSetStringNonConverted_InScope(Collection<String> stringNonConvertedList) {
        regINS(CK_INS, cTL(stringNonConvertedList), xgetCValueStringNonConverted(), "STRING_NON_CONVERTED");
    }

    /**
     * NotInScope {not in ('a', 'b')}. And NullOrEmptyIgnored, NullOrEmptyElementIgnored, SeveralRegistered. <br>
     * STRING_NON_CONVERTED: {varchar(200)}
     * @param stringNonConvertedList The collection of stringNonConverted as notInScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringNonConverted_NotInScope(Collection<String> stringNonConvertedList) {
        doSetStringNonConverted_NotInScope(stringNonConvertedList);
    }

    protected void doSetStringNonConverted_NotInScope(Collection<String> stringNonConvertedList) {
        regINS(CK_NINS, cTL(stringNonConvertedList), xgetCValueStringNonConverted(), "STRING_NON_CONVERTED");
    }

    /**
     * LikeSearch with various options. (versatile) {like '%xxx%' escape ...}. And NullOrEmptyIgnored, SeveralRegistered. <br>
     * STRING_NON_CONVERTED: {varchar(200)} <br>
     * <pre>e.g. setStringNonConverted_LikeSearch("xxx", op <span style="color: #90226C; font-weight: bold"><span style="font-size: 120%">-</span>&gt;</span> op.<span style="color: #CC4747">likeContain()</span>);</pre>
     * @param stringNonConverted The value of stringNonConverted as likeSearch. (NullAllowed: if null (or empty), no condition)
     * @param opLambda The callback for option of like-search. (NotNull)
     */
    public void setStringNonConverted_LikeSearch(String stringNonConverted, ConditionOptionCall<LikeSearchOption> opLambda) {
        setStringNonConverted_LikeSearch(stringNonConverted, xcLSOP(opLambda));
    }

    /**
     * LikeSearch with various options. (versatile) {like '%xxx%' escape ...}. And NullOrEmptyIgnored, SeveralRegistered. <br>
     * STRING_NON_CONVERTED: {varchar(200)} <br>
     * <pre>e.g. setStringNonConverted_LikeSearch("xxx", new <span style="color: #CC4747">LikeSearchOption</span>().likeContain());</pre>
     * @param stringNonConverted The value of stringNonConverted as likeSearch. (NullAllowed: if null (or empty), no condition)
     * @param likeSearchOption The option of like-search. (NotNull)
     */
    protected void setStringNonConverted_LikeSearch(String stringNonConverted, LikeSearchOption likeSearchOption) {
        regLSQ(CK_LS, fRES(stringNonConverted), xgetCValueStringNonConverted(), "STRING_NON_CONVERTED", likeSearchOption);
    }

    /**
     * NotLikeSearch with various options. (versatile) {not like 'xxx%' escape ...} <br>
     * And NullOrEmptyIgnored, SeveralRegistered. <br>
     * STRING_NON_CONVERTED: {varchar(200)}
     * @param stringNonConverted The value of stringNonConverted as notLikeSearch. (NullAllowed: if null (or empty), no condition)
     * @param opLambda The callback for option of like-search. (NotNull)
     */
    public void setStringNonConverted_NotLikeSearch(String stringNonConverted, ConditionOptionCall<LikeSearchOption> opLambda) {
        setStringNonConverted_NotLikeSearch(stringNonConverted, xcLSOP(opLambda));
    }

    /**
     * NotLikeSearch with various options. (versatile) {not like 'xxx%' escape ...} <br>
     * And NullOrEmptyIgnored, SeveralRegistered. <br>
     * STRING_NON_CONVERTED: {varchar(200)}
     * @param stringNonConverted The value of stringNonConverted as notLikeSearch. (NullAllowed: if null (or empty), no condition)
     * @param likeSearchOption The option of not-like-search. (NotNull)
     */
    protected void setStringNonConverted_NotLikeSearch(String stringNonConverted, LikeSearchOption likeSearchOption) {
        regLSQ(CK_NLS, fRES(stringNonConverted), xgetCValueStringNonConverted(), "STRING_NON_CONVERTED", likeSearchOption);
    }

    /**
     * IsNull {is null}. And OnlyOnceRegistered. <br>
     * STRING_NON_CONVERTED: {varchar(200)}
     */
    public void setStringNonConverted_IsNull() { regStringNonConverted(CK_ISN, DOBJ); }

    /**
     * IsNullOrEmpty {is null or empty}. And OnlyOnceRegistered. <br>
     * STRING_NON_CONVERTED: {varchar(200)}
     */
    public void setStringNonConverted_IsNullOrEmpty() { regStringNonConverted(CK_ISNOE, DOBJ); }

    /**
     * IsNotNull {is not null}. And OnlyOnceRegistered. <br>
     * STRING_NON_CONVERTED: {varchar(200)}
     */
    public void setStringNonConverted_IsNotNull() { regStringNonConverted(CK_ISNN, DOBJ); }

    protected void regStringNonConverted(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueStringNonConverted(), "STRING_NON_CONVERTED"); }
    protected abstract ConditionValue xgetCValueStringNonConverted();

    /**
     * Equal(=). And NullIgnored, OnlyOnceRegistered. <br>
     * DATE_DEFAULT: {NotNull, datetime(23, 3)}
     * @param dateDefault The value of dateDefault as equal. (NullAllowed: if null, no condition)
     */
    public void setDateDefault_Equal(java.time.LocalDateTime dateDefault) {
        regDateDefault(CK_EQ,  dateDefault);
    }

    /**
     * GreaterThan(&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * DATE_DEFAULT: {NotNull, datetime(23, 3)}
     * @param dateDefault The value of dateDefault as greaterThan. (NullAllowed: if null, no condition)
     */
    public void setDateDefault_GreaterThan(java.time.LocalDateTime dateDefault) {
        regDateDefault(CK_GT,  dateDefault);
    }

    /**
     * LessThan(&lt;). And NullIgnored, OnlyOnceRegistered. <br>
     * DATE_DEFAULT: {NotNull, datetime(23, 3)}
     * @param dateDefault The value of dateDefault as lessThan. (NullAllowed: if null, no condition)
     */
    public void setDateDefault_LessThan(java.time.LocalDateTime dateDefault) {
        regDateDefault(CK_LT,  dateDefault);
    }

    /**
     * GreaterEqual(&gt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * DATE_DEFAULT: {NotNull, datetime(23, 3)}
     * @param dateDefault The value of dateDefault as greaterEqual. (NullAllowed: if null, no condition)
     */
    public void setDateDefault_GreaterEqual(java.time.LocalDateTime dateDefault) {
        regDateDefault(CK_GE,  dateDefault);
    }

    /**
     * LessEqual(&lt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * DATE_DEFAULT: {NotNull, datetime(23, 3)}
     * @param dateDefault The value of dateDefault as lessEqual. (NullAllowed: if null, no condition)
     */
    public void setDateDefault_LessEqual(java.time.LocalDateTime dateDefault) {
        regDateDefault(CK_LE, dateDefault);
    }

    /**
     * FromTo with various options. (versatile) {(default) fromDatetime &lt;= column &lt;= toDatetime} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * DATE_DEFAULT: {NotNull, datetime(23, 3)}
     * <pre>e.g. setDateDefault_FromTo(fromDate, toDate, op <span style="color: #90226C; font-weight: bold"><span style="font-size: 120%">-</span>&gt;</span> op.<span style="color: #CC4747">compareAsDate()</span>);</pre>
     * @param fromDatetime The from-datetime(yyyy/MM/dd HH:mm:ss.SSS) of dateDefault. (NullAllowed: if null, no from-condition)
     * @param toDatetime The to-datetime(yyyy/MM/dd HH:mm:ss.SSS) of dateDefault. (NullAllowed: if null, no to-condition)
     * @param opLambda The callback for option of from-to. (NotNull)
     */
    public void setDateDefault_FromTo(java.time.LocalDateTime fromDatetime, java.time.LocalDateTime toDatetime, ConditionOptionCall<FromToOption> opLambda) {
        setDateDefault_FromTo(fromDatetime, toDatetime, xcFTOP(opLambda));
    }

    /**
     * FromTo with various options. (versatile) {(default) fromDatetime &lt;= column &lt;= toDatetime} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * DATE_DEFAULT: {NotNull, datetime(23, 3)}
     * <pre>e.g. setDateDefault_FromTo(fromDate, toDate, new <span style="color: #CC4747">FromToOption</span>().compareAsDate());</pre>
     * @param fromDatetime The from-datetime(yyyy/MM/dd HH:mm:ss.SSS) of dateDefault. (NullAllowed: if null, no from-condition)
     * @param toDatetime The to-datetime(yyyy/MM/dd HH:mm:ss.SSS) of dateDefault. (NullAllowed: if null, no to-condition)
     * @param fromToOption The option of from-to. (NotNull)
     */
    protected void setDateDefault_FromTo(java.time.LocalDateTime fromDatetime, java.time.LocalDateTime toDatetime, FromToOption fromToOption) {
        String nm = "DATE_DEFAULT"; FromToOption op = fromToOption;
        regFTQ(xfFTHD(fromDatetime, nm, op), xfFTHD(toDatetime, nm, op), xgetCValueDateDefault(), nm, op);
    }

    protected void regDateDefault(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueDateDefault(), "DATE_DEFAULT"); }
    protected abstract ConditionValue xgetCValueDateDefault();

    // ===================================================================================
    //                                                                     ScalarCondition
    //                                                                     ===============
    /**
     * Prepare ScalarCondition as equal. <br>
     * {where FOO = (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_Equal()</span>.max(new SubQuery&lt;WhiteDelimiterCB&gt;() {
     *     public void query(WhiteDelimiterCB subCB) {
     *         subCB.specify().setXxx... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setYyy...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<WhiteDelimiterCB> scalar_Equal() {
        return xcreateSSQFunction(CK_EQ, WhiteDelimiterCB.class);
    }

    /**
     * Prepare ScalarCondition as equal. <br>
     * {where FOO &lt;&gt; (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_NotEqual()</span>.max(new SubQuery&lt;WhiteDelimiterCB&gt;() {
     *     public void query(WhiteDelimiterCB subCB) {
     *         subCB.specify().setXxx... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setYyy...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<WhiteDelimiterCB> scalar_NotEqual() {
        return xcreateSSQFunction(CK_NES, WhiteDelimiterCB.class);
    }

    /**
     * Prepare ScalarCondition as greaterThan. <br>
     * {where FOO &gt; (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_GreaterThan()</span>.max(new SubQuery&lt;WhiteDelimiterCB&gt;() {
     *     public void query(WhiteDelimiterCB subCB) {
     *         subCB.specify().setFoo... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setBar...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<WhiteDelimiterCB> scalar_GreaterThan() {
        return xcreateSSQFunction(CK_GT, WhiteDelimiterCB.class);
    }

    /**
     * Prepare ScalarCondition as lessThan. <br>
     * {where FOO &lt; (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_LessThan()</span>.max(new SubQuery&lt;WhiteDelimiterCB&gt;() {
     *     public void query(WhiteDelimiterCB subCB) {
     *         subCB.specify().setFoo... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setBar...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<WhiteDelimiterCB> scalar_LessThan() {
        return xcreateSSQFunction(CK_LT, WhiteDelimiterCB.class);
    }

    /**
     * Prepare ScalarCondition as greaterEqual. <br>
     * {where FOO &gt;= (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_GreaterEqual()</span>.max(new SubQuery&lt;WhiteDelimiterCB&gt;() {
     *     public void query(WhiteDelimiterCB subCB) {
     *         subCB.specify().setFoo... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setBar...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<WhiteDelimiterCB> scalar_GreaterEqual() {
        return xcreateSSQFunction(CK_GE, WhiteDelimiterCB.class);
    }

    /**
     * Prepare ScalarCondition as lessEqual. <br>
     * {where FOO &lt;= (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_LessEqual()</span>.max(new SubQuery&lt;WhiteDelimiterCB&gt;() {
     *     public void query(WhiteDelimiterCB subCB) {
     *         subCB.specify().setFoo... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setBar...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<WhiteDelimiterCB> scalar_LessEqual() {
        return xcreateSSQFunction(CK_LE, WhiteDelimiterCB.class);
    }

    @SuppressWarnings("unchecked")
    protected <CB extends ConditionBean> void xscalarCondition(String fn, SubQuery<CB> sq, String rd, HpSSQOption<CB> op) {
        assertObjectNotNull("subQuery", sq);
        WhiteDelimiterCB cb = xcreateScalarConditionCB(); sq.query((CB)cb);
        String pp = keepScalarCondition(cb.query()); // for saving query-value
        op.setPartitionByCBean((CB)xcreateScalarConditionPartitionByCB()); // for using partition-by
        registerScalarCondition(fn, cb.query(), pp, rd, op);
    }
    public abstract String keepScalarCondition(WhiteDelimiterCQ sq);

    protected WhiteDelimiterCB xcreateScalarConditionCB() {
        WhiteDelimiterCB cb = newMyCB(); cb.xsetupForScalarCondition(this); return cb;
    }

    protected WhiteDelimiterCB xcreateScalarConditionPartitionByCB() {
        WhiteDelimiterCB cb = newMyCB(); cb.xsetupForScalarConditionPartitionBy(this); return cb;
    }

    // ===================================================================================
    //                                                                       MyselfDerived
    //                                                                       =============
    public void xsmyselfDerive(String fn, SubQuery<WhiteDelimiterCB> sq, String al, DerivedReferrerOption op) {
        assertObjectNotNull("subQuery", sq);
        WhiteDelimiterCB cb = new WhiteDelimiterCB(); cb.xsetupForDerivedReferrer(this);
        lockCall(() -> sq.query(cb)); String pp = keepSpecifyMyselfDerived(cb.query()); String pk = "DELIMITER_ID";
        registerSpecifyMyselfDerived(fn, cb.query(), pk, pk, pp, "myselfDerived", al, op);
    }
    public abstract String keepSpecifyMyselfDerived(WhiteDelimiterCQ sq);

    /**
     * Prepare for (Query)MyselfDerived (correlated sub-query).
     * @return The object to set up a function for myself table. (NotNull)
     */
    public HpQDRFunction<WhiteDelimiterCB> myselfDerived() {
        return xcreateQDRFunctionMyselfDerived(WhiteDelimiterCB.class);
    }
    @SuppressWarnings("unchecked")
    protected <CB extends ConditionBean> void xqderiveMyselfDerived(String fn, SubQuery<CB> sq, String rd, Object vl, DerivedReferrerOption op) {
        assertObjectNotNull("subQuery", sq);
        WhiteDelimiterCB cb = new WhiteDelimiterCB(); cb.xsetupForDerivedReferrer(this); sq.query((CB)cb);
        String pk = "DELIMITER_ID";
        String sqpp = keepQueryMyselfDerived(cb.query()); // for saving query-value.
        String prpp = keepQueryMyselfDerivedParameter(vl);
        registerQueryMyselfDerived(fn, cb.query(), pk, pk, sqpp, "myselfDerived", rd, vl, prpp, op);
    }
    public abstract String keepQueryMyselfDerived(WhiteDelimiterCQ sq);
    public abstract String keepQueryMyselfDerivedParameter(Object vl);

    // ===================================================================================
    //                                                                        MyselfExists
    //                                                                        ============
    /**
     * Prepare for MyselfExists (correlated sub-query).
     * @param subCBLambda The implementation of sub-query. (NotNull)
     */
    public void myselfExists(SubQuery<WhiteDelimiterCB> subCBLambda) {
        assertObjectNotNull("subCBLambda", subCBLambda);
        WhiteDelimiterCB cb = new WhiteDelimiterCB(); cb.xsetupForMyselfExists(this);
        lockCall(() -> subCBLambda.query(cb)); String pp = keepMyselfExists(cb.query());
        registerMyselfExists(cb.query(), pp);
    }
    public abstract String keepMyselfExists(WhiteDelimiterCQ sq);

    // ===================================================================================
    //                                                                        Manual Order
    //                                                                        ============
    /**
     * Order along manual ordering information.
     * <pre>
     * cb.query().addOrderBy_Birthdate_Asc().<span style="color: #CC4747">withManualOrder</span>(<span style="color: #553000">op</span> <span style="color: #90226C; font-weight: bold"><span style="font-size: 120%">-</span>&gt;</span> {
     *     <span style="color: #553000">op</span>.<span style="color: #CC4747">when_GreaterEqual</span>(priorityDate); <span style="color: #3F7E5E">// e.g. 2000/01/01</span>
     * });
     * <span style="color: #3F7E5E">// order by </span>
     * <span style="color: #3F7E5E">//   case</span>
     * <span style="color: #3F7E5E">//     when BIRTHDATE &gt;= '2000/01/01' then 0</span>
     * <span style="color: #3F7E5E">//     else 1</span>
     * <span style="color: #3F7E5E">//   end asc, ...</span>
     *
     * cb.query().addOrderBy_MemberStatusCode_Asc().<span style="color: #CC4747">withManualOrder</span>(<span style="color: #553000">op</span> <span style="color: #90226C; font-weight: bold"><span style="font-size: 120%">-</span>&gt;</span> {
     *     <span style="color: #553000">op</span>.<span style="color: #CC4747">when_GreaterEqual</span>(priorityDate); <span style="color: #3F7E5E">// e.g. 2000/01/01</span>
     *     <span style="color: #553000">op</span>.<span style="color: #CC4747">when_Equal</span>(CDef.MemberStatus.Withdrawal);
     *     <span style="color: #553000">op</span>.<span style="color: #CC4747">when_Equal</span>(CDef.MemberStatus.Formalized);
     *     <span style="color: #553000">op</span>.<span style="color: #CC4747">when_Equal</span>(CDef.MemberStatus.Provisional);
     * });
     * <span style="color: #3F7E5E">// order by </span>
     * <span style="color: #3F7E5E">//   case</span>
     * <span style="color: #3F7E5E">//     when MEMBER_STATUS_CODE = 'WDL' then 0</span>
     * <span style="color: #3F7E5E">//     when MEMBER_STATUS_CODE = 'FML' then 1</span>
     * <span style="color: #3F7E5E">//     when MEMBER_STATUS_CODE = 'PRV' then 2</span>
     * <span style="color: #3F7E5E">//     else 3</span>
     * <span style="color: #3F7E5E">//   end asc, ...</span>
     * </pre>
     * <p>This function with Union is unsupported!</p>
     * <p>The order values are bound (treated as bind parameter).</p>
     * @param opLambda The callback for option of manual-order containing order values. (NotNull)
     */
    public void withManualOrder(ManualOrderOptionCall opLambda) { // is user public!
        xdoWithManualOrder(cMOO(opLambda));
    }

    // ===================================================================================
    //                                                                    Small Adjustment
    //                                                                    ================
    // ===================================================================================
    //                                                                       Very Internal
    //                                                                       =============
    protected WhiteDelimiterCB newMyCB() {
        return new WhiteDelimiterCB();
    }
    // very internal (for suppressing warn about 'Not Use Import')
    protected String xabUDT() { return Date.class.getName(); }
    protected String xabCQ() { return WhiteDelimiterCQ.class.getName(); }
    protected String xabLSO() { return LikeSearchOption.class.getName(); }
    protected String xabSSQS() { return HpSSQSetupper.class.getName(); }
    protected String xabSCP() { return SubQuery.class.getName(); }
}

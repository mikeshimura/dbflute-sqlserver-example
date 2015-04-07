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
 * The abstract condition-query of vendor_large_data.
 * @author DBFlute(AutoGenerator)
 */
public abstract class AbstractBsVendorLargeDataCQ extends AbstractConditionQuery {

    // ===================================================================================
    //                                                                         Constructor
    //                                                                         ===========
    public AbstractBsVendorLargeDataCQ(ConditionQuery referrerQuery, SqlClause sqlClause, String aliasName, int nestLevel) {
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
        return "vendor_large_data";
    }

    // ===================================================================================
    //                                                                               Query
    //                                                                               =====
    /**
     * Equal(=). And NullIgnored, OnlyOnceRegistered. <br>
     * LARGE_DATA_ID: {PK, NotNull, BIGINT(19)}
     * @param largeDataId The value of largeDataId as equal. (NullAllowed: if null, no condition)
     */
    public void setLargeDataId_Equal(Long largeDataId) {
        doSetLargeDataId_Equal(largeDataId);
    }

    protected void doSetLargeDataId_Equal(Long largeDataId) {
        regLargeDataId(CK_EQ, largeDataId);
    }

    /**
     * NotEqual(&lt;&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * LARGE_DATA_ID: {PK, NotNull, BIGINT(19)}
     * @param largeDataId The value of largeDataId as notEqual. (NullAllowed: if null, no condition)
     */
    public void setLargeDataId_NotEqual(Long largeDataId) {
        doSetLargeDataId_NotEqual(largeDataId);
    }

    protected void doSetLargeDataId_NotEqual(Long largeDataId) {
        regLargeDataId(CK_NES, largeDataId);
    }

    /**
     * GreaterThan(&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * LARGE_DATA_ID: {PK, NotNull, BIGINT(19)}
     * @param largeDataId The value of largeDataId as greaterThan. (NullAllowed: if null, no condition)
     */
    public void setLargeDataId_GreaterThan(Long largeDataId) {
        regLargeDataId(CK_GT, largeDataId);
    }

    /**
     * LessThan(&lt;). And NullIgnored, OnlyOnceRegistered. <br>
     * LARGE_DATA_ID: {PK, NotNull, BIGINT(19)}
     * @param largeDataId The value of largeDataId as lessThan. (NullAllowed: if null, no condition)
     */
    public void setLargeDataId_LessThan(Long largeDataId) {
        regLargeDataId(CK_LT, largeDataId);
    }

    /**
     * GreaterEqual(&gt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * LARGE_DATA_ID: {PK, NotNull, BIGINT(19)}
     * @param largeDataId The value of largeDataId as greaterEqual. (NullAllowed: if null, no condition)
     */
    public void setLargeDataId_GreaterEqual(Long largeDataId) {
        regLargeDataId(CK_GE, largeDataId);
    }

    /**
     * LessEqual(&lt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * LARGE_DATA_ID: {PK, NotNull, BIGINT(19)}
     * @param largeDataId The value of largeDataId as lessEqual. (NullAllowed: if null, no condition)
     */
    public void setLargeDataId_LessEqual(Long largeDataId) {
        regLargeDataId(CK_LE, largeDataId);
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * LARGE_DATA_ID: {PK, NotNull, BIGINT(19)}
     * @param minNumber The min number of largeDataId. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of largeDataId. (NullAllowed: if null, no to-condition)
     * @param opLambda The callback for option of range-of. (NotNull)
     */
    public void setLargeDataId_RangeOf(Long minNumber, Long maxNumber, ConditionOptionCall<RangeOfOption> opLambda) {
        setLargeDataId_RangeOf(minNumber, maxNumber, xcROOP(opLambda));
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * LARGE_DATA_ID: {PK, NotNull, BIGINT(19)}
     * @param minNumber The min number of largeDataId. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of largeDataId. (NullAllowed: if null, no to-condition)
     * @param rangeOfOption The option of range-of. (NotNull)
     */
    protected void setLargeDataId_RangeOf(Long minNumber, Long maxNumber, RangeOfOption rangeOfOption) {
        regROO(minNumber, maxNumber, xgetCValueLargeDataId(), "LARGE_DATA_ID", rangeOfOption);
    }

    /**
     * InScope {in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * LARGE_DATA_ID: {PK, NotNull, BIGINT(19)}
     * @param largeDataIdList The collection of largeDataId as inScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setLargeDataId_InScope(Collection<Long> largeDataIdList) {
        doSetLargeDataId_InScope(largeDataIdList);
    }

    protected void doSetLargeDataId_InScope(Collection<Long> largeDataIdList) {
        regINS(CK_INS, cTL(largeDataIdList), xgetCValueLargeDataId(), "LARGE_DATA_ID");
    }

    /**
     * NotInScope {not in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * LARGE_DATA_ID: {PK, NotNull, BIGINT(19)}
     * @param largeDataIdList The collection of largeDataId as notInScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setLargeDataId_NotInScope(Collection<Long> largeDataIdList) {
        doSetLargeDataId_NotInScope(largeDataIdList);
    }

    protected void doSetLargeDataId_NotInScope(Collection<Long> largeDataIdList) {
        regINS(CK_NINS, cTL(largeDataIdList), xgetCValueLargeDataId(), "LARGE_DATA_ID");
    }

    /**
     * IsNull {is null}. And OnlyOnceRegistered. <br>
     * LARGE_DATA_ID: {PK, NotNull, BIGINT(19)}
     */
    public void setLargeDataId_IsNull() { regLargeDataId(CK_ISN, DOBJ); }

    /**
     * IsNotNull {is not null}. And OnlyOnceRegistered. <br>
     * LARGE_DATA_ID: {PK, NotNull, BIGINT(19)}
     */
    public void setLargeDataId_IsNotNull() { regLargeDataId(CK_ISNN, DOBJ); }

    protected void regLargeDataId(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueLargeDataId(), "LARGE_DATA_ID"); }
    protected abstract ConditionValue xgetCValueLargeDataId();

    /**
     * Equal(=). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_INDEX: {NotNull, VARCHAR(180)}
     * @param stringIndex The value of stringIndex as equal. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringIndex_Equal(String stringIndex) {
        doSetStringIndex_Equal(fRES(stringIndex));
    }

    protected void doSetStringIndex_Equal(String stringIndex) {
        regStringIndex(CK_EQ, stringIndex);
    }

    /**
     * NotEqual(&lt;&gt;). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_INDEX: {NotNull, VARCHAR(180)}
     * @param stringIndex The value of stringIndex as notEqual. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringIndex_NotEqual(String stringIndex) {
        doSetStringIndex_NotEqual(fRES(stringIndex));
    }

    protected void doSetStringIndex_NotEqual(String stringIndex) {
        regStringIndex(CK_NES, stringIndex);
    }

    /**
     * GreaterThan(&gt;). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_INDEX: {NotNull, VARCHAR(180)}
     * @param stringIndex The value of stringIndex as greaterThan. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringIndex_GreaterThan(String stringIndex) {
        regStringIndex(CK_GT, fRES(stringIndex));
    }

    /**
     * LessThan(&lt;). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_INDEX: {NotNull, VARCHAR(180)}
     * @param stringIndex The value of stringIndex as lessThan. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringIndex_LessThan(String stringIndex) {
        regStringIndex(CK_LT, fRES(stringIndex));
    }

    /**
     * GreaterEqual(&gt;=). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_INDEX: {NotNull, VARCHAR(180)}
     * @param stringIndex The value of stringIndex as greaterEqual. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringIndex_GreaterEqual(String stringIndex) {
        regStringIndex(CK_GE, fRES(stringIndex));
    }

    /**
     * LessEqual(&lt;=). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_INDEX: {NotNull, VARCHAR(180)}
     * @param stringIndex The value of stringIndex as lessEqual. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringIndex_LessEqual(String stringIndex) {
        regStringIndex(CK_LE, fRES(stringIndex));
    }

    /**
     * InScope {in ('a', 'b')}. And NullOrEmptyIgnored, NullOrEmptyElementIgnored, SeveralRegistered. <br>
     * STRING_INDEX: {NotNull, VARCHAR(180)}
     * @param stringIndexList The collection of stringIndex as inScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringIndex_InScope(Collection<String> stringIndexList) {
        doSetStringIndex_InScope(stringIndexList);
    }

    protected void doSetStringIndex_InScope(Collection<String> stringIndexList) {
        regINS(CK_INS, cTL(stringIndexList), xgetCValueStringIndex(), "STRING_INDEX");
    }

    /**
     * NotInScope {not in ('a', 'b')}. And NullOrEmptyIgnored, NullOrEmptyElementIgnored, SeveralRegistered. <br>
     * STRING_INDEX: {NotNull, VARCHAR(180)}
     * @param stringIndexList The collection of stringIndex as notInScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringIndex_NotInScope(Collection<String> stringIndexList) {
        doSetStringIndex_NotInScope(stringIndexList);
    }

    protected void doSetStringIndex_NotInScope(Collection<String> stringIndexList) {
        regINS(CK_NINS, cTL(stringIndexList), xgetCValueStringIndex(), "STRING_INDEX");
    }

    /**
     * LikeSearch with various options. (versatile) {like '%xxx%' escape ...}. And NullOrEmptyIgnored, SeveralRegistered. <br>
     * STRING_INDEX: {NotNull, VARCHAR(180)} <br>
     * <pre>e.g. setStringIndex_LikeSearch("xxx", op <span style="color: #90226C; font-weight: bold"><span style="font-size: 120%">-</span>&gt;</span> op.<span style="color: #CC4747">likeContain()</span>);</pre>
     * @param stringIndex The value of stringIndex as likeSearch. (NullAllowed: if null (or empty), no condition)
     * @param opLambda The callback for option of like-search. (NotNull)
     */
    public void setStringIndex_LikeSearch(String stringIndex, ConditionOptionCall<LikeSearchOption> opLambda) {
        setStringIndex_LikeSearch(stringIndex, xcLSOP(opLambda));
    }

    /**
     * LikeSearch with various options. (versatile) {like '%xxx%' escape ...}. And NullOrEmptyIgnored, SeveralRegistered. <br>
     * STRING_INDEX: {NotNull, VARCHAR(180)} <br>
     * <pre>e.g. setStringIndex_LikeSearch("xxx", new <span style="color: #CC4747">LikeSearchOption</span>().likeContain());</pre>
     * @param stringIndex The value of stringIndex as likeSearch. (NullAllowed: if null (or empty), no condition)
     * @param likeSearchOption The option of like-search. (NotNull)
     */
    protected void setStringIndex_LikeSearch(String stringIndex, LikeSearchOption likeSearchOption) {
        regLSQ(CK_LS, fRES(stringIndex), xgetCValueStringIndex(), "STRING_INDEX", likeSearchOption);
    }

    /**
     * NotLikeSearch with various options. (versatile) {not like 'xxx%' escape ...} <br>
     * And NullOrEmptyIgnored, SeveralRegistered. <br>
     * STRING_INDEX: {NotNull, VARCHAR(180)}
     * @param stringIndex The value of stringIndex as notLikeSearch. (NullAllowed: if null (or empty), no condition)
     * @param opLambda The callback for option of like-search. (NotNull)
     */
    public void setStringIndex_NotLikeSearch(String stringIndex, ConditionOptionCall<LikeSearchOption> opLambda) {
        setStringIndex_NotLikeSearch(stringIndex, xcLSOP(opLambda));
    }

    /**
     * NotLikeSearch with various options. (versatile) {not like 'xxx%' escape ...} <br>
     * And NullOrEmptyIgnored, SeveralRegistered. <br>
     * STRING_INDEX: {NotNull, VARCHAR(180)}
     * @param stringIndex The value of stringIndex as notLikeSearch. (NullAllowed: if null (or empty), no condition)
     * @param likeSearchOption The option of not-like-search. (NotNull)
     */
    protected void setStringIndex_NotLikeSearch(String stringIndex, LikeSearchOption likeSearchOption) {
        regLSQ(CK_NLS, fRES(stringIndex), xgetCValueStringIndex(), "STRING_INDEX", likeSearchOption);
    }

    protected void regStringIndex(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueStringIndex(), "STRING_INDEX"); }
    protected abstract ConditionValue xgetCValueStringIndex();

    /**
     * Equal(=). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_NO_INDEX: {NotNull, VARCHAR(180)}
     * @param stringNoIndex The value of stringNoIndex as equal. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringNoIndex_Equal(String stringNoIndex) {
        doSetStringNoIndex_Equal(fRES(stringNoIndex));
    }

    protected void doSetStringNoIndex_Equal(String stringNoIndex) {
        regStringNoIndex(CK_EQ, stringNoIndex);
    }

    /**
     * NotEqual(&lt;&gt;). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_NO_INDEX: {NotNull, VARCHAR(180)}
     * @param stringNoIndex The value of stringNoIndex as notEqual. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringNoIndex_NotEqual(String stringNoIndex) {
        doSetStringNoIndex_NotEqual(fRES(stringNoIndex));
    }

    protected void doSetStringNoIndex_NotEqual(String stringNoIndex) {
        regStringNoIndex(CK_NES, stringNoIndex);
    }

    /**
     * GreaterThan(&gt;). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_NO_INDEX: {NotNull, VARCHAR(180)}
     * @param stringNoIndex The value of stringNoIndex as greaterThan. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringNoIndex_GreaterThan(String stringNoIndex) {
        regStringNoIndex(CK_GT, fRES(stringNoIndex));
    }

    /**
     * LessThan(&lt;). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_NO_INDEX: {NotNull, VARCHAR(180)}
     * @param stringNoIndex The value of stringNoIndex as lessThan. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringNoIndex_LessThan(String stringNoIndex) {
        regStringNoIndex(CK_LT, fRES(stringNoIndex));
    }

    /**
     * GreaterEqual(&gt;=). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_NO_INDEX: {NotNull, VARCHAR(180)}
     * @param stringNoIndex The value of stringNoIndex as greaterEqual. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringNoIndex_GreaterEqual(String stringNoIndex) {
        regStringNoIndex(CK_GE, fRES(stringNoIndex));
    }

    /**
     * LessEqual(&lt;=). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_NO_INDEX: {NotNull, VARCHAR(180)}
     * @param stringNoIndex The value of stringNoIndex as lessEqual. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringNoIndex_LessEqual(String stringNoIndex) {
        regStringNoIndex(CK_LE, fRES(stringNoIndex));
    }

    /**
     * InScope {in ('a', 'b')}. And NullOrEmptyIgnored, NullOrEmptyElementIgnored, SeveralRegistered. <br>
     * STRING_NO_INDEX: {NotNull, VARCHAR(180)}
     * @param stringNoIndexList The collection of stringNoIndex as inScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringNoIndex_InScope(Collection<String> stringNoIndexList) {
        doSetStringNoIndex_InScope(stringNoIndexList);
    }

    protected void doSetStringNoIndex_InScope(Collection<String> stringNoIndexList) {
        regINS(CK_INS, cTL(stringNoIndexList), xgetCValueStringNoIndex(), "STRING_NO_INDEX");
    }

    /**
     * NotInScope {not in ('a', 'b')}. And NullOrEmptyIgnored, NullOrEmptyElementIgnored, SeveralRegistered. <br>
     * STRING_NO_INDEX: {NotNull, VARCHAR(180)}
     * @param stringNoIndexList The collection of stringNoIndex as notInScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringNoIndex_NotInScope(Collection<String> stringNoIndexList) {
        doSetStringNoIndex_NotInScope(stringNoIndexList);
    }

    protected void doSetStringNoIndex_NotInScope(Collection<String> stringNoIndexList) {
        regINS(CK_NINS, cTL(stringNoIndexList), xgetCValueStringNoIndex(), "STRING_NO_INDEX");
    }

    /**
     * LikeSearch with various options. (versatile) {like '%xxx%' escape ...}. And NullOrEmptyIgnored, SeveralRegistered. <br>
     * STRING_NO_INDEX: {NotNull, VARCHAR(180)} <br>
     * <pre>e.g. setStringNoIndex_LikeSearch("xxx", op <span style="color: #90226C; font-weight: bold"><span style="font-size: 120%">-</span>&gt;</span> op.<span style="color: #CC4747">likeContain()</span>);</pre>
     * @param stringNoIndex The value of stringNoIndex as likeSearch. (NullAllowed: if null (or empty), no condition)
     * @param opLambda The callback for option of like-search. (NotNull)
     */
    public void setStringNoIndex_LikeSearch(String stringNoIndex, ConditionOptionCall<LikeSearchOption> opLambda) {
        setStringNoIndex_LikeSearch(stringNoIndex, xcLSOP(opLambda));
    }

    /**
     * LikeSearch with various options. (versatile) {like '%xxx%' escape ...}. And NullOrEmptyIgnored, SeveralRegistered. <br>
     * STRING_NO_INDEX: {NotNull, VARCHAR(180)} <br>
     * <pre>e.g. setStringNoIndex_LikeSearch("xxx", new <span style="color: #CC4747">LikeSearchOption</span>().likeContain());</pre>
     * @param stringNoIndex The value of stringNoIndex as likeSearch. (NullAllowed: if null (or empty), no condition)
     * @param likeSearchOption The option of like-search. (NotNull)
     */
    protected void setStringNoIndex_LikeSearch(String stringNoIndex, LikeSearchOption likeSearchOption) {
        regLSQ(CK_LS, fRES(stringNoIndex), xgetCValueStringNoIndex(), "STRING_NO_INDEX", likeSearchOption);
    }

    /**
     * NotLikeSearch with various options. (versatile) {not like 'xxx%' escape ...} <br>
     * And NullOrEmptyIgnored, SeveralRegistered. <br>
     * STRING_NO_INDEX: {NotNull, VARCHAR(180)}
     * @param stringNoIndex The value of stringNoIndex as notLikeSearch. (NullAllowed: if null (or empty), no condition)
     * @param opLambda The callback for option of like-search. (NotNull)
     */
    public void setStringNoIndex_NotLikeSearch(String stringNoIndex, ConditionOptionCall<LikeSearchOption> opLambda) {
        setStringNoIndex_NotLikeSearch(stringNoIndex, xcLSOP(opLambda));
    }

    /**
     * NotLikeSearch with various options. (versatile) {not like 'xxx%' escape ...} <br>
     * And NullOrEmptyIgnored, SeveralRegistered. <br>
     * STRING_NO_INDEX: {NotNull, VARCHAR(180)}
     * @param stringNoIndex The value of stringNoIndex as notLikeSearch. (NullAllowed: if null (or empty), no condition)
     * @param likeSearchOption The option of not-like-search. (NotNull)
     */
    protected void setStringNoIndex_NotLikeSearch(String stringNoIndex, LikeSearchOption likeSearchOption) {
        regLSQ(CK_NLS, fRES(stringNoIndex), xgetCValueStringNoIndex(), "STRING_NO_INDEX", likeSearchOption);
    }

    protected void regStringNoIndex(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueStringNoIndex(), "STRING_NO_INDEX"); }
    protected abstract ConditionValue xgetCValueStringNoIndex();

    /**
     * Equal(=). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_UNIQUE_INDEX: {UQ, NotNull, VARCHAR(180)}
     * @param stringUniqueIndex The value of stringUniqueIndex as equal. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringUniqueIndex_Equal(String stringUniqueIndex) {
        doSetStringUniqueIndex_Equal(fRES(stringUniqueIndex));
    }

    protected void doSetStringUniqueIndex_Equal(String stringUniqueIndex) {
        regStringUniqueIndex(CK_EQ, stringUniqueIndex);
    }

    /**
     * NotEqual(&lt;&gt;). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_UNIQUE_INDEX: {UQ, NotNull, VARCHAR(180)}
     * @param stringUniqueIndex The value of stringUniqueIndex as notEqual. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringUniqueIndex_NotEqual(String stringUniqueIndex) {
        doSetStringUniqueIndex_NotEqual(fRES(stringUniqueIndex));
    }

    protected void doSetStringUniqueIndex_NotEqual(String stringUniqueIndex) {
        regStringUniqueIndex(CK_NES, stringUniqueIndex);
    }

    /**
     * GreaterThan(&gt;). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_UNIQUE_INDEX: {UQ, NotNull, VARCHAR(180)}
     * @param stringUniqueIndex The value of stringUniqueIndex as greaterThan. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringUniqueIndex_GreaterThan(String stringUniqueIndex) {
        regStringUniqueIndex(CK_GT, fRES(stringUniqueIndex));
    }

    /**
     * LessThan(&lt;). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_UNIQUE_INDEX: {UQ, NotNull, VARCHAR(180)}
     * @param stringUniqueIndex The value of stringUniqueIndex as lessThan. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringUniqueIndex_LessThan(String stringUniqueIndex) {
        regStringUniqueIndex(CK_LT, fRES(stringUniqueIndex));
    }

    /**
     * GreaterEqual(&gt;=). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_UNIQUE_INDEX: {UQ, NotNull, VARCHAR(180)}
     * @param stringUniqueIndex The value of stringUniqueIndex as greaterEqual. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringUniqueIndex_GreaterEqual(String stringUniqueIndex) {
        regStringUniqueIndex(CK_GE, fRES(stringUniqueIndex));
    }

    /**
     * LessEqual(&lt;=). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * STRING_UNIQUE_INDEX: {UQ, NotNull, VARCHAR(180)}
     * @param stringUniqueIndex The value of stringUniqueIndex as lessEqual. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringUniqueIndex_LessEqual(String stringUniqueIndex) {
        regStringUniqueIndex(CK_LE, fRES(stringUniqueIndex));
    }

    /**
     * InScope {in ('a', 'b')}. And NullOrEmptyIgnored, NullOrEmptyElementIgnored, SeveralRegistered. <br>
     * STRING_UNIQUE_INDEX: {UQ, NotNull, VARCHAR(180)}
     * @param stringUniqueIndexList The collection of stringUniqueIndex as inScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringUniqueIndex_InScope(Collection<String> stringUniqueIndexList) {
        doSetStringUniqueIndex_InScope(stringUniqueIndexList);
    }

    protected void doSetStringUniqueIndex_InScope(Collection<String> stringUniqueIndexList) {
        regINS(CK_INS, cTL(stringUniqueIndexList), xgetCValueStringUniqueIndex(), "STRING_UNIQUE_INDEX");
    }

    /**
     * NotInScope {not in ('a', 'b')}. And NullOrEmptyIgnored, NullOrEmptyElementIgnored, SeveralRegistered. <br>
     * STRING_UNIQUE_INDEX: {UQ, NotNull, VARCHAR(180)}
     * @param stringUniqueIndexList The collection of stringUniqueIndex as notInScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setStringUniqueIndex_NotInScope(Collection<String> stringUniqueIndexList) {
        doSetStringUniqueIndex_NotInScope(stringUniqueIndexList);
    }

    protected void doSetStringUniqueIndex_NotInScope(Collection<String> stringUniqueIndexList) {
        regINS(CK_NINS, cTL(stringUniqueIndexList), xgetCValueStringUniqueIndex(), "STRING_UNIQUE_INDEX");
    }

    /**
     * LikeSearch with various options. (versatile) {like '%xxx%' escape ...}. And NullOrEmptyIgnored, SeveralRegistered. <br>
     * STRING_UNIQUE_INDEX: {UQ, NotNull, VARCHAR(180)} <br>
     * <pre>e.g. setStringUniqueIndex_LikeSearch("xxx", op <span style="color: #90226C; font-weight: bold"><span style="font-size: 120%">-</span>&gt;</span> op.<span style="color: #CC4747">likeContain()</span>);</pre>
     * @param stringUniqueIndex The value of stringUniqueIndex as likeSearch. (NullAllowed: if null (or empty), no condition)
     * @param opLambda The callback for option of like-search. (NotNull)
     */
    public void setStringUniqueIndex_LikeSearch(String stringUniqueIndex, ConditionOptionCall<LikeSearchOption> opLambda) {
        setStringUniqueIndex_LikeSearch(stringUniqueIndex, xcLSOP(opLambda));
    }

    /**
     * LikeSearch with various options. (versatile) {like '%xxx%' escape ...}. And NullOrEmptyIgnored, SeveralRegistered. <br>
     * STRING_UNIQUE_INDEX: {UQ, NotNull, VARCHAR(180)} <br>
     * <pre>e.g. setStringUniqueIndex_LikeSearch("xxx", new <span style="color: #CC4747">LikeSearchOption</span>().likeContain());</pre>
     * @param stringUniqueIndex The value of stringUniqueIndex as likeSearch. (NullAllowed: if null (or empty), no condition)
     * @param likeSearchOption The option of like-search. (NotNull)
     */
    protected void setStringUniqueIndex_LikeSearch(String stringUniqueIndex, LikeSearchOption likeSearchOption) {
        regLSQ(CK_LS, fRES(stringUniqueIndex), xgetCValueStringUniqueIndex(), "STRING_UNIQUE_INDEX", likeSearchOption);
    }

    /**
     * NotLikeSearch with various options. (versatile) {not like 'xxx%' escape ...} <br>
     * And NullOrEmptyIgnored, SeveralRegistered. <br>
     * STRING_UNIQUE_INDEX: {UQ, NotNull, VARCHAR(180)}
     * @param stringUniqueIndex The value of stringUniqueIndex as notLikeSearch. (NullAllowed: if null (or empty), no condition)
     * @param opLambda The callback for option of like-search. (NotNull)
     */
    public void setStringUniqueIndex_NotLikeSearch(String stringUniqueIndex, ConditionOptionCall<LikeSearchOption> opLambda) {
        setStringUniqueIndex_NotLikeSearch(stringUniqueIndex, xcLSOP(opLambda));
    }

    /**
     * NotLikeSearch with various options. (versatile) {not like 'xxx%' escape ...} <br>
     * And NullOrEmptyIgnored, SeveralRegistered. <br>
     * STRING_UNIQUE_INDEX: {UQ, NotNull, VARCHAR(180)}
     * @param stringUniqueIndex The value of stringUniqueIndex as notLikeSearch. (NullAllowed: if null (or empty), no condition)
     * @param likeSearchOption The option of not-like-search. (NotNull)
     */
    protected void setStringUniqueIndex_NotLikeSearch(String stringUniqueIndex, LikeSearchOption likeSearchOption) {
        regLSQ(CK_NLS, fRES(stringUniqueIndex), xgetCValueStringUniqueIndex(), "STRING_UNIQUE_INDEX", likeSearchOption);
    }

    protected void regStringUniqueIndex(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueStringUniqueIndex(), "STRING_UNIQUE_INDEX"); }
    protected abstract ConditionValue xgetCValueStringUniqueIndex();

    /**
     * Equal(=). And NullIgnored, OnlyOnceRegistered. <br>
     * INTFLG_INDEX: {NotNull, INT(10)}
     * @param intflgIndex The value of intflgIndex as equal. (NullAllowed: if null, no condition)
     */
    public void setIntflgIndex_Equal(Integer intflgIndex) {
        doSetIntflgIndex_Equal(intflgIndex);
    }

    protected void doSetIntflgIndex_Equal(Integer intflgIndex) {
        regIntflgIndex(CK_EQ, intflgIndex);
    }

    /**
     * NotEqual(&lt;&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * INTFLG_INDEX: {NotNull, INT(10)}
     * @param intflgIndex The value of intflgIndex as notEqual. (NullAllowed: if null, no condition)
     */
    public void setIntflgIndex_NotEqual(Integer intflgIndex) {
        doSetIntflgIndex_NotEqual(intflgIndex);
    }

    protected void doSetIntflgIndex_NotEqual(Integer intflgIndex) {
        regIntflgIndex(CK_NES, intflgIndex);
    }

    /**
     * GreaterThan(&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * INTFLG_INDEX: {NotNull, INT(10)}
     * @param intflgIndex The value of intflgIndex as greaterThan. (NullAllowed: if null, no condition)
     */
    public void setIntflgIndex_GreaterThan(Integer intflgIndex) {
        regIntflgIndex(CK_GT, intflgIndex);
    }

    /**
     * LessThan(&lt;). And NullIgnored, OnlyOnceRegistered. <br>
     * INTFLG_INDEX: {NotNull, INT(10)}
     * @param intflgIndex The value of intflgIndex as lessThan. (NullAllowed: if null, no condition)
     */
    public void setIntflgIndex_LessThan(Integer intflgIndex) {
        regIntflgIndex(CK_LT, intflgIndex);
    }

    /**
     * GreaterEqual(&gt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * INTFLG_INDEX: {NotNull, INT(10)}
     * @param intflgIndex The value of intflgIndex as greaterEqual. (NullAllowed: if null, no condition)
     */
    public void setIntflgIndex_GreaterEqual(Integer intflgIndex) {
        regIntflgIndex(CK_GE, intflgIndex);
    }

    /**
     * LessEqual(&lt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * INTFLG_INDEX: {NotNull, INT(10)}
     * @param intflgIndex The value of intflgIndex as lessEqual. (NullAllowed: if null, no condition)
     */
    public void setIntflgIndex_LessEqual(Integer intflgIndex) {
        regIntflgIndex(CK_LE, intflgIndex);
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * INTFLG_INDEX: {NotNull, INT(10)}
     * @param minNumber The min number of intflgIndex. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of intflgIndex. (NullAllowed: if null, no to-condition)
     * @param opLambda The callback for option of range-of. (NotNull)
     */
    public void setIntflgIndex_RangeOf(Integer minNumber, Integer maxNumber, ConditionOptionCall<RangeOfOption> opLambda) {
        setIntflgIndex_RangeOf(minNumber, maxNumber, xcROOP(opLambda));
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * INTFLG_INDEX: {NotNull, INT(10)}
     * @param minNumber The min number of intflgIndex. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of intflgIndex. (NullAllowed: if null, no to-condition)
     * @param rangeOfOption The option of range-of. (NotNull)
     */
    protected void setIntflgIndex_RangeOf(Integer minNumber, Integer maxNumber, RangeOfOption rangeOfOption) {
        regROO(minNumber, maxNumber, xgetCValueIntflgIndex(), "INTFLG_INDEX", rangeOfOption);
    }

    /**
     * InScope {in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * INTFLG_INDEX: {NotNull, INT(10)}
     * @param intflgIndexList The collection of intflgIndex as inScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setIntflgIndex_InScope(Collection<Integer> intflgIndexList) {
        doSetIntflgIndex_InScope(intflgIndexList);
    }

    protected void doSetIntflgIndex_InScope(Collection<Integer> intflgIndexList) {
        regINS(CK_INS, cTL(intflgIndexList), xgetCValueIntflgIndex(), "INTFLG_INDEX");
    }

    /**
     * NotInScope {not in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * INTFLG_INDEX: {NotNull, INT(10)}
     * @param intflgIndexList The collection of intflgIndex as notInScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setIntflgIndex_NotInScope(Collection<Integer> intflgIndexList) {
        doSetIntflgIndex_NotInScope(intflgIndexList);
    }

    protected void doSetIntflgIndex_NotInScope(Collection<Integer> intflgIndexList) {
        regINS(CK_NINS, cTL(intflgIndexList), xgetCValueIntflgIndex(), "INTFLG_INDEX");
    }

    protected void regIntflgIndex(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueIntflgIndex(), "INTFLG_INDEX"); }
    protected abstract ConditionValue xgetCValueIntflgIndex();

    /**
     * Equal(=). And NullIgnored, OnlyOnceRegistered. <br>
     * NUMERIC_INTEGER_INDEX: {NotNull, DECIMAL(8)}
     * @param numericIntegerIndex The value of numericIntegerIndex as equal. (NullAllowed: if null, no condition)
     */
    public void setNumericIntegerIndex_Equal(Integer numericIntegerIndex) {
        doSetNumericIntegerIndex_Equal(numericIntegerIndex);
    }

    protected void doSetNumericIntegerIndex_Equal(Integer numericIntegerIndex) {
        regNumericIntegerIndex(CK_EQ, numericIntegerIndex);
    }

    /**
     * NotEqual(&lt;&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * NUMERIC_INTEGER_INDEX: {NotNull, DECIMAL(8)}
     * @param numericIntegerIndex The value of numericIntegerIndex as notEqual. (NullAllowed: if null, no condition)
     */
    public void setNumericIntegerIndex_NotEqual(Integer numericIntegerIndex) {
        doSetNumericIntegerIndex_NotEqual(numericIntegerIndex);
    }

    protected void doSetNumericIntegerIndex_NotEqual(Integer numericIntegerIndex) {
        regNumericIntegerIndex(CK_NES, numericIntegerIndex);
    }

    /**
     * GreaterThan(&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * NUMERIC_INTEGER_INDEX: {NotNull, DECIMAL(8)}
     * @param numericIntegerIndex The value of numericIntegerIndex as greaterThan. (NullAllowed: if null, no condition)
     */
    public void setNumericIntegerIndex_GreaterThan(Integer numericIntegerIndex) {
        regNumericIntegerIndex(CK_GT, numericIntegerIndex);
    }

    /**
     * LessThan(&lt;). And NullIgnored, OnlyOnceRegistered. <br>
     * NUMERIC_INTEGER_INDEX: {NotNull, DECIMAL(8)}
     * @param numericIntegerIndex The value of numericIntegerIndex as lessThan. (NullAllowed: if null, no condition)
     */
    public void setNumericIntegerIndex_LessThan(Integer numericIntegerIndex) {
        regNumericIntegerIndex(CK_LT, numericIntegerIndex);
    }

    /**
     * GreaterEqual(&gt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * NUMERIC_INTEGER_INDEX: {NotNull, DECIMAL(8)}
     * @param numericIntegerIndex The value of numericIntegerIndex as greaterEqual. (NullAllowed: if null, no condition)
     */
    public void setNumericIntegerIndex_GreaterEqual(Integer numericIntegerIndex) {
        regNumericIntegerIndex(CK_GE, numericIntegerIndex);
    }

    /**
     * LessEqual(&lt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * NUMERIC_INTEGER_INDEX: {NotNull, DECIMAL(8)}
     * @param numericIntegerIndex The value of numericIntegerIndex as lessEqual. (NullAllowed: if null, no condition)
     */
    public void setNumericIntegerIndex_LessEqual(Integer numericIntegerIndex) {
        regNumericIntegerIndex(CK_LE, numericIntegerIndex);
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * NUMERIC_INTEGER_INDEX: {NotNull, DECIMAL(8)}
     * @param minNumber The min number of numericIntegerIndex. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of numericIntegerIndex. (NullAllowed: if null, no to-condition)
     * @param opLambda The callback for option of range-of. (NotNull)
     */
    public void setNumericIntegerIndex_RangeOf(Integer minNumber, Integer maxNumber, ConditionOptionCall<RangeOfOption> opLambda) {
        setNumericIntegerIndex_RangeOf(minNumber, maxNumber, xcROOP(opLambda));
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * NUMERIC_INTEGER_INDEX: {NotNull, DECIMAL(8)}
     * @param minNumber The min number of numericIntegerIndex. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of numericIntegerIndex. (NullAllowed: if null, no to-condition)
     * @param rangeOfOption The option of range-of. (NotNull)
     */
    protected void setNumericIntegerIndex_RangeOf(Integer minNumber, Integer maxNumber, RangeOfOption rangeOfOption) {
        regROO(minNumber, maxNumber, xgetCValueNumericIntegerIndex(), "NUMERIC_INTEGER_INDEX", rangeOfOption);
    }

    /**
     * InScope {in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * NUMERIC_INTEGER_INDEX: {NotNull, DECIMAL(8)}
     * @param numericIntegerIndexList The collection of numericIntegerIndex as inScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setNumericIntegerIndex_InScope(Collection<Integer> numericIntegerIndexList) {
        doSetNumericIntegerIndex_InScope(numericIntegerIndexList);
    }

    protected void doSetNumericIntegerIndex_InScope(Collection<Integer> numericIntegerIndexList) {
        regINS(CK_INS, cTL(numericIntegerIndexList), xgetCValueNumericIntegerIndex(), "NUMERIC_INTEGER_INDEX");
    }

    /**
     * NotInScope {not in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * NUMERIC_INTEGER_INDEX: {NotNull, DECIMAL(8)}
     * @param numericIntegerIndexList The collection of numericIntegerIndex as notInScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setNumericIntegerIndex_NotInScope(Collection<Integer> numericIntegerIndexList) {
        doSetNumericIntegerIndex_NotInScope(numericIntegerIndexList);
    }

    protected void doSetNumericIntegerIndex_NotInScope(Collection<Integer> numericIntegerIndexList) {
        regINS(CK_NINS, cTL(numericIntegerIndexList), xgetCValueNumericIntegerIndex(), "NUMERIC_INTEGER_INDEX");
    }

    protected void regNumericIntegerIndex(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueNumericIntegerIndex(), "NUMERIC_INTEGER_INDEX"); }
    protected abstract ConditionValue xgetCValueNumericIntegerIndex();

    /**
     * Equal(=). And NullIgnored, OnlyOnceRegistered. <br>
     * NUMERIC_INTEGER_NO_INDEX: {NotNull, DECIMAL(8)}
     * @param numericIntegerNoIndex The value of numericIntegerNoIndex as equal. (NullAllowed: if null, no condition)
     */
    public void setNumericIntegerNoIndex_Equal(Integer numericIntegerNoIndex) {
        doSetNumericIntegerNoIndex_Equal(numericIntegerNoIndex);
    }

    protected void doSetNumericIntegerNoIndex_Equal(Integer numericIntegerNoIndex) {
        regNumericIntegerNoIndex(CK_EQ, numericIntegerNoIndex);
    }

    /**
     * NotEqual(&lt;&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * NUMERIC_INTEGER_NO_INDEX: {NotNull, DECIMAL(8)}
     * @param numericIntegerNoIndex The value of numericIntegerNoIndex as notEqual. (NullAllowed: if null, no condition)
     */
    public void setNumericIntegerNoIndex_NotEqual(Integer numericIntegerNoIndex) {
        doSetNumericIntegerNoIndex_NotEqual(numericIntegerNoIndex);
    }

    protected void doSetNumericIntegerNoIndex_NotEqual(Integer numericIntegerNoIndex) {
        regNumericIntegerNoIndex(CK_NES, numericIntegerNoIndex);
    }

    /**
     * GreaterThan(&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * NUMERIC_INTEGER_NO_INDEX: {NotNull, DECIMAL(8)}
     * @param numericIntegerNoIndex The value of numericIntegerNoIndex as greaterThan. (NullAllowed: if null, no condition)
     */
    public void setNumericIntegerNoIndex_GreaterThan(Integer numericIntegerNoIndex) {
        regNumericIntegerNoIndex(CK_GT, numericIntegerNoIndex);
    }

    /**
     * LessThan(&lt;). And NullIgnored, OnlyOnceRegistered. <br>
     * NUMERIC_INTEGER_NO_INDEX: {NotNull, DECIMAL(8)}
     * @param numericIntegerNoIndex The value of numericIntegerNoIndex as lessThan. (NullAllowed: if null, no condition)
     */
    public void setNumericIntegerNoIndex_LessThan(Integer numericIntegerNoIndex) {
        regNumericIntegerNoIndex(CK_LT, numericIntegerNoIndex);
    }

    /**
     * GreaterEqual(&gt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * NUMERIC_INTEGER_NO_INDEX: {NotNull, DECIMAL(8)}
     * @param numericIntegerNoIndex The value of numericIntegerNoIndex as greaterEqual. (NullAllowed: if null, no condition)
     */
    public void setNumericIntegerNoIndex_GreaterEqual(Integer numericIntegerNoIndex) {
        regNumericIntegerNoIndex(CK_GE, numericIntegerNoIndex);
    }

    /**
     * LessEqual(&lt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * NUMERIC_INTEGER_NO_INDEX: {NotNull, DECIMAL(8)}
     * @param numericIntegerNoIndex The value of numericIntegerNoIndex as lessEqual. (NullAllowed: if null, no condition)
     */
    public void setNumericIntegerNoIndex_LessEqual(Integer numericIntegerNoIndex) {
        regNumericIntegerNoIndex(CK_LE, numericIntegerNoIndex);
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * NUMERIC_INTEGER_NO_INDEX: {NotNull, DECIMAL(8)}
     * @param minNumber The min number of numericIntegerNoIndex. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of numericIntegerNoIndex. (NullAllowed: if null, no to-condition)
     * @param opLambda The callback for option of range-of. (NotNull)
     */
    public void setNumericIntegerNoIndex_RangeOf(Integer minNumber, Integer maxNumber, ConditionOptionCall<RangeOfOption> opLambda) {
        setNumericIntegerNoIndex_RangeOf(minNumber, maxNumber, xcROOP(opLambda));
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * NUMERIC_INTEGER_NO_INDEX: {NotNull, DECIMAL(8)}
     * @param minNumber The min number of numericIntegerNoIndex. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of numericIntegerNoIndex. (NullAllowed: if null, no to-condition)
     * @param rangeOfOption The option of range-of. (NotNull)
     */
    protected void setNumericIntegerNoIndex_RangeOf(Integer minNumber, Integer maxNumber, RangeOfOption rangeOfOption) {
        regROO(minNumber, maxNumber, xgetCValueNumericIntegerNoIndex(), "NUMERIC_INTEGER_NO_INDEX", rangeOfOption);
    }

    /**
     * InScope {in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * NUMERIC_INTEGER_NO_INDEX: {NotNull, DECIMAL(8)}
     * @param numericIntegerNoIndexList The collection of numericIntegerNoIndex as inScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setNumericIntegerNoIndex_InScope(Collection<Integer> numericIntegerNoIndexList) {
        doSetNumericIntegerNoIndex_InScope(numericIntegerNoIndexList);
    }

    protected void doSetNumericIntegerNoIndex_InScope(Collection<Integer> numericIntegerNoIndexList) {
        regINS(CK_INS, cTL(numericIntegerNoIndexList), xgetCValueNumericIntegerNoIndex(), "NUMERIC_INTEGER_NO_INDEX");
    }

    /**
     * NotInScope {not in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * NUMERIC_INTEGER_NO_INDEX: {NotNull, DECIMAL(8)}
     * @param numericIntegerNoIndexList The collection of numericIntegerNoIndex as notInScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setNumericIntegerNoIndex_NotInScope(Collection<Integer> numericIntegerNoIndexList) {
        doSetNumericIntegerNoIndex_NotInScope(numericIntegerNoIndexList);
    }

    protected void doSetNumericIntegerNoIndex_NotInScope(Collection<Integer> numericIntegerNoIndexList) {
        regINS(CK_NINS, cTL(numericIntegerNoIndexList), xgetCValueNumericIntegerNoIndex(), "NUMERIC_INTEGER_NO_INDEX");
    }

    protected void regNumericIntegerNoIndex(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueNumericIntegerNoIndex(), "NUMERIC_INTEGER_NO_INDEX"); }
    protected abstract ConditionValue xgetCValueNumericIntegerNoIndex();

    // ===================================================================================
    //                                                                     ScalarCondition
    //                                                                     ===============
    /**
     * Prepare ScalarCondition as equal. <br>
     * {where FOO = (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_Equal()</span>.max(new SubQuery&lt;VendorLargeDataCB&gt;() {
     *     public void query(VendorLargeDataCB subCB) {
     *         subCB.specify().setXxx... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setYyy...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorLargeDataCB> scalar_Equal() {
        return xcreateSSQFunction(CK_EQ, VendorLargeDataCB.class);
    }

    /**
     * Prepare ScalarCondition as equal. <br>
     * {where FOO &lt;&gt; (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_NotEqual()</span>.max(new SubQuery&lt;VendorLargeDataCB&gt;() {
     *     public void query(VendorLargeDataCB subCB) {
     *         subCB.specify().setXxx... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setYyy...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorLargeDataCB> scalar_NotEqual() {
        return xcreateSSQFunction(CK_NES, VendorLargeDataCB.class);
    }

    /**
     * Prepare ScalarCondition as greaterThan. <br>
     * {where FOO &gt; (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_GreaterThan()</span>.max(new SubQuery&lt;VendorLargeDataCB&gt;() {
     *     public void query(VendorLargeDataCB subCB) {
     *         subCB.specify().setFoo... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setBar...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorLargeDataCB> scalar_GreaterThan() {
        return xcreateSSQFunction(CK_GT, VendorLargeDataCB.class);
    }

    /**
     * Prepare ScalarCondition as lessThan. <br>
     * {where FOO &lt; (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_LessThan()</span>.max(new SubQuery&lt;VendorLargeDataCB&gt;() {
     *     public void query(VendorLargeDataCB subCB) {
     *         subCB.specify().setFoo... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setBar...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorLargeDataCB> scalar_LessThan() {
        return xcreateSSQFunction(CK_LT, VendorLargeDataCB.class);
    }

    /**
     * Prepare ScalarCondition as greaterEqual. <br>
     * {where FOO &gt;= (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_GreaterEqual()</span>.max(new SubQuery&lt;VendorLargeDataCB&gt;() {
     *     public void query(VendorLargeDataCB subCB) {
     *         subCB.specify().setFoo... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setBar...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorLargeDataCB> scalar_GreaterEqual() {
        return xcreateSSQFunction(CK_GE, VendorLargeDataCB.class);
    }

    /**
     * Prepare ScalarCondition as lessEqual. <br>
     * {where FOO &lt;= (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_LessEqual()</span>.max(new SubQuery&lt;VendorLargeDataCB&gt;() {
     *     public void query(VendorLargeDataCB subCB) {
     *         subCB.specify().setFoo... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setBar...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorLargeDataCB> scalar_LessEqual() {
        return xcreateSSQFunction(CK_LE, VendorLargeDataCB.class);
    }

    @SuppressWarnings("unchecked")
    protected <CB extends ConditionBean> void xscalarCondition(String fn, SubQuery<CB> sq, String rd, HpSSQOption<CB> op) {
        assertObjectNotNull("subQuery", sq);
        VendorLargeDataCB cb = xcreateScalarConditionCB(); sq.query((CB)cb);
        String pp = keepScalarCondition(cb.query()); // for saving query-value
        op.setPartitionByCBean((CB)xcreateScalarConditionPartitionByCB()); // for using partition-by
        registerScalarCondition(fn, cb.query(), pp, rd, op);
    }
    public abstract String keepScalarCondition(VendorLargeDataCQ sq);

    protected VendorLargeDataCB xcreateScalarConditionCB() {
        VendorLargeDataCB cb = newMyCB(); cb.xsetupForScalarCondition(this); return cb;
    }

    protected VendorLargeDataCB xcreateScalarConditionPartitionByCB() {
        VendorLargeDataCB cb = newMyCB(); cb.xsetupForScalarConditionPartitionBy(this); return cb;
    }

    // ===================================================================================
    //                                                                       MyselfDerived
    //                                                                       =============
    public void xsmyselfDerive(String fn, SubQuery<VendorLargeDataCB> sq, String al, DerivedReferrerOption op) {
        assertObjectNotNull("subQuery", sq);
        VendorLargeDataCB cb = new VendorLargeDataCB(); cb.xsetupForDerivedReferrer(this);
        lockCall(() -> sq.query(cb)); String pp = keepSpecifyMyselfDerived(cb.query()); String pk = "LARGE_DATA_ID";
        registerSpecifyMyselfDerived(fn, cb.query(), pk, pk, pp, "myselfDerived", al, op);
    }
    public abstract String keepSpecifyMyselfDerived(VendorLargeDataCQ sq);

    /**
     * Prepare for (Query)MyselfDerived (correlated sub-query).
     * @return The object to set up a function for myself table. (NotNull)
     */
    public HpQDRFunction<VendorLargeDataCB> myselfDerived() {
        return xcreateQDRFunctionMyselfDerived(VendorLargeDataCB.class);
    }
    @SuppressWarnings("unchecked")
    protected <CB extends ConditionBean> void xqderiveMyselfDerived(String fn, SubQuery<CB> sq, String rd, Object vl, DerivedReferrerOption op) {
        assertObjectNotNull("subQuery", sq);
        VendorLargeDataCB cb = new VendorLargeDataCB(); cb.xsetupForDerivedReferrer(this); sq.query((CB)cb);
        String pk = "LARGE_DATA_ID";
        String sqpp = keepQueryMyselfDerived(cb.query()); // for saving query-value.
        String prpp = keepQueryMyselfDerivedParameter(vl);
        registerQueryMyselfDerived(fn, cb.query(), pk, pk, sqpp, "myselfDerived", rd, vl, prpp, op);
    }
    public abstract String keepQueryMyselfDerived(VendorLargeDataCQ sq);
    public abstract String keepQueryMyselfDerivedParameter(Object vl);

    // ===================================================================================
    //                                                                        MyselfExists
    //                                                                        ============
    /**
     * Prepare for MyselfExists (correlated sub-query).
     * @param subCBLambda The implementation of sub-query. (NotNull)
     */
    public void myselfExists(SubQuery<VendorLargeDataCB> subCBLambda) {
        assertObjectNotNull("subCBLambda", subCBLambda);
        VendorLargeDataCB cb = new VendorLargeDataCB(); cb.xsetupForMyselfExists(this);
        lockCall(() -> subCBLambda.query(cb)); String pp = keepMyselfExists(cb.query());
        registerMyselfExists(cb.query(), pp);
    }
    public abstract String keepMyselfExists(VendorLargeDataCQ sq);

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
    protected VendorLargeDataCB newMyCB() {
        return new VendorLargeDataCB();
    }
    // very internal (for suppressing warn about 'Not Use Import')
    protected String xabUDT() { return Date.class.getName(); }
    protected String xabCQ() { return VendorLargeDataCQ.class.getName(); }
    protected String xabLSO() { return LikeSearchOption.class.getName(); }
    protected String xabSSQS() { return HpSSQSetupper.class.getName(); }
    protected String xabSCP() { return SubQuery.class.getName(); }
}

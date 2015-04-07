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
 * The abstract condition-query of vendor_large_data_ref.
 * @author DBFlute(AutoGenerator)
 */
public abstract class AbstractBsVendorLargeDataRefCQ extends AbstractConditionQuery {

    // ===================================================================================
    //                                                                         Constructor
    //                                                                         ===========
    public AbstractBsVendorLargeDataRefCQ(ConditionQuery referrerQuery, SqlClause sqlClause, String aliasName, int nestLevel) {
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
        return "vendor_large_data_ref";
    }

    // ===================================================================================
    //                                                                               Query
    //                                                                               =====
    /**
     * Equal(=). And NullIgnored, OnlyOnceRegistered. <br>
     * LARGE_DATA_REF_ID: {PK, NotNull, BIGINT(19)}
     * @param largeDataRefId The value of largeDataRefId as equal. (NullAllowed: if null, no condition)
     */
    public void setLargeDataRefId_Equal(Long largeDataRefId) {
        doSetLargeDataRefId_Equal(largeDataRefId);
    }

    protected void doSetLargeDataRefId_Equal(Long largeDataRefId) {
        regLargeDataRefId(CK_EQ, largeDataRefId);
    }

    /**
     * NotEqual(&lt;&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * LARGE_DATA_REF_ID: {PK, NotNull, BIGINT(19)}
     * @param largeDataRefId The value of largeDataRefId as notEqual. (NullAllowed: if null, no condition)
     */
    public void setLargeDataRefId_NotEqual(Long largeDataRefId) {
        doSetLargeDataRefId_NotEqual(largeDataRefId);
    }

    protected void doSetLargeDataRefId_NotEqual(Long largeDataRefId) {
        regLargeDataRefId(CK_NES, largeDataRefId);
    }

    /**
     * GreaterThan(&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * LARGE_DATA_REF_ID: {PK, NotNull, BIGINT(19)}
     * @param largeDataRefId The value of largeDataRefId as greaterThan. (NullAllowed: if null, no condition)
     */
    public void setLargeDataRefId_GreaterThan(Long largeDataRefId) {
        regLargeDataRefId(CK_GT, largeDataRefId);
    }

    /**
     * LessThan(&lt;). And NullIgnored, OnlyOnceRegistered. <br>
     * LARGE_DATA_REF_ID: {PK, NotNull, BIGINT(19)}
     * @param largeDataRefId The value of largeDataRefId as lessThan. (NullAllowed: if null, no condition)
     */
    public void setLargeDataRefId_LessThan(Long largeDataRefId) {
        regLargeDataRefId(CK_LT, largeDataRefId);
    }

    /**
     * GreaterEqual(&gt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * LARGE_DATA_REF_ID: {PK, NotNull, BIGINT(19)}
     * @param largeDataRefId The value of largeDataRefId as greaterEqual. (NullAllowed: if null, no condition)
     */
    public void setLargeDataRefId_GreaterEqual(Long largeDataRefId) {
        regLargeDataRefId(CK_GE, largeDataRefId);
    }

    /**
     * LessEqual(&lt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * LARGE_DATA_REF_ID: {PK, NotNull, BIGINT(19)}
     * @param largeDataRefId The value of largeDataRefId as lessEqual. (NullAllowed: if null, no condition)
     */
    public void setLargeDataRefId_LessEqual(Long largeDataRefId) {
        regLargeDataRefId(CK_LE, largeDataRefId);
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * LARGE_DATA_REF_ID: {PK, NotNull, BIGINT(19)}
     * @param minNumber The min number of largeDataRefId. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of largeDataRefId. (NullAllowed: if null, no to-condition)
     * @param opLambda The callback for option of range-of. (NotNull)
     */
    public void setLargeDataRefId_RangeOf(Long minNumber, Long maxNumber, ConditionOptionCall<RangeOfOption> opLambda) {
        setLargeDataRefId_RangeOf(minNumber, maxNumber, xcROOP(opLambda));
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * LARGE_DATA_REF_ID: {PK, NotNull, BIGINT(19)}
     * @param minNumber The min number of largeDataRefId. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of largeDataRefId. (NullAllowed: if null, no to-condition)
     * @param rangeOfOption The option of range-of. (NotNull)
     */
    protected void setLargeDataRefId_RangeOf(Long minNumber, Long maxNumber, RangeOfOption rangeOfOption) {
        regROO(minNumber, maxNumber, xgetCValueLargeDataRefId(), "LARGE_DATA_REF_ID", rangeOfOption);
    }

    /**
     * InScope {in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * LARGE_DATA_REF_ID: {PK, NotNull, BIGINT(19)}
     * @param largeDataRefIdList The collection of largeDataRefId as inScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setLargeDataRefId_InScope(Collection<Long> largeDataRefIdList) {
        doSetLargeDataRefId_InScope(largeDataRefIdList);
    }

    protected void doSetLargeDataRefId_InScope(Collection<Long> largeDataRefIdList) {
        regINS(CK_INS, cTL(largeDataRefIdList), xgetCValueLargeDataRefId(), "LARGE_DATA_REF_ID");
    }

    /**
     * NotInScope {not in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * LARGE_DATA_REF_ID: {PK, NotNull, BIGINT(19)}
     * @param largeDataRefIdList The collection of largeDataRefId as notInScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setLargeDataRefId_NotInScope(Collection<Long> largeDataRefIdList) {
        doSetLargeDataRefId_NotInScope(largeDataRefIdList);
    }

    protected void doSetLargeDataRefId_NotInScope(Collection<Long> largeDataRefIdList) {
        regINS(CK_NINS, cTL(largeDataRefIdList), xgetCValueLargeDataRefId(), "LARGE_DATA_REF_ID");
    }

    /**
     * IsNull {is null}. And OnlyOnceRegistered. <br>
     * LARGE_DATA_REF_ID: {PK, NotNull, BIGINT(19)}
     */
    public void setLargeDataRefId_IsNull() { regLargeDataRefId(CK_ISN, DOBJ); }

    /**
     * IsNotNull {is not null}. And OnlyOnceRegistered. <br>
     * LARGE_DATA_REF_ID: {PK, NotNull, BIGINT(19)}
     */
    public void setLargeDataRefId_IsNotNull() { regLargeDataRefId(CK_ISNN, DOBJ); }

    protected void regLargeDataRefId(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueLargeDataRefId(), "LARGE_DATA_REF_ID"); }
    protected abstract ConditionValue xgetCValueLargeDataRefId();

    /**
     * Equal(=). And NullIgnored, OnlyOnceRegistered. <br>
     * LARGE_DATA_ID: {NotNull, BIGINT(19)}
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
     * LARGE_DATA_ID: {NotNull, BIGINT(19)}
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
     * LARGE_DATA_ID: {NotNull, BIGINT(19)}
     * @param largeDataId The value of largeDataId as greaterThan. (NullAllowed: if null, no condition)
     */
    public void setLargeDataId_GreaterThan(Long largeDataId) {
        regLargeDataId(CK_GT, largeDataId);
    }

    /**
     * LessThan(&lt;). And NullIgnored, OnlyOnceRegistered. <br>
     * LARGE_DATA_ID: {NotNull, BIGINT(19)}
     * @param largeDataId The value of largeDataId as lessThan. (NullAllowed: if null, no condition)
     */
    public void setLargeDataId_LessThan(Long largeDataId) {
        regLargeDataId(CK_LT, largeDataId);
    }

    /**
     * GreaterEqual(&gt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * LARGE_DATA_ID: {NotNull, BIGINT(19)}
     * @param largeDataId The value of largeDataId as greaterEqual. (NullAllowed: if null, no condition)
     */
    public void setLargeDataId_GreaterEqual(Long largeDataId) {
        regLargeDataId(CK_GE, largeDataId);
    }

    /**
     * LessEqual(&lt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * LARGE_DATA_ID: {NotNull, BIGINT(19)}
     * @param largeDataId The value of largeDataId as lessEqual. (NullAllowed: if null, no condition)
     */
    public void setLargeDataId_LessEqual(Long largeDataId) {
        regLargeDataId(CK_LE, largeDataId);
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * LARGE_DATA_ID: {NotNull, BIGINT(19)}
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
     * LARGE_DATA_ID: {NotNull, BIGINT(19)}
     * @param minNumber The min number of largeDataId. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of largeDataId. (NullAllowed: if null, no to-condition)
     * @param rangeOfOption The option of range-of. (NotNull)
     */
    protected void setLargeDataId_RangeOf(Long minNumber, Long maxNumber, RangeOfOption rangeOfOption) {
        regROO(minNumber, maxNumber, xgetCValueLargeDataId(), "LARGE_DATA_ID", rangeOfOption);
    }

    /**
     * InScope {in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * LARGE_DATA_ID: {NotNull, BIGINT(19)}
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
     * LARGE_DATA_ID: {NotNull, BIGINT(19)}
     * @param largeDataIdList The collection of largeDataId as notInScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setLargeDataId_NotInScope(Collection<Long> largeDataIdList) {
        doSetLargeDataId_NotInScope(largeDataIdList);
    }

    protected void doSetLargeDataId_NotInScope(Collection<Long> largeDataIdList) {
        regINS(CK_NINS, cTL(largeDataIdList), xgetCValueLargeDataId(), "LARGE_DATA_ID");
    }

    protected void regLargeDataId(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueLargeDataId(), "LARGE_DATA_ID"); }
    protected abstract ConditionValue xgetCValueLargeDataId();

    /**
     * Equal(=). And NullIgnored, OnlyOnceRegistered. <br>
     * DATE_INDEX: {NotNull, DATE(10)}
     * @param dateIndex The value of dateIndex as equal. (NullAllowed: if null, no condition)
     */
    public void setDateIndex_Equal(java.time.LocalDate dateIndex) {
        regDateIndex(CK_EQ,  dateIndex);
    }

    /**
     * GreaterThan(&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * DATE_INDEX: {NotNull, DATE(10)}
     * @param dateIndex The value of dateIndex as greaterThan. (NullAllowed: if null, no condition)
     */
    public void setDateIndex_GreaterThan(java.time.LocalDate dateIndex) {
        regDateIndex(CK_GT,  dateIndex);
    }

    /**
     * LessThan(&lt;). And NullIgnored, OnlyOnceRegistered. <br>
     * DATE_INDEX: {NotNull, DATE(10)}
     * @param dateIndex The value of dateIndex as lessThan. (NullAllowed: if null, no condition)
     */
    public void setDateIndex_LessThan(java.time.LocalDate dateIndex) {
        regDateIndex(CK_LT,  dateIndex);
    }

    /**
     * GreaterEqual(&gt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * DATE_INDEX: {NotNull, DATE(10)}
     * @param dateIndex The value of dateIndex as greaterEqual. (NullAllowed: if null, no condition)
     */
    public void setDateIndex_GreaterEqual(java.time.LocalDate dateIndex) {
        regDateIndex(CK_GE,  dateIndex);
    }

    /**
     * LessEqual(&lt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * DATE_INDEX: {NotNull, DATE(10)}
     * @param dateIndex The value of dateIndex as lessEqual. (NullAllowed: if null, no condition)
     */
    public void setDateIndex_LessEqual(java.time.LocalDate dateIndex) {
        regDateIndex(CK_LE, dateIndex);
    }

    /**
     * FromTo with various options. (versatile) {(default) fromDatetime &lt;= column &lt;= toDatetime} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * DATE_INDEX: {NotNull, DATE(10)}
     * <pre>e.g. setDateIndex_FromTo(fromDate, toDate, op <span style="color: #90226C; font-weight: bold"><span style="font-size: 120%">-</span>&gt;</span> op.<span style="color: #CC4747">compareAsDate()</span>);</pre>
     * @param fromDatetime The from-datetime(yyyy/MM/dd HH:mm:ss.SSS) of dateIndex. (NullAllowed: if null, no from-condition)
     * @param toDatetime The to-datetime(yyyy/MM/dd HH:mm:ss.SSS) of dateIndex. (NullAllowed: if null, no to-condition)
     * @param opLambda The callback for option of from-to. (NotNull)
     */
    public void setDateIndex_FromTo(java.time.LocalDate fromDatetime, java.time.LocalDate toDatetime, ConditionOptionCall<FromToOption> opLambda) {
        setDateIndex_FromTo(fromDatetime, toDatetime, xcFTOP(opLambda));
    }

    /**
     * FromTo with various options. (versatile) {(default) fromDatetime &lt;= column &lt;= toDatetime} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * DATE_INDEX: {NotNull, DATE(10)}
     * <pre>e.g. setDateIndex_FromTo(fromDate, toDate, new <span style="color: #CC4747">FromToOption</span>().compareAsDate());</pre>
     * @param fromDatetime The from-datetime(yyyy/MM/dd HH:mm:ss.SSS) of dateIndex. (NullAllowed: if null, no from-condition)
     * @param toDatetime The to-datetime(yyyy/MM/dd HH:mm:ss.SSS) of dateIndex. (NullAllowed: if null, no to-condition)
     * @param fromToOption The option of from-to. (NotNull)
     */
    protected void setDateIndex_FromTo(java.time.LocalDate fromDatetime, java.time.LocalDate toDatetime, FromToOption fromToOption) {
        String nm = "DATE_INDEX"; FromToOption op = fromToOption;
        regFTQ(xfFTHD(fromDatetime, nm, op), xfFTHD(toDatetime, nm, op), xgetCValueDateIndex(), nm, op);
    }

    protected void regDateIndex(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueDateIndex(), "DATE_INDEX"); }
    protected abstract ConditionValue xgetCValueDateIndex();

    /**
     * Equal(=). And NullIgnored, OnlyOnceRegistered. <br>
     * DATE_NO_INDEX: {NotNull, DATE(10)}
     * @param dateNoIndex The value of dateNoIndex as equal. (NullAllowed: if null, no condition)
     */
    public void setDateNoIndex_Equal(java.time.LocalDate dateNoIndex) {
        regDateNoIndex(CK_EQ,  dateNoIndex);
    }

    /**
     * GreaterThan(&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * DATE_NO_INDEX: {NotNull, DATE(10)}
     * @param dateNoIndex The value of dateNoIndex as greaterThan. (NullAllowed: if null, no condition)
     */
    public void setDateNoIndex_GreaterThan(java.time.LocalDate dateNoIndex) {
        regDateNoIndex(CK_GT,  dateNoIndex);
    }

    /**
     * LessThan(&lt;). And NullIgnored, OnlyOnceRegistered. <br>
     * DATE_NO_INDEX: {NotNull, DATE(10)}
     * @param dateNoIndex The value of dateNoIndex as lessThan. (NullAllowed: if null, no condition)
     */
    public void setDateNoIndex_LessThan(java.time.LocalDate dateNoIndex) {
        regDateNoIndex(CK_LT,  dateNoIndex);
    }

    /**
     * GreaterEqual(&gt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * DATE_NO_INDEX: {NotNull, DATE(10)}
     * @param dateNoIndex The value of dateNoIndex as greaterEqual. (NullAllowed: if null, no condition)
     */
    public void setDateNoIndex_GreaterEqual(java.time.LocalDate dateNoIndex) {
        regDateNoIndex(CK_GE,  dateNoIndex);
    }

    /**
     * LessEqual(&lt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * DATE_NO_INDEX: {NotNull, DATE(10)}
     * @param dateNoIndex The value of dateNoIndex as lessEqual. (NullAllowed: if null, no condition)
     */
    public void setDateNoIndex_LessEqual(java.time.LocalDate dateNoIndex) {
        regDateNoIndex(CK_LE, dateNoIndex);
    }

    /**
     * FromTo with various options. (versatile) {(default) fromDatetime &lt;= column &lt;= toDatetime} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * DATE_NO_INDEX: {NotNull, DATE(10)}
     * <pre>e.g. setDateNoIndex_FromTo(fromDate, toDate, op <span style="color: #90226C; font-weight: bold"><span style="font-size: 120%">-</span>&gt;</span> op.<span style="color: #CC4747">compareAsDate()</span>);</pre>
     * @param fromDatetime The from-datetime(yyyy/MM/dd HH:mm:ss.SSS) of dateNoIndex. (NullAllowed: if null, no from-condition)
     * @param toDatetime The to-datetime(yyyy/MM/dd HH:mm:ss.SSS) of dateNoIndex. (NullAllowed: if null, no to-condition)
     * @param opLambda The callback for option of from-to. (NotNull)
     */
    public void setDateNoIndex_FromTo(java.time.LocalDate fromDatetime, java.time.LocalDate toDatetime, ConditionOptionCall<FromToOption> opLambda) {
        setDateNoIndex_FromTo(fromDatetime, toDatetime, xcFTOP(opLambda));
    }

    /**
     * FromTo with various options. (versatile) {(default) fromDatetime &lt;= column &lt;= toDatetime} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * DATE_NO_INDEX: {NotNull, DATE(10)}
     * <pre>e.g. setDateNoIndex_FromTo(fromDate, toDate, new <span style="color: #CC4747">FromToOption</span>().compareAsDate());</pre>
     * @param fromDatetime The from-datetime(yyyy/MM/dd HH:mm:ss.SSS) of dateNoIndex. (NullAllowed: if null, no from-condition)
     * @param toDatetime The to-datetime(yyyy/MM/dd HH:mm:ss.SSS) of dateNoIndex. (NullAllowed: if null, no to-condition)
     * @param fromToOption The option of from-to. (NotNull)
     */
    protected void setDateNoIndex_FromTo(java.time.LocalDate fromDatetime, java.time.LocalDate toDatetime, FromToOption fromToOption) {
        String nm = "DATE_NO_INDEX"; FromToOption op = fromToOption;
        regFTQ(xfFTHD(fromDatetime, nm, op), xfFTHD(toDatetime, nm, op), xgetCValueDateNoIndex(), nm, op);
    }

    protected void regDateNoIndex(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueDateNoIndex(), "DATE_NO_INDEX"); }
    protected abstract ConditionValue xgetCValueDateNoIndex();

    /**
     * Equal(=). And NullIgnored, OnlyOnceRegistered. <br>
     * TIMESTAMP_INDEX: {NotNull, DATETIME(19)}
     * @param timestampIndex The value of timestampIndex as equal. (NullAllowed: if null, no condition)
     */
    public void setTimestampIndex_Equal(java.time.LocalDateTime timestampIndex) {
        regTimestampIndex(CK_EQ,  timestampIndex);
    }

    /**
     * GreaterThan(&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * TIMESTAMP_INDEX: {NotNull, DATETIME(19)}
     * @param timestampIndex The value of timestampIndex as greaterThan. (NullAllowed: if null, no condition)
     */
    public void setTimestampIndex_GreaterThan(java.time.LocalDateTime timestampIndex) {
        regTimestampIndex(CK_GT,  timestampIndex);
    }

    /**
     * LessThan(&lt;). And NullIgnored, OnlyOnceRegistered. <br>
     * TIMESTAMP_INDEX: {NotNull, DATETIME(19)}
     * @param timestampIndex The value of timestampIndex as lessThan. (NullAllowed: if null, no condition)
     */
    public void setTimestampIndex_LessThan(java.time.LocalDateTime timestampIndex) {
        regTimestampIndex(CK_LT,  timestampIndex);
    }

    /**
     * GreaterEqual(&gt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * TIMESTAMP_INDEX: {NotNull, DATETIME(19)}
     * @param timestampIndex The value of timestampIndex as greaterEqual. (NullAllowed: if null, no condition)
     */
    public void setTimestampIndex_GreaterEqual(java.time.LocalDateTime timestampIndex) {
        regTimestampIndex(CK_GE,  timestampIndex);
    }

    /**
     * LessEqual(&lt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * TIMESTAMP_INDEX: {NotNull, DATETIME(19)}
     * @param timestampIndex The value of timestampIndex as lessEqual. (NullAllowed: if null, no condition)
     */
    public void setTimestampIndex_LessEqual(java.time.LocalDateTime timestampIndex) {
        regTimestampIndex(CK_LE, timestampIndex);
    }

    /**
     * FromTo with various options. (versatile) {(default) fromDatetime &lt;= column &lt;= toDatetime} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * TIMESTAMP_INDEX: {NotNull, DATETIME(19)}
     * <pre>e.g. setTimestampIndex_FromTo(fromDate, toDate, op <span style="color: #90226C; font-weight: bold"><span style="font-size: 120%">-</span>&gt;</span> op.<span style="color: #CC4747">compareAsDate()</span>);</pre>
     * @param fromDatetime The from-datetime(yyyy/MM/dd HH:mm:ss.SSS) of timestampIndex. (NullAllowed: if null, no from-condition)
     * @param toDatetime The to-datetime(yyyy/MM/dd HH:mm:ss.SSS) of timestampIndex. (NullAllowed: if null, no to-condition)
     * @param opLambda The callback for option of from-to. (NotNull)
     */
    public void setTimestampIndex_FromTo(java.time.LocalDateTime fromDatetime, java.time.LocalDateTime toDatetime, ConditionOptionCall<FromToOption> opLambda) {
        setTimestampIndex_FromTo(fromDatetime, toDatetime, xcFTOP(opLambda));
    }

    /**
     * FromTo with various options. (versatile) {(default) fromDatetime &lt;= column &lt;= toDatetime} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * TIMESTAMP_INDEX: {NotNull, DATETIME(19)}
     * <pre>e.g. setTimestampIndex_FromTo(fromDate, toDate, new <span style="color: #CC4747">FromToOption</span>().compareAsDate());</pre>
     * @param fromDatetime The from-datetime(yyyy/MM/dd HH:mm:ss.SSS) of timestampIndex. (NullAllowed: if null, no from-condition)
     * @param toDatetime The to-datetime(yyyy/MM/dd HH:mm:ss.SSS) of timestampIndex. (NullAllowed: if null, no to-condition)
     * @param fromToOption The option of from-to. (NotNull)
     */
    protected void setTimestampIndex_FromTo(java.time.LocalDateTime fromDatetime, java.time.LocalDateTime toDatetime, FromToOption fromToOption) {
        String nm = "TIMESTAMP_INDEX"; FromToOption op = fromToOption;
        regFTQ(xfFTHD(fromDatetime, nm, op), xfFTHD(toDatetime, nm, op), xgetCValueTimestampIndex(), nm, op);
    }

    protected void regTimestampIndex(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueTimestampIndex(), "TIMESTAMP_INDEX"); }
    protected abstract ConditionValue xgetCValueTimestampIndex();

    /**
     * Equal(=). And NullIgnored, OnlyOnceRegistered. <br>
     * TIMESTAMP_NO_INDEX: {NotNull, DATETIME(19)}
     * @param timestampNoIndex The value of timestampNoIndex as equal. (NullAllowed: if null, no condition)
     */
    public void setTimestampNoIndex_Equal(java.time.LocalDateTime timestampNoIndex) {
        regTimestampNoIndex(CK_EQ,  timestampNoIndex);
    }

    /**
     * GreaterThan(&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * TIMESTAMP_NO_INDEX: {NotNull, DATETIME(19)}
     * @param timestampNoIndex The value of timestampNoIndex as greaterThan. (NullAllowed: if null, no condition)
     */
    public void setTimestampNoIndex_GreaterThan(java.time.LocalDateTime timestampNoIndex) {
        regTimestampNoIndex(CK_GT,  timestampNoIndex);
    }

    /**
     * LessThan(&lt;). And NullIgnored, OnlyOnceRegistered. <br>
     * TIMESTAMP_NO_INDEX: {NotNull, DATETIME(19)}
     * @param timestampNoIndex The value of timestampNoIndex as lessThan. (NullAllowed: if null, no condition)
     */
    public void setTimestampNoIndex_LessThan(java.time.LocalDateTime timestampNoIndex) {
        regTimestampNoIndex(CK_LT,  timestampNoIndex);
    }

    /**
     * GreaterEqual(&gt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * TIMESTAMP_NO_INDEX: {NotNull, DATETIME(19)}
     * @param timestampNoIndex The value of timestampNoIndex as greaterEqual. (NullAllowed: if null, no condition)
     */
    public void setTimestampNoIndex_GreaterEqual(java.time.LocalDateTime timestampNoIndex) {
        regTimestampNoIndex(CK_GE,  timestampNoIndex);
    }

    /**
     * LessEqual(&lt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * TIMESTAMP_NO_INDEX: {NotNull, DATETIME(19)}
     * @param timestampNoIndex The value of timestampNoIndex as lessEqual. (NullAllowed: if null, no condition)
     */
    public void setTimestampNoIndex_LessEqual(java.time.LocalDateTime timestampNoIndex) {
        regTimestampNoIndex(CK_LE, timestampNoIndex);
    }

    /**
     * FromTo with various options. (versatile) {(default) fromDatetime &lt;= column &lt;= toDatetime} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * TIMESTAMP_NO_INDEX: {NotNull, DATETIME(19)}
     * <pre>e.g. setTimestampNoIndex_FromTo(fromDate, toDate, op <span style="color: #90226C; font-weight: bold"><span style="font-size: 120%">-</span>&gt;</span> op.<span style="color: #CC4747">compareAsDate()</span>);</pre>
     * @param fromDatetime The from-datetime(yyyy/MM/dd HH:mm:ss.SSS) of timestampNoIndex. (NullAllowed: if null, no from-condition)
     * @param toDatetime The to-datetime(yyyy/MM/dd HH:mm:ss.SSS) of timestampNoIndex. (NullAllowed: if null, no to-condition)
     * @param opLambda The callback for option of from-to. (NotNull)
     */
    public void setTimestampNoIndex_FromTo(java.time.LocalDateTime fromDatetime, java.time.LocalDateTime toDatetime, ConditionOptionCall<FromToOption> opLambda) {
        setTimestampNoIndex_FromTo(fromDatetime, toDatetime, xcFTOP(opLambda));
    }

    /**
     * FromTo with various options. (versatile) {(default) fromDatetime &lt;= column &lt;= toDatetime} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * TIMESTAMP_NO_INDEX: {NotNull, DATETIME(19)}
     * <pre>e.g. setTimestampNoIndex_FromTo(fromDate, toDate, new <span style="color: #CC4747">FromToOption</span>().compareAsDate());</pre>
     * @param fromDatetime The from-datetime(yyyy/MM/dd HH:mm:ss.SSS) of timestampNoIndex. (NullAllowed: if null, no from-condition)
     * @param toDatetime The to-datetime(yyyy/MM/dd HH:mm:ss.SSS) of timestampNoIndex. (NullAllowed: if null, no to-condition)
     * @param fromToOption The option of from-to. (NotNull)
     */
    protected void setTimestampNoIndex_FromTo(java.time.LocalDateTime fromDatetime, java.time.LocalDateTime toDatetime, FromToOption fromToOption) {
        String nm = "TIMESTAMP_NO_INDEX"; FromToOption op = fromToOption;
        regFTQ(xfFTHD(fromDatetime, nm, op), xfFTHD(toDatetime, nm, op), xgetCValueTimestampNoIndex(), nm, op);
    }

    protected void regTimestampNoIndex(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueTimestampNoIndex(), "TIMESTAMP_NO_INDEX"); }
    protected abstract ConditionValue xgetCValueTimestampNoIndex();

    /**
     * Equal(=). And NullIgnored, OnlyOnceRegistered. <br>
     * NULLABLE_DECIMAL_INDEX: {DECIMAL(12, 3)}
     * @param nullableDecimalIndex The value of nullableDecimalIndex as equal. (NullAllowed: if null, no condition)
     */
    public void setNullableDecimalIndex_Equal(java.math.BigDecimal nullableDecimalIndex) {
        doSetNullableDecimalIndex_Equal(nullableDecimalIndex);
    }

    protected void doSetNullableDecimalIndex_Equal(java.math.BigDecimal nullableDecimalIndex) {
        regNullableDecimalIndex(CK_EQ, nullableDecimalIndex);
    }

    /**
     * NotEqual(&lt;&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * NULLABLE_DECIMAL_INDEX: {DECIMAL(12, 3)}
     * @param nullableDecimalIndex The value of nullableDecimalIndex as notEqual. (NullAllowed: if null, no condition)
     */
    public void setNullableDecimalIndex_NotEqual(java.math.BigDecimal nullableDecimalIndex) {
        doSetNullableDecimalIndex_NotEqual(nullableDecimalIndex);
    }

    protected void doSetNullableDecimalIndex_NotEqual(java.math.BigDecimal nullableDecimalIndex) {
        regNullableDecimalIndex(CK_NES, nullableDecimalIndex);
    }

    /**
     * GreaterThan(&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * NULLABLE_DECIMAL_INDEX: {DECIMAL(12, 3)}
     * @param nullableDecimalIndex The value of nullableDecimalIndex as greaterThan. (NullAllowed: if null, no condition)
     */
    public void setNullableDecimalIndex_GreaterThan(java.math.BigDecimal nullableDecimalIndex) {
        regNullableDecimalIndex(CK_GT, nullableDecimalIndex);
    }

    /**
     * LessThan(&lt;). And NullIgnored, OnlyOnceRegistered. <br>
     * NULLABLE_DECIMAL_INDEX: {DECIMAL(12, 3)}
     * @param nullableDecimalIndex The value of nullableDecimalIndex as lessThan. (NullAllowed: if null, no condition)
     */
    public void setNullableDecimalIndex_LessThan(java.math.BigDecimal nullableDecimalIndex) {
        regNullableDecimalIndex(CK_LT, nullableDecimalIndex);
    }

    /**
     * GreaterEqual(&gt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * NULLABLE_DECIMAL_INDEX: {DECIMAL(12, 3)}
     * @param nullableDecimalIndex The value of nullableDecimalIndex as greaterEqual. (NullAllowed: if null, no condition)
     */
    public void setNullableDecimalIndex_GreaterEqual(java.math.BigDecimal nullableDecimalIndex) {
        regNullableDecimalIndex(CK_GE, nullableDecimalIndex);
    }

    /**
     * LessEqual(&lt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * NULLABLE_DECIMAL_INDEX: {DECIMAL(12, 3)}
     * @param nullableDecimalIndex The value of nullableDecimalIndex as lessEqual. (NullAllowed: if null, no condition)
     */
    public void setNullableDecimalIndex_LessEqual(java.math.BigDecimal nullableDecimalIndex) {
        regNullableDecimalIndex(CK_LE, nullableDecimalIndex);
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * NULLABLE_DECIMAL_INDEX: {DECIMAL(12, 3)}
     * @param minNumber The min number of nullableDecimalIndex. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of nullableDecimalIndex. (NullAllowed: if null, no to-condition)
     * @param opLambda The callback for option of range-of. (NotNull)
     */
    public void setNullableDecimalIndex_RangeOf(java.math.BigDecimal minNumber, java.math.BigDecimal maxNumber, ConditionOptionCall<RangeOfOption> opLambda) {
        setNullableDecimalIndex_RangeOf(minNumber, maxNumber, xcROOP(opLambda));
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * NULLABLE_DECIMAL_INDEX: {DECIMAL(12, 3)}
     * @param minNumber The min number of nullableDecimalIndex. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of nullableDecimalIndex. (NullAllowed: if null, no to-condition)
     * @param rangeOfOption The option of range-of. (NotNull)
     */
    protected void setNullableDecimalIndex_RangeOf(java.math.BigDecimal minNumber, java.math.BigDecimal maxNumber, RangeOfOption rangeOfOption) {
        regROO(minNumber, maxNumber, xgetCValueNullableDecimalIndex(), "NULLABLE_DECIMAL_INDEX", rangeOfOption);
    }

    /**
     * InScope {in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * NULLABLE_DECIMAL_INDEX: {DECIMAL(12, 3)}
     * @param nullableDecimalIndexList The collection of nullableDecimalIndex as inScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setNullableDecimalIndex_InScope(Collection<java.math.BigDecimal> nullableDecimalIndexList) {
        doSetNullableDecimalIndex_InScope(nullableDecimalIndexList);
    }

    protected void doSetNullableDecimalIndex_InScope(Collection<java.math.BigDecimal> nullableDecimalIndexList) {
        regINS(CK_INS, cTL(nullableDecimalIndexList), xgetCValueNullableDecimalIndex(), "NULLABLE_DECIMAL_INDEX");
    }

    /**
     * NotInScope {not in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * NULLABLE_DECIMAL_INDEX: {DECIMAL(12, 3)}
     * @param nullableDecimalIndexList The collection of nullableDecimalIndex as notInScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setNullableDecimalIndex_NotInScope(Collection<java.math.BigDecimal> nullableDecimalIndexList) {
        doSetNullableDecimalIndex_NotInScope(nullableDecimalIndexList);
    }

    protected void doSetNullableDecimalIndex_NotInScope(Collection<java.math.BigDecimal> nullableDecimalIndexList) {
        regINS(CK_NINS, cTL(nullableDecimalIndexList), xgetCValueNullableDecimalIndex(), "NULLABLE_DECIMAL_INDEX");
    }

    /**
     * IsNull {is null}. And OnlyOnceRegistered. <br>
     * NULLABLE_DECIMAL_INDEX: {DECIMAL(12, 3)}
     */
    public void setNullableDecimalIndex_IsNull() { regNullableDecimalIndex(CK_ISN, DOBJ); }

    /**
     * IsNotNull {is not null}. And OnlyOnceRegistered. <br>
     * NULLABLE_DECIMAL_INDEX: {DECIMAL(12, 3)}
     */
    public void setNullableDecimalIndex_IsNotNull() { regNullableDecimalIndex(CK_ISNN, DOBJ); }

    protected void regNullableDecimalIndex(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueNullableDecimalIndex(), "NULLABLE_DECIMAL_INDEX"); }
    protected abstract ConditionValue xgetCValueNullableDecimalIndex();

    /**
     * Equal(=). And NullIgnored, OnlyOnceRegistered. <br>
     * NULLABLE_DECIMAL_NO_INDEX: {DECIMAL(12, 3)}
     * @param nullableDecimalNoIndex The value of nullableDecimalNoIndex as equal. (NullAllowed: if null, no condition)
     */
    public void setNullableDecimalNoIndex_Equal(java.math.BigDecimal nullableDecimalNoIndex) {
        doSetNullableDecimalNoIndex_Equal(nullableDecimalNoIndex);
    }

    protected void doSetNullableDecimalNoIndex_Equal(java.math.BigDecimal nullableDecimalNoIndex) {
        regNullableDecimalNoIndex(CK_EQ, nullableDecimalNoIndex);
    }

    /**
     * NotEqual(&lt;&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * NULLABLE_DECIMAL_NO_INDEX: {DECIMAL(12, 3)}
     * @param nullableDecimalNoIndex The value of nullableDecimalNoIndex as notEqual. (NullAllowed: if null, no condition)
     */
    public void setNullableDecimalNoIndex_NotEqual(java.math.BigDecimal nullableDecimalNoIndex) {
        doSetNullableDecimalNoIndex_NotEqual(nullableDecimalNoIndex);
    }

    protected void doSetNullableDecimalNoIndex_NotEqual(java.math.BigDecimal nullableDecimalNoIndex) {
        regNullableDecimalNoIndex(CK_NES, nullableDecimalNoIndex);
    }

    /**
     * GreaterThan(&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * NULLABLE_DECIMAL_NO_INDEX: {DECIMAL(12, 3)}
     * @param nullableDecimalNoIndex The value of nullableDecimalNoIndex as greaterThan. (NullAllowed: if null, no condition)
     */
    public void setNullableDecimalNoIndex_GreaterThan(java.math.BigDecimal nullableDecimalNoIndex) {
        regNullableDecimalNoIndex(CK_GT, nullableDecimalNoIndex);
    }

    /**
     * LessThan(&lt;). And NullIgnored, OnlyOnceRegistered. <br>
     * NULLABLE_DECIMAL_NO_INDEX: {DECIMAL(12, 3)}
     * @param nullableDecimalNoIndex The value of nullableDecimalNoIndex as lessThan. (NullAllowed: if null, no condition)
     */
    public void setNullableDecimalNoIndex_LessThan(java.math.BigDecimal nullableDecimalNoIndex) {
        regNullableDecimalNoIndex(CK_LT, nullableDecimalNoIndex);
    }

    /**
     * GreaterEqual(&gt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * NULLABLE_DECIMAL_NO_INDEX: {DECIMAL(12, 3)}
     * @param nullableDecimalNoIndex The value of nullableDecimalNoIndex as greaterEqual. (NullAllowed: if null, no condition)
     */
    public void setNullableDecimalNoIndex_GreaterEqual(java.math.BigDecimal nullableDecimalNoIndex) {
        regNullableDecimalNoIndex(CK_GE, nullableDecimalNoIndex);
    }

    /**
     * LessEqual(&lt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * NULLABLE_DECIMAL_NO_INDEX: {DECIMAL(12, 3)}
     * @param nullableDecimalNoIndex The value of nullableDecimalNoIndex as lessEqual. (NullAllowed: if null, no condition)
     */
    public void setNullableDecimalNoIndex_LessEqual(java.math.BigDecimal nullableDecimalNoIndex) {
        regNullableDecimalNoIndex(CK_LE, nullableDecimalNoIndex);
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * NULLABLE_DECIMAL_NO_INDEX: {DECIMAL(12, 3)}
     * @param minNumber The min number of nullableDecimalNoIndex. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of nullableDecimalNoIndex. (NullAllowed: if null, no to-condition)
     * @param opLambda The callback for option of range-of. (NotNull)
     */
    public void setNullableDecimalNoIndex_RangeOf(java.math.BigDecimal minNumber, java.math.BigDecimal maxNumber, ConditionOptionCall<RangeOfOption> opLambda) {
        setNullableDecimalNoIndex_RangeOf(minNumber, maxNumber, xcROOP(opLambda));
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * NULLABLE_DECIMAL_NO_INDEX: {DECIMAL(12, 3)}
     * @param minNumber The min number of nullableDecimalNoIndex. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of nullableDecimalNoIndex. (NullAllowed: if null, no to-condition)
     * @param rangeOfOption The option of range-of. (NotNull)
     */
    protected void setNullableDecimalNoIndex_RangeOf(java.math.BigDecimal minNumber, java.math.BigDecimal maxNumber, RangeOfOption rangeOfOption) {
        regROO(minNumber, maxNumber, xgetCValueNullableDecimalNoIndex(), "NULLABLE_DECIMAL_NO_INDEX", rangeOfOption);
    }

    /**
     * InScope {in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * NULLABLE_DECIMAL_NO_INDEX: {DECIMAL(12, 3)}
     * @param nullableDecimalNoIndexList The collection of nullableDecimalNoIndex as inScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setNullableDecimalNoIndex_InScope(Collection<java.math.BigDecimal> nullableDecimalNoIndexList) {
        doSetNullableDecimalNoIndex_InScope(nullableDecimalNoIndexList);
    }

    protected void doSetNullableDecimalNoIndex_InScope(Collection<java.math.BigDecimal> nullableDecimalNoIndexList) {
        regINS(CK_INS, cTL(nullableDecimalNoIndexList), xgetCValueNullableDecimalNoIndex(), "NULLABLE_DECIMAL_NO_INDEX");
    }

    /**
     * NotInScope {not in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * NULLABLE_DECIMAL_NO_INDEX: {DECIMAL(12, 3)}
     * @param nullableDecimalNoIndexList The collection of nullableDecimalNoIndex as notInScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setNullableDecimalNoIndex_NotInScope(Collection<java.math.BigDecimal> nullableDecimalNoIndexList) {
        doSetNullableDecimalNoIndex_NotInScope(nullableDecimalNoIndexList);
    }

    protected void doSetNullableDecimalNoIndex_NotInScope(Collection<java.math.BigDecimal> nullableDecimalNoIndexList) {
        regINS(CK_NINS, cTL(nullableDecimalNoIndexList), xgetCValueNullableDecimalNoIndex(), "NULLABLE_DECIMAL_NO_INDEX");
    }

    /**
     * IsNull {is null}. And OnlyOnceRegistered. <br>
     * NULLABLE_DECIMAL_NO_INDEX: {DECIMAL(12, 3)}
     */
    public void setNullableDecimalNoIndex_IsNull() { regNullableDecimalNoIndex(CK_ISN, DOBJ); }

    /**
     * IsNotNull {is not null}. And OnlyOnceRegistered. <br>
     * NULLABLE_DECIMAL_NO_INDEX: {DECIMAL(12, 3)}
     */
    public void setNullableDecimalNoIndex_IsNotNull() { regNullableDecimalNoIndex(CK_ISNN, DOBJ); }

    protected void regNullableDecimalNoIndex(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueNullableDecimalNoIndex(), "NULLABLE_DECIMAL_NO_INDEX"); }
    protected abstract ConditionValue xgetCValueNullableDecimalNoIndex();

    /**
     * Equal(=). And NullIgnored, OnlyOnceRegistered. <br>
     * SELF_PARENT_ID: {BIGINT(19)}
     * @param selfParentId The value of selfParentId as equal. (NullAllowed: if null, no condition)
     */
    public void setSelfParentId_Equal(Long selfParentId) {
        doSetSelfParentId_Equal(selfParentId);
    }

    protected void doSetSelfParentId_Equal(Long selfParentId) {
        regSelfParentId(CK_EQ, selfParentId);
    }

    /**
     * NotEqual(&lt;&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * SELF_PARENT_ID: {BIGINT(19)}
     * @param selfParentId The value of selfParentId as notEqual. (NullAllowed: if null, no condition)
     */
    public void setSelfParentId_NotEqual(Long selfParentId) {
        doSetSelfParentId_NotEqual(selfParentId);
    }

    protected void doSetSelfParentId_NotEqual(Long selfParentId) {
        regSelfParentId(CK_NES, selfParentId);
    }

    /**
     * GreaterThan(&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * SELF_PARENT_ID: {BIGINT(19)}
     * @param selfParentId The value of selfParentId as greaterThan. (NullAllowed: if null, no condition)
     */
    public void setSelfParentId_GreaterThan(Long selfParentId) {
        regSelfParentId(CK_GT, selfParentId);
    }

    /**
     * LessThan(&lt;). And NullIgnored, OnlyOnceRegistered. <br>
     * SELF_PARENT_ID: {BIGINT(19)}
     * @param selfParentId The value of selfParentId as lessThan. (NullAllowed: if null, no condition)
     */
    public void setSelfParentId_LessThan(Long selfParentId) {
        regSelfParentId(CK_LT, selfParentId);
    }

    /**
     * GreaterEqual(&gt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * SELF_PARENT_ID: {BIGINT(19)}
     * @param selfParentId The value of selfParentId as greaterEqual. (NullAllowed: if null, no condition)
     */
    public void setSelfParentId_GreaterEqual(Long selfParentId) {
        regSelfParentId(CK_GE, selfParentId);
    }

    /**
     * LessEqual(&lt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * SELF_PARENT_ID: {BIGINT(19)}
     * @param selfParentId The value of selfParentId as lessEqual. (NullAllowed: if null, no condition)
     */
    public void setSelfParentId_LessEqual(Long selfParentId) {
        regSelfParentId(CK_LE, selfParentId);
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * SELF_PARENT_ID: {BIGINT(19)}
     * @param minNumber The min number of selfParentId. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of selfParentId. (NullAllowed: if null, no to-condition)
     * @param opLambda The callback for option of range-of. (NotNull)
     */
    public void setSelfParentId_RangeOf(Long minNumber, Long maxNumber, ConditionOptionCall<RangeOfOption> opLambda) {
        setSelfParentId_RangeOf(minNumber, maxNumber, xcROOP(opLambda));
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * SELF_PARENT_ID: {BIGINT(19)}
     * @param minNumber The min number of selfParentId. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of selfParentId. (NullAllowed: if null, no to-condition)
     * @param rangeOfOption The option of range-of. (NotNull)
     */
    protected void setSelfParentId_RangeOf(Long minNumber, Long maxNumber, RangeOfOption rangeOfOption) {
        regROO(minNumber, maxNumber, xgetCValueSelfParentId(), "SELF_PARENT_ID", rangeOfOption);
    }

    /**
     * InScope {in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * SELF_PARENT_ID: {BIGINT(19)}
     * @param selfParentIdList The collection of selfParentId as inScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setSelfParentId_InScope(Collection<Long> selfParentIdList) {
        doSetSelfParentId_InScope(selfParentIdList);
    }

    protected void doSetSelfParentId_InScope(Collection<Long> selfParentIdList) {
        regINS(CK_INS, cTL(selfParentIdList), xgetCValueSelfParentId(), "SELF_PARENT_ID");
    }

    /**
     * NotInScope {not in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * SELF_PARENT_ID: {BIGINT(19)}
     * @param selfParentIdList The collection of selfParentId as notInScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setSelfParentId_NotInScope(Collection<Long> selfParentIdList) {
        doSetSelfParentId_NotInScope(selfParentIdList);
    }

    protected void doSetSelfParentId_NotInScope(Collection<Long> selfParentIdList) {
        regINS(CK_NINS, cTL(selfParentIdList), xgetCValueSelfParentId(), "SELF_PARENT_ID");
    }

    /**
     * IsNull {is null}. And OnlyOnceRegistered. <br>
     * SELF_PARENT_ID: {BIGINT(19)}
     */
    public void setSelfParentId_IsNull() { regSelfParentId(CK_ISN, DOBJ); }

    /**
     * IsNotNull {is not null}. And OnlyOnceRegistered. <br>
     * SELF_PARENT_ID: {BIGINT(19)}
     */
    public void setSelfParentId_IsNotNull() { regSelfParentId(CK_ISNN, DOBJ); }

    protected void regSelfParentId(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueSelfParentId(), "SELF_PARENT_ID"); }
    protected abstract ConditionValue xgetCValueSelfParentId();

    // ===================================================================================
    //                                                                     ScalarCondition
    //                                                                     ===============
    /**
     * Prepare ScalarCondition as equal. <br>
     * {where FOO = (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_Equal()</span>.max(new SubQuery&lt;VendorLargeDataRefCB&gt;() {
     *     public void query(VendorLargeDataRefCB subCB) {
     *         subCB.specify().setXxx... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setYyy...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorLargeDataRefCB> scalar_Equal() {
        return xcreateSSQFunction(CK_EQ, VendorLargeDataRefCB.class);
    }

    /**
     * Prepare ScalarCondition as equal. <br>
     * {where FOO &lt;&gt; (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_NotEqual()</span>.max(new SubQuery&lt;VendorLargeDataRefCB&gt;() {
     *     public void query(VendorLargeDataRefCB subCB) {
     *         subCB.specify().setXxx... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setYyy...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorLargeDataRefCB> scalar_NotEqual() {
        return xcreateSSQFunction(CK_NES, VendorLargeDataRefCB.class);
    }

    /**
     * Prepare ScalarCondition as greaterThan. <br>
     * {where FOO &gt; (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_GreaterThan()</span>.max(new SubQuery&lt;VendorLargeDataRefCB&gt;() {
     *     public void query(VendorLargeDataRefCB subCB) {
     *         subCB.specify().setFoo... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setBar...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorLargeDataRefCB> scalar_GreaterThan() {
        return xcreateSSQFunction(CK_GT, VendorLargeDataRefCB.class);
    }

    /**
     * Prepare ScalarCondition as lessThan. <br>
     * {where FOO &lt; (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_LessThan()</span>.max(new SubQuery&lt;VendorLargeDataRefCB&gt;() {
     *     public void query(VendorLargeDataRefCB subCB) {
     *         subCB.specify().setFoo... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setBar...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorLargeDataRefCB> scalar_LessThan() {
        return xcreateSSQFunction(CK_LT, VendorLargeDataRefCB.class);
    }

    /**
     * Prepare ScalarCondition as greaterEqual. <br>
     * {where FOO &gt;= (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_GreaterEqual()</span>.max(new SubQuery&lt;VendorLargeDataRefCB&gt;() {
     *     public void query(VendorLargeDataRefCB subCB) {
     *         subCB.specify().setFoo... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setBar...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorLargeDataRefCB> scalar_GreaterEqual() {
        return xcreateSSQFunction(CK_GE, VendorLargeDataRefCB.class);
    }

    /**
     * Prepare ScalarCondition as lessEqual. <br>
     * {where FOO &lt;= (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_LessEqual()</span>.max(new SubQuery&lt;VendorLargeDataRefCB&gt;() {
     *     public void query(VendorLargeDataRefCB subCB) {
     *         subCB.specify().setFoo... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setBar...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorLargeDataRefCB> scalar_LessEqual() {
        return xcreateSSQFunction(CK_LE, VendorLargeDataRefCB.class);
    }

    @SuppressWarnings("unchecked")
    protected <CB extends ConditionBean> void xscalarCondition(String fn, SubQuery<CB> sq, String rd, HpSSQOption<CB> op) {
        assertObjectNotNull("subQuery", sq);
        VendorLargeDataRefCB cb = xcreateScalarConditionCB(); sq.query((CB)cb);
        String pp = keepScalarCondition(cb.query()); // for saving query-value
        op.setPartitionByCBean((CB)xcreateScalarConditionPartitionByCB()); // for using partition-by
        registerScalarCondition(fn, cb.query(), pp, rd, op);
    }
    public abstract String keepScalarCondition(VendorLargeDataRefCQ sq);

    protected VendorLargeDataRefCB xcreateScalarConditionCB() {
        VendorLargeDataRefCB cb = newMyCB(); cb.xsetupForScalarCondition(this); return cb;
    }

    protected VendorLargeDataRefCB xcreateScalarConditionPartitionByCB() {
        VendorLargeDataRefCB cb = newMyCB(); cb.xsetupForScalarConditionPartitionBy(this); return cb;
    }

    // ===================================================================================
    //                                                                       MyselfDerived
    //                                                                       =============
    public void xsmyselfDerive(String fn, SubQuery<VendorLargeDataRefCB> sq, String al, DerivedReferrerOption op) {
        assertObjectNotNull("subQuery", sq);
        VendorLargeDataRefCB cb = new VendorLargeDataRefCB(); cb.xsetupForDerivedReferrer(this);
        lockCall(() -> sq.query(cb)); String pp = keepSpecifyMyselfDerived(cb.query()); String pk = "LARGE_DATA_REF_ID";
        registerSpecifyMyselfDerived(fn, cb.query(), pk, pk, pp, "myselfDerived", al, op);
    }
    public abstract String keepSpecifyMyselfDerived(VendorLargeDataRefCQ sq);

    /**
     * Prepare for (Query)MyselfDerived (correlated sub-query).
     * @return The object to set up a function for myself table. (NotNull)
     */
    public HpQDRFunction<VendorLargeDataRefCB> myselfDerived() {
        return xcreateQDRFunctionMyselfDerived(VendorLargeDataRefCB.class);
    }
    @SuppressWarnings("unchecked")
    protected <CB extends ConditionBean> void xqderiveMyselfDerived(String fn, SubQuery<CB> sq, String rd, Object vl, DerivedReferrerOption op) {
        assertObjectNotNull("subQuery", sq);
        VendorLargeDataRefCB cb = new VendorLargeDataRefCB(); cb.xsetupForDerivedReferrer(this); sq.query((CB)cb);
        String pk = "LARGE_DATA_REF_ID";
        String sqpp = keepQueryMyselfDerived(cb.query()); // for saving query-value.
        String prpp = keepQueryMyselfDerivedParameter(vl);
        registerQueryMyselfDerived(fn, cb.query(), pk, pk, sqpp, "myselfDerived", rd, vl, prpp, op);
    }
    public abstract String keepQueryMyselfDerived(VendorLargeDataRefCQ sq);
    public abstract String keepQueryMyselfDerivedParameter(Object vl);

    // ===================================================================================
    //                                                                        MyselfExists
    //                                                                        ============
    /**
     * Prepare for MyselfExists (correlated sub-query).
     * @param subCBLambda The implementation of sub-query. (NotNull)
     */
    public void myselfExists(SubQuery<VendorLargeDataRefCB> subCBLambda) {
        assertObjectNotNull("subCBLambda", subCBLambda);
        VendorLargeDataRefCB cb = new VendorLargeDataRefCB(); cb.xsetupForMyselfExists(this);
        lockCall(() -> subCBLambda.query(cb)); String pp = keepMyselfExists(cb.query());
        registerMyselfExists(cb.query(), pp);
    }
    public abstract String keepMyselfExists(VendorLargeDataRefCQ sq);

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
    protected VendorLargeDataRefCB newMyCB() {
        return new VendorLargeDataRefCB();
    }
    // very internal (for suppressing warn about 'Not Use Import')
    protected String xabUDT() { return Date.class.getName(); }
    protected String xabCQ() { return VendorLargeDataRefCQ.class.getName(); }
    protected String xabLSO() { return LikeSearchOption.class.getName(); }
    protected String xabSSQS() { return HpSSQSetupper.class.getName(); }
    protected String xabSCP() { return SubQuery.class.getName(); }
}

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
 * The abstract condition-query of VENDOR_SYMMETRIC.
 * @author DBFlute(AutoGenerator)
 */
public abstract class AbstractBsVendorSymmetricCQ extends AbstractConditionQuery {

    // ===================================================================================
    //                                                                         Constructor
    //                                                                         ===========
    public AbstractBsVendorSymmetricCQ(ConditionQuery referrerQuery, SqlClause sqlClause, String aliasName, int nestLevel) {
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
        return "VENDOR_SYMMETRIC";
    }

    // ===================================================================================
    //                                                                               Query
    //                                                                               =====
    /**
     * Equal(=). And NullIgnored, OnlyOnceRegistered. <br>
     * VENDOR_SYMMETRIC_ID: {PK, NotNull, numeric(16)}
     * @param vendorSymmetricId The value of vendorSymmetricId as equal. (NullAllowed: if null, no condition)
     */
    public void setVendorSymmetricId_Equal(Long vendorSymmetricId) {
        doSetVendorSymmetricId_Equal(vendorSymmetricId);
    }

    protected void doSetVendorSymmetricId_Equal(Long vendorSymmetricId) {
        regVendorSymmetricId(CK_EQ, vendorSymmetricId);
    }

    /**
     * NotEqual(&lt;&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * VENDOR_SYMMETRIC_ID: {PK, NotNull, numeric(16)}
     * @param vendorSymmetricId The value of vendorSymmetricId as notEqual. (NullAllowed: if null, no condition)
     */
    public void setVendorSymmetricId_NotEqual(Long vendorSymmetricId) {
        doSetVendorSymmetricId_NotEqual(vendorSymmetricId);
    }

    protected void doSetVendorSymmetricId_NotEqual(Long vendorSymmetricId) {
        regVendorSymmetricId(CK_NES, vendorSymmetricId);
    }

    /**
     * GreaterThan(&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * VENDOR_SYMMETRIC_ID: {PK, NotNull, numeric(16)}
     * @param vendorSymmetricId The value of vendorSymmetricId as greaterThan. (NullAllowed: if null, no condition)
     */
    public void setVendorSymmetricId_GreaterThan(Long vendorSymmetricId) {
        regVendorSymmetricId(CK_GT, vendorSymmetricId);
    }

    /**
     * LessThan(&lt;). And NullIgnored, OnlyOnceRegistered. <br>
     * VENDOR_SYMMETRIC_ID: {PK, NotNull, numeric(16)}
     * @param vendorSymmetricId The value of vendorSymmetricId as lessThan. (NullAllowed: if null, no condition)
     */
    public void setVendorSymmetricId_LessThan(Long vendorSymmetricId) {
        regVendorSymmetricId(CK_LT, vendorSymmetricId);
    }

    /**
     * GreaterEqual(&gt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * VENDOR_SYMMETRIC_ID: {PK, NotNull, numeric(16)}
     * @param vendorSymmetricId The value of vendorSymmetricId as greaterEqual. (NullAllowed: if null, no condition)
     */
    public void setVendorSymmetricId_GreaterEqual(Long vendorSymmetricId) {
        regVendorSymmetricId(CK_GE, vendorSymmetricId);
    }

    /**
     * LessEqual(&lt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * VENDOR_SYMMETRIC_ID: {PK, NotNull, numeric(16)}
     * @param vendorSymmetricId The value of vendorSymmetricId as lessEqual. (NullAllowed: if null, no condition)
     */
    public void setVendorSymmetricId_LessEqual(Long vendorSymmetricId) {
        regVendorSymmetricId(CK_LE, vendorSymmetricId);
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * VENDOR_SYMMETRIC_ID: {PK, NotNull, numeric(16)}
     * @param minNumber The min number of vendorSymmetricId. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of vendorSymmetricId. (NullAllowed: if null, no to-condition)
     * @param opLambda The callback for option of range-of. (NotNull)
     */
    public void setVendorSymmetricId_RangeOf(Long minNumber, Long maxNumber, ConditionOptionCall<RangeOfOption> opLambda) {
        setVendorSymmetricId_RangeOf(minNumber, maxNumber, xcROOP(opLambda));
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * VENDOR_SYMMETRIC_ID: {PK, NotNull, numeric(16)}
     * @param minNumber The min number of vendorSymmetricId. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of vendorSymmetricId. (NullAllowed: if null, no to-condition)
     * @param rangeOfOption The option of range-of. (NotNull)
     */
    protected void setVendorSymmetricId_RangeOf(Long minNumber, Long maxNumber, RangeOfOption rangeOfOption) {
        regROO(minNumber, maxNumber, xgetCValueVendorSymmetricId(), "VENDOR_SYMMETRIC_ID", rangeOfOption);
    }

    /**
     * InScope {in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * VENDOR_SYMMETRIC_ID: {PK, NotNull, numeric(16)}
     * @param vendorSymmetricIdList The collection of vendorSymmetricId as inScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setVendorSymmetricId_InScope(Collection<Long> vendorSymmetricIdList) {
        doSetVendorSymmetricId_InScope(vendorSymmetricIdList);
    }

    protected void doSetVendorSymmetricId_InScope(Collection<Long> vendorSymmetricIdList) {
        regINS(CK_INS, cTL(vendorSymmetricIdList), xgetCValueVendorSymmetricId(), "VENDOR_SYMMETRIC_ID");
    }

    /**
     * NotInScope {not in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * VENDOR_SYMMETRIC_ID: {PK, NotNull, numeric(16)}
     * @param vendorSymmetricIdList The collection of vendorSymmetricId as notInScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setVendorSymmetricId_NotInScope(Collection<Long> vendorSymmetricIdList) {
        doSetVendorSymmetricId_NotInScope(vendorSymmetricIdList);
    }

    protected void doSetVendorSymmetricId_NotInScope(Collection<Long> vendorSymmetricIdList) {
        regINS(CK_NINS, cTL(vendorSymmetricIdList), xgetCValueVendorSymmetricId(), "VENDOR_SYMMETRIC_ID");
    }

    /**
     * IsNull {is null}. And OnlyOnceRegistered. <br>
     * VENDOR_SYMMETRIC_ID: {PK, NotNull, numeric(16)}
     */
    public void setVendorSymmetricId_IsNull() { regVendorSymmetricId(CK_ISN, DOBJ); }

    /**
     * IsNotNull {is not null}. And OnlyOnceRegistered. <br>
     * VENDOR_SYMMETRIC_ID: {PK, NotNull, numeric(16)}
     */
    public void setVendorSymmetricId_IsNotNull() { regVendorSymmetricId(CK_ISNN, DOBJ); }

    protected void regVendorSymmetricId(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueVendorSymmetricId(), "VENDOR_SYMMETRIC_ID"); }
    protected abstract ConditionValue xgetCValueVendorSymmetricId();

    /**
     * Equal(=). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * PLAIN_TEXT: {nvarchar(100)}
     * @param plainText The value of plainText as equal. (NullAllowed: if null (or empty), no condition)
     */
    public void setPlainText_Equal(String plainText) {
        doSetPlainText_Equal(fRES(plainText));
    }

    protected void doSetPlainText_Equal(String plainText) {
        regPlainText(CK_EQ, plainText);
    }

    /**
     * NotEqual(&lt;&gt;). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * PLAIN_TEXT: {nvarchar(100)}
     * @param plainText The value of plainText as notEqual. (NullAllowed: if null (or empty), no condition)
     */
    public void setPlainText_NotEqual(String plainText) {
        doSetPlainText_NotEqual(fRES(plainText));
    }

    protected void doSetPlainText_NotEqual(String plainText) {
        regPlainText(CK_NES, plainText);
    }

    /**
     * GreaterThan(&gt;). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * PLAIN_TEXT: {nvarchar(100)}
     * @param plainText The value of plainText as greaterThan. (NullAllowed: if null (or empty), no condition)
     */
    public void setPlainText_GreaterThan(String plainText) {
        regPlainText(CK_GT, fRES(plainText));
    }

    /**
     * LessThan(&lt;). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * PLAIN_TEXT: {nvarchar(100)}
     * @param plainText The value of plainText as lessThan. (NullAllowed: if null (or empty), no condition)
     */
    public void setPlainText_LessThan(String plainText) {
        regPlainText(CK_LT, fRES(plainText));
    }

    /**
     * GreaterEqual(&gt;=). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * PLAIN_TEXT: {nvarchar(100)}
     * @param plainText The value of plainText as greaterEqual. (NullAllowed: if null (or empty), no condition)
     */
    public void setPlainText_GreaterEqual(String plainText) {
        regPlainText(CK_GE, fRES(plainText));
    }

    /**
     * LessEqual(&lt;=). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * PLAIN_TEXT: {nvarchar(100)}
     * @param plainText The value of plainText as lessEqual. (NullAllowed: if null (or empty), no condition)
     */
    public void setPlainText_LessEqual(String plainText) {
        regPlainText(CK_LE, fRES(plainText));
    }

    /**
     * InScope {in ('a', 'b')}. And NullOrEmptyIgnored, NullOrEmptyElementIgnored, SeveralRegistered. <br>
     * PLAIN_TEXT: {nvarchar(100)}
     * @param plainTextList The collection of plainText as inScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setPlainText_InScope(Collection<String> plainTextList) {
        doSetPlainText_InScope(plainTextList);
    }

    protected void doSetPlainText_InScope(Collection<String> plainTextList) {
        regINS(CK_INS, cTL(plainTextList), xgetCValuePlainText(), "PLAIN_TEXT");
    }

    /**
     * NotInScope {not in ('a', 'b')}. And NullOrEmptyIgnored, NullOrEmptyElementIgnored, SeveralRegistered. <br>
     * PLAIN_TEXT: {nvarchar(100)}
     * @param plainTextList The collection of plainText as notInScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setPlainText_NotInScope(Collection<String> plainTextList) {
        doSetPlainText_NotInScope(plainTextList);
    }

    protected void doSetPlainText_NotInScope(Collection<String> plainTextList) {
        regINS(CK_NINS, cTL(plainTextList), xgetCValuePlainText(), "PLAIN_TEXT");
    }

    /**
     * LikeSearch with various options. (versatile) {like '%xxx%' escape ...}. And NullOrEmptyIgnored, SeveralRegistered. <br>
     * PLAIN_TEXT: {nvarchar(100)} <br>
     * <pre>e.g. setPlainText_LikeSearch("xxx", op <span style="color: #90226C; font-weight: bold"><span style="font-size: 120%">-</span>&gt;</span> op.<span style="color: #CC4747">likeContain()</span>);</pre>
     * @param plainText The value of plainText as likeSearch. (NullAllowed: if null (or empty), no condition)
     * @param opLambda The callback for option of like-search. (NotNull)
     */
    public void setPlainText_LikeSearch(String plainText, ConditionOptionCall<LikeSearchOption> opLambda) {
        setPlainText_LikeSearch(plainText, xcLSOP(opLambda));
    }

    /**
     * LikeSearch with various options. (versatile) {like '%xxx%' escape ...}. And NullOrEmptyIgnored, SeveralRegistered. <br>
     * PLAIN_TEXT: {nvarchar(100)} <br>
     * <pre>e.g. setPlainText_LikeSearch("xxx", new <span style="color: #CC4747">LikeSearchOption</span>().likeContain());</pre>
     * @param plainText The value of plainText as likeSearch. (NullAllowed: if null (or empty), no condition)
     * @param likeSearchOption The option of like-search. (NotNull)
     */
    protected void setPlainText_LikeSearch(String plainText, LikeSearchOption likeSearchOption) {
        regLSQ(CK_LS, fRES(plainText), xgetCValuePlainText(), "PLAIN_TEXT", likeSearchOption);
    }

    /**
     * NotLikeSearch with various options. (versatile) {not like 'xxx%' escape ...} <br>
     * And NullOrEmptyIgnored, SeveralRegistered. <br>
     * PLAIN_TEXT: {nvarchar(100)}
     * @param plainText The value of plainText as notLikeSearch. (NullAllowed: if null (or empty), no condition)
     * @param opLambda The callback for option of like-search. (NotNull)
     */
    public void setPlainText_NotLikeSearch(String plainText, ConditionOptionCall<LikeSearchOption> opLambda) {
        setPlainText_NotLikeSearch(plainText, xcLSOP(opLambda));
    }

    /**
     * NotLikeSearch with various options. (versatile) {not like 'xxx%' escape ...} <br>
     * And NullOrEmptyIgnored, SeveralRegistered. <br>
     * PLAIN_TEXT: {nvarchar(100)}
     * @param plainText The value of plainText as notLikeSearch. (NullAllowed: if null (or empty), no condition)
     * @param likeSearchOption The option of not-like-search. (NotNull)
     */
    protected void setPlainText_NotLikeSearch(String plainText, LikeSearchOption likeSearchOption) {
        regLSQ(CK_NLS, fRES(plainText), xgetCValuePlainText(), "PLAIN_TEXT", likeSearchOption);
    }

    /**
     * IsNull {is null}. And OnlyOnceRegistered. <br>
     * PLAIN_TEXT: {nvarchar(100)}
     */
    public void setPlainText_IsNull() { regPlainText(CK_ISN, DOBJ); }

    /**
     * IsNullOrEmpty {is null or empty}. And OnlyOnceRegistered. <br>
     * PLAIN_TEXT: {nvarchar(100)}
     */
    public void setPlainText_IsNullOrEmpty() { regPlainText(CK_ISNOE, DOBJ); }

    /**
     * IsNotNull {is not null}. And OnlyOnceRegistered. <br>
     * PLAIN_TEXT: {nvarchar(100)}
     */
    public void setPlainText_IsNotNull() { regPlainText(CK_ISNN, DOBJ); }

    protected void regPlainText(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValuePlainText(), "PLAIN_TEXT"); }
    protected abstract ConditionValue xgetCValuePlainText();


    /**
     * IsNull {is null}. And OnlyOnceRegistered. <br>
     * ENCRYPTED_DATA: {varbinary(2147483647)}
     */
    public void setEncryptedData_IsNull() { regEncryptedData(CK_ISN, DOBJ); }

    /**
     * IsNotNull {is not null}. And OnlyOnceRegistered. <br>
     * ENCRYPTED_DATA: {varbinary(2147483647)}
     */
    public void setEncryptedData_IsNotNull() { regEncryptedData(CK_ISNN, DOBJ); }

    protected void regEncryptedData(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueEncryptedData(), "ENCRYPTED_DATA"); }
    protected abstract ConditionValue xgetCValueEncryptedData();

    // ===================================================================================
    //                                                                     ScalarCondition
    //                                                                     ===============
    /**
     * Prepare ScalarCondition as equal. <br>
     * {where FOO = (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_Equal()</span>.max(new SubQuery&lt;VendorSymmetricCB&gt;() {
     *     public void query(VendorSymmetricCB subCB) {
     *         subCB.specify().setXxx... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setYyy...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorSymmetricCB> scalar_Equal() {
        return xcreateSSQFunction(CK_EQ, VendorSymmetricCB.class);
    }

    /**
     * Prepare ScalarCondition as equal. <br>
     * {where FOO &lt;&gt; (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_NotEqual()</span>.max(new SubQuery&lt;VendorSymmetricCB&gt;() {
     *     public void query(VendorSymmetricCB subCB) {
     *         subCB.specify().setXxx... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setYyy...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorSymmetricCB> scalar_NotEqual() {
        return xcreateSSQFunction(CK_NES, VendorSymmetricCB.class);
    }

    /**
     * Prepare ScalarCondition as greaterThan. <br>
     * {where FOO &gt; (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_GreaterThan()</span>.max(new SubQuery&lt;VendorSymmetricCB&gt;() {
     *     public void query(VendorSymmetricCB subCB) {
     *         subCB.specify().setFoo... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setBar...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorSymmetricCB> scalar_GreaterThan() {
        return xcreateSSQFunction(CK_GT, VendorSymmetricCB.class);
    }

    /**
     * Prepare ScalarCondition as lessThan. <br>
     * {where FOO &lt; (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_LessThan()</span>.max(new SubQuery&lt;VendorSymmetricCB&gt;() {
     *     public void query(VendorSymmetricCB subCB) {
     *         subCB.specify().setFoo... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setBar...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorSymmetricCB> scalar_LessThan() {
        return xcreateSSQFunction(CK_LT, VendorSymmetricCB.class);
    }

    /**
     * Prepare ScalarCondition as greaterEqual. <br>
     * {where FOO &gt;= (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_GreaterEqual()</span>.max(new SubQuery&lt;VendorSymmetricCB&gt;() {
     *     public void query(VendorSymmetricCB subCB) {
     *         subCB.specify().setFoo... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setBar...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorSymmetricCB> scalar_GreaterEqual() {
        return xcreateSSQFunction(CK_GE, VendorSymmetricCB.class);
    }

    /**
     * Prepare ScalarCondition as lessEqual. <br>
     * {where FOO &lt;= (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_LessEqual()</span>.max(new SubQuery&lt;VendorSymmetricCB&gt;() {
     *     public void query(VendorSymmetricCB subCB) {
     *         subCB.specify().setFoo... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setBar...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorSymmetricCB> scalar_LessEqual() {
        return xcreateSSQFunction(CK_LE, VendorSymmetricCB.class);
    }

    @SuppressWarnings("unchecked")
    protected <CB extends ConditionBean> void xscalarCondition(String fn, SubQuery<CB> sq, String rd, HpSSQOption<CB> op) {
        assertObjectNotNull("subQuery", sq);
        VendorSymmetricCB cb = xcreateScalarConditionCB(); sq.query((CB)cb);
        String pp = keepScalarCondition(cb.query()); // for saving query-value
        op.setPartitionByCBean((CB)xcreateScalarConditionPartitionByCB()); // for using partition-by
        registerScalarCondition(fn, cb.query(), pp, rd, op);
    }
    public abstract String keepScalarCondition(VendorSymmetricCQ sq);

    protected VendorSymmetricCB xcreateScalarConditionCB() {
        VendorSymmetricCB cb = newMyCB(); cb.xsetupForScalarCondition(this); return cb;
    }

    protected VendorSymmetricCB xcreateScalarConditionPartitionByCB() {
        VendorSymmetricCB cb = newMyCB(); cb.xsetupForScalarConditionPartitionBy(this); return cb;
    }

    // ===================================================================================
    //                                                                       MyselfDerived
    //                                                                       =============
    public void xsmyselfDerive(String fn, SubQuery<VendorSymmetricCB> sq, String al, DerivedReferrerOption op) {
        assertObjectNotNull("subQuery", sq);
        VendorSymmetricCB cb = new VendorSymmetricCB(); cb.xsetupForDerivedReferrer(this);
        lockCall(() -> sq.query(cb)); String pp = keepSpecifyMyselfDerived(cb.query()); String pk = "VENDOR_SYMMETRIC_ID";
        registerSpecifyMyselfDerived(fn, cb.query(), pk, pk, pp, "myselfDerived", al, op);
    }
    public abstract String keepSpecifyMyselfDerived(VendorSymmetricCQ sq);

    /**
     * Prepare for (Query)MyselfDerived (correlated sub-query).
     * @return The object to set up a function for myself table. (NotNull)
     */
    public HpQDRFunction<VendorSymmetricCB> myselfDerived() {
        return xcreateQDRFunctionMyselfDerived(VendorSymmetricCB.class);
    }
    @SuppressWarnings("unchecked")
    protected <CB extends ConditionBean> void xqderiveMyselfDerived(String fn, SubQuery<CB> sq, String rd, Object vl, DerivedReferrerOption op) {
        assertObjectNotNull("subQuery", sq);
        VendorSymmetricCB cb = new VendorSymmetricCB(); cb.xsetupForDerivedReferrer(this); sq.query((CB)cb);
        String pk = "VENDOR_SYMMETRIC_ID";
        String sqpp = keepQueryMyselfDerived(cb.query()); // for saving query-value.
        String prpp = keepQueryMyselfDerivedParameter(vl);
        registerQueryMyselfDerived(fn, cb.query(), pk, pk, sqpp, "myselfDerived", rd, vl, prpp, op);
    }
    public abstract String keepQueryMyselfDerived(VendorSymmetricCQ sq);
    public abstract String keepQueryMyselfDerivedParameter(Object vl);

    // ===================================================================================
    //                                                                        MyselfExists
    //                                                                        ============
    /**
     * Prepare for MyselfExists (correlated sub-query).
     * @param subCBLambda The implementation of sub-query. (NotNull)
     */
    public void myselfExists(SubQuery<VendorSymmetricCB> subCBLambda) {
        assertObjectNotNull("subCBLambda", subCBLambda);
        VendorSymmetricCB cb = new VendorSymmetricCB(); cb.xsetupForMyselfExists(this);
        lockCall(() -> subCBLambda.query(cb)); String pp = keepMyselfExists(cb.query());
        registerMyselfExists(cb.query(), pp);
    }
    public abstract String keepMyselfExists(VendorSymmetricCQ sq);

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
    protected VendorSymmetricCB newMyCB() {
        return new VendorSymmetricCB();
    }
    // very internal (for suppressing warn about 'Not Use Import')
    protected String xabUDT() { return Date.class.getName(); }
    protected String xabCQ() { return VendorSymmetricCQ.class.getName(); }
    protected String xabLSO() { return LikeSearchOption.class.getName(); }
    protected String xabSSQS() { return HpSSQSetupper.class.getName(); }
    protected String xabSCP() { return SubQuery.class.getName(); }
}

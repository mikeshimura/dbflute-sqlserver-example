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
 * The abstract condition-query of vendor_constraint_name_auto_qux.
 * @author DBFlute(AutoGenerator)
 */
public abstract class AbstractBsVendorConstraintNameAutoQuxCQ extends AbstractConditionQuery {

    // ===================================================================================
    //                                                                         Constructor
    //                                                                         ===========
    public AbstractBsVendorConstraintNameAutoQuxCQ(ConditionQuery referrerQuery, SqlClause sqlClause, String aliasName, int nestLevel) {
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
        return "vendor_constraint_name_auto_qux";
    }

    // ===================================================================================
    //                                                                               Query
    //                                                                               =====
    /**
     * Equal(=). And NullIgnored, OnlyOnceRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_ID: {PK, NotNull, DECIMAL(16)}
     * @param constraintNameAutoQuxId The value of constraintNameAutoQuxId as equal. (NullAllowed: if null, no condition)
     */
    public void setConstraintNameAutoQuxId_Equal(Long constraintNameAutoQuxId) {
        doSetConstraintNameAutoQuxId_Equal(constraintNameAutoQuxId);
    }

    protected void doSetConstraintNameAutoQuxId_Equal(Long constraintNameAutoQuxId) {
        regConstraintNameAutoQuxId(CK_EQ, constraintNameAutoQuxId);
    }

    /**
     * NotEqual(&lt;&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_ID: {PK, NotNull, DECIMAL(16)}
     * @param constraintNameAutoQuxId The value of constraintNameAutoQuxId as notEqual. (NullAllowed: if null, no condition)
     */
    public void setConstraintNameAutoQuxId_NotEqual(Long constraintNameAutoQuxId) {
        doSetConstraintNameAutoQuxId_NotEqual(constraintNameAutoQuxId);
    }

    protected void doSetConstraintNameAutoQuxId_NotEqual(Long constraintNameAutoQuxId) {
        regConstraintNameAutoQuxId(CK_NES, constraintNameAutoQuxId);
    }

    /**
     * GreaterThan(&gt;). And NullIgnored, OnlyOnceRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_ID: {PK, NotNull, DECIMAL(16)}
     * @param constraintNameAutoQuxId The value of constraintNameAutoQuxId as greaterThan. (NullAllowed: if null, no condition)
     */
    public void setConstraintNameAutoQuxId_GreaterThan(Long constraintNameAutoQuxId) {
        regConstraintNameAutoQuxId(CK_GT, constraintNameAutoQuxId);
    }

    /**
     * LessThan(&lt;). And NullIgnored, OnlyOnceRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_ID: {PK, NotNull, DECIMAL(16)}
     * @param constraintNameAutoQuxId The value of constraintNameAutoQuxId as lessThan. (NullAllowed: if null, no condition)
     */
    public void setConstraintNameAutoQuxId_LessThan(Long constraintNameAutoQuxId) {
        regConstraintNameAutoQuxId(CK_LT, constraintNameAutoQuxId);
    }

    /**
     * GreaterEqual(&gt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_ID: {PK, NotNull, DECIMAL(16)}
     * @param constraintNameAutoQuxId The value of constraintNameAutoQuxId as greaterEqual. (NullAllowed: if null, no condition)
     */
    public void setConstraintNameAutoQuxId_GreaterEqual(Long constraintNameAutoQuxId) {
        regConstraintNameAutoQuxId(CK_GE, constraintNameAutoQuxId);
    }

    /**
     * LessEqual(&lt;=). And NullIgnored, OnlyOnceRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_ID: {PK, NotNull, DECIMAL(16)}
     * @param constraintNameAutoQuxId The value of constraintNameAutoQuxId as lessEqual. (NullAllowed: if null, no condition)
     */
    public void setConstraintNameAutoQuxId_LessEqual(Long constraintNameAutoQuxId) {
        regConstraintNameAutoQuxId(CK_LE, constraintNameAutoQuxId);
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_ID: {PK, NotNull, DECIMAL(16)}
     * @param minNumber The min number of constraintNameAutoQuxId. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of constraintNameAutoQuxId. (NullAllowed: if null, no to-condition)
     * @param opLambda The callback for option of range-of. (NotNull)
     */
    public void setConstraintNameAutoQuxId_RangeOf(Long minNumber, Long maxNumber, ConditionOptionCall<RangeOfOption> opLambda) {
        setConstraintNameAutoQuxId_RangeOf(minNumber, maxNumber, xcROOP(opLambda));
    }

    /**
     * RangeOf with various options. (versatile) <br>
     * {(default) minNumber &lt;= column &lt;= maxNumber} <br>
     * And NullIgnored, OnlyOnceRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_ID: {PK, NotNull, DECIMAL(16)}
     * @param minNumber The min number of constraintNameAutoQuxId. (NullAllowed: if null, no from-condition)
     * @param maxNumber The max number of constraintNameAutoQuxId. (NullAllowed: if null, no to-condition)
     * @param rangeOfOption The option of range-of. (NotNull)
     */
    protected void setConstraintNameAutoQuxId_RangeOf(Long minNumber, Long maxNumber, RangeOfOption rangeOfOption) {
        regROO(minNumber, maxNumber, xgetCValueConstraintNameAutoQuxId(), "CONSTRAINT_NAME_AUTO_QUX_ID", rangeOfOption);
    }

    /**
     * InScope {in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_ID: {PK, NotNull, DECIMAL(16)}
     * @param constraintNameAutoQuxIdList The collection of constraintNameAutoQuxId as inScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setConstraintNameAutoQuxId_InScope(Collection<Long> constraintNameAutoQuxIdList) {
        doSetConstraintNameAutoQuxId_InScope(constraintNameAutoQuxIdList);
    }

    protected void doSetConstraintNameAutoQuxId_InScope(Collection<Long> constraintNameAutoQuxIdList) {
        regINS(CK_INS, cTL(constraintNameAutoQuxIdList), xgetCValueConstraintNameAutoQuxId(), "CONSTRAINT_NAME_AUTO_QUX_ID");
    }

    /**
     * NotInScope {not in (1, 2)}. And NullIgnored, NullElementIgnored, SeveralRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_ID: {PK, NotNull, DECIMAL(16)}
     * @param constraintNameAutoQuxIdList The collection of constraintNameAutoQuxId as notInScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setConstraintNameAutoQuxId_NotInScope(Collection<Long> constraintNameAutoQuxIdList) {
        doSetConstraintNameAutoQuxId_NotInScope(constraintNameAutoQuxIdList);
    }

    protected void doSetConstraintNameAutoQuxId_NotInScope(Collection<Long> constraintNameAutoQuxIdList) {
        regINS(CK_NINS, cTL(constraintNameAutoQuxIdList), xgetCValueConstraintNameAutoQuxId(), "CONSTRAINT_NAME_AUTO_QUX_ID");
    }

    /**
     * Set up ExistsReferrer (correlated sub-query). <br>
     * {exists (select CONSTRAINT_NAME_AUTO_QUX_ID from vendor_constraint_name_auto_ref where ...)} <br>
     * vendor_constraint_name_auto_ref by CONSTRAINT_NAME_AUTO_QUX_ID, named 'vendorConstraintNameAutoRefAsOne'.
     * <pre>
     * cb.query().<span style="color: #CC4747">existsVendorConstraintNameAutoRef</span>(refCB <span style="color: #90226C; font-weight: bold"><span style="font-size: 120%">-</span>&gt;</span> {
     *     refCB.query().set...
     * });
     * </pre>
     * @param subCBLambda The callback for sub-query of VendorConstraintNameAutoRefList for 'exists'. (NotNull)
     */
    public void existsVendorConstraintNameAutoRef(SubQuery<VendorConstraintNameAutoRefCB> subCBLambda) {
        assertObjectNotNull("subCBLambda", subCBLambda);
        VendorConstraintNameAutoRefCB cb = new VendorConstraintNameAutoRefCB(); cb.xsetupForExistsReferrer(this);
        lockCall(() -> subCBLambda.query(cb)); String pp = keepConstraintNameAutoQuxId_ExistsReferrer_VendorConstraintNameAutoRefList(cb.query());
        registerExistsReferrer(cb.query(), "CONSTRAINT_NAME_AUTO_QUX_ID", "CONSTRAINT_NAME_AUTO_QUX_ID", pp, "vendorConstraintNameAutoRefList");
    }
    public abstract String keepConstraintNameAutoQuxId_ExistsReferrer_VendorConstraintNameAutoRefList(VendorConstraintNameAutoRefCQ sq);

    /**
     * Set up NotExistsReferrer (correlated sub-query). <br>
     * {not exists (select CONSTRAINT_NAME_AUTO_QUX_ID from vendor_constraint_name_auto_ref where ...)} <br>
     * vendor_constraint_name_auto_ref by CONSTRAINT_NAME_AUTO_QUX_ID, named 'vendorConstraintNameAutoRefAsOne'.
     * <pre>
     * cb.query().<span style="color: #CC4747">notExistsVendorConstraintNameAutoRef</span>(refCB <span style="color: #90226C; font-weight: bold"><span style="font-size: 120%">-</span>&gt;</span> {
     *     refCB.query().set...
     * });
     * </pre>
     * @param subCBLambda The callback for sub-query of ConstraintNameAutoQuxId_NotExistsReferrer_VendorConstraintNameAutoRefList for 'not exists'. (NotNull)
     */
    public void notExistsVendorConstraintNameAutoRef(SubQuery<VendorConstraintNameAutoRefCB> subCBLambda) {
        assertObjectNotNull("subCBLambda", subCBLambda);
        VendorConstraintNameAutoRefCB cb = new VendorConstraintNameAutoRefCB(); cb.xsetupForExistsReferrer(this);
        lockCall(() -> subCBLambda.query(cb)); String pp = keepConstraintNameAutoQuxId_NotExistsReferrer_VendorConstraintNameAutoRefList(cb.query());
        registerNotExistsReferrer(cb.query(), "CONSTRAINT_NAME_AUTO_QUX_ID", "CONSTRAINT_NAME_AUTO_QUX_ID", pp, "vendorConstraintNameAutoRefList");
    }
    public abstract String keepConstraintNameAutoQuxId_NotExistsReferrer_VendorConstraintNameAutoRefList(VendorConstraintNameAutoRefCQ sq);

    public void xsderiveVendorConstraintNameAutoRefList(String fn, SubQuery<VendorConstraintNameAutoRefCB> sq, String al, DerivedReferrerOption op) {
        assertObjectNotNull("subQuery", sq);
        VendorConstraintNameAutoRefCB cb = new VendorConstraintNameAutoRefCB(); cb.xsetupForDerivedReferrer(this);
        lockCall(() -> sq.query(cb)); String pp = keepConstraintNameAutoQuxId_SpecifyDerivedReferrer_VendorConstraintNameAutoRefList(cb.query());
        registerSpecifyDerivedReferrer(fn, cb.query(), "CONSTRAINT_NAME_AUTO_QUX_ID", "CONSTRAINT_NAME_AUTO_QUX_ID", pp, "vendorConstraintNameAutoRefList", al, op);
    }
    public abstract String keepConstraintNameAutoQuxId_SpecifyDerivedReferrer_VendorConstraintNameAutoRefList(VendorConstraintNameAutoRefCQ sq);

    /**
     * Prepare for (Query)DerivedReferrer (correlated sub-query). <br>
     * {FOO &lt;= (select max(BAR) from vendor_constraint_name_auto_ref where ...)} <br>
     * vendor_constraint_name_auto_ref by CONSTRAINT_NAME_AUTO_QUX_ID, named 'vendorConstraintNameAutoRefAsOne'.
     * <pre>
     * cb.query().<span style="color: #CC4747">derivedVendorConstraintNameAutoRef()</span>.<span style="color: #CC4747">max</span>(refCB <span style="color: #90226C; font-weight: bold"><span style="font-size: 120%">-</span>&gt;</span> {
     *     refCB.specify().<span style="color: #CC4747">columnFoo...</span> <span style="color: #3F7E5E">// derived column by function</span>
     *     refCB.query().setBar... <span style="color: #3F7E5E">// referrer condition</span>
     * }).<span style="color: #CC4747">greaterEqual</span>(123); <span style="color: #3F7E5E">// condition to derived column</span>
     * </pre>
     * @return The object to set up a function for referrer table. (NotNull)
     */
    public HpQDRFunction<VendorConstraintNameAutoRefCB> derivedVendorConstraintNameAutoRef() {
        return xcreateQDRFunctionVendorConstraintNameAutoRefList();
    }
    protected HpQDRFunction<VendorConstraintNameAutoRefCB> xcreateQDRFunctionVendorConstraintNameAutoRefList() {
        return xcQDRFunc((fn, sq, rd, vl, op) -> xqderiveVendorConstraintNameAutoRefList(fn, sq, rd, vl, op));
    }
    public void xqderiveVendorConstraintNameAutoRefList(String fn, SubQuery<VendorConstraintNameAutoRefCB> sq, String rd, Object vl, DerivedReferrerOption op) {
        assertObjectNotNull("subQuery", sq);
        VendorConstraintNameAutoRefCB cb = new VendorConstraintNameAutoRefCB(); cb.xsetupForDerivedReferrer(this);
        lockCall(() -> sq.query(cb)); String sqpp = keepConstraintNameAutoQuxId_QueryDerivedReferrer_VendorConstraintNameAutoRefList(cb.query()); String prpp = keepConstraintNameAutoQuxId_QueryDerivedReferrer_VendorConstraintNameAutoRefListParameter(vl);
        registerQueryDerivedReferrer(fn, cb.query(), "CONSTRAINT_NAME_AUTO_QUX_ID", "CONSTRAINT_NAME_AUTO_QUX_ID", sqpp, "vendorConstraintNameAutoRefList", rd, vl, prpp, op);
    }
    public abstract String keepConstraintNameAutoQuxId_QueryDerivedReferrer_VendorConstraintNameAutoRefList(VendorConstraintNameAutoRefCQ sq);
    public abstract String keepConstraintNameAutoQuxId_QueryDerivedReferrer_VendorConstraintNameAutoRefListParameter(Object vl);

    /**
     * IsNull {is null}. And OnlyOnceRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_ID: {PK, NotNull, DECIMAL(16)}
     */
    public void setConstraintNameAutoQuxId_IsNull() { regConstraintNameAutoQuxId(CK_ISN, DOBJ); }

    /**
     * IsNotNull {is not null}. And OnlyOnceRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_ID: {PK, NotNull, DECIMAL(16)}
     */
    public void setConstraintNameAutoQuxId_IsNotNull() { regConstraintNameAutoQuxId(CK_ISNN, DOBJ); }

    protected void regConstraintNameAutoQuxId(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueConstraintNameAutoQuxId(), "CONSTRAINT_NAME_AUTO_QUX_ID"); }
    protected abstract ConditionValue xgetCValueConstraintNameAutoQuxId();

    /**
     * Equal(=). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_NAME: {UQ, NotNull, VARCHAR(50)}
     * @param constraintNameAutoQuxName The value of constraintNameAutoQuxName as equal. (NullAllowed: if null (or empty), no condition)
     */
    public void setConstraintNameAutoQuxName_Equal(String constraintNameAutoQuxName) {
        doSetConstraintNameAutoQuxName_Equal(fRES(constraintNameAutoQuxName));
    }

    protected void doSetConstraintNameAutoQuxName_Equal(String constraintNameAutoQuxName) {
        regConstraintNameAutoQuxName(CK_EQ, constraintNameAutoQuxName);
    }

    /**
     * NotEqual(&lt;&gt;). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_NAME: {UQ, NotNull, VARCHAR(50)}
     * @param constraintNameAutoQuxName The value of constraintNameAutoQuxName as notEqual. (NullAllowed: if null (or empty), no condition)
     */
    public void setConstraintNameAutoQuxName_NotEqual(String constraintNameAutoQuxName) {
        doSetConstraintNameAutoQuxName_NotEqual(fRES(constraintNameAutoQuxName));
    }

    protected void doSetConstraintNameAutoQuxName_NotEqual(String constraintNameAutoQuxName) {
        regConstraintNameAutoQuxName(CK_NES, constraintNameAutoQuxName);
    }

    /**
     * GreaterThan(&gt;). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_NAME: {UQ, NotNull, VARCHAR(50)}
     * @param constraintNameAutoQuxName The value of constraintNameAutoQuxName as greaterThan. (NullAllowed: if null (or empty), no condition)
     */
    public void setConstraintNameAutoQuxName_GreaterThan(String constraintNameAutoQuxName) {
        regConstraintNameAutoQuxName(CK_GT, fRES(constraintNameAutoQuxName));
    }

    /**
     * LessThan(&lt;). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_NAME: {UQ, NotNull, VARCHAR(50)}
     * @param constraintNameAutoQuxName The value of constraintNameAutoQuxName as lessThan. (NullAllowed: if null (or empty), no condition)
     */
    public void setConstraintNameAutoQuxName_LessThan(String constraintNameAutoQuxName) {
        regConstraintNameAutoQuxName(CK_LT, fRES(constraintNameAutoQuxName));
    }

    /**
     * GreaterEqual(&gt;=). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_NAME: {UQ, NotNull, VARCHAR(50)}
     * @param constraintNameAutoQuxName The value of constraintNameAutoQuxName as greaterEqual. (NullAllowed: if null (or empty), no condition)
     */
    public void setConstraintNameAutoQuxName_GreaterEqual(String constraintNameAutoQuxName) {
        regConstraintNameAutoQuxName(CK_GE, fRES(constraintNameAutoQuxName));
    }

    /**
     * LessEqual(&lt;=). And NullOrEmptyIgnored, OnlyOnceRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_NAME: {UQ, NotNull, VARCHAR(50)}
     * @param constraintNameAutoQuxName The value of constraintNameAutoQuxName as lessEqual. (NullAllowed: if null (or empty), no condition)
     */
    public void setConstraintNameAutoQuxName_LessEqual(String constraintNameAutoQuxName) {
        regConstraintNameAutoQuxName(CK_LE, fRES(constraintNameAutoQuxName));
    }

    /**
     * InScope {in ('a', 'b')}. And NullOrEmptyIgnored, NullOrEmptyElementIgnored, SeveralRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_NAME: {UQ, NotNull, VARCHAR(50)}
     * @param constraintNameAutoQuxNameList The collection of constraintNameAutoQuxName as inScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setConstraintNameAutoQuxName_InScope(Collection<String> constraintNameAutoQuxNameList) {
        doSetConstraintNameAutoQuxName_InScope(constraintNameAutoQuxNameList);
    }

    protected void doSetConstraintNameAutoQuxName_InScope(Collection<String> constraintNameAutoQuxNameList) {
        regINS(CK_INS, cTL(constraintNameAutoQuxNameList), xgetCValueConstraintNameAutoQuxName(), "CONSTRAINT_NAME_AUTO_QUX_NAME");
    }

    /**
     * NotInScope {not in ('a', 'b')}. And NullOrEmptyIgnored, NullOrEmptyElementIgnored, SeveralRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_NAME: {UQ, NotNull, VARCHAR(50)}
     * @param constraintNameAutoQuxNameList The collection of constraintNameAutoQuxName as notInScope. (NullAllowed: if null (or empty), no condition)
     */
    public void setConstraintNameAutoQuxName_NotInScope(Collection<String> constraintNameAutoQuxNameList) {
        doSetConstraintNameAutoQuxName_NotInScope(constraintNameAutoQuxNameList);
    }

    protected void doSetConstraintNameAutoQuxName_NotInScope(Collection<String> constraintNameAutoQuxNameList) {
        regINS(CK_NINS, cTL(constraintNameAutoQuxNameList), xgetCValueConstraintNameAutoQuxName(), "CONSTRAINT_NAME_AUTO_QUX_NAME");
    }

    /**
     * LikeSearch with various options. (versatile) {like '%xxx%' escape ...}. And NullOrEmptyIgnored, SeveralRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_NAME: {UQ, NotNull, VARCHAR(50)} <br>
     * <pre>e.g. setConstraintNameAutoQuxName_LikeSearch("xxx", op <span style="color: #90226C; font-weight: bold"><span style="font-size: 120%">-</span>&gt;</span> op.<span style="color: #CC4747">likeContain()</span>);</pre>
     * @param constraintNameAutoQuxName The value of constraintNameAutoQuxName as likeSearch. (NullAllowed: if null (or empty), no condition)
     * @param opLambda The callback for option of like-search. (NotNull)
     */
    public void setConstraintNameAutoQuxName_LikeSearch(String constraintNameAutoQuxName, ConditionOptionCall<LikeSearchOption> opLambda) {
        setConstraintNameAutoQuxName_LikeSearch(constraintNameAutoQuxName, xcLSOP(opLambda));
    }

    /**
     * LikeSearch with various options. (versatile) {like '%xxx%' escape ...}. And NullOrEmptyIgnored, SeveralRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_NAME: {UQ, NotNull, VARCHAR(50)} <br>
     * <pre>e.g. setConstraintNameAutoQuxName_LikeSearch("xxx", new <span style="color: #CC4747">LikeSearchOption</span>().likeContain());</pre>
     * @param constraintNameAutoQuxName The value of constraintNameAutoQuxName as likeSearch. (NullAllowed: if null (or empty), no condition)
     * @param likeSearchOption The option of like-search. (NotNull)
     */
    protected void setConstraintNameAutoQuxName_LikeSearch(String constraintNameAutoQuxName, LikeSearchOption likeSearchOption) {
        regLSQ(CK_LS, fRES(constraintNameAutoQuxName), xgetCValueConstraintNameAutoQuxName(), "CONSTRAINT_NAME_AUTO_QUX_NAME", likeSearchOption);
    }

    /**
     * NotLikeSearch with various options. (versatile) {not like 'xxx%' escape ...} <br>
     * And NullOrEmptyIgnored, SeveralRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_NAME: {UQ, NotNull, VARCHAR(50)}
     * @param constraintNameAutoQuxName The value of constraintNameAutoQuxName as notLikeSearch. (NullAllowed: if null (or empty), no condition)
     * @param opLambda The callback for option of like-search. (NotNull)
     */
    public void setConstraintNameAutoQuxName_NotLikeSearch(String constraintNameAutoQuxName, ConditionOptionCall<LikeSearchOption> opLambda) {
        setConstraintNameAutoQuxName_NotLikeSearch(constraintNameAutoQuxName, xcLSOP(opLambda));
    }

    /**
     * NotLikeSearch with various options. (versatile) {not like 'xxx%' escape ...} <br>
     * And NullOrEmptyIgnored, SeveralRegistered. <br>
     * CONSTRAINT_NAME_AUTO_QUX_NAME: {UQ, NotNull, VARCHAR(50)}
     * @param constraintNameAutoQuxName The value of constraintNameAutoQuxName as notLikeSearch. (NullAllowed: if null (or empty), no condition)
     * @param likeSearchOption The option of not-like-search. (NotNull)
     */
    protected void setConstraintNameAutoQuxName_NotLikeSearch(String constraintNameAutoQuxName, LikeSearchOption likeSearchOption) {
        regLSQ(CK_NLS, fRES(constraintNameAutoQuxName), xgetCValueConstraintNameAutoQuxName(), "CONSTRAINT_NAME_AUTO_QUX_NAME", likeSearchOption);
    }

    protected void regConstraintNameAutoQuxName(ConditionKey ky, Object vl) { regQ(ky, vl, xgetCValueConstraintNameAutoQuxName(), "CONSTRAINT_NAME_AUTO_QUX_NAME"); }
    protected abstract ConditionValue xgetCValueConstraintNameAutoQuxName();

    // ===================================================================================
    //                                                                     ScalarCondition
    //                                                                     ===============
    /**
     * Prepare ScalarCondition as equal. <br>
     * {where FOO = (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_Equal()</span>.max(new SubQuery&lt;VendorConstraintNameAutoQuxCB&gt;() {
     *     public void query(VendorConstraintNameAutoQuxCB subCB) {
     *         subCB.specify().setXxx... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setYyy...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorConstraintNameAutoQuxCB> scalar_Equal() {
        return xcreateSSQFunction(CK_EQ, VendorConstraintNameAutoQuxCB.class);
    }

    /**
     * Prepare ScalarCondition as equal. <br>
     * {where FOO &lt;&gt; (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_NotEqual()</span>.max(new SubQuery&lt;VendorConstraintNameAutoQuxCB&gt;() {
     *     public void query(VendorConstraintNameAutoQuxCB subCB) {
     *         subCB.specify().setXxx... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setYyy...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorConstraintNameAutoQuxCB> scalar_NotEqual() {
        return xcreateSSQFunction(CK_NES, VendorConstraintNameAutoQuxCB.class);
    }

    /**
     * Prepare ScalarCondition as greaterThan. <br>
     * {where FOO &gt; (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_GreaterThan()</span>.max(new SubQuery&lt;VendorConstraintNameAutoQuxCB&gt;() {
     *     public void query(VendorConstraintNameAutoQuxCB subCB) {
     *         subCB.specify().setFoo... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setBar...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorConstraintNameAutoQuxCB> scalar_GreaterThan() {
        return xcreateSSQFunction(CK_GT, VendorConstraintNameAutoQuxCB.class);
    }

    /**
     * Prepare ScalarCondition as lessThan. <br>
     * {where FOO &lt; (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_LessThan()</span>.max(new SubQuery&lt;VendorConstraintNameAutoQuxCB&gt;() {
     *     public void query(VendorConstraintNameAutoQuxCB subCB) {
     *         subCB.specify().setFoo... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setBar...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorConstraintNameAutoQuxCB> scalar_LessThan() {
        return xcreateSSQFunction(CK_LT, VendorConstraintNameAutoQuxCB.class);
    }

    /**
     * Prepare ScalarCondition as greaterEqual. <br>
     * {where FOO &gt;= (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_GreaterEqual()</span>.max(new SubQuery&lt;VendorConstraintNameAutoQuxCB&gt;() {
     *     public void query(VendorConstraintNameAutoQuxCB subCB) {
     *         subCB.specify().setFoo... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setBar...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorConstraintNameAutoQuxCB> scalar_GreaterEqual() {
        return xcreateSSQFunction(CK_GE, VendorConstraintNameAutoQuxCB.class);
    }

    /**
     * Prepare ScalarCondition as lessEqual. <br>
     * {where FOO &lt;= (select max(BAR) from ...)
     * <pre>
     * cb.query().<span style="color: #CC4747">scalar_LessEqual()</span>.max(new SubQuery&lt;VendorConstraintNameAutoQuxCB&gt;() {
     *     public void query(VendorConstraintNameAutoQuxCB subCB) {
     *         subCB.specify().setFoo... <span style="color: #3F7E5E">// derived column for function</span>
     *         subCB.query().setBar...
     *     }
     * });
     * </pre>
     * @return The object to set up a function. (NotNull)
     */
    public HpSSQFunction<VendorConstraintNameAutoQuxCB> scalar_LessEqual() {
        return xcreateSSQFunction(CK_LE, VendorConstraintNameAutoQuxCB.class);
    }

    @SuppressWarnings("unchecked")
    protected <CB extends ConditionBean> void xscalarCondition(String fn, SubQuery<CB> sq, String rd, HpSSQOption<CB> op) {
        assertObjectNotNull("subQuery", sq);
        VendorConstraintNameAutoQuxCB cb = xcreateScalarConditionCB(); sq.query((CB)cb);
        String pp = keepScalarCondition(cb.query()); // for saving query-value
        op.setPartitionByCBean((CB)xcreateScalarConditionPartitionByCB()); // for using partition-by
        registerScalarCondition(fn, cb.query(), pp, rd, op);
    }
    public abstract String keepScalarCondition(VendorConstraintNameAutoQuxCQ sq);

    protected VendorConstraintNameAutoQuxCB xcreateScalarConditionCB() {
        VendorConstraintNameAutoQuxCB cb = newMyCB(); cb.xsetupForScalarCondition(this); return cb;
    }

    protected VendorConstraintNameAutoQuxCB xcreateScalarConditionPartitionByCB() {
        VendorConstraintNameAutoQuxCB cb = newMyCB(); cb.xsetupForScalarConditionPartitionBy(this); return cb;
    }

    // ===================================================================================
    //                                                                       MyselfDerived
    //                                                                       =============
    public void xsmyselfDerive(String fn, SubQuery<VendorConstraintNameAutoQuxCB> sq, String al, DerivedReferrerOption op) {
        assertObjectNotNull("subQuery", sq);
        VendorConstraintNameAutoQuxCB cb = new VendorConstraintNameAutoQuxCB(); cb.xsetupForDerivedReferrer(this);
        lockCall(() -> sq.query(cb)); String pp = keepSpecifyMyselfDerived(cb.query()); String pk = "CONSTRAINT_NAME_AUTO_QUX_ID";
        registerSpecifyMyselfDerived(fn, cb.query(), pk, pk, pp, "myselfDerived", al, op);
    }
    public abstract String keepSpecifyMyselfDerived(VendorConstraintNameAutoQuxCQ sq);

    /**
     * Prepare for (Query)MyselfDerived (correlated sub-query).
     * @return The object to set up a function for myself table. (NotNull)
     */
    public HpQDRFunction<VendorConstraintNameAutoQuxCB> myselfDerived() {
        return xcreateQDRFunctionMyselfDerived(VendorConstraintNameAutoQuxCB.class);
    }
    @SuppressWarnings("unchecked")
    protected <CB extends ConditionBean> void xqderiveMyselfDerived(String fn, SubQuery<CB> sq, String rd, Object vl, DerivedReferrerOption op) {
        assertObjectNotNull("subQuery", sq);
        VendorConstraintNameAutoQuxCB cb = new VendorConstraintNameAutoQuxCB(); cb.xsetupForDerivedReferrer(this); sq.query((CB)cb);
        String pk = "CONSTRAINT_NAME_AUTO_QUX_ID";
        String sqpp = keepQueryMyselfDerived(cb.query()); // for saving query-value.
        String prpp = keepQueryMyselfDerivedParameter(vl);
        registerQueryMyselfDerived(fn, cb.query(), pk, pk, sqpp, "myselfDerived", rd, vl, prpp, op);
    }
    public abstract String keepQueryMyselfDerived(VendorConstraintNameAutoQuxCQ sq);
    public abstract String keepQueryMyselfDerivedParameter(Object vl);

    // ===================================================================================
    //                                                                        MyselfExists
    //                                                                        ============
    /**
     * Prepare for MyselfExists (correlated sub-query).
     * @param subCBLambda The implementation of sub-query. (NotNull)
     */
    public void myselfExists(SubQuery<VendorConstraintNameAutoQuxCB> subCBLambda) {
        assertObjectNotNull("subCBLambda", subCBLambda);
        VendorConstraintNameAutoQuxCB cb = new VendorConstraintNameAutoQuxCB(); cb.xsetupForMyselfExists(this);
        lockCall(() -> subCBLambda.query(cb)); String pp = keepMyselfExists(cb.query());
        registerMyselfExists(cb.query(), pp);
    }
    public abstract String keepMyselfExists(VendorConstraintNameAutoQuxCQ sq);

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
    protected VendorConstraintNameAutoQuxCB newMyCB() {
        return new VendorConstraintNameAutoQuxCB();
    }
    // very internal (for suppressing warn about 'Not Use Import')
    protected String xabUDT() { return Date.class.getName(); }
    protected String xabCQ() { return VendorConstraintNameAutoQuxCQ.class.getName(); }
    protected String xabLSO() { return LikeSearchOption.class.getName(); }
    protected String xabSSQS() { return HpSSQSetupper.class.getName(); }
    protected String xabSCP() { return SubQuery.class.getName(); }
}

package com.mssoftech.dbflute.cbean.cq.bs;

import java.util.Map;

import org.dbflute.cbean.*;
import org.dbflute.cbean.chelper.*;
import org.dbflute.cbean.coption.*;
import org.dbflute.cbean.cvalue.ConditionValue;
import org.dbflute.cbean.sqlclause.SqlClause;
import org.dbflute.exception.IllegalConditionBeanOperationException;
import com.mssoftech.dbflute.cbean.cq.ciq.*;
import com.mssoftech.dbflute.cbean.*;
import com.mssoftech.dbflute.cbean.cq.*;

/**
 * The base condition-query of VENDOR_SYMMETRIC.
 * @author DBFlute(AutoGenerator)
 */
public class BsVendorSymmetricCQ extends AbstractBsVendorSymmetricCQ {

    // ===================================================================================
    //                                                                           Attribute
    //                                                                           =========
    protected VendorSymmetricCIQ _inlineQuery;

    // ===================================================================================
    //                                                                         Constructor
    //                                                                         ===========
    public BsVendorSymmetricCQ(ConditionQuery referrerQuery, SqlClause sqlClause, String aliasName, int nestLevel) {
        super(referrerQuery, sqlClause, aliasName, nestLevel);
    }

    // ===================================================================================
    //                                                                 InlineView/OrClause
    //                                                                 ===================
    /**
     * Prepare InlineView query. <br>
     * {select ... from ... left outer join (select * from VENDOR_SYMMETRIC) where FOO = [value] ...}
     * <pre>
     * cb.query().queryMemberStatus().<span style="color: #CC4747">inline()</span>.setFoo...;
     * </pre>
     * @return The condition-query for InlineView query. (NotNull)
     */
    public VendorSymmetricCIQ inline() {
        if (_inlineQuery == null) { _inlineQuery = xcreateCIQ(); }
        _inlineQuery.xsetOnClause(false); return _inlineQuery;
    }

    protected VendorSymmetricCIQ xcreateCIQ() {
        VendorSymmetricCIQ ciq = xnewCIQ();
        ciq.xsetBaseCB(_baseCB);
        return ciq;
    }

    protected VendorSymmetricCIQ xnewCIQ() {
        return new VendorSymmetricCIQ(xgetReferrerQuery(), xgetSqlClause(), xgetAliasName(), xgetNestLevel(), this);
    }

    /**
     * Prepare OnClause query. <br>
     * {select ... from ... left outer join VENDOR_SYMMETRIC on ... and FOO = [value] ...}
     * <pre>
     * cb.query().queryMemberStatus().<span style="color: #CC4747">on()</span>.setFoo...;
     * </pre>
     * @return The condition-query for OnClause query. (NotNull)
     * @throws IllegalConditionBeanOperationException When this condition-query is base query.
     */
    public VendorSymmetricCIQ on() {
        if (isBaseQuery()) { throw new IllegalConditionBeanOperationException("OnClause for local table is unavailable!"); }
        VendorSymmetricCIQ inlineQuery = inline(); inlineQuery.xsetOnClause(true); return inlineQuery;
    }

    // ===================================================================================
    //                                                                               Query
    //                                                                               =====
    protected ConditionValue _vendorSymmetricId;
    public ConditionValue xdfgetVendorSymmetricId()
    { if (_vendorSymmetricId == null) { _vendorSymmetricId = nCV(); }
      return _vendorSymmetricId; }
    protected ConditionValue xgetCValueVendorSymmetricId() { return xdfgetVendorSymmetricId(); }

    /** 
     * Add order-by as ascend. <br>
     * VENDOR_SYMMETRIC_ID: {PK, NotNull, numeric(16)}
     * @return this. (NotNull)
     */
    public BsVendorSymmetricCQ addOrderBy_VendorSymmetricId_Asc() { regOBA("VENDOR_SYMMETRIC_ID"); return this; }

    /**
     * Add order-by as descend. <br>
     * VENDOR_SYMMETRIC_ID: {PK, NotNull, numeric(16)}
     * @return this. (NotNull)
     */
    public BsVendorSymmetricCQ addOrderBy_VendorSymmetricId_Desc() { regOBD("VENDOR_SYMMETRIC_ID"); return this; }

    protected ConditionValue _plainText;
    public ConditionValue xdfgetPlainText()
    { if (_plainText == null) { _plainText = nCV(); }
      return _plainText; }
    protected ConditionValue xgetCValuePlainText() { return xdfgetPlainText(); }

    /** 
     * Add order-by as ascend. <br>
     * PLAIN_TEXT: {nvarchar(100)}
     * @return this. (NotNull)
     */
    public BsVendorSymmetricCQ addOrderBy_PlainText_Asc() { regOBA("PLAIN_TEXT"); return this; }

    /**
     * Add order-by as descend. <br>
     * PLAIN_TEXT: {nvarchar(100)}
     * @return this. (NotNull)
     */
    public BsVendorSymmetricCQ addOrderBy_PlainText_Desc() { regOBD("PLAIN_TEXT"); return this; }

    protected ConditionValue _encryptedData;
    public ConditionValue xdfgetEncryptedData()
    { if (_encryptedData == null) { _encryptedData = nCV(); }
      return _encryptedData; }
    protected ConditionValue xgetCValueEncryptedData() { return xdfgetEncryptedData(); }

    /** 
     * Add order-by as ascend. <br>
     * ENCRYPTED_DATA: {varbinary(2147483647)}
     * @return this. (NotNull)
     */
    public BsVendorSymmetricCQ addOrderBy_EncryptedData_Asc() { regOBA("ENCRYPTED_DATA"); return this; }

    /**
     * Add order-by as descend. <br>
     * ENCRYPTED_DATA: {varbinary(2147483647)}
     * @return this. (NotNull)
     */
    public BsVendorSymmetricCQ addOrderBy_EncryptedData_Desc() { regOBD("ENCRYPTED_DATA"); return this; }

    // ===================================================================================
    //                                                             SpecifiedDerivedOrderBy
    //                                                             =======================
    /**
     * Add order-by for specified derived column as ascend.
     * <pre>
     * cb.specify().derivedPurchaseList().max(new SubQuery&lt;PurchaseCB&gt;() {
     *     public void query(PurchaseCB subCB) {
     *         subCB.specify().columnPurchaseDatetime();
     *     }
     * }, <span style="color: #CC4747">aliasName</span>);
     * <span style="color: #3F7E5E">// order by [alias-name] asc</span>
     * cb.<span style="color: #CC4747">addSpecifiedDerivedOrderBy_Asc</span>(<span style="color: #CC4747">aliasName</span>);
     * </pre>
     * @param aliasName The alias name specified at (Specify)DerivedReferrer. (NotNull)
     * @return this. (NotNull)
     */
    public BsVendorSymmetricCQ addSpecifiedDerivedOrderBy_Asc(String aliasName) { registerSpecifiedDerivedOrderBy_Asc(aliasName); return this; }

    /**
     * Add order-by for specified derived column as descend.
     * <pre>
     * cb.specify().derivedPurchaseList().max(new SubQuery&lt;PurchaseCB&gt;() {
     *     public void query(PurchaseCB subCB) {
     *         subCB.specify().columnPurchaseDatetime();
     *     }
     * }, <span style="color: #CC4747">aliasName</span>);
     * <span style="color: #3F7E5E">// order by [alias-name] desc</span>
     * cb.<span style="color: #CC4747">addSpecifiedDerivedOrderBy_Desc</span>(<span style="color: #CC4747">aliasName</span>);
     * </pre>
     * @param aliasName The alias name specified at (Specify)DerivedReferrer. (NotNull)
     * @return this. (NotNull)
     */
    public BsVendorSymmetricCQ addSpecifiedDerivedOrderBy_Desc(String aliasName) { registerSpecifiedDerivedOrderBy_Desc(aliasName); return this; }

    // ===================================================================================
    //                                                                         Union Query
    //                                                                         ===========
    public void reflectRelationOnUnionQuery(ConditionQuery bqs, ConditionQuery uqs) {
    }

    // ===================================================================================
    //                                                                       Foreign Query
    //                                                                       =============
    protected Map<String, Object> xfindFixedConditionDynamicParameterMap(String property) {
        return null;
    }

    // ===================================================================================
    //                                                                     ScalarCondition
    //                                                                     ===============
    public Map<String, VendorSymmetricCQ> xdfgetScalarCondition() { return xgetSQueMap("scalarCondition"); }
    public String keepScalarCondition(VendorSymmetricCQ sq) { return xkeepSQue("scalarCondition", sq); }

    // ===================================================================================
    //                                                                       MyselfDerived
    //                                                                       =============
    public Map<String, VendorSymmetricCQ> xdfgetSpecifyMyselfDerived() { return xgetSQueMap("specifyMyselfDerived"); }
    public String keepSpecifyMyselfDerived(VendorSymmetricCQ sq) { return xkeepSQue("specifyMyselfDerived", sq); }

    public Map<String, VendorSymmetricCQ> xdfgetQueryMyselfDerived() { return xgetSQueMap("queryMyselfDerived"); }
    public String keepQueryMyselfDerived(VendorSymmetricCQ sq) { return xkeepSQue("queryMyselfDerived", sq); }
    public Map<String, Object> xdfgetQueryMyselfDerivedParameter() { return xgetSQuePmMap("queryMyselfDerived"); }
    public String keepQueryMyselfDerivedParameter(Object pm) { return xkeepSQuePm("queryMyselfDerived", pm); }

    // ===================================================================================
    //                                                                        MyselfExists
    //                                                                        ============
    protected Map<String, VendorSymmetricCQ> _myselfExistsMap;
    public Map<String, VendorSymmetricCQ> xdfgetMyselfExists() { return xgetSQueMap("myselfExists"); }
    public String keepMyselfExists(VendorSymmetricCQ sq) { return xkeepSQue("myselfExists", sq); }

    // ===================================================================================
    //                                                                       MyselfInScope
    //                                                                       =============
    public Map<String, VendorSymmetricCQ> xdfgetMyselfInScope() { return xgetSQueMap("myselfInScope"); }
    public String keepMyselfInScope(VendorSymmetricCQ sq) { return xkeepSQue("myselfInScope", sq); }

    // ===================================================================================
    //                                                                       Very Internal
    //                                                                       =============
    // very internal (for suppressing warn about 'Not Use Import')
    protected String xCB() { return VendorSymmetricCB.class.getName(); }
    protected String xCQ() { return VendorSymmetricCQ.class.getName(); }
    protected String xCHp() { return HpQDRFunction.class.getName(); }
    protected String xCOp() { return ConditionOption.class.getName(); }
    protected String xMap() { return Map.class.getName(); }
}

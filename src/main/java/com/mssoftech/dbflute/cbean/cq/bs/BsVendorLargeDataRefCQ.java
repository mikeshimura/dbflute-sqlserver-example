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
 * The base condition-query of vendor_large_data_ref.
 * @author DBFlute(AutoGenerator)
 */
public class BsVendorLargeDataRefCQ extends AbstractBsVendorLargeDataRefCQ {

    // ===================================================================================
    //                                                                           Attribute
    //                                                                           =========
    protected VendorLargeDataRefCIQ _inlineQuery;

    // ===================================================================================
    //                                                                         Constructor
    //                                                                         ===========
    public BsVendorLargeDataRefCQ(ConditionQuery referrerQuery, SqlClause sqlClause, String aliasName, int nestLevel) {
        super(referrerQuery, sqlClause, aliasName, nestLevel);
    }

    // ===================================================================================
    //                                                                 InlineView/OrClause
    //                                                                 ===================
    /**
     * Prepare InlineView query. <br>
     * {select ... from ... left outer join (select * from vendor_large_data_ref) where FOO = [value] ...}
     * <pre>
     * cb.query().queryMemberStatus().<span style="color: #CC4747">inline()</span>.setFoo...;
     * </pre>
     * @return The condition-query for InlineView query. (NotNull)
     */
    public VendorLargeDataRefCIQ inline() {
        if (_inlineQuery == null) { _inlineQuery = xcreateCIQ(); }
        _inlineQuery.xsetOnClause(false); return _inlineQuery;
    }

    protected VendorLargeDataRefCIQ xcreateCIQ() {
        VendorLargeDataRefCIQ ciq = xnewCIQ();
        ciq.xsetBaseCB(_baseCB);
        return ciq;
    }

    protected VendorLargeDataRefCIQ xnewCIQ() {
        return new VendorLargeDataRefCIQ(xgetReferrerQuery(), xgetSqlClause(), xgetAliasName(), xgetNestLevel(), this);
    }

    /**
     * Prepare OnClause query. <br>
     * {select ... from ... left outer join vendor_large_data_ref on ... and FOO = [value] ...}
     * <pre>
     * cb.query().queryMemberStatus().<span style="color: #CC4747">on()</span>.setFoo...;
     * </pre>
     * @return The condition-query for OnClause query. (NotNull)
     * @throws IllegalConditionBeanOperationException When this condition-query is base query.
     */
    public VendorLargeDataRefCIQ on() {
        if (isBaseQuery()) { throw new IllegalConditionBeanOperationException("OnClause for local table is unavailable!"); }
        VendorLargeDataRefCIQ inlineQuery = inline(); inlineQuery.xsetOnClause(true); return inlineQuery;
    }

    // ===================================================================================
    //                                                                               Query
    //                                                                               =====
    protected ConditionValue _largeDataRefId;
    public ConditionValue xdfgetLargeDataRefId()
    { if (_largeDataRefId == null) { _largeDataRefId = nCV(); }
      return _largeDataRefId; }
    protected ConditionValue xgetCValueLargeDataRefId() { return xdfgetLargeDataRefId(); }

    /** 
     * Add order-by as ascend. <br>
     * LARGE_DATA_REF_ID: {PK, NotNull, BIGINT(19)}
     * @return this. (NotNull)
     */
    public BsVendorLargeDataRefCQ addOrderBy_LargeDataRefId_Asc() { regOBA("LARGE_DATA_REF_ID"); return this; }

    /**
     * Add order-by as descend. <br>
     * LARGE_DATA_REF_ID: {PK, NotNull, BIGINT(19)}
     * @return this. (NotNull)
     */
    public BsVendorLargeDataRefCQ addOrderBy_LargeDataRefId_Desc() { regOBD("LARGE_DATA_REF_ID"); return this; }

    protected ConditionValue _largeDataId;
    public ConditionValue xdfgetLargeDataId()
    { if (_largeDataId == null) { _largeDataId = nCV(); }
      return _largeDataId; }
    protected ConditionValue xgetCValueLargeDataId() { return xdfgetLargeDataId(); }

    /** 
     * Add order-by as ascend. <br>
     * LARGE_DATA_ID: {NotNull, BIGINT(19)}
     * @return this. (NotNull)
     */
    public BsVendorLargeDataRefCQ addOrderBy_LargeDataId_Asc() { regOBA("LARGE_DATA_ID"); return this; }

    /**
     * Add order-by as descend. <br>
     * LARGE_DATA_ID: {NotNull, BIGINT(19)}
     * @return this. (NotNull)
     */
    public BsVendorLargeDataRefCQ addOrderBy_LargeDataId_Desc() { regOBD("LARGE_DATA_ID"); return this; }

    protected ConditionValue _dateIndex;
    public ConditionValue xdfgetDateIndex()
    { if (_dateIndex == null) { _dateIndex = nCV(); }
      return _dateIndex; }
    protected ConditionValue xgetCValueDateIndex() { return xdfgetDateIndex(); }

    /** 
     * Add order-by as ascend. <br>
     * DATE_INDEX: {NotNull, DATE(10)}
     * @return this. (NotNull)
     */
    public BsVendorLargeDataRefCQ addOrderBy_DateIndex_Asc() { regOBA("DATE_INDEX"); return this; }

    /**
     * Add order-by as descend. <br>
     * DATE_INDEX: {NotNull, DATE(10)}
     * @return this. (NotNull)
     */
    public BsVendorLargeDataRefCQ addOrderBy_DateIndex_Desc() { regOBD("DATE_INDEX"); return this; }

    protected ConditionValue _dateNoIndex;
    public ConditionValue xdfgetDateNoIndex()
    { if (_dateNoIndex == null) { _dateNoIndex = nCV(); }
      return _dateNoIndex; }
    protected ConditionValue xgetCValueDateNoIndex() { return xdfgetDateNoIndex(); }

    /** 
     * Add order-by as ascend. <br>
     * DATE_NO_INDEX: {NotNull, DATE(10)}
     * @return this. (NotNull)
     */
    public BsVendorLargeDataRefCQ addOrderBy_DateNoIndex_Asc() { regOBA("DATE_NO_INDEX"); return this; }

    /**
     * Add order-by as descend. <br>
     * DATE_NO_INDEX: {NotNull, DATE(10)}
     * @return this. (NotNull)
     */
    public BsVendorLargeDataRefCQ addOrderBy_DateNoIndex_Desc() { regOBD("DATE_NO_INDEX"); return this; }

    protected ConditionValue _timestampIndex;
    public ConditionValue xdfgetTimestampIndex()
    { if (_timestampIndex == null) { _timestampIndex = nCV(); }
      return _timestampIndex; }
    protected ConditionValue xgetCValueTimestampIndex() { return xdfgetTimestampIndex(); }

    /** 
     * Add order-by as ascend. <br>
     * TIMESTAMP_INDEX: {NotNull, DATETIME(19)}
     * @return this. (NotNull)
     */
    public BsVendorLargeDataRefCQ addOrderBy_TimestampIndex_Asc() { regOBA("TIMESTAMP_INDEX"); return this; }

    /**
     * Add order-by as descend. <br>
     * TIMESTAMP_INDEX: {NotNull, DATETIME(19)}
     * @return this. (NotNull)
     */
    public BsVendorLargeDataRefCQ addOrderBy_TimestampIndex_Desc() { regOBD("TIMESTAMP_INDEX"); return this; }

    protected ConditionValue _timestampNoIndex;
    public ConditionValue xdfgetTimestampNoIndex()
    { if (_timestampNoIndex == null) { _timestampNoIndex = nCV(); }
      return _timestampNoIndex; }
    protected ConditionValue xgetCValueTimestampNoIndex() { return xdfgetTimestampNoIndex(); }

    /** 
     * Add order-by as ascend. <br>
     * TIMESTAMP_NO_INDEX: {NotNull, DATETIME(19)}
     * @return this. (NotNull)
     */
    public BsVendorLargeDataRefCQ addOrderBy_TimestampNoIndex_Asc() { regOBA("TIMESTAMP_NO_INDEX"); return this; }

    /**
     * Add order-by as descend. <br>
     * TIMESTAMP_NO_INDEX: {NotNull, DATETIME(19)}
     * @return this. (NotNull)
     */
    public BsVendorLargeDataRefCQ addOrderBy_TimestampNoIndex_Desc() { regOBD("TIMESTAMP_NO_INDEX"); return this; }

    protected ConditionValue _nullableDecimalIndex;
    public ConditionValue xdfgetNullableDecimalIndex()
    { if (_nullableDecimalIndex == null) { _nullableDecimalIndex = nCV(); }
      return _nullableDecimalIndex; }
    protected ConditionValue xgetCValueNullableDecimalIndex() { return xdfgetNullableDecimalIndex(); }

    /** 
     * Add order-by as ascend. <br>
     * NULLABLE_DECIMAL_INDEX: {DECIMAL(12, 3)}
     * @return this. (NotNull)
     */
    public BsVendorLargeDataRefCQ addOrderBy_NullableDecimalIndex_Asc() { regOBA("NULLABLE_DECIMAL_INDEX"); return this; }

    /**
     * Add order-by as descend. <br>
     * NULLABLE_DECIMAL_INDEX: {DECIMAL(12, 3)}
     * @return this. (NotNull)
     */
    public BsVendorLargeDataRefCQ addOrderBy_NullableDecimalIndex_Desc() { regOBD("NULLABLE_DECIMAL_INDEX"); return this; }

    protected ConditionValue _nullableDecimalNoIndex;
    public ConditionValue xdfgetNullableDecimalNoIndex()
    { if (_nullableDecimalNoIndex == null) { _nullableDecimalNoIndex = nCV(); }
      return _nullableDecimalNoIndex; }
    protected ConditionValue xgetCValueNullableDecimalNoIndex() { return xdfgetNullableDecimalNoIndex(); }

    /** 
     * Add order-by as ascend. <br>
     * NULLABLE_DECIMAL_NO_INDEX: {DECIMAL(12, 3)}
     * @return this. (NotNull)
     */
    public BsVendorLargeDataRefCQ addOrderBy_NullableDecimalNoIndex_Asc() { regOBA("NULLABLE_DECIMAL_NO_INDEX"); return this; }

    /**
     * Add order-by as descend. <br>
     * NULLABLE_DECIMAL_NO_INDEX: {DECIMAL(12, 3)}
     * @return this. (NotNull)
     */
    public BsVendorLargeDataRefCQ addOrderBy_NullableDecimalNoIndex_Desc() { regOBD("NULLABLE_DECIMAL_NO_INDEX"); return this; }

    protected ConditionValue _selfParentId;
    public ConditionValue xdfgetSelfParentId()
    { if (_selfParentId == null) { _selfParentId = nCV(); }
      return _selfParentId; }
    protected ConditionValue xgetCValueSelfParentId() { return xdfgetSelfParentId(); }

    /** 
     * Add order-by as ascend. <br>
     * SELF_PARENT_ID: {BIGINT(19)}
     * @return this. (NotNull)
     */
    public BsVendorLargeDataRefCQ addOrderBy_SelfParentId_Asc() { regOBA("SELF_PARENT_ID"); return this; }

    /**
     * Add order-by as descend. <br>
     * SELF_PARENT_ID: {BIGINT(19)}
     * @return this. (NotNull)
     */
    public BsVendorLargeDataRefCQ addOrderBy_SelfParentId_Desc() { regOBD("SELF_PARENT_ID"); return this; }

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
    public BsVendorLargeDataRefCQ addSpecifiedDerivedOrderBy_Asc(String aliasName) { registerSpecifiedDerivedOrderBy_Asc(aliasName); return this; }

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
    public BsVendorLargeDataRefCQ addSpecifiedDerivedOrderBy_Desc(String aliasName) { registerSpecifiedDerivedOrderBy_Desc(aliasName); return this; }

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
    public Map<String, VendorLargeDataRefCQ> xdfgetScalarCondition() { return xgetSQueMap("scalarCondition"); }
    public String keepScalarCondition(VendorLargeDataRefCQ sq) { return xkeepSQue("scalarCondition", sq); }

    // ===================================================================================
    //                                                                       MyselfDerived
    //                                                                       =============
    public Map<String, VendorLargeDataRefCQ> xdfgetSpecifyMyselfDerived() { return xgetSQueMap("specifyMyselfDerived"); }
    public String keepSpecifyMyselfDerived(VendorLargeDataRefCQ sq) { return xkeepSQue("specifyMyselfDerived", sq); }

    public Map<String, VendorLargeDataRefCQ> xdfgetQueryMyselfDerived() { return xgetSQueMap("queryMyselfDerived"); }
    public String keepQueryMyselfDerived(VendorLargeDataRefCQ sq) { return xkeepSQue("queryMyselfDerived", sq); }
    public Map<String, Object> xdfgetQueryMyselfDerivedParameter() { return xgetSQuePmMap("queryMyselfDerived"); }
    public String keepQueryMyselfDerivedParameter(Object pm) { return xkeepSQuePm("queryMyselfDerived", pm); }

    // ===================================================================================
    //                                                                        MyselfExists
    //                                                                        ============
    protected Map<String, VendorLargeDataRefCQ> _myselfExistsMap;
    public Map<String, VendorLargeDataRefCQ> xdfgetMyselfExists() { return xgetSQueMap("myselfExists"); }
    public String keepMyselfExists(VendorLargeDataRefCQ sq) { return xkeepSQue("myselfExists", sq); }

    // ===================================================================================
    //                                                                       MyselfInScope
    //                                                                       =============
    public Map<String, VendorLargeDataRefCQ> xdfgetMyselfInScope() { return xgetSQueMap("myselfInScope"); }
    public String keepMyselfInScope(VendorLargeDataRefCQ sq) { return xkeepSQue("myselfInScope", sq); }

    // ===================================================================================
    //                                                                       Very Internal
    //                                                                       =============
    // very internal (for suppressing warn about 'Not Use Import')
    protected String xCB() { return VendorLargeDataRefCB.class.getName(); }
    protected String xCQ() { return VendorLargeDataRefCQ.class.getName(); }
    protected String xCHp() { return HpQDRFunction.class.getName(); }
    protected String xCOp() { return ConditionOption.class.getName(); }
    protected String xMap() { return Map.class.getName(); }
}

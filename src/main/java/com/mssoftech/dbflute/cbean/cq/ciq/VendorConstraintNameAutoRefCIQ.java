package com.mssoftech.dbflute.cbean.cq.ciq;

import java.util.Map;
import org.dbflute.cbean.*;
import org.dbflute.cbean.ckey.*;
import org.dbflute.cbean.coption.ConditionOption;
import org.dbflute.cbean.cvalue.ConditionValue;
import org.dbflute.cbean.sqlclause.SqlClause;
import org.dbflute.exception.IllegalConditionBeanOperationException;
import com.mssoftech.dbflute.cbean.*;
import com.mssoftech.dbflute.cbean.cq.bs.*;
import com.mssoftech.dbflute.cbean.cq.*;

/**
 * The condition-query for in-line of vendor_constraint_name_auto_ref.
 * @author DBFlute(AutoGenerator)
 */
public class VendorConstraintNameAutoRefCIQ extends AbstractBsVendorConstraintNameAutoRefCQ {

    // ===================================================================================
    //                                                                           Attribute
    //                                                                           =========
    protected BsVendorConstraintNameAutoRefCQ _myCQ;

    // ===================================================================================
    //                                                                         Constructor
    //                                                                         ===========
    public VendorConstraintNameAutoRefCIQ(ConditionQuery referrerQuery, SqlClause sqlClause
                        , String aliasName, int nestLevel, BsVendorConstraintNameAutoRefCQ myCQ) {
        super(referrerQuery, sqlClause, aliasName, nestLevel);
        _myCQ = myCQ;
        _foreignPropertyName = _myCQ.xgetForeignPropertyName(); // accept foreign property name
        _relationPath = _myCQ.xgetRelationPath(); // accept relation path
        _inline = true;
    }

    // ===================================================================================
    //                                                             Override about Register
    //                                                             =======================
    protected void reflectRelationOnUnionQuery(ConditionQuery bq, ConditionQuery uq)
    { throw new IllegalConditionBeanOperationException("InlineView cannot use Union: " + bq + " : " + uq); }

    @Override
    protected void setupConditionValueAndRegisterWhereClause(ConditionKey k, Object v, ConditionValue cv, String col)
    { regIQ(k, v, cv, col); }

    @Override
    protected void setupConditionValueAndRegisterWhereClause(ConditionKey k, Object v, ConditionValue cv, String col, ConditionOption op)
    { regIQ(k, v, cv, col, op); }

    @Override
    protected void registerWhereClause(String wc)
    { registerInlineWhereClause(wc); }

    @Override
    protected boolean isInScopeRelationSuppressLocalAliasName() {
        if (_onClause) { throw new IllegalConditionBeanOperationException("InScopeRelation on OnClause is unsupported."); }
        return true;
    }

    // ===================================================================================
    //                                                                Override about Query
    //                                                                ====================
    protected ConditionValue xgetCValueConstraintNameAutoRefId() { return _myCQ.xdfgetConstraintNameAutoRefId(); }
    protected ConditionValue xgetCValueConstraintNameAutoFooId() { return _myCQ.xdfgetConstraintNameAutoFooId(); }
    protected ConditionValue xgetCValueConstraintNameAutoBarId() { return _myCQ.xdfgetConstraintNameAutoBarId(); }
    protected ConditionValue xgetCValueConstraintNameAutoQuxId() { return _myCQ.xdfgetConstraintNameAutoQuxId(); }
    protected ConditionValue xgetCValueConstraintNameAutoCorgeId() { return _myCQ.xdfgetConstraintNameAutoCorgeId(); }
    protected ConditionValue xgetCValueConstraintNameAutoUnique() { return _myCQ.xdfgetConstraintNameAutoUnique(); }
    protected Map<String, Object> xfindFixedConditionDynamicParameterMap(String pp) { return null; }
    public String keepScalarCondition(VendorConstraintNameAutoRefCQ sq)
    { throwIICBOE("ScalarCondition"); return null; }
    public String keepSpecifyMyselfDerived(VendorConstraintNameAutoRefCQ sq)
    { throwIICBOE("(Specify)MyselfDerived"); return null;}
    public String keepQueryMyselfDerived(VendorConstraintNameAutoRefCQ sq)
    { throwIICBOE("(Query)MyselfDerived"); return null;}
    public String keepQueryMyselfDerivedParameter(Object vl)
    { throwIICBOE("(Query)MyselfDerived"); return null;}
    public String keepMyselfExists(VendorConstraintNameAutoRefCQ sq)
    { throwIICBOE("MyselfExists"); return null;}

    protected void throwIICBOE(String name)
    { throw new IllegalConditionBeanOperationException(name + " at InlineView is unsupported."); }

    // ===================================================================================
    //                                                                       Very Internal
    //                                                                       =============
    // very internal (for suppressing warn about 'Not Use Import')
    protected String xinCB() { return VendorConstraintNameAutoRefCB.class.getName(); }
    protected String xinCQ() { return VendorConstraintNameAutoRefCQ.class.getName(); }
}

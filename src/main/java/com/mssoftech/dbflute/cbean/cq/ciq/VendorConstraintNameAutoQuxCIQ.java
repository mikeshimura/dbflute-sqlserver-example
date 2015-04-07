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
 * The condition-query for in-line of vendor_constraint_name_auto_qux.
 * @author DBFlute(AutoGenerator)
 */
public class VendorConstraintNameAutoQuxCIQ extends AbstractBsVendorConstraintNameAutoQuxCQ {

    // ===================================================================================
    //                                                                           Attribute
    //                                                                           =========
    protected BsVendorConstraintNameAutoQuxCQ _myCQ;

    // ===================================================================================
    //                                                                         Constructor
    //                                                                         ===========
    public VendorConstraintNameAutoQuxCIQ(ConditionQuery referrerQuery, SqlClause sqlClause
                        , String aliasName, int nestLevel, BsVendorConstraintNameAutoQuxCQ myCQ) {
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
    protected ConditionValue xgetCValueConstraintNameAutoQuxId() { return _myCQ.xdfgetConstraintNameAutoQuxId(); }
    public String keepConstraintNameAutoQuxId_ExistsReferrer_VendorConstraintNameAutoRefList(VendorConstraintNameAutoRefCQ sq)
    { throwIICBOE("ExistsReferrer"); return null; }
    public String keepConstraintNameAutoQuxId_NotExistsReferrer_VendorConstraintNameAutoRefList(VendorConstraintNameAutoRefCQ sq)
    { throwIICBOE("NotExistsReferrer"); return null; }
    public String keepConstraintNameAutoQuxId_SpecifyDerivedReferrer_VendorConstraintNameAutoRefList(VendorConstraintNameAutoRefCQ sq)
    { throwIICBOE("(Specify)DerivedReferrer"); return null; }
    public String keepConstraintNameAutoQuxId_QueryDerivedReferrer_VendorConstraintNameAutoRefList(VendorConstraintNameAutoRefCQ sq)
    { throwIICBOE("(Query)DerivedReferrer"); return null; }
    public String keepConstraintNameAutoQuxId_QueryDerivedReferrer_VendorConstraintNameAutoRefListParameter(Object vl)
    { throwIICBOE("(Query)DerivedReferrer"); return null; }
    protected ConditionValue xgetCValueConstraintNameAutoQuxName() { return _myCQ.xdfgetConstraintNameAutoQuxName(); }
    protected Map<String, Object> xfindFixedConditionDynamicParameterMap(String pp) { return null; }
    public String keepScalarCondition(VendorConstraintNameAutoQuxCQ sq)
    { throwIICBOE("ScalarCondition"); return null; }
    public String keepSpecifyMyselfDerived(VendorConstraintNameAutoQuxCQ sq)
    { throwIICBOE("(Specify)MyselfDerived"); return null;}
    public String keepQueryMyselfDerived(VendorConstraintNameAutoQuxCQ sq)
    { throwIICBOE("(Query)MyselfDerived"); return null;}
    public String keepQueryMyselfDerivedParameter(Object vl)
    { throwIICBOE("(Query)MyselfDerived"); return null;}
    public String keepMyselfExists(VendorConstraintNameAutoQuxCQ sq)
    { throwIICBOE("MyselfExists"); return null;}

    protected void throwIICBOE(String name)
    { throw new IllegalConditionBeanOperationException(name + " at InlineView is unsupported."); }

    // ===================================================================================
    //                                                                       Very Internal
    //                                                                       =============
    // very internal (for suppressing warn about 'Not Use Import')
    protected String xinCB() { return VendorConstraintNameAutoQuxCB.class.getName(); }
    protected String xinCQ() { return VendorConstraintNameAutoQuxCQ.class.getName(); }
}

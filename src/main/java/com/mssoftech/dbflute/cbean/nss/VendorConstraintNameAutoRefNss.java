package com.mssoftech.dbflute.cbean.nss;

import com.mssoftech.dbflute.cbean.cq.VendorConstraintNameAutoRefCQ;

/**
 * The nest select set-upper of vendor_constraint_name_auto_ref.
 * @author DBFlute(AutoGenerator)
 */
public class VendorConstraintNameAutoRefNss {

    // ===================================================================================
    //                                                                           Attribute
    //                                                                           =========
    protected final VendorConstraintNameAutoRefCQ _query;
    public VendorConstraintNameAutoRefNss(VendorConstraintNameAutoRefCQ query) { _query = query; }
    public boolean hasConditionQuery() { return _query != null; }

    // ===================================================================================
    //                                                                     Nested Relation
    //                                                                     ===============
    /**
     * With nested relation columns to select clause. <br>
     * vendor_constraint_name_auto_bar by my CONSTRAINT_NAME_AUTO_BAR_ID, named 'vendorConstraintNameAutoBar'.
     */
    public void withVendorConstraintNameAutoBar() {
        _query.xdoNss(() -> _query.queryVendorConstraintNameAutoBar());
    }
    /**
     * With nested relation columns to select clause. <br>
     * vendor_constraint_name_auto_foo by my CONSTRAINT_NAME_AUTO_FOO_ID, named 'vendorConstraintNameAutoFoo'.
     */
    public void withVendorConstraintNameAutoFoo() {
        _query.xdoNss(() -> _query.queryVendorConstraintNameAutoFoo());
    }
    /**
     * With nested relation columns to select clause. <br>
     * vendor_constraint_name_auto_qux by my CONSTRAINT_NAME_AUTO_QUX_ID, named 'vendorConstraintNameAutoQux'.
     */
    public void withVendorConstraintNameAutoQux() {
        _query.xdoNss(() -> _query.queryVendorConstraintNameAutoQux());
    }
}

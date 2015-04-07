package com.mssoftech.dbflute.bsentity;

import java.util.List;
import java.util.ArrayList;

import org.dbflute.dbmeta.DBMeta;
import org.dbflute.dbmeta.AbstractEntity;
import org.dbflute.dbmeta.accessory.DomainEntity;
import com.mssoftech.dbflute.allcommon.DBMetaInstanceHandler;
import com.mssoftech.dbflute.exentity.*;

/**
 * The entity of vendor_constraint_name_auto_foo as TABLE. <br>
 * <pre>
 * [primary-key]
 *     CONSTRAINT_NAME_AUTO_FOO_ID
 * 
 * [column]
 *     CONSTRAINT_NAME_AUTO_FOO_ID, CONSTRAINT_NAME_AUTO_FOO_NAME
 * 
 * [sequence]
 *     
 * 
 * [identity]
 *     
 * 
 * [version-no]
 *     
 * 
 * [foreign table]
 *     
 * 
 * [referrer table]
 *     vendor_constraint_name_auto_ref
 * 
 * [foreign property]
 *     
 * 
 * [referrer property]
 *     vendorConstraintNameAutoRefList
 * 
 * [get/set template]
 * /= = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
 * Long constraintNameAutoFooId = entity.getConstraintNameAutoFooId();
 * String constraintNameAutoFooName = entity.getConstraintNameAutoFooName();
 * entity.setConstraintNameAutoFooId(constraintNameAutoFooId);
 * entity.setConstraintNameAutoFooName(constraintNameAutoFooName);
 * = = = = = = = = = =/
 * </pre>
 * @author DBFlute(AutoGenerator)
 */
public abstract class BsVendorConstraintNameAutoFoo extends AbstractEntity implements DomainEntity {

    // ===================================================================================
    //                                                                          Definition
    //                                                                          ==========
    /** The serial version UID for object serialization. (Default) */
    private static final long serialVersionUID = 1L;

    // ===================================================================================
    //                                                                           Attribute
    //                                                                           =========
    /** CONSTRAINT_NAME_AUTO_FOO_ID: {PK, NotNull, DECIMAL(16)} */
    protected Long _constraintNameAutoFooId;

    /** CONSTRAINT_NAME_AUTO_FOO_NAME: {UQ, NotNull, VARCHAR(50)} */
    protected String _constraintNameAutoFooName;

    // ===================================================================================
    //                                                                             DB Meta
    //                                                                             =======
    /** {@inheritDoc} */
    public DBMeta asDBMeta() {
        return DBMetaInstanceHandler.findDBMeta(asTableDbName());
    }

    /** {@inheritDoc} */
    public String asTableDbName() {
        return "vendor_constraint_name_auto_foo";
    }

    // ===================================================================================
    //                                                                        Key Handling
    //                                                                        ============
    /** {@inheritDoc} */
    public boolean hasPrimaryKeyValue() {
        if (_constraintNameAutoFooId == null) { return false; }
        return true;
    }

    /**
     * To be unique by the unique column. <br>
     * You can update the entity by the key when entity update (NOT batch update).
     * @param constraintNameAutoFooName : UQ, NotNull, VARCHAR(50). (NotNull)
     */
    public void uniqueBy(String constraintNameAutoFooName) {
        __uniqueDrivenProperties.clear();
        __uniqueDrivenProperties.addPropertyName("constraintNameAutoFooName");
        setConstraintNameAutoFooName(constraintNameAutoFooName);
    }

    // ===================================================================================
    //                                                                    Foreign Property
    //                                                                    ================
    // ===================================================================================
    //                                                                   Referrer Property
    //                                                                   =================
    /** vendor_constraint_name_auto_ref by CONSTRAINT_NAME_AUTO_FOO_ID, named 'vendorConstraintNameAutoRefList'. */
    protected List<VendorConstraintNameAutoRef> _vendorConstraintNameAutoRefList;

    /**
     * [get] vendor_constraint_name_auto_ref by CONSTRAINT_NAME_AUTO_FOO_ID, named 'vendorConstraintNameAutoRefList'.
     * @return The entity list of referrer property 'vendorConstraintNameAutoRefList'. (NotNull: even if no loading, returns empty list)
     */
    public List<VendorConstraintNameAutoRef> getVendorConstraintNameAutoRefList() {
        if (_vendorConstraintNameAutoRefList == null) { _vendorConstraintNameAutoRefList = newReferrerList(); }
        return _vendorConstraintNameAutoRefList;
    }

    /**
     * [set] vendor_constraint_name_auto_ref by CONSTRAINT_NAME_AUTO_FOO_ID, named 'vendorConstraintNameAutoRefList'.
     * @param vendorConstraintNameAutoRefList The entity list of referrer property 'vendorConstraintNameAutoRefList'. (NullAllowed)
     */
    public void setVendorConstraintNameAutoRefList(List<VendorConstraintNameAutoRef> vendorConstraintNameAutoRefList) {
        _vendorConstraintNameAutoRefList = vendorConstraintNameAutoRefList;
    }

    protected <ELEMENT> List<ELEMENT> newReferrerList() {
        return new ArrayList<ELEMENT>();
    }

    // ===================================================================================
    //                                                                      Basic Override
    //                                                                      ==============
    @Override
    protected boolean doEquals(Object obj) {
        if (obj instanceof BsVendorConstraintNameAutoFoo) {
            BsVendorConstraintNameAutoFoo other = (BsVendorConstraintNameAutoFoo)obj;
            if (!xSV(_constraintNameAutoFooId, other._constraintNameAutoFooId)) { return false; }
            return true;
        } else {
            return false;
        }
    }

    @Override
    protected int doHashCode(int initial) {
        int hs = initial;
        hs = xCH(hs, asTableDbName());
        hs = xCH(hs, _constraintNameAutoFooId);
        return hs;
    }

    @Override
    protected String doBuildStringWithRelation(String li) {
        StringBuilder sb = new StringBuilder();
        if (_vendorConstraintNameAutoRefList != null) { for (VendorConstraintNameAutoRef et : _vendorConstraintNameAutoRefList)
        { if (et != null) { sb.append(li).append(xbRDS(et, "vendorConstraintNameAutoRefList")); } } }
        return sb.toString();
    }

    @Override
    protected String doBuildColumnString(String dm) {
        StringBuilder sb = new StringBuilder();
        sb.append(dm).append(xfND(_constraintNameAutoFooId));
        sb.append(dm).append(xfND(_constraintNameAutoFooName));
        if (sb.length() > dm.length()) {
            sb.delete(0, dm.length());
        }
        sb.insert(0, "{").append("}");
        return sb.toString();
    }

    @Override
    protected String doBuildRelationString(String dm) {
        StringBuilder sb = new StringBuilder();
        if (_vendorConstraintNameAutoRefList != null && !_vendorConstraintNameAutoRefList.isEmpty())
        { sb.append(dm).append("vendorConstraintNameAutoRefList"); }
        if (sb.length() > dm.length()) {
            sb.delete(0, dm.length()).insert(0, "(").append(")");
        }
        return sb.toString();
    }

    @Override
    public VendorConstraintNameAutoFoo clone() {
        return (VendorConstraintNameAutoFoo)super.clone();
    }

    // ===================================================================================
    //                                                                            Accessor
    //                                                                            ========
    /**
     * [get] CONSTRAINT_NAME_AUTO_FOO_ID: {PK, NotNull, DECIMAL(16)} <br>
     * @return The value of the column 'CONSTRAINT_NAME_AUTO_FOO_ID'. (basically NotNull if selected: for the constraint)
     */
    public Long getConstraintNameAutoFooId() {
        checkSpecifiedProperty("constraintNameAutoFooId");
        return _constraintNameAutoFooId;
    }

    /**
     * [set] CONSTRAINT_NAME_AUTO_FOO_ID: {PK, NotNull, DECIMAL(16)} <br>
     * @param constraintNameAutoFooId The value of the column 'CONSTRAINT_NAME_AUTO_FOO_ID'. (basically NotNull if update: for the constraint)
     */
    public void setConstraintNameAutoFooId(Long constraintNameAutoFooId) {
        registerModifiedProperty("constraintNameAutoFooId");
        _constraintNameAutoFooId = constraintNameAutoFooId;
    }

    /**
     * [get] CONSTRAINT_NAME_AUTO_FOO_NAME: {UQ, NotNull, VARCHAR(50)} <br>
     * @return The value of the column 'CONSTRAINT_NAME_AUTO_FOO_NAME'. (basically NotNull if selected: for the constraint)
     */
    public String getConstraintNameAutoFooName() {
        checkSpecifiedProperty("constraintNameAutoFooName");
        return _constraintNameAutoFooName;
    }

    /**
     * [set] CONSTRAINT_NAME_AUTO_FOO_NAME: {UQ, NotNull, VARCHAR(50)} <br>
     * @param constraintNameAutoFooName The value of the column 'CONSTRAINT_NAME_AUTO_FOO_NAME'. (basically NotNull if update: for the constraint)
     */
    public void setConstraintNameAutoFooName(String constraintNameAutoFooName) {
        registerModifiedProperty("constraintNameAutoFooName");
        _constraintNameAutoFooName = constraintNameAutoFooName;
    }
}

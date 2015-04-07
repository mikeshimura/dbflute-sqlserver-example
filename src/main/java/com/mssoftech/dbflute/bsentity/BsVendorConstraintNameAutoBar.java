package com.mssoftech.dbflute.bsentity;

import java.util.List;
import java.util.ArrayList;

import org.dbflute.dbmeta.DBMeta;
import org.dbflute.dbmeta.AbstractEntity;
import org.dbflute.dbmeta.accessory.DomainEntity;
import com.mssoftech.dbflute.allcommon.DBMetaInstanceHandler;
import com.mssoftech.dbflute.exentity.*;

/**
 * The entity of vendor_constraint_name_auto_bar as TABLE. <br>
 * <pre>
 * [primary-key]
 *     CONSTRAINT_NAME_AUTO_BAR_ID
 * 
 * [column]
 *     CONSTRAINT_NAME_AUTO_BAR_ID, CONSTRAINT_NAME_AUTO_BAR_NAME
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
 * Long constraintNameAutoBarId = entity.getConstraintNameAutoBarId();
 * String constraintNameAutoBarName = entity.getConstraintNameAutoBarName();
 * entity.setConstraintNameAutoBarId(constraintNameAutoBarId);
 * entity.setConstraintNameAutoBarName(constraintNameAutoBarName);
 * = = = = = = = = = =/
 * </pre>
 * @author DBFlute(AutoGenerator)
 */
public abstract class BsVendorConstraintNameAutoBar extends AbstractEntity implements DomainEntity {

    // ===================================================================================
    //                                                                          Definition
    //                                                                          ==========
    /** The serial version UID for object serialization. (Default) */
    private static final long serialVersionUID = 1L;

    // ===================================================================================
    //                                                                           Attribute
    //                                                                           =========
    /** CONSTRAINT_NAME_AUTO_BAR_ID: {PK, NotNull, DECIMAL(16)} */
    protected Long _constraintNameAutoBarId;

    /** CONSTRAINT_NAME_AUTO_BAR_NAME: {UQ, NotNull, VARCHAR(50)} */
    protected String _constraintNameAutoBarName;

    // ===================================================================================
    //                                                                             DB Meta
    //                                                                             =======
    /** {@inheritDoc} */
    public DBMeta asDBMeta() {
        return DBMetaInstanceHandler.findDBMeta(asTableDbName());
    }

    /** {@inheritDoc} */
    public String asTableDbName() {
        return "vendor_constraint_name_auto_bar";
    }

    // ===================================================================================
    //                                                                        Key Handling
    //                                                                        ============
    /** {@inheritDoc} */
    public boolean hasPrimaryKeyValue() {
        if (_constraintNameAutoBarId == null) { return false; }
        return true;
    }

    /**
     * To be unique by the unique column. <br>
     * You can update the entity by the key when entity update (NOT batch update).
     * @param constraintNameAutoBarName : UQ, NotNull, VARCHAR(50). (NotNull)
     */
    public void uniqueBy(String constraintNameAutoBarName) {
        __uniqueDrivenProperties.clear();
        __uniqueDrivenProperties.addPropertyName("constraintNameAutoBarName");
        setConstraintNameAutoBarName(constraintNameAutoBarName);
    }

    // ===================================================================================
    //                                                                    Foreign Property
    //                                                                    ================
    // ===================================================================================
    //                                                                   Referrer Property
    //                                                                   =================
    /** vendor_constraint_name_auto_ref by CONSTRAINT_NAME_AUTO_BAR_ID, named 'vendorConstraintNameAutoRefList'. */
    protected List<VendorConstraintNameAutoRef> _vendorConstraintNameAutoRefList;

    /**
     * [get] vendor_constraint_name_auto_ref by CONSTRAINT_NAME_AUTO_BAR_ID, named 'vendorConstraintNameAutoRefList'.
     * @return The entity list of referrer property 'vendorConstraintNameAutoRefList'. (NotNull: even if no loading, returns empty list)
     */
    public List<VendorConstraintNameAutoRef> getVendorConstraintNameAutoRefList() {
        if (_vendorConstraintNameAutoRefList == null) { _vendorConstraintNameAutoRefList = newReferrerList(); }
        return _vendorConstraintNameAutoRefList;
    }

    /**
     * [set] vendor_constraint_name_auto_ref by CONSTRAINT_NAME_AUTO_BAR_ID, named 'vendorConstraintNameAutoRefList'.
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
        if (obj instanceof BsVendorConstraintNameAutoBar) {
            BsVendorConstraintNameAutoBar other = (BsVendorConstraintNameAutoBar)obj;
            if (!xSV(_constraintNameAutoBarId, other._constraintNameAutoBarId)) { return false; }
            return true;
        } else {
            return false;
        }
    }

    @Override
    protected int doHashCode(int initial) {
        int hs = initial;
        hs = xCH(hs, asTableDbName());
        hs = xCH(hs, _constraintNameAutoBarId);
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
        sb.append(dm).append(xfND(_constraintNameAutoBarId));
        sb.append(dm).append(xfND(_constraintNameAutoBarName));
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
    public VendorConstraintNameAutoBar clone() {
        return (VendorConstraintNameAutoBar)super.clone();
    }

    // ===================================================================================
    //                                                                            Accessor
    //                                                                            ========
    /**
     * [get] CONSTRAINT_NAME_AUTO_BAR_ID: {PK, NotNull, DECIMAL(16)} <br>
     * @return The value of the column 'CONSTRAINT_NAME_AUTO_BAR_ID'. (basically NotNull if selected: for the constraint)
     */
    public Long getConstraintNameAutoBarId() {
        checkSpecifiedProperty("constraintNameAutoBarId");
        return _constraintNameAutoBarId;
    }

    /**
     * [set] CONSTRAINT_NAME_AUTO_BAR_ID: {PK, NotNull, DECIMAL(16)} <br>
     * @param constraintNameAutoBarId The value of the column 'CONSTRAINT_NAME_AUTO_BAR_ID'. (basically NotNull if update: for the constraint)
     */
    public void setConstraintNameAutoBarId(Long constraintNameAutoBarId) {
        registerModifiedProperty("constraintNameAutoBarId");
        _constraintNameAutoBarId = constraintNameAutoBarId;
    }

    /**
     * [get] CONSTRAINT_NAME_AUTO_BAR_NAME: {UQ, NotNull, VARCHAR(50)} <br>
     * @return The value of the column 'CONSTRAINT_NAME_AUTO_BAR_NAME'. (basically NotNull if selected: for the constraint)
     */
    public String getConstraintNameAutoBarName() {
        checkSpecifiedProperty("constraintNameAutoBarName");
        return _constraintNameAutoBarName;
    }

    /**
     * [set] CONSTRAINT_NAME_AUTO_BAR_NAME: {UQ, NotNull, VARCHAR(50)} <br>
     * @param constraintNameAutoBarName The value of the column 'CONSTRAINT_NAME_AUTO_BAR_NAME'. (basically NotNull if update: for the constraint)
     */
    public void setConstraintNameAutoBarName(String constraintNameAutoBarName) {
        registerModifiedProperty("constraintNameAutoBarName");
        _constraintNameAutoBarName = constraintNameAutoBarName;
    }
}

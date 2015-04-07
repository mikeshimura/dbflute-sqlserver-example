package com.mssoftech.dbflute.bsentity;

import java.util.List;
import java.util.ArrayList;

import org.dbflute.dbmeta.DBMeta;
import org.dbflute.dbmeta.AbstractEntity;
import org.dbflute.dbmeta.accessory.DomainEntity;
import com.mssoftech.dbflute.allcommon.DBMetaInstanceHandler;
import com.mssoftech.dbflute.exentity.*;

/**
 * The entity of vendor_large_data as TABLE. <br>
 * <pre>
 * [primary-key]
 *     LARGE_DATA_ID
 * 
 * [column]
 *     LARGE_DATA_ID, STRING_INDEX, STRING_NO_INDEX, STRING_UNIQUE_INDEX, INTFLG_INDEX, NUMERIC_INTEGER_INDEX, NUMERIC_INTEGER_NO_INDEX
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
 *     
 * 
 * [foreign property]
 *     
 * 
 * [referrer property]
 *     
 * 
 * [get/set template]
 * /= = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
 * Long largeDataId = entity.getLargeDataId();
 * String stringIndex = entity.getStringIndex();
 * String stringNoIndex = entity.getStringNoIndex();
 * String stringUniqueIndex = entity.getStringUniqueIndex();
 * Integer intflgIndex = entity.getIntflgIndex();
 * Integer numericIntegerIndex = entity.getNumericIntegerIndex();
 * Integer numericIntegerNoIndex = entity.getNumericIntegerNoIndex();
 * entity.setLargeDataId(largeDataId);
 * entity.setStringIndex(stringIndex);
 * entity.setStringNoIndex(stringNoIndex);
 * entity.setStringUniqueIndex(stringUniqueIndex);
 * entity.setIntflgIndex(intflgIndex);
 * entity.setNumericIntegerIndex(numericIntegerIndex);
 * entity.setNumericIntegerNoIndex(numericIntegerNoIndex);
 * = = = = = = = = = =/
 * </pre>
 * @author DBFlute(AutoGenerator)
 */
public abstract class BsVendorLargeData extends AbstractEntity implements DomainEntity {

    // ===================================================================================
    //                                                                          Definition
    //                                                                          ==========
    /** The serial version UID for object serialization. (Default) */
    private static final long serialVersionUID = 1L;

    // ===================================================================================
    //                                                                           Attribute
    //                                                                           =========
    /** LARGE_DATA_ID: {PK, NotNull, BIGINT(19)} */
    protected Long _largeDataId;

    /** STRING_INDEX: {NotNull, VARCHAR(180)} */
    protected String _stringIndex;

    /** STRING_NO_INDEX: {NotNull, VARCHAR(180)} */
    protected String _stringNoIndex;

    /** STRING_UNIQUE_INDEX: {UQ, NotNull, VARCHAR(180)} */
    protected String _stringUniqueIndex;

    /** INTFLG_INDEX: {NotNull, INT(10)} */
    protected Integer _intflgIndex;

    /** NUMERIC_INTEGER_INDEX: {NotNull, DECIMAL(8)} */
    protected Integer _numericIntegerIndex;

    /** NUMERIC_INTEGER_NO_INDEX: {NotNull, DECIMAL(8)} */
    protected Integer _numericIntegerNoIndex;

    // ===================================================================================
    //                                                                             DB Meta
    //                                                                             =======
    /** {@inheritDoc} */
    public DBMeta asDBMeta() {
        return DBMetaInstanceHandler.findDBMeta(asTableDbName());
    }

    /** {@inheritDoc} */
    public String asTableDbName() {
        return "vendor_large_data";
    }

    // ===================================================================================
    //                                                                        Key Handling
    //                                                                        ============
    /** {@inheritDoc} */
    public boolean hasPrimaryKeyValue() {
        if (_largeDataId == null) { return false; }
        return true;
    }

    /**
     * To be unique by the unique column. <br>
     * You can update the entity by the key when entity update (NOT batch update).
     * @param stringUniqueIndex : UQ, NotNull, VARCHAR(180). (NotNull)
     */
    public void uniqueBy(String stringUniqueIndex) {
        __uniqueDrivenProperties.clear();
        __uniqueDrivenProperties.addPropertyName("stringUniqueIndex");
        setStringUniqueIndex(stringUniqueIndex);
    }

    // ===================================================================================
    //                                                                    Foreign Property
    //                                                                    ================
    // ===================================================================================
    //                                                                   Referrer Property
    //                                                                   =================
    protected <ELEMENT> List<ELEMENT> newReferrerList() {
        return new ArrayList<ELEMENT>();
    }

    // ===================================================================================
    //                                                                      Basic Override
    //                                                                      ==============
    @Override
    protected boolean doEquals(Object obj) {
        if (obj instanceof BsVendorLargeData) {
            BsVendorLargeData other = (BsVendorLargeData)obj;
            if (!xSV(_largeDataId, other._largeDataId)) { return false; }
            return true;
        } else {
            return false;
        }
    }

    @Override
    protected int doHashCode(int initial) {
        int hs = initial;
        hs = xCH(hs, asTableDbName());
        hs = xCH(hs, _largeDataId);
        return hs;
    }

    @Override
    protected String doBuildStringWithRelation(String li) {
        return "";
    }

    @Override
    protected String doBuildColumnString(String dm) {
        StringBuilder sb = new StringBuilder();
        sb.append(dm).append(xfND(_largeDataId));
        sb.append(dm).append(xfND(_stringIndex));
        sb.append(dm).append(xfND(_stringNoIndex));
        sb.append(dm).append(xfND(_stringUniqueIndex));
        sb.append(dm).append(xfND(_intflgIndex));
        sb.append(dm).append(xfND(_numericIntegerIndex));
        sb.append(dm).append(xfND(_numericIntegerNoIndex));
        if (sb.length() > dm.length()) {
            sb.delete(0, dm.length());
        }
        sb.insert(0, "{").append("}");
        return sb.toString();
    }

    @Override
    protected String doBuildRelationString(String dm) {
        return "";
    }

    @Override
    public VendorLargeData clone() {
        return (VendorLargeData)super.clone();
    }

    // ===================================================================================
    //                                                                            Accessor
    //                                                                            ========
    /**
     * [get] LARGE_DATA_ID: {PK, NotNull, BIGINT(19)} <br>
     * @return The value of the column 'LARGE_DATA_ID'. (basically NotNull if selected: for the constraint)
     */
    public Long getLargeDataId() {
        checkSpecifiedProperty("largeDataId");
        return _largeDataId;
    }

    /**
     * [set] LARGE_DATA_ID: {PK, NotNull, BIGINT(19)} <br>
     * @param largeDataId The value of the column 'LARGE_DATA_ID'. (basically NotNull if update: for the constraint)
     */
    public void setLargeDataId(Long largeDataId) {
        registerModifiedProperty("largeDataId");
        _largeDataId = largeDataId;
    }

    /**
     * [get] STRING_INDEX: {NotNull, VARCHAR(180)} <br>
     * @return The value of the column 'STRING_INDEX'. (basically NotNull if selected: for the constraint)
     */
    public String getStringIndex() {
        checkSpecifiedProperty("stringIndex");
        return _stringIndex;
    }

    /**
     * [set] STRING_INDEX: {NotNull, VARCHAR(180)} <br>
     * @param stringIndex The value of the column 'STRING_INDEX'. (basically NotNull if update: for the constraint)
     */
    public void setStringIndex(String stringIndex) {
        registerModifiedProperty("stringIndex");
        _stringIndex = stringIndex;
    }

    /**
     * [get] STRING_NO_INDEX: {NotNull, VARCHAR(180)} <br>
     * @return The value of the column 'STRING_NO_INDEX'. (basically NotNull if selected: for the constraint)
     */
    public String getStringNoIndex() {
        checkSpecifiedProperty("stringNoIndex");
        return _stringNoIndex;
    }

    /**
     * [set] STRING_NO_INDEX: {NotNull, VARCHAR(180)} <br>
     * @param stringNoIndex The value of the column 'STRING_NO_INDEX'. (basically NotNull if update: for the constraint)
     */
    public void setStringNoIndex(String stringNoIndex) {
        registerModifiedProperty("stringNoIndex");
        _stringNoIndex = stringNoIndex;
    }

    /**
     * [get] STRING_UNIQUE_INDEX: {UQ, NotNull, VARCHAR(180)} <br>
     * @return The value of the column 'STRING_UNIQUE_INDEX'. (basically NotNull if selected: for the constraint)
     */
    public String getStringUniqueIndex() {
        checkSpecifiedProperty("stringUniqueIndex");
        return _stringUniqueIndex;
    }

    /**
     * [set] STRING_UNIQUE_INDEX: {UQ, NotNull, VARCHAR(180)} <br>
     * @param stringUniqueIndex The value of the column 'STRING_UNIQUE_INDEX'. (basically NotNull if update: for the constraint)
     */
    public void setStringUniqueIndex(String stringUniqueIndex) {
        registerModifiedProperty("stringUniqueIndex");
        _stringUniqueIndex = stringUniqueIndex;
    }

    /**
     * [get] INTFLG_INDEX: {NotNull, INT(10)} <br>
     * @return The value of the column 'INTFLG_INDEX'. (basically NotNull if selected: for the constraint)
     */
    public Integer getIntflgIndex() {
        checkSpecifiedProperty("intflgIndex");
        return _intflgIndex;
    }

    /**
     * [set] INTFLG_INDEX: {NotNull, INT(10)} <br>
     * @param intflgIndex The value of the column 'INTFLG_INDEX'. (basically NotNull if update: for the constraint)
     */
    public void setIntflgIndex(Integer intflgIndex) {
        registerModifiedProperty("intflgIndex");
        _intflgIndex = intflgIndex;
    }

    /**
     * [get] NUMERIC_INTEGER_INDEX: {NotNull, DECIMAL(8)} <br>
     * @return The value of the column 'NUMERIC_INTEGER_INDEX'. (basically NotNull if selected: for the constraint)
     */
    public Integer getNumericIntegerIndex() {
        checkSpecifiedProperty("numericIntegerIndex");
        return _numericIntegerIndex;
    }

    /**
     * [set] NUMERIC_INTEGER_INDEX: {NotNull, DECIMAL(8)} <br>
     * @param numericIntegerIndex The value of the column 'NUMERIC_INTEGER_INDEX'. (basically NotNull if update: for the constraint)
     */
    public void setNumericIntegerIndex(Integer numericIntegerIndex) {
        registerModifiedProperty("numericIntegerIndex");
        _numericIntegerIndex = numericIntegerIndex;
    }

    /**
     * [get] NUMERIC_INTEGER_NO_INDEX: {NotNull, DECIMAL(8)} <br>
     * @return The value of the column 'NUMERIC_INTEGER_NO_INDEX'. (basically NotNull if selected: for the constraint)
     */
    public Integer getNumericIntegerNoIndex() {
        checkSpecifiedProperty("numericIntegerNoIndex");
        return _numericIntegerNoIndex;
    }

    /**
     * [set] NUMERIC_INTEGER_NO_INDEX: {NotNull, DECIMAL(8)} <br>
     * @param numericIntegerNoIndex The value of the column 'NUMERIC_INTEGER_NO_INDEX'. (basically NotNull if update: for the constraint)
     */
    public void setNumericIntegerNoIndex(Integer numericIntegerNoIndex) {
        registerModifiedProperty("numericIntegerNoIndex");
        _numericIntegerNoIndex = numericIntegerNoIndex;
    }
}

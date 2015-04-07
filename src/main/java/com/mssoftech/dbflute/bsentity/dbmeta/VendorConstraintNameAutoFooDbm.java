package com.mssoftech.dbflute.bsentity.dbmeta;

import java.util.List;
import java.util.Map;

import org.dbflute.Entity;
import org.dbflute.dbmeta.AbstractDBMeta;
import org.dbflute.dbmeta.info.*;
import org.dbflute.dbmeta.name.*;
import org.dbflute.dbmeta.property.PropertyGateway;
import org.dbflute.dbway.DBDef;
import com.mssoftech.dbflute.allcommon.*;
import com.mssoftech.dbflute.exentity.*;

/**
 * The DB meta of vendor_constraint_name_auto_foo. (Singleton)
 * @author DBFlute(AutoGenerator)
 */
public class VendorConstraintNameAutoFooDbm extends AbstractDBMeta {

    // ===================================================================================
    //                                                                           Singleton
    //                                                                           =========
    private static final VendorConstraintNameAutoFooDbm _instance = new VendorConstraintNameAutoFooDbm();
    private VendorConstraintNameAutoFooDbm() {}
    public static VendorConstraintNameAutoFooDbm getInstance() { return _instance; }

    // ===================================================================================
    //                                                                       Current DBDef
    //                                                                       =============
    public String getProjectName() { return DBCurrent.getInstance().projectName(); }
    public String getProjectPrefix() { return DBCurrent.getInstance().projectPrefix(); }
    public String getGenerationGapBasePrefix() { return DBCurrent.getInstance().generationGapBasePrefix(); }
    public DBDef getCurrentDBDef() { return DBCurrent.getInstance().currentDBDef(); }

    // ===================================================================================
    //                                                                    Property Gateway
    //                                                                    ================
    // -----------------------------------------------------
    //                                       Column Property
    //                                       ---------------
    protected final Map<String, PropertyGateway> _epgMap = newHashMap();
    { xsetupEpg(); }
    protected void xsetupEpg() {
        setupEpg(_epgMap, et -> ((VendorConstraintNameAutoFoo)et).getConstraintNameAutoFooId(), (et, vl) -> ((VendorConstraintNameAutoFoo)et).setConstraintNameAutoFooId(ctl(vl)), "constraintNameAutoFooId");
        setupEpg(_epgMap, et -> ((VendorConstraintNameAutoFoo)et).getConstraintNameAutoFooName(), (et, vl) -> ((VendorConstraintNameAutoFoo)et).setConstraintNameAutoFooName((String)vl), "constraintNameAutoFooName");
    }
    public PropertyGateway findPropertyGateway(String prop)
    { return doFindEpg(_epgMap, prop); }

    // ===================================================================================
    //                                                                          Table Info
    //                                                                          ==========
    protected final String _tableDbName = "vendor_constraint_name_auto_foo";
    protected final String _tableDispName = "vendor_constraint_name_auto_foo";
    protected final String _tablePropertyName = "vendorConstraintNameAutoFoo";
    protected final TableSqlName _tableSqlName = new TableSqlName("vendor_constraint_name_auto_foo", _tableDbName);
    { _tableSqlName.xacceptFilter(DBFluteConfig.getInstance().getTableSqlNameFilter()); }
    public String getTableDbName() { return _tableDbName; }
    public String getTableDispName() { return _tableDispName; }
    public String getTablePropertyName() { return _tablePropertyName; }
    public TableSqlName getTableSqlName() { return _tableSqlName; }

    // ===================================================================================
    //                                                                         Column Info
    //                                                                         ===========
    protected final ColumnInfo _columnConstraintNameAutoFooId = cci("CONSTRAINT_NAME_AUTO_FOO_ID", "CONSTRAINT_NAME_AUTO_FOO_ID", null, null, Long.class, "constraintNameAutoFooId", null, true, false, true, "DECIMAL", 16, 0, null, false, null, null, null, "vendorConstraintNameAutoRefList", null, false);
//"df.Numeric"
    protected final ColumnInfo _columnConstraintNameAutoFooName = cci("CONSTRAINT_NAME_AUTO_FOO_NAME", "CONSTRAINT_NAME_AUTO_FOO_NAME", null, null, String.class, "constraintNameAutoFooName", null, false, false, true, "VARCHAR", 50, 0, null, false, null, null, null, null, null, false);
//"string"

    /**
     * CONSTRAINT_NAME_AUTO_FOO_ID: {PK, NotNull, DECIMAL(16)}
     * @return The information object of specified column. (NotNull)
     */
    public ColumnInfo columnConstraintNameAutoFooId() { return _columnConstraintNameAutoFooId; }
    /**
     * CONSTRAINT_NAME_AUTO_FOO_NAME: {UQ, NotNull, VARCHAR(50)}
     * @return The information object of specified column. (NotNull)
     */
    public ColumnInfo columnConstraintNameAutoFooName() { return _columnConstraintNameAutoFooName; }

    protected List<ColumnInfo> ccil() {
        List<ColumnInfo> ls = newArrayList();
        ls.add(columnConstraintNameAutoFooId());
        ls.add(columnConstraintNameAutoFooName());
        return ls;
    }

    { initializeInformationResource(); }

    // ===================================================================================
    //                                                                         Unique Info
    //                                                                         ===========
    // -----------------------------------------------------
    //                                       Primary Element
    //                                       ---------------
    protected UniqueInfo cpui() { return hpcpui(columnConstraintNameAutoFooId()); }
    public boolean hasPrimaryKey() { return true; }
    public boolean hasCompoundPrimaryKey() { return false; }

    // -----------------------------------------------------
    //                                        Unique Element
    //                                        --------------
    public UniqueInfo uniqueOf() { return hpcui(columnConstraintNameAutoFooName()); }

    // ===================================================================================
    //                                                                       Relation Info
    //                                                                       =============
    // cannot cache because it uses related DB meta instance while booting
    // (instead, cached by super's collection)
    // -----------------------------------------------------
    //                                      Foreign Property
    //                                      ----------------

    // -----------------------------------------------------
    //                                     Referrer Property
    //                                     -----------------
    /**
     * vendor_constraint_name_auto_ref by CONSTRAINT_NAME_AUTO_FOO_ID, named 'vendorConstraintNameAutoRefList'.
     * @return The information object of referrer property. (NotNull)
     */
    public ReferrerInfo referrerVendorConstraintNameAutoRefList() {
        Map<ColumnInfo, ColumnInfo> mp = newLinkedHashMap(columnConstraintNameAutoFooId(), VendorConstraintNameAutoRefDbm.getInstance().columnConstraintNameAutoFooId());
        return cri("vendor_constraint_name_auto_ref_ibfk_1", "vendorConstraintNameAutoRefList", this, VendorConstraintNameAutoRefDbm.getInstance(), mp, false, "vendorConstraintNameAutoFoo");
    }

    // ===================================================================================
    //                                                                        Various Info
    //                                                                        ============

    // ===================================================================================
    //                                                                           Type Name
    //                                                                           =========
    public String getEntityTypeName() { return "com.mssoftech.dbflute.exentity.VendorConstraintNameAutoFoo"; }
    public String getConditionBeanTypeName() { return "com.mssoftech.dbflute.cbean.VendorConstraintNameAutoFooCB"; }
    public String getBehaviorTypeName() { return "com.mssoftech.dbflute.exbhv.VendorConstraintNameAutoFooBhv"; }

    // ===================================================================================
    //                                                                         Object Type
    //                                                                         ===========
    public Class<VendorConstraintNameAutoFoo> getEntityType() { return VendorConstraintNameAutoFoo.class; }

    // ===================================================================================
    //                                                                     Object Instance
    //                                                                     ===============
    public VendorConstraintNameAutoFoo newEntity() { return new VendorConstraintNameAutoFoo(); }

    // ===================================================================================
    //                                                                   Map Communication
    //                                                                   =================
    public void acceptPrimaryKeyMap(Entity et, Map<String, ? extends Object> mp)
    { doAcceptPrimaryKeyMap((VendorConstraintNameAutoFoo)et, mp); }
    public void acceptAllColumnMap(Entity et, Map<String, ? extends Object> mp)
    { doAcceptAllColumnMap((VendorConstraintNameAutoFoo)et, mp); }
    public Map<String, Object> extractPrimaryKeyMap(Entity et) { return doExtractPrimaryKeyMap(et); }
    public Map<String, Object> extractAllColumnMap(Entity et) { return doExtractAllColumnMap(et); }
}

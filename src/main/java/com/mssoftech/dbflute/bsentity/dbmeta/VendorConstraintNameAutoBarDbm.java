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
 * The DB meta of vendor_constraint_name_auto_bar. (Singleton)
 * @author DBFlute(AutoGenerator)
 */
public class VendorConstraintNameAutoBarDbm extends AbstractDBMeta {

    // ===================================================================================
    //                                                                           Singleton
    //                                                                           =========
    private static final VendorConstraintNameAutoBarDbm _instance = new VendorConstraintNameAutoBarDbm();
    private VendorConstraintNameAutoBarDbm() {}
    public static VendorConstraintNameAutoBarDbm getInstance() { return _instance; }

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
        setupEpg(_epgMap, et -> ((VendorConstraintNameAutoBar)et).getConstraintNameAutoBarId(), (et, vl) -> ((VendorConstraintNameAutoBar)et).setConstraintNameAutoBarId(ctl(vl)), "constraintNameAutoBarId");
        setupEpg(_epgMap, et -> ((VendorConstraintNameAutoBar)et).getConstraintNameAutoBarName(), (et, vl) -> ((VendorConstraintNameAutoBar)et).setConstraintNameAutoBarName((String)vl), "constraintNameAutoBarName");
    }
    public PropertyGateway findPropertyGateway(String prop)
    { return doFindEpg(_epgMap, prop); }

    // ===================================================================================
    //                                                                          Table Info
    //                                                                          ==========
    protected final String _tableDbName = "vendor_constraint_name_auto_bar";
    protected final String _tableDispName = "vendor_constraint_name_auto_bar";
    protected final String _tablePropertyName = "vendorConstraintNameAutoBar";
    protected final TableSqlName _tableSqlName = new TableSqlName("vendor_constraint_name_auto_bar", _tableDbName);
    { _tableSqlName.xacceptFilter(DBFluteConfig.getInstance().getTableSqlNameFilter()); }
    public String getTableDbName() { return _tableDbName; }
    public String getTableDispName() { return _tableDispName; }
    public String getTablePropertyName() { return _tablePropertyName; }
    public TableSqlName getTableSqlName() { return _tableSqlName; }

    // ===================================================================================
    //                                                                         Column Info
    //                                                                         ===========
    protected final ColumnInfo _columnConstraintNameAutoBarId = cci("CONSTRAINT_NAME_AUTO_BAR_ID", "CONSTRAINT_NAME_AUTO_BAR_ID", null, null, Long.class, "constraintNameAutoBarId", null, true, false, true, "DECIMAL", 16, 0, null, false, null, null, null, "vendorConstraintNameAutoRefList", null, false);
//"df.Numeric"
    protected final ColumnInfo _columnConstraintNameAutoBarName = cci("CONSTRAINT_NAME_AUTO_BAR_NAME", "CONSTRAINT_NAME_AUTO_BAR_NAME", null, null, String.class, "constraintNameAutoBarName", null, false, false, true, "VARCHAR", 50, 0, null, false, null, null, null, null, null, false);
//"string"

    /**
     * CONSTRAINT_NAME_AUTO_BAR_ID: {PK, NotNull, DECIMAL(16)}
     * @return The information object of specified column. (NotNull)
     */
    public ColumnInfo columnConstraintNameAutoBarId() { return _columnConstraintNameAutoBarId; }
    /**
     * CONSTRAINT_NAME_AUTO_BAR_NAME: {UQ, NotNull, VARCHAR(50)}
     * @return The information object of specified column. (NotNull)
     */
    public ColumnInfo columnConstraintNameAutoBarName() { return _columnConstraintNameAutoBarName; }

    protected List<ColumnInfo> ccil() {
        List<ColumnInfo> ls = newArrayList();
        ls.add(columnConstraintNameAutoBarId());
        ls.add(columnConstraintNameAutoBarName());
        return ls;
    }

    { initializeInformationResource(); }

    // ===================================================================================
    //                                                                         Unique Info
    //                                                                         ===========
    // -----------------------------------------------------
    //                                       Primary Element
    //                                       ---------------
    protected UniqueInfo cpui() { return hpcpui(columnConstraintNameAutoBarId()); }
    public boolean hasPrimaryKey() { return true; }
    public boolean hasCompoundPrimaryKey() { return false; }

    // -----------------------------------------------------
    //                                        Unique Element
    //                                        --------------
    public UniqueInfo uniqueOf() { return hpcui(columnConstraintNameAutoBarName()); }

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
     * vendor_constraint_name_auto_ref by CONSTRAINT_NAME_AUTO_BAR_ID, named 'vendorConstraintNameAutoRefList'.
     * @return The information object of referrer property. (NotNull)
     */
    public ReferrerInfo referrerVendorConstraintNameAutoRefList() {
        Map<ColumnInfo, ColumnInfo> mp = newLinkedHashMap(columnConstraintNameAutoBarId(), VendorConstraintNameAutoRefDbm.getInstance().columnConstraintNameAutoBarId());
        return cri("vendor_constraint_name_auto_ref_ibfk_2", "vendorConstraintNameAutoRefList", this, VendorConstraintNameAutoRefDbm.getInstance(), mp, false, "vendorConstraintNameAutoBar");
    }

    // ===================================================================================
    //                                                                        Various Info
    //                                                                        ============

    // ===================================================================================
    //                                                                           Type Name
    //                                                                           =========
    public String getEntityTypeName() { return "com.mssoftech.dbflute.exentity.VendorConstraintNameAutoBar"; }
    public String getConditionBeanTypeName() { return "com.mssoftech.dbflute.cbean.VendorConstraintNameAutoBarCB"; }
    public String getBehaviorTypeName() { return "com.mssoftech.dbflute.exbhv.VendorConstraintNameAutoBarBhv"; }

    // ===================================================================================
    //                                                                         Object Type
    //                                                                         ===========
    public Class<VendorConstraintNameAutoBar> getEntityType() { return VendorConstraintNameAutoBar.class; }

    // ===================================================================================
    //                                                                     Object Instance
    //                                                                     ===============
    public VendorConstraintNameAutoBar newEntity() { return new VendorConstraintNameAutoBar(); }

    // ===================================================================================
    //                                                                   Map Communication
    //                                                                   =================
    public void acceptPrimaryKeyMap(Entity et, Map<String, ? extends Object> mp)
    { doAcceptPrimaryKeyMap((VendorConstraintNameAutoBar)et, mp); }
    public void acceptAllColumnMap(Entity et, Map<String, ? extends Object> mp)
    { doAcceptAllColumnMap((VendorConstraintNameAutoBar)et, mp); }
    public Map<String, Object> extractPrimaryKeyMap(Entity et) { return doExtractPrimaryKeyMap(et); }
    public Map<String, Object> extractAllColumnMap(Entity et) { return doExtractAllColumnMap(et); }
}

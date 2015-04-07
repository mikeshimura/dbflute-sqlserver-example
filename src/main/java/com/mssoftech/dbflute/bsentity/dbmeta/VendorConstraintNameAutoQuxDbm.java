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
 * The DB meta of vendor_constraint_name_auto_qux. (Singleton)
 * @author DBFlute(AutoGenerator)
 */
public class VendorConstraintNameAutoQuxDbm extends AbstractDBMeta {

    // ===================================================================================
    //                                                                           Singleton
    //                                                                           =========
    private static final VendorConstraintNameAutoQuxDbm _instance = new VendorConstraintNameAutoQuxDbm();
    private VendorConstraintNameAutoQuxDbm() {}
    public static VendorConstraintNameAutoQuxDbm getInstance() { return _instance; }

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
        setupEpg(_epgMap, et -> ((VendorConstraintNameAutoQux)et).getConstraintNameAutoQuxId(), (et, vl) -> ((VendorConstraintNameAutoQux)et).setConstraintNameAutoQuxId(ctl(vl)), "constraintNameAutoQuxId");
        setupEpg(_epgMap, et -> ((VendorConstraintNameAutoQux)et).getConstraintNameAutoQuxName(), (et, vl) -> ((VendorConstraintNameAutoQux)et).setConstraintNameAutoQuxName((String)vl), "constraintNameAutoQuxName");
    }
    public PropertyGateway findPropertyGateway(String prop)
    { return doFindEpg(_epgMap, prop); }

    // ===================================================================================
    //                                                                          Table Info
    //                                                                          ==========
    protected final String _tableDbName = "vendor_constraint_name_auto_qux";
    protected final String _tableDispName = "vendor_constraint_name_auto_qux";
    protected final String _tablePropertyName = "vendorConstraintNameAutoQux";
    protected final TableSqlName _tableSqlName = new TableSqlName("vendor_constraint_name_auto_qux", _tableDbName);
    { _tableSqlName.xacceptFilter(DBFluteConfig.getInstance().getTableSqlNameFilter()); }
    public String getTableDbName() { return _tableDbName; }
    public String getTableDispName() { return _tableDispName; }
    public String getTablePropertyName() { return _tablePropertyName; }
    public TableSqlName getTableSqlName() { return _tableSqlName; }

    // ===================================================================================
    //                                                                         Column Info
    //                                                                         ===========
    protected final ColumnInfo _columnConstraintNameAutoQuxId = cci("CONSTRAINT_NAME_AUTO_QUX_ID", "CONSTRAINT_NAME_AUTO_QUX_ID", null, null, Long.class, "constraintNameAutoQuxId", null, true, false, true, "DECIMAL", 16, 0, null, false, null, null, null, "vendorConstraintNameAutoRefList", null, false);
//"df.Numeric"
    protected final ColumnInfo _columnConstraintNameAutoQuxName = cci("CONSTRAINT_NAME_AUTO_QUX_NAME", "CONSTRAINT_NAME_AUTO_QUX_NAME", null, null, String.class, "constraintNameAutoQuxName", null, false, false, true, "VARCHAR", 50, 0, null, false, null, null, null, null, null, false);
//"string"

    /**
     * CONSTRAINT_NAME_AUTO_QUX_ID: {PK, NotNull, DECIMAL(16)}
     * @return The information object of specified column. (NotNull)
     */
    public ColumnInfo columnConstraintNameAutoQuxId() { return _columnConstraintNameAutoQuxId; }
    /**
     * CONSTRAINT_NAME_AUTO_QUX_NAME: {UQ, NotNull, VARCHAR(50)}
     * @return The information object of specified column. (NotNull)
     */
    public ColumnInfo columnConstraintNameAutoQuxName() { return _columnConstraintNameAutoQuxName; }

    protected List<ColumnInfo> ccil() {
        List<ColumnInfo> ls = newArrayList();
        ls.add(columnConstraintNameAutoQuxId());
        ls.add(columnConstraintNameAutoQuxName());
        return ls;
    }

    { initializeInformationResource(); }

    // ===================================================================================
    //                                                                         Unique Info
    //                                                                         ===========
    // -----------------------------------------------------
    //                                       Primary Element
    //                                       ---------------
    protected UniqueInfo cpui() { return hpcpui(columnConstraintNameAutoQuxId()); }
    public boolean hasPrimaryKey() { return true; }
    public boolean hasCompoundPrimaryKey() { return false; }

    // -----------------------------------------------------
    //                                        Unique Element
    //                                        --------------
    public UniqueInfo uniqueOf() { return hpcui(columnConstraintNameAutoQuxName()); }

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
     * vendor_constraint_name_auto_ref by CONSTRAINT_NAME_AUTO_QUX_ID, named 'vendorConstraintNameAutoRefList'.
     * @return The information object of referrer property. (NotNull)
     */
    public ReferrerInfo referrerVendorConstraintNameAutoRefList() {
        Map<ColumnInfo, ColumnInfo> mp = newLinkedHashMap(columnConstraintNameAutoQuxId(), VendorConstraintNameAutoRefDbm.getInstance().columnConstraintNameAutoQuxId());
        return cri("vendor_constraint_name_auto_ref_ibfk_3", "vendorConstraintNameAutoRefList", this, VendorConstraintNameAutoRefDbm.getInstance(), mp, false, "vendorConstraintNameAutoQux");
    }

    // ===================================================================================
    //                                                                        Various Info
    //                                                                        ============

    // ===================================================================================
    //                                                                           Type Name
    //                                                                           =========
    public String getEntityTypeName() { return "com.mssoftech.dbflute.exentity.VendorConstraintNameAutoQux"; }
    public String getConditionBeanTypeName() { return "com.mssoftech.dbflute.cbean.VendorConstraintNameAutoQuxCB"; }
    public String getBehaviorTypeName() { return "com.mssoftech.dbflute.exbhv.VendorConstraintNameAutoQuxBhv"; }

    // ===================================================================================
    //                                                                         Object Type
    //                                                                         ===========
    public Class<VendorConstraintNameAutoQux> getEntityType() { return VendorConstraintNameAutoQux.class; }

    // ===================================================================================
    //                                                                     Object Instance
    //                                                                     ===============
    public VendorConstraintNameAutoQux newEntity() { return new VendorConstraintNameAutoQux(); }

    // ===================================================================================
    //                                                                   Map Communication
    //                                                                   =================
    public void acceptPrimaryKeyMap(Entity et, Map<String, ? extends Object> mp)
    { doAcceptPrimaryKeyMap((VendorConstraintNameAutoQux)et, mp); }
    public void acceptAllColumnMap(Entity et, Map<String, ? extends Object> mp)
    { doAcceptAllColumnMap((VendorConstraintNameAutoQux)et, mp); }
    public Map<String, Object> extractPrimaryKeyMap(Entity et) { return doExtractPrimaryKeyMap(et); }
    public Map<String, Object> extractAllColumnMap(Entity et) { return doExtractAllColumnMap(et); }
}

package meta

import (
	"github.com/mikeshimura/dbflute/df"
)

func MetaInit() {
	Create_MemberDbm()
	var Member_meta df.DBMeta = MemberDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["Member"] = &Member_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("Member", "Member")
	Create_MemberAddressDbm()
	var MemberAddress_meta df.DBMeta = MemberAddressDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["MemberAddress"] = &MemberAddress_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("MemberAddress", "MemberAddress")
	Create_MemberLoginDbm()
	var MemberLogin_meta df.DBMeta = MemberLoginDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["MemberLogin"] = &MemberLogin_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("MemberLogin", "MemberLogin")
	Create_MemberSecurityDbm()
	var MemberSecurity_meta df.DBMeta = MemberSecurityDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["MemberSecurity"] = &MemberSecurity_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("MemberSecurity", "MemberSecurity")
	Create_MemberServiceDbm()
	var MemberService_meta df.DBMeta = MemberServiceDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["MemberService"] = &MemberService_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("MemberService", "MemberService")
	Create_MemberStatusDbm()
	var MemberStatus_meta df.DBMeta = MemberStatusDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["MemberStatus"] = &MemberStatus_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("MemberStatus", "MemberStatus")
	Create_MemberWithdrawalDbm()
	var MemberWithdrawal_meta df.DBMeta = MemberWithdrawalDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["MemberWithdrawal"] = &MemberWithdrawal_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("MemberWithdrawal", "MemberWithdrawal")
	Create_ProductDbm()
	var Product_meta df.DBMeta = ProductDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["Product"] = &Product_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("Product", "Product")
	Create_ProductCategoryDbm()
	var ProductCategory_meta df.DBMeta = ProductCategoryDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["ProductCategory"] = &ProductCategory_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("ProductCategory", "ProductCategory")
	Create_ProductStatusDbm()
	var ProductStatus_meta df.DBMeta = ProductStatusDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["ProductStatus"] = &ProductStatus_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("ProductStatus", "ProductStatus")
	Create_PurchaseDbm()
	var Purchase_meta df.DBMeta = PurchaseDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["Purchase"] = &Purchase_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("Purchase", "Purchase")
	Create_PurchasePaymentDbm()
	var PurchasePayment_meta df.DBMeta = PurchasePaymentDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["PurchasePayment"] = &PurchasePayment_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("PurchasePayment", "PurchasePayment")
	Create_RegionDbm()
	var Region_meta df.DBMeta = RegionDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["Region"] = &Region_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("Region", "Region")
	Create_ServiceRankDbm()
	var ServiceRank_meta df.DBMeta = ServiceRankDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["ServiceRank"] = &ServiceRank_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("ServiceRank", "ServiceRank")
	Create_SummaryProductDbm()
	var SummaryProduct_meta df.DBMeta = SummaryProductDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["SummaryProduct"] = &SummaryProduct_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("SummaryProduct", "SummaryProduct")
	Create_SummaryWithdrawalDbm()
	var SummaryWithdrawal_meta df.DBMeta = SummaryWithdrawalDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["SummaryWithdrawal"] = &SummaryWithdrawal_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("SummaryWithdrawal", "SummaryWithdrawal")
	Create_VendorConstraintNameAutoBarDbm()
	var VendorConstraintNameAutoBar_meta df.DBMeta = VendorConstraintNameAutoBarDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["VendorConstraintNameAutoBar"] = &VendorConstraintNameAutoBar_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("VendorConstraintNameAutoBar", "VendorConstraintNameAutoBar")
	Create_VendorConstraintNameAutoFooDbm()
	var VendorConstraintNameAutoFoo_meta df.DBMeta = VendorConstraintNameAutoFooDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["VendorConstraintNameAutoFoo"] = &VendorConstraintNameAutoFoo_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("VendorConstraintNameAutoFoo", "VendorConstraintNameAutoFoo")
	Create_VendorConstraintNameAutoQuxDbm()
	var VendorConstraintNameAutoQux_meta df.DBMeta = VendorConstraintNameAutoQuxDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["VendorConstraintNameAutoQux"] = &VendorConstraintNameAutoQux_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("VendorConstraintNameAutoQux", "VendorConstraintNameAutoQux")
	Create_VendorConstraintNameAutoRefDbm()
	var VendorConstraintNameAutoRef_meta df.DBMeta = VendorConstraintNameAutoRefDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["VendorConstraintNameAutoRef"] = &VendorConstraintNameAutoRef_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("VendorConstraintNameAutoRef", "VendorConstraintNameAutoRef")
	Create_VendorLargeDataDbm()
	var VendorLargeData_meta df.DBMeta = VendorLargeDataDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["VendorLargeData"] = &VendorLargeData_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("VendorLargeData", "VendorLargeData")
	Create_VendorLargeDataRefDbm()
	var VendorLargeDataRef_meta df.DBMeta = VendorLargeDataRefDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["VendorLargeDataRef"] = &VendorLargeDataRef_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("VendorLargeDataRef", "VendorLargeDataRef")
	Create_VendorTheLongAndWindingTableAndColumnDbm()
	var VendorTheLongAndWindingTableAndColumn_meta df.DBMeta = VendorTheLongAndWindingTableAndColumnDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["VendorTheLongAndWindingTableAndColumn"] = &VendorTheLongAndWindingTableAndColumn_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("VendorTheLongAndWindingTableAndColumn", "VendorTheLongAndWindingTableAndColumn")
	Create_VendorTheLongAndWindingTableAndColumnRefDbm()
	var VendorTheLongAndWindingTableAndColumnRef_meta df.DBMeta = VendorTheLongAndWindingTableAndColumnRefDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["VendorTheLongAndWindingTableAndColumnRef"] = &VendorTheLongAndWindingTableAndColumnRef_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("VendorTheLongAndWindingTableAndColumnRef", "VendorTheLongAndWindingTableAndColumnRef")
	Create_WithdrawalReasonDbm()
	var WithdrawalReason_meta df.DBMeta = WithdrawalReasonDbm
	df.DBMetaProvider_I.TableDbNameInstanceMap["WithdrawalReason"] = &WithdrawalReason_meta
	df.DBMetaProvider_I.TableDbNameFlexibleMap.Put("WithdrawalReason", "WithdrawalReason")
}
package bhv

import (
	"dbflute/adf"
	"github.com/mikeshimura/dbflute/df"
	"dbflute/adf/entity"
	"dbflute/adf/meta"
	_	"github.com/lib/pq"  //Don't eraze require to use this package
)
//To fix init sequence, all init routines are called from here.
func init() {
	adf.AdfInit()
	entity.EntityInit()
	meta.MetaInit()
	MemberBhv_I = new(MemberBhv)
	MemberBhv_I.BaseBehavior = new(df.BaseBehavior)
	MemberBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	MemberBhv_I.BaseBehavior.TableDbName = "Member"
	var member df.Behavior =MemberBhv_I
	MemberBhv_I.BaseBehavior.Behavior=&member
	MemberBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	MemberAddressBhv_I = new(MemberAddressBhv)
	MemberAddressBhv_I.BaseBehavior = new(df.BaseBehavior)
	MemberAddressBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	MemberAddressBhv_I.BaseBehavior.TableDbName = "MemberAddress"
	var memberAddress df.Behavior =MemberAddressBhv_I
	MemberAddressBhv_I.BaseBehavior.Behavior=&memberAddress
	MemberAddressBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	MemberLoginBhv_I = new(MemberLoginBhv)
	MemberLoginBhv_I.BaseBehavior = new(df.BaseBehavior)
	MemberLoginBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	MemberLoginBhv_I.BaseBehavior.TableDbName = "MemberLogin"
	var memberLogin df.Behavior =MemberLoginBhv_I
	MemberLoginBhv_I.BaseBehavior.Behavior=&memberLogin
	MemberLoginBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	MemberSecurityBhv_I = new(MemberSecurityBhv)
	MemberSecurityBhv_I.BaseBehavior = new(df.BaseBehavior)
	MemberSecurityBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	MemberSecurityBhv_I.BaseBehavior.TableDbName = "MemberSecurity"
	var memberSecurity df.Behavior =MemberSecurityBhv_I
	MemberSecurityBhv_I.BaseBehavior.Behavior=&memberSecurity
	MemberSecurityBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	MemberServiceBhv_I = new(MemberServiceBhv)
	MemberServiceBhv_I.BaseBehavior = new(df.BaseBehavior)
	MemberServiceBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	MemberServiceBhv_I.BaseBehavior.TableDbName = "MemberService"
	var memberService df.Behavior =MemberServiceBhv_I
	MemberServiceBhv_I.BaseBehavior.Behavior=&memberService
	MemberServiceBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	MemberStatusBhv_I = new(MemberStatusBhv)
	MemberStatusBhv_I.BaseBehavior = new(df.BaseBehavior)
	MemberStatusBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	MemberStatusBhv_I.BaseBehavior.TableDbName = "MemberStatus"
	var memberStatus df.Behavior =MemberStatusBhv_I
	MemberStatusBhv_I.BaseBehavior.Behavior=&memberStatus
	MemberStatusBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	MemberWithdrawalBhv_I = new(MemberWithdrawalBhv)
	MemberWithdrawalBhv_I.BaseBehavior = new(df.BaseBehavior)
	MemberWithdrawalBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	MemberWithdrawalBhv_I.BaseBehavior.TableDbName = "MemberWithdrawal"
	var memberWithdrawal df.Behavior =MemberWithdrawalBhv_I
	MemberWithdrawalBhv_I.BaseBehavior.Behavior=&memberWithdrawal
	MemberWithdrawalBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	ProductBhv_I = new(ProductBhv)
	ProductBhv_I.BaseBehavior = new(df.BaseBehavior)
	ProductBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	ProductBhv_I.BaseBehavior.TableDbName = "Product"
	var product df.Behavior =ProductBhv_I
	ProductBhv_I.BaseBehavior.Behavior=&product
	ProductBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	ProductCategoryBhv_I = new(ProductCategoryBhv)
	ProductCategoryBhv_I.BaseBehavior = new(df.BaseBehavior)
	ProductCategoryBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	ProductCategoryBhv_I.BaseBehavior.TableDbName = "ProductCategory"
	var productCategory df.Behavior =ProductCategoryBhv_I
	ProductCategoryBhv_I.BaseBehavior.Behavior=&productCategory
	ProductCategoryBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	ProductStatusBhv_I = new(ProductStatusBhv)
	ProductStatusBhv_I.BaseBehavior = new(df.BaseBehavior)
	ProductStatusBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	ProductStatusBhv_I.BaseBehavior.TableDbName = "ProductStatus"
	var productStatus df.Behavior =ProductStatusBhv_I
	ProductStatusBhv_I.BaseBehavior.Behavior=&productStatus
	ProductStatusBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	PurchaseBhv_I = new(PurchaseBhv)
	PurchaseBhv_I.BaseBehavior = new(df.BaseBehavior)
	PurchaseBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	PurchaseBhv_I.BaseBehavior.TableDbName = "Purchase"
	var purchase df.Behavior =PurchaseBhv_I
	PurchaseBhv_I.BaseBehavior.Behavior=&purchase
	PurchaseBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	PurchasePaymentBhv_I = new(PurchasePaymentBhv)
	PurchasePaymentBhv_I.BaseBehavior = new(df.BaseBehavior)
	PurchasePaymentBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	PurchasePaymentBhv_I.BaseBehavior.TableDbName = "PurchasePayment"
	var purchasePayment df.Behavior =PurchasePaymentBhv_I
	PurchasePaymentBhv_I.BaseBehavior.Behavior=&purchasePayment
	PurchasePaymentBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	RegionBhv_I = new(RegionBhv)
	RegionBhv_I.BaseBehavior = new(df.BaseBehavior)
	RegionBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	RegionBhv_I.BaseBehavior.TableDbName = "Region"
	var region df.Behavior =RegionBhv_I
	RegionBhv_I.BaseBehavior.Behavior=&region
	RegionBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	ServiceRankBhv_I = new(ServiceRankBhv)
	ServiceRankBhv_I.BaseBehavior = new(df.BaseBehavior)
	ServiceRankBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	ServiceRankBhv_I.BaseBehavior.TableDbName = "ServiceRank"
	var serviceRank df.Behavior =ServiceRankBhv_I
	ServiceRankBhv_I.BaseBehavior.Behavior=&serviceRank
	ServiceRankBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	SummaryProductBhv_I = new(SummaryProductBhv)
	SummaryProductBhv_I.BaseBehavior = new(df.BaseBehavior)
	SummaryProductBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	SummaryProductBhv_I.BaseBehavior.TableDbName = "SummaryProduct"
	var summaryProduct df.Behavior =SummaryProductBhv_I
	SummaryProductBhv_I.BaseBehavior.Behavior=&summaryProduct
	SummaryProductBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	SummaryWithdrawalBhv_I = new(SummaryWithdrawalBhv)
	SummaryWithdrawalBhv_I.BaseBehavior = new(df.BaseBehavior)
	SummaryWithdrawalBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	SummaryWithdrawalBhv_I.BaseBehavior.TableDbName = "SummaryWithdrawal"
	var summaryWithdrawal df.Behavior =SummaryWithdrawalBhv_I
	SummaryWithdrawalBhv_I.BaseBehavior.Behavior=&summaryWithdrawal
	SummaryWithdrawalBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	VendorConstraintNameAutoBarBhv_I = new(VendorConstraintNameAutoBarBhv)
	VendorConstraintNameAutoBarBhv_I.BaseBehavior = new(df.BaseBehavior)
	VendorConstraintNameAutoBarBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	VendorConstraintNameAutoBarBhv_I.BaseBehavior.TableDbName = "VendorConstraintNameAutoBar"
	var vendorConstraintNameAutoBar df.Behavior =VendorConstraintNameAutoBarBhv_I
	VendorConstraintNameAutoBarBhv_I.BaseBehavior.Behavior=&vendorConstraintNameAutoBar
	VendorConstraintNameAutoBarBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	VendorConstraintNameAutoFooBhv_I = new(VendorConstraintNameAutoFooBhv)
	VendorConstraintNameAutoFooBhv_I.BaseBehavior = new(df.BaseBehavior)
	VendorConstraintNameAutoFooBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	VendorConstraintNameAutoFooBhv_I.BaseBehavior.TableDbName = "VendorConstraintNameAutoFoo"
	var vendorConstraintNameAutoFoo df.Behavior =VendorConstraintNameAutoFooBhv_I
	VendorConstraintNameAutoFooBhv_I.BaseBehavior.Behavior=&vendorConstraintNameAutoFoo
	VendorConstraintNameAutoFooBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	VendorConstraintNameAutoQuxBhv_I = new(VendorConstraintNameAutoQuxBhv)
	VendorConstraintNameAutoQuxBhv_I.BaseBehavior = new(df.BaseBehavior)
	VendorConstraintNameAutoQuxBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	VendorConstraintNameAutoQuxBhv_I.BaseBehavior.TableDbName = "VendorConstraintNameAutoQux"
	var vendorConstraintNameAutoQux df.Behavior =VendorConstraintNameAutoQuxBhv_I
	VendorConstraintNameAutoQuxBhv_I.BaseBehavior.Behavior=&vendorConstraintNameAutoQux
	VendorConstraintNameAutoQuxBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	VendorConstraintNameAutoRefBhv_I = new(VendorConstraintNameAutoRefBhv)
	VendorConstraintNameAutoRefBhv_I.BaseBehavior = new(df.BaseBehavior)
	VendorConstraintNameAutoRefBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	VendorConstraintNameAutoRefBhv_I.BaseBehavior.TableDbName = "VendorConstraintNameAutoRef"
	var vendorConstraintNameAutoRef df.Behavior =VendorConstraintNameAutoRefBhv_I
	VendorConstraintNameAutoRefBhv_I.BaseBehavior.Behavior=&vendorConstraintNameAutoRef
	VendorConstraintNameAutoRefBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	VendorLargeDataBhv_I = new(VendorLargeDataBhv)
	VendorLargeDataBhv_I.BaseBehavior = new(df.BaseBehavior)
	VendorLargeDataBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	VendorLargeDataBhv_I.BaseBehavior.TableDbName = "VendorLargeData"
	var vendorLargeData df.Behavior =VendorLargeDataBhv_I
	VendorLargeDataBhv_I.BaseBehavior.Behavior=&vendorLargeData
	VendorLargeDataBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	VendorLargeDataRefBhv_I = new(VendorLargeDataRefBhv)
	VendorLargeDataRefBhv_I.BaseBehavior = new(df.BaseBehavior)
	VendorLargeDataRefBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	VendorLargeDataRefBhv_I.BaseBehavior.TableDbName = "VendorLargeDataRef"
	var vendorLargeDataRef df.Behavior =VendorLargeDataRefBhv_I
	VendorLargeDataRefBhv_I.BaseBehavior.Behavior=&vendorLargeDataRef
	VendorLargeDataRefBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	VendorTheLongAndWindingTableAndColumnBhv_I = new(VendorTheLongAndWindingTableAndColumnBhv)
	VendorTheLongAndWindingTableAndColumnBhv_I.BaseBehavior = new(df.BaseBehavior)
	VendorTheLongAndWindingTableAndColumnBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	VendorTheLongAndWindingTableAndColumnBhv_I.BaseBehavior.TableDbName = "VendorTheLongAndWindingTableAndColumn"
	var vendorTheLongAndWindingTableAndColumn df.Behavior =VendorTheLongAndWindingTableAndColumnBhv_I
	VendorTheLongAndWindingTableAndColumnBhv_I.BaseBehavior.Behavior=&vendorTheLongAndWindingTableAndColumn
	VendorTheLongAndWindingTableAndColumnBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	VendorTheLongAndWindingTableAndColumnRefBhv_I = new(VendorTheLongAndWindingTableAndColumnRefBhv)
	VendorTheLongAndWindingTableAndColumnRefBhv_I.BaseBehavior = new(df.BaseBehavior)
	VendorTheLongAndWindingTableAndColumnRefBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	VendorTheLongAndWindingTableAndColumnRefBhv_I.BaseBehavior.TableDbName = "VendorTheLongAndWindingTableAndColumnRef"
	var vendorTheLongAndWindingTableAndColumnRef df.Behavior =VendorTheLongAndWindingTableAndColumnRefBhv_I
	VendorTheLongAndWindingTableAndColumnRefBhv_I.BaseBehavior.Behavior=&vendorTheLongAndWindingTableAndColumnRef
	VendorTheLongAndWindingTableAndColumnRefBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
	WithdrawalReasonBhv_I = new(WithdrawalReasonBhv)
	WithdrawalReasonBhv_I.BaseBehavior = new(df.BaseBehavior)
	WithdrawalReasonBhv_I.BaseBehavior.CreateBehaviorCommandInvoker()
	WithdrawalReasonBhv_I.BaseBehavior.TableDbName = "WithdrawalReason"
	var withdrawalReason df.Behavior =WithdrawalReasonBhv_I
	WithdrawalReasonBhv_I.BaseBehavior.Behavior=&withdrawalReason
	WithdrawalReasonBhv_I.BaseBehavior.BehaviorCommandInvoker = df.Bci
}
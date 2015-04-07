package entity

import (
	"github.com/mikeshimura/dbflute/df"
)

func EntityInit() {
	Member := func() *df.Entity {
		var te df.Entity = new(Member)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("Member", Member)
	MemberAddress := func() *df.Entity {
		var te df.Entity = new(MemberAddress)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("MemberAddress", MemberAddress)
	MemberLogin := func() *df.Entity {
		var te df.Entity = new(MemberLogin)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("MemberLogin", MemberLogin)
	MemberSecurity := func() *df.Entity {
		var te df.Entity = new(MemberSecurity)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("MemberSecurity", MemberSecurity)
	MemberService := func() *df.Entity {
		var te df.Entity = new(MemberService)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("MemberService", MemberService)
	MemberStatus := func() *df.Entity {
		var te df.Entity = new(MemberStatus)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("MemberStatus", MemberStatus)
	MemberWithdrawal := func() *df.Entity {
		var te df.Entity = new(MemberWithdrawal)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("MemberWithdrawal", MemberWithdrawal)
	Product := func() *df.Entity {
		var te df.Entity = new(Product)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("Product", Product)
	ProductCategory := func() *df.Entity {
		var te df.Entity = new(ProductCategory)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("ProductCategory", ProductCategory)
	ProductStatus := func() *df.Entity {
		var te df.Entity = new(ProductStatus)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("ProductStatus", ProductStatus)
	Purchase := func() *df.Entity {
		var te df.Entity = new(Purchase)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("Purchase", Purchase)
	PurchasePayment := func() *df.Entity {
		var te df.Entity = new(PurchasePayment)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("PurchasePayment", PurchasePayment)
	Region := func() *df.Entity {
		var te df.Entity = new(Region)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("Region", Region)
	ServiceRank := func() *df.Entity {
		var te df.Entity = new(ServiceRank)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("ServiceRank", ServiceRank)
	SummaryProduct := func() *df.Entity {
		var te df.Entity = new(SummaryProduct)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("SummaryProduct", SummaryProduct)
	SummaryWithdrawal := func() *df.Entity {
		var te df.Entity = new(SummaryWithdrawal)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("SummaryWithdrawal", SummaryWithdrawal)
	VendorConstraintNameAutoBar := func() *df.Entity {
		var te df.Entity = new(VendorConstraintNameAutoBar)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("VendorConstraintNameAutoBar", VendorConstraintNameAutoBar)
	VendorConstraintNameAutoFoo := func() *df.Entity {
		var te df.Entity = new(VendorConstraintNameAutoFoo)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("VendorConstraintNameAutoFoo", VendorConstraintNameAutoFoo)
	VendorConstraintNameAutoQux := func() *df.Entity {
		var te df.Entity = new(VendorConstraintNameAutoQux)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("VendorConstraintNameAutoQux", VendorConstraintNameAutoQux)
	VendorConstraintNameAutoRef := func() *df.Entity {
		var te df.Entity = new(VendorConstraintNameAutoRef)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("VendorConstraintNameAutoRef", VendorConstraintNameAutoRef)
	VendorLargeData := func() *df.Entity {
		var te df.Entity = new(VendorLargeData)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("VendorLargeData", VendorLargeData)
	VendorLargeDataRef := func() *df.Entity {
		var te df.Entity = new(VendorLargeDataRef)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("VendorLargeDataRef", VendorLargeDataRef)
	VendorTheLongAndWindingTableAndColumn := func() *df.Entity {
		var te df.Entity = new(VendorTheLongAndWindingTableAndColumn)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("VendorTheLongAndWindingTableAndColumn", VendorTheLongAndWindingTableAndColumn)
	VendorTheLongAndWindingTableAndColumnRef := func() *df.Entity {
		var te df.Entity = new(VendorTheLongAndWindingTableAndColumnRef)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("VendorTheLongAndWindingTableAndColumnRef", VendorTheLongAndWindingTableAndColumnRef)
	WithdrawalReason := func() *df.Entity {
		var te df.Entity = new(WithdrawalReason)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("WithdrawalReason", WithdrawalReason)
}
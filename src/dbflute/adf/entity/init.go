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
	WhiteDelimiter := func() *df.Entity {
		var te df.Entity = new(WhiteDelimiter)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("WhiteDelimiter", WhiteDelimiter)
	WithdrawalReason := func() *df.Entity {
		var te df.Entity = new(WithdrawalReason)
		te.SetUp()
		return &te
	}
	df.BhvUtil_I.AddEntity("WithdrawalReason", WithdrawalReason)
}
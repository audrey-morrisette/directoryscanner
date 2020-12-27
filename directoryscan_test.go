package directoryscanner_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/audrey-morrisette/directoryscanner"
)

// Test for generic sensitive data
var _ = Describe("Directoryscan", func() {
	Describe("Scanning a folder with 8 known findings", func() {
		Context("less than 100 MB in size", func() {
			It("should return 8 findings", func() {
				results, _ := directoryscanner.Scan("./secret")
				Expect(len(results)).To(Equal(8))
			})
		})
	})
})

// Test for credit cards
var _ = Describe("CreditCardScan", func() {
	Describe("Scanning a folder with 2 known Credit Card findings", func() {
		Context("less than 100 MB in size", func() {
			It("should return 2 findings", func() {
				results, _ := directoryscanner.Find("./secret", "Credit Card")
				Expect(len(results)).To(Equal(2))
			})
		})
	})
})

// Test for SSN
var _ = Describe("SSNScan", func() {
	Describe("Scanning a folder with 4 known SSN findings", func() {
		Context("less than 100 MB in size", func() {
			It("should return 4 findings", func() {
				results, _ := directoryscanner.Find("./secret", "SSN")
				Expect(len(results)).To(Equal(4))
			})
		})
	})
})

// Test for Word Password
var _ = Describe("PasswordScan", func() {
	Describe("Scanning a folder with 1 known \"Word Password\" finding", func() {
		Context("less than 100 MB in size", func() {
			It("should return 1 finding", func() {
				results, _ := directoryscanner.Find("./secret", "Word Password")
				Expect(len(results)).To(Equal(1))
			})
		})
	})
})

// Test for Word Username
var _ = Describe("UsernameScan", func() {
	Describe("Scanning a folder with 1 known \"Word Username\" finding", func() {
		Context("less than 100 MB in size", func() {
			It("should return 1 finding", func() {
				results, _ := directoryscanner.Find("./secret", "Word Username")
				Expect(len(results)).To(Equal(1))
			})
		})
	})
})

// Test for Multiscan
var _ = Describe("MultipleScan", func() {
	Describe("Scanning a folder with 7 known total findings of SSN, Word Username, and Credit Card types", func() {
		Context("less than 100 MB in size", func() {
			It("should return 7 findings", func() {
				results, _ := directoryscanner.Find("./secret", "SSN", "Word Username", "Credit Card")
				Expect(len(results)).To(Equal(7))
			})
		})
	})
})
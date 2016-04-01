package cpi_test

import (
	"github.com/tscolari/bosh-c3pi/cloud/fakes"
	"github.com/tscolari/bosh-c3pi/cpi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dispatcher", func() {
	var cpiClient *fakes.FakeCloud
	var subject *cpi.Dispatcher

	BeforeEach(func() {
		cpiClient = new(fakes.FakeCloud)
		subject = cpi.NewDispatcher(cpiClient, logger)
	})

	Describe("Dispatch", func() {
		Context("When the method name is invalid", func() {
			It("returns an error", func() {
				_, err := subject.Dispatch("invalid_method", nil)
				Expect(err).To(MatchError("Invalid cpi method: 'invalid_method'"))
			})
		})
	})
})

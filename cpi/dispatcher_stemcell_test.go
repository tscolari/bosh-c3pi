package cpi_test

import (
	"errors"

	"github.com/tscolari/bosh-c3pi/cloud"
	"github.com/tscolari/bosh-c3pi/cloud/fakes"
	"github.com/tscolari/bosh-c3pi/cpi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dispatcher", func() {
	var cpiClient *fakes.FakeCloud
	var subject *cpi.Dispatcher
	var methodName string
	var arguments cpi.RequestArguments

	BeforeEach(func() {
		cpiClient = new(fakes.FakeCloud)
		subject = cpi.NewDispatcher(cpiClient, logger)
	})

	Context("Stemcell Operations", func() {
		Describe("create_stemcell", func() {
			var cloudProperties cloud.CloudProperties

			BeforeEach(func() {
				methodName = "create_stemcell"

				cloudProperties = cloud.CloudProperties{"key": "value"}
				arguments = []interface{}{
					"image-path-1",
					cloudProperties,
				}
			})

			It("calls the correct method in the cloud object", func() {
				subject.Dispatch(methodName, arguments)
				Expect(cpiClient.CreateStemcellCallCount()).To(Equal(1))
			})

			It("calls the cloud object method with the correct arguments", func() {
				subject.Dispatch(methodName, arguments)
				imagePath, cloudProperties_ := cpiClient.CreateStemcellArgsForCall(0)
				Expect(imagePath).To(Equal("image-path-1"))
				Expect(cloudProperties_).To(Equal(cloudProperties))
			})

			It("returns no error", func() {
				_, err := subject.Dispatch(methodName, arguments)
				Expect(err).ToNot(HaveOccurred())
			})

			It("returns the stemcell id as result", func() {
				cpiClient.CreateStemcellReturns("stemcell-id-1", nil)
				result, _ := subject.Dispatch(methodName, arguments)
				Expect(result).To(Equal("stemcell-id-1"))
			})

			Context("when the cloud client fails", func() {
				var expectedError error

				BeforeEach(func() {
					expectedError = errors.New("failed here")
					cpiClient.CreateStemcellReturns("", expectedError)
				})

				It("returns the error", func() {
					_, err := subject.Dispatch(methodName, arguments)
					Expect(err).To(MatchError("failed here"))
				})
			})
		})

		Describe("delete_stemcell", func() {
			BeforeEach(func() {
				methodName = "delete_stemcell"
				arguments = []interface{}{
					"stemcell-id-1",
				}
			})

			It("calls the correct method in the cloud object", func() {
				subject.Dispatch(methodName, arguments)
				Expect(cpiClient.DeleteStemcellCallCount()).To(Equal(1))
			})

			It("calls the cloud object method with the correct arguments", func() {
				subject.Dispatch(methodName, arguments)
				stemcellID := cpiClient.DeleteStemcellArgsForCall(0)
				Expect(stemcellID).To(Equal("stemcell-id-1"))
			})

			It("returns no error", func() {
				_, err := subject.Dispatch(methodName, arguments)
				Expect(err).ToNot(HaveOccurred())
			})

			Context("when the cloud client fails", func() {
				var expectedError error

				BeforeEach(func() {
					expectedError = errors.New("failed here")
					cpiClient.DeleteStemcellReturns(expectedError)
				})

				It("returns the error", func() {
					_, err := subject.Dispatch(methodName, arguments)
					Expect(err).To(MatchError("failed here"))
				})
			})
		})
	})
})

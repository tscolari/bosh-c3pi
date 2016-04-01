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

	Context("Snapshot Operations", func() {
		Describe("snapshot_disk", func() {
			var metadata cloud.Metadata

			BeforeEach(func() {
				methodName = "snapshot_disk"

				metadata = cloud.Metadata{"key": "value"}
				arguments = []interface{}{
					"disk-id-1",
					metadata,
				}
			})

			It("calls the correct method in the cloud object", func() {
				subject.Dispatch(methodName, arguments)
				Expect(cpiClient.SnapshotDiskCallCount()).To(Equal(1))
			})

			It("calls the cloud object method with the correct arguments", func() {
				subject.Dispatch(methodName, arguments)
				diskID, metadata_ := cpiClient.SnapshotDiskArgsForCall(0)
				Expect(diskID).To(Equal("disk-id-1"))
				Expect(metadata_).To(Equal(metadata))
			})

			It("returns no error", func() {
				_, err := subject.Dispatch(methodName, arguments)
				Expect(err).ToNot(HaveOccurred())
			})

			It("returns the stemcell id as result", func() {
				cpiClient.SnapshotDiskReturns("snapshot-id-1", nil)
				result, _ := subject.Dispatch(methodName, arguments)
				Expect(result).To(Equal("snapshot-id-1"))
			})

			Context("when the cloud client fails", func() {
				var expectedError error

				BeforeEach(func() {
					expectedError = errors.New("failed here")
					cpiClient.SnapshotDiskReturns("", expectedError)
				})

				It("returns the error", func() {
					_, err := subject.Dispatch(methodName, arguments)
					Expect(err).To(MatchError("failed here"))
				})
			})
		})

		Describe("delete_snapshot", func() {
			BeforeEach(func() {
				methodName = "delete_snapshot"

				arguments = []interface{}{
					"snapshot-id-1",
				}
			})

			It("calls the correct method in the cloud object", func() {
				subject.Dispatch(methodName, arguments)
				Expect(cpiClient.DeleteSnapshotCallCount()).To(Equal(1))
			})

			It("calls the cloud object method with the correct arguments", func() {
				subject.Dispatch(methodName, arguments)
				stemcellID := cpiClient.DeleteSnapshotArgsForCall(0)
				Expect(stemcellID).To(Equal("snapshot-id-1"))
			})

			It("returns no error", func() {
				_, err := subject.Dispatch(methodName, arguments)
				Expect(err).ToNot(HaveOccurred())
			})

			Context("when the cloud client fails", func() {
				var expectedError error

				BeforeEach(func() {
					expectedError = errors.New("failed here")
					cpiClient.DeleteSnapshotReturns(expectedError)
				})

				It("returns the error", func() {
					_, err := subject.Dispatch(methodName, arguments)
					Expect(err).To(MatchError("failed here"))
				})
			})
		})
	})
})

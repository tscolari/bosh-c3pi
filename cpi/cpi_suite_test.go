package cpi_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

var logger boshlog.Logger

func TestCpi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cpi Suite")
}

var _ = BeforeEach(func() {
	logger = boshlog.NewWriterLogger(boshlog.LevelDebug, os.Stderr, os.Stderr)
})

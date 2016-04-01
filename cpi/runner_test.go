package cpi_test

import (
	"bytes"
	"encoding/json"
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tscolari/bosh-c3pi/cpi"
	"github.com/tscolari/bosh-c3pi/cpi/fakes"
)

var _ = Describe("Runner", func() {
	var dispatcher *fakes.FakeCpiDispatcher
	var runner *cpi.Runner

	var stdin *bytes.Buffer
	var stdout *bytes.Buffer

	BeforeEach(func() {
		dispatcher = new(fakes.FakeCpiDispatcher)
		runner = cpi.NewRunnerWithDispatcher(dispatcher, logger)
	})

	Describe("Run", func() {
		var input []byte

		BeforeEach(func() {
			input = []byte(`{"method":"cpi_test", "arguments":["arg1", "arg2", "arg3"]}   `)
			stdout = new(bytes.Buffer)
		})

		JustBeforeEach(func() {
			stdin = bytes.NewBuffer(input)
			runner.Run(stdin, stdout)
		})

		It("uses the dispatcher", func() {
			Expect(dispatcher.DispatchCallCount()).To(Equal(1))
		})

		It("sends the correct methodName to the dispatcher", func() {
			methodName, _ := dispatcher.DispatchArgsForCall(0)
			Expect(methodName).To(Equal("cpi_test"))
		})

		It("sends the correct arguments to the dispatcher", func() {
			_, arguments := dispatcher.DispatchArgsForCall(0)
			Expect(arguments).To(Equal(cpi.RequestArguments{"arg1", "arg2", "arg3"}))
		})

		Context("when the input is not a valid json", func() {
			BeforeEach(func() {
				input = []byte(`{"method":"cpi_test, "arguments":["arg1", "arg2", "arg3"]}`)
			})

			It("prints out a response object with the error", func() {
				var response cpi.Response
				err := json.NewDecoder(stdout).Decode(&response)
				Expect(err).ToNot(HaveOccurred())

				Expect(response.Error.Message).To(MatchRegexp("invalid character"))
			})
		})

		Context("Output", func() {
			BeforeEach(func() {
				dispatcher.DispatchReturns("dispatcher result", nil)
			})

			It("prints a response object with the dispatcher result", func() {
				var response cpi.Response
				err := json.NewDecoder(stdout).Decode(&response)
				Expect(err).ToNot(HaveOccurred())

				Expect(response.Result).To(Equal("dispatcher result"))
			})

			Context("when the dispatcher returns an error", func() {
				BeforeEach(func() {
					dispatcher.DispatchReturns(nil, errors.New("failed!"))
				})

				It("prints a response object with the dispatcher error", func() {
					var response cpi.Response
					err := json.NewDecoder(stdout).Decode(&response)
					Expect(err).ToNot(HaveOccurred())

					Expect(response.Error.Message).To(Equal("failed!"))
				})
			})
		})
	})
})

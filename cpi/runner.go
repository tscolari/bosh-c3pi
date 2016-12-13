package cpi

import (
	"encoding/json"
	"io"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/tscolari/bosh-c3pi/cloud"
)

type CpiDispatcher interface {
	Dispatch(methodName string, Arguments RequestArguments) (interface{}, error)
}

func NewRunner(client cloud.Cloud, logger boshlog.Logger) *Runner {
	dispatcher := NewDispatcher(client, logger)
	return NewRunnerWithDispatcher(dispatcher, logger)
}

func NewRunnerWithDispatcher(dispatcher CpiDispatcher, logger boshlog.Logger) *Runner {
	return &Runner{
		dispatcher: dispatcher,
		logger:     logger,
	}
}

type Runner struct {
	dispatcher CpiDispatcher
	logger     boshlog.Logger
}

func (r *Runner) Run(stdin io.Reader, stdout io.Writer) {
	var request Request

	err := json.NewDecoder(stdin).Decode(&request)
	if err != nil {
		r.printResponse(stdout, nil, err)
		return
	}

	r.logger.Info("runner", "Called with request: %#v", request)
	result, err := r.dispatcher.Dispatch(request.Method, request.Arguments)

	r.printResponse(stdout, result, err)
}

func (r *Runner) printResponse(output io.Writer, result interface{}, err error) {
	var responseError *ResponseError

	if err != nil {
		responseError = &ResponseError{
			Type:    "CPIError",
			Message: err.Error(),
		}
	}

	response := Response{
		Result: result,
		Error:  responseError,
	}

	json.NewEncoder(output).Encode(response)
}

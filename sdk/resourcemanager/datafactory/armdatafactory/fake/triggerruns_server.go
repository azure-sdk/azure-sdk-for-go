// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v9"
	"net/http"
	"net/url"
	"regexp"
)

// TriggerRunsServer is a fake server for instances of the armdatafactory.TriggerRunsClient type.
type TriggerRunsServer struct {
	// Cancel is the fake for method TriggerRunsClient.Cancel
	// HTTP status codes to indicate success: http.StatusOK
	Cancel func(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, runID string, options *armdatafactory.TriggerRunsClientCancelOptions) (resp azfake.Responder[armdatafactory.TriggerRunsClientCancelResponse], errResp azfake.ErrorResponder)

	// QueryByFactory is the fake for method TriggerRunsClient.QueryByFactory
	// HTTP status codes to indicate success: http.StatusOK
	QueryByFactory func(ctx context.Context, resourceGroupName string, factoryName string, filterParameters armdatafactory.RunFilterParameters, options *armdatafactory.TriggerRunsClientQueryByFactoryOptions) (resp azfake.Responder[armdatafactory.TriggerRunsClientQueryByFactoryResponse], errResp azfake.ErrorResponder)

	// Rerun is the fake for method TriggerRunsClient.Rerun
	// HTTP status codes to indicate success: http.StatusOK
	Rerun func(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, runID string, options *armdatafactory.TriggerRunsClientRerunOptions) (resp azfake.Responder[armdatafactory.TriggerRunsClientRerunResponse], errResp azfake.ErrorResponder)
}

// NewTriggerRunsServerTransport creates a new instance of TriggerRunsServerTransport with the provided implementation.
// The returned TriggerRunsServerTransport instance is connected to an instance of armdatafactory.TriggerRunsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTriggerRunsServerTransport(srv *TriggerRunsServer) *TriggerRunsServerTransport {
	return &TriggerRunsServerTransport{srv: srv}
}

// TriggerRunsServerTransport connects instances of armdatafactory.TriggerRunsClient to instances of TriggerRunsServer.
// Don't use this type directly, use NewTriggerRunsServerTransport instead.
type TriggerRunsServerTransport struct {
	srv *TriggerRunsServer
}

// Do implements the policy.Transporter interface for TriggerRunsServerTransport.
func (t *TriggerRunsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return t.dispatchToMethodFake(req, method)
}

func (t *TriggerRunsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if triggerRunsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = triggerRunsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "TriggerRunsClient.Cancel":
				res.resp, res.err = t.dispatchCancel(req)
			case "TriggerRunsClient.QueryByFactory":
				res.resp, res.err = t.dispatchQueryByFactory(req)
			case "TriggerRunsClient.Rerun":
				res.resp, res.err = t.dispatchRerun(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (t *TriggerRunsServerTransport) dispatchCancel(req *http.Request) (*http.Response, error) {
	if t.srv.Cancel == nil {
		return nil, &nonRetriableError{errors.New("fake for method Cancel not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggers/(?P<triggerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggerRuns/(?P<runId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cancel`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	triggerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("triggerName")])
	if err != nil {
		return nil, err
	}
	runIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("runId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Cancel(req.Context(), resourceGroupNameParam, factoryNameParam, triggerNameParam, runIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TriggerRunsServerTransport) dispatchQueryByFactory(req *http.Request) (*http.Response, error) {
	if t.srv.QueryByFactory == nil {
		return nil, &nonRetriableError{errors.New("fake for method QueryByFactory not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queryTriggerRuns`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdatafactory.RunFilterParameters](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.QueryByFactory(req.Context(), resourceGroupNameParam, factoryNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TriggerRunsQueryResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TriggerRunsServerTransport) dispatchRerun(req *http.Request) (*http.Response, error) {
	if t.srv.Rerun == nil {
		return nil, &nonRetriableError{errors.New("fake for method Rerun not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggers/(?P<triggerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggerRuns/(?P<runId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/rerun`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	triggerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("triggerName")])
	if err != nil {
		return nil, err
	}
	runIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("runId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Rerun(req.Context(), resourceGroupNameParam, factoryNameParam, triggerNameParam, runIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to TriggerRunsServerTransport
var triggerRunsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

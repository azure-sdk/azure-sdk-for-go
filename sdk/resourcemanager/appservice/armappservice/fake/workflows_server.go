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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v4"
	"net/http"
	"net/url"
	"regexp"
)

// WorkflowsServer is a fake server for instances of the armappservice.WorkflowsClient type.
type WorkflowsServer struct {
	// RegenerateAccessKey is the fake for method WorkflowsClient.RegenerateAccessKey
	// HTTP status codes to indicate success: http.StatusOK
	RegenerateAccessKey func(ctx context.Context, resourceGroupName string, name string, workflowName string, keyType armappservice.RegenerateActionParameter, options *armappservice.WorkflowsClientRegenerateAccessKeyOptions) (resp azfake.Responder[armappservice.WorkflowsClientRegenerateAccessKeyResponse], errResp azfake.ErrorResponder)

	// Validate is the fake for method WorkflowsClient.Validate
	// HTTP status codes to indicate success: http.StatusOK
	Validate func(ctx context.Context, resourceGroupName string, name string, workflowName string, validate armappservice.Workflow, options *armappservice.WorkflowsClientValidateOptions) (resp azfake.Responder[armappservice.WorkflowsClientValidateResponse], errResp azfake.ErrorResponder)
}

// NewWorkflowsServerTransport creates a new instance of WorkflowsServerTransport with the provided implementation.
// The returned WorkflowsServerTransport instance is connected to an instance of armappservice.WorkflowsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkflowsServerTransport(srv *WorkflowsServer) *WorkflowsServerTransport {
	return &WorkflowsServerTransport{srv: srv}
}

// WorkflowsServerTransport connects instances of armappservice.WorkflowsClient to instances of WorkflowsServer.
// Don't use this type directly, use NewWorkflowsServerTransport instead.
type WorkflowsServerTransport struct {
	srv *WorkflowsServer
}

// Do implements the policy.Transporter interface for WorkflowsServerTransport.
func (w *WorkflowsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return w.dispatchToMethodFake(req, method)
}

func (w *WorkflowsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if workflowsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = workflowsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "WorkflowsClient.RegenerateAccessKey":
				res.resp, res.err = w.dispatchRegenerateAccessKey(req)
			case "WorkflowsClient.Validate":
				res.resp, res.err = w.dispatchValidate(req)
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

func (w *WorkflowsServerTransport) dispatchRegenerateAccessKey(req *http.Request) (*http.Response, error) {
	if w.srv.RegenerateAccessKey == nil {
		return nil, &nonRetriableError{errors.New("fake for method RegenerateAccessKey not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/sites/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hostruntime/runtime/webhooks/workflow/api/management/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/regenerateAccessKey`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armappservice.RegenerateActionParameter](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.RegenerateAccessKey(req.Context(), resourceGroupNameParam, nameParam, workflowNameParam, body, nil)
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

func (w *WorkflowsServerTransport) dispatchValidate(req *http.Request) (*http.Response, error) {
	if w.srv.Validate == nil {
		return nil, &nonRetriableError{errors.New("fake for method Validate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/sites/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hostruntime/runtime/webhooks/workflow/api/management/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/validate`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armappservice.Workflow](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Validate(req.Context(), resourceGroupNameParam, nameParam, workflowNameParam, body, nil)
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

// set this to conditionally intercept incoming requests to WorkflowsServerTransport
var workflowsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

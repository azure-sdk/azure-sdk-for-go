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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// WorkflowRunsServer is a fake server for instances of the armappservice.WorkflowRunsClient type.
type WorkflowRunsServer struct {
	// Cancel is the fake for method WorkflowRunsClient.Cancel
	// HTTP status codes to indicate success: http.StatusOK
	Cancel func(ctx context.Context, resourceGroupName string, name string, workflowName string, runName string, options *armappservice.WorkflowRunsClientCancelOptions) (resp azfake.Responder[armappservice.WorkflowRunsClientCancelResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method WorkflowRunsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, name string, workflowName string, runName string, options *armappservice.WorkflowRunsClientGetOptions) (resp azfake.Responder[armappservice.WorkflowRunsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method WorkflowRunsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, name string, workflowName string, options *armappservice.WorkflowRunsClientListOptions) (resp azfake.PagerResponder[armappservice.WorkflowRunsClientListResponse])
}

// NewWorkflowRunsServerTransport creates a new instance of WorkflowRunsServerTransport with the provided implementation.
// The returned WorkflowRunsServerTransport instance is connected to an instance of armappservice.WorkflowRunsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkflowRunsServerTransport(srv *WorkflowRunsServer) *WorkflowRunsServerTransport {
	return &WorkflowRunsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armappservice.WorkflowRunsClientListResponse]](),
	}
}

// WorkflowRunsServerTransport connects instances of armappservice.WorkflowRunsClient to instances of WorkflowRunsServer.
// Don't use this type directly, use NewWorkflowRunsServerTransport instead.
type WorkflowRunsServerTransport struct {
	srv          *WorkflowRunsServer
	newListPager *tracker[azfake.PagerResponder[armappservice.WorkflowRunsClientListResponse]]
}

// Do implements the policy.Transporter interface for WorkflowRunsServerTransport.
func (w *WorkflowRunsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return w.dispatchToMethodFake(req, method)
}

func (w *WorkflowRunsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if workflowRunsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = workflowRunsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "WorkflowRunsClient.Cancel":
				res.resp, res.err = w.dispatchCancel(req)
			case "WorkflowRunsClient.Get":
				res.resp, res.err = w.dispatchGet(req)
			case "WorkflowRunsClient.NewListPager":
				res.resp, res.err = w.dispatchNewListPager(req)
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

func (w *WorkflowRunsServerTransport) dispatchCancel(req *http.Request) (*http.Response, error) {
	if w.srv.Cancel == nil {
		return nil, &nonRetriableError{errors.New("fake for method Cancel not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/sites/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hostruntime/runtime/webhooks/workflow/api/management/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runs/(?P<runName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cancel`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
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
	runNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("runName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Cancel(req.Context(), resourceGroupNameParam, nameParam, workflowNameParam, runNameParam, nil)
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

func (w *WorkflowRunsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/sites/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hostruntime/runtime/webhooks/workflow/api/management/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runs/(?P<runName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
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
	runNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("runName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, nameParam, workflowNameParam, runNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WorkflowRun, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkflowRunsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := w.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/sites/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hostruntime/runtime/webhooks/workflow/api/management/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runs`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
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
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armappservice.WorkflowRunsClientListOptions
		if topParam != nil || filterParam != nil {
			options = &armappservice.WorkflowRunsClientListOptions{
				Top:    topParam,
				Filter: filterParam,
			}
		}
		resp := w.srv.NewListPager(resourceGroupNameParam, nameParam, workflowNameParam, options)
		newListPager = &resp
		w.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armappservice.WorkflowRunsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		w.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to WorkflowRunsServerTransport
var workflowRunsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

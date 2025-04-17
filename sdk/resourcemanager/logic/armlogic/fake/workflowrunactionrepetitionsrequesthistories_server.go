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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
	"net/http"
	"net/url"
	"regexp"
)

// WorkflowRunActionRepetitionsRequestHistoriesServer is a fake server for instances of the armlogic.WorkflowRunActionRepetitionsRequestHistoriesClient type.
type WorkflowRunActionRepetitionsRequestHistoriesServer struct {
	// Get is the fake for method WorkflowRunActionRepetitionsRequestHistoriesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, repetitionName string, requestHistoryName string, options *armlogic.WorkflowRunActionRepetitionsRequestHistoriesClientGetOptions) (resp azfake.Responder[armlogic.WorkflowRunActionRepetitionsRequestHistoriesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method WorkflowRunActionRepetitionsRequestHistoriesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workflowName string, runName string, actionName string, repetitionName string, options *armlogic.WorkflowRunActionRepetitionsRequestHistoriesClientListOptions) (resp azfake.PagerResponder[armlogic.WorkflowRunActionRepetitionsRequestHistoriesClientListResponse])
}

// NewWorkflowRunActionRepetitionsRequestHistoriesServerTransport creates a new instance of WorkflowRunActionRepetitionsRequestHistoriesServerTransport with the provided implementation.
// The returned WorkflowRunActionRepetitionsRequestHistoriesServerTransport instance is connected to an instance of armlogic.WorkflowRunActionRepetitionsRequestHistoriesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkflowRunActionRepetitionsRequestHistoriesServerTransport(srv *WorkflowRunActionRepetitionsRequestHistoriesServer) *WorkflowRunActionRepetitionsRequestHistoriesServerTransport {
	return &WorkflowRunActionRepetitionsRequestHistoriesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armlogic.WorkflowRunActionRepetitionsRequestHistoriesClientListResponse]](),
	}
}

// WorkflowRunActionRepetitionsRequestHistoriesServerTransport connects instances of armlogic.WorkflowRunActionRepetitionsRequestHistoriesClient to instances of WorkflowRunActionRepetitionsRequestHistoriesServer.
// Don't use this type directly, use NewWorkflowRunActionRepetitionsRequestHistoriesServerTransport instead.
type WorkflowRunActionRepetitionsRequestHistoriesServerTransport struct {
	srv          *WorkflowRunActionRepetitionsRequestHistoriesServer
	newListPager *tracker[azfake.PagerResponder[armlogic.WorkflowRunActionRepetitionsRequestHistoriesClientListResponse]]
}

// Do implements the policy.Transporter interface for WorkflowRunActionRepetitionsRequestHistoriesServerTransport.
func (w *WorkflowRunActionRepetitionsRequestHistoriesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return w.dispatchToMethodFake(req, method)
}

func (w *WorkflowRunActionRepetitionsRequestHistoriesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if workflowRunActionRepetitionsRequestHistoriesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = workflowRunActionRepetitionsRequestHistoriesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "WorkflowRunActionRepetitionsRequestHistoriesClient.Get":
				res.resp, res.err = w.dispatchGet(req)
			case "WorkflowRunActionRepetitionsRequestHistoriesClient.NewListPager":
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

func (w *WorkflowRunActionRepetitionsRequestHistoriesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runs/(?P<runName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/actions/(?P<actionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/repetitions/(?P<repetitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/requestHistories/(?P<requestHistoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 7 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
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
	actionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("actionName")])
	if err != nil {
		return nil, err
	}
	repetitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("repetitionName")])
	if err != nil {
		return nil, err
	}
	requestHistoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("requestHistoryName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, workflowNameParam, runNameParam, actionNameParam, repetitionNameParam, requestHistoryNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RequestHistory, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkflowRunActionRepetitionsRequestHistoriesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := w.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runs/(?P<runName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/actions/(?P<actionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/repetitions/(?P<repetitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/requestHistories`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
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
		actionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("actionName")])
		if err != nil {
			return nil, err
		}
		repetitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("repetitionName")])
		if err != nil {
			return nil, err
		}
		resp := w.srv.NewListPager(resourceGroupNameParam, workflowNameParam, runNameParam, actionNameParam, repetitionNameParam, nil)
		newListPager = &resp
		w.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armlogic.WorkflowRunActionRepetitionsRequestHistoriesClientListResponse, createLink func() string) {
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

// set this to conditionally intercept incoming requests to WorkflowRunActionRepetitionsRequestHistoriesServerTransport
var workflowRunActionRepetitionsRequestHistoriesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ExascaleDbNodesServer is a fake server for instances of the armoracledatabase.ExascaleDbNodesClient type.
type ExascaleDbNodesServer struct {
	// BeginAction is the fake for method ExascaleDbNodesClient.BeginAction
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginAction func(ctx context.Context, resourceGroupName string, exadbVMClusterName string, exascaleDbNodeName string, body armoracledatabase.DbNodeAction, options *armoracledatabase.ExascaleDbNodesClientBeginActionOptions) (resp azfake.PollerResponder[armoracledatabase.ExascaleDbNodesClientActionResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ExascaleDbNodesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, exadbVMClusterName string, exascaleDbNodeName string, options *armoracledatabase.ExascaleDbNodesClientGetOptions) (resp azfake.Responder[armoracledatabase.ExascaleDbNodesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByParentPager is the fake for method ExascaleDbNodesClient.NewListByParentPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByParentPager func(resourceGroupName string, exadbVMClusterName string, options *armoracledatabase.ExascaleDbNodesClientListByParentOptions) (resp azfake.PagerResponder[armoracledatabase.ExascaleDbNodesClientListByParentResponse])
}

// NewExascaleDbNodesServerTransport creates a new instance of ExascaleDbNodesServerTransport with the provided implementation.
// The returned ExascaleDbNodesServerTransport instance is connected to an instance of armoracledatabase.ExascaleDbNodesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExascaleDbNodesServerTransport(srv *ExascaleDbNodesServer) *ExascaleDbNodesServerTransport {
	return &ExascaleDbNodesServerTransport{
		srv:                  srv,
		beginAction:          newTracker[azfake.PollerResponder[armoracledatabase.ExascaleDbNodesClientActionResponse]](),
		newListByParentPager: newTracker[azfake.PagerResponder[armoracledatabase.ExascaleDbNodesClientListByParentResponse]](),
	}
}

// ExascaleDbNodesServerTransport connects instances of armoracledatabase.ExascaleDbNodesClient to instances of ExascaleDbNodesServer.
// Don't use this type directly, use NewExascaleDbNodesServerTransport instead.
type ExascaleDbNodesServerTransport struct {
	srv                  *ExascaleDbNodesServer
	beginAction          *tracker[azfake.PollerResponder[armoracledatabase.ExascaleDbNodesClientActionResponse]]
	newListByParentPager *tracker[azfake.PagerResponder[armoracledatabase.ExascaleDbNodesClientListByParentResponse]]
}

// Do implements the policy.Transporter interface for ExascaleDbNodesServerTransport.
func (e *ExascaleDbNodesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return e.dispatchToMethodFake(req, method)
}

func (e *ExascaleDbNodesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if exascaleDbNodesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = exascaleDbNodesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ExascaleDbNodesClient.BeginAction":
				res.resp, res.err = e.dispatchBeginAction(req)
			case "ExascaleDbNodesClient.Get":
				res.resp, res.err = e.dispatchGet(req)
			case "ExascaleDbNodesClient.NewListByParentPager":
				res.resp, res.err = e.dispatchNewListByParentPager(req)
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

func (e *ExascaleDbNodesServerTransport) dispatchBeginAction(req *http.Request) (*http.Response, error) {
	if e.srv.BeginAction == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginAction not implemented")}
	}
	beginAction := e.beginAction.get(req)
	if beginAction == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/exadbVmClusters/(?P<exadbVmClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dbNodes/(?P<exascaleDbNodeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/action`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armoracledatabase.DbNodeAction](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		exadbVMClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("exadbVmClusterName")])
		if err != nil {
			return nil, err
		}
		exascaleDbNodeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("exascaleDbNodeName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginAction(req.Context(), resourceGroupNameParam, exadbVMClusterNameParam, exascaleDbNodeNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginAction = &respr
		e.beginAction.add(req, beginAction)
	}

	resp, err := server.PollerResponderNext(beginAction, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		e.beginAction.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginAction) {
		e.beginAction.remove(req)
	}

	return resp, nil
}

func (e *ExascaleDbNodesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/exadbVmClusters/(?P<exadbVmClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dbNodes/(?P<exascaleDbNodeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	exadbVMClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("exadbVmClusterName")])
	if err != nil {
		return nil, err
	}
	exascaleDbNodeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("exascaleDbNodeName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameParam, exadbVMClusterNameParam, exascaleDbNodeNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExascaleDbNode, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExascaleDbNodesServerTransport) dispatchNewListByParentPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListByParentPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByParentPager not implemented")}
	}
	newListByParentPager := e.newListByParentPager.get(req)
	if newListByParentPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/exadbVmClusters/(?P<exadbVmClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dbNodes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		exadbVMClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("exadbVmClusterName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListByParentPager(resourceGroupNameParam, exadbVMClusterNameParam, nil)
		newListByParentPager = &resp
		e.newListByParentPager.add(req, newListByParentPager)
		server.PagerResponderInjectNextLinks(newListByParentPager, req, func(page *armoracledatabase.ExascaleDbNodesClientListByParentResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByParentPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListByParentPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByParentPager) {
		e.newListByParentPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ExascaleDbNodesServerTransport
var exascaleDbNodesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

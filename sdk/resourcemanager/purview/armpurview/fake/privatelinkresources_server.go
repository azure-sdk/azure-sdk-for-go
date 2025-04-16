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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateLinkResourcesServer is a fake server for instances of the armpurview.PrivateLinkResourcesClient type.
type PrivateLinkResourcesServer struct {
	// GetByGroupID is the fake for method PrivateLinkResourcesClient.GetByGroupID
	// HTTP status codes to indicate success: http.StatusOK
	GetByGroupID func(ctx context.Context, resourceGroupName string, accountName string, groupID string, options *armpurview.PrivateLinkResourcesClientGetByGroupIDOptions) (resp azfake.Responder[armpurview.PrivateLinkResourcesClientGetByGroupIDResponse], errResp azfake.ErrorResponder)

	// NewListByAccountPager is the fake for method PrivateLinkResourcesClient.NewListByAccountPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByAccountPager func(resourceGroupName string, accountName string, options *armpurview.PrivateLinkResourcesClientListByAccountOptions) (resp azfake.PagerResponder[armpurview.PrivateLinkResourcesClientListByAccountResponse])
}

// NewPrivateLinkResourcesServerTransport creates a new instance of PrivateLinkResourcesServerTransport with the provided implementation.
// The returned PrivateLinkResourcesServerTransport instance is connected to an instance of armpurview.PrivateLinkResourcesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateLinkResourcesServerTransport(srv *PrivateLinkResourcesServer) *PrivateLinkResourcesServerTransport {
	return &PrivateLinkResourcesServerTransport{
		srv:                   srv,
		newListByAccountPager: newTracker[azfake.PagerResponder[armpurview.PrivateLinkResourcesClientListByAccountResponse]](),
	}
}

// PrivateLinkResourcesServerTransport connects instances of armpurview.PrivateLinkResourcesClient to instances of PrivateLinkResourcesServer.
// Don't use this type directly, use NewPrivateLinkResourcesServerTransport instead.
type PrivateLinkResourcesServerTransport struct {
	srv                   *PrivateLinkResourcesServer
	newListByAccountPager *tracker[azfake.PagerResponder[armpurview.PrivateLinkResourcesClientListByAccountResponse]]
}

// Do implements the policy.Transporter interface for PrivateLinkResourcesServerTransport.
func (p *PrivateLinkResourcesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PrivateLinkResourcesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if privateLinkResourcesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = privateLinkResourcesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "PrivateLinkResourcesClient.GetByGroupID":
				res.resp, res.err = p.dispatchGetByGroupID(req)
			case "PrivateLinkResourcesClient.NewListByAccountPager":
				res.resp, res.err = p.dispatchNewListByAccountPager(req)
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

func (p *PrivateLinkResourcesServerTransport) dispatchGetByGroupID(req *http.Request) (*http.Response, error) {
	if p.srv.GetByGroupID == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByGroupID not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Purview/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources/(?P<groupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	groupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.GetByGroupID(req.Context(), resourceGroupNameParam, accountNameParam, groupIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateLinkResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateLinkResourcesServerTransport) dispatchNewListByAccountPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByAccountPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByAccountPager not implemented")}
	}
	newListByAccountPager := p.newListByAccountPager.get(req)
	if newListByAccountPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Purview/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListByAccountPager(resourceGroupNameParam, accountNameParam, nil)
		newListByAccountPager = &resp
		p.newListByAccountPager.add(req, newListByAccountPager)
		server.PagerResponderInjectNextLinks(newListByAccountPager, req, func(page *armpurview.PrivateLinkResourcesClientListByAccountResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByAccountPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByAccountPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByAccountPager) {
		p.newListByAccountPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to PrivateLinkResourcesServerTransport
var privateLinkResourcesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

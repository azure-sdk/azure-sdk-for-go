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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ActiveSessionHostConfigurationsServer is a fake server for instances of the armdesktopvirtualization.ActiveSessionHostConfigurationsClient type.
type ActiveSessionHostConfigurationsServer struct {
	// Get is the fake for method ActiveSessionHostConfigurationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, hostPoolName string, options *armdesktopvirtualization.ActiveSessionHostConfigurationsClientGetOptions) (resp azfake.Responder[armdesktopvirtualization.ActiveSessionHostConfigurationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByHostPoolPager is the fake for method ActiveSessionHostConfigurationsClient.NewListByHostPoolPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByHostPoolPager func(resourceGroupName string, hostPoolName string, options *armdesktopvirtualization.ActiveSessionHostConfigurationsClientListByHostPoolOptions) (resp azfake.PagerResponder[armdesktopvirtualization.ActiveSessionHostConfigurationsClientListByHostPoolResponse])
}

// NewActiveSessionHostConfigurationsServerTransport creates a new instance of ActiveSessionHostConfigurationsServerTransport with the provided implementation.
// The returned ActiveSessionHostConfigurationsServerTransport instance is connected to an instance of armdesktopvirtualization.ActiveSessionHostConfigurationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewActiveSessionHostConfigurationsServerTransport(srv *ActiveSessionHostConfigurationsServer) *ActiveSessionHostConfigurationsServerTransport {
	return &ActiveSessionHostConfigurationsServerTransport{
		srv:                    srv,
		newListByHostPoolPager: newTracker[azfake.PagerResponder[armdesktopvirtualization.ActiveSessionHostConfigurationsClientListByHostPoolResponse]](),
	}
}

// ActiveSessionHostConfigurationsServerTransport connects instances of armdesktopvirtualization.ActiveSessionHostConfigurationsClient to instances of ActiveSessionHostConfigurationsServer.
// Don't use this type directly, use NewActiveSessionHostConfigurationsServerTransport instead.
type ActiveSessionHostConfigurationsServerTransport struct {
	srv                    *ActiveSessionHostConfigurationsServer
	newListByHostPoolPager *tracker[azfake.PagerResponder[armdesktopvirtualization.ActiveSessionHostConfigurationsClientListByHostPoolResponse]]
}

// Do implements the policy.Transporter interface for ActiveSessionHostConfigurationsServerTransport.
func (a *ActiveSessionHostConfigurationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *ActiveSessionHostConfigurationsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if activeSessionHostConfigurationsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = activeSessionHostConfigurationsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ActiveSessionHostConfigurationsClient.Get":
				res.resp, res.err = a.dispatchGet(req)
			case "ActiveSessionHostConfigurationsClient.NewListByHostPoolPager":
				res.resp, res.err = a.dispatchNewListByHostPoolPager(req)
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

func (a *ActiveSessionHostConfigurationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/activeSessionHostConfigurations/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hostPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hostPoolName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, hostPoolNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ActiveSessionHostConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ActiveSessionHostConfigurationsServerTransport) dispatchNewListByHostPoolPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByHostPoolPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByHostPoolPager not implemented")}
	}
	newListByHostPoolPager := a.newListByHostPoolPager.get(req)
	if newListByHostPoolPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/activeSessionHostConfigurations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hostPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hostPoolName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByHostPoolPager(resourceGroupNameParam, hostPoolNameParam, nil)
		newListByHostPoolPager = &resp
		a.newListByHostPoolPager.add(req, newListByHostPoolPager)
		server.PagerResponderInjectNextLinks(newListByHostPoolPager, req, func(page *armdesktopvirtualization.ActiveSessionHostConfigurationsClientListByHostPoolResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByHostPoolPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByHostPoolPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByHostPoolPager) {
		a.newListByHostPoolPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ActiveSessionHostConfigurationsServerTransport
var activeSessionHostConfigurationsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

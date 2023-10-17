//go:build go1.18
// +build go1.18

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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"net/http"
	"net/url"
	"regexp"
)

// ExpressRoutePortAuthorizationsServer is a fake server for instances of the armnetwork.ExpressRoutePortAuthorizationsClient type.
type ExpressRoutePortAuthorizationsServer struct {
	// BeginCreateOrUpdate is the fake for method ExpressRoutePortAuthorizationsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, expressRoutePortName string, authorizationName string, authorizationParameters armnetwork.ExpressRoutePortAuthorization, options *armnetwork.ExpressRoutePortAuthorizationsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.ExpressRoutePortAuthorizationsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ExpressRoutePortAuthorizationsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, expressRoutePortName string, authorizationName string, options *armnetwork.ExpressRoutePortAuthorizationsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.ExpressRoutePortAuthorizationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ExpressRoutePortAuthorizationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, expressRoutePortName string, authorizationName string, options *armnetwork.ExpressRoutePortAuthorizationsClientGetOptions) (resp azfake.Responder[armnetwork.ExpressRoutePortAuthorizationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ExpressRoutePortAuthorizationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, expressRoutePortName string, options *armnetwork.ExpressRoutePortAuthorizationsClientListOptions) (resp azfake.PagerResponder[armnetwork.ExpressRoutePortAuthorizationsClientListResponse])
}

// NewExpressRoutePortAuthorizationsServerTransport creates a new instance of ExpressRoutePortAuthorizationsServerTransport with the provided implementation.
// The returned ExpressRoutePortAuthorizationsServerTransport instance is connected to an instance of armnetwork.ExpressRoutePortAuthorizationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExpressRoutePortAuthorizationsServerTransport(srv *ExpressRoutePortAuthorizationsServer) *ExpressRoutePortAuthorizationsServerTransport {
	return &ExpressRoutePortAuthorizationsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnetwork.ExpressRoutePortAuthorizationsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armnetwork.ExpressRoutePortAuthorizationsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armnetwork.ExpressRoutePortAuthorizationsClientListResponse]](),
	}
}

// ExpressRoutePortAuthorizationsServerTransport connects instances of armnetwork.ExpressRoutePortAuthorizationsClient to instances of ExpressRoutePortAuthorizationsServer.
// Don't use this type directly, use NewExpressRoutePortAuthorizationsServerTransport instead.
type ExpressRoutePortAuthorizationsServerTransport struct {
	srv                 *ExpressRoutePortAuthorizationsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnetwork.ExpressRoutePortAuthorizationsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armnetwork.ExpressRoutePortAuthorizationsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armnetwork.ExpressRoutePortAuthorizationsClientListResponse]]
}

// Do implements the policy.Transporter interface for ExpressRoutePortAuthorizationsServerTransport.
func (e *ExpressRoutePortAuthorizationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ExpressRoutePortAuthorizationsClient.BeginCreateOrUpdate":
		resp, err = e.dispatchBeginCreateOrUpdate(req)
	case "ExpressRoutePortAuthorizationsClient.BeginDelete":
		resp, err = e.dispatchBeginDelete(req)
	case "ExpressRoutePortAuthorizationsClient.Get":
		resp, err = e.dispatchGet(req)
	case "ExpressRoutePortAuthorizationsClient.NewListPager":
		resp, err = e.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *ExpressRoutePortAuthorizationsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := e.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRoutePorts/(?P<expressRoutePortName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizations/(?P<authorizationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.ExpressRoutePortAuthorization](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		expressRoutePortNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("expressRoutePortName")])
		if err != nil {
			return nil, err
		}
		authorizationNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, expressRoutePortNameUnescaped, authorizationNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		e.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		e.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		e.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (e *ExpressRoutePortAuthorizationsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if e.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := e.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRoutePorts/(?P<expressRoutePortName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizations/(?P<authorizationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		expressRoutePortNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("expressRoutePortName")])
		if err != nil {
			return nil, err
		}
		authorizationNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, expressRoutePortNameUnescaped, authorizationNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		e.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		e.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		e.beginDelete.remove(req)
	}

	return resp, nil
}

func (e *ExpressRoutePortAuthorizationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRoutePorts/(?P<expressRoutePortName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizations/(?P<authorizationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	expressRoutePortNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("expressRoutePortName")])
	if err != nil {
		return nil, err
	}
	authorizationNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameUnescaped, expressRoutePortNameUnescaped, authorizationNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExpressRoutePortAuthorization, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExpressRoutePortAuthorizationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := e.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/expressRoutePorts/(?P<expressRoutePortName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		expressRoutePortNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("expressRoutePortName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListPager(resourceGroupNameUnescaped, expressRoutePortNameUnescaped, nil)
		newListPager = &resp
		e.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.ExpressRoutePortAuthorizationsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		e.newListPager.remove(req)
	}
	return resp, nil
}

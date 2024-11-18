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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh"
	"net/http"
	"net/url"
	"regexp"
)

// GatewayServer is a fake server for instances of the armservicefabricmesh.GatewayClient type.
type GatewayServer struct {
	// Create is the fake for method GatewayClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	Create func(ctx context.Context, resourceGroupName string, gatewayResourceName string, gatewayResourceDescription armservicefabricmesh.GatewayResourceDescription, options *armservicefabricmesh.GatewayClientCreateOptions) (resp azfake.Responder[armservicefabricmesh.GatewayClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method GatewayClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, gatewayResourceName string, options *armservicefabricmesh.GatewayClientDeleteOptions) (resp azfake.Responder[armservicefabricmesh.GatewayClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method GatewayClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, gatewayResourceName string, options *armservicefabricmesh.GatewayClientGetOptions) (resp azfake.Responder[armservicefabricmesh.GatewayClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method GatewayClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armservicefabricmesh.GatewayClientListByResourceGroupOptions) (resp azfake.PagerResponder[armservicefabricmesh.GatewayClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method GatewayClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armservicefabricmesh.GatewayClientListBySubscriptionOptions) (resp azfake.PagerResponder[armservicefabricmesh.GatewayClientListBySubscriptionResponse])
}

// NewGatewayServerTransport creates a new instance of GatewayServerTransport with the provided implementation.
// The returned GatewayServerTransport instance is connected to an instance of armservicefabricmesh.GatewayClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGatewayServerTransport(srv *GatewayServer) *GatewayServerTransport {
	return &GatewayServerTransport{
		srv:                         srv,
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armservicefabricmesh.GatewayClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armservicefabricmesh.GatewayClientListBySubscriptionResponse]](),
	}
}

// GatewayServerTransport connects instances of armservicefabricmesh.GatewayClient to instances of GatewayServer.
// Don't use this type directly, use NewGatewayServerTransport instead.
type GatewayServerTransport struct {
	srv                         *GatewayServer
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armservicefabricmesh.GatewayClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armservicefabricmesh.GatewayClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for GatewayServerTransport.
func (g *GatewayServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GatewayClient.Create":
		resp, err = g.dispatchCreate(req)
	case "GatewayClient.Delete":
		resp, err = g.dispatchDelete(req)
	case "GatewayClient.Get":
		resp, err = g.dispatchGet(req)
	case "GatewayClient.NewListByResourceGroupPager":
		resp, err = g.dispatchNewListByResourceGroupPager(req)
	case "GatewayClient.NewListBySubscriptionPager":
		resp, err = g.dispatchNewListBySubscriptionPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GatewayServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if g.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceFabricMesh/gateways/(?P<gatewayResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armservicefabricmesh.GatewayResourceDescription](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	gatewayResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Create(req.Context(), resourceGroupNameParam, gatewayResourceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GatewayResourceDescription, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GatewayServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if g.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceFabricMesh/gateways/(?P<gatewayResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	gatewayResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Delete(req.Context(), resourceGroupNameParam, gatewayResourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GatewayServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if g.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceFabricMesh/gateways/(?P<gatewayResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	gatewayResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Get(req.Context(), resourceGroupNameParam, gatewayResourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GatewayResourceDescription, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GatewayServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := g.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceFabricMesh/gateways`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := g.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		g.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armservicefabricmesh.GatewayClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		g.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (g *GatewayServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := g.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceFabricMesh/gateways`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := g.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		g.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armservicefabricmesh.GatewayClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		g.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

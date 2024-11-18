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

// NetworkServer is a fake server for instances of the armservicefabricmesh.NetworkClient type.
type NetworkServer struct {
	// Create is the fake for method NetworkClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	Create func(ctx context.Context, resourceGroupName string, networkResourceName string, networkResourceDescription armservicefabricmesh.NetworkResourceDescription, options *armservicefabricmesh.NetworkClientCreateOptions) (resp azfake.Responder[armservicefabricmesh.NetworkClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method NetworkClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, networkResourceName string, options *armservicefabricmesh.NetworkClientDeleteOptions) (resp azfake.Responder[armservicefabricmesh.NetworkClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method NetworkClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkResourceName string, options *armservicefabricmesh.NetworkClientGetOptions) (resp azfake.Responder[armservicefabricmesh.NetworkClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method NetworkClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armservicefabricmesh.NetworkClientListByResourceGroupOptions) (resp azfake.PagerResponder[armservicefabricmesh.NetworkClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method NetworkClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armservicefabricmesh.NetworkClientListBySubscriptionOptions) (resp azfake.PagerResponder[armservicefabricmesh.NetworkClientListBySubscriptionResponse])
}

// NewNetworkServerTransport creates a new instance of NetworkServerTransport with the provided implementation.
// The returned NetworkServerTransport instance is connected to an instance of armservicefabricmesh.NetworkClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNetworkServerTransport(srv *NetworkServer) *NetworkServerTransport {
	return &NetworkServerTransport{
		srv:                         srv,
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armservicefabricmesh.NetworkClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armservicefabricmesh.NetworkClientListBySubscriptionResponse]](),
	}
}

// NetworkServerTransport connects instances of armservicefabricmesh.NetworkClient to instances of NetworkServer.
// Don't use this type directly, use NewNetworkServerTransport instead.
type NetworkServerTransport struct {
	srv                         *NetworkServer
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armservicefabricmesh.NetworkClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armservicefabricmesh.NetworkClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for NetworkServerTransport.
func (n *NetworkServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "NetworkClient.Create":
		resp, err = n.dispatchCreate(req)
	case "NetworkClient.Delete":
		resp, err = n.dispatchDelete(req)
	case "NetworkClient.Get":
		resp, err = n.dispatchGet(req)
	case "NetworkClient.NewListByResourceGroupPager":
		resp, err = n.dispatchNewListByResourceGroupPager(req)
	case "NetworkClient.NewListBySubscriptionPager":
		resp, err = n.dispatchNewListBySubscriptionPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NetworkServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if n.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceFabricMesh/networks/(?P<networkResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armservicefabricmesh.NetworkResourceDescription](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Create(req.Context(), resourceGroupNameParam, networkResourceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NetworkResourceDescription, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NetworkServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if n.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceFabricMesh/networks/(?P<networkResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Delete(req.Context(), resourceGroupNameParam, networkResourceNameParam, nil)
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

func (n *NetworkServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if n.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceFabricMesh/networks/(?P<networkResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Get(req.Context(), resourceGroupNameParam, networkResourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NetworkResourceDescription, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NetworkServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := n.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceFabricMesh/networks`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := n.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		n.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armservicefabricmesh.NetworkClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		n.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (n *NetworkServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := n.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceFabricMesh/networks`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := n.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		n.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armservicefabricmesh.NetworkClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		n.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}
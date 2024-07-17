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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
	"net/http"
	"net/url"
	"regexp"
)

// VirtualApplianceConnectionsServer is a fake server for instances of the armnetwork.VirtualApplianceConnectionsClient type.
type VirtualApplianceConnectionsServer struct {
	// BeginCreateOrUpdate is the fake for method VirtualApplianceConnectionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, connectionName string, networkVirtualApplianceConnectionParameters armnetwork.VirtualApplianceConnection, options *armnetwork.VirtualApplianceConnectionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.VirtualApplianceConnectionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualApplianceConnectionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, connectionName string, options *armnetwork.VirtualApplianceConnectionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.VirtualApplianceConnectionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualApplianceConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, connectionName string, options *armnetwork.VirtualApplianceConnectionsClientGetOptions) (resp azfake.Responder[armnetwork.VirtualApplianceConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VirtualApplianceConnectionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkVirtualApplianceName string, options *armnetwork.VirtualApplianceConnectionsClientListOptions) (resp azfake.PagerResponder[armnetwork.VirtualApplianceConnectionsClientListResponse])
}

// NewVirtualApplianceConnectionsServerTransport creates a new instance of VirtualApplianceConnectionsServerTransport with the provided implementation.
// The returned VirtualApplianceConnectionsServerTransport instance is connected to an instance of armnetwork.VirtualApplianceConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualApplianceConnectionsServerTransport(srv *VirtualApplianceConnectionsServer) *VirtualApplianceConnectionsServerTransport {
	return &VirtualApplianceConnectionsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnetwork.VirtualApplianceConnectionsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armnetwork.VirtualApplianceConnectionsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armnetwork.VirtualApplianceConnectionsClientListResponse]](),
	}
}

// VirtualApplianceConnectionsServerTransport connects instances of armnetwork.VirtualApplianceConnectionsClient to instances of VirtualApplianceConnectionsServer.
// Don't use this type directly, use NewVirtualApplianceConnectionsServerTransport instead.
type VirtualApplianceConnectionsServerTransport struct {
	srv                 *VirtualApplianceConnectionsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnetwork.VirtualApplianceConnectionsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armnetwork.VirtualApplianceConnectionsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armnetwork.VirtualApplianceConnectionsClientListResponse]]
}

// Do implements the policy.Transporter interface for VirtualApplianceConnectionsServerTransport.
func (v *VirtualApplianceConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualApplianceConnectionsClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VirtualApplianceConnectionsClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VirtualApplianceConnectionsClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualApplianceConnectionsClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualApplianceConnectionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := v.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkVirtualAppliances/(?P<networkVirtualApplianceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkVirtualApplianceConnections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VirtualApplianceConnection](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkVirtualApplianceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkVirtualApplianceName")])
		if err != nil {
			return nil, err
		}
		connectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, networkVirtualApplianceNameParam, connectionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		v.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		v.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		v.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (v *VirtualApplianceConnectionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkVirtualAppliances/(?P<networkVirtualApplianceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkVirtualApplianceConnections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkVirtualApplianceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkVirtualApplianceName")])
		if err != nil {
			return nil, err
		}
		connectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameParam, networkVirtualApplianceNameParam, connectionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		v.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		v.beginDelete.remove(req)
	}

	return resp, nil
}

func (v *VirtualApplianceConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkVirtualAppliances/(?P<networkVirtualApplianceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkVirtualApplianceConnections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkVirtualApplianceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkVirtualApplianceName")])
	if err != nil {
		return nil, err
	}
	connectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameParam, networkVirtualApplianceNameParam, connectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualApplianceConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualApplianceConnectionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := v.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkVirtualAppliances/(?P<networkVirtualApplianceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkVirtualApplianceConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkVirtualApplianceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkVirtualApplianceName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListPager(resourceGroupNameParam, networkVirtualApplianceNameParam, nil)
		newListPager = &resp
		v.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.VirtualApplianceConnectionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		v.newListPager.remove(req)
	}
	return resp, nil
}

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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"net/http"
	"net/url"
	"regexp"
)

// VPNServerConfigurationsAssociatedWithVirtualWanServer is a fake server for instances of the armnetwork.VPNServerConfigurationsAssociatedWithVirtualWanClient type.
type VPNServerConfigurationsAssociatedWithVirtualWanServer struct {
	// BeginList is the fake for method VPNServerConfigurationsAssociatedWithVirtualWanClient.BeginList
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginList func(ctx context.Context, resourceGroupName string, virtualWANName string, options *armnetwork.VPNServerConfigurationsAssociatedWithVirtualWanClientBeginListOptions) (resp azfake.PollerResponder[armnetwork.VPNServerConfigurationsAssociatedWithVirtualWanClientListResponse], errResp azfake.ErrorResponder)
}

// NewVPNServerConfigurationsAssociatedWithVirtualWanServerTransport creates a new instance of VPNServerConfigurationsAssociatedWithVirtualWanServerTransport with the provided implementation.
// The returned VPNServerConfigurationsAssociatedWithVirtualWanServerTransport instance is connected to an instance of armnetwork.VPNServerConfigurationsAssociatedWithVirtualWanClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVPNServerConfigurationsAssociatedWithVirtualWanServerTransport(srv *VPNServerConfigurationsAssociatedWithVirtualWanServer) *VPNServerConfigurationsAssociatedWithVirtualWanServerTransport {
	return &VPNServerConfigurationsAssociatedWithVirtualWanServerTransport{
		srv:       srv,
		beginList: newTracker[azfake.PollerResponder[armnetwork.VPNServerConfigurationsAssociatedWithVirtualWanClientListResponse]](),
	}
}

// VPNServerConfigurationsAssociatedWithVirtualWanServerTransport connects instances of armnetwork.VPNServerConfigurationsAssociatedWithVirtualWanClient to instances of VPNServerConfigurationsAssociatedWithVirtualWanServer.
// Don't use this type directly, use NewVPNServerConfigurationsAssociatedWithVirtualWanServerTransport instead.
type VPNServerConfigurationsAssociatedWithVirtualWanServerTransport struct {
	srv       *VPNServerConfigurationsAssociatedWithVirtualWanServer
	beginList *tracker[azfake.PollerResponder[armnetwork.VPNServerConfigurationsAssociatedWithVirtualWanClientListResponse]]
}

// Do implements the policy.Transporter interface for VPNServerConfigurationsAssociatedWithVirtualWanServerTransport.
func (v *VPNServerConfigurationsAssociatedWithVirtualWanServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VPNServerConfigurationsAssociatedWithVirtualWanClient.BeginList":
		resp, err = v.dispatchBeginList(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VPNServerConfigurationsAssociatedWithVirtualWanServerTransport) dispatchBeginList(req *http.Request) (*http.Response, error) {
	if v.srv.BeginList == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginList not implemented")}
	}
	beginList := v.beginList.get(req)
	if beginList == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualWans/(?P<virtualWANName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/vpnServerConfigurations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualWANNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualWANName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginList(req.Context(), resourceGroupNameUnescaped, virtualWANNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginList = &respr
		v.beginList.add(req, beginList)
	}

	resp, err := server.PollerResponderNext(beginList, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginList.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginList) {
		v.beginList.remove(req)
	}

	return resp, nil
}

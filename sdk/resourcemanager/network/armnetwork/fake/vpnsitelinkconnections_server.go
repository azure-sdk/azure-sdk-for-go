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

// VPNSiteLinkConnectionsServer is a fake server for instances of the armnetwork.VPNSiteLinkConnectionsClient type.
type VPNSiteLinkConnectionsServer struct {
	// Get is the fake for method VPNSiteLinkConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string, options *armnetwork.VPNSiteLinkConnectionsClientGetOptions) (resp azfake.Responder[armnetwork.VPNSiteLinkConnectionsClientGetResponse], errResp azfake.ErrorResponder)
}

// NewVPNSiteLinkConnectionsServerTransport creates a new instance of VPNSiteLinkConnectionsServerTransport with the provided implementation.
// The returned VPNSiteLinkConnectionsServerTransport instance is connected to an instance of armnetwork.VPNSiteLinkConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVPNSiteLinkConnectionsServerTransport(srv *VPNSiteLinkConnectionsServer) *VPNSiteLinkConnectionsServerTransport {
	return &VPNSiteLinkConnectionsServerTransport{srv: srv}
}

// VPNSiteLinkConnectionsServerTransport connects instances of armnetwork.VPNSiteLinkConnectionsClient to instances of VPNSiteLinkConnectionsServer.
// Don't use this type directly, use NewVPNSiteLinkConnectionsServerTransport instead.
type VPNSiteLinkConnectionsServerTransport struct {
	srv *VPNSiteLinkConnectionsServer
}

// Do implements the policy.Transporter interface for VPNSiteLinkConnectionsServerTransport.
func (v *VPNSiteLinkConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VPNSiteLinkConnectionsClient.Get":
		resp, err = v.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VPNSiteLinkConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/vpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/vpnConnections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/vpnLinkConnections/(?P<linkConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	gatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
	if err != nil {
		return nil, err
	}
	connectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
	if err != nil {
		return nil, err
	}
	linkConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("linkConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameUnescaped, gatewayNameUnescaped, connectionNameUnescaped, linkConnectionNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VPNSiteLinkConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

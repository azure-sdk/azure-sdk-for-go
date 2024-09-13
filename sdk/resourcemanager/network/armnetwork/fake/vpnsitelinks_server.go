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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
	"net/http"
	"net/url"
	"regexp"
)

// VPNSiteLinksServer is a fake server for instances of the armnetwork.VPNSiteLinksClient type.
type VPNSiteLinksServer struct {
	// Get is the fake for method VPNSiteLinksClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteLinkName string, options *armnetwork.VPNSiteLinksClientGetOptions) (resp azfake.Responder[armnetwork.VPNSiteLinksClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByVPNSitePager is the fake for method VPNSiteLinksClient.NewListByVPNSitePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByVPNSitePager func(resourceGroupName string, vpnSiteName string, options *armnetwork.VPNSiteLinksClientListByVPNSiteOptions) (resp azfake.PagerResponder[armnetwork.VPNSiteLinksClientListByVPNSiteResponse])
}

// NewVPNSiteLinksServerTransport creates a new instance of VPNSiteLinksServerTransport with the provided implementation.
// The returned VPNSiteLinksServerTransport instance is connected to an instance of armnetwork.VPNSiteLinksClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVPNSiteLinksServerTransport(srv *VPNSiteLinksServer) *VPNSiteLinksServerTransport {
	return &VPNSiteLinksServerTransport{
		srv:                   srv,
		newListByVPNSitePager: newTracker[azfake.PagerResponder[armnetwork.VPNSiteLinksClientListByVPNSiteResponse]](),
	}
}

// VPNSiteLinksServerTransport connects instances of armnetwork.VPNSiteLinksClient to instances of VPNSiteLinksServer.
// Don't use this type directly, use NewVPNSiteLinksServerTransport instead.
type VPNSiteLinksServerTransport struct {
	srv                   *VPNSiteLinksServer
	newListByVPNSitePager *tracker[azfake.PagerResponder[armnetwork.VPNSiteLinksClientListByVPNSiteResponse]]
}

// Do implements the policy.Transporter interface for VPNSiteLinksServerTransport.
func (v *VPNSiteLinksServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VPNSiteLinksClient.Get":
		resp, err = v.dispatchGet(req)
	case "VPNSiteLinksClient.NewListByVPNSitePager":
		resp, err = v.dispatchNewListByVPNSitePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VPNSiteLinksServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/vpnSites/(?P<vpnSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/vpnSiteLinks/(?P<vpnSiteLinkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vpnSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vpnSiteName")])
	if err != nil {
		return nil, err
	}
	vpnSiteLinkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vpnSiteLinkName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameParam, vpnSiteNameParam, vpnSiteLinkNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VPNSiteLink, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VPNSiteLinksServerTransport) dispatchNewListByVPNSitePager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListByVPNSitePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByVPNSitePager not implemented")}
	}
	newListByVPNSitePager := v.newListByVPNSitePager.get(req)
	if newListByVPNSitePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/vpnSites/(?P<vpnSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/vpnSiteLinks`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vpnSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vpnSiteName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListByVPNSitePager(resourceGroupNameParam, vpnSiteNameParam, nil)
		newListByVPNSitePager = &resp
		v.newListByVPNSitePager.add(req, newListByVPNSitePager)
		server.PagerResponderInjectNextLinks(newListByVPNSitePager, req, func(page *armnetwork.VPNSiteLinksClientListByVPNSiteResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByVPNSitePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListByVPNSitePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByVPNSitePager) {
		v.newListByVPNSitePager.remove(req)
	}
	return resp, nil
}

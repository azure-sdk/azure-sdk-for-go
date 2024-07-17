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

// VirtualApplianceSKUsServer is a fake server for instances of the armnetwork.VirtualApplianceSKUsClient type.
type VirtualApplianceSKUsServer struct {
	// Get is the fake for method VirtualApplianceSKUsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, skuName string, options *armnetwork.VirtualApplianceSKUsClientGetOptions) (resp azfake.Responder[armnetwork.VirtualApplianceSKUsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VirtualApplianceSKUsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armnetwork.VirtualApplianceSKUsClientListOptions) (resp azfake.PagerResponder[armnetwork.VirtualApplianceSKUsClientListResponse])
}

// NewVirtualApplianceSKUsServerTransport creates a new instance of VirtualApplianceSKUsServerTransport with the provided implementation.
// The returned VirtualApplianceSKUsServerTransport instance is connected to an instance of armnetwork.VirtualApplianceSKUsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualApplianceSKUsServerTransport(srv *VirtualApplianceSKUsServer) *VirtualApplianceSKUsServerTransport {
	return &VirtualApplianceSKUsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armnetwork.VirtualApplianceSKUsClientListResponse]](),
	}
}

// VirtualApplianceSKUsServerTransport connects instances of armnetwork.VirtualApplianceSKUsClient to instances of VirtualApplianceSKUsServer.
// Don't use this type directly, use NewVirtualApplianceSKUsServerTransport instead.
type VirtualApplianceSKUsServerTransport struct {
	srv          *VirtualApplianceSKUsServer
	newListPager *tracker[azfake.PagerResponder[armnetwork.VirtualApplianceSKUsClientListResponse]]
}

// Do implements the policy.Transporter interface for VirtualApplianceSKUsServerTransport.
func (v *VirtualApplianceSKUsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualApplianceSKUsClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualApplianceSKUsClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualApplianceSKUsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkVirtualApplianceSkus/(?P<skuName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	skuNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("skuName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), skuNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualApplianceSKU, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualApplianceSKUsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := v.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkVirtualApplianceSkus`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := v.srv.NewListPager(nil)
		newListPager = &resp
		v.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.VirtualApplianceSKUsClientListResponse, createLink func() string) {
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

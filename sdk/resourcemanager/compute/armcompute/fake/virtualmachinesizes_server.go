//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
	"net/http"
	"net/url"
	"regexp"
)

// VirtualMachineSizesServer is a fake server for instances of the armcompute.VirtualMachineSizesClient type.
type VirtualMachineSizesServer struct {
	// NewListPager is the fake for method VirtualMachineSizesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(location string, options *armcompute.VirtualMachineSizesClientListOptions) (resp azfake.PagerResponder[armcompute.VirtualMachineSizesClientListResponse])
}

// NewVirtualMachineSizesServerTransport creates a new instance of VirtualMachineSizesServerTransport with the provided implementation.
// The returned VirtualMachineSizesServerTransport instance is connected to an instance of armcompute.VirtualMachineSizesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualMachineSizesServerTransport(srv *VirtualMachineSizesServer) *VirtualMachineSizesServerTransport {
	return &VirtualMachineSizesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armcompute.VirtualMachineSizesClientListResponse]](),
	}
}

// VirtualMachineSizesServerTransport connects instances of armcompute.VirtualMachineSizesClient to instances of VirtualMachineSizesServer.
// Don't use this type directly, use NewVirtualMachineSizesServerTransport instead.
type VirtualMachineSizesServerTransport struct {
	srv          *VirtualMachineSizesServer
	newListPager *tracker[azfake.PagerResponder[armcompute.VirtualMachineSizesClientListResponse]]
}

// Do implements the policy.Transporter interface for VirtualMachineSizesServerTransport.
func (v *VirtualMachineSizesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualMachineSizesClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualMachineSizesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := v.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/vmSizes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListPager(locationParam, nil)
		newListPager = &resp
		v.newListPager.add(req, newListPager)
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

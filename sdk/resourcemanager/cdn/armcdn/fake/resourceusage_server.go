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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v3"
	"net/http"
	"regexp"
)

// ResourceUsageServer is a fake server for instances of the armcdn.ResourceUsageClient type.
type ResourceUsageServer struct {
	// NewListPager is the fake for method ResourceUsageClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armcdn.ResourceUsageClientListOptions) (resp azfake.PagerResponder[armcdn.ResourceUsageClientListResponse])
}

// NewResourceUsageServerTransport creates a new instance of ResourceUsageServerTransport with the provided implementation.
// The returned ResourceUsageServerTransport instance is connected to an instance of armcdn.ResourceUsageClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewResourceUsageServerTransport(srv *ResourceUsageServer) *ResourceUsageServerTransport {
	return &ResourceUsageServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armcdn.ResourceUsageClientListResponse]](),
	}
}

// ResourceUsageServerTransport connects instances of armcdn.ResourceUsageClient to instances of ResourceUsageServer.
// Don't use this type directly, use NewResourceUsageServerTransport instead.
type ResourceUsageServerTransport struct {
	srv          *ResourceUsageServer
	newListPager *tracker[azfake.PagerResponder[armcdn.ResourceUsageClientListResponse]]
}

// Do implements the policy.Transporter interface for ResourceUsageServerTransport.
func (r *ResourceUsageServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ResourceUsageClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ResourceUsageServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/checkResourceUsage`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := r.srv.NewListPager(nil)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcdn.ResourceUsageClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}

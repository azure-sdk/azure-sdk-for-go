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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
	"net/http"
	"net/url"
	"regexp"
)

// WorkspaceLinksServer is a fake server for instances of the armapimanagement.WorkspaceLinksClient type.
type WorkspaceLinksServer struct {
	// NewListByServicePager is the fake for method WorkspaceLinksClient.NewListByServicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServicePager func(resourceGroupName string, serviceName string, options *armapimanagement.WorkspaceLinksClientListByServiceOptions) (resp azfake.PagerResponder[armapimanagement.WorkspaceLinksClientListByServiceResponse])
}

// NewWorkspaceLinksServerTransport creates a new instance of WorkspaceLinksServerTransport with the provided implementation.
// The returned WorkspaceLinksServerTransport instance is connected to an instance of armapimanagement.WorkspaceLinksClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkspaceLinksServerTransport(srv *WorkspaceLinksServer) *WorkspaceLinksServerTransport {
	return &WorkspaceLinksServerTransport{
		srv:                   srv,
		newListByServicePager: newTracker[azfake.PagerResponder[armapimanagement.WorkspaceLinksClientListByServiceResponse]](),
	}
}

// WorkspaceLinksServerTransport connects instances of armapimanagement.WorkspaceLinksClient to instances of WorkspaceLinksServer.
// Don't use this type directly, use NewWorkspaceLinksServerTransport instead.
type WorkspaceLinksServerTransport struct {
	srv                   *WorkspaceLinksServer
	newListByServicePager *tracker[azfake.PagerResponder[armapimanagement.WorkspaceLinksClientListByServiceResponse]]
}

// Do implements the policy.Transporter interface for WorkspaceLinksServerTransport.
func (w *WorkspaceLinksServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WorkspaceLinksClient.NewListByServicePager":
		resp, err = w.dispatchNewListByServicePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WorkspaceLinksServerTransport) dispatchNewListByServicePager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListByServicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServicePager not implemented")}
	}
	newListByServicePager := w.newListByServicePager.get(req)
	if newListByServicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaceLinks`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		resp := w.srv.NewListByServicePager(resourceGroupNameParam, serviceNameParam, nil)
		newListByServicePager = &resp
		w.newListByServicePager.add(req, newListByServicePager)
		server.PagerResponderInjectNextLinks(newListByServicePager, req, func(page *armapimanagement.WorkspaceLinksClientListByServiceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListByServicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServicePager) {
		w.newListByServicePager.remove(req)
	}
	return resp, nil
}
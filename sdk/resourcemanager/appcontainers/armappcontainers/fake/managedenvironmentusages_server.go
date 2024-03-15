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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ManagedEnvironmentUsagesServer is a fake server for instances of the armappcontainers.ManagedEnvironmentUsagesClient type.
type ManagedEnvironmentUsagesServer struct {
	// NewListPager is the fake for method ManagedEnvironmentUsagesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, environmentName string, options *armappcontainers.ManagedEnvironmentUsagesClientListOptions) (resp azfake.PagerResponder[armappcontainers.ManagedEnvironmentUsagesClientListResponse])
}

// NewManagedEnvironmentUsagesServerTransport creates a new instance of ManagedEnvironmentUsagesServerTransport with the provided implementation.
// The returned ManagedEnvironmentUsagesServerTransport instance is connected to an instance of armappcontainers.ManagedEnvironmentUsagesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedEnvironmentUsagesServerTransport(srv *ManagedEnvironmentUsagesServer) *ManagedEnvironmentUsagesServerTransport {
	return &ManagedEnvironmentUsagesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armappcontainers.ManagedEnvironmentUsagesClientListResponse]](),
	}
}

// ManagedEnvironmentUsagesServerTransport connects instances of armappcontainers.ManagedEnvironmentUsagesClient to instances of ManagedEnvironmentUsagesServer.
// Don't use this type directly, use NewManagedEnvironmentUsagesServerTransport instead.
type ManagedEnvironmentUsagesServerTransport struct {
	srv          *ManagedEnvironmentUsagesServer
	newListPager *tracker[azfake.PagerResponder[armappcontainers.ManagedEnvironmentUsagesClientListResponse]]
}

// Do implements the policy.Transporter interface for ManagedEnvironmentUsagesServerTransport.
func (m *ManagedEnvironmentUsagesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagedEnvironmentUsagesClient.NewListPager":
		resp, err = m.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagedEnvironmentUsagesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := m.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/usages`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
		if err != nil {
			return nil, err
		}
		resp := m.srv.NewListPager(resourceGroupNameParam, environmentNameParam, nil)
		newListPager = &resp
		m.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armappcontainers.ManagedEnvironmentUsagesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		m.newListPager.remove(req)
	}
	return resp, nil
}

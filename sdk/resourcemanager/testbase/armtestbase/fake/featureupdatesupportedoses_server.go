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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
	"net/http"
	"net/url"
	"regexp"
)

// FeatureUpdateSupportedOsesServer is a fake server for instances of the armtestbase.FeatureUpdateSupportedOsesClient type.
type FeatureUpdateSupportedOsesServer struct {
	// NewListPager is the fake for method FeatureUpdateSupportedOsesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, testBaseAccountName string, options *armtestbase.FeatureUpdateSupportedOsesClientListOptions) (resp azfake.PagerResponder[armtestbase.FeatureUpdateSupportedOsesClientListResponse])
}

// NewFeatureUpdateSupportedOsesServerTransport creates a new instance of FeatureUpdateSupportedOsesServerTransport with the provided implementation.
// The returned FeatureUpdateSupportedOsesServerTransport instance is connected to an instance of armtestbase.FeatureUpdateSupportedOsesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFeatureUpdateSupportedOsesServerTransport(srv *FeatureUpdateSupportedOsesServer) *FeatureUpdateSupportedOsesServerTransport {
	return &FeatureUpdateSupportedOsesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armtestbase.FeatureUpdateSupportedOsesClientListResponse]](),
	}
}

// FeatureUpdateSupportedOsesServerTransport connects instances of armtestbase.FeatureUpdateSupportedOsesClient to instances of FeatureUpdateSupportedOsesServer.
// Don't use this type directly, use NewFeatureUpdateSupportedOsesServerTransport instead.
type FeatureUpdateSupportedOsesServerTransport struct {
	srv          *FeatureUpdateSupportedOsesServer
	newListPager *tracker[azfake.PagerResponder[armtestbase.FeatureUpdateSupportedOsesClientListResponse]]
}

// Do implements the policy.Transporter interface for FeatureUpdateSupportedOsesServerTransport.
func (f *FeatureUpdateSupportedOsesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FeatureUpdateSupportedOsesClient.NewListPager":
		resp, err = f.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FeatureUpdateSupportedOsesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := f.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.TestBase/testBaseAccounts/(?P<testBaseAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/featureUpdateSupportedOses`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		testBaseAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("testBaseAccountName")])
		if err != nil {
			return nil, err
		}
		resp := f.srv.NewListPager(resourceGroupNameParam, testBaseAccountNameParam, nil)
		newListPager = &resp
		f.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armtestbase.FeatureUpdateSupportedOsesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		f.newListPager.remove(req)
	}
	return resp, nil
}

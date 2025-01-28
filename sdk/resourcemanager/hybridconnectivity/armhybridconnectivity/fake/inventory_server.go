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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
	"net/http"
	"net/url"
	"regexp"
)

// InventoryServer is a fake server for instances of the armhybridconnectivity.InventoryClient type.
type InventoryServer struct {
	// Get is the fake for method InventoryClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceURI string, solutionConfiguration string, inventoryID string, options *armhybridconnectivity.InventoryClientGetOptions) (resp azfake.Responder[armhybridconnectivity.InventoryClientGetResponse], errResp azfake.ErrorResponder)

	// NewListBySolutionConfigurationPager is the fake for method InventoryClient.NewListBySolutionConfigurationPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySolutionConfigurationPager func(resourceURI string, solutionConfiguration string, options *armhybridconnectivity.InventoryClientListBySolutionConfigurationOptions) (resp azfake.PagerResponder[armhybridconnectivity.InventoryClientListBySolutionConfigurationResponse])
}

// NewInventoryServerTransport creates a new instance of InventoryServerTransport with the provided implementation.
// The returned InventoryServerTransport instance is connected to an instance of armhybridconnectivity.InventoryClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewInventoryServerTransport(srv *InventoryServer) *InventoryServerTransport {
	return &InventoryServerTransport{
		srv:                                 srv,
		newListBySolutionConfigurationPager: newTracker[azfake.PagerResponder[armhybridconnectivity.InventoryClientListBySolutionConfigurationResponse]](),
	}
}

// InventoryServerTransport connects instances of armhybridconnectivity.InventoryClient to instances of InventoryServer.
// Don't use this type directly, use NewInventoryServerTransport instead.
type InventoryServerTransport struct {
	srv                                 *InventoryServer
	newListBySolutionConfigurationPager *tracker[azfake.PagerResponder[armhybridconnectivity.InventoryClientListBySolutionConfigurationResponse]]
}

// Do implements the policy.Transporter interface for InventoryServerTransport.
func (i *InventoryServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "InventoryClient.Get":
		resp, err = i.dispatchGet(req)
	case "InventoryClient.NewListBySolutionConfigurationPager":
		resp, err = i.dispatchNewListBySolutionConfigurationPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *InventoryServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridConnectivity/solutionConfigurations/(?P<solutionConfiguration>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/inventory/(?P<inventoryId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
	if err != nil {
		return nil, err
	}
	solutionConfigurationParam, err := url.PathUnescape(matches[regex.SubexpIndex("solutionConfiguration")])
	if err != nil {
		return nil, err
	}
	inventoryIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("inventoryId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceURIParam, solutionConfigurationParam, inventoryIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).InventoryResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *InventoryServerTransport) dispatchNewListBySolutionConfigurationPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListBySolutionConfigurationPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySolutionConfigurationPager not implemented")}
	}
	newListBySolutionConfigurationPager := i.newListBySolutionConfigurationPager.get(req)
	if newListBySolutionConfigurationPager == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridConnectivity/solutionConfigurations/(?P<solutionConfiguration>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/inventory`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		solutionConfigurationParam, err := url.PathUnescape(matches[regex.SubexpIndex("solutionConfiguration")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListBySolutionConfigurationPager(resourceURIParam, solutionConfigurationParam, nil)
		newListBySolutionConfigurationPager = &resp
		i.newListBySolutionConfigurationPager.add(req, newListBySolutionConfigurationPager)
		server.PagerResponderInjectNextLinks(newListBySolutionConfigurationPager, req, func(page *armhybridconnectivity.InventoryClientListBySolutionConfigurationResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySolutionConfigurationPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListBySolutionConfigurationPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySolutionConfigurationPager) {
		i.newListBySolutionConfigurationPager.remove(req)
	}
	return resp, nil
}

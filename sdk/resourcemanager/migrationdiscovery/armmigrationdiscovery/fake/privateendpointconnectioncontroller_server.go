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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationdiscovery/armmigrationdiscovery"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// PrivateEndpointConnectionControllerServer is a fake server for instances of the armmigrationdiscovery.PrivateEndpointConnectionControllerClient type.
type PrivateEndpointConnectionControllerServer struct {
	// Create is the fake for method PrivateEndpointConnectionControllerClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, siteName string, peConnectionName string, body armmigrationdiscovery.PrivateEndpointConnection, options *armmigrationdiscovery.PrivateEndpointConnectionControllerClientCreateOptions) (resp azfake.Responder[armmigrationdiscovery.PrivateEndpointConnectionControllerClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method PrivateEndpointConnectionControllerClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, siteName string, peConnectionName string, options *armmigrationdiscovery.PrivateEndpointConnectionControllerClientDeleteOptions) (resp azfake.Responder[armmigrationdiscovery.PrivateEndpointConnectionControllerClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PrivateEndpointConnectionControllerClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, siteName string, peConnectionName string, options *armmigrationdiscovery.PrivateEndpointConnectionControllerClientGetOptions) (resp azfake.Responder[armmigrationdiscovery.PrivateEndpointConnectionControllerClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByMasterSitePager is the fake for method PrivateEndpointConnectionControllerClient.NewListByMasterSitePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByMasterSitePager func(resourceGroupName string, siteName string, options *armmigrationdiscovery.PrivateEndpointConnectionControllerClientListByMasterSiteOptions) (resp azfake.PagerResponder[armmigrationdiscovery.PrivateEndpointConnectionControllerClientListByMasterSiteResponse])
}

// NewPrivateEndpointConnectionControllerServerTransport creates a new instance of PrivateEndpointConnectionControllerServerTransport with the provided implementation.
// The returned PrivateEndpointConnectionControllerServerTransport instance is connected to an instance of armmigrationdiscovery.PrivateEndpointConnectionControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateEndpointConnectionControllerServerTransport(srv *PrivateEndpointConnectionControllerServer) *PrivateEndpointConnectionControllerServerTransport {
	return &PrivateEndpointConnectionControllerServerTransport{
		srv:                      srv,
		newListByMasterSitePager: newTracker[azfake.PagerResponder[armmigrationdiscovery.PrivateEndpointConnectionControllerClientListByMasterSiteResponse]](),
	}
}

// PrivateEndpointConnectionControllerServerTransport connects instances of armmigrationdiscovery.PrivateEndpointConnectionControllerClient to instances of PrivateEndpointConnectionControllerServer.
// Don't use this type directly, use NewPrivateEndpointConnectionControllerServerTransport instead.
type PrivateEndpointConnectionControllerServerTransport struct {
	srv                      *PrivateEndpointConnectionControllerServer
	newListByMasterSitePager *tracker[azfake.PagerResponder[armmigrationdiscovery.PrivateEndpointConnectionControllerClientListByMasterSiteResponse]]
}

// Do implements the policy.Transporter interface for PrivateEndpointConnectionControllerServerTransport.
func (p *PrivateEndpointConnectionControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PrivateEndpointConnectionControllerClient.Create":
		resp, err = p.dispatchCreate(req)
	case "PrivateEndpointConnectionControllerClient.Delete":
		resp, err = p.dispatchDelete(req)
	case "PrivateEndpointConnectionControllerClient.Get":
		resp, err = p.dispatchGet(req)
	case "PrivateEndpointConnectionControllerClient.NewListByMasterSitePager":
		resp, err = p.dispatchNewListByMasterSitePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PrivateEndpointConnectionControllerServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if p.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<peConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmigrationdiscovery.PrivateEndpointConnection](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
	if err != nil {
		return nil, err
	}
	peConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("peConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Create(req.Context(), resourceGroupNameParam, siteNameParam, peConnectionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).RetryAfter; val != nil {
		resp.Header.Set("Retry-After", strconv.FormatInt(int64(*val), 10))
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionControllerServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if p.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<peConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
	if err != nil {
		return nil, err
	}
	peConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("peConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Delete(req.Context(), resourceGroupNameParam, siteNameParam, peConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionControllerServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<peConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
	if err != nil {
		return nil, err
	}
	peConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("peConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, siteNameParam, peConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionControllerServerTransport) dispatchNewListByMasterSitePager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByMasterSitePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByMasterSitePager not implemented")}
	}
	newListByMasterSitePager := p.newListByMasterSitePager.get(req)
	if newListByMasterSitePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListByMasterSitePager(resourceGroupNameParam, siteNameParam, nil)
		newListByMasterSitePager = &resp
		p.newListByMasterSitePager.add(req, newListByMasterSitePager)
		server.PagerResponderInjectNextLinks(newListByMasterSitePager, req, func(page *armmigrationdiscovery.PrivateEndpointConnectionControllerClientListByMasterSiteResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByMasterSitePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByMasterSitePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByMasterSitePager) {
		p.newListByMasterSitePager.remove(req)
	}
	return resp, nil
}

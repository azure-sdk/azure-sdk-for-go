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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate/v2"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateLinkResourcesControllerServer is a fake server for instances of the armmigrate.PrivateLinkResourcesControllerClient type.
type PrivateLinkResourcesControllerServer struct {
	// Get is the fake for method PrivateLinkResourcesControllerClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, siteName string, privateLinkResourceName string, options *armmigrate.PrivateLinkResourcesControllerClientGetOptions) (resp azfake.Responder[armmigrate.PrivateLinkResourcesControllerClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByMasterSitePager is the fake for method PrivateLinkResourcesControllerClient.NewListByMasterSitePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByMasterSitePager func(resourceGroupName string, siteName string, options *armmigrate.PrivateLinkResourcesControllerClientListByMasterSiteOptions) (resp azfake.PagerResponder[armmigrate.PrivateLinkResourcesControllerClientListByMasterSiteResponse])
}

// NewPrivateLinkResourcesControllerServerTransport creates a new instance of PrivateLinkResourcesControllerServerTransport with the provided implementation.
// The returned PrivateLinkResourcesControllerServerTransport instance is connected to an instance of armmigrate.PrivateLinkResourcesControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateLinkResourcesControllerServerTransport(srv *PrivateLinkResourcesControllerServer) *PrivateLinkResourcesControllerServerTransport {
	return &PrivateLinkResourcesControllerServerTransport{
		srv:                      srv,
		newListByMasterSitePager: newTracker[azfake.PagerResponder[armmigrate.PrivateLinkResourcesControllerClientListByMasterSiteResponse]](),
	}
}

// PrivateLinkResourcesControllerServerTransport connects instances of armmigrate.PrivateLinkResourcesControllerClient to instances of PrivateLinkResourcesControllerServer.
// Don't use this type directly, use NewPrivateLinkResourcesControllerServerTransport instead.
type PrivateLinkResourcesControllerServerTransport struct {
	srv                      *PrivateLinkResourcesControllerServer
	newListByMasterSitePager *tracker[azfake.PagerResponder[armmigrate.PrivateLinkResourcesControllerClientListByMasterSiteResponse]]
}

// Do implements the policy.Transporter interface for PrivateLinkResourcesControllerServerTransport.
func (p *PrivateLinkResourcesControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PrivateLinkResourcesControllerClient.Get":
		resp, err = p.dispatchGet(req)
	case "PrivateLinkResourcesControllerClient.NewListByMasterSitePager":
		resp, err = p.dispatchNewListByMasterSitePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PrivateLinkResourcesControllerServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources/(?P<privateLinkResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	privateLinkResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateLinkResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, siteNameParam, privateLinkResourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateLinkResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateLinkResourcesControllerServerTransport) dispatchNewListByMasterSitePager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByMasterSitePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByMasterSitePager not implemented")}
	}
	newListByMasterSitePager := p.newListByMasterSitePager.get(req)
	if newListByMasterSitePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources`
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
		server.PagerResponderInjectNextLinks(newListByMasterSitePager, req, func(page *armmigrate.PrivateLinkResourcesControllerClientListByMasterSiteResponse, createLink func() string) {
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
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
)

// WebAppRunAsAccountsControllerServer is a fake server for instances of the armmigrationdiscovery.WebAppRunAsAccountsControllerClient type.
type WebAppRunAsAccountsControllerServer struct {
	// Get is the fake for method WebAppRunAsAccountsControllerClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, accountName string, options *armmigrationdiscovery.WebAppRunAsAccountsControllerClientGetOptions) (resp azfake.Responder[armmigrationdiscovery.WebAppRunAsAccountsControllerClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByWebAppSitePager is the fake for method WebAppRunAsAccountsControllerClient.NewListByWebAppSitePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByWebAppSitePager func(resourceGroupName string, siteName string, webAppSiteName string, options *armmigrationdiscovery.WebAppRunAsAccountsControllerClientListByWebAppSiteOptions) (resp azfake.PagerResponder[armmigrationdiscovery.WebAppRunAsAccountsControllerClientListByWebAppSiteResponse])
}

// NewWebAppRunAsAccountsControllerServerTransport creates a new instance of WebAppRunAsAccountsControllerServerTransport with the provided implementation.
// The returned WebAppRunAsAccountsControllerServerTransport instance is connected to an instance of armmigrationdiscovery.WebAppRunAsAccountsControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWebAppRunAsAccountsControllerServerTransport(srv *WebAppRunAsAccountsControllerServer) *WebAppRunAsAccountsControllerServerTransport {
	return &WebAppRunAsAccountsControllerServerTransport{
		srv:                      srv,
		newListByWebAppSitePager: newTracker[azfake.PagerResponder[armmigrationdiscovery.WebAppRunAsAccountsControllerClientListByWebAppSiteResponse]](),
	}
}

// WebAppRunAsAccountsControllerServerTransport connects instances of armmigrationdiscovery.WebAppRunAsAccountsControllerClient to instances of WebAppRunAsAccountsControllerServer.
// Don't use this type directly, use NewWebAppRunAsAccountsControllerServerTransport instead.
type WebAppRunAsAccountsControllerServerTransport struct {
	srv                      *WebAppRunAsAccountsControllerServer
	newListByWebAppSitePager *tracker[azfake.PagerResponder[armmigrationdiscovery.WebAppRunAsAccountsControllerClientListByWebAppSiteResponse]]
}

// Do implements the policy.Transporter interface for WebAppRunAsAccountsControllerServerTransport.
func (w *WebAppRunAsAccountsControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WebAppRunAsAccountsControllerClient.Get":
		resp, err = w.dispatchGet(req)
	case "WebAppRunAsAccountsControllerClient.NewListByWebAppSitePager":
		resp, err = w.dispatchNewListByWebAppSitePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WebAppRunAsAccountsControllerServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webAppSites/(?P<webAppSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runasaccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	webAppSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("webAppSiteName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, siteNameParam, webAppSiteNameParam, accountNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WebAppRunAsAccount, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WebAppRunAsAccountsControllerServerTransport) dispatchNewListByWebAppSitePager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListByWebAppSitePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByWebAppSitePager not implemented")}
	}
	newListByWebAppSitePager := w.newListByWebAppSitePager.get(req)
	if newListByWebAppSitePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webAppSites/(?P<webAppSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runasaccounts`
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
		webAppSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("webAppSiteName")])
		if err != nil {
			return nil, err
		}
		resp := w.srv.NewListByWebAppSitePager(resourceGroupNameParam, siteNameParam, webAppSiteNameParam, nil)
		newListByWebAppSitePager = &resp
		w.newListByWebAppSitePager.add(req, newListByWebAppSitePager)
		server.PagerResponderInjectNextLinks(newListByWebAppSitePager, req, func(page *armmigrationdiscovery.WebAppRunAsAccountsControllerClientListByWebAppSiteResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByWebAppSitePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListByWebAppSitePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByWebAppSitePager) {
		w.newListByWebAppSitePager.remove(req)
	}
	return resp, nil
}

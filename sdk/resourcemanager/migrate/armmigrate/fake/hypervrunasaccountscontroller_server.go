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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
	"net/http"
	"net/url"
	"regexp"
)

// HypervRunAsAccountsControllerServer is a fake server for instances of the armmigrate.HypervRunAsAccountsControllerClient type.
type HypervRunAsAccountsControllerServer struct {
	// Get is the fake for method HypervRunAsAccountsControllerClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, siteName string, accountName string, options *armmigrate.HypervRunAsAccountsControllerClientGetOptions) (resp azfake.Responder[armmigrate.HypervRunAsAccountsControllerClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByHypervSitePager is the fake for method HypervRunAsAccountsControllerClient.NewListByHypervSitePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByHypervSitePager func(resourceGroupName string, siteName string, options *armmigrate.HypervRunAsAccountsControllerClientListByHypervSiteOptions) (resp azfake.PagerResponder[armmigrate.HypervRunAsAccountsControllerClientListByHypervSiteResponse])
}

// NewHypervRunAsAccountsControllerServerTransport creates a new instance of HypervRunAsAccountsControllerServerTransport with the provided implementation.
// The returned HypervRunAsAccountsControllerServerTransport instance is connected to an instance of armmigrate.HypervRunAsAccountsControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHypervRunAsAccountsControllerServerTransport(srv *HypervRunAsAccountsControllerServer) *HypervRunAsAccountsControllerServerTransport {
	return &HypervRunAsAccountsControllerServerTransport{
		srv:                      srv,
		newListByHypervSitePager: newTracker[azfake.PagerResponder[armmigrate.HypervRunAsAccountsControllerClientListByHypervSiteResponse]](),
	}
}

// HypervRunAsAccountsControllerServerTransport connects instances of armmigrate.HypervRunAsAccountsControllerClient to instances of HypervRunAsAccountsControllerServer.
// Don't use this type directly, use NewHypervRunAsAccountsControllerServerTransport instead.
type HypervRunAsAccountsControllerServerTransport struct {
	srv                      *HypervRunAsAccountsControllerServer
	newListByHypervSitePager *tracker[azfake.PagerResponder[armmigrate.HypervRunAsAccountsControllerClientListByHypervSiteResponse]]
}

// Do implements the policy.Transporter interface for HypervRunAsAccountsControllerServerTransport.
func (h *HypervRunAsAccountsControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return h.dispatchToMethodFake(req, method)
}

func (h *HypervRunAsAccountsControllerServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if hypervRunAsAccountsControllerServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = hypervRunAsAccountsControllerServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "HypervRunAsAccountsControllerClient.Get":
				res.resp, res.err = h.dispatchGet(req)
			case "HypervRunAsAccountsControllerClient.NewListByHypervSitePager":
				res.resp, res.err = h.dispatchNewListByHypervSitePager(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (h *HypervRunAsAccountsControllerServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if h.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runAsAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Get(req.Context(), resourceGroupNameParam, siteNameParam, accountNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HypervRunAsAccountResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HypervRunAsAccountsControllerServerTransport) dispatchNewListByHypervSitePager(req *http.Request) (*http.Response, error) {
	if h.srv.NewListByHypervSitePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByHypervSitePager not implemented")}
	}
	newListByHypervSitePager := h.newListByHypervSitePager.get(req)
	if newListByHypervSitePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runAsAccounts`
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
		resp := h.srv.NewListByHypervSitePager(resourceGroupNameParam, siteNameParam, nil)
		newListByHypervSitePager = &resp
		h.newListByHypervSitePager.add(req, newListByHypervSitePager)
		server.PagerResponderInjectNextLinks(newListByHypervSitePager, req, func(page *armmigrate.HypervRunAsAccountsControllerClientListByHypervSiteResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByHypervSitePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		h.newListByHypervSitePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByHypervSitePager) {
		h.newListByHypervSitePager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to HypervRunAsAccountsControllerServerTransport
var hypervRunAsAccountsControllerServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

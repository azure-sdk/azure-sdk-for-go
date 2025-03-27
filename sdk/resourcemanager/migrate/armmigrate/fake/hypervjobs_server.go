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

// HypervJobsServer is a fake server for instances of the armmigrate.HypervJobsClient type.
type HypervJobsServer struct {
	// Get is the fake for method HypervJobsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, siteName string, jobName string, options *armmigrate.HypervJobsClientGetOptions) (resp azfake.Responder[armmigrate.HypervJobsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByHypervSitePager is the fake for method HypervJobsClient.NewListByHypervSitePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByHypervSitePager func(resourceGroupName string, siteName string, options *armmigrate.HypervJobsClientListByHypervSiteOptions) (resp azfake.PagerResponder[armmigrate.HypervJobsClientListByHypervSiteResponse])
}

// NewHypervJobsServerTransport creates a new instance of HypervJobsServerTransport with the provided implementation.
// The returned HypervJobsServerTransport instance is connected to an instance of armmigrate.HypervJobsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHypervJobsServerTransport(srv *HypervJobsServer) *HypervJobsServerTransport {
	return &HypervJobsServerTransport{
		srv:                      srv,
		newListByHypervSitePager: newTracker[azfake.PagerResponder[armmigrate.HypervJobsClientListByHypervSiteResponse]](),
	}
}

// HypervJobsServerTransport connects instances of armmigrate.HypervJobsClient to instances of HypervJobsServer.
// Don't use this type directly, use NewHypervJobsServerTransport instead.
type HypervJobsServerTransport struct {
	srv                      *HypervJobsServer
	newListByHypervSitePager *tracker[azfake.PagerResponder[armmigrate.HypervJobsClientListByHypervSiteResponse]]
}

// Do implements the policy.Transporter interface for HypervJobsServerTransport.
func (h *HypervJobsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return h.dispatchToMethodFake(req, method)
}

func (h *HypervJobsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if hypervJobsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = hypervJobsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "HypervJobsClient.Get":
				res.resp, res.err = h.dispatchGet(req)
			case "HypervJobsClient.NewListByHypervSitePager":
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

func (h *HypervJobsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if h.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Get(req.Context(), resourceGroupNameParam, siteNameParam, jobNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HypervJob, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HypervJobsServerTransport) dispatchNewListByHypervSitePager(req *http.Request) (*http.Response, error) {
	if h.srv.NewListByHypervSitePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByHypervSitePager not implemented")}
	}
	newListByHypervSitePager := h.newListByHypervSitePager.get(req)
	if newListByHypervSitePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs`
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
		server.PagerResponderInjectNextLinks(newListByHypervSitePager, req, func(page *armmigrate.HypervJobsClientListByHypervSiteResponse, createLink func() string) {
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

// set this to conditionally intercept incoming requests to HypervJobsServerTransport
var hypervJobsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

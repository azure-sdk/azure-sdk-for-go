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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
	"net/http"
	"net/url"
	"regexp"
)

// HypervJobsControllerServer is a fake server for instances of the armmigrate.HypervJobsControllerClient type.
type HypervJobsControllerServer struct {
	// Get is the fake for method HypervJobsControllerClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, siteName string, jobName string, options *armmigrate.HypervJobsControllerClientGetOptions) (resp azfake.Responder[armmigrate.HypervJobsControllerClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByVmwareSitePager is the fake for method HypervJobsControllerClient.NewListByVmwareSitePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByVmwareSitePager func(resourceGroupName string, siteName string, options *armmigrate.HypervJobsControllerClientListByVmwareSiteOptions) (resp azfake.PagerResponder[armmigrate.HypervJobsControllerClientListByVmwareSiteResponse])
}

// NewHypervJobsControllerServerTransport creates a new instance of HypervJobsControllerServerTransport with the provided implementation.
// The returned HypervJobsControllerServerTransport instance is connected to an instance of armmigrate.HypervJobsControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHypervJobsControllerServerTransport(srv *HypervJobsControllerServer) *HypervJobsControllerServerTransport {
	return &HypervJobsControllerServerTransport{
		srv:                      srv,
		newListByVmwareSitePager: newTracker[azfake.PagerResponder[armmigrate.HypervJobsControllerClientListByVmwareSiteResponse]](),
	}
}

// HypervJobsControllerServerTransport connects instances of armmigrate.HypervJobsControllerClient to instances of HypervJobsControllerServer.
// Don't use this type directly, use NewHypervJobsControllerServerTransport instead.
type HypervJobsControllerServerTransport struct {
	srv                      *HypervJobsControllerServer
	newListByVmwareSitePager *tracker[azfake.PagerResponder[armmigrate.HypervJobsControllerClientListByVmwareSiteResponse]]
}

// Do implements the policy.Transporter interface for HypervJobsControllerServerTransport.
func (h *HypervJobsControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "HypervJobsControllerClient.Get":
		resp, err = h.dispatchGet(req)
	case "HypervJobsControllerClient.NewListByVmwareSitePager":
		resp, err = h.dispatchNewListByVmwareSitePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HypervJobsControllerServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if h.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VmwareJob, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HypervJobsControllerServerTransport) dispatchNewListByVmwareSitePager(req *http.Request) (*http.Response, error) {
	if h.srv.NewListByVmwareSitePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByVmwareSitePager not implemented")}
	}
	newListByVmwareSitePager := h.newListByVmwareSitePager.get(req)
	if newListByVmwareSitePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs`
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
		resp := h.srv.NewListByVmwareSitePager(resourceGroupNameParam, siteNameParam, nil)
		newListByVmwareSitePager = &resp
		h.newListByVmwareSitePager.add(req, newListByVmwareSitePager)
		server.PagerResponderInjectNextLinks(newListByVmwareSitePager, req, func(page *armmigrate.HypervJobsControllerClientListByVmwareSiteResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByVmwareSitePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		h.newListByVmwareSitePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByVmwareSitePager) {
		h.newListByVmwareSitePager.remove(req)
	}
	return resp, nil
}

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
	"net/http"
	"net/url"
	"regexp"
)

// HcxEnterpriseSitesServer is a fake server for instances of the armavs.HcxEnterpriseSitesClient type.
type HcxEnterpriseSitesServer struct {
	// CreateOrUpdate is the fake for method HcxEnterpriseSitesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, privateCloudName string, hcxEnterpriseSiteName string, hcxEnterpriseSite armavs.HcxEnterpriseSite, options *armavs.HcxEnterpriseSitesClientCreateOrUpdateOptions) (resp azfake.Responder[armavs.HcxEnterpriseSitesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method HcxEnterpriseSitesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, privateCloudName string, hcxEnterpriseSiteName string, options *armavs.HcxEnterpriseSitesClientDeleteOptions) (resp azfake.Responder[armavs.HcxEnterpriseSitesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method HcxEnterpriseSitesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, privateCloudName string, hcxEnterpriseSiteName string, options *armavs.HcxEnterpriseSitesClientGetOptions) (resp azfake.Responder[armavs.HcxEnterpriseSitesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method HcxEnterpriseSitesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, privateCloudName string, options *armavs.HcxEnterpriseSitesClientListOptions) (resp azfake.PagerResponder[armavs.HcxEnterpriseSitesClientListResponse])
}

// NewHcxEnterpriseSitesServerTransport creates a new instance of HcxEnterpriseSitesServerTransport with the provided implementation.
// The returned HcxEnterpriseSitesServerTransport instance is connected to an instance of armavs.HcxEnterpriseSitesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHcxEnterpriseSitesServerTransport(srv *HcxEnterpriseSitesServer) *HcxEnterpriseSitesServerTransport {
	return &HcxEnterpriseSitesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armavs.HcxEnterpriseSitesClientListResponse]](),
	}
}

// HcxEnterpriseSitesServerTransport connects instances of armavs.HcxEnterpriseSitesClient to instances of HcxEnterpriseSitesServer.
// Don't use this type directly, use NewHcxEnterpriseSitesServerTransport instead.
type HcxEnterpriseSitesServerTransport struct {
	srv          *HcxEnterpriseSitesServer
	newListPager *tracker[azfake.PagerResponder[armavs.HcxEnterpriseSitesClientListResponse]]
}

// Do implements the policy.Transporter interface for HcxEnterpriseSitesServerTransport.
func (h *HcxEnterpriseSitesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return h.dispatchToMethodFake(req, method)
}

func (h *HcxEnterpriseSitesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "HcxEnterpriseSitesClient.CreateOrUpdate":
		resp, err = h.dispatchCreateOrUpdate(req)
	case "HcxEnterpriseSitesClient.Delete":
		resp, err = h.dispatchDelete(req)
	case "HcxEnterpriseSitesClient.Get":
		resp, err = h.dispatchGet(req)
	case "HcxEnterpriseSitesClient.NewListPager":
		resp, err = h.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (h *HcxEnterpriseSitesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if h.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hcxEnterpriseSites/(?P<hcxEnterpriseSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armavs.HcxEnterpriseSite](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
	if err != nil {
		return nil, err
	}
	hcxEnterpriseSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hcxEnterpriseSiteName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, privateCloudNameParam, hcxEnterpriseSiteNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HcxEnterpriseSite, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HcxEnterpriseSitesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if h.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hcxEnterpriseSites/(?P<hcxEnterpriseSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
	if err != nil {
		return nil, err
	}
	hcxEnterpriseSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hcxEnterpriseSiteName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Delete(req.Context(), resourceGroupNameParam, privateCloudNameParam, hcxEnterpriseSiteNameParam, nil)
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

func (h *HcxEnterpriseSitesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if h.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hcxEnterpriseSites/(?P<hcxEnterpriseSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
	if err != nil {
		return nil, err
	}
	hcxEnterpriseSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hcxEnterpriseSiteName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Get(req.Context(), resourceGroupNameParam, privateCloudNameParam, hcxEnterpriseSiteNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HcxEnterpriseSite, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HcxEnterpriseSitesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if h.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := h.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hcxEnterpriseSites`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
		if err != nil {
			return nil, err
		}
		resp := h.srv.NewListPager(resourceGroupNameParam, privateCloudNameParam, nil)
		newListPager = &resp
		h.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armavs.HcxEnterpriseSitesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		h.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		h.newListPager.remove(req)
	}
	return resp, nil
}

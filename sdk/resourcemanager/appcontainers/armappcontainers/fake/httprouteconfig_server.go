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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v4"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// HTTPRouteConfigServer is a fake server for instances of the armappcontainers.HTTPRouteConfigClient type.
type HTTPRouteConfigServer struct {
	// CreateOrUpdate is the fake for method HTTPRouteConfigClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, environmentName string, httpRouteName string, options *armappcontainers.HTTPRouteConfigClientCreateOrUpdateOptions) (resp azfake.Responder[armappcontainers.HTTPRouteConfigClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method HTTPRouteConfigClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, environmentName string, httpRouteName string, options *armappcontainers.HTTPRouteConfigClientDeleteOptions) (resp azfake.Responder[armappcontainers.HTTPRouteConfigClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method HTTPRouteConfigClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, environmentName string, httpRouteName string, options *armappcontainers.HTTPRouteConfigClientGetOptions) (resp azfake.Responder[armappcontainers.HTTPRouteConfigClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method HTTPRouteConfigClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, environmentName string, options *armappcontainers.HTTPRouteConfigClientListOptions) (resp azfake.PagerResponder[armappcontainers.HTTPRouteConfigClientListResponse])

	// Update is the fake for method HTTPRouteConfigClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, environmentName string, httpRouteName string, httpRouteConfigEnvelope armappcontainers.HTTPRouteConfig, options *armappcontainers.HTTPRouteConfigClientUpdateOptions) (resp azfake.Responder[armappcontainers.HTTPRouteConfigClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewHTTPRouteConfigServerTransport creates a new instance of HTTPRouteConfigServerTransport with the provided implementation.
// The returned HTTPRouteConfigServerTransport instance is connected to an instance of armappcontainers.HTTPRouteConfigClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHTTPRouteConfigServerTransport(srv *HTTPRouteConfigServer) *HTTPRouteConfigServerTransport {
	return &HTTPRouteConfigServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armappcontainers.HTTPRouteConfigClientListResponse]](),
	}
}

// HTTPRouteConfigServerTransport connects instances of armappcontainers.HTTPRouteConfigClient to instances of HTTPRouteConfigServer.
// Don't use this type directly, use NewHTTPRouteConfigServerTransport instead.
type HTTPRouteConfigServerTransport struct {
	srv          *HTTPRouteConfigServer
	newListPager *tracker[azfake.PagerResponder[armappcontainers.HTTPRouteConfigClientListResponse]]
}

// Do implements the policy.Transporter interface for HTTPRouteConfigServerTransport.
func (h *HTTPRouteConfigServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "HTTPRouteConfigClient.CreateOrUpdate":
		resp, err = h.dispatchCreateOrUpdate(req)
	case "HTTPRouteConfigClient.Delete":
		resp, err = h.dispatchDelete(req)
	case "HTTPRouteConfigClient.Get":
		resp, err = h.dispatchGet(req)
	case "HTTPRouteConfigClient.NewListPager":
		resp, err = h.dispatchNewListPager(req)
	case "HTTPRouteConfigClient.Update":
		resp, err = h.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HTTPRouteConfigServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if h.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/httpRouteConfigs/(?P<httpRouteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armappcontainers.HTTPRouteConfig](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
	if err != nil {
		return nil, err
	}
	httpRouteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("httpRouteName")])
	if err != nil {
		return nil, err
	}
	var options *armappcontainers.HTTPRouteConfigClientCreateOrUpdateOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armappcontainers.HTTPRouteConfigClientCreateOrUpdateOptions{
			HTTPRouteConfigEnvelope: &body,
		}
	}
	respr, errRespr := h.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, environmentNameParam, httpRouteNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HTTPRouteConfig, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPRouteConfigServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if h.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/httpRouteConfigs/(?P<httpRouteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
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
	httpRouteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("httpRouteName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Delete(req.Context(), resourceGroupNameParam, environmentNameParam, httpRouteNameParam, nil)
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

func (h *HTTPRouteConfigServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if h.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/httpRouteConfigs/(?P<httpRouteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
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
	httpRouteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("httpRouteName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Get(req.Context(), resourceGroupNameParam, environmentNameParam, httpRouteNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HTTPRouteConfig, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HTTPRouteConfigServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if h.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := h.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/httpRouteConfigs`
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
		resp := h.srv.NewListPager(resourceGroupNameParam, environmentNameParam, nil)
		newListPager = &resp
		h.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armappcontainers.HTTPRouteConfigClientListResponse, createLink func() string) {
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

func (h *HTTPRouteConfigServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if h.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/httpRouteConfigs/(?P<httpRouteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armappcontainers.HTTPRouteConfig](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
	if err != nil {
		return nil, err
	}
	httpRouteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("httpRouteName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Update(req.Context(), resourceGroupNameParam, environmentNameParam, httpRouteNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HTTPRouteConfig, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

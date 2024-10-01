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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azureadexternalidentities/armazureadexternalidentities"
	"net/http"
	"net/url"
	"regexp"
)

// GuestUsagesServer is a fake server for instances of the armazureadexternalidentities.GuestUsagesClient type.
type GuestUsagesServer struct {
	// Create is the fake for method GuestUsagesClient.Create
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, resourceGroupName string, resourceName string, resource armazureadexternalidentities.GuestUsagesResource, options *armazureadexternalidentities.GuestUsagesClientCreateOptions) (resp azfake.Responder[armazureadexternalidentities.GuestUsagesClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method GuestUsagesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, resourceName string, options *armazureadexternalidentities.GuestUsagesClientDeleteOptions) (resp azfake.Responder[armazureadexternalidentities.GuestUsagesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method GuestUsagesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, options *armazureadexternalidentities.GuestUsagesClientGetOptions) (resp azfake.Responder[armazureadexternalidentities.GuestUsagesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method GuestUsagesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armazureadexternalidentities.GuestUsagesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armazureadexternalidentities.GuestUsagesClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method GuestUsagesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armazureadexternalidentities.GuestUsagesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armazureadexternalidentities.GuestUsagesClientListBySubscriptionResponse])

	// Update is the fake for method GuestUsagesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, resourceName string, resourcePatch armazureadexternalidentities.GuestUsagesResourcePatch, options *armazureadexternalidentities.GuestUsagesClientUpdateOptions) (resp azfake.Responder[armazureadexternalidentities.GuestUsagesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewGuestUsagesServerTransport creates a new instance of GuestUsagesServerTransport with the provided implementation.
// The returned GuestUsagesServerTransport instance is connected to an instance of armazureadexternalidentities.GuestUsagesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGuestUsagesServerTransport(srv *GuestUsagesServer) *GuestUsagesServerTransport {
	return &GuestUsagesServerTransport{
		srv:                         srv,
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armazureadexternalidentities.GuestUsagesClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armazureadexternalidentities.GuestUsagesClientListBySubscriptionResponse]](),
	}
}

// GuestUsagesServerTransport connects instances of armazureadexternalidentities.GuestUsagesClient to instances of GuestUsagesServer.
// Don't use this type directly, use NewGuestUsagesServerTransport instead.
type GuestUsagesServerTransport struct {
	srv                         *GuestUsagesServer
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armazureadexternalidentities.GuestUsagesClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armazureadexternalidentities.GuestUsagesClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for GuestUsagesServerTransport.
func (g *GuestUsagesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GuestUsagesClient.Create":
		resp, err = g.dispatchCreate(req)
	case "GuestUsagesClient.Delete":
		resp, err = g.dispatchDelete(req)
	case "GuestUsagesClient.Get":
		resp, err = g.dispatchGet(req)
	case "GuestUsagesClient.NewListByResourceGroupPager":
		resp, err = g.dispatchNewListByResourceGroupPager(req)
	case "GuestUsagesClient.NewListBySubscriptionPager":
		resp, err = g.dispatchNewListBySubscriptionPager(req)
	case "GuestUsagesClient.Update":
		resp, err = g.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GuestUsagesServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if g.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureActiveDirectory/guestUsages/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armazureadexternalidentities.GuestUsagesResource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Create(req.Context(), resourceGroupNameParam, resourceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GuestUsagesResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GuestUsagesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if g.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureActiveDirectory/guestUsages/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Delete(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
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

func (g *GuestUsagesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if g.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureActiveDirectory/guestUsages/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GuestUsagesResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GuestUsagesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := g.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureActiveDirectory/guestUsages`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := g.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		g.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armazureadexternalidentities.GuestUsagesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		g.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (g *GuestUsagesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := g.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureActiveDirectory/guestUsages`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := g.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		g.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		g.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (g *GuestUsagesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureActiveDirectory/guestUsages/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armazureadexternalidentities.GuestUsagesResourcePatch](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Update(req.Context(), resourceGroupNameParam, resourceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GuestUsagesResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// AzureMonitorWorkspacesServer is a fake server for instances of the armmonitor.AzureMonitorWorkspacesClient type.
type AzureMonitorWorkspacesServer struct {
	// Create is the fake for method AzureMonitorWorkspacesClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, azureMonitorWorkspaceName string, azureMonitorWorkspaceProperties armmonitor.AzureMonitorWorkspaceResource, options *armmonitor.AzureMonitorWorkspacesClientCreateOptions) (resp azfake.Responder[armmonitor.AzureMonitorWorkspacesClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method AzureMonitorWorkspacesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, azureMonitorWorkspaceName string, options *armmonitor.AzureMonitorWorkspacesClientBeginDeleteOptions) (resp azfake.PollerResponder[armmonitor.AzureMonitorWorkspacesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AzureMonitorWorkspacesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, azureMonitorWorkspaceName string, options *armmonitor.AzureMonitorWorkspacesClientGetOptions) (resp azfake.Responder[armmonitor.AzureMonitorWorkspacesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method AzureMonitorWorkspacesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armmonitor.AzureMonitorWorkspacesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armmonitor.AzureMonitorWorkspacesClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method AzureMonitorWorkspacesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armmonitor.AzureMonitorWorkspacesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armmonitor.AzureMonitorWorkspacesClientListBySubscriptionResponse])

	// Update is the fake for method AzureMonitorWorkspacesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, azureMonitorWorkspaceName string, options *armmonitor.AzureMonitorWorkspacesClientUpdateOptions) (resp azfake.Responder[armmonitor.AzureMonitorWorkspacesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewAzureMonitorWorkspacesServerTransport creates a new instance of AzureMonitorWorkspacesServerTransport with the provided implementation.
// The returned AzureMonitorWorkspacesServerTransport instance is connected to an instance of armmonitor.AzureMonitorWorkspacesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAzureMonitorWorkspacesServerTransport(srv *AzureMonitorWorkspacesServer) *AzureMonitorWorkspacesServerTransport {
	return &AzureMonitorWorkspacesServerTransport{
		srv:                         srv,
		beginDelete:                 newTracker[azfake.PollerResponder[armmonitor.AzureMonitorWorkspacesClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armmonitor.AzureMonitorWorkspacesClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armmonitor.AzureMonitorWorkspacesClientListBySubscriptionResponse]](),
	}
}

// AzureMonitorWorkspacesServerTransport connects instances of armmonitor.AzureMonitorWorkspacesClient to instances of AzureMonitorWorkspacesServer.
// Don't use this type directly, use NewAzureMonitorWorkspacesServerTransport instead.
type AzureMonitorWorkspacesServerTransport struct {
	srv                         *AzureMonitorWorkspacesServer
	beginDelete                 *tracker[azfake.PollerResponder[armmonitor.AzureMonitorWorkspacesClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armmonitor.AzureMonitorWorkspacesClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armmonitor.AzureMonitorWorkspacesClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for AzureMonitorWorkspacesServerTransport.
func (a *AzureMonitorWorkspacesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AzureMonitorWorkspacesClient.Create":
		resp, err = a.dispatchCreate(req)
	case "AzureMonitorWorkspacesClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "AzureMonitorWorkspacesClient.Get":
		resp, err = a.dispatchGet(req)
	case "AzureMonitorWorkspacesClient.NewListByResourceGroupPager":
		resp, err = a.dispatchNewListByResourceGroupPager(req)
	case "AzureMonitorWorkspacesClient.NewListBySubscriptionPager":
		resp, err = a.dispatchNewListBySubscriptionPager(req)
	case "AzureMonitorWorkspacesClient.Update":
		resp, err = a.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AzureMonitorWorkspacesServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if a.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Monitor/accounts/(?P<azureMonitorWorkspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmonitor.AzureMonitorWorkspaceResource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	azureMonitorWorkspaceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("azureMonitorWorkspaceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Create(req.Context(), resourceGroupNameUnescaped, azureMonitorWorkspaceNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureMonitorWorkspaceResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AzureMonitorWorkspacesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Monitor/accounts/(?P<azureMonitorWorkspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		azureMonitorWorkspaceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("azureMonitorWorkspaceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, azureMonitorWorkspaceNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		a.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		a.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		a.beginDelete.remove(req)
	}

	return resp, nil
}

func (a *AzureMonitorWorkspacesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Monitor/accounts/(?P<azureMonitorWorkspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	azureMonitorWorkspaceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("azureMonitorWorkspaceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameUnescaped, azureMonitorWorkspaceNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureMonitorWorkspaceResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AzureMonitorWorkspacesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := a.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Monitor/accounts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByResourceGroupPager(resourceGroupNameUnescaped, nil)
		newListByResourceGroupPager = &resp
		a.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armmonitor.AzureMonitorWorkspacesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		a.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (a *AzureMonitorWorkspacesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := a.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Monitor/accounts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		a.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armmonitor.AzureMonitorWorkspacesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		a.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (a *AzureMonitorWorkspacesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Monitor/accounts/(?P<azureMonitorWorkspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmonitor.AzureMonitorWorkspaceResourceForUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	azureMonitorWorkspaceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("azureMonitorWorkspaceName")])
	if err != nil {
		return nil, err
	}
	var options *armmonitor.AzureMonitorWorkspacesClientUpdateOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armmonitor.AzureMonitorWorkspacesClientUpdateOptions{
			AzureMonitorWorkspaceProperties: &body,
		}
	}
	respr, errRespr := a.srv.Update(req.Context(), resourceGroupNameUnescaped, azureMonitorWorkspaceNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureMonitorWorkspaceResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

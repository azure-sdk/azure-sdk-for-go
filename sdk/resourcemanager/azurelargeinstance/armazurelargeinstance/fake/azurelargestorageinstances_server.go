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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurelargeinstance/armazurelargeinstance"
	"net/http"
	"net/url"
	"regexp"
)

// AzureLargeStorageInstancesServer is a fake server for instances of the armazurelargeinstance.AzureLargeStorageInstancesClient type.
type AzureLargeStorageInstancesServer struct {
	// Get is the fake for method AzureLargeStorageInstancesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, azureLargeStorageInstanceName string, options *armazurelargeinstance.AzureLargeStorageInstancesClientGetOptions) (resp azfake.Responder[armazurelargeinstance.AzureLargeStorageInstancesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method AzureLargeStorageInstancesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armazurelargeinstance.AzureLargeStorageInstancesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armazurelargeinstance.AzureLargeStorageInstancesClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method AzureLargeStorageInstancesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armazurelargeinstance.AzureLargeStorageInstancesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armazurelargeinstance.AzureLargeStorageInstancesClientListBySubscriptionResponse])

	// Update is the fake for method AzureLargeStorageInstancesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, azureLargeStorageInstanceName string, properties armazurelargeinstance.AzureLargeStorageInstanceTagsUpdate, options *armazurelargeinstance.AzureLargeStorageInstancesClientUpdateOptions) (resp azfake.Responder[armazurelargeinstance.AzureLargeStorageInstancesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewAzureLargeStorageInstancesServerTransport creates a new instance of AzureLargeStorageInstancesServerTransport with the provided implementation.
// The returned AzureLargeStorageInstancesServerTransport instance is connected to an instance of armazurelargeinstance.AzureLargeStorageInstancesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAzureLargeStorageInstancesServerTransport(srv *AzureLargeStorageInstancesServer) *AzureLargeStorageInstancesServerTransport {
	return &AzureLargeStorageInstancesServerTransport{
		srv:                         srv,
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armazurelargeinstance.AzureLargeStorageInstancesClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armazurelargeinstance.AzureLargeStorageInstancesClientListBySubscriptionResponse]](),
	}
}

// AzureLargeStorageInstancesServerTransport connects instances of armazurelargeinstance.AzureLargeStorageInstancesClient to instances of AzureLargeStorageInstancesServer.
// Don't use this type directly, use NewAzureLargeStorageInstancesServerTransport instead.
type AzureLargeStorageInstancesServerTransport struct {
	srv                         *AzureLargeStorageInstancesServer
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armazurelargeinstance.AzureLargeStorageInstancesClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armazurelargeinstance.AzureLargeStorageInstancesClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for AzureLargeStorageInstancesServerTransport.
func (a *AzureLargeStorageInstancesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AzureLargeStorageInstancesClient.Get":
		resp, err = a.dispatchGet(req)
	case "AzureLargeStorageInstancesClient.NewListByResourceGroupPager":
		resp, err = a.dispatchNewListByResourceGroupPager(req)
	case "AzureLargeStorageInstancesClient.NewListBySubscriptionPager":
		resp, err = a.dispatchNewListBySubscriptionPager(req)
	case "AzureLargeStorageInstancesClient.Update":
		resp, err = a.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AzureLargeStorageInstancesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureLargeInstance/azureLargeStorageInstances/(?P<azureLargeStorageInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	azureLargeStorageInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureLargeStorageInstanceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, azureLargeStorageInstanceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureLargeStorageInstance, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AzureLargeStorageInstancesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := a.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureLargeInstance/azureLargeStorageInstances`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		a.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armazurelargeinstance.AzureLargeStorageInstancesClientListByResourceGroupResponse, createLink func() string) {
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

func (a *AzureLargeStorageInstancesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := a.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureLargeInstance/azureLargeStorageInstances`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		a.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armazurelargeinstance.AzureLargeStorageInstancesClientListBySubscriptionResponse, createLink func() string) {
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

func (a *AzureLargeStorageInstancesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureLargeInstance/azureLargeStorageInstances/(?P<azureLargeStorageInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armazurelargeinstance.AzureLargeStorageInstanceTagsUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	azureLargeStorageInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureLargeStorageInstanceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Update(req.Context(), resourceGroupNameParam, azureLargeStorageInstanceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureLargeStorageInstance, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

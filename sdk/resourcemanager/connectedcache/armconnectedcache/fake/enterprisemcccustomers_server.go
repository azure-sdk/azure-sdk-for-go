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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedcache/armconnectedcache"
	"net/http"
	"net/url"
	"regexp"
)

// EnterpriseMccCustomersServer is a fake server for instances of the armconnectedcache.EnterpriseMccCustomersClient type.
type EnterpriseMccCustomersServer struct {
	// BeginCreateOrUpdate is the fake for method EnterpriseMccCustomersClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, customerResourceName string, resource armconnectedcache.EnterpriseMccCustomerResource, options *armconnectedcache.EnterpriseMccCustomersClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armconnectedcache.EnterpriseMccCustomersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method EnterpriseMccCustomersClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, customerResourceName string, options *armconnectedcache.EnterpriseMccCustomersClientBeginDeleteOptions) (resp azfake.PollerResponder[armconnectedcache.EnterpriseMccCustomersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method EnterpriseMccCustomersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, customerResourceName string, options *armconnectedcache.EnterpriseMccCustomersClientGetOptions) (resp azfake.Responder[armconnectedcache.EnterpriseMccCustomersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method EnterpriseMccCustomersClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armconnectedcache.EnterpriseMccCustomersClientListByResourceGroupOptions) (resp azfake.PagerResponder[armconnectedcache.EnterpriseMccCustomersClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method EnterpriseMccCustomersClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armconnectedcache.EnterpriseMccCustomersClientListBySubscriptionOptions) (resp azfake.PagerResponder[armconnectedcache.EnterpriseMccCustomersClientListBySubscriptionResponse])

	// Update is the fake for method EnterpriseMccCustomersClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, customerResourceName string, properties armconnectedcache.PatchResource, options *armconnectedcache.EnterpriseMccCustomersClientUpdateOptions) (resp azfake.Responder[armconnectedcache.EnterpriseMccCustomersClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewEnterpriseMccCustomersServerTransport creates a new instance of EnterpriseMccCustomersServerTransport with the provided implementation.
// The returned EnterpriseMccCustomersServerTransport instance is connected to an instance of armconnectedcache.EnterpriseMccCustomersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewEnterpriseMccCustomersServerTransport(srv *EnterpriseMccCustomersServer) *EnterpriseMccCustomersServerTransport {
	return &EnterpriseMccCustomersServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armconnectedcache.EnterpriseMccCustomersClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armconnectedcache.EnterpriseMccCustomersClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armconnectedcache.EnterpriseMccCustomersClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armconnectedcache.EnterpriseMccCustomersClientListBySubscriptionResponse]](),
	}
}

// EnterpriseMccCustomersServerTransport connects instances of armconnectedcache.EnterpriseMccCustomersClient to instances of EnterpriseMccCustomersServer.
// Don't use this type directly, use NewEnterpriseMccCustomersServerTransport instead.
type EnterpriseMccCustomersServerTransport struct {
	srv                         *EnterpriseMccCustomersServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armconnectedcache.EnterpriseMccCustomersClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armconnectedcache.EnterpriseMccCustomersClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armconnectedcache.EnterpriseMccCustomersClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armconnectedcache.EnterpriseMccCustomersClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for EnterpriseMccCustomersServerTransport.
func (e *EnterpriseMccCustomersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "EnterpriseMccCustomersClient.BeginCreateOrUpdate":
		resp, err = e.dispatchBeginCreateOrUpdate(req)
	case "EnterpriseMccCustomersClient.BeginDelete":
		resp, err = e.dispatchBeginDelete(req)
	case "EnterpriseMccCustomersClient.Get":
		resp, err = e.dispatchGet(req)
	case "EnterpriseMccCustomersClient.NewListByResourceGroupPager":
		resp, err = e.dispatchNewListByResourceGroupPager(req)
	case "EnterpriseMccCustomersClient.NewListBySubscriptionPager":
		resp, err = e.dispatchNewListBySubscriptionPager(req)
	case "EnterpriseMccCustomersClient.Update":
		resp, err = e.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *EnterpriseMccCustomersServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := e.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedCache/enterpriseMccCustomers/(?P<customerResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armconnectedcache.EnterpriseMccCustomerResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		customerResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("customerResourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, customerResourceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		e.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		e.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		e.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (e *EnterpriseMccCustomersServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if e.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := e.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedCache/enterpriseMccCustomers/(?P<customerResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		customerResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("customerResourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginDelete(req.Context(), resourceGroupNameParam, customerResourceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		e.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		e.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		e.beginDelete.remove(req)
	}

	return resp, nil
}

func (e *EnterpriseMccCustomersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedCache/enterpriseMccCustomers/(?P<customerResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	customerResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("customerResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameParam, customerResourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EnterpriseMccCustomerResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EnterpriseMccCustomersServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := e.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedCache/enterpriseMccCustomers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		e.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armconnectedcache.EnterpriseMccCustomersClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		e.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (e *EnterpriseMccCustomersServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := e.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedCache/enterpriseMccCustomers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := e.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		e.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armconnectedcache.EnterpriseMccCustomersClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		e.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (e *EnterpriseMccCustomersServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedCache/enterpriseMccCustomers/(?P<customerResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armconnectedcache.PatchResource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	customerResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("customerResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Update(req.Context(), resourceGroupNameParam, customerResourceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EnterpriseMccCustomerResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

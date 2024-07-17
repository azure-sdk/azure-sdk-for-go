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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/fabric/armfabric"
	"net/http"
	"net/url"
	"regexp"
)

// CapacitiesServer is a fake server for instances of the armfabric.CapacitiesClient type.
type CapacitiesServer struct {
	// CheckNameAvailability is the fake for method CapacitiesClient.CheckNameAvailability
	// HTTP status codes to indicate success: http.StatusOK
	CheckNameAvailability func(ctx context.Context, location string, body armfabric.CheckNameAvailabilityRequest, options *armfabric.CapacitiesClientCheckNameAvailabilityOptions) (resp azfake.Responder[armfabric.CapacitiesClientCheckNameAvailabilityResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method CapacitiesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, capacityName string, resource armfabric.Capacity, options *armfabric.CapacitiesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armfabric.CapacitiesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method CapacitiesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, capacityName string, options *armfabric.CapacitiesClientBeginDeleteOptions) (resp azfake.PollerResponder[armfabric.CapacitiesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method CapacitiesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, capacityName string, options *armfabric.CapacitiesClientGetOptions) (resp azfake.Responder[armfabric.CapacitiesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method CapacitiesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armfabric.CapacitiesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armfabric.CapacitiesClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method CapacitiesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armfabric.CapacitiesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armfabric.CapacitiesClientListBySubscriptionResponse])

	// NewListSKUsPager is the fake for method CapacitiesClient.NewListSKUsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListSKUsPager func(options *armfabric.CapacitiesClientListSKUsOptions) (resp azfake.PagerResponder[armfabric.CapacitiesClientListSKUsResponse])

	// NewListSKUsForCapacityPager is the fake for method CapacitiesClient.NewListSKUsForCapacityPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListSKUsForCapacityPager func(resourceGroupName string, capacityName string, options *armfabric.CapacitiesClientListSKUsForCapacityOptions) (resp azfake.PagerResponder[armfabric.CapacitiesClientListSKUsForCapacityResponse])

	// BeginResume is the fake for method CapacitiesClient.BeginResume
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginResume func(ctx context.Context, resourceGroupName string, capacityName string, options *armfabric.CapacitiesClientBeginResumeOptions) (resp azfake.PollerResponder[armfabric.CapacitiesClientResumeResponse], errResp azfake.ErrorResponder)

	// BeginSuspend is the fake for method CapacitiesClient.BeginSuspend
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginSuspend func(ctx context.Context, resourceGroupName string, capacityName string, options *armfabric.CapacitiesClientBeginSuspendOptions) (resp azfake.PollerResponder[armfabric.CapacitiesClientSuspendResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method CapacitiesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, capacityName string, properties armfabric.CapacityUpdate, options *armfabric.CapacitiesClientBeginUpdateOptions) (resp azfake.PollerResponder[armfabric.CapacitiesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewCapacitiesServerTransport creates a new instance of CapacitiesServerTransport with the provided implementation.
// The returned CapacitiesServerTransport instance is connected to an instance of armfabric.CapacitiesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCapacitiesServerTransport(srv *CapacitiesServer) *CapacitiesServerTransport {
	return &CapacitiesServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armfabric.CapacitiesClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armfabric.CapacitiesClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armfabric.CapacitiesClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armfabric.CapacitiesClientListBySubscriptionResponse]](),
		newListSKUsPager:            newTracker[azfake.PagerResponder[armfabric.CapacitiesClientListSKUsResponse]](),
		newListSKUsForCapacityPager: newTracker[azfake.PagerResponder[armfabric.CapacitiesClientListSKUsForCapacityResponse]](),
		beginResume:                 newTracker[azfake.PollerResponder[armfabric.CapacitiesClientResumeResponse]](),
		beginSuspend:                newTracker[azfake.PollerResponder[armfabric.CapacitiesClientSuspendResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armfabric.CapacitiesClientUpdateResponse]](),
	}
}

// CapacitiesServerTransport connects instances of armfabric.CapacitiesClient to instances of CapacitiesServer.
// Don't use this type directly, use NewCapacitiesServerTransport instead.
type CapacitiesServerTransport struct {
	srv                         *CapacitiesServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armfabric.CapacitiesClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armfabric.CapacitiesClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armfabric.CapacitiesClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armfabric.CapacitiesClientListBySubscriptionResponse]]
	newListSKUsPager            *tracker[azfake.PagerResponder[armfabric.CapacitiesClientListSKUsResponse]]
	newListSKUsForCapacityPager *tracker[azfake.PagerResponder[armfabric.CapacitiesClientListSKUsForCapacityResponse]]
	beginResume                 *tracker[azfake.PollerResponder[armfabric.CapacitiesClientResumeResponse]]
	beginSuspend                *tracker[azfake.PollerResponder[armfabric.CapacitiesClientSuspendResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armfabric.CapacitiesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for CapacitiesServerTransport.
func (c *CapacitiesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CapacitiesClient.CheckNameAvailability":
		resp, err = c.dispatchCheckNameAvailability(req)
	case "CapacitiesClient.BeginCreateOrUpdate":
		resp, err = c.dispatchBeginCreateOrUpdate(req)
	case "CapacitiesClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "CapacitiesClient.Get":
		resp, err = c.dispatchGet(req)
	case "CapacitiesClient.NewListByResourceGroupPager":
		resp, err = c.dispatchNewListByResourceGroupPager(req)
	case "CapacitiesClient.NewListBySubscriptionPager":
		resp, err = c.dispatchNewListBySubscriptionPager(req)
	case "CapacitiesClient.NewListSKUsPager":
		resp, err = c.dispatchNewListSKUsPager(req)
	case "CapacitiesClient.NewListSKUsForCapacityPager":
		resp, err = c.dispatchNewListSKUsForCapacityPager(req)
	case "CapacitiesClient.BeginResume":
		resp, err = c.dispatchBeginResume(req)
	case "CapacitiesClient.BeginSuspend":
		resp, err = c.dispatchBeginSuspend(req)
	case "CapacitiesClient.BeginUpdate":
		resp, err = c.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CapacitiesServerTransport) dispatchCheckNameAvailability(req *http.Request) (*http.Response, error) {
	if c.srv.CheckNameAvailability == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckNameAvailability not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Fabric/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/checkNameAvailability`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armfabric.CheckNameAvailabilityRequest](req)
	if err != nil {
		return nil, err
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.CheckNameAvailability(req.Context(), locationParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CheckNameAvailabilityResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CapacitiesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Fabric/capacities/(?P<capacityName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armfabric.Capacity](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		capacityNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("capacityName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, capacityNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		c.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		c.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		c.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (c *CapacitiesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Fabric/capacities/(?P<capacityName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		capacityNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("capacityName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, capacityNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		c.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		c.beginDelete.remove(req)
	}

	return resp, nil
}

func (c *CapacitiesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Fabric/capacities/(?P<capacityName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	capacityNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("capacityName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, capacityNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Capacity, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CapacitiesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := c.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Fabric/capacities`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		c.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armfabric.CapacitiesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		c.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (c *CapacitiesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := c.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Fabric/capacities`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := c.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		c.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armfabric.CapacitiesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		c.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (c *CapacitiesServerTransport) dispatchNewListSKUsPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListSKUsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListSKUsPager not implemented")}
	}
	newListSKUsPager := c.newListSKUsPager.get(req)
	if newListSKUsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Fabric/skus`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := c.srv.NewListSKUsPager(nil)
		newListSKUsPager = &resp
		c.newListSKUsPager.add(req, newListSKUsPager)
		server.PagerResponderInjectNextLinks(newListSKUsPager, req, func(page *armfabric.CapacitiesClientListSKUsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListSKUsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListSKUsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListSKUsPager) {
		c.newListSKUsPager.remove(req)
	}
	return resp, nil
}

func (c *CapacitiesServerTransport) dispatchNewListSKUsForCapacityPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListSKUsForCapacityPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListSKUsForCapacityPager not implemented")}
	}
	newListSKUsForCapacityPager := c.newListSKUsForCapacityPager.get(req)
	if newListSKUsForCapacityPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Fabric/capacities/(?P<capacityName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/skus`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		capacityNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("capacityName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListSKUsForCapacityPager(resourceGroupNameParam, capacityNameParam, nil)
		newListSKUsForCapacityPager = &resp
		c.newListSKUsForCapacityPager.add(req, newListSKUsForCapacityPager)
		server.PagerResponderInjectNextLinks(newListSKUsForCapacityPager, req, func(page *armfabric.CapacitiesClientListSKUsForCapacityResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListSKUsForCapacityPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListSKUsForCapacityPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListSKUsForCapacityPager) {
		c.newListSKUsForCapacityPager.remove(req)
	}
	return resp, nil
}

func (c *CapacitiesServerTransport) dispatchBeginResume(req *http.Request) (*http.Response, error) {
	if c.srv.BeginResume == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginResume not implemented")}
	}
	beginResume := c.beginResume.get(req)
	if beginResume == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Fabric/capacities/(?P<capacityName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resume`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		capacityNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("capacityName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginResume(req.Context(), resourceGroupNameParam, capacityNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginResume = &respr
		c.beginResume.add(req, beginResume)
	}

	resp, err := server.PollerResponderNext(beginResume, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		c.beginResume.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginResume) {
		c.beginResume.remove(req)
	}

	return resp, nil
}

func (c *CapacitiesServerTransport) dispatchBeginSuspend(req *http.Request) (*http.Response, error) {
	if c.srv.BeginSuspend == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginSuspend not implemented")}
	}
	beginSuspend := c.beginSuspend.get(req)
	if beginSuspend == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Fabric/capacities/(?P<capacityName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/suspend`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		capacityNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("capacityName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginSuspend(req.Context(), resourceGroupNameParam, capacityNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginSuspend = &respr
		c.beginSuspend.add(req, beginSuspend)
	}

	resp, err := server.PollerResponderNext(beginSuspend, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		c.beginSuspend.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginSuspend) {
		c.beginSuspend.remove(req)
	}

	return resp, nil
}

func (c *CapacitiesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := c.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Fabric/capacities/(?P<capacityName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armfabric.CapacityUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		capacityNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("capacityName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginUpdate(req.Context(), resourceGroupNameParam, capacityNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		c.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		c.beginUpdate.remove(req)
	}

	return resp, nil
}

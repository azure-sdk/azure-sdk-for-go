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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ResourceHealthMetadataServer is a fake server for instances of the armappservice.ResourceHealthMetadataClient type.
type ResourceHealthMetadataServer struct {
	// GetBySite is the fake for method ResourceHealthMetadataClient.GetBySite
	// HTTP status codes to indicate success: http.StatusOK
	GetBySite func(ctx context.Context, resourceGroupName string, name string, options *armappservice.ResourceHealthMetadataClientGetBySiteOptions) (resp azfake.Responder[armappservice.ResourceHealthMetadataClientGetBySiteResponse], errResp azfake.ErrorResponder)

	// GetBySiteSlot is the fake for method ResourceHealthMetadataClient.GetBySiteSlot
	// HTTP status codes to indicate success: http.StatusOK
	GetBySiteSlot func(ctx context.Context, resourceGroupName string, name string, slot string, options *armappservice.ResourceHealthMetadataClientGetBySiteSlotOptions) (resp azfake.Responder[armappservice.ResourceHealthMetadataClientGetBySiteSlotResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ResourceHealthMetadataClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armappservice.ResourceHealthMetadataClientListOptions) (resp azfake.PagerResponder[armappservice.ResourceHealthMetadataClientListResponse])

	// NewListByResourceGroupPager is the fake for method ResourceHealthMetadataClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armappservice.ResourceHealthMetadataClientListByResourceGroupOptions) (resp azfake.PagerResponder[armappservice.ResourceHealthMetadataClientListByResourceGroupResponse])

	// NewListBySitePager is the fake for method ResourceHealthMetadataClient.NewListBySitePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySitePager func(resourceGroupName string, name string, options *armappservice.ResourceHealthMetadataClientListBySiteOptions) (resp azfake.PagerResponder[armappservice.ResourceHealthMetadataClientListBySiteResponse])

	// NewListBySiteSlotPager is the fake for method ResourceHealthMetadataClient.NewListBySiteSlotPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySiteSlotPager func(resourceGroupName string, name string, slot string, options *armappservice.ResourceHealthMetadataClientListBySiteSlotOptions) (resp azfake.PagerResponder[armappservice.ResourceHealthMetadataClientListBySiteSlotResponse])
}

// NewResourceHealthMetadataServerTransport creates a new instance of ResourceHealthMetadataServerTransport with the provided implementation.
// The returned ResourceHealthMetadataServerTransport instance is connected to an instance of armappservice.ResourceHealthMetadataClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewResourceHealthMetadataServerTransport(srv *ResourceHealthMetadataServer) *ResourceHealthMetadataServerTransport {
	return &ResourceHealthMetadataServerTransport{
		srv:                         srv,
		newListPager:                newTracker[azfake.PagerResponder[armappservice.ResourceHealthMetadataClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armappservice.ResourceHealthMetadataClientListByResourceGroupResponse]](),
		newListBySitePager:          newTracker[azfake.PagerResponder[armappservice.ResourceHealthMetadataClientListBySiteResponse]](),
		newListBySiteSlotPager:      newTracker[azfake.PagerResponder[armappservice.ResourceHealthMetadataClientListBySiteSlotResponse]](),
	}
}

// ResourceHealthMetadataServerTransport connects instances of armappservice.ResourceHealthMetadataClient to instances of ResourceHealthMetadataServer.
// Don't use this type directly, use NewResourceHealthMetadataServerTransport instead.
type ResourceHealthMetadataServerTransport struct {
	srv                         *ResourceHealthMetadataServer
	newListPager                *tracker[azfake.PagerResponder[armappservice.ResourceHealthMetadataClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armappservice.ResourceHealthMetadataClientListByResourceGroupResponse]]
	newListBySitePager          *tracker[azfake.PagerResponder[armappservice.ResourceHealthMetadataClientListBySiteResponse]]
	newListBySiteSlotPager      *tracker[azfake.PagerResponder[armappservice.ResourceHealthMetadataClientListBySiteSlotResponse]]
}

// Do implements the policy.Transporter interface for ResourceHealthMetadataServerTransport.
func (r *ResourceHealthMetadataServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ResourceHealthMetadataClient.GetBySite":
		resp, err = r.dispatchGetBySite(req)
	case "ResourceHealthMetadataClient.GetBySiteSlot":
		resp, err = r.dispatchGetBySiteSlot(req)
	case "ResourceHealthMetadataClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	case "ResourceHealthMetadataClient.NewListByResourceGroupPager":
		resp, err = r.dispatchNewListByResourceGroupPager(req)
	case "ResourceHealthMetadataClient.NewListBySitePager":
		resp, err = r.dispatchNewListBySitePager(req)
	case "ResourceHealthMetadataClient.NewListBySiteSlotPager":
		resp, err = r.dispatchNewListBySiteSlotPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ResourceHealthMetadataServerTransport) dispatchGetBySite(req *http.Request) (*http.Response, error) {
	if r.srv.GetBySite == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetBySite not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/sites/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceHealthMetadata/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.GetBySite(req.Context(), resourceGroupNameParam, nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ResourceHealthMetadata, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ResourceHealthMetadataServerTransport) dispatchGetBySiteSlot(req *http.Request) (*http.Response, error) {
	if r.srv.GetBySiteSlot == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetBySiteSlot not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/sites/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/slots/(?P<slot>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceHealthMetadata/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	slotParam, err := url.PathUnescape(matches[regex.SubexpIndex("slot")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.GetBySiteSlot(req.Context(), resourceGroupNameParam, nameParam, slotParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ResourceHealthMetadata, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ResourceHealthMetadataServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/resourceHealthMetadata`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := r.srv.NewListPager(nil)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armappservice.ResourceHealthMetadataClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}

func (r *ResourceHealthMetadataServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := r.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/resourceHealthMetadata`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		r.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armappservice.ResourceHealthMetadataClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		r.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (r *ResourceHealthMetadataServerTransport) dispatchNewListBySitePager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListBySitePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySitePager not implemented")}
	}
	newListBySitePager := r.newListBySitePager.get(req)
	if newListBySitePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/sites/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceHealthMetadata`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListBySitePager(resourceGroupNameParam, nameParam, nil)
		newListBySitePager = &resp
		r.newListBySitePager.add(req, newListBySitePager)
		server.PagerResponderInjectNextLinks(newListBySitePager, req, func(page *armappservice.ResourceHealthMetadataClientListBySiteResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySitePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListBySitePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySitePager) {
		r.newListBySitePager.remove(req)
	}
	return resp, nil
}

func (r *ResourceHealthMetadataServerTransport) dispatchNewListBySiteSlotPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListBySiteSlotPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySiteSlotPager not implemented")}
	}
	newListBySiteSlotPager := r.newListBySiteSlotPager.get(req)
	if newListBySiteSlotPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/sites/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/slots/(?P<slot>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceHealthMetadata`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		slotParam, err := url.PathUnescape(matches[regex.SubexpIndex("slot")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListBySiteSlotPager(resourceGroupNameParam, nameParam, slotParam, nil)
		newListBySiteSlotPager = &resp
		r.newListBySiteSlotPager.add(req, newListBySiteSlotPager)
		server.PagerResponderInjectNextLinks(newListBySiteSlotPager, req, func(page *armappservice.ResourceHealthMetadataClientListBySiteSlotResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySiteSlotPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListBySiteSlotPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySiteSlotPager) {
		r.newListBySiteSlotPager.remove(req)
	}
	return resp, nil
}

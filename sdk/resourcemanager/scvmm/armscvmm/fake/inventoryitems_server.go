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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm/v2"
	"net/http"
	"net/url"
	"regexp"
)

// InventoryItemsServer is a fake server for instances of the armscvmm.InventoryItemsClient type.
type InventoryItemsServer struct {
	// Create is the fake for method InventoryItemsClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, vmmServerName string, inventoryItemResourceName string, resource armscvmm.InventoryItem, options *armscvmm.InventoryItemsClientCreateOptions) (resp azfake.Responder[armscvmm.InventoryItemsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method InventoryItemsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, vmmServerName string, inventoryItemResourceName string, options *armscvmm.InventoryItemsClientDeleteOptions) (resp azfake.Responder[armscvmm.InventoryItemsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method InventoryItemsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vmmServerName string, inventoryItemResourceName string, options *armscvmm.InventoryItemsClientGetOptions) (resp azfake.Responder[armscvmm.InventoryItemsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByVmmServerPager is the fake for method InventoryItemsClient.NewListByVmmServerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByVmmServerPager func(resourceGroupName string, vmmServerName string, options *armscvmm.InventoryItemsClientListByVmmServerOptions) (resp azfake.PagerResponder[armscvmm.InventoryItemsClientListByVmmServerResponse])
}

// NewInventoryItemsServerTransport creates a new instance of InventoryItemsServerTransport with the provided implementation.
// The returned InventoryItemsServerTransport instance is connected to an instance of armscvmm.InventoryItemsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewInventoryItemsServerTransport(srv *InventoryItemsServer) *InventoryItemsServerTransport {
	return &InventoryItemsServerTransport{
		srv:                     srv,
		newListByVmmServerPager: newTracker[azfake.PagerResponder[armscvmm.InventoryItemsClientListByVmmServerResponse]](),
	}
}

// InventoryItemsServerTransport connects instances of armscvmm.InventoryItemsClient to instances of InventoryItemsServer.
// Don't use this type directly, use NewInventoryItemsServerTransport instead.
type InventoryItemsServerTransport struct {
	srv                     *InventoryItemsServer
	newListByVmmServerPager *tracker[azfake.PagerResponder[armscvmm.InventoryItemsClientListByVmmServerResponse]]
}

// Do implements the policy.Transporter interface for InventoryItemsServerTransport.
func (i *InventoryItemsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return i.dispatchToMethodFake(req, method)
}

func (i *InventoryItemsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "InventoryItemsClient.Create":
		resp, err = i.dispatchCreate(req)
	case "InventoryItemsClient.Delete":
		resp, err = i.dispatchDelete(req)
	case "InventoryItemsClient.Get":
		resp, err = i.dispatchGet(req)
	case "InventoryItemsClient.NewListByVmmServerPager":
		resp, err = i.dispatchNewListByVmmServerPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (i *InventoryItemsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if i.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ScVmm/vmmServers/(?P<vmmServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/inventoryItems/(?P<inventoryItemResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armscvmm.InventoryItem](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vmmServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmmServerName")])
	if err != nil {
		return nil, err
	}
	inventoryItemResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("inventoryItemResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Create(req.Context(), resourceGroupNameParam, vmmServerNameParam, inventoryItemResourceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).InventoryItem, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *InventoryItemsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if i.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ScVmm/vmmServers/(?P<vmmServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/inventoryItems/(?P<inventoryItemResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vmmServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmmServerName")])
	if err != nil {
		return nil, err
	}
	inventoryItemResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("inventoryItemResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Delete(req.Context(), resourceGroupNameParam, vmmServerNameParam, inventoryItemResourceNameParam, nil)
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

func (i *InventoryItemsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ScVmm/vmmServers/(?P<vmmServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/inventoryItems/(?P<inventoryItemResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vmmServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmmServerName")])
	if err != nil {
		return nil, err
	}
	inventoryItemResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("inventoryItemResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, vmmServerNameParam, inventoryItemResourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).InventoryItem, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *InventoryItemsServerTransport) dispatchNewListByVmmServerPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListByVmmServerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByVmmServerPager not implemented")}
	}
	newListByVmmServerPager := i.newListByVmmServerPager.get(req)
	if newListByVmmServerPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ScVmm/vmmServers/(?P<vmmServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/inventoryItems`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmmServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmmServerName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListByVmmServerPager(resourceGroupNameParam, vmmServerNameParam, nil)
		newListByVmmServerPager = &resp
		i.newListByVmmServerPager.add(req, newListByVmmServerPager)
		server.PagerResponderInjectNextLinks(newListByVmmServerPager, req, func(page *armscvmm.InventoryItemsClientListByVmmServerResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByVmmServerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListByVmmServerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByVmmServerPager) {
		i.newListByVmmServerPager.remove(req)
	}
	return resp, nil
}

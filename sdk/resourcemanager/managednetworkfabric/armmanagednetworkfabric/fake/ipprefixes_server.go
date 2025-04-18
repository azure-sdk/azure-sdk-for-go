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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
	"net/http"
	"net/url"
	"regexp"
)

// IPPrefixesServer is a fake server for instances of the armmanagednetworkfabric.IPPrefixesClient type.
type IPPrefixesServer struct {
	// BeginCreate is the fake for method IPPrefixesClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, ipPrefixName string, body armmanagednetworkfabric.IPPrefix, options *armmanagednetworkfabric.IPPrefixesClientBeginCreateOptions) (resp azfake.PollerResponder[armmanagednetworkfabric.IPPrefixesClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method IPPrefixesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, ipPrefixName string, options *armmanagednetworkfabric.IPPrefixesClientBeginDeleteOptions) (resp azfake.PollerResponder[armmanagednetworkfabric.IPPrefixesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method IPPrefixesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, ipPrefixName string, options *armmanagednetworkfabric.IPPrefixesClientGetOptions) (resp azfake.Responder[armmanagednetworkfabric.IPPrefixesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method IPPrefixesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armmanagednetworkfabric.IPPrefixesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armmanagednetworkfabric.IPPrefixesClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method IPPrefixesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armmanagednetworkfabric.IPPrefixesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armmanagednetworkfabric.IPPrefixesClientListBySubscriptionResponse])

	// BeginUpdate is the fake for method IPPrefixesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, ipPrefixName string, body armmanagednetworkfabric.IPPrefixPatch, options *armmanagednetworkfabric.IPPrefixesClientBeginUpdateOptions) (resp azfake.PollerResponder[armmanagednetworkfabric.IPPrefixesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewIPPrefixesServerTransport creates a new instance of IPPrefixesServerTransport with the provided implementation.
// The returned IPPrefixesServerTransport instance is connected to an instance of armmanagednetworkfabric.IPPrefixesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewIPPrefixesServerTransport(srv *IPPrefixesServer) *IPPrefixesServerTransport {
	return &IPPrefixesServerTransport{
		srv:                         srv,
		beginCreate:                 newTracker[azfake.PollerResponder[armmanagednetworkfabric.IPPrefixesClientCreateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armmanagednetworkfabric.IPPrefixesClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armmanagednetworkfabric.IPPrefixesClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armmanagednetworkfabric.IPPrefixesClientListBySubscriptionResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armmanagednetworkfabric.IPPrefixesClientUpdateResponse]](),
	}
}

// IPPrefixesServerTransport connects instances of armmanagednetworkfabric.IPPrefixesClient to instances of IPPrefixesServer.
// Don't use this type directly, use NewIPPrefixesServerTransport instead.
type IPPrefixesServerTransport struct {
	srv                         *IPPrefixesServer
	beginCreate                 *tracker[azfake.PollerResponder[armmanagednetworkfabric.IPPrefixesClientCreateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armmanagednetworkfabric.IPPrefixesClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armmanagednetworkfabric.IPPrefixesClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armmanagednetworkfabric.IPPrefixesClientListBySubscriptionResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armmanagednetworkfabric.IPPrefixesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for IPPrefixesServerTransport.
func (i *IPPrefixesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return i.dispatchToMethodFake(req, method)
}

func (i *IPPrefixesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if ipPrefixesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = ipPrefixesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "IPPrefixesClient.BeginCreate":
				res.resp, res.err = i.dispatchBeginCreate(req)
			case "IPPrefixesClient.BeginDelete":
				res.resp, res.err = i.dispatchBeginDelete(req)
			case "IPPrefixesClient.Get":
				res.resp, res.err = i.dispatchGet(req)
			case "IPPrefixesClient.NewListByResourceGroupPager":
				res.resp, res.err = i.dispatchNewListByResourceGroupPager(req)
			case "IPPrefixesClient.NewListBySubscriptionPager":
				res.resp, res.err = i.dispatchNewListBySubscriptionPager(req)
			case "IPPrefixesClient.BeginUpdate":
				res.resp, res.err = i.dispatchBeginUpdate(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (i *IPPrefixesServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if i.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := i.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/ipPrefixes/(?P<ipPrefixName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmanagednetworkfabric.IPPrefix](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		ipPrefixNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ipPrefixName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginCreate(req.Context(), resourceGroupNameParam, ipPrefixNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		i.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		i.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		i.beginCreate.remove(req)
	}

	return resp, nil
}

func (i *IPPrefixesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if i.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := i.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/ipPrefixes/(?P<ipPrefixName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		ipPrefixNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ipPrefixName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginDelete(req.Context(), resourceGroupNameParam, ipPrefixNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		i.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		i.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		i.beginDelete.remove(req)
	}

	return resp, nil
}

func (i *IPPrefixesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/ipPrefixes/(?P<ipPrefixName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	ipPrefixNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ipPrefixName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, ipPrefixNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IPPrefix, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IPPrefixesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := i.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/ipPrefixes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		i.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armmanagednetworkfabric.IPPrefixesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		i.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (i *IPPrefixesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := i.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/ipPrefixes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := i.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		i.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armmanagednetworkfabric.IPPrefixesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		i.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (i *IPPrefixesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := i.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/ipPrefixes/(?P<ipPrefixName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmanagednetworkfabric.IPPrefixPatch](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		ipPrefixNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ipPrefixName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginUpdate(req.Context(), resourceGroupNameParam, ipPrefixNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		i.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		i.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		i.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to IPPrefixesServerTransport
var ipPrefixesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateAccessesServer is a fake server for instances of the armchaos.PrivateAccessesClient type.
type PrivateAccessesServer struct {
	// BeginCreateOrUpdate is the fake for method PrivateAccessesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, privateAccessName string, privateAccess armchaos.PrivateAccess, options *armchaos.PrivateAccessesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armchaos.PrivateAccessesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PrivateAccessesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, privateAccessName string, options *armchaos.PrivateAccessesClientBeginDeleteOptions) (resp azfake.PollerResponder[armchaos.PrivateAccessesClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginDeleteAPrivateEndpointConnection is the fake for method PrivateAccessesClient.BeginDeleteAPrivateEndpointConnection
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDeleteAPrivateEndpointConnection func(ctx context.Context, resourceGroupName string, privateAccessName string, privateEndpointConnectionName string, options *armchaos.PrivateAccessesClientBeginDeleteAPrivateEndpointConnectionOptions) (resp azfake.PollerResponder[armchaos.PrivateAccessesClientDeleteAPrivateEndpointConnectionResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PrivateAccessesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, privateAccessName string, options *armchaos.PrivateAccessesClientGetOptions) (resp azfake.Responder[armchaos.PrivateAccessesClientGetResponse], errResp azfake.ErrorResponder)

	// GetAPrivateEndpointConnection is the fake for method PrivateAccessesClient.GetAPrivateEndpointConnection
	// HTTP status codes to indicate success: http.StatusOK
	GetAPrivateEndpointConnection func(ctx context.Context, resourceGroupName string, privateAccessName string, privateEndpointConnectionName string, options *armchaos.PrivateAccessesClientGetAPrivateEndpointConnectionOptions) (resp azfake.Responder[armchaos.PrivateAccessesClientGetAPrivateEndpointConnectionResponse], errResp azfake.ErrorResponder)

	// NewGetPrivateLinkResourcesPager is the fake for method PrivateAccessesClient.NewGetPrivateLinkResourcesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetPrivateLinkResourcesPager func(resourceGroupName string, privateAccessName string, options *armchaos.PrivateAccessesClientGetPrivateLinkResourcesOptions) (resp azfake.PagerResponder[armchaos.PrivateAccessesClientGetPrivateLinkResourcesResponse])

	// NewListPager is the fake for method PrivateAccessesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armchaos.PrivateAccessesClientListOptions) (resp azfake.PagerResponder[armchaos.PrivateAccessesClientListResponse])

	// NewListAllPager is the fake for method PrivateAccessesClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armchaos.PrivateAccessesClientListAllOptions) (resp azfake.PagerResponder[armchaos.PrivateAccessesClientListAllResponse])

	// NewListPrivateEndpointConnectionsPager is the fake for method PrivateAccessesClient.NewListPrivateEndpointConnectionsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPrivateEndpointConnectionsPager func(resourceGroupName string, privateAccessName string, options *armchaos.PrivateAccessesClientListPrivateEndpointConnectionsOptions) (resp azfake.PagerResponder[armchaos.PrivateAccessesClientListPrivateEndpointConnectionsResponse])

	// BeginUpdate is the fake for method PrivateAccessesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, privateAccessName string, privateAccessPatch armchaos.PrivateAccessPatch, options *armchaos.PrivateAccessesClientBeginUpdateOptions) (resp azfake.PollerResponder[armchaos.PrivateAccessesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewPrivateAccessesServerTransport creates a new instance of PrivateAccessesServerTransport with the provided implementation.
// The returned PrivateAccessesServerTransport instance is connected to an instance of armchaos.PrivateAccessesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateAccessesServerTransport(srv *PrivateAccessesServer) *PrivateAccessesServerTransport {
	return &PrivateAccessesServerTransport{
		srv:                                    srv,
		beginCreateOrUpdate:                    newTracker[azfake.PollerResponder[armchaos.PrivateAccessesClientCreateOrUpdateResponse]](),
		beginDelete:                            newTracker[azfake.PollerResponder[armchaos.PrivateAccessesClientDeleteResponse]](),
		beginDeleteAPrivateEndpointConnection:  newTracker[azfake.PollerResponder[armchaos.PrivateAccessesClientDeleteAPrivateEndpointConnectionResponse]](),
		newGetPrivateLinkResourcesPager:        newTracker[azfake.PagerResponder[armchaos.PrivateAccessesClientGetPrivateLinkResourcesResponse]](),
		newListPager:                           newTracker[azfake.PagerResponder[armchaos.PrivateAccessesClientListResponse]](),
		newListAllPager:                        newTracker[azfake.PagerResponder[armchaos.PrivateAccessesClientListAllResponse]](),
		newListPrivateEndpointConnectionsPager: newTracker[azfake.PagerResponder[armchaos.PrivateAccessesClientListPrivateEndpointConnectionsResponse]](),
		beginUpdate:                            newTracker[azfake.PollerResponder[armchaos.PrivateAccessesClientUpdateResponse]](),
	}
}

// PrivateAccessesServerTransport connects instances of armchaos.PrivateAccessesClient to instances of PrivateAccessesServer.
// Don't use this type directly, use NewPrivateAccessesServerTransport instead.
type PrivateAccessesServerTransport struct {
	srv                                    *PrivateAccessesServer
	beginCreateOrUpdate                    *tracker[azfake.PollerResponder[armchaos.PrivateAccessesClientCreateOrUpdateResponse]]
	beginDelete                            *tracker[azfake.PollerResponder[armchaos.PrivateAccessesClientDeleteResponse]]
	beginDeleteAPrivateEndpointConnection  *tracker[azfake.PollerResponder[armchaos.PrivateAccessesClientDeleteAPrivateEndpointConnectionResponse]]
	newGetPrivateLinkResourcesPager        *tracker[azfake.PagerResponder[armchaos.PrivateAccessesClientGetPrivateLinkResourcesResponse]]
	newListPager                           *tracker[azfake.PagerResponder[armchaos.PrivateAccessesClientListResponse]]
	newListAllPager                        *tracker[azfake.PagerResponder[armchaos.PrivateAccessesClientListAllResponse]]
	newListPrivateEndpointConnectionsPager *tracker[azfake.PagerResponder[armchaos.PrivateAccessesClientListPrivateEndpointConnectionsResponse]]
	beginUpdate                            *tracker[azfake.PollerResponder[armchaos.PrivateAccessesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for PrivateAccessesServerTransport.
func (p *PrivateAccessesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PrivateAccessesClient.BeginCreateOrUpdate":
		resp, err = p.dispatchBeginCreateOrUpdate(req)
	case "PrivateAccessesClient.BeginDelete":
		resp, err = p.dispatchBeginDelete(req)
	case "PrivateAccessesClient.BeginDeleteAPrivateEndpointConnection":
		resp, err = p.dispatchBeginDeleteAPrivateEndpointConnection(req)
	case "PrivateAccessesClient.Get":
		resp, err = p.dispatchGet(req)
	case "PrivateAccessesClient.GetAPrivateEndpointConnection":
		resp, err = p.dispatchGetAPrivateEndpointConnection(req)
	case "PrivateAccessesClient.NewGetPrivateLinkResourcesPager":
		resp, err = p.dispatchNewGetPrivateLinkResourcesPager(req)
	case "PrivateAccessesClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	case "PrivateAccessesClient.NewListAllPager":
		resp, err = p.dispatchNewListAllPager(req)
	case "PrivateAccessesClient.NewListPrivateEndpointConnectionsPager":
		resp, err = p.dispatchNewListPrivateEndpointConnectionsPager(req)
	case "PrivateAccessesClient.BeginUpdate":
		resp, err = p.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PrivateAccessesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := p.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Chaos/privateAccesses/(?P<privateAccessName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armchaos.PrivateAccess](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateAccessNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateAccessName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, privateAccessNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		p.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		p.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		p.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (p *PrivateAccessesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Chaos/privateAccesses/(?P<privateAccessName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateAccessNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateAccessName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, privateAccessNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		p.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		p.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		p.beginDelete.remove(req)
	}

	return resp, nil
}

func (p *PrivateAccessesServerTransport) dispatchBeginDeleteAPrivateEndpointConnection(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDeleteAPrivateEndpointConnection == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDeleteAPrivateEndpointConnection not implemented")}
	}
	beginDeleteAPrivateEndpointConnection := p.beginDeleteAPrivateEndpointConnection.get(req)
	if beginDeleteAPrivateEndpointConnection == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Chaos/privateAccesses/(?P<privateAccessName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateAccessNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateAccessName")])
		if err != nil {
			return nil, err
		}
		privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDeleteAPrivateEndpointConnection(req.Context(), resourceGroupNameParam, privateAccessNameParam, privateEndpointConnectionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDeleteAPrivateEndpointConnection = &respr
		p.beginDeleteAPrivateEndpointConnection.add(req, beginDeleteAPrivateEndpointConnection)
	}

	resp, err := server.PollerResponderNext(beginDeleteAPrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		p.beginDeleteAPrivateEndpointConnection.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDeleteAPrivateEndpointConnection) {
		p.beginDeleteAPrivateEndpointConnection.remove(req)
	}

	return resp, nil
}

func (p *PrivateAccessesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Chaos/privateAccesses/(?P<privateAccessName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	privateAccessNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateAccessName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, privateAccessNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateAccess, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateAccessesServerTransport) dispatchGetAPrivateEndpointConnection(req *http.Request) (*http.Response, error) {
	if p.srv.GetAPrivateEndpointConnection == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAPrivateEndpointConnection not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Chaos/privateAccesses/(?P<privateAccessName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	privateAccessNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateAccessName")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.GetAPrivateEndpointConnection(req.Context(), resourceGroupNameParam, privateAccessNameParam, privateEndpointConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateAccessesServerTransport) dispatchNewGetPrivateLinkResourcesPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetPrivateLinkResourcesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetPrivateLinkResourcesPager not implemented")}
	}
	newGetPrivateLinkResourcesPager := p.newGetPrivateLinkResourcesPager.get(req)
	if newGetPrivateLinkResourcesPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Chaos/privateAccesses/(?P<privateAccessName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateAccessNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateAccessName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewGetPrivateLinkResourcesPager(resourceGroupNameParam, privateAccessNameParam, nil)
		newGetPrivateLinkResourcesPager = &resp
		p.newGetPrivateLinkResourcesPager.add(req, newGetPrivateLinkResourcesPager)
		server.PagerResponderInjectNextLinks(newGetPrivateLinkResourcesPager, req, func(page *armchaos.PrivateAccessesClientGetPrivateLinkResourcesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newGetPrivateLinkResourcesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newGetPrivateLinkResourcesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newGetPrivateLinkResourcesPager) {
		p.newGetPrivateLinkResourcesPager.remove(req)
	}
	return resp, nil
}

func (p *PrivateAccessesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Chaos/privateAccesses`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		continuationTokenUnescaped, err := url.QueryUnescape(qp.Get("continuationToken"))
		if err != nil {
			return nil, err
		}
		continuationTokenParam := getOptional(continuationTokenUnescaped)
		var options *armchaos.PrivateAccessesClientListOptions
		if continuationTokenParam != nil {
			options = &armchaos.PrivateAccessesClientListOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := p.srv.NewListPager(resourceGroupNameParam, options)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armchaos.PrivateAccessesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		p.newListPager.remove(req)
	}
	return resp, nil
}

func (p *PrivateAccessesServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAllPager not implemented")}
	}
	newListAllPager := p.newListAllPager.get(req)
	if newListAllPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Chaos/privateAccesses`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		continuationTokenUnescaped, err := url.QueryUnescape(qp.Get("continuationToken"))
		if err != nil {
			return nil, err
		}
		continuationTokenParam := getOptional(continuationTokenUnescaped)
		var options *armchaos.PrivateAccessesClientListAllOptions
		if continuationTokenParam != nil {
			options = &armchaos.PrivateAccessesClientListAllOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := p.srv.NewListAllPager(options)
		newListAllPager = &resp
		p.newListAllPager.add(req, newListAllPager)
		server.PagerResponderInjectNextLinks(newListAllPager, req, func(page *armchaos.PrivateAccessesClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListAllPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListAllPager) {
		p.newListAllPager.remove(req)
	}
	return resp, nil
}

func (p *PrivateAccessesServerTransport) dispatchNewListPrivateEndpointConnectionsPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPrivateEndpointConnectionsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPrivateEndpointConnectionsPager not implemented")}
	}
	newListPrivateEndpointConnectionsPager := p.newListPrivateEndpointConnectionsPager.get(req)
	if newListPrivateEndpointConnectionsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Chaos/privateAccesses/(?P<privateAccessName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateAccessNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateAccessName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListPrivateEndpointConnectionsPager(resourceGroupNameParam, privateAccessNameParam, nil)
		newListPrivateEndpointConnectionsPager = &resp
		p.newListPrivateEndpointConnectionsPager.add(req, newListPrivateEndpointConnectionsPager)
		server.PagerResponderInjectNextLinks(newListPrivateEndpointConnectionsPager, req, func(page *armchaos.PrivateAccessesClientListPrivateEndpointConnectionsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPrivateEndpointConnectionsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListPrivateEndpointConnectionsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPrivateEndpointConnectionsPager) {
		p.newListPrivateEndpointConnectionsPager.remove(req)
	}
	return resp, nil
}

func (p *PrivateAccessesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := p.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Chaos/privateAccesses/(?P<privateAccessName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armchaos.PrivateAccessPatch](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateAccessNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateAccessName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginUpdate(req.Context(), resourceGroupNameParam, privateAccessNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		p.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		p.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		p.beginUpdate.remove(req)
	}

	return resp, nil
}

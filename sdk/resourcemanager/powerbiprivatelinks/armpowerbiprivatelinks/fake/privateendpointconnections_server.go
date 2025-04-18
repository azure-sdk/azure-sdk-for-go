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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks/v2"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateEndpointConnectionsServer is a fake server for instances of the armpowerbiprivatelinks.PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsServer struct {
	// Create is the fake for method PrivateEndpointConnectionsClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, azureResourceName string, privateEndpointName string, privateEndpointConnection armpowerbiprivatelinks.PrivateEndpointConnection, options *armpowerbiprivatelinks.PrivateEndpointConnectionsClientCreateOptions) (resp azfake.Responder[armpowerbiprivatelinks.PrivateEndpointConnectionsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PrivateEndpointConnectionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, azureResourceName string, privateEndpointName string, options *armpowerbiprivatelinks.PrivateEndpointConnectionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armpowerbiprivatelinks.PrivateEndpointConnectionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PrivateEndpointConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, azureResourceName string, privateEndpointName string, options *armpowerbiprivatelinks.PrivateEndpointConnectionsClientGetOptions) (resp azfake.Responder[armpowerbiprivatelinks.PrivateEndpointConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourcePager is the fake for method PrivateEndpointConnectionsClient.NewListByResourcePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourcePager func(resourceGroupName string, azureResourceName string, options *armpowerbiprivatelinks.PrivateEndpointConnectionsClientListByResourceOptions) (resp azfake.PagerResponder[armpowerbiprivatelinks.PrivateEndpointConnectionsClientListByResourceResponse])
}

// NewPrivateEndpointConnectionsServerTransport creates a new instance of PrivateEndpointConnectionsServerTransport with the provided implementation.
// The returned PrivateEndpointConnectionsServerTransport instance is connected to an instance of armpowerbiprivatelinks.PrivateEndpointConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateEndpointConnectionsServerTransport(srv *PrivateEndpointConnectionsServer) *PrivateEndpointConnectionsServerTransport {
	return &PrivateEndpointConnectionsServerTransport{
		srv:                    srv,
		beginDelete:            newTracker[azfake.PollerResponder[armpowerbiprivatelinks.PrivateEndpointConnectionsClientDeleteResponse]](),
		newListByResourcePager: newTracker[azfake.PagerResponder[armpowerbiprivatelinks.PrivateEndpointConnectionsClientListByResourceResponse]](),
	}
}

// PrivateEndpointConnectionsServerTransport connects instances of armpowerbiprivatelinks.PrivateEndpointConnectionsClient to instances of PrivateEndpointConnectionsServer.
// Don't use this type directly, use NewPrivateEndpointConnectionsServerTransport instead.
type PrivateEndpointConnectionsServerTransport struct {
	srv                    *PrivateEndpointConnectionsServer
	beginDelete            *tracker[azfake.PollerResponder[armpowerbiprivatelinks.PrivateEndpointConnectionsClientDeleteResponse]]
	newListByResourcePager *tracker[azfake.PagerResponder[armpowerbiprivatelinks.PrivateEndpointConnectionsClientListByResourceResponse]]
}

// Do implements the policy.Transporter interface for PrivateEndpointConnectionsServerTransport.
func (p *PrivateEndpointConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PrivateEndpointConnectionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if privateEndpointConnectionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = privateEndpointConnectionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "PrivateEndpointConnectionsClient.Create":
				res.resp, res.err = p.dispatchCreate(req)
			case "PrivateEndpointConnectionsClient.BeginDelete":
				res.resp, res.err = p.dispatchBeginDelete(req)
			case "PrivateEndpointConnectionsClient.Get":
				res.resp, res.err = p.dispatchGet(req)
			case "PrivateEndpointConnectionsClient.NewListByResourcePager":
				res.resp, res.err = p.dispatchNewListByResourcePager(req)
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

func (p *PrivateEndpointConnectionsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if p.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.PowerBI/privateLinkServicesForPowerBI/(?P<azureResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armpowerbiprivatelinks.PrivateEndpointConnection](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	azureResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureResourceName")])
	if err != nil {
		return nil, err
	}
	privateEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Create(req.Context(), resourceGroupNameParam, azureResourceNameParam, privateEndpointNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.PowerBI/privateLinkServicesForPowerBI/(?P<azureResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		azureResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureResourceName")])
		if err != nil {
			return nil, err
		}
		privateEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, azureResourceNameParam, privateEndpointNameParam, nil)
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

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		p.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		p.beginDelete.remove(req)
	}

	return resp, nil
}

func (p *PrivateEndpointConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.PowerBI/privateLinkServicesForPowerBI/(?P<azureResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	azureResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureResourceName")])
	if err != nil {
		return nil, err
	}
	privateEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, azureResourceNameParam, privateEndpointNameParam, nil)
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

func (p *PrivateEndpointConnectionsServerTransport) dispatchNewListByResourcePager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByResourcePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourcePager not implemented")}
	}
	newListByResourcePager := p.newListByResourcePager.get(req)
	if newListByResourcePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.PowerBI/privateLinkServicesForPowerBI/(?P<azureResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		azureResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureResourceName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListByResourcePager(resourceGroupNameParam, azureResourceNameParam, nil)
		newListByResourcePager = &resp
		p.newListByResourcePager.add(req, newListByResourcePager)
		server.PagerResponderInjectNextLinks(newListByResourcePager, req, func(page *armpowerbiprivatelinks.PrivateEndpointConnectionsClientListByResourceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourcePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByResourcePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourcePager) {
		p.newListByResourcePager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to PrivateEndpointConnectionsServerTransport
var privateEndpointConnectionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

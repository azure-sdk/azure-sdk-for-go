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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
	"net/http"
	"net/url"
	"regexp"
)

// ApplicationGatewayPrivateEndpointConnectionsServer is a fake server for instances of the armnetwork.ApplicationGatewayPrivateEndpointConnectionsClient type.
type ApplicationGatewayPrivateEndpointConnectionsServer struct {
	// BeginDelete is the fake for method ApplicationGatewayPrivateEndpointConnectionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, options *armnetwork.ApplicationGatewayPrivateEndpointConnectionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.ApplicationGatewayPrivateEndpointConnectionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ApplicationGatewayPrivateEndpointConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, options *armnetwork.ApplicationGatewayPrivateEndpointConnectionsClientGetOptions) (resp azfake.Responder[armnetwork.ApplicationGatewayPrivateEndpointConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ApplicationGatewayPrivateEndpointConnectionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, applicationGatewayName string, options *armnetwork.ApplicationGatewayPrivateEndpointConnectionsClientListOptions) (resp azfake.PagerResponder[armnetwork.ApplicationGatewayPrivateEndpointConnectionsClientListResponse])

	// BeginUpdate is the fake for method ApplicationGatewayPrivateEndpointConnectionsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, applicationGatewayName string, connectionName string, parameters armnetwork.ApplicationGatewayPrivateEndpointConnection, options *armnetwork.ApplicationGatewayPrivateEndpointConnectionsClientBeginUpdateOptions) (resp azfake.PollerResponder[armnetwork.ApplicationGatewayPrivateEndpointConnectionsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewApplicationGatewayPrivateEndpointConnectionsServerTransport creates a new instance of ApplicationGatewayPrivateEndpointConnectionsServerTransport with the provided implementation.
// The returned ApplicationGatewayPrivateEndpointConnectionsServerTransport instance is connected to an instance of armnetwork.ApplicationGatewayPrivateEndpointConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewApplicationGatewayPrivateEndpointConnectionsServerTransport(srv *ApplicationGatewayPrivateEndpointConnectionsServer) *ApplicationGatewayPrivateEndpointConnectionsServerTransport {
	return &ApplicationGatewayPrivateEndpointConnectionsServerTransport{
		srv:          srv,
		beginDelete:  newTracker[azfake.PollerResponder[armnetwork.ApplicationGatewayPrivateEndpointConnectionsClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armnetwork.ApplicationGatewayPrivateEndpointConnectionsClientListResponse]](),
		beginUpdate:  newTracker[azfake.PollerResponder[armnetwork.ApplicationGatewayPrivateEndpointConnectionsClientUpdateResponse]](),
	}
}

// ApplicationGatewayPrivateEndpointConnectionsServerTransport connects instances of armnetwork.ApplicationGatewayPrivateEndpointConnectionsClient to instances of ApplicationGatewayPrivateEndpointConnectionsServer.
// Don't use this type directly, use NewApplicationGatewayPrivateEndpointConnectionsServerTransport instead.
type ApplicationGatewayPrivateEndpointConnectionsServerTransport struct {
	srv          *ApplicationGatewayPrivateEndpointConnectionsServer
	beginDelete  *tracker[azfake.PollerResponder[armnetwork.ApplicationGatewayPrivateEndpointConnectionsClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armnetwork.ApplicationGatewayPrivateEndpointConnectionsClientListResponse]]
	beginUpdate  *tracker[azfake.PollerResponder[armnetwork.ApplicationGatewayPrivateEndpointConnectionsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for ApplicationGatewayPrivateEndpointConnectionsServerTransport.
func (a *ApplicationGatewayPrivateEndpointConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *ApplicationGatewayPrivateEndpointConnectionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if applicationGatewayPrivateEndpointConnectionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = applicationGatewayPrivateEndpointConnectionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ApplicationGatewayPrivateEndpointConnectionsClient.BeginDelete":
				res.resp, res.err = a.dispatchBeginDelete(req)
			case "ApplicationGatewayPrivateEndpointConnectionsClient.Get":
				res.resp, res.err = a.dispatchGet(req)
			case "ApplicationGatewayPrivateEndpointConnectionsClient.NewListPager":
				res.resp, res.err = a.dispatchNewListPager(req)
			case "ApplicationGatewayPrivateEndpointConnectionsClient.BeginUpdate":
				res.resp, res.err = a.dispatchBeginUpdate(req)
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

func (a *ApplicationGatewayPrivateEndpointConnectionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/applicationGateways/(?P<applicationGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		applicationGatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationGatewayName")])
		if err != nil {
			return nil, err
		}
		connectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameParam, applicationGatewayNameParam, connectionNameParam, nil)
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

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		a.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		a.beginDelete.remove(req)
	}

	return resp, nil
}

func (a *ApplicationGatewayPrivateEndpointConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/applicationGateways/(?P<applicationGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	applicationGatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationGatewayName")])
	if err != nil {
		return nil, err
	}
	connectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, applicationGatewayNameParam, connectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ApplicationGatewayPrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationGatewayPrivateEndpointConnectionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/applicationGateways/(?P<applicationGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		applicationGatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationGatewayName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListPager(resourceGroupNameParam, applicationGatewayNameParam, nil)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.ApplicationGatewayPrivateEndpointConnectionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}

func (a *ApplicationGatewayPrivateEndpointConnectionsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := a.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/applicationGateways/(?P<applicationGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.ApplicationGatewayPrivateEndpointConnection](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		applicationGatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationGatewayName")])
		if err != nil {
			return nil, err
		}
		connectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginUpdate(req.Context(), resourceGroupNameParam, applicationGatewayNameParam, connectionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		a.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		a.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ApplicationGatewayPrivateEndpointConnectionsServerTransport
var applicationGatewayPrivateEndpointConnectionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

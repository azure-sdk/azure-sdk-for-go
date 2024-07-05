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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate/v2"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateEndpointConnectionProxiesServer is a fake server for instances of the armdeviceupdate.PrivateEndpointConnectionProxiesClient type.
type PrivateEndpointConnectionProxiesServer struct {
	// BeginCreateOrUpdate is the fake for method PrivateEndpointConnectionProxiesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionProxyID string, privateEndpointConnectionProxy armdeviceupdate.PrivateEndpointConnectionProxy, options *armdeviceupdate.PrivateEndpointConnectionProxiesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armdeviceupdate.PrivateEndpointConnectionProxiesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PrivateEndpointConnectionProxiesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionProxyID string, options *armdeviceupdate.PrivateEndpointConnectionProxiesClientBeginDeleteOptions) (resp azfake.PollerResponder[armdeviceupdate.PrivateEndpointConnectionProxiesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PrivateEndpointConnectionProxiesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionProxyID string, options *armdeviceupdate.PrivateEndpointConnectionProxiesClientGetOptions) (resp azfake.Responder[armdeviceupdate.PrivateEndpointConnectionProxiesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByAccountPager is the fake for method PrivateEndpointConnectionProxiesClient.NewListByAccountPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByAccountPager func(resourceGroupName string, accountName string, options *armdeviceupdate.PrivateEndpointConnectionProxiesClientListByAccountOptions) (resp azfake.PagerResponder[armdeviceupdate.PrivateEndpointConnectionProxiesClientListByAccountResponse])

	// UpdatePrivateEndpointProperties is the fake for method PrivateEndpointConnectionProxiesClient.UpdatePrivateEndpointProperties
	// HTTP status codes to indicate success: http.StatusOK
	UpdatePrivateEndpointProperties func(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionProxyID string, privateEndpointUpdate armdeviceupdate.PrivateEndpointUpdate, options *armdeviceupdate.PrivateEndpointConnectionProxiesClientUpdatePrivateEndpointPropertiesOptions) (resp azfake.Responder[armdeviceupdate.PrivateEndpointConnectionProxiesClientUpdatePrivateEndpointPropertiesResponse], errResp azfake.ErrorResponder)

	// Validate is the fake for method PrivateEndpointConnectionProxiesClient.Validate
	// HTTP status codes to indicate success: http.StatusOK
	Validate func(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionProxyID string, privateEndpointConnectionProxy armdeviceupdate.PrivateEndpointConnectionProxy, options *armdeviceupdate.PrivateEndpointConnectionProxiesClientValidateOptions) (resp azfake.Responder[armdeviceupdate.PrivateEndpointConnectionProxiesClientValidateResponse], errResp azfake.ErrorResponder)
}

// NewPrivateEndpointConnectionProxiesServerTransport creates a new instance of PrivateEndpointConnectionProxiesServerTransport with the provided implementation.
// The returned PrivateEndpointConnectionProxiesServerTransport instance is connected to an instance of armdeviceupdate.PrivateEndpointConnectionProxiesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateEndpointConnectionProxiesServerTransport(srv *PrivateEndpointConnectionProxiesServer) *PrivateEndpointConnectionProxiesServerTransport {
	return &PrivateEndpointConnectionProxiesServerTransport{
		srv:                   srv,
		beginCreateOrUpdate:   newTracker[azfake.PollerResponder[armdeviceupdate.PrivateEndpointConnectionProxiesClientCreateOrUpdateResponse]](),
		beginDelete:           newTracker[azfake.PollerResponder[armdeviceupdate.PrivateEndpointConnectionProxiesClientDeleteResponse]](),
		newListByAccountPager: newTracker[azfake.PagerResponder[armdeviceupdate.PrivateEndpointConnectionProxiesClientListByAccountResponse]](),
	}
}

// PrivateEndpointConnectionProxiesServerTransport connects instances of armdeviceupdate.PrivateEndpointConnectionProxiesClient to instances of PrivateEndpointConnectionProxiesServer.
// Don't use this type directly, use NewPrivateEndpointConnectionProxiesServerTransport instead.
type PrivateEndpointConnectionProxiesServerTransport struct {
	srv                   *PrivateEndpointConnectionProxiesServer
	beginCreateOrUpdate   *tracker[azfake.PollerResponder[armdeviceupdate.PrivateEndpointConnectionProxiesClientCreateOrUpdateResponse]]
	beginDelete           *tracker[azfake.PollerResponder[armdeviceupdate.PrivateEndpointConnectionProxiesClientDeleteResponse]]
	newListByAccountPager *tracker[azfake.PagerResponder[armdeviceupdate.PrivateEndpointConnectionProxiesClientListByAccountResponse]]
}

// Do implements the policy.Transporter interface for PrivateEndpointConnectionProxiesServerTransport.
func (p *PrivateEndpointConnectionProxiesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PrivateEndpointConnectionProxiesClient.BeginCreateOrUpdate":
		resp, err = p.dispatchBeginCreateOrUpdate(req)
	case "PrivateEndpointConnectionProxiesClient.BeginDelete":
		resp, err = p.dispatchBeginDelete(req)
	case "PrivateEndpointConnectionProxiesClient.Get":
		resp, err = p.dispatchGet(req)
	case "PrivateEndpointConnectionProxiesClient.NewListByAccountPager":
		resp, err = p.dispatchNewListByAccountPager(req)
	case "PrivateEndpointConnectionProxiesClient.UpdatePrivateEndpointProperties":
		resp, err = p.dispatchUpdatePrivateEndpointProperties(req)
	case "PrivateEndpointConnectionProxiesClient.Validate":
		resp, err = p.dispatchValidate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PrivateEndpointConnectionProxiesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := p.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeviceUpdate/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnectionProxies/(?P<privateEndpointConnectionProxyId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdeviceupdate.PrivateEndpointConnectionProxy](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		privateEndpointConnectionProxyIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionProxyId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, accountNameParam, privateEndpointConnectionProxyIDParam, body, nil)
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

	if !contains([]int{http.StatusCreated}, resp.StatusCode) {
		p.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		p.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (p *PrivateEndpointConnectionProxiesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeviceUpdate/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnectionProxies/(?P<privateEndpointConnectionProxyId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		privateEndpointConnectionProxyIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionProxyId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, accountNameParam, privateEndpointConnectionProxyIDParam, nil)
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

func (p *PrivateEndpointConnectionProxiesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeviceUpdate/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnectionProxies/(?P<privateEndpointConnectionProxyId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionProxyIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionProxyId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, accountNameParam, privateEndpointConnectionProxyIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateEndpointConnectionProxy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionProxiesServerTransport) dispatchNewListByAccountPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByAccountPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByAccountPager not implemented")}
	}
	newListByAccountPager := p.newListByAccountPager.get(req)
	if newListByAccountPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeviceUpdate/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnectionProxies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListByAccountPager(resourceGroupNameParam, accountNameParam, nil)
		newListByAccountPager = &resp
		p.newListByAccountPager.add(req, newListByAccountPager)
	}
	resp, err := server.PagerResponderNext(newListByAccountPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByAccountPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByAccountPager) {
		p.newListByAccountPager.remove(req)
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionProxiesServerTransport) dispatchUpdatePrivateEndpointProperties(req *http.Request) (*http.Response, error) {
	if p.srv.UpdatePrivateEndpointProperties == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdatePrivateEndpointProperties not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeviceUpdate/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnectionProxies/(?P<privateEndpointConnectionProxyId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updatePrivateEndpointProperties`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdeviceupdate.PrivateEndpointUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionProxyIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionProxyId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.UpdatePrivateEndpointProperties(req.Context(), resourceGroupNameParam, accountNameParam, privateEndpointConnectionProxyIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionProxiesServerTransport) dispatchValidate(req *http.Request) (*http.Response, error) {
	if p.srv.Validate == nil {
		return nil, &nonRetriableError{errors.New("fake for method Validate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DeviceUpdate/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnectionProxies/(?P<privateEndpointConnectionProxyId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/validate`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdeviceupdate.PrivateEndpointConnectionProxy](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionProxyIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionProxyId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Validate(req.Context(), resourceGroupNameParam, accountNameParam, privateEndpointConnectionProxyIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

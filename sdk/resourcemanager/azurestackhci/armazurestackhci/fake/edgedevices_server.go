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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
	"net/http"
	"net/url"
	"regexp"
)

// EdgeDevicesServer is a fake server for instances of the armazurestackhci.EdgeDevicesClient type.
type EdgeDevicesServer struct {
	// BeginCreateOrUpdate is the fake for method EdgeDevicesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceURI string, edgeDeviceName string, resource armazurestackhci.EdgeDevice, options *armazurestackhci.EdgeDevicesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armazurestackhci.EdgeDevicesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method EdgeDevicesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceURI string, edgeDeviceName string, options *armazurestackhci.EdgeDevicesClientBeginDeleteOptions) (resp azfake.PollerResponder[armazurestackhci.EdgeDevicesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method EdgeDevicesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceURI string, edgeDeviceName string, options *armazurestackhci.EdgeDevicesClientGetOptions) (resp azfake.Responder[armazurestackhci.EdgeDevicesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method EdgeDevicesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceURI string, options *armazurestackhci.EdgeDevicesClientListOptions) (resp azfake.PagerResponder[armazurestackhci.EdgeDevicesClientListResponse])

	// BeginValidate is the fake for method EdgeDevicesClient.BeginValidate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginValidate func(ctx context.Context, resourceURI string, edgeDeviceName string, validateRequest armazurestackhci.ValidateRequest, options *armazurestackhci.EdgeDevicesClientBeginValidateOptions) (resp azfake.PollerResponder[armazurestackhci.EdgeDevicesClientValidateResponse], errResp azfake.ErrorResponder)
}

// NewEdgeDevicesServerTransport creates a new instance of EdgeDevicesServerTransport with the provided implementation.
// The returned EdgeDevicesServerTransport instance is connected to an instance of armazurestackhci.EdgeDevicesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewEdgeDevicesServerTransport(srv *EdgeDevicesServer) *EdgeDevicesServerTransport {
	return &EdgeDevicesServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armazurestackhci.EdgeDevicesClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armazurestackhci.EdgeDevicesClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armazurestackhci.EdgeDevicesClientListResponse]](),
		beginValidate:       newTracker[azfake.PollerResponder[armazurestackhci.EdgeDevicesClientValidateResponse]](),
	}
}

// EdgeDevicesServerTransport connects instances of armazurestackhci.EdgeDevicesClient to instances of EdgeDevicesServer.
// Don't use this type directly, use NewEdgeDevicesServerTransport instead.
type EdgeDevicesServerTransport struct {
	srv                 *EdgeDevicesServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armazurestackhci.EdgeDevicesClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armazurestackhci.EdgeDevicesClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armazurestackhci.EdgeDevicesClientListResponse]]
	beginValidate       *tracker[azfake.PollerResponder[armazurestackhci.EdgeDevicesClientValidateResponse]]
}

// Do implements the policy.Transporter interface for EdgeDevicesServerTransport.
func (e *EdgeDevicesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "EdgeDevicesClient.BeginCreateOrUpdate":
		resp, err = e.dispatchBeginCreateOrUpdate(req)
	case "EdgeDevicesClient.BeginDelete":
		resp, err = e.dispatchBeginDelete(req)
	case "EdgeDevicesClient.Get":
		resp, err = e.dispatchGet(req)
	case "EdgeDevicesClient.NewListPager":
		resp, err = e.dispatchNewListPager(req)
	case "EdgeDevicesClient.BeginValidate":
		resp, err = e.dispatchBeginValidate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *EdgeDevicesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := e.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureStackHCI/edgeDevices/(?P<edgeDeviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armazurestackhci.EdgeDevice](req)
		if err != nil {
			return nil, err
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		edgeDeviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("edgeDeviceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginCreateOrUpdate(req.Context(), resourceURIParam, edgeDeviceNameParam, body, nil)
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

func (e *EdgeDevicesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if e.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := e.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureStackHCI/edgeDevices/(?P<edgeDeviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		edgeDeviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("edgeDeviceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginDelete(req.Context(), resourceURIParam, edgeDeviceNameParam, nil)
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

func (e *EdgeDevicesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureStackHCI/edgeDevices/(?P<edgeDeviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
	if err != nil {
		return nil, err
	}
	edgeDeviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("edgeDeviceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceURIParam, edgeDeviceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EdgeDevice, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EdgeDevicesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := e.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureStackHCI/edgeDevices`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListPager(resourceURIParam, nil)
		newListPager = &resp
		e.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armazurestackhci.EdgeDevicesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		e.newListPager.remove(req)
	}
	return resp, nil
}

func (e *EdgeDevicesServerTransport) dispatchBeginValidate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginValidate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginValidate not implemented")}
	}
	beginValidate := e.beginValidate.get(req)
	if beginValidate == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureStackHCI/edgeDevices/(?P<edgeDeviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/validate`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armazurestackhci.ValidateRequest](req)
		if err != nil {
			return nil, err
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		edgeDeviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("edgeDeviceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginValidate(req.Context(), resourceURIParam, edgeDeviceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginValidate = &respr
		e.beginValidate.add(req, beginValidate)
	}

	resp, err := server.PollerResponderNext(beginValidate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		e.beginValidate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginValidate) {
		e.beginValidate.remove(req)
	}

	return resp, nil
}

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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw"
	"net/http"
	"regexp"
)

// PaloAltoNetworksCloudngfwServer is a fake server for instances of the armpanngfw.PaloAltoNetworksCloudngfwClient type.
type PaloAltoNetworksCloudngfwServer struct {
	// CreateProductSerialNumber is the fake for method PaloAltoNetworksCloudngfwClient.CreateProductSerialNumber
	// HTTP status codes to indicate success: http.StatusOK
	CreateProductSerialNumber func(ctx context.Context, options *armpanngfw.PaloAltoNetworksCloudngfwClientCreateProductSerialNumberOptions) (resp azfake.Responder[armpanngfw.PaloAltoNetworksCloudngfwClientCreateProductSerialNumberResponse], errResp azfake.ErrorResponder)

	// ListCloudManagerTenants is the fake for method PaloAltoNetworksCloudngfwClient.ListCloudManagerTenants
	// HTTP status codes to indicate success: http.StatusOK
	ListCloudManagerTenants func(ctx context.Context, options *armpanngfw.PaloAltoNetworksCloudngfwClientListCloudManagerTenantsOptions) (resp azfake.Responder[armpanngfw.PaloAltoNetworksCloudngfwClientListCloudManagerTenantsResponse], errResp azfake.ErrorResponder)

	// ListProductSerialNumberStatus is the fake for method PaloAltoNetworksCloudngfwClient.ListProductSerialNumberStatus
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNotFound
	ListProductSerialNumberStatus func(ctx context.Context, options *armpanngfw.PaloAltoNetworksCloudngfwClientListProductSerialNumberStatusOptions) (resp azfake.Responder[armpanngfw.PaloAltoNetworksCloudngfwClientListProductSerialNumberStatusResponse], errResp azfake.ErrorResponder)

	// ListSupportInfo is the fake for method PaloAltoNetworksCloudngfwClient.ListSupportInfo
	// HTTP status codes to indicate success: http.StatusOK
	ListSupportInfo func(ctx context.Context, options *armpanngfw.PaloAltoNetworksCloudngfwClientListSupportInfoOptions) (resp azfake.Responder[armpanngfw.PaloAltoNetworksCloudngfwClientListSupportInfoResponse], errResp azfake.ErrorResponder)
}

// NewPaloAltoNetworksCloudngfwServerTransport creates a new instance of PaloAltoNetworksCloudngfwServerTransport with the provided implementation.
// The returned PaloAltoNetworksCloudngfwServerTransport instance is connected to an instance of armpanngfw.PaloAltoNetworksCloudngfwClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPaloAltoNetworksCloudngfwServerTransport(srv *PaloAltoNetworksCloudngfwServer) *PaloAltoNetworksCloudngfwServerTransport {
	return &PaloAltoNetworksCloudngfwServerTransport{srv: srv}
}

// PaloAltoNetworksCloudngfwServerTransport connects instances of armpanngfw.PaloAltoNetworksCloudngfwClient to instances of PaloAltoNetworksCloudngfwServer.
// Don't use this type directly, use NewPaloAltoNetworksCloudngfwServerTransport instead.
type PaloAltoNetworksCloudngfwServerTransport struct {
	srv *PaloAltoNetworksCloudngfwServer
}

// Do implements the policy.Transporter interface for PaloAltoNetworksCloudngfwServerTransport.
func (p *PaloAltoNetworksCloudngfwServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PaloAltoNetworksCloudngfwClient.CreateProductSerialNumber":
		resp, err = p.dispatchCreateProductSerialNumber(req)
	case "PaloAltoNetworksCloudngfwClient.ListCloudManagerTenants":
		resp, err = p.dispatchListCloudManagerTenants(req)
	case "PaloAltoNetworksCloudngfwClient.ListProductSerialNumberStatus":
		resp, err = p.dispatchListProductSerialNumberStatus(req)
	case "PaloAltoNetworksCloudngfwClient.ListSupportInfo":
		resp, err = p.dispatchListSupportInfo(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PaloAltoNetworksCloudngfwServerTransport) dispatchCreateProductSerialNumber(req *http.Request) (*http.Response, error) {
	if p.srv.CreateProductSerialNumber == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateProductSerialNumber not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/PaloAltoNetworks\.Cloudngfw/createProductSerialNumber`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := p.srv.CreateProductSerialNumber(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProductSerialNumberRequestStatus, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PaloAltoNetworksCloudngfwServerTransport) dispatchListCloudManagerTenants(req *http.Request) (*http.Response, error) {
	if p.srv.ListCloudManagerTenants == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListCloudManagerTenants not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/PaloAltoNetworks\.Cloudngfw/listCloudManagerTenants`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := p.srv.ListCloudManagerTenants(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CloudManagerTenantList, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PaloAltoNetworksCloudngfwServerTransport) dispatchListProductSerialNumberStatus(req *http.Request) (*http.Response, error) {
	if p.srv.ListProductSerialNumberStatus == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListProductSerialNumberStatus not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/PaloAltoNetworks\.Cloudngfw/listProductSerialNumberStatus`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := p.srv.ListProductSerialNumberStatus(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNotFound}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNotFound", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProductSerialNumberStatus, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PaloAltoNetworksCloudngfwServerTransport) dispatchListSupportInfo(req *http.Request) (*http.Response, error) {
	if p.srv.ListSupportInfo == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListSupportInfo not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/PaloAltoNetworks\.Cloudngfw/listSupportInfo`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := p.srv.ListSupportInfo(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SupportInfoModel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
	"net/http"
	"net/url"
	"regexp"
)

// NetworkStatusServer is a fake server for instances of the armapimanagement.NetworkStatusClient type.
type NetworkStatusServer struct {
	// ListByLocation is the fake for method NetworkStatusClient.ListByLocation
	// HTTP status codes to indicate success: http.StatusOK
	ListByLocation func(ctx context.Context, resourceGroupName string, serviceName string, locationName string, options *armapimanagement.NetworkStatusClientListByLocationOptions) (resp azfake.Responder[armapimanagement.NetworkStatusClientListByLocationResponse], errResp azfake.ErrorResponder)

	// ListByService is the fake for method NetworkStatusClient.ListByService
	// HTTP status codes to indicate success: http.StatusOK
	ListByService func(ctx context.Context, resourceGroupName string, serviceName string, options *armapimanagement.NetworkStatusClientListByServiceOptions) (resp azfake.Responder[armapimanagement.NetworkStatusClientListByServiceResponse], errResp azfake.ErrorResponder)
}

// NewNetworkStatusServerTransport creates a new instance of NetworkStatusServerTransport with the provided implementation.
// The returned NetworkStatusServerTransport instance is connected to an instance of armapimanagement.NetworkStatusClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNetworkStatusServerTransport(srv *NetworkStatusServer) *NetworkStatusServerTransport {
	return &NetworkStatusServerTransport{srv: srv}
}

// NetworkStatusServerTransport connects instances of armapimanagement.NetworkStatusClient to instances of NetworkStatusServer.
// Don't use this type directly, use NewNetworkStatusServerTransport instead.
type NetworkStatusServerTransport struct {
	srv *NetworkStatusServer
}

// Do implements the policy.Transporter interface for NetworkStatusServerTransport.
func (n *NetworkStatusServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "NetworkStatusClient.ListByLocation":
		resp, err = n.dispatchListByLocation(req)
	case "NetworkStatusClient.ListByService":
		resp, err = n.dispatchListByService(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NetworkStatusServerTransport) dispatchListByLocation(req *http.Request) (*http.Response, error) {
	if n.srv.ListByLocation == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListByLocation not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/locations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkstatus`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.ListByLocation(req.Context(), resourceGroupNameParam, serviceNameParam, locationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NetworkStatusContract, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NetworkStatusServerTransport) dispatchListByService(req *http.Request) (*http.Response, error) {
	if n.srv.ListByService == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListByService not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkstatus`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.ListByService(req.Context(), resourceGroupNameParam, serviceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NetworkStatusContractByLocationArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

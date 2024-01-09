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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
	"net/http"
	"net/url"
	"regexp"
)

// ServiceAssociationLinksServer is a fake server for instances of the armnetwork.ServiceAssociationLinksClient type.
type ServiceAssociationLinksServer struct {
	// List is the fake for method ServiceAssociationLinksClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *armnetwork.ServiceAssociationLinksClientListOptions) (resp azfake.Responder[armnetwork.ServiceAssociationLinksClientListResponse], errResp azfake.ErrorResponder)
}

// NewServiceAssociationLinksServerTransport creates a new instance of ServiceAssociationLinksServerTransport with the provided implementation.
// The returned ServiceAssociationLinksServerTransport instance is connected to an instance of armnetwork.ServiceAssociationLinksClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServiceAssociationLinksServerTransport(srv *ServiceAssociationLinksServer) *ServiceAssociationLinksServerTransport {
	return &ServiceAssociationLinksServerTransport{srv: srv}
}

// ServiceAssociationLinksServerTransport connects instances of armnetwork.ServiceAssociationLinksClient to instances of ServiceAssociationLinksServer.
// Don't use this type directly, use NewServiceAssociationLinksServerTransport instead.
type ServiceAssociationLinksServerTransport struct {
	srv *ServiceAssociationLinksServer
}

// Do implements the policy.Transporter interface for ServiceAssociationLinksServerTransport.
func (s *ServiceAssociationLinksServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ServiceAssociationLinksClient.List":
		resp, err = s.dispatchList(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServiceAssociationLinksServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if s.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/virtualNetworks/(?P<virtualNetworkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subnets/(?P<subnetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ServiceAssociationLinks`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	virtualNetworkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkName")])
	if err != nil {
		return nil, err
	}
	subnetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("subnetName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.List(req.Context(), resourceGroupNameParam, virtualNetworkNameParam, subnetNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServiceAssociationLinksListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

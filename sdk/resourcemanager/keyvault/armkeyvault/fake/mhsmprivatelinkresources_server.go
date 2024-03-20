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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault/v2"
	"net/http"
	"net/url"
	"regexp"
)

// MHSMPrivateLinkResourcesServer is a fake server for instances of the armkeyvault.MHSMPrivateLinkResourcesClient type.
type MHSMPrivateLinkResourcesServer struct {
	// ListByMHSMResource is the fake for method MHSMPrivateLinkResourcesClient.ListByMHSMResource
	// HTTP status codes to indicate success: http.StatusOK
	ListByMHSMResource func(ctx context.Context, resourceGroupName string, name string, options *armkeyvault.MHSMPrivateLinkResourcesClientListByMHSMResourceOptions) (resp azfake.Responder[armkeyvault.MHSMPrivateLinkResourcesClientListByMHSMResourceResponse], errResp azfake.ErrorResponder)
}

// NewMHSMPrivateLinkResourcesServerTransport creates a new instance of MHSMPrivateLinkResourcesServerTransport with the provided implementation.
// The returned MHSMPrivateLinkResourcesServerTransport instance is connected to an instance of armkeyvault.MHSMPrivateLinkResourcesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewMHSMPrivateLinkResourcesServerTransport(srv *MHSMPrivateLinkResourcesServer) *MHSMPrivateLinkResourcesServerTransport {
	return &MHSMPrivateLinkResourcesServerTransport{srv: srv}
}

// MHSMPrivateLinkResourcesServerTransport connects instances of armkeyvault.MHSMPrivateLinkResourcesClient to instances of MHSMPrivateLinkResourcesServer.
// Don't use this type directly, use NewMHSMPrivateLinkResourcesServerTransport instead.
type MHSMPrivateLinkResourcesServerTransport struct {
	srv *MHSMPrivateLinkResourcesServer
}

// Do implements the policy.Transporter interface for MHSMPrivateLinkResourcesServerTransport.
func (m *MHSMPrivateLinkResourcesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "MHSMPrivateLinkResourcesClient.ListByMHSMResource":
		resp, err = m.dispatchListByMHSMResource(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *MHSMPrivateLinkResourcesServerTransport) dispatchListByMHSMResource(req *http.Request) (*http.Response, error) {
	if m.srv.ListByMHSMResource == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListByMHSMResource not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/managedHSMs/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.ListByMHSMResource(req.Context(), resourceGroupNameParam, nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).MHSMPrivateLinkResourceListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

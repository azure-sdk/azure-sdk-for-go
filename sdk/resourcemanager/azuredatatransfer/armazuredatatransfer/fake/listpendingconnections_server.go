//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredatatransfer/armazuredatatransfer"
	"net/http"
	"net/url"
	"regexp"
)

// ListPendingConnectionsServer is a fake server for instances of the armazuredatatransfer.ListPendingConnectionsClient type.
type ListPendingConnectionsServer struct {
	// NewListPager is the fake for method ListPendingConnectionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, connectionName string, options *armazuredatatransfer.ListPendingConnectionsClientListOptions) (resp azfake.PagerResponder[armazuredatatransfer.ListPendingConnectionsClientListResponse])
}

// NewListPendingConnectionsServerTransport creates a new instance of ListPendingConnectionsServerTransport with the provided implementation.
// The returned ListPendingConnectionsServerTransport instance is connected to an instance of armazuredatatransfer.ListPendingConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewListPendingConnectionsServerTransport(srv *ListPendingConnectionsServer) *ListPendingConnectionsServerTransport {
	return &ListPendingConnectionsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armazuredatatransfer.ListPendingConnectionsClientListResponse]](),
	}
}

// ListPendingConnectionsServerTransport connects instances of armazuredatatransfer.ListPendingConnectionsClient to instances of ListPendingConnectionsServer.
// Don't use this type directly, use NewListPendingConnectionsServerTransport instead.
type ListPendingConnectionsServerTransport struct {
	srv          *ListPendingConnectionsServer
	newListPager *tracker[azfake.PagerResponder[armazuredatatransfer.ListPendingConnectionsClientListResponse]]
}

// Do implements the policy.Transporter interface for ListPendingConnectionsServerTransport.
func (l *ListPendingConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ListPendingConnectionsClient.NewListPager":
		resp, err = l.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *ListPendingConnectionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := l.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureDataTransfer/connections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listPendingConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		connectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
		if err != nil {
			return nil, err
		}
		resp := l.srv.NewListPager(resourceGroupNameParam, connectionNameParam, nil)
		newListPager = &resp
		l.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armazuredatatransfer.ListPendingConnectionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		l.newListPager.remove(req)
	}
	return resp, nil
}

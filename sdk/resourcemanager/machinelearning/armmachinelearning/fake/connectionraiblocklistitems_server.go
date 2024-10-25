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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
	"net/http"
	"net/url"
	"regexp"
)

// ConnectionRaiBlocklistItemsServer is a fake server for instances of the armmachinelearning.ConnectionRaiBlocklistItemsClient type.
type ConnectionRaiBlocklistItemsServer struct {
	// NewListPager is the fake for method ConnectionRaiBlocklistItemsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, connectionName string, raiBlocklistName string, options *armmachinelearning.ConnectionRaiBlocklistItemsClientListOptions) (resp azfake.PagerResponder[armmachinelearning.ConnectionRaiBlocklistItemsClientListResponse])
}

// NewConnectionRaiBlocklistItemsServerTransport creates a new instance of ConnectionRaiBlocklistItemsServerTransport with the provided implementation.
// The returned ConnectionRaiBlocklistItemsServerTransport instance is connected to an instance of armmachinelearning.ConnectionRaiBlocklistItemsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewConnectionRaiBlocklistItemsServerTransport(srv *ConnectionRaiBlocklistItemsServer) *ConnectionRaiBlocklistItemsServerTransport {
	return &ConnectionRaiBlocklistItemsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armmachinelearning.ConnectionRaiBlocklistItemsClientListResponse]](),
	}
}

// ConnectionRaiBlocklistItemsServerTransport connects instances of armmachinelearning.ConnectionRaiBlocklistItemsClient to instances of ConnectionRaiBlocklistItemsServer.
// Don't use this type directly, use NewConnectionRaiBlocklistItemsServerTransport instead.
type ConnectionRaiBlocklistItemsServerTransport struct {
	srv          *ConnectionRaiBlocklistItemsServer
	newListPager *tracker[azfake.PagerResponder[armmachinelearning.ConnectionRaiBlocklistItemsClientListResponse]]
}

// Do implements the policy.Transporter interface for ConnectionRaiBlocklistItemsServerTransport.
func (c *ConnectionRaiBlocklistItemsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ConnectionRaiBlocklistItemsClient.NewListPager":
		resp, err = c.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ConnectionRaiBlocklistItemsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := c.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/raiBlocklists/(?P<raiBlocklistName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/raiBlocklistItems`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		connectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
		if err != nil {
			return nil, err
		}
		raiBlocklistNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("raiBlocklistName")])
		if err != nil {
			return nil, err
		}
		proxyAPIVersionUnescaped, err := url.QueryUnescape(qp.Get("proxy-api-version"))
		if err != nil {
			return nil, err
		}
		proxyAPIVersionParam := getOptional(proxyAPIVersionUnescaped)
		var options *armmachinelearning.ConnectionRaiBlocklistItemsClientListOptions
		if proxyAPIVersionParam != nil {
			options = &armmachinelearning.ConnectionRaiBlocklistItemsClientListOptions{
				ProxyAPIVersion: proxyAPIVersionParam,
			}
		}
		resp := c.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, connectionNameParam, raiBlocklistNameParam, options)
		newListPager = &resp
		c.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmachinelearning.ConnectionRaiBlocklistItemsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		c.newListPager.remove(req)
	}
	return resp, nil
}

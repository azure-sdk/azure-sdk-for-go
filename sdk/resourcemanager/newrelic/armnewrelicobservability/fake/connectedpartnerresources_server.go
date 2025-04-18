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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/newrelic/armnewrelicobservability"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// ConnectedPartnerResourcesServer is a fake server for instances of the armnewrelicobservability.ConnectedPartnerResourcesClient type.
type ConnectedPartnerResourcesServer struct {
	// NewListPager is the fake for method ConnectedPartnerResourcesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, monitorName string, options *armnewrelicobservability.ConnectedPartnerResourcesClientListOptions) (resp azfake.PagerResponder[armnewrelicobservability.ConnectedPartnerResourcesClientListResponse])
}

// NewConnectedPartnerResourcesServerTransport creates a new instance of ConnectedPartnerResourcesServerTransport with the provided implementation.
// The returned ConnectedPartnerResourcesServerTransport instance is connected to an instance of armnewrelicobservability.ConnectedPartnerResourcesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewConnectedPartnerResourcesServerTransport(srv *ConnectedPartnerResourcesServer) *ConnectedPartnerResourcesServerTransport {
	return &ConnectedPartnerResourcesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armnewrelicobservability.ConnectedPartnerResourcesClientListResponse]](),
	}
}

// ConnectedPartnerResourcesServerTransport connects instances of armnewrelicobservability.ConnectedPartnerResourcesClient to instances of ConnectedPartnerResourcesServer.
// Don't use this type directly, use NewConnectedPartnerResourcesServerTransport instead.
type ConnectedPartnerResourcesServerTransport struct {
	srv          *ConnectedPartnerResourcesServer
	newListPager *tracker[azfake.PagerResponder[armnewrelicobservability.ConnectedPartnerResourcesClientListResponse]]
}

// Do implements the policy.Transporter interface for ConnectedPartnerResourcesServerTransport.
func (c *ConnectedPartnerResourcesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *ConnectedPartnerResourcesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if connectedPartnerResourcesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = connectedPartnerResourcesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ConnectedPartnerResourcesClient.NewListPager":
				res.resp, res.err = c.dispatchNewListPager(req)
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

func (c *ConnectedPartnerResourcesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := c.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/NewRelic\.Observability/monitors/(?P<monitorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listConnectedPartnerResources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[string](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		monitorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("monitorName")])
		if err != nil {
			return nil, err
		}
		var options *armnewrelicobservability.ConnectedPartnerResourcesClientListOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armnewrelicobservability.ConnectedPartnerResourcesClientListOptions{
				Body: &body,
			}
		}
		resp := c.srv.NewListPager(resourceGroupNameParam, monitorNameParam, options)
		newListPager = &resp
		c.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnewrelicobservability.ConnectedPartnerResourcesClientListResponse, createLink func() string) {
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

// set this to conditionally intercept incoming requests to ConnectedPartnerResourcesServerTransport
var connectedPartnerResourcesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// SubscriptionNetworkManagerConnectionsServer is a fake server for instances of the armnetwork.SubscriptionNetworkManagerConnectionsClient type.
type SubscriptionNetworkManagerConnectionsServer struct {
	// CreateOrUpdate is the fake for method SubscriptionNetworkManagerConnectionsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, networkManagerConnectionName string, parameters armnetwork.ManagerConnection, options *armnetwork.SubscriptionNetworkManagerConnectionsClientCreateOrUpdateOptions) (resp azfake.Responder[armnetwork.SubscriptionNetworkManagerConnectionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method SubscriptionNetworkManagerConnectionsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, networkManagerConnectionName string, options *armnetwork.SubscriptionNetworkManagerConnectionsClientDeleteOptions) (resp azfake.Responder[armnetwork.SubscriptionNetworkManagerConnectionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SubscriptionNetworkManagerConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, networkManagerConnectionName string, options *armnetwork.SubscriptionNetworkManagerConnectionsClientGetOptions) (resp azfake.Responder[armnetwork.SubscriptionNetworkManagerConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method SubscriptionNetworkManagerConnectionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armnetwork.SubscriptionNetworkManagerConnectionsClientListOptions) (resp azfake.PagerResponder[armnetwork.SubscriptionNetworkManagerConnectionsClientListResponse])
}

// NewSubscriptionNetworkManagerConnectionsServerTransport creates a new instance of SubscriptionNetworkManagerConnectionsServerTransport with the provided implementation.
// The returned SubscriptionNetworkManagerConnectionsServerTransport instance is connected to an instance of armnetwork.SubscriptionNetworkManagerConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSubscriptionNetworkManagerConnectionsServerTransport(srv *SubscriptionNetworkManagerConnectionsServer) *SubscriptionNetworkManagerConnectionsServerTransport {
	return &SubscriptionNetworkManagerConnectionsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armnetwork.SubscriptionNetworkManagerConnectionsClientListResponse]](),
	}
}

// SubscriptionNetworkManagerConnectionsServerTransport connects instances of armnetwork.SubscriptionNetworkManagerConnectionsClient to instances of SubscriptionNetworkManagerConnectionsServer.
// Don't use this type directly, use NewSubscriptionNetworkManagerConnectionsServerTransport instead.
type SubscriptionNetworkManagerConnectionsServerTransport struct {
	srv          *SubscriptionNetworkManagerConnectionsServer
	newListPager *tracker[azfake.PagerResponder[armnetwork.SubscriptionNetworkManagerConnectionsClientListResponse]]
}

// Do implements the policy.Transporter interface for SubscriptionNetworkManagerConnectionsServerTransport.
func (s *SubscriptionNetworkManagerConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SubscriptionNetworkManagerConnectionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if subscriptionNetworkManagerConnectionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = subscriptionNetworkManagerConnectionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "SubscriptionNetworkManagerConnectionsClient.CreateOrUpdate":
				res.resp, res.err = s.dispatchCreateOrUpdate(req)
			case "SubscriptionNetworkManagerConnectionsClient.Delete":
				res.resp, res.err = s.dispatchDelete(req)
			case "SubscriptionNetworkManagerConnectionsClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "SubscriptionNetworkManagerConnectionsClient.NewListPager":
				res.resp, res.err = s.dispatchNewListPager(req)
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

func (s *SubscriptionNetworkManagerConnectionsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagerConnections/(?P<networkManagerConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.ManagerConnection](req)
	if err != nil {
		return nil, err
	}
	networkManagerConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.CreateOrUpdate(req.Context(), networkManagerConnectionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagerConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SubscriptionNetworkManagerConnectionsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if s.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagerConnections/(?P<networkManagerConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	networkManagerConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Delete(req.Context(), networkManagerConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SubscriptionNetworkManagerConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagerConnections/(?P<networkManagerConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	networkManagerConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), networkManagerConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagerConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SubscriptionNetworkManagerConnectionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagerConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		var options *armnetwork.SubscriptionNetworkManagerConnectionsClientListOptions
		if topParam != nil || skipTokenParam != nil {
			options = &armnetwork.SubscriptionNetworkManagerConnectionsClientListOptions{
				Top:       topParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := s.srv.NewListPager(options)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.SubscriptionNetworkManagerConnectionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		s.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to SubscriptionNetworkManagerConnectionsServerTransport
var subscriptionNetworkManagerConnectionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

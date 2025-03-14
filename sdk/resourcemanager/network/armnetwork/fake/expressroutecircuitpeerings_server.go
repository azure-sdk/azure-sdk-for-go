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
)

// ExpressRouteCircuitPeeringsServer is a fake server for instances of the armnetwork.ExpressRouteCircuitPeeringsClient type.
type ExpressRouteCircuitPeeringsServer struct {
	// BeginCreateOrUpdate is the fake for method ExpressRouteCircuitPeeringsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, peeringParameters armnetwork.ExpressRouteCircuitPeering, options *armnetwork.ExpressRouteCircuitPeeringsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.ExpressRouteCircuitPeeringsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ExpressRouteCircuitPeeringsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *armnetwork.ExpressRouteCircuitPeeringsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.ExpressRouteCircuitPeeringsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ExpressRouteCircuitPeeringsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *armnetwork.ExpressRouteCircuitPeeringsClientGetOptions) (resp azfake.Responder[armnetwork.ExpressRouteCircuitPeeringsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ExpressRouteCircuitPeeringsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, circuitName string, options *armnetwork.ExpressRouteCircuitPeeringsClientListOptions) (resp azfake.PagerResponder[armnetwork.ExpressRouteCircuitPeeringsClientListResponse])
}

// NewExpressRouteCircuitPeeringsServerTransport creates a new instance of ExpressRouteCircuitPeeringsServerTransport with the provided implementation.
// The returned ExpressRouteCircuitPeeringsServerTransport instance is connected to an instance of armnetwork.ExpressRouteCircuitPeeringsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExpressRouteCircuitPeeringsServerTransport(srv *ExpressRouteCircuitPeeringsServer) *ExpressRouteCircuitPeeringsServerTransport {
	return &ExpressRouteCircuitPeeringsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnetwork.ExpressRouteCircuitPeeringsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armnetwork.ExpressRouteCircuitPeeringsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armnetwork.ExpressRouteCircuitPeeringsClientListResponse]](),
	}
}

// ExpressRouteCircuitPeeringsServerTransport connects instances of armnetwork.ExpressRouteCircuitPeeringsClient to instances of ExpressRouteCircuitPeeringsServer.
// Don't use this type directly, use NewExpressRouteCircuitPeeringsServerTransport instead.
type ExpressRouteCircuitPeeringsServerTransport struct {
	srv                 *ExpressRouteCircuitPeeringsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnetwork.ExpressRouteCircuitPeeringsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armnetwork.ExpressRouteCircuitPeeringsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armnetwork.ExpressRouteCircuitPeeringsClientListResponse]]
}

// Do implements the policy.Transporter interface for ExpressRouteCircuitPeeringsServerTransport.
func (e *ExpressRouteCircuitPeeringsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return e.dispatchToMethodFake(req, method)
}

func (e *ExpressRouteCircuitPeeringsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if expressRouteCircuitPeeringsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = expressRouteCircuitPeeringsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ExpressRouteCircuitPeeringsClient.BeginCreateOrUpdate":
				res.resp, res.err = e.dispatchBeginCreateOrUpdate(req)
			case "ExpressRouteCircuitPeeringsClient.BeginDelete":
				res.resp, res.err = e.dispatchBeginDelete(req)
			case "ExpressRouteCircuitPeeringsClient.Get":
				res.resp, res.err = e.dispatchGet(req)
			case "ExpressRouteCircuitPeeringsClient.NewListPager":
				res.resp, res.err = e.dispatchNewListPager(req)
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

func (e *ExpressRouteCircuitPeeringsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := e.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/expressRouteCircuits/(?P<circuitName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerings/(?P<peeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.ExpressRouteCircuitPeering](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		circuitNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("circuitName")])
		if err != nil {
			return nil, err
		}
		peeringNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("peeringName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, circuitNameParam, peeringNameParam, body, nil)
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

func (e *ExpressRouteCircuitPeeringsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if e.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := e.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/expressRouteCircuits/(?P<circuitName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerings/(?P<peeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		circuitNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("circuitName")])
		if err != nil {
			return nil, err
		}
		peeringNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("peeringName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginDelete(req.Context(), resourceGroupNameParam, circuitNameParam, peeringNameParam, nil)
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

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		e.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		e.beginDelete.remove(req)
	}

	return resp, nil
}

func (e *ExpressRouteCircuitPeeringsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/expressRouteCircuits/(?P<circuitName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerings/(?P<peeringName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	circuitNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("circuitName")])
	if err != nil {
		return nil, err
	}
	peeringNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("peeringName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameParam, circuitNameParam, peeringNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExpressRouteCircuitPeering, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExpressRouteCircuitPeeringsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := e.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/expressRouteCircuits/(?P<circuitName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/peerings`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		circuitNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("circuitName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListPager(resourceGroupNameParam, circuitNameParam, nil)
		newListPager = &resp
		e.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.ExpressRouteCircuitPeeringsClientListResponse, createLink func() string) {
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

// set this to conditionally intercept incoming requests to ExpressRouteCircuitPeeringsServerTransport
var expressRouteCircuitPeeringsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

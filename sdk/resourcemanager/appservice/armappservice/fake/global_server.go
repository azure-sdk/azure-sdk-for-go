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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
	"net/http"
	"net/url"
	"regexp"
)

// GlobalServer is a fake server for instances of the armappservice.GlobalClient type.
type GlobalServer struct {
	// GetDeletedWebApp is the fake for method GlobalClient.GetDeletedWebApp
	// HTTP status codes to indicate success: http.StatusOK
	GetDeletedWebApp func(ctx context.Context, deletedSiteID string, options *armappservice.GlobalClientGetDeletedWebAppOptions) (resp azfake.Responder[armappservice.GlobalClientGetDeletedWebAppResponse], errResp azfake.ErrorResponder)

	// GetDeletedWebAppSnapshots is the fake for method GlobalClient.GetDeletedWebAppSnapshots
	// HTTP status codes to indicate success: http.StatusOK
	GetDeletedWebAppSnapshots func(ctx context.Context, deletedSiteID string, options *armappservice.GlobalClientGetDeletedWebAppSnapshotsOptions) (resp azfake.Responder[armappservice.GlobalClientGetDeletedWebAppSnapshotsResponse], errResp azfake.ErrorResponder)

	// GetSubscriptionOperationWithAsyncResponse is the fake for method GlobalClient.GetSubscriptionOperationWithAsyncResponse
	// HTTP status codes to indicate success: http.StatusNoContent
	GetSubscriptionOperationWithAsyncResponse func(ctx context.Context, location string, operationID string, options *armappservice.GlobalClientGetSubscriptionOperationWithAsyncResponseOptions) (resp azfake.Responder[armappservice.GlobalClientGetSubscriptionOperationWithAsyncResponseResponse], errResp azfake.ErrorResponder)
}

// NewGlobalServerTransport creates a new instance of GlobalServerTransport with the provided implementation.
// The returned GlobalServerTransport instance is connected to an instance of armappservice.GlobalClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGlobalServerTransport(srv *GlobalServer) *GlobalServerTransport {
	return &GlobalServerTransport{srv: srv}
}

// GlobalServerTransport connects instances of armappservice.GlobalClient to instances of GlobalServer.
// Don't use this type directly, use NewGlobalServerTransport instead.
type GlobalServerTransport struct {
	srv *GlobalServer
}

// Do implements the policy.Transporter interface for GlobalServerTransport.
func (g *GlobalServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return g.dispatchToMethodFake(req, method)
}

func (g *GlobalServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if globalServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = globalServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "GlobalClient.GetDeletedWebApp":
				res.resp, res.err = g.dispatchGetDeletedWebApp(req)
			case "GlobalClient.GetDeletedWebAppSnapshots":
				res.resp, res.err = g.dispatchGetDeletedWebAppSnapshots(req)
			case "GlobalClient.GetSubscriptionOperationWithAsyncResponse":
				res.resp, res.err = g.dispatchGetSubscriptionOperationWithAsyncResponse(req)
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

func (g *GlobalServerTransport) dispatchGetDeletedWebApp(req *http.Request) (*http.Response, error) {
	if g.srv.GetDeletedWebApp == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDeletedWebApp not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/deletedSites/(?P<deletedSiteId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	deletedSiteIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("deletedSiteId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.GetDeletedWebApp(req.Context(), deletedSiteIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeletedSite, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GlobalServerTransport) dispatchGetDeletedWebAppSnapshots(req *http.Request) (*http.Response, error) {
	if g.srv.GetDeletedWebAppSnapshots == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDeletedWebAppSnapshots not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/deletedSites/(?P<deletedSiteId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/snapshots`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	deletedSiteIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("deletedSiteId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.GetDeletedWebAppSnapshots(req.Context(), deletedSiteIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SnapshotArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GlobalServerTransport) dispatchGetSubscriptionOperationWithAsyncResponse(req *http.Request) (*http.Response, error) {
	if g.srv.GetSubscriptionOperationWithAsyncResponse == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetSubscriptionOperationWithAsyncResponse not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/operations/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.GetSubscriptionOperationWithAsyncResponse(req.Context(), locationParam, operationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to GlobalServerTransport
var globalServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

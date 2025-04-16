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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"
	"net/http"
	"net/url"
	"regexp"
)

// LocationsServer is a fake server for instances of the armdatalakestore.LocationsClient type.
type LocationsServer struct {
	// GetCapability is the fake for method LocationsClient.GetCapability
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNotFound
	GetCapability func(ctx context.Context, location string, options *armdatalakestore.LocationsClientGetCapabilityOptions) (resp azfake.Responder[armdatalakestore.LocationsClientGetCapabilityResponse], errResp azfake.ErrorResponder)

	// NewGetUsagePager is the fake for method LocationsClient.NewGetUsagePager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetUsagePager func(location string, options *armdatalakestore.LocationsClientGetUsageOptions) (resp azfake.PagerResponder[armdatalakestore.LocationsClientGetUsageResponse])
}

// NewLocationsServerTransport creates a new instance of LocationsServerTransport with the provided implementation.
// The returned LocationsServerTransport instance is connected to an instance of armdatalakestore.LocationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLocationsServerTransport(srv *LocationsServer) *LocationsServerTransport {
	return &LocationsServerTransport{
		srv:              srv,
		newGetUsagePager: newTracker[azfake.PagerResponder[armdatalakestore.LocationsClientGetUsageResponse]](),
	}
}

// LocationsServerTransport connects instances of armdatalakestore.LocationsClient to instances of LocationsServer.
// Don't use this type directly, use NewLocationsServerTransport instead.
type LocationsServerTransport struct {
	srv              *LocationsServer
	newGetUsagePager *tracker[azfake.PagerResponder[armdatalakestore.LocationsClientGetUsageResponse]]
}

// Do implements the policy.Transporter interface for LocationsServerTransport.
func (l *LocationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return l.dispatchToMethodFake(req, method)
}

func (l *LocationsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if locationsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = locationsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "LocationsClient.GetCapability":
				res.resp, res.err = l.dispatchGetCapability(req)
			case "LocationsClient.NewGetUsagePager":
				res.resp, res.err = l.dispatchNewGetUsagePager(req)
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

func (l *LocationsServerTransport) dispatchGetCapability(req *http.Request) (*http.Response, error) {
	if l.srv.GetCapability == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetCapability not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataLakeStore/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/capability`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.GetCapability(req.Context(), locationParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNotFound}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNotFound", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CapabilityInformation, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LocationsServerTransport) dispatchNewGetUsagePager(req *http.Request) (*http.Response, error) {
	if l.srv.NewGetUsagePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetUsagePager not implemented")}
	}
	newGetUsagePager := l.newGetUsagePager.get(req)
	if newGetUsagePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataLakeStore/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/usages`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resp := l.srv.NewGetUsagePager(locationParam, nil)
		newGetUsagePager = &resp
		l.newGetUsagePager.add(req, newGetUsagePager)
	}
	resp, err := server.PagerResponderNext(newGetUsagePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newGetUsagePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newGetUsagePager) {
		l.newGetUsagePager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to LocationsServerTransport
var locationsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

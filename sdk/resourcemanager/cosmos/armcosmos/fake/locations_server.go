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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
	"net/http"
	"net/url"
	"regexp"
)

// LocationsServer is a fake server for instances of the armcosmos.LocationsClient type.
type LocationsServer struct {
	// Get is the fake for method LocationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, options *armcosmos.LocationsClientGetOptions) (resp azfake.Responder[armcosmos.LocationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method LocationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armcosmos.LocationsClientListOptions) (resp azfake.PagerResponder[armcosmos.LocationsClientListResponse])
}

// NewLocationsServerTransport creates a new instance of LocationsServerTransport with the provided implementation.
// The returned LocationsServerTransport instance is connected to an instance of armcosmos.LocationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLocationsServerTransport(srv *LocationsServer) *LocationsServerTransport {
	return &LocationsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armcosmos.LocationsClientListResponse]](),
	}
}

// LocationsServerTransport connects instances of armcosmos.LocationsClient to instances of LocationsServer.
// Don't use this type directly, use NewLocationsServerTransport instead.
type LocationsServerTransport struct {
	srv          *LocationsServer
	newListPager *tracker[azfake.PagerResponder[armcosmos.LocationsClientListResponse]]
}

// Do implements the policy.Transporter interface for LocationsServerTransport.
func (l *LocationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LocationsClient.Get":
		resp, err = l.dispatchGet(req)
	case "LocationsClient.NewListPager":
		resp, err = l.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LocationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if l.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.Get(req.Context(), locationParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LocationGetResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LocationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := l.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/locations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := l.srv.NewListPager(nil)
		newListPager = &resp
		l.newListPager.add(req, newListPager)
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

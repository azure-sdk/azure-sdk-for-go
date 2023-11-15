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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
	"net/http"
	"net/url"
	"regexp"
)

// DeletedServersServer is a fake server for instances of the armsql.DeletedServersClient type.
type DeletedServersServer struct {
	// Get is the fake for method DeletedServersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, locationName string, deletedServerName string, options *armsql.DeletedServersClientGetOptions) (resp azfake.Responder[armsql.DeletedServersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method DeletedServersClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armsql.DeletedServersClientListOptions) (resp azfake.PagerResponder[armsql.DeletedServersClientListResponse])

	// NewListByLocationPager is the fake for method DeletedServersClient.NewListByLocationPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByLocationPager func(locationName string, options *armsql.DeletedServersClientListByLocationOptions) (resp azfake.PagerResponder[armsql.DeletedServersClientListByLocationResponse])

	// BeginRecover is the fake for method DeletedServersClient.BeginRecover
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRecover func(ctx context.Context, locationName string, deletedServerName string, options *armsql.DeletedServersClientBeginRecoverOptions) (resp azfake.PollerResponder[armsql.DeletedServersClientRecoverResponse], errResp azfake.ErrorResponder)
}

// NewDeletedServersServerTransport creates a new instance of DeletedServersServerTransport with the provided implementation.
// The returned DeletedServersServerTransport instance is connected to an instance of armsql.DeletedServersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDeletedServersServerTransport(srv *DeletedServersServer) *DeletedServersServerTransport {
	return &DeletedServersServerTransport{
		srv:                    srv,
		newListPager:           newTracker[azfake.PagerResponder[armsql.DeletedServersClientListResponse]](),
		newListByLocationPager: newTracker[azfake.PagerResponder[armsql.DeletedServersClientListByLocationResponse]](),
		beginRecover:           newTracker[azfake.PollerResponder[armsql.DeletedServersClientRecoverResponse]](),
	}
}

// DeletedServersServerTransport connects instances of armsql.DeletedServersClient to instances of DeletedServersServer.
// Don't use this type directly, use NewDeletedServersServerTransport instead.
type DeletedServersServerTransport struct {
	srv                    *DeletedServersServer
	newListPager           *tracker[azfake.PagerResponder[armsql.DeletedServersClientListResponse]]
	newListByLocationPager *tracker[azfake.PagerResponder[armsql.DeletedServersClientListByLocationResponse]]
	beginRecover           *tracker[azfake.PollerResponder[armsql.DeletedServersClientRecoverResponse]]
}

// Do implements the policy.Transporter interface for DeletedServersServerTransport.
func (d *DeletedServersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DeletedServersClient.Get":
		resp, err = d.dispatchGet(req)
	case "DeletedServersClient.NewListPager":
		resp, err = d.dispatchNewListPager(req)
	case "DeletedServersClient.NewListByLocationPager":
		resp, err = d.dispatchNewListByLocationPager(req)
	case "DeletedServersClient.BeginRecover":
		resp, err = d.dispatchBeginRecover(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DeletedServersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/locations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deletedServers/(?P<deletedServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
	if err != nil {
		return nil, err
	}
	deletedServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deletedServerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), locationNameParam, deletedServerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeletedServer, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DeletedServersServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := d.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/deletedServers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := d.srv.NewListPager(nil)
		newListPager = &resp
		d.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsql.DeletedServersClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		d.newListPager.remove(req)
	}
	return resp, nil
}

func (d *DeletedServersServerTransport) dispatchNewListByLocationPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByLocationPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByLocationPager not implemented")}
	}
	newListByLocationPager := d.newListByLocationPager.get(req)
	if newListByLocationPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/locations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deletedServers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListByLocationPager(locationNameParam, nil)
		newListByLocationPager = &resp
		d.newListByLocationPager.add(req, newListByLocationPager)
		server.PagerResponderInjectNextLinks(newListByLocationPager, req, func(page *armsql.DeletedServersClientListByLocationResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByLocationPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListByLocationPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByLocationPager) {
		d.newListByLocationPager.remove(req)
	}
	return resp, nil
}

func (d *DeletedServersServerTransport) dispatchBeginRecover(req *http.Request) (*http.Response, error) {
	if d.srv.BeginRecover == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRecover not implemented")}
	}
	beginRecover := d.beginRecover.get(req)
	if beginRecover == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/locations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deletedServers/(?P<deletedServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/recover`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
		if err != nil {
			return nil, err
		}
		deletedServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deletedServerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginRecover(req.Context(), locationNameParam, deletedServerNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRecover = &respr
		d.beginRecover.add(req, beginRecover)
	}

	resp, err := server.PollerResponderNext(beginRecover, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginRecover.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRecover) {
		d.beginRecover.remove(req)
	}

	return resp, nil
}

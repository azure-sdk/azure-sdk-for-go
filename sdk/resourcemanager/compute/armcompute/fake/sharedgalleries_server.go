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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	"net/http"
	"net/url"
	"regexp"
)

// SharedGalleriesServer is a fake server for instances of the armcompute.SharedGalleriesClient type.
type SharedGalleriesServer struct {
	// Get is the fake for method SharedGalleriesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, galleryUniqueName string, options *armcompute.SharedGalleriesClientGetOptions) (resp azfake.Responder[armcompute.SharedGalleriesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method SharedGalleriesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(location string, options *armcompute.SharedGalleriesClientListOptions) (resp azfake.PagerResponder[armcompute.SharedGalleriesClientListResponse])
}

// NewSharedGalleriesServerTransport creates a new instance of SharedGalleriesServerTransport with the provided implementation.
// The returned SharedGalleriesServerTransport instance is connected to an instance of armcompute.SharedGalleriesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSharedGalleriesServerTransport(srv *SharedGalleriesServer) *SharedGalleriesServerTransport {
	return &SharedGalleriesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armcompute.SharedGalleriesClientListResponse]](),
	}
}

// SharedGalleriesServerTransport connects instances of armcompute.SharedGalleriesClient to instances of SharedGalleriesServer.
// Don't use this type directly, use NewSharedGalleriesServerTransport instead.
type SharedGalleriesServerTransport struct {
	srv          *SharedGalleriesServer
	newListPager *tracker[azfake.PagerResponder[armcompute.SharedGalleriesClientListResponse]]
}

// Do implements the policy.Transporter interface for SharedGalleriesServerTransport.
func (s *SharedGalleriesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SharedGalleriesClient.Get":
		resp, err = s.dispatchGet(req)
	case "SharedGalleriesClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SharedGalleriesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sharedGalleries/(?P<galleryUniqueName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	galleryUniqueNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("galleryUniqueName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), locationUnescaped, galleryUniqueNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SharedGallery, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SharedGalleriesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sharedGalleries`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		sharedToUnescaped, err := url.QueryUnescape(qp.Get("sharedTo"))
		if err != nil {
			return nil, err
		}
		sharedToParam := getOptional(armcompute.SharedToValues(sharedToUnescaped))
		var options *armcompute.SharedGalleriesClientListOptions
		if sharedToParam != nil {
			options = &armcompute.SharedGalleriesClientListOptions{
				SharedTo: sharedToParam,
			}
		}
		resp := s.srv.NewListPager(locationUnescaped, options)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcompute.SharedGalleriesClientListResponse, createLink func() string) {
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

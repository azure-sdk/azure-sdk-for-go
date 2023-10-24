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

// CommunityGalleryImagesServer is a fake server for instances of the armcompute.CommunityGalleryImagesClient type.
type CommunityGalleryImagesServer struct {
	// Get is the fake for method CommunityGalleryImagesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, publicGalleryName string, galleryImageName string, options *armcompute.CommunityGalleryImagesClientGetOptions) (resp azfake.Responder[armcompute.CommunityGalleryImagesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method CommunityGalleryImagesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(location string, publicGalleryName string, options *armcompute.CommunityGalleryImagesClientListOptions) (resp azfake.PagerResponder[armcompute.CommunityGalleryImagesClientListResponse])
}

// NewCommunityGalleryImagesServerTransport creates a new instance of CommunityGalleryImagesServerTransport with the provided implementation.
// The returned CommunityGalleryImagesServerTransport instance is connected to an instance of armcompute.CommunityGalleryImagesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCommunityGalleryImagesServerTransport(srv *CommunityGalleryImagesServer) *CommunityGalleryImagesServerTransport {
	return &CommunityGalleryImagesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armcompute.CommunityGalleryImagesClientListResponse]](),
	}
}

// CommunityGalleryImagesServerTransport connects instances of armcompute.CommunityGalleryImagesClient to instances of CommunityGalleryImagesServer.
// Don't use this type directly, use NewCommunityGalleryImagesServerTransport instead.
type CommunityGalleryImagesServerTransport struct {
	srv          *CommunityGalleryImagesServer
	newListPager *tracker[azfake.PagerResponder[armcompute.CommunityGalleryImagesClientListResponse]]
}

// Do implements the policy.Transporter interface for CommunityGalleryImagesServerTransport.
func (c *CommunityGalleryImagesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CommunityGalleryImagesClient.Get":
		resp, err = c.dispatchGet(req)
	case "CommunityGalleryImagesClient.NewListPager":
		resp, err = c.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CommunityGalleryImagesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/communityGalleries/(?P<publicGalleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/images/(?P<galleryImageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	publicGalleryNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("publicGalleryName")])
	if err != nil {
		return nil, err
	}
	galleryImageNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("galleryImageName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), locationUnescaped, publicGalleryNameUnescaped, galleryImageNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CommunityGalleryImage, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CommunityGalleryImagesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := c.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/communityGalleries/(?P<publicGalleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/images`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		publicGalleryNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("publicGalleryName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListPager(locationUnescaped, publicGalleryNameUnescaped, nil)
		newListPager = &resp
		c.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcompute.CommunityGalleryImagesClientListResponse, createLink func() string) {
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

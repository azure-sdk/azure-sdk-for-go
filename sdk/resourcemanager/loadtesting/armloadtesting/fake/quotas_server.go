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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtesting/armloadtesting/v2"
	"net/http"
	"net/url"
	"regexp"
)

// QuotasServer is a fake server for instances of the armloadtesting.QuotasClient type.
type QuotasServer struct {
	// CheckAvailability is the fake for method QuotasClient.CheckAvailability
	// HTTP status codes to indicate success: http.StatusOK
	CheckAvailability func(ctx context.Context, location string, quotaBucketName string, body armloadtesting.QuotaBucketRequest, options *armloadtesting.QuotasClientCheckAvailabilityOptions) (resp azfake.Responder[armloadtesting.QuotasClientCheckAvailabilityResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method QuotasClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, quotaBucketName string, options *armloadtesting.QuotasClientGetOptions) (resp azfake.Responder[armloadtesting.QuotasClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method QuotasClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(location string, options *armloadtesting.QuotasClientListOptions) (resp azfake.PagerResponder[armloadtesting.QuotasClientListResponse])
}

// NewQuotasServerTransport creates a new instance of QuotasServerTransport with the provided implementation.
// The returned QuotasServerTransport instance is connected to an instance of armloadtesting.QuotasClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewQuotasServerTransport(srv *QuotasServer) *QuotasServerTransport {
	return &QuotasServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armloadtesting.QuotasClientListResponse]](),
	}
}

// QuotasServerTransport connects instances of armloadtesting.QuotasClient to instances of QuotasServer.
// Don't use this type directly, use NewQuotasServerTransport instead.
type QuotasServerTransport struct {
	srv          *QuotasServer
	newListPager *tracker[azfake.PagerResponder[armloadtesting.QuotasClientListResponse]]
}

// Do implements the policy.Transporter interface for QuotasServerTransport.
func (q *QuotasServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "QuotasClient.CheckAvailability":
		resp, err = q.dispatchCheckAvailability(req)
	case "QuotasClient.Get":
		resp, err = q.dispatchGet(req)
	case "QuotasClient.NewListPager":
		resp, err = q.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (q *QuotasServerTransport) dispatchCheckAvailability(req *http.Request) (*http.Response, error) {
	if q.srv.CheckAvailability == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckAvailability not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LoadTestService/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/quotas/(?P<quotaBucketName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/checkAvailability`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armloadtesting.QuotaBucketRequest](req)
	if err != nil {
		return nil, err
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	quotaBucketNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("quotaBucketName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := q.srv.CheckAvailability(req.Context(), locationParam, quotaBucketNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CheckQuotaAvailabilityResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *QuotasServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if q.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LoadTestService/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/quotas/(?P<quotaBucketName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	quotaBucketNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("quotaBucketName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := q.srv.Get(req.Context(), locationParam, quotaBucketNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).QuotaResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *QuotasServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if q.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := q.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LoadTestService/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/quotas`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resp := q.srv.NewListPager(locationParam, nil)
		newListPager = &resp
		q.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armloadtesting.QuotasClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		q.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		q.newListPager.remove(req)
	}
	return resp, nil
}

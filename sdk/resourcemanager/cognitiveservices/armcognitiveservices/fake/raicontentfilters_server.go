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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
	"net/http"
	"net/url"
	"regexp"
)

// RaiContentFiltersServer is a fake server for instances of the armcognitiveservices.RaiContentFiltersClient type.
type RaiContentFiltersServer struct {
	// Get is the fake for method RaiContentFiltersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, filterName string, options *armcognitiveservices.RaiContentFiltersClientGetOptions) (resp azfake.Responder[armcognitiveservices.RaiContentFiltersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method RaiContentFiltersClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(location string, options *armcognitiveservices.RaiContentFiltersClientListOptions) (resp azfake.PagerResponder[armcognitiveservices.RaiContentFiltersClientListResponse])
}

// NewRaiContentFiltersServerTransport creates a new instance of RaiContentFiltersServerTransport with the provided implementation.
// The returned RaiContentFiltersServerTransport instance is connected to an instance of armcognitiveservices.RaiContentFiltersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRaiContentFiltersServerTransport(srv *RaiContentFiltersServer) *RaiContentFiltersServerTransport {
	return &RaiContentFiltersServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armcognitiveservices.RaiContentFiltersClientListResponse]](),
	}
}

// RaiContentFiltersServerTransport connects instances of armcognitiveservices.RaiContentFiltersClient to instances of RaiContentFiltersServer.
// Don't use this type directly, use NewRaiContentFiltersServerTransport instead.
type RaiContentFiltersServerTransport struct {
	srv          *RaiContentFiltersServer
	newListPager *tracker[azfake.PagerResponder[armcognitiveservices.RaiContentFiltersClientListResponse]]
}

// Do implements the policy.Transporter interface for RaiContentFiltersServerTransport.
func (r *RaiContentFiltersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *RaiContentFiltersServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if raiContentFiltersServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = raiContentFiltersServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "RaiContentFiltersClient.Get":
				res.resp, res.err = r.dispatchGet(req)
			case "RaiContentFiltersClient.NewListPager":
				res.resp, res.err = r.dispatchNewListPager(req)
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

func (r *RaiContentFiltersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CognitiveServices/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/raiContentFilters/(?P<filterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	filterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("filterName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), locationParam, filterNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RaiContentFilter, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RaiContentFiltersServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CognitiveServices/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/raiContentFilters`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListPager(locationParam, nil)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcognitiveservices.RaiContentFiltersClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to RaiContentFiltersServerTransport
var raiContentFiltersServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

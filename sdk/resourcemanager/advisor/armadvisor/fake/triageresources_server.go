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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
	"net/http"
	"net/url"
	"regexp"
)

// TriageResourcesServer is a fake server for instances of the armadvisor.TriageResourcesClient type.
type TriageResourcesServer struct {
	// Get is the fake for method TriageResourcesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, reviewID string, recommendationID string, recommendationResourceID string, options *armadvisor.TriageResourcesClientGetOptions) (resp azfake.Responder[armadvisor.TriageResourcesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method TriageResourcesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(reviewID string, recommendationID string, options *armadvisor.TriageResourcesClientListOptions) (resp azfake.PagerResponder[armadvisor.TriageResourcesClientListResponse])
}

// NewTriageResourcesServerTransport creates a new instance of TriageResourcesServerTransport with the provided implementation.
// The returned TriageResourcesServerTransport instance is connected to an instance of armadvisor.TriageResourcesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTriageResourcesServerTransport(srv *TriageResourcesServer) *TriageResourcesServerTransport {
	return &TriageResourcesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armadvisor.TriageResourcesClientListResponse]](),
	}
}

// TriageResourcesServerTransport connects instances of armadvisor.TriageResourcesClient to instances of TriageResourcesServer.
// Don't use this type directly, use NewTriageResourcesServerTransport instead.
type TriageResourcesServerTransport struct {
	srv          *TriageResourcesServer
	newListPager *tracker[azfake.PagerResponder[armadvisor.TriageResourcesClientListResponse]]
}

// Do implements the policy.Transporter interface for TriageResourcesServerTransport.
func (t *TriageResourcesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return t.dispatchToMethodFake(req, method)
}

func (t *TriageResourcesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if triageResourcesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = triageResourcesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "TriageResourcesClient.Get":
				res.resp, res.err = t.dispatchGet(req)
			case "TriageResourcesClient.NewListPager":
				res.resp, res.err = t.dispatchNewListPager(req)
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

func (t *TriageResourcesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if t.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Advisor/resiliencyReviews/(?P<reviewId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Advisor/triageRecommendations/(?P<recommendationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Advisor/triageResources/(?P<recommendationResourceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	reviewIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("reviewId")])
	if err != nil {
		return nil, err
	}
	recommendationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("recommendationId")])
	if err != nil {
		return nil, err
	}
	recommendationResourceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("recommendationResourceId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Get(req.Context(), reviewIDParam, recommendationIDParam, recommendationResourceIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TriageResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TriageResourcesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := t.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Advisor/resiliencyReviews/(?P<reviewId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Advisor/triageRecommendations/(?P<recommendationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Advisor/triageResources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		reviewIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("reviewId")])
		if err != nil {
			return nil, err
		}
		recommendationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("recommendationId")])
		if err != nil {
			return nil, err
		}
		resp := t.srv.NewListPager(reviewIDParam, recommendationIDParam, nil)
		newListPager = &resp
		t.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armadvisor.TriageResourcesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		t.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to TriageResourcesServerTransport
var triageResourcesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

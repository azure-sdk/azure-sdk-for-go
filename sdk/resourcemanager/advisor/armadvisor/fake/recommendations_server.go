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
	"strconv"
)

// RecommendationsServer is a fake server for instances of the armadvisor.RecommendationsClient type.
type RecommendationsServer struct {
	// Generate is the fake for method RecommendationsClient.Generate
	// HTTP status codes to indicate success: http.StatusAccepted
	Generate func(ctx context.Context, options *armadvisor.RecommendationsClientGenerateOptions) (resp azfake.Responder[armadvisor.RecommendationsClientGenerateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RecommendationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceURI string, recommendationID string, options *armadvisor.RecommendationsClientGetOptions) (resp azfake.Responder[armadvisor.RecommendationsClientGetResponse], errResp azfake.ErrorResponder)

	// GetGenerateStatus is the fake for method RecommendationsClient.GetGenerateStatus
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	GetGenerateStatus func(ctx context.Context, operationID string, options *armadvisor.RecommendationsClientGetGenerateStatusOptions) (resp azfake.Responder[armadvisor.RecommendationsClientGetGenerateStatusResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method RecommendationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armadvisor.RecommendationsClientListOptions) (resp azfake.PagerResponder[armadvisor.RecommendationsClientListResponse])

	// Patch is the fake for method RecommendationsClient.Patch
	// HTTP status codes to indicate success: http.StatusOK
	Patch func(ctx context.Context, resourceURI string, recommendationID string, trackedProperties armadvisor.TrackedRecommendationPropertiesPayload, options *armadvisor.RecommendationsClientPatchOptions) (resp azfake.Responder[armadvisor.RecommendationsClientPatchResponse], errResp azfake.ErrorResponder)
}

// NewRecommendationsServerTransport creates a new instance of RecommendationsServerTransport with the provided implementation.
// The returned RecommendationsServerTransport instance is connected to an instance of armadvisor.RecommendationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRecommendationsServerTransport(srv *RecommendationsServer) *RecommendationsServerTransport {
	return &RecommendationsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armadvisor.RecommendationsClientListResponse]](),
	}
}

// RecommendationsServerTransport connects instances of armadvisor.RecommendationsClient to instances of RecommendationsServer.
// Don't use this type directly, use NewRecommendationsServerTransport instead.
type RecommendationsServerTransport struct {
	srv          *RecommendationsServer
	newListPager *tracker[azfake.PagerResponder[armadvisor.RecommendationsClientListResponse]]
}

// Do implements the policy.Transporter interface for RecommendationsServerTransport.
func (r *RecommendationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *RecommendationsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if recommendationsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = recommendationsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "RecommendationsClient.Generate":
				res.resp, res.err = r.dispatchGenerate(req)
			case "RecommendationsClient.Get":
				res.resp, res.err = r.dispatchGet(req)
			case "RecommendationsClient.GetGenerateStatus":
				res.resp, res.err = r.dispatchGetGenerateStatus(req)
			case "RecommendationsClient.NewListPager":
				res.resp, res.err = r.dispatchNewListPager(req)
			case "RecommendationsClient.Patch":
				res.resp, res.err = r.dispatchPatch(req)
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

func (r *RecommendationsServerTransport) dispatchGenerate(req *http.Request) (*http.Response, error) {
	if r.srv.Generate == nil {
		return nil, &nonRetriableError{errors.New("fake for method Generate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Advisor/generateRecommendations`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := r.srv.Generate(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).Location; val != nil {
		resp.Header.Set("Location", *val)
	}
	if val := server.GetResponse(respr).RetryAfter; val != nil {
		resp.Header.Set("Retry-After", *val)
	}
	return resp, nil
}

func (r *RecommendationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Advisor/recommendations/(?P<recommendationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
	if err != nil {
		return nil, err
	}
	recommendationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("recommendationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceURIParam, recommendationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ResourceRecommendationBase, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RecommendationsServerTransport) dispatchGetGenerateStatus(req *http.Request) (*http.Response, error) {
	if r.srv.GetGenerateStatus == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetGenerateStatus not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Advisor/generateRecommendations/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.GetGenerateStatus(req.Context(), operationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RecommendationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Advisor/recommendations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		var options *armadvisor.RecommendationsClientListOptions
		if filterParam != nil || topParam != nil || skipTokenParam != nil {
			options = &armadvisor.RecommendationsClientListOptions{
				Filter:    filterParam,
				Top:       topParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := r.srv.NewListPager(options)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armadvisor.RecommendationsClientListResponse, createLink func() string) {
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

func (r *RecommendationsServerTransport) dispatchPatch(req *http.Request) (*http.Response, error) {
	if r.srv.Patch == nil {
		return nil, &nonRetriableError{errors.New("fake for method Patch not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Advisor/recommendations/(?P<recommendationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armadvisor.TrackedRecommendationPropertiesPayload](req)
	if err != nil {
		return nil, err
	}
	resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
	if err != nil {
		return nil, err
	}
	recommendationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("recommendationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Patch(req.Context(), resourceURIParam, recommendationIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ResourceRecommendationBase, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to RecommendationsServerTransport
var recommendationsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

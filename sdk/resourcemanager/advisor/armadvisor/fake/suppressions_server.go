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

// SuppressionsServer is a fake server for instances of the armadvisor.SuppressionsClient type.
type SuppressionsServer struct {
	// Create is the fake for method SuppressionsClient.Create
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, resourceURI string, recommendationID string, name string, suppressionContract armadvisor.SuppressionContract, options *armadvisor.SuppressionsClientCreateOptions) (resp azfake.Responder[armadvisor.SuppressionsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method SuppressionsClient.Delete
	// HTTP status codes to indicate success: http.StatusNoContent
	Delete func(ctx context.Context, resourceURI string, recommendationID string, name string, options *armadvisor.SuppressionsClientDeleteOptions) (resp azfake.Responder[armadvisor.SuppressionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SuppressionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceURI string, recommendationID string, name string, options *armadvisor.SuppressionsClientGetOptions) (resp azfake.Responder[armadvisor.SuppressionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method SuppressionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armadvisor.SuppressionsClientListOptions) (resp azfake.PagerResponder[armadvisor.SuppressionsClientListResponse])
}

// NewSuppressionsServerTransport creates a new instance of SuppressionsServerTransport with the provided implementation.
// The returned SuppressionsServerTransport instance is connected to an instance of armadvisor.SuppressionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSuppressionsServerTransport(srv *SuppressionsServer) *SuppressionsServerTransport {
	return &SuppressionsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armadvisor.SuppressionsClientListResponse]](),
	}
}

// SuppressionsServerTransport connects instances of armadvisor.SuppressionsClient to instances of SuppressionsServer.
// Don't use this type directly, use NewSuppressionsServerTransport instead.
type SuppressionsServerTransport struct {
	srv          *SuppressionsServer
	newListPager *tracker[azfake.PagerResponder[armadvisor.SuppressionsClientListResponse]]
}

// Do implements the policy.Transporter interface for SuppressionsServerTransport.
func (s *SuppressionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SuppressionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if suppressionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = suppressionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "SuppressionsClient.Create":
				res.resp, res.err = s.dispatchCreate(req)
			case "SuppressionsClient.Delete":
				res.resp, res.err = s.dispatchDelete(req)
			case "SuppressionsClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "SuppressionsClient.NewListPager":
				res.resp, res.err = s.dispatchNewListPager(req)
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

func (s *SuppressionsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if s.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Advisor/recommendations/(?P<recommendationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/suppressions/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armadvisor.SuppressionContract](req)
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
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Create(req.Context(), resourceURIParam, recommendationIDParam, nameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SuppressionContract, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SuppressionsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if s.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Advisor/recommendations/(?P<recommendationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/suppressions/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
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
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Delete(req.Context(), resourceURIParam, recommendationIDParam, nameParam, nil)
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

func (s *SuppressionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Advisor/recommendations/(?P<recommendationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/suppressions/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
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
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceURIParam, recommendationIDParam, nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SuppressionContract, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SuppressionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Advisor/suppressions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
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
		var options *armadvisor.SuppressionsClientListOptions
		if topParam != nil || skipTokenParam != nil {
			options = &armadvisor.SuppressionsClientListOptions{
				Top:       topParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := s.srv.NewListPager(options)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armadvisor.SuppressionsClientListResponse, createLink func() string) {
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

// set this to conditionally intercept incoming requests to SuppressionsServerTransport
var suppressionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

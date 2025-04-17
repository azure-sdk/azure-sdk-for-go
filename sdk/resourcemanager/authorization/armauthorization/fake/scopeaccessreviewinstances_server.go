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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ScopeAccessReviewInstancesServer is a fake server for instances of the armauthorization.ScopeAccessReviewInstancesClient type.
type ScopeAccessReviewInstancesServer struct {
	// Create is the fake for method ScopeAccessReviewInstancesClient.Create
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, scope string, scheduleDefinitionID string, id string, properties armauthorization.AccessReviewInstanceProperties, options *armauthorization.ScopeAccessReviewInstancesClientCreateOptions) (resp azfake.Responder[armauthorization.ScopeAccessReviewInstancesClientCreateResponse], errResp azfake.ErrorResponder)

	// GetByID is the fake for method ScopeAccessReviewInstancesClient.GetByID
	// HTTP status codes to indicate success: http.StatusOK
	GetByID func(ctx context.Context, scope string, scheduleDefinitionID string, id string, options *armauthorization.ScopeAccessReviewInstancesClientGetByIDOptions) (resp azfake.Responder[armauthorization.ScopeAccessReviewInstancesClientGetByIDResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ScopeAccessReviewInstancesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(scope string, scheduleDefinitionID string, options *armauthorization.ScopeAccessReviewInstancesClientListOptions) (resp azfake.PagerResponder[armauthorization.ScopeAccessReviewInstancesClientListResponse])
}

// NewScopeAccessReviewInstancesServerTransport creates a new instance of ScopeAccessReviewInstancesServerTransport with the provided implementation.
// The returned ScopeAccessReviewInstancesServerTransport instance is connected to an instance of armauthorization.ScopeAccessReviewInstancesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewScopeAccessReviewInstancesServerTransport(srv *ScopeAccessReviewInstancesServer) *ScopeAccessReviewInstancesServerTransport {
	return &ScopeAccessReviewInstancesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armauthorization.ScopeAccessReviewInstancesClientListResponse]](),
	}
}

// ScopeAccessReviewInstancesServerTransport connects instances of armauthorization.ScopeAccessReviewInstancesClient to instances of ScopeAccessReviewInstancesServer.
// Don't use this type directly, use NewScopeAccessReviewInstancesServerTransport instead.
type ScopeAccessReviewInstancesServerTransport struct {
	srv          *ScopeAccessReviewInstancesServer
	newListPager *tracker[azfake.PagerResponder[armauthorization.ScopeAccessReviewInstancesClientListResponse]]
}

// Do implements the policy.Transporter interface for ScopeAccessReviewInstancesServerTransport.
func (s *ScopeAccessReviewInstancesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ScopeAccessReviewInstancesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if scopeAccessReviewInstancesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = scopeAccessReviewInstancesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ScopeAccessReviewInstancesClient.Create":
				res.resp, res.err = s.dispatchCreate(req)
			case "ScopeAccessReviewInstancesClient.GetByID":
				res.resp, res.err = s.dispatchGetByID(req)
			case "ScopeAccessReviewInstancesClient.NewListPager":
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

func (s *ScopeAccessReviewInstancesServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if s.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewScheduleDefinitions/(?P<scheduleDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/instances/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armauthorization.AccessReviewInstanceProperties](req)
	if err != nil {
		return nil, err
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	scheduleDefinitionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("scheduleDefinitionId")])
	if err != nil {
		return nil, err
	}
	idParam, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Create(req.Context(), scopeParam, scheduleDefinitionIDParam, idParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccessReviewInstance, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScopeAccessReviewInstancesServerTransport) dispatchGetByID(req *http.Request) (*http.Response, error) {
	if s.srv.GetByID == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByID not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewScheduleDefinitions/(?P<scheduleDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/instances/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	scheduleDefinitionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("scheduleDefinitionId")])
	if err != nil {
		return nil, err
	}
	idParam, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetByID(req.Context(), scopeParam, scheduleDefinitionIDParam, idParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccessReviewInstance, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScopeAccessReviewInstancesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewScheduleDefinitions/(?P<scheduleDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/instances`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
		if err != nil {
			return nil, err
		}
		scheduleDefinitionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("scheduleDefinitionId")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armauthorization.ScopeAccessReviewInstancesClientListOptions
		if filterParam != nil {
			options = &armauthorization.ScopeAccessReviewInstancesClientListOptions{
				Filter: filterParam,
			}
		}
		resp := s.srv.NewListPager(scopeParam, scheduleDefinitionIDParam, options)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armauthorization.ScopeAccessReviewInstancesClientListResponse, createLink func() string) {
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

// set this to conditionally intercept incoming requests to ScopeAccessReviewInstancesServerTransport
var scopeAccessReviewInstancesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

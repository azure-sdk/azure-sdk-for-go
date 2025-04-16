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

// ScopeAccessReviewHistoryDefinitionsServer is a fake server for instances of the armauthorization.ScopeAccessReviewHistoryDefinitionsClient type.
type ScopeAccessReviewHistoryDefinitionsServer struct {
	// GetByID is the fake for method ScopeAccessReviewHistoryDefinitionsClient.GetByID
	// HTTP status codes to indicate success: http.StatusOK
	GetByID func(ctx context.Context, scope string, historyDefinitionID string, options *armauthorization.ScopeAccessReviewHistoryDefinitionsClientGetByIDOptions) (resp azfake.Responder[armauthorization.ScopeAccessReviewHistoryDefinitionsClientGetByIDResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ScopeAccessReviewHistoryDefinitionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(scope string, options *armauthorization.ScopeAccessReviewHistoryDefinitionsClientListOptions) (resp azfake.PagerResponder[armauthorization.ScopeAccessReviewHistoryDefinitionsClientListResponse])
}

// NewScopeAccessReviewHistoryDefinitionsServerTransport creates a new instance of ScopeAccessReviewHistoryDefinitionsServerTransport with the provided implementation.
// The returned ScopeAccessReviewHistoryDefinitionsServerTransport instance is connected to an instance of armauthorization.ScopeAccessReviewHistoryDefinitionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewScopeAccessReviewHistoryDefinitionsServerTransport(srv *ScopeAccessReviewHistoryDefinitionsServer) *ScopeAccessReviewHistoryDefinitionsServerTransport {
	return &ScopeAccessReviewHistoryDefinitionsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armauthorization.ScopeAccessReviewHistoryDefinitionsClientListResponse]](),
	}
}

// ScopeAccessReviewHistoryDefinitionsServerTransport connects instances of armauthorization.ScopeAccessReviewHistoryDefinitionsClient to instances of ScopeAccessReviewHistoryDefinitionsServer.
// Don't use this type directly, use NewScopeAccessReviewHistoryDefinitionsServerTransport instead.
type ScopeAccessReviewHistoryDefinitionsServerTransport struct {
	srv          *ScopeAccessReviewHistoryDefinitionsServer
	newListPager *tracker[azfake.PagerResponder[armauthorization.ScopeAccessReviewHistoryDefinitionsClientListResponse]]
}

// Do implements the policy.Transporter interface for ScopeAccessReviewHistoryDefinitionsServerTransport.
func (s *ScopeAccessReviewHistoryDefinitionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ScopeAccessReviewHistoryDefinitionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if scopeAccessReviewHistoryDefinitionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = scopeAccessReviewHistoryDefinitionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ScopeAccessReviewHistoryDefinitionsClient.GetByID":
				res.resp, res.err = s.dispatchGetByID(req)
			case "ScopeAccessReviewHistoryDefinitionsClient.NewListPager":
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

func (s *ScopeAccessReviewHistoryDefinitionsServerTransport) dispatchGetByID(req *http.Request) (*http.Response, error) {
	if s.srv.GetByID == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByID not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewHistoryDefinitions/(?P<historyDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	historyDefinitionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("historyDefinitionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetByID(req.Context(), scopeParam, historyDefinitionIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccessReviewHistoryDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScopeAccessReviewHistoryDefinitionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewHistoryDefinitions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armauthorization.ScopeAccessReviewHistoryDefinitionsClientListOptions
		if filterParam != nil {
			options = &armauthorization.ScopeAccessReviewHistoryDefinitionsClientListOptions{
				Filter: filterParam,
			}
		}
		resp := s.srv.NewListPager(scopeParam, options)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armauthorization.ScopeAccessReviewHistoryDefinitionsClientListResponse, createLink func() string) {
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

// set this to conditionally intercept incoming requests to ScopeAccessReviewHistoryDefinitionsServerTransport
var scopeAccessReviewHistoryDefinitionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

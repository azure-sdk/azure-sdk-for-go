// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
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

// ScopeAccessReviewInstanceContactedReviewersServer is a fake server for instances of the armauthorization.ScopeAccessReviewInstanceContactedReviewersClient type.
type ScopeAccessReviewInstanceContactedReviewersServer struct {
	// NewListPager is the fake for method ScopeAccessReviewInstanceContactedReviewersClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(scope string, scheduleDefinitionID string, id string, options *armauthorization.ScopeAccessReviewInstanceContactedReviewersClientListOptions) (resp azfake.PagerResponder[armauthorization.ScopeAccessReviewInstanceContactedReviewersClientListResponse])
}

// NewScopeAccessReviewInstanceContactedReviewersServerTransport creates a new instance of ScopeAccessReviewInstanceContactedReviewersServerTransport with the provided implementation.
// The returned ScopeAccessReviewInstanceContactedReviewersServerTransport instance is connected to an instance of armauthorization.ScopeAccessReviewInstanceContactedReviewersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewScopeAccessReviewInstanceContactedReviewersServerTransport(srv *ScopeAccessReviewInstanceContactedReviewersServer) *ScopeAccessReviewInstanceContactedReviewersServerTransport {
	return &ScopeAccessReviewInstanceContactedReviewersServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armauthorization.ScopeAccessReviewInstanceContactedReviewersClientListResponse]](),
	}
}

// ScopeAccessReviewInstanceContactedReviewersServerTransport connects instances of armauthorization.ScopeAccessReviewInstanceContactedReviewersClient to instances of ScopeAccessReviewInstanceContactedReviewersServer.
// Don't use this type directly, use NewScopeAccessReviewInstanceContactedReviewersServerTransport instead.
type ScopeAccessReviewInstanceContactedReviewersServerTransport struct {
	srv          *ScopeAccessReviewInstanceContactedReviewersServer
	newListPager *tracker[azfake.PagerResponder[armauthorization.ScopeAccessReviewInstanceContactedReviewersClientListResponse]]
}

// Do implements the policy.Transporter interface for ScopeAccessReviewInstanceContactedReviewersServerTransport.
func (s *ScopeAccessReviewInstanceContactedReviewersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ScopeAccessReviewInstanceContactedReviewersServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if scopeAccessReviewInstanceContactedReviewersServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = scopeAccessReviewInstanceContactedReviewersServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ScopeAccessReviewInstanceContactedReviewersClient.NewListPager":
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

func (s *ScopeAccessReviewInstanceContactedReviewersServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewScheduleDefinitions/(?P<scheduleDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/instances/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/contactedReviewers`
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
		resp := s.srv.NewListPager(scopeParam, scheduleDefinitionIDParam, idParam, nil)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armauthorization.ScopeAccessReviewInstanceContactedReviewersClientListResponse, createLink func() string) {
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

// set this to conditionally intercept incoming requests to ScopeAccessReviewInstanceContactedReviewersServerTransport
var scopeAccessReviewInstanceContactedReviewersServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

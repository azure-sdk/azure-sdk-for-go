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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ScopeAccessReviewDefaultSettingsServer is a fake server for instances of the armauthorization.ScopeAccessReviewDefaultSettingsClient type.
type ScopeAccessReviewDefaultSettingsServer struct {
	// Get is the fake for method ScopeAccessReviewDefaultSettingsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, scope string, options *armauthorization.ScopeAccessReviewDefaultSettingsClientGetOptions) (resp azfake.Responder[armauthorization.ScopeAccessReviewDefaultSettingsClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method ScopeAccessReviewDefaultSettingsClient.Put
	// HTTP status codes to indicate success: http.StatusOK
	Put func(ctx context.Context, scope string, properties armauthorization.AccessReviewScheduleSettings, options *armauthorization.ScopeAccessReviewDefaultSettingsClientPutOptions) (resp azfake.Responder[armauthorization.ScopeAccessReviewDefaultSettingsClientPutResponse], errResp azfake.ErrorResponder)
}

// NewScopeAccessReviewDefaultSettingsServerTransport creates a new instance of ScopeAccessReviewDefaultSettingsServerTransport with the provided implementation.
// The returned ScopeAccessReviewDefaultSettingsServerTransport instance is connected to an instance of armauthorization.ScopeAccessReviewDefaultSettingsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewScopeAccessReviewDefaultSettingsServerTransport(srv *ScopeAccessReviewDefaultSettingsServer) *ScopeAccessReviewDefaultSettingsServerTransport {
	return &ScopeAccessReviewDefaultSettingsServerTransport{srv: srv}
}

// ScopeAccessReviewDefaultSettingsServerTransport connects instances of armauthorization.ScopeAccessReviewDefaultSettingsClient to instances of ScopeAccessReviewDefaultSettingsServer.
// Don't use this type directly, use NewScopeAccessReviewDefaultSettingsServerTransport instead.
type ScopeAccessReviewDefaultSettingsServerTransport struct {
	srv *ScopeAccessReviewDefaultSettingsServer
}

// Do implements the policy.Transporter interface for ScopeAccessReviewDefaultSettingsServerTransport.
func (s *ScopeAccessReviewDefaultSettingsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ScopeAccessReviewDefaultSettingsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if scopeAccessReviewDefaultSettingsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = scopeAccessReviewDefaultSettingsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ScopeAccessReviewDefaultSettingsClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "ScopeAccessReviewDefaultSettingsClient.Put":
				res.resp, res.err = s.dispatchPut(req)
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

func (s *ScopeAccessReviewDefaultSettingsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewScheduleSettings/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), scopeParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccessReviewDefaultSettings, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScopeAccessReviewDefaultSettingsServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if s.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewScheduleSettings/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armauthorization.AccessReviewScheduleSettings](req)
	if err != nil {
		return nil, err
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Put(req.Context(), scopeParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccessReviewDefaultSettings, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ScopeAccessReviewDefaultSettingsServerTransport
var scopeAccessReviewDefaultSettingsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

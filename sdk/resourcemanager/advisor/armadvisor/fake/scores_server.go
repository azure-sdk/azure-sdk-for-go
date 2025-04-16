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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
	"net/http"
	"net/url"
	"regexp"
)

// ScoresServer is a fake server for instances of the armadvisor.ScoresClient type.
type ScoresServer struct {
	// Get is the fake for method ScoresClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, name string, options *armadvisor.ScoresClientGetOptions) (resp azfake.Responder[armadvisor.ScoresClientGetResponse], errResp azfake.ErrorResponder)

	// List is the fake for method ScoresClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, options *armadvisor.ScoresClientListOptions) (resp azfake.Responder[armadvisor.ScoresClientListResponse], errResp azfake.ErrorResponder)
}

// NewScoresServerTransport creates a new instance of ScoresServerTransport with the provided implementation.
// The returned ScoresServerTransport instance is connected to an instance of armadvisor.ScoresClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewScoresServerTransport(srv *ScoresServer) *ScoresServerTransport {
	return &ScoresServerTransport{srv: srv}
}

// ScoresServerTransport connects instances of armadvisor.ScoresClient to instances of ScoresServer.
// Don't use this type directly, use NewScoresServerTransport instead.
type ScoresServerTransport struct {
	srv *ScoresServer
}

// Do implements the policy.Transporter interface for ScoresServerTransport.
func (s *ScoresServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ScoresServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if scoresServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = scoresServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ScoresClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "ScoresClient.List":
				res.resp, res.err = s.dispatchList(req)
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

func (s *ScoresServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Advisor/advisorScore/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ScoreEntityForAdvisor, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScoresServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if s.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Advisor/advisorScore`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := s.srv.List(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ScoreResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ScoresServerTransport
var scoresServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

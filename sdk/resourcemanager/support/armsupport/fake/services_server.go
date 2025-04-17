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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
	"net/http"
	"net/url"
	"regexp"
)

// ServicesServer is a fake server for instances of the armsupport.ServicesClient type.
type ServicesServer struct {
	// Get is the fake for method ServicesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, serviceName string, options *armsupport.ServicesClientGetOptions) (resp azfake.Responder[armsupport.ServicesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ServicesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armsupport.ServicesClientListOptions) (resp azfake.PagerResponder[armsupport.ServicesClientListResponse])
}

// NewServicesServerTransport creates a new instance of ServicesServerTransport with the provided implementation.
// The returned ServicesServerTransport instance is connected to an instance of armsupport.ServicesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServicesServerTransport(srv *ServicesServer) *ServicesServerTransport {
	return &ServicesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsupport.ServicesClientListResponse]](),
	}
}

// ServicesServerTransport connects instances of armsupport.ServicesClient to instances of ServicesServer.
// Don't use this type directly, use NewServicesServerTransport instead.
type ServicesServerTransport struct {
	srv          *ServicesServer
	newListPager *tracker[azfake.PagerResponder[armsupport.ServicesClientListResponse]]
}

// Do implements the policy.Transporter interface for ServicesServerTransport.
func (s *ServicesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ServicesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if servicesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = servicesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ServicesClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "ServicesClient.NewListPager":
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

func (s *ServicesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Support/services/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), serviceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Service, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServicesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		resp := s.srv.NewListPager(nil)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
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

// set this to conditionally intercept incoming requests to ServicesServerTransport
var servicesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

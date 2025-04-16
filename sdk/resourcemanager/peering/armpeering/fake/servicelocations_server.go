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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
	"net/http"
	"net/url"
	"regexp"
)

// ServiceLocationsServer is a fake server for instances of the armpeering.ServiceLocationsClient type.
type ServiceLocationsServer struct {
	// NewListPager is the fake for method ServiceLocationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armpeering.ServiceLocationsClientListOptions) (resp azfake.PagerResponder[armpeering.ServiceLocationsClientListResponse])
}

// NewServiceLocationsServerTransport creates a new instance of ServiceLocationsServerTransport with the provided implementation.
// The returned ServiceLocationsServerTransport instance is connected to an instance of armpeering.ServiceLocationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServiceLocationsServerTransport(srv *ServiceLocationsServer) *ServiceLocationsServerTransport {
	return &ServiceLocationsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armpeering.ServiceLocationsClientListResponse]](),
	}
}

// ServiceLocationsServerTransport connects instances of armpeering.ServiceLocationsClient to instances of ServiceLocationsServer.
// Don't use this type directly, use NewServiceLocationsServerTransport instead.
type ServiceLocationsServerTransport struct {
	srv          *ServiceLocationsServer
	newListPager *tracker[azfake.PagerResponder[armpeering.ServiceLocationsClientListResponse]]
}

// Do implements the policy.Transporter interface for ServiceLocationsServerTransport.
func (s *ServiceLocationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ServiceLocationsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if serviceLocationsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = serviceLocationsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ServiceLocationsClient.NewListPager":
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

func (s *ServiceLocationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Peering/peeringServiceLocations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		countryUnescaped, err := url.QueryUnescape(qp.Get("country"))
		if err != nil {
			return nil, err
		}
		countryParam := getOptional(countryUnescaped)
		var options *armpeering.ServiceLocationsClientListOptions
		if countryParam != nil {
			options = &armpeering.ServiceLocationsClientListOptions{
				Country: countryParam,
			}
		}
		resp := s.srv.NewListPager(options)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armpeering.ServiceLocationsClientListResponse, createLink func() string) {
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

// set this to conditionally intercept incoming requests to ServiceLocationsServerTransport
var serviceLocationsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
	"net/http"
	"net/url"
	"regexp"
)

// LocationBasedPerformanceTierServer is a fake server for instances of the armmariadb.LocationBasedPerformanceTierClient type.
type LocationBasedPerformanceTierServer struct {
	// NewListPager is the fake for method LocationBasedPerformanceTierClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(locationName string, options *armmariadb.LocationBasedPerformanceTierClientListOptions) (resp azfake.PagerResponder[armmariadb.LocationBasedPerformanceTierClientListResponse])
}

// NewLocationBasedPerformanceTierServerTransport creates a new instance of LocationBasedPerformanceTierServerTransport with the provided implementation.
// The returned LocationBasedPerformanceTierServerTransport instance is connected to an instance of armmariadb.LocationBasedPerformanceTierClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLocationBasedPerformanceTierServerTransport(srv *LocationBasedPerformanceTierServer) *LocationBasedPerformanceTierServerTransport {
	return &LocationBasedPerformanceTierServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armmariadb.LocationBasedPerformanceTierClientListResponse]](),
	}
}

// LocationBasedPerformanceTierServerTransport connects instances of armmariadb.LocationBasedPerformanceTierClient to instances of LocationBasedPerformanceTierServer.
// Don't use this type directly, use NewLocationBasedPerformanceTierServerTransport instead.
type LocationBasedPerformanceTierServerTransport struct {
	srv          *LocationBasedPerformanceTierServer
	newListPager *tracker[azfake.PagerResponder[armmariadb.LocationBasedPerformanceTierClientListResponse]]
}

// Do implements the policy.Transporter interface for LocationBasedPerformanceTierServerTransport.
func (l *LocationBasedPerformanceTierServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return l.dispatchToMethodFake(req, method)
}

func (l *LocationBasedPerformanceTierServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if locationBasedPerformanceTierServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = locationBasedPerformanceTierServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "LocationBasedPerformanceTierClient.NewListPager":
				res.resp, res.err = l.dispatchNewListPager(req)
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

func (l *LocationBasedPerformanceTierServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := l.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforMariaDB/locations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/performanceTiers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
		if err != nil {
			return nil, err
		}
		resp := l.srv.NewListPager(locationNameParam, nil)
		newListPager = &resp
		l.newListPager.add(req, newListPager)
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		l.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to LocationBasedPerformanceTierServerTransport
var locationBasedPerformanceTierServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

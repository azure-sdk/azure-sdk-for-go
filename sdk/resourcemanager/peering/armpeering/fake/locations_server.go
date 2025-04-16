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

// LocationsServer is a fake server for instances of the armpeering.LocationsClient type.
type LocationsServer struct {
	// NewListPager is the fake for method LocationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(kind armpeering.PeeringLocationsKind, options *armpeering.LocationsClientListOptions) (resp azfake.PagerResponder[armpeering.LocationsClientListResponse])
}

// NewLocationsServerTransport creates a new instance of LocationsServerTransport with the provided implementation.
// The returned LocationsServerTransport instance is connected to an instance of armpeering.LocationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLocationsServerTransport(srv *LocationsServer) *LocationsServerTransport {
	return &LocationsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armpeering.LocationsClientListResponse]](),
	}
}

// LocationsServerTransport connects instances of armpeering.LocationsClient to instances of LocationsServer.
// Don't use this type directly, use NewLocationsServerTransport instead.
type LocationsServerTransport struct {
	srv          *LocationsServer
	newListPager *tracker[azfake.PagerResponder[armpeering.LocationsClientListResponse]]
}

// Do implements the policy.Transporter interface for LocationsServerTransport.
func (l *LocationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return l.dispatchToMethodFake(req, method)
}

func (l *LocationsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if locationsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = locationsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "LocationsClient.NewListPager":
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

func (l *LocationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := l.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Peering/peeringLocations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		kindParam, err := parseWithCast(qp.Get("kind"), func(v string) (armpeering.PeeringLocationsKind, error) {
			p, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armpeering.PeeringLocationsKind(p), nil
		})
		if err != nil {
			return nil, err
		}
		directPeeringTypeUnescaped, err := url.QueryUnescape(qp.Get("directPeeringType"))
		if err != nil {
			return nil, err
		}
		directPeeringTypeParam := getOptional(armpeering.PeeringLocationsDirectPeeringType(directPeeringTypeUnescaped))
		var options *armpeering.LocationsClientListOptions
		if directPeeringTypeParam != nil {
			options = &armpeering.LocationsClientListOptions{
				DirectPeeringType: directPeeringTypeParam,
			}
		}
		resp := l.srv.NewListPager(kindParam, options)
		newListPager = &resp
		l.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armpeering.LocationsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
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

// set this to conditionally intercept incoming requests to LocationsServerTransport
var locationsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

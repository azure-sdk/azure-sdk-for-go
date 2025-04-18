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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// MarketplacesServer is a fake server for instances of the armconsumption.MarketplacesClient type.
type MarketplacesServer struct {
	// NewListPager is the fake for method MarketplacesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	NewListPager func(scope string, options *armconsumption.MarketplacesClientListOptions) (resp azfake.PagerResponder[armconsumption.MarketplacesClientListResponse])
}

// NewMarketplacesServerTransport creates a new instance of MarketplacesServerTransport with the provided implementation.
// The returned MarketplacesServerTransport instance is connected to an instance of armconsumption.MarketplacesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewMarketplacesServerTransport(srv *MarketplacesServer) *MarketplacesServerTransport {
	return &MarketplacesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armconsumption.MarketplacesClientListResponse]](),
	}
}

// MarketplacesServerTransport connects instances of armconsumption.MarketplacesClient to instances of MarketplacesServer.
// Don't use this type directly, use NewMarketplacesServerTransport instead.
type MarketplacesServerTransport struct {
	srv          *MarketplacesServer
	newListPager *tracker[azfake.PagerResponder[armconsumption.MarketplacesClientListResponse]]
}

// Do implements the policy.Transporter interface for MarketplacesServerTransport.
func (m *MarketplacesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return m.dispatchToMethodFake(req, method)
}

func (m *MarketplacesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if marketplacesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = marketplacesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "MarketplacesClient.NewListPager":
				res.resp, res.err = m.dispatchNewListPager(req)
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

func (m *MarketplacesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := m.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Consumption/marketplaces`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
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
		skiptokenUnescaped, err := url.QueryUnescape(qp.Get("$skiptoken"))
		if err != nil {
			return nil, err
		}
		skiptokenParam := getOptional(skiptokenUnescaped)
		scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
		if err != nil {
			return nil, err
		}
		var options *armconsumption.MarketplacesClientListOptions
		if filterParam != nil || topParam != nil || skiptokenParam != nil {
			options = &armconsumption.MarketplacesClientListOptions{
				Filter:    filterParam,
				Top:       topParam,
				Skiptoken: skiptokenParam,
			}
		}
		resp := m.srv.NewListPager(scopeParam, options)
		newListPager = &resp
		m.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armconsumption.MarketplacesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK, http.StatusNoContent}, resp.StatusCode) {
		m.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		m.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to MarketplacesServerTransport
var marketplacesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

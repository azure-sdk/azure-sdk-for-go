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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
	"net/http"
	"net/url"
	"regexp"
)

// QuotaUsagesServer is a fake server for instances of the armpostgresqlflexibleservers.QuotaUsagesClient type.
type QuotaUsagesServer struct {
	// NewListPager is the fake for method QuotaUsagesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(locationName string, options *armpostgresqlflexibleservers.QuotaUsagesClientListOptions) (resp azfake.PagerResponder[armpostgresqlflexibleservers.QuotaUsagesClientListResponse])
}

// NewQuotaUsagesServerTransport creates a new instance of QuotaUsagesServerTransport with the provided implementation.
// The returned QuotaUsagesServerTransport instance is connected to an instance of armpostgresqlflexibleservers.QuotaUsagesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewQuotaUsagesServerTransport(srv *QuotaUsagesServer) *QuotaUsagesServerTransport {
	return &QuotaUsagesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armpostgresqlflexibleservers.QuotaUsagesClientListResponse]](),
	}
}

// QuotaUsagesServerTransport connects instances of armpostgresqlflexibleservers.QuotaUsagesClient to instances of QuotaUsagesServer.
// Don't use this type directly, use NewQuotaUsagesServerTransport instead.
type QuotaUsagesServerTransport struct {
	srv          *QuotaUsagesServer
	newListPager *tracker[azfake.PagerResponder[armpostgresqlflexibleservers.QuotaUsagesClientListResponse]]
}

// Do implements the policy.Transporter interface for QuotaUsagesServerTransport.
func (q *QuotaUsagesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return q.dispatchToMethodFake(req, method)
}

func (q *QuotaUsagesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if quotaUsagesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = quotaUsagesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "QuotaUsagesClient.NewListPager":
				res.resp, res.err = q.dispatchNewListPager(req)
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

func (q *QuotaUsagesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if q.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := q.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/locations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceType/flexibleServers/usages`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
		if err != nil {
			return nil, err
		}
		resp := q.srv.NewListPager(locationNameParam, nil)
		newListPager = &resp
		q.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armpostgresqlflexibleservers.QuotaUsagesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		q.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		q.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to QuotaUsagesServerTransport
var quotaUsagesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ProductPackagesServer is a fake server for instances of the armsecurityinsights.ProductPackagesClient type.
type ProductPackagesServer struct {
	// NewListPager is the fake for method ProductPackagesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, options *armsecurityinsights.ProductPackagesClientListOptions) (resp azfake.PagerResponder[armsecurityinsights.ProductPackagesClientListResponse])
}

// NewProductPackagesServerTransport creates a new instance of ProductPackagesServerTransport with the provided implementation.
// The returned ProductPackagesServerTransport instance is connected to an instance of armsecurityinsights.ProductPackagesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProductPackagesServerTransport(srv *ProductPackagesServer) *ProductPackagesServerTransport {
	return &ProductPackagesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsecurityinsights.ProductPackagesClientListResponse]](),
	}
}

// ProductPackagesServerTransport connects instances of armsecurityinsights.ProductPackagesClient to instances of ProductPackagesServer.
// Don't use this type directly, use NewProductPackagesServerTransport instead.
type ProductPackagesServerTransport struct {
	srv          *ProductPackagesServer
	newListPager *tracker[azfake.PagerResponder[armsecurityinsights.ProductPackagesClientListResponse]]
}

// Do implements the policy.Transporter interface for ProductPackagesServerTransport.
func (p *ProductPackagesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *ProductPackagesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if productPackagesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = productPackagesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ProductPackagesClient.NewListPager":
				res.resp, res.err = p.dispatchNewListPager(req)
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

func (p *ProductPackagesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/contentProductPackages`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		orderbyUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
		if err != nil {
			return nil, err
		}
		orderbyParam := getOptional(orderbyUnescaped)
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
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		searchUnescaped, err := url.QueryUnescape(qp.Get("$search"))
		if err != nil {
			return nil, err
		}
		searchParam := getOptional(searchUnescaped)
		var options *armsecurityinsights.ProductPackagesClientListOptions
		if filterParam != nil || orderbyParam != nil || topParam != nil || skipTokenParam != nil || searchParam != nil {
			options = &armsecurityinsights.ProductPackagesClientListOptions{
				Filter:    filterParam,
				Orderby:   orderbyParam,
				Top:       topParam,
				SkipToken: skipTokenParam,
				Search:    searchParam,
			}
		}
		resp := p.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, options)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsecurityinsights.ProductPackagesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		p.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ProductPackagesServerTransport
var productPackagesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

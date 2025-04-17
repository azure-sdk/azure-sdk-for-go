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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcemover/armresourcemover"
	"net/http"
	"net/url"
	"regexp"
)

// UnresolvedDependenciesServer is a fake server for instances of the armresourcemover.UnresolvedDependenciesClient type.
type UnresolvedDependenciesServer struct {
	// NewGetPager is the fake for method UnresolvedDependenciesClient.NewGetPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetPager func(resourceGroupName string, moveCollectionName string, options *armresourcemover.UnresolvedDependenciesClientGetOptions) (resp azfake.PagerResponder[armresourcemover.UnresolvedDependenciesClientGetResponse])
}

// NewUnresolvedDependenciesServerTransport creates a new instance of UnresolvedDependenciesServerTransport with the provided implementation.
// The returned UnresolvedDependenciesServerTransport instance is connected to an instance of armresourcemover.UnresolvedDependenciesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewUnresolvedDependenciesServerTransport(srv *UnresolvedDependenciesServer) *UnresolvedDependenciesServerTransport {
	return &UnresolvedDependenciesServerTransport{
		srv:         srv,
		newGetPager: newTracker[azfake.PagerResponder[armresourcemover.UnresolvedDependenciesClientGetResponse]](),
	}
}

// UnresolvedDependenciesServerTransport connects instances of armresourcemover.UnresolvedDependenciesClient to instances of UnresolvedDependenciesServer.
// Don't use this type directly, use NewUnresolvedDependenciesServerTransport instead.
type UnresolvedDependenciesServerTransport struct {
	srv         *UnresolvedDependenciesServer
	newGetPager *tracker[azfake.PagerResponder[armresourcemover.UnresolvedDependenciesClientGetResponse]]
}

// Do implements the policy.Transporter interface for UnresolvedDependenciesServerTransport.
func (u *UnresolvedDependenciesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return u.dispatchToMethodFake(req, method)
}

func (u *UnresolvedDependenciesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if unresolvedDependenciesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = unresolvedDependenciesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "UnresolvedDependenciesClient.NewGetPager":
				res.resp, res.err = u.dispatchNewGetPager(req)
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

func (u *UnresolvedDependenciesServerTransport) dispatchNewGetPager(req *http.Request) (*http.Response, error) {
	if u.srv.NewGetPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetPager not implemented")}
	}
	newGetPager := u.newGetPager.get(req)
	if newGetPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/moveCollections/(?P<moveCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/unresolvedDependencies`
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
		moveCollectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("moveCollectionName")])
		if err != nil {
			return nil, err
		}
		dependencyLevelUnescaped, err := url.QueryUnescape(qp.Get("dependencyLevel"))
		if err != nil {
			return nil, err
		}
		dependencyLevelParam := getOptional(armresourcemover.DependencyLevel(dependencyLevelUnescaped))
		orderbyUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
		if err != nil {
			return nil, err
		}
		orderbyParam := getOptional(orderbyUnescaped)
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armresourcemover.UnresolvedDependenciesClientGetOptions
		if dependencyLevelParam != nil || orderbyParam != nil || filterParam != nil {
			options = &armresourcemover.UnresolvedDependenciesClientGetOptions{
				DependencyLevel: dependencyLevelParam,
				Orderby:         orderbyParam,
				Filter:          filterParam,
			}
		}
		resp := u.srv.NewGetPager(resourceGroupNameParam, moveCollectionNameParam, options)
		newGetPager = &resp
		u.newGetPager.add(req, newGetPager)
		server.PagerResponderInjectNextLinks(newGetPager, req, func(page *armresourcemover.UnresolvedDependenciesClientGetResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newGetPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		u.newGetPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newGetPager) {
		u.newGetPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to UnresolvedDependenciesServerTransport
var unresolvedDependenciesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

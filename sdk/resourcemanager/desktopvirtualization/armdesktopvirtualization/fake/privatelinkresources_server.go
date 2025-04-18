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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// PrivateLinkResourcesServer is a fake server for instances of the armdesktopvirtualization.PrivateLinkResourcesClient type.
type PrivateLinkResourcesServer struct {
	// NewListByHostPoolPager is the fake for method PrivateLinkResourcesClient.NewListByHostPoolPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByHostPoolPager func(resourceGroupName string, hostPoolName string, options *armdesktopvirtualization.PrivateLinkResourcesClientListByHostPoolOptions) (resp azfake.PagerResponder[armdesktopvirtualization.PrivateLinkResourcesClientListByHostPoolResponse])

	// NewListByWorkspacePager is the fake for method PrivateLinkResourcesClient.NewListByWorkspacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByWorkspacePager func(resourceGroupName string, workspaceName string, options *armdesktopvirtualization.PrivateLinkResourcesClientListByWorkspaceOptions) (resp azfake.PagerResponder[armdesktopvirtualization.PrivateLinkResourcesClientListByWorkspaceResponse])
}

// NewPrivateLinkResourcesServerTransport creates a new instance of PrivateLinkResourcesServerTransport with the provided implementation.
// The returned PrivateLinkResourcesServerTransport instance is connected to an instance of armdesktopvirtualization.PrivateLinkResourcesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateLinkResourcesServerTransport(srv *PrivateLinkResourcesServer) *PrivateLinkResourcesServerTransport {
	return &PrivateLinkResourcesServerTransport{
		srv:                     srv,
		newListByHostPoolPager:  newTracker[azfake.PagerResponder[armdesktopvirtualization.PrivateLinkResourcesClientListByHostPoolResponse]](),
		newListByWorkspacePager: newTracker[azfake.PagerResponder[armdesktopvirtualization.PrivateLinkResourcesClientListByWorkspaceResponse]](),
	}
}

// PrivateLinkResourcesServerTransport connects instances of armdesktopvirtualization.PrivateLinkResourcesClient to instances of PrivateLinkResourcesServer.
// Don't use this type directly, use NewPrivateLinkResourcesServerTransport instead.
type PrivateLinkResourcesServerTransport struct {
	srv                     *PrivateLinkResourcesServer
	newListByHostPoolPager  *tracker[azfake.PagerResponder[armdesktopvirtualization.PrivateLinkResourcesClientListByHostPoolResponse]]
	newListByWorkspacePager *tracker[azfake.PagerResponder[armdesktopvirtualization.PrivateLinkResourcesClientListByWorkspaceResponse]]
}

// Do implements the policy.Transporter interface for PrivateLinkResourcesServerTransport.
func (p *PrivateLinkResourcesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PrivateLinkResourcesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if privateLinkResourcesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = privateLinkResourcesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "PrivateLinkResourcesClient.NewListByHostPoolPager":
				res.resp, res.err = p.dispatchNewListByHostPoolPager(req)
			case "PrivateLinkResourcesClient.NewListByWorkspacePager":
				res.resp, res.err = p.dispatchNewListByWorkspacePager(req)
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

func (p *PrivateLinkResourcesServerTransport) dispatchNewListByHostPoolPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByHostPoolPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByHostPoolPager not implemented")}
	}
	newListByHostPoolPager := p.newListByHostPoolPager.get(req)
	if newListByHostPoolPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources`
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
		hostPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hostPoolName")])
		if err != nil {
			return nil, err
		}
		pageSizeUnescaped, err := url.QueryUnescape(qp.Get("pageSize"))
		if err != nil {
			return nil, err
		}
		pageSizeParam, err := parseOptional(pageSizeUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		isDescendingUnescaped, err := url.QueryUnescape(qp.Get("isDescending"))
		if err != nil {
			return nil, err
		}
		isDescendingParam, err := parseOptional(isDescendingUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		initialSkipUnescaped, err := url.QueryUnescape(qp.Get("initialSkip"))
		if err != nil {
			return nil, err
		}
		initialSkipParam, err := parseOptional(initialSkipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armdesktopvirtualization.PrivateLinkResourcesClientListByHostPoolOptions
		if pageSizeParam != nil || isDescendingParam != nil || initialSkipParam != nil {
			options = &armdesktopvirtualization.PrivateLinkResourcesClientListByHostPoolOptions{
				PageSize:     pageSizeParam,
				IsDescending: isDescendingParam,
				InitialSkip:  initialSkipParam,
			}
		}
		resp := p.srv.NewListByHostPoolPager(resourceGroupNameParam, hostPoolNameParam, options)
		newListByHostPoolPager = &resp
		p.newListByHostPoolPager.add(req, newListByHostPoolPager)
		server.PagerResponderInjectNextLinks(newListByHostPoolPager, req, func(page *armdesktopvirtualization.PrivateLinkResourcesClientListByHostPoolResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByHostPoolPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByHostPoolPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByHostPoolPager) {
		p.newListByHostPoolPager.remove(req)
	}
	return resp, nil
}

func (p *PrivateLinkResourcesServerTransport) dispatchNewListByWorkspacePager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByWorkspacePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByWorkspacePager not implemented")}
	}
	newListByWorkspacePager := p.newListByWorkspacePager.get(req)
	if newListByWorkspacePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources`
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
		pageSizeUnescaped, err := url.QueryUnescape(qp.Get("pageSize"))
		if err != nil {
			return nil, err
		}
		pageSizeParam, err := parseOptional(pageSizeUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		isDescendingUnescaped, err := url.QueryUnescape(qp.Get("isDescending"))
		if err != nil {
			return nil, err
		}
		isDescendingParam, err := parseOptional(isDescendingUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		initialSkipUnescaped, err := url.QueryUnescape(qp.Get("initialSkip"))
		if err != nil {
			return nil, err
		}
		initialSkipParam, err := parseOptional(initialSkipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armdesktopvirtualization.PrivateLinkResourcesClientListByWorkspaceOptions
		if pageSizeParam != nil || isDescendingParam != nil || initialSkipParam != nil {
			options = &armdesktopvirtualization.PrivateLinkResourcesClientListByWorkspaceOptions{
				PageSize:     pageSizeParam,
				IsDescending: isDescendingParam,
				InitialSkip:  initialSkipParam,
			}
		}
		resp := p.srv.NewListByWorkspacePager(resourceGroupNameParam, workspaceNameParam, options)
		newListByWorkspacePager = &resp
		p.newListByWorkspacePager.add(req, newListByWorkspacePager)
		server.PagerResponderInjectNextLinks(newListByWorkspacePager, req, func(page *armdesktopvirtualization.PrivateLinkResourcesClientListByWorkspaceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByWorkspacePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByWorkspacePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByWorkspacePager) {
		p.newListByWorkspacePager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to PrivateLinkResourcesServerTransport
var privateLinkResourcesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

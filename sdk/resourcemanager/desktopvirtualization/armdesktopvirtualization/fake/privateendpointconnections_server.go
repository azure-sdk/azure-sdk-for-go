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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v3"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// PrivateEndpointConnectionsServer is a fake server for instances of the armdesktopvirtualization.PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsServer struct {
	// DeleteByHostPool is the fake for method PrivateEndpointConnectionsClient.DeleteByHostPool
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	DeleteByHostPool func(ctx context.Context, resourceGroupName string, hostPoolName string, privateEndpointConnectionName string, options *armdesktopvirtualization.PrivateEndpointConnectionsClientDeleteByHostPoolOptions) (resp azfake.Responder[armdesktopvirtualization.PrivateEndpointConnectionsClientDeleteByHostPoolResponse], errResp azfake.ErrorResponder)

	// DeleteByWorkspace is the fake for method PrivateEndpointConnectionsClient.DeleteByWorkspace
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	DeleteByWorkspace func(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string, options *armdesktopvirtualization.PrivateEndpointConnectionsClientDeleteByWorkspaceOptions) (resp azfake.Responder[armdesktopvirtualization.PrivateEndpointConnectionsClientDeleteByWorkspaceResponse], errResp azfake.ErrorResponder)

	// GetByHostPool is the fake for method PrivateEndpointConnectionsClient.GetByHostPool
	// HTTP status codes to indicate success: http.StatusOK
	GetByHostPool func(ctx context.Context, resourceGroupName string, hostPoolName string, privateEndpointConnectionName string, options *armdesktopvirtualization.PrivateEndpointConnectionsClientGetByHostPoolOptions) (resp azfake.Responder[armdesktopvirtualization.PrivateEndpointConnectionsClientGetByHostPoolResponse], errResp azfake.ErrorResponder)

	// GetByWorkspace is the fake for method PrivateEndpointConnectionsClient.GetByWorkspace
	// HTTP status codes to indicate success: http.StatusOK
	GetByWorkspace func(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string, options *armdesktopvirtualization.PrivateEndpointConnectionsClientGetByWorkspaceOptions) (resp azfake.Responder[armdesktopvirtualization.PrivateEndpointConnectionsClientGetByWorkspaceResponse], errResp azfake.ErrorResponder)

	// NewListByHostPoolPager is the fake for method PrivateEndpointConnectionsClient.NewListByHostPoolPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByHostPoolPager func(resourceGroupName string, hostPoolName string, options *armdesktopvirtualization.PrivateEndpointConnectionsClientListByHostPoolOptions) (resp azfake.PagerResponder[armdesktopvirtualization.PrivateEndpointConnectionsClientListByHostPoolResponse])

	// NewListByWorkspacePager is the fake for method PrivateEndpointConnectionsClient.NewListByWorkspacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByWorkspacePager func(resourceGroupName string, workspaceName string, options *armdesktopvirtualization.PrivateEndpointConnectionsClientListByWorkspaceOptions) (resp azfake.PagerResponder[armdesktopvirtualization.PrivateEndpointConnectionsClientListByWorkspaceResponse])

	// UpdateByHostPool is the fake for method PrivateEndpointConnectionsClient.UpdateByHostPool
	// HTTP status codes to indicate success: http.StatusOK
	UpdateByHostPool func(ctx context.Context, resourceGroupName string, hostPoolName string, privateEndpointConnectionName string, connection armdesktopvirtualization.PrivateEndpointConnection, options *armdesktopvirtualization.PrivateEndpointConnectionsClientUpdateByHostPoolOptions) (resp azfake.Responder[armdesktopvirtualization.PrivateEndpointConnectionsClientUpdateByHostPoolResponse], errResp azfake.ErrorResponder)

	// UpdateByWorkspace is the fake for method PrivateEndpointConnectionsClient.UpdateByWorkspace
	// HTTP status codes to indicate success: http.StatusOK
	UpdateByWorkspace func(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string, connection armdesktopvirtualization.PrivateEndpointConnection, options *armdesktopvirtualization.PrivateEndpointConnectionsClientUpdateByWorkspaceOptions) (resp azfake.Responder[armdesktopvirtualization.PrivateEndpointConnectionsClientUpdateByWorkspaceResponse], errResp azfake.ErrorResponder)
}

// NewPrivateEndpointConnectionsServerTransport creates a new instance of PrivateEndpointConnectionsServerTransport with the provided implementation.
// The returned PrivateEndpointConnectionsServerTransport instance is connected to an instance of armdesktopvirtualization.PrivateEndpointConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateEndpointConnectionsServerTransport(srv *PrivateEndpointConnectionsServer) *PrivateEndpointConnectionsServerTransport {
	return &PrivateEndpointConnectionsServerTransport{
		srv:                     srv,
		newListByHostPoolPager:  newTracker[azfake.PagerResponder[armdesktopvirtualization.PrivateEndpointConnectionsClientListByHostPoolResponse]](),
		newListByWorkspacePager: newTracker[azfake.PagerResponder[armdesktopvirtualization.PrivateEndpointConnectionsClientListByWorkspaceResponse]](),
	}
}

// PrivateEndpointConnectionsServerTransport connects instances of armdesktopvirtualization.PrivateEndpointConnectionsClient to instances of PrivateEndpointConnectionsServer.
// Don't use this type directly, use NewPrivateEndpointConnectionsServerTransport instead.
type PrivateEndpointConnectionsServerTransport struct {
	srv                     *PrivateEndpointConnectionsServer
	newListByHostPoolPager  *tracker[azfake.PagerResponder[armdesktopvirtualization.PrivateEndpointConnectionsClientListByHostPoolResponse]]
	newListByWorkspacePager *tracker[azfake.PagerResponder[armdesktopvirtualization.PrivateEndpointConnectionsClientListByWorkspaceResponse]]
}

// Do implements the policy.Transporter interface for PrivateEndpointConnectionsServerTransport.
func (p *PrivateEndpointConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PrivateEndpointConnectionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if privateEndpointConnectionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = privateEndpointConnectionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "PrivateEndpointConnectionsClient.DeleteByHostPool":
				res.resp, res.err = p.dispatchDeleteByHostPool(req)
			case "PrivateEndpointConnectionsClient.DeleteByWorkspace":
				res.resp, res.err = p.dispatchDeleteByWorkspace(req)
			case "PrivateEndpointConnectionsClient.GetByHostPool":
				res.resp, res.err = p.dispatchGetByHostPool(req)
			case "PrivateEndpointConnectionsClient.GetByWorkspace":
				res.resp, res.err = p.dispatchGetByWorkspace(req)
			case "PrivateEndpointConnectionsClient.NewListByHostPoolPager":
				res.resp, res.err = p.dispatchNewListByHostPoolPager(req)
			case "PrivateEndpointConnectionsClient.NewListByWorkspacePager":
				res.resp, res.err = p.dispatchNewListByWorkspacePager(req)
			case "PrivateEndpointConnectionsClient.UpdateByHostPool":
				res.resp, res.err = p.dispatchUpdateByHostPool(req)
			case "PrivateEndpointConnectionsClient.UpdateByWorkspace":
				res.resp, res.err = p.dispatchUpdateByWorkspace(req)
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

func (p *PrivateEndpointConnectionsServerTransport) dispatchDeleteByHostPool(req *http.Request) (*http.Response, error) {
	if p.srv.DeleteByHostPool == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteByHostPool not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hostPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hostPoolName")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.DeleteByHostPool(req.Context(), resourceGroupNameParam, hostPoolNameParam, privateEndpointConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionsServerTransport) dispatchDeleteByWorkspace(req *http.Request) (*http.Response, error) {
	if p.srv.DeleteByWorkspace == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteByWorkspace not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.DeleteByWorkspace(req.Context(), resourceGroupNameParam, workspaceNameParam, privateEndpointConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionsServerTransport) dispatchGetByHostPool(req *http.Request) (*http.Response, error) {
	if p.srv.GetByHostPool == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByHostPool not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hostPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hostPoolName")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.GetByHostPool(req.Context(), resourceGroupNameParam, hostPoolNameParam, privateEndpointConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateEndpointConnectionWithSystemData, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionsServerTransport) dispatchGetByWorkspace(req *http.Request) (*http.Response, error) {
	if p.srv.GetByWorkspace == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByWorkspace not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.GetByWorkspace(req.Context(), resourceGroupNameParam, workspaceNameParam, privateEndpointConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateEndpointConnectionWithSystemData, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionsServerTransport) dispatchNewListByHostPoolPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByHostPoolPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByHostPoolPager not implemented")}
	}
	newListByHostPoolPager := p.newListByHostPoolPager.get(req)
	if newListByHostPoolPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections`
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
		var options *armdesktopvirtualization.PrivateEndpointConnectionsClientListByHostPoolOptions
		if pageSizeParam != nil || isDescendingParam != nil || initialSkipParam != nil {
			options = &armdesktopvirtualization.PrivateEndpointConnectionsClientListByHostPoolOptions{
				PageSize:     pageSizeParam,
				IsDescending: isDescendingParam,
				InitialSkip:  initialSkipParam,
			}
		}
		resp := p.srv.NewListByHostPoolPager(resourceGroupNameParam, hostPoolNameParam, options)
		newListByHostPoolPager = &resp
		p.newListByHostPoolPager.add(req, newListByHostPoolPager)
		server.PagerResponderInjectNextLinks(newListByHostPoolPager, req, func(page *armdesktopvirtualization.PrivateEndpointConnectionsClientListByHostPoolResponse, createLink func() string) {
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

func (p *PrivateEndpointConnectionsServerTransport) dispatchNewListByWorkspacePager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByWorkspacePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByWorkspacePager not implemented")}
	}
	newListByWorkspacePager := p.newListByWorkspacePager.get(req)
	if newListByWorkspacePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListByWorkspacePager(resourceGroupNameParam, workspaceNameParam, nil)
		newListByWorkspacePager = &resp
		p.newListByWorkspacePager.add(req, newListByWorkspacePager)
		server.PagerResponderInjectNextLinks(newListByWorkspacePager, req, func(page *armdesktopvirtualization.PrivateEndpointConnectionsClientListByWorkspaceResponse, createLink func() string) {
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

func (p *PrivateEndpointConnectionsServerTransport) dispatchUpdateByHostPool(req *http.Request) (*http.Response, error) {
	if p.srv.UpdateByHostPool == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateByHostPool not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdesktopvirtualization.PrivateEndpointConnection](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hostPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hostPoolName")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.UpdateByHostPool(req.Context(), resourceGroupNameParam, hostPoolNameParam, privateEndpointConnectionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateEndpointConnectionWithSystemData, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionsServerTransport) dispatchUpdateByWorkspace(req *http.Request) (*http.Response, error) {
	if p.srv.UpdateByWorkspace == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateByWorkspace not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdesktopvirtualization.PrivateEndpointConnection](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.UpdateByWorkspace(req.Context(), resourceGroupNameParam, workspaceNameParam, privateEndpointConnectionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateEndpointConnectionWithSystemData, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to PrivateEndpointConnectionsServerTransport
var privateEndpointConnectionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

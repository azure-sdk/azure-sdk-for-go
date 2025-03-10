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
	"reflect"
	"regexp"
	"strconv"
)

// SessionHostsServer is a fake server for instances of the armdesktopvirtualization.SessionHostsClient type.
type SessionHostsServer struct {
	// Delete is the fake for method SessionHostsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, options *armdesktopvirtualization.SessionHostsClientDeleteOptions) (resp azfake.Responder[armdesktopvirtualization.SessionHostsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SessionHostsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, options *armdesktopvirtualization.SessionHostsClientGetOptions) (resp azfake.Responder[armdesktopvirtualization.SessionHostsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method SessionHostsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, hostPoolName string, options *armdesktopvirtualization.SessionHostsClientListOptions) (resp azfake.PagerResponder[armdesktopvirtualization.SessionHostsClientListResponse])

	// RetryProvisioning is the fake for method SessionHostsClient.RetryProvisioning
	// HTTP status codes to indicate success: http.StatusNoContent
	RetryProvisioning func(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, options *armdesktopvirtualization.SessionHostsClientRetryProvisioningOptions) (resp azfake.Responder[armdesktopvirtualization.SessionHostsClientRetryProvisioningResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method SessionHostsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, options *armdesktopvirtualization.SessionHostsClientUpdateOptions) (resp azfake.Responder[armdesktopvirtualization.SessionHostsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewSessionHostsServerTransport creates a new instance of SessionHostsServerTransport with the provided implementation.
// The returned SessionHostsServerTransport instance is connected to an instance of armdesktopvirtualization.SessionHostsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSessionHostsServerTransport(srv *SessionHostsServer) *SessionHostsServerTransport {
	return &SessionHostsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armdesktopvirtualization.SessionHostsClientListResponse]](),
	}
}

// SessionHostsServerTransport connects instances of armdesktopvirtualization.SessionHostsClient to instances of SessionHostsServer.
// Don't use this type directly, use NewSessionHostsServerTransport instead.
type SessionHostsServerTransport struct {
	srv          *SessionHostsServer
	newListPager *tracker[azfake.PagerResponder[armdesktopvirtualization.SessionHostsClientListResponse]]
}

// Do implements the policy.Transporter interface for SessionHostsServerTransport.
func (s *SessionHostsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SessionHostsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if sessionHostsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = sessionHostsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "SessionHostsClient.Delete":
				res.resp, res.err = s.dispatchDelete(req)
			case "SessionHostsClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "SessionHostsClient.NewListPager":
				res.resp, res.err = s.dispatchNewListPager(req)
			case "SessionHostsClient.RetryProvisioning":
				res.resp, res.err = s.dispatchRetryProvisioning(req)
			case "SessionHostsClient.Update":
				res.resp, res.err = s.dispatchUpdate(req)
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

func (s *SessionHostsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if s.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sessionHosts/(?P<sessionHostName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
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
	sessionHostNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sessionHostName")])
	if err != nil {
		return nil, err
	}
	forceUnescaped, err := url.QueryUnescape(qp.Get("force"))
	if err != nil {
		return nil, err
	}
	forceParam, err := parseOptional(forceUnescaped, strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	var options *armdesktopvirtualization.SessionHostsClientDeleteOptions
	if forceParam != nil {
		options = &armdesktopvirtualization.SessionHostsClientDeleteOptions{
			Force: forceParam,
		}
	}
	respr, errRespr := s.srv.Delete(req.Context(), resourceGroupNameParam, hostPoolNameParam, sessionHostNameParam, options)
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

func (s *SessionHostsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sessionHosts/(?P<sessionHostName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	sessionHostNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sessionHostName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, hostPoolNameParam, sessionHostNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SessionHost, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SessionHostsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sessionHosts`
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
		vMPathUnescaped, err := url.QueryUnescape(qp.Get("vmPath"))
		if err != nil {
			return nil, err
		}
		vMPathParam := getOptional(vMPathUnescaped)
		var options *armdesktopvirtualization.SessionHostsClientListOptions
		if pageSizeParam != nil || isDescendingParam != nil || initialSkipParam != nil || vMPathParam != nil {
			options = &armdesktopvirtualization.SessionHostsClientListOptions{
				PageSize:     pageSizeParam,
				IsDescending: isDescendingParam,
				InitialSkip:  initialSkipParam,
				VMPath:       vMPathParam,
			}
		}
		resp := s.srv.NewListPager(resourceGroupNameParam, hostPoolNameParam, options)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armdesktopvirtualization.SessionHostsClientListResponse, createLink func() string) {
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

func (s *SessionHostsServerTransport) dispatchRetryProvisioning(req *http.Request) (*http.Response, error) {
	if s.srv.RetryProvisioning == nil {
		return nil, &nonRetriableError{errors.New("fake for method RetryProvisioning not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sessionHosts/(?P<sessionHostName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/retryProvisioning`
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
	sessionHostNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sessionHostName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.RetryProvisioning(req.Context(), resourceGroupNameParam, hostPoolNameParam, sessionHostNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SessionHostsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sessionHosts/(?P<sessionHostName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[armdesktopvirtualization.SessionHostPatch](req)
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
	sessionHostNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sessionHostName")])
	if err != nil {
		return nil, err
	}
	forceUnescaped, err := url.QueryUnescape(qp.Get("force"))
	if err != nil {
		return nil, err
	}
	forceParam, err := parseOptional(forceUnescaped, strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	var options *armdesktopvirtualization.SessionHostsClientUpdateOptions
	if forceParam != nil || !reflect.ValueOf(body).IsZero() {
		options = &armdesktopvirtualization.SessionHostsClientUpdateOptions{
			Force:       forceParam,
			SessionHost: &body,
		}
	}
	respr, errRespr := s.srv.Update(req.Context(), resourceGroupNameParam, hostPoolNameParam, sessionHostNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SessionHost, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to SessionHostsServerTransport
var sessionHostsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

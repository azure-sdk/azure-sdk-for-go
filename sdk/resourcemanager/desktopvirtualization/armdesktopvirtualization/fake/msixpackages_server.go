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

// MSIXPackagesServer is a fake server for instances of the armdesktopvirtualization.MSIXPackagesClient type.
type MSIXPackagesServer struct {
	// CreateOrUpdate is the fake for method MSIXPackagesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, hostPoolName string, msixPackageFullName string, msixPackage armdesktopvirtualization.MSIXPackage, options *armdesktopvirtualization.MSIXPackagesClientCreateOrUpdateOptions) (resp azfake.Responder[armdesktopvirtualization.MSIXPackagesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method MSIXPackagesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, hostPoolName string, msixPackageFullName string, options *armdesktopvirtualization.MSIXPackagesClientDeleteOptions) (resp azfake.Responder[armdesktopvirtualization.MSIXPackagesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method MSIXPackagesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, hostPoolName string, msixPackageFullName string, options *armdesktopvirtualization.MSIXPackagesClientGetOptions) (resp azfake.Responder[armdesktopvirtualization.MSIXPackagesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method MSIXPackagesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, hostPoolName string, options *armdesktopvirtualization.MSIXPackagesClientListOptions) (resp azfake.PagerResponder[armdesktopvirtualization.MSIXPackagesClientListResponse])

	// Update is the fake for method MSIXPackagesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, hostPoolName string, msixPackageFullName string, msixPackage armdesktopvirtualization.MSIXPackagePatch, options *armdesktopvirtualization.MSIXPackagesClientUpdateOptions) (resp azfake.Responder[armdesktopvirtualization.MSIXPackagesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewMSIXPackagesServerTransport creates a new instance of MSIXPackagesServerTransport with the provided implementation.
// The returned MSIXPackagesServerTransport instance is connected to an instance of armdesktopvirtualization.MSIXPackagesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewMSIXPackagesServerTransport(srv *MSIXPackagesServer) *MSIXPackagesServerTransport {
	return &MSIXPackagesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armdesktopvirtualization.MSIXPackagesClientListResponse]](),
	}
}

// MSIXPackagesServerTransport connects instances of armdesktopvirtualization.MSIXPackagesClient to instances of MSIXPackagesServer.
// Don't use this type directly, use NewMSIXPackagesServerTransport instead.
type MSIXPackagesServerTransport struct {
	srv          *MSIXPackagesServer
	newListPager *tracker[azfake.PagerResponder[armdesktopvirtualization.MSIXPackagesClientListResponse]]
}

// Do implements the policy.Transporter interface for MSIXPackagesServerTransport.
func (m *MSIXPackagesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return m.dispatchToMethodFake(req, method)
}

func (m *MSIXPackagesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if msixPackagesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = msixPackagesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "MSIXPackagesClient.CreateOrUpdate":
				res.resp, res.err = m.dispatchCreateOrUpdate(req)
			case "MSIXPackagesClient.Delete":
				res.resp, res.err = m.dispatchDelete(req)
			case "MSIXPackagesClient.Get":
				res.resp, res.err = m.dispatchGet(req)
			case "MSIXPackagesClient.NewListPager":
				res.resp, res.err = m.dispatchNewListPager(req)
			case "MSIXPackagesClient.Update":
				res.resp, res.err = m.dispatchUpdate(req)
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

func (m *MSIXPackagesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/msixPackages/(?P<msixPackageFullName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdesktopvirtualization.MSIXPackage](req)
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
	msixPackageFullNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("msixPackageFullName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, hostPoolNameParam, msixPackageFullNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).MSIXPackage, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MSIXPackagesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if m.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/msixPackages/(?P<msixPackageFullName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	msixPackageFullNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("msixPackageFullName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Delete(req.Context(), resourceGroupNameParam, hostPoolNameParam, msixPackageFullNameParam, nil)
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

func (m *MSIXPackagesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/msixPackages/(?P<msixPackageFullName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	msixPackageFullNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("msixPackageFullName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, hostPoolNameParam, msixPackageFullNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).MSIXPackage, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MSIXPackagesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := m.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/msixPackages`
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
		var options *armdesktopvirtualization.MSIXPackagesClientListOptions
		if pageSizeParam != nil || isDescendingParam != nil || initialSkipParam != nil {
			options = &armdesktopvirtualization.MSIXPackagesClientListOptions{
				PageSize:     pageSizeParam,
				IsDescending: isDescendingParam,
				InitialSkip:  initialSkipParam,
			}
		}
		resp := m.srv.NewListPager(resourceGroupNameParam, hostPoolNameParam, options)
		newListPager = &resp
		m.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armdesktopvirtualization.MSIXPackagesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		m.newListPager.remove(req)
	}
	return resp, nil
}

func (m *MSIXPackagesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DesktopVirtualization/hostPools/(?P<hostPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/msixPackages/(?P<msixPackageFullName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdesktopvirtualization.MSIXPackagePatch](req)
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
	msixPackageFullNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("msixPackageFullName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Update(req.Context(), resourceGroupNameParam, hostPoolNameParam, msixPackageFullNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).MSIXPackage, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to MSIXPackagesServerTransport
var msixPackagesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

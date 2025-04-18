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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// UsersServer is a fake server for instances of the armdevtestlabs.UsersClient type.
type UsersServer struct {
	// BeginCreateOrUpdate is the fake for method UsersClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, labName string, name string, userParam armdevtestlabs.User, options *armdevtestlabs.UsersClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armdevtestlabs.UsersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method UsersClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, labName string, name string, options *armdevtestlabs.UsersClientBeginDeleteOptions) (resp azfake.PollerResponder[armdevtestlabs.UsersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method UsersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, labName string, name string, options *armdevtestlabs.UsersClientGetOptions) (resp azfake.Responder[armdevtestlabs.UsersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method UsersClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, labName string, options *armdevtestlabs.UsersClientListOptions) (resp azfake.PagerResponder[armdevtestlabs.UsersClientListResponse])

	// Update is the fake for method UsersClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, labName string, name string, userParam armdevtestlabs.UserFragment, options *armdevtestlabs.UsersClientUpdateOptions) (resp azfake.Responder[armdevtestlabs.UsersClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewUsersServerTransport creates a new instance of UsersServerTransport with the provided implementation.
// The returned UsersServerTransport instance is connected to an instance of armdevtestlabs.UsersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewUsersServerTransport(srv *UsersServer) *UsersServerTransport {
	return &UsersServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armdevtestlabs.UsersClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armdevtestlabs.UsersClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armdevtestlabs.UsersClientListResponse]](),
	}
}

// UsersServerTransport connects instances of armdevtestlabs.UsersClient to instances of UsersServer.
// Don't use this type directly, use NewUsersServerTransport instead.
type UsersServerTransport struct {
	srv                 *UsersServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armdevtestlabs.UsersClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armdevtestlabs.UsersClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armdevtestlabs.UsersClientListResponse]]
}

// Do implements the policy.Transporter interface for UsersServerTransport.
func (u *UsersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return u.dispatchToMethodFake(req, method)
}

func (u *UsersServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if usersServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = usersServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "UsersClient.BeginCreateOrUpdate":
				res.resp, res.err = u.dispatchBeginCreateOrUpdate(req)
			case "UsersClient.BeginDelete":
				res.resp, res.err = u.dispatchBeginDelete(req)
			case "UsersClient.Get":
				res.resp, res.err = u.dispatchGet(req)
			case "UsersClient.NewListPager":
				res.resp, res.err = u.dispatchNewListPager(req)
			case "UsersClient.Update":
				res.resp, res.err = u.dispatchUpdate(req)
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

func (u *UsersServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if u.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := u.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdevtestlabs.User](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := u.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, labNameParam, nameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		u.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		u.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		u.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (u *UsersServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if u.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := u.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := u.srv.BeginDelete(req.Context(), resourceGroupNameParam, labNameParam, nameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		u.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		u.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		u.beginDelete.remove(req)
	}

	return resp, nil
}

func (u *UsersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if u.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armdevtestlabs.UsersClientGetOptions
	if expandParam != nil {
		options = &armdevtestlabs.UsersClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := u.srv.Get(req.Context(), resourceGroupNameParam, labNameParam, nameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).User, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *UsersServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if u.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := u.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users`
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
		labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
		if err != nil {
			return nil, err
		}
		expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
		if err != nil {
			return nil, err
		}
		expandParam := getOptional(expandUnescaped)
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
		orderbyUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
		if err != nil {
			return nil, err
		}
		orderbyParam := getOptional(orderbyUnescaped)
		var options *armdevtestlabs.UsersClientListOptions
		if expandParam != nil || filterParam != nil || topParam != nil || orderbyParam != nil {
			options = &armdevtestlabs.UsersClientListOptions{
				Expand:  expandParam,
				Filter:  filterParam,
				Top:     topParam,
				Orderby: orderbyParam,
			}
		}
		resp := u.srv.NewListPager(resourceGroupNameParam, labNameParam, options)
		newListPager = &resp
		u.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armdevtestlabs.UsersClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		u.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		u.newListPager.remove(req)
	}
	return resp, nil
}

func (u *UsersServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if u.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdevtestlabs.UserFragment](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := u.srv.Update(req.Context(), resourceGroupNameParam, labNameParam, nameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).User, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to UsersServerTransport
var usersServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}

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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
	"net/http"
	"net/url"
	"regexp"
)

// UsersServer is a fake server for instances of the armdataboxedge.UsersClient type.
type UsersServer struct {
	// BeginCreateOrUpdate is the fake for method UsersClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, deviceName string, name string, resourceGroupName string, userParam armdataboxedge.User, options *armdataboxedge.UsersClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armdataboxedge.UsersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method UsersClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, deviceName string, name string, resourceGroupName string, options *armdataboxedge.UsersClientBeginDeleteOptions) (resp azfake.PollerResponder[armdataboxedge.UsersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method UsersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, deviceName string, name string, resourceGroupName string, options *armdataboxedge.UsersClientGetOptions) (resp azfake.Responder[armdataboxedge.UsersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByDataBoxEdgeDevicePager is the fake for method UsersClient.NewListByDataBoxEdgeDevicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDataBoxEdgeDevicePager func(deviceName string, resourceGroupName string, options *armdataboxedge.UsersClientListByDataBoxEdgeDeviceOptions) (resp azfake.PagerResponder[armdataboxedge.UsersClientListByDataBoxEdgeDeviceResponse])
}

// NewUsersServerTransport creates a new instance of UsersServerTransport with the provided implementation.
// The returned UsersServerTransport instance is connected to an instance of armdataboxedge.UsersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewUsersServerTransport(srv *UsersServer) *UsersServerTransport {
	return &UsersServerTransport{
		srv:                             srv,
		beginCreateOrUpdate:             newTracker[azfake.PollerResponder[armdataboxedge.UsersClientCreateOrUpdateResponse]](),
		beginDelete:                     newTracker[azfake.PollerResponder[armdataboxedge.UsersClientDeleteResponse]](),
		newListByDataBoxEdgeDevicePager: newTracker[azfake.PagerResponder[armdataboxedge.UsersClientListByDataBoxEdgeDeviceResponse]](),
	}
}

// UsersServerTransport connects instances of armdataboxedge.UsersClient to instances of UsersServer.
// Don't use this type directly, use NewUsersServerTransport instead.
type UsersServerTransport struct {
	srv                             *UsersServer
	beginCreateOrUpdate             *tracker[azfake.PollerResponder[armdataboxedge.UsersClientCreateOrUpdateResponse]]
	beginDelete                     *tracker[azfake.PollerResponder[armdataboxedge.UsersClientDeleteResponse]]
	newListByDataBoxEdgeDevicePager *tracker[azfake.PagerResponder[armdataboxedge.UsersClientListByDataBoxEdgeDeviceResponse]]
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
			case "UsersClient.NewListByDataBoxEdgeDevicePager":
				res.resp, res.err = u.dispatchNewListByDataBoxEdgeDevicePager(req)
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
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdataboxedge.User](req)
		if err != nil {
			return nil, err
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := u.srv.BeginCreateOrUpdate(req.Context(), deviceNameParam, nameParam, resourceGroupNameParam, body, nil)
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

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		u.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
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
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := u.srv.BeginDelete(req.Context(), deviceNameParam, nameParam, resourceGroupNameParam, nil)
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
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := u.srv.Get(req.Context(), deviceNameParam, nameParam, resourceGroupNameParam, nil)
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

func (u *UsersServerTransport) dispatchNewListByDataBoxEdgeDevicePager(req *http.Request) (*http.Response, error) {
	if u.srv.NewListByDataBoxEdgeDevicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByDataBoxEdgeDevicePager not implemented")}
	}
	newListByDataBoxEdgeDevicePager := u.newListByDataBoxEdgeDevicePager.get(req)
	if newListByDataBoxEdgeDevicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armdataboxedge.UsersClientListByDataBoxEdgeDeviceOptions
		if filterParam != nil {
			options = &armdataboxedge.UsersClientListByDataBoxEdgeDeviceOptions{
				Filter: filterParam,
			}
		}
		resp := u.srv.NewListByDataBoxEdgeDevicePager(deviceNameParam, resourceGroupNameParam, options)
		newListByDataBoxEdgeDevicePager = &resp
		u.newListByDataBoxEdgeDevicePager.add(req, newListByDataBoxEdgeDevicePager)
		server.PagerResponderInjectNextLinks(newListByDataBoxEdgeDevicePager, req, func(page *armdataboxedge.UsersClientListByDataBoxEdgeDeviceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByDataBoxEdgeDevicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		u.newListByDataBoxEdgeDevicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByDataBoxEdgeDevicePager) {
		u.newListByDataBoxEdgeDevicePager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to UsersServerTransport
var usersServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
